package router

import (
	"book/docs"
	"book/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	//跨域
	r.Use(Cors())
	router := r.Group("/api")
	//swagger调试
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//静态文件
	router.StaticFS("imgs", http.Dir("imgs"))

	//管理员
	router.POST("/code", services.SendCode)
	router.POST("/register", services.Register)
	router.POST("/login", services.Login)
	router.GET("/admin_list", services.GetAdminList)
	//用户
	router.POST("/user_create", services.UserCreate)
	router.GET("/user_list", services.GetUserList)
	router.POST("/user_account", services.UserAccount)
	router.GET("/user_detail", services.GetUserDetail)
	router.DELETE("/user_delete", services.UserDelete)
	router.PUT("/user_update", services.UserUpdate)
	//分类
	router.POST("/category_create", services.CategoryCreate)
	router.GET("/category_list", services.GetCategoryList)
	router.GET("/category_detail", services.GetCategoryDetail)
	router.DELETE("/category_delete", services.CategoryDelete)
	router.PUT("/category_update", services.CategoryUpdate)
	//图书
	router.POST("/bookimgup", services.CoverPhotoUp)
	router.POST("/book_create", services.BookCreate)
	router.GET("/book_list", services.GetBookList)
	router.GET("/book_detail", services.GetBookDetail)
	router.DELETE("/book_delete", services.BookDelete)
	router.PUT("/book_update", services.BookUpdate)
	//借书
	router.POST("/borrow_create", services.BorrowCreate)
	router.GET("/borrow_list", services.GetBorrowList)
	router.DELETE("/borrow_delete", services.BorrowDelete)
	//还书
	router.POST("/retur_create", services.ReturCreate)
	router.GET("/retur_list", services.GetReturList)
	router.DELETE("/retur_delete", services.ReturDelete)
	//主页echarts数据
	router.GET("/echart_data", services.EchartsData)
	return r
}
