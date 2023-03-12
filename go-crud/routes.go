package main

import (
	"github.com/gin-gonic/gin"
	"john/gin-curd/controller"
	"john/gin-curd/middlerware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/login", middlerware.AuthMiddleware(), controller.Login)
	return r
}
