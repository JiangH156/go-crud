package controller

import "github.com/gin-gonic/gin"

// 提供restful接口的抽象接口
type RestControoler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Query(ctx *gin.Context)
}
