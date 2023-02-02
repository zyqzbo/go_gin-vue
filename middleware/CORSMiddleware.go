package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc { // 后端解决前端跨域问题
	return func(cxt *gin.Context) {
		cxt.Writer.Header().Set("Access-Control-Allow-Origin", "http://loccalhost:8080")
		cxt.Writer.Header().Set("Access-Control-Max-Age", "86400")   // 设置缓存时间
		cxt.Writer.Header().Set("Access-Control-Allow-Methods", "*") // 允许请求所有的方法
		cxt.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		cxt.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 判断是否是 Options 请求 是的话直接返回200 否则继续向下判定
		if cxt.Request.Method == http.MethodOptions {
			cxt.AbortWithStatus(200)
		} else {
			cxt.Next()
		}
	}
}
