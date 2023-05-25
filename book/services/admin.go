package services

import (
	"book/config"
	"book/models"
	"book/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SendCode
// @Tags 管理员方法
// @Summary 发送邮箱验证码
// @Param email formData string false "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/code [post]
func SendCode(c *gin.Context) {
	var data models.Admin
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}

	email := strings.ToLower(data.Email)
	code := utils.GetRand()
	err = utils.SendCode(email, code) //发送验证码给邮箱
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "send code Error:" + err.Error(),
		})
		return
	}
	models.RDB.Set(c, email, code, time.Second*300) //保存验证码到redis
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags 管理员方法
// @Summary 管理员注册
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param password2 formData string true "password2"
// @Param phone formData string false "phone"
// @Param email formData string true "email"
// @Param code formData string true "code"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/register [post]
func Register(c *gin.Context) {
	type temp struct {
		Password2 string `form:"password2"`
		Code      string `form:"code"`
	}
	var data models.Admin
	var temp_data temp
	err := c.ShouldBind(&data)
	err = c.ShouldBind(&temp_data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "bind error",
		})
		return
	}
	//密码是否一致
	password := data.Password
	password2 := temp_data.Password2
	if password2 != password {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "密码不一致",
		})
		return
	}

	//判断邮箱是否已存在
	data.Email = strings.ToLower(data.Email)
	email := data.Email
	var cnt int64
	err = models.DB.Where("email=?", email).Model(new(models.Admin)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get user error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册",
		})
		return
	}

	//获取验证码
	syscode, err := models.RDB.Get(c, email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请重新获取验证码",
		})
		return
	}
	usercode := temp_data.Code
	if syscode != usercode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}

	//数据插入
	data.Uuid = utils.GetUUID()
	data.Password = utils.GetMd5(data.Password)
	err = models.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "create user error" + err.Error(),
		})
		return
	}

	//生成token
	token, err := utils.GenerateToken(data.Uuid, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "generate token error" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// Login
// @Tags 管理员方法
// @Summary 管理员登录
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var data models.Admin
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "bind error",
		})
	}
	data.Password = utils.GetMd5(data.Password)
	err = models.DB.Where("name =? And password =?", data.Name, data.Password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get AdminBasic Error" + err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(data.Uuid, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// GetAdminList
// @Tags 管理员方法
// @Summary 管理员列表
// @Param page query string false "page"
// @Param size query string false "size"
// @Param name query string false "name"
// @Param phone query string false "phone"
// @Param email query string false "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/admin_list [get]
func GetAdminList(c *gin.Context) {
	type _param struct {
		Size  string `form:"size" json:"size"`
		Page  string `form:"page" json:"page"`
		Name  string `form:"name" json:"name"`
		Phone string `form:"phone" json:"phone"`
		Email string `form:"email" json:"email"`
	}
	data := &_param{
		Size: config.DefaultSize,
		Page: config.DefaultPage,
	}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "bind error" + err.Error(),
		})
		return
	}
	page, _ := strconv.Atoi(data.Page)
	size, _ := strconv.Atoi(data.Size)
	page = (page - 1) * size
	var count int64

	list := make([]*models.Admin, 0)

	tx := models.GetAdminList(data.Name, data.Phone, data.Email)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Rank List Error:" + err.Error(),
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
