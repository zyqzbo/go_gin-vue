package middleware

// 为了拦截重复创建分类是出现的报错信息显示出来 for CategoryController Create
//定义的拦截方法  for CategoryController Create 重复创建出现报错的时候配合panic使用
//gin框架也有在main.go 的 r := gin.Default() 的 Default 点进去的 recover()拦截显示方法

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin+vue/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(context, nil, fmt.Sprint(err))
			}
		}()

		context.Next()
	}
}
