package services

import (
	"book/config"
	"book/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// ReturCreate
// @Tags 还书方法
// @Summary 还书创建
// @Param book_id formData string true "book_id"
// @Param user_id formData string true "user_id"
// @Param days formData string true "days"
// @Param id formData string true "id"
// @Param return_date formData string true "return_date"
// @Param created_at formData string true "created_at"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/retur_create [post]
func ReturCreate(c *gin.Context) {
	var retur models.Retur
	err := c.ShouldBind(&retur)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	book_id := retur.BookID
	borrow_id := retur.ID //传递当前要还的书的id
	retur.ID = 0
	retur.Books = nil
	retur.Users = nil

	tx := models.DB.Begin()
	var book models.Book
	tx.Where("id=?", book_id).First(&book)
	book.Nums += 1
	tx.Updates(&book)
	tx.Where("id=?", borrow_id).First(new(models.Borrow)).Select("note").Update("note", "已归还")
	tx.Model(new(models.Retur)).Create(&retur)
	err = tx.Commit().Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "图书不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "create retur Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "归还成功",
	})
}

// GetReturList
// @Tags 还书方法
// @Summary 还书列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param user_name query string false "user_name"
// @Param book_name query string false "book_name"
// @Param book_id query string false "book_id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/retur_list [get]
func GetReturList(c *gin.Context) {
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
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
	}
	list := make([]models.Retur, 0)
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64
	tx := models.GetReturList(data.UserName, data.BookName, data.BookID)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "find data error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}

// ReturDelete
// @Tags 还书方法
// @Summary 还书删除
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/retur_delete [delete]
func ReturDelete(c *gin.Context) {
	var data models.Retur
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	err = models.DB.Where("id=?", data.ID).First(&data).Delete(&data).Error
	if err != nil {
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
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "删除成功",
	})
}
