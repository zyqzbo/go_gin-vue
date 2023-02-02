package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) { // 封装请求格式
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) { // 根据上面封装方法定义请求成功时的返回字段格式
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) { // 根据上面封装方法定义请求失败时的返回字段格式
	Response(ctx, http.StatusOK, 400, data, msg)
}
