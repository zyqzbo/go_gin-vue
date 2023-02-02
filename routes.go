package main

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/controller"
	"go_gin+vue/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware()) // 跨域处理 、 报错捕捉显示

	// 注册、登陆、获取用户基本信息 API
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // AuthMiddleware token验证

	// 分类 增删改查 API
	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	// 还有一个PATCH 方法 是局部修改的意思 而PUT是整个替换的意思

	// 文章 增删改查 API
	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware()) // 使用中间键 获取用户信息
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	// 分页查询
	postRoutes.POST("page/list", postController.PageList)

	return r

}
