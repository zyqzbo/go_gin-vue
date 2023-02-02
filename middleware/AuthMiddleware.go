package middleware

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
	"go_gin+vue/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header （获取授权）
		tokenSting := ctx.GetHeader("Authorization")

		// 验证tonken的格式 如果为空 或者不是以Bearer开头
		if tokenSting == "" || !strings.HasPrefix(tokenSting, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() // 抛弃此次请求
			return
		}
		// 提取token的有效部分 从第7开始截取 才是有用部分，因为 Bearer 加上后面的空格一共七位
		tokenSting = tokenSting[7:]
		// 对tokenString进行解析
		token, claims, err := common.ParseToken(tokenSting)
		// 如果解析失败，或者解析后的token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() // 抛弃此次请求
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId) // 从数据查询到第一条后返回

		// 如果用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
		}

		// 用户存在 将user 的信息写入上下文
		ctx.Set("user", user)

		//Next()只能在中间件中使用，会挂起当前中间件(也就是Next()后面的代码先不执行)，开始执行后面的中间件和最后的handler。
		//后续的中间件和最后的handler执行完毕后，再回到调用Next()的中间件，继续执行Next()之后的代码。
		ctx.Next()
	}
}
