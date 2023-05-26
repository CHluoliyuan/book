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

// CategoryCreate
// @Tags 分类方法
// @Summary 添加分类
// @Param name formData string true "name"
// @Param remark formData string false "remark"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /category_create [post]
func CategoryCreate(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Category
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Errorln("bind error")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}

	utils.Log.Infoln("getting data")
	err = models.DB.Model(new(models.Category)).Create(&data).Error
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

// GetCategoryList
// @Tags 分类方法
// @Summary 分类列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param name query string false "name"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category_list [get]
func GetCategoryList(c *gin.Context) {

	utils.Log.Infoln("enter")
	type _param struct {
		Size string `form:"size" json:"size"`
		Page string `form:"page" json:"page"`
		Name string `form:"name" json:"name"`
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
	list := make([]models.Category, 0)
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64

	utils.Log.Infoln("getting data")
	tx := models.GetCategoryList(data.Name)
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

// GetCategoryDetail
// @Tags 分类方法
// @Summary 分类详细
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category_detail [get]
func GetCategoryDetail(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Category
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Errorln("bind error")
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	utils.Log.Infoln("getting data ")
	err = models.DB.Where("id=?", data.ID).First(&data).Error
	if err != nil {
		utils.Log.Errorln("get data error")
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "分类不存在",
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

// CategoryDelete
// @Tags 分类方法
// @Summary 分类删除
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category_delete [delete]
func CategoryDelete(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Category
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
				"msg":  "分类不存在",
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

// CategoryUpdate
// @Tags 分类方法
// @Summary 分类修改
// @Param id formData int true "id"
// @Param name formData string false "name"
// @Param remark formData string false "remark"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category_update [put]
func CategoryUpdate(c *gin.Context) {
	utils.Log.Infoln("enter")
	var data models.Category
	err := c.ShouldBind(&data)
	if err != nil {
		utils.Log.Errorln("bind error")
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error " + err.Error(),
		})

		return
	}
	utils.Log.Infoln("getting data")
	err = models.DB.Model(new(models.Category)).Where("id=?", data.ID).Updates(&data).Error
	if err != nil {
		utils.Log.Infoln("get data error")
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
