package services

import (
	"book/config"
	"book/models"
	"book/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// CoverPhotoUp
// @Tags 图书方法
// @Summary 头像上传
// @Param token query string true "token"
// @Param file formData string true "file"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /bookimgup [post]
func CoverPhotoUp(c *gin.Context) {
	utils.Log.Infoln("enter")
	token := c.Query("token")
	userClaim, err := utils.AnalyseToken(token)
	if err != nil {
		utils.Log.Errorln("analyse token error")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "未认证:Unauthorized Authorization",
		})
		return
	}
	if userClaim == nil {
		utils.Log.Infoln("unauthor admin")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "用户无权限:Unauthorized Admin",
		})
		return
	}
	utils.Log.Infoln("uping file")
	// 单文件
	file, _ := c.FormFile("file")
	dst := "./imgs/" + utils.GetUUID() + file.Filename
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.Log.Infoln("up file error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "save error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("return")
	c.JSON(200, gin.H{
		"code": "200",
		"data": dst,
	})
}

// BookCreate
// @Tags 图书方法
// @Summary 图书创建
// @Param name formData string true "name"
// @Param description formData string true "description"
// @Param publish_date formData string true "publish_date"
// @Param author formData string true "author"
// @Param publisher formData string true "publisher"
// @Param bookNo formData string true "bookNo"
// @Param score formData string true "score"
// @Param nums formData string true "nums"
// @Param category_id formData string true "category_id"
// @Param cover formData string true "cover"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /book_create [post]
func BookCreate(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Book
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Infoln("bind error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("getting data")
	err = models.DB.Model(new(models.Book)).Create(&data).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "create error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("return")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "创建成功",
	})
}

// GetBookList
// @Tags 图书方法
// @Summary 图书列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param name query string false "name"
// @Param id query string false "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /book_list [get]
func GetBookList(c *gin.Context) {
	utils.Log.Infoln("enter")
	type _param struct {
		Size   string `form:"size" json:"size"`
		Page   string `form:"page" json:"page"`
		Name   string `form:"name" json:"name"`
		BookNo string `form:"bookNo" json:"bookNo"`
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
	list := make([]models.Book, 0)
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64

	utils.Log.Infoln("getting data")
	tx := models.GetBookList(data.Name, data.BookNo)
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

// GetBookDetail
// @Tags 图书方法
// @Summary 图书详细
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /book_detail [get]
func GetBookDetail(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Book
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
	err = models.DB.Where("id=?", data.ID).First(&data).Error
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
			"msg":  "Get Detail Error:" + err.Error(),
		})
		return
	}

	utils.Log.Infoln("return")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})
}

// BookDelete
// @Tags 图书方法
// @Summary 图书删除
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /book_delete [delete]
func BookDelete(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Book
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

// BookUpdate
// @Tags 图书方法
// @Summary 图书修改
// @Param id formData int true "id"
// @Param name formData string false "name"
// @Param description formData string false "description"
// @Param publish_date formData string false "publish_date"
// @Param author formData string false "author"
// @Param publisher formData string false "publisher"
// @Param bookNo formData string false "bookNo"
// @Param score formData string false "score"
// @Param nums formData string false "nums"
// @Param category_id formData string false "category_id"
// @Param cover formData string false "cover"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /book_update [put]
func BookUpdate(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Book
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
	err = models.DB.Model(new(models.Book)).Where("id=?", data.ID).Updates(&data).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "update error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("return")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "修改成功",
	})
}
