package services

import (
	"book/config"
	"book/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// UserCreate
// @Tags 用户方法
// @Summary 用户注册
// @Param name formData string true "name"
// @Param age formData string true "age"
// @Param sex formData string true "sex"
// @Param phone formData string false "phone"
// @Param address formData string true "address"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /api/user_create [post]
func UserCreate(c *gin.Context) {
	var data models.User
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	err = models.DB.Model(new(models.User)).Create(&data).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "create error" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "创建成功",
	})
}

// GetUserList
// @Tags 用户方法
// @Summary 用户列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param name query string false "name"
// @Param phone query string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/user_list [get]
func GetUserList(c *gin.Context) {
	type _param struct {
		Size  string `form:"size" json:"size"`
		Page  string `form:"page" json:"page"`
		Name  string `form:"name" json:"name"`
		Phone string `form:"phone" json:"phone"`
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
	list := make([]models.User, 0)
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64
	tx := models.GetUserList(data.Phone, data.Name)
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

// UserAccount
// @Tags 用户方法
// @Summary 账号充值
// @Param id formData string true "id"
// @Param score formData int true "score"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/user_account [post]
func UserAccount(c *gin.Context) {
	var data models.User
	err := c.ShouldBind(&data)
	type params struct {
		Score int `form:"score"`
	}
	var param_data params
	err = c.ShouldBind(&param_data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	score := param_data.Score

	err = models.DB.Where("id=?", data.ID).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Detail Error:" + err.Error(),
		})
		return
	}
	data.Account = data.Account + score
	err = models.DB.Save(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Detail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})
}

// GetUserDetail
// @Tags 用户方法
// @Summary 用户详细
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/user_detail [get]
func GetUserDetail(c *gin.Context) {
	var data models.User
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	err = models.DB.Where("id=?", data.ID).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Detail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})
}

// UserDelete
// @Tags 用户方法
// @Summary 用户删除
// @Param id query string true "id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/user_delete [delete]
func UserDelete(c *gin.Context) {
	var data models.User
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
				"msg":  "用户不存在",
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

// UserUpdate
// @Tags 用户方法
// @Summary 用户修改
// @Param id formData int true "id"
// @Param name formData string false "name"
// @Param age formData string false "age"
// @Param address formData string false "address"
// @Param sex formData string false "sex"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/user_update [put]
func UserUpdate(c *gin.Context) {
	var data models.User
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	err = models.DB.Model(new(models.User)).Where("id=?", data.ID).Updates(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "update error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "修改成功",
	})
}
