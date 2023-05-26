package services

import (
	"book/config"
	"book/models"
	"book/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// BorrowCreate
// @Tags 借书方法
// @Summary 新建借书
// @Param book_id formData string true "book_id"
// @Param user_id formData string true "user_id"
// @Param days formData string true "days"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /borrow_create [post]
func BorrowCreate(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Borrow
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	user_id := data.UserID
	book_id := data.BookID
	data.Books = nil
	data.Users = nil
	data.Note = "外租中"
	data.ReturnDate = time.Now().AddDate(0, 0, data.Days).Format("2006-01-02 15:04:05")
	utils.Log.Infoln("getting data")
	tx := models.DB.Begin()
	err = tx.Model(new(models.Borrow)).Create(&data).Error

	var book models.Book
	err = tx.Where("id=?", book_id).First(&book).Error
	book.Nums = book.Nums - 1
	book.PublishDate = book.PublishDate[:10]
	if book.Nums < 0 {
		utils.Log.Errorln("get data error")
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "图书数量不足",
		})
		return
	}
	var user models.User
	err = tx.Where("id=?", user_id).First(&user).Error
	user.Account = user.Account - book.Score
	if user.Account < 0 {
		utils.Log.Errorln("get data error")
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户余额不足",
		})
		return
	}
	err = tx.Updates(&user).Error
	err = tx.Updates(&book).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "create error" + err.Error(),
		})
		return
	}

	tx.Commit()
	utils.Log.Infoln("return")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "创建成功",
	})

}

// GetBorrowList
// @Tags 借书方法
// @Summary 借书列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param user_name query string false "user_name"
// @Param book_name query string false "book_name"
// @Param book_id query string false "book_id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /borrow_list [get]
func GetBorrowList(c *gin.Context) {
	utils.Log.Infoln("info")
	type _param struct {
		Size     string `form:"size" json:"size"`
		Page     string `form:"page" json:"page"`
		UserName string `form:"user_name" json:"user_name"`
		BookName string `form:"book_name" json:"book_name"`
		BookID   string `form:"book_id" json:"book_id"`
	}
	data := _param{
		Size: config.DefaultSize,
		Page: config.DefaultPage,
	}
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Errorln("bind error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
	}
	list := make([]models.Borrow, 0)
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64

	utils.Log.Infoln("getting data")
	tx := models.GetBorrowList(data.UserName, data.BookName, data.BookID)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "find data error" + err.Error(),
		})
		return
	}

	utils.Log.Infoln("return")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}

// BorrowDelete
// @Tags 借书方法
// @Summary 借书删除
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /borrow_delete [delete]
func BorrowDelete(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Borrow
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Errorln("bind error")
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("getting data")
	err = models.DB.Where("id=?", data.ID).First(&data).Delete(&data).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "图书不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "delete Detail Error:" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("return")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "删除成功",
	})
}
