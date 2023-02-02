package controller

import "github.com/gin-gonic/gin"

type RestController interface { // 多处要用增删改查，封装起来用
	Create(cxt *gin.Context)
	Update(cxt *gin.Context)
	Show(cxt *gin.Context)
	Delete(cxt *gin.Context)
}
