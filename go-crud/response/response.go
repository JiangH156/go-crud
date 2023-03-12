package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response
// @Description: 处理返回响应
// @param ctx *gin.Context 上下文
// @param httpStatus int http状态码
// @param code int 返回前端状态码
// @param msg string 信息
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}
