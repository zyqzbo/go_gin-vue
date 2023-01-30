package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
	"go_gin+vue/model"
	"go_gin+vue/util"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		// http.StatusUnprocessableEntity 是http的常量 相当于状态码
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号码必须为11位数"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码不能少于6位数"})
		return
	}

	// 如果没有名称传，就随机生成字符串作为名字
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)

	if isPhoneExist(DB, phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "用户存在不允许创建用户"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	log.Println("34")
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
