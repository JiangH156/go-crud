package main

import (
	"github.com/gin-gonic/gin"
	"john/gin-curd/controller"
	"john/gin-curd/middlerware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 解决跨域问题
	r.Use(middlerware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", middlerware.AuthMiddleware(), controller.Login)
	r.GET("/api/auth/info", middlerware.AuthMiddleware(), controller.Info)

	// category
	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("/create", categoryController.Create)
	categoryRoutes.DELETE("/delete/:id", categoryController.Delete)
	categoryRoutes.PUT("/update/:id", categoryController.Update)
	categoryRoutes.GET("/query/:id", categoryController.Query)

	// post
	postRoutes := r.Group("/post")
	r.Use(middlerware.CORSMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("/create", postController.Create)
	postRoutes.DELETE("/delete/:id", postController.Delete)
	postRoutes.PUT("/update/:id", postController.Update)
	postRoutes.GET("/query/:id", postController.Query)

	return r
}
