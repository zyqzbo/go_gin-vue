package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
	"go_gin+vue/dto"
	"go_gin+vue/model"
	"go_gin+vue/response"
	"go_gin+vue/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) { // 注册
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		// http.StatusUnprocessableEntity 是http的常量 相当于状态码
		// 用response文件统一封装的请求格式
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位数")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位数")
		return
	}

	// 如果没有名称传，就随机生成字符串作为名字
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)

	if isPhoneExist(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户存在不允许创建用户")
		return
	}

	// 创建用户
	// 给注册的用户密码做加密处理
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(fromPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) { // 源码要求的固定的写法，因为gin关于http请求的handler 方法是定义是需要传入这样一个参数
	DB := common.GetDB() // 引进db
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(phone) != 11 {
		// http.StatusUnprocessableEntity 是http的常量 相当于状态码
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号码必须为11位数"})
		// 用response文件统一封装的请求格式
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位数")
		return
	}

	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码不能少于6位数"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位数")
		return
	}

	// 判断手机号码是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 442, "msg": "用户不存在"})
		response.Response(ctx, http.StatusInternalServerError, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		response.Response(ctx, http.StatusInternalServerError, 422, nil, "用户不存在")
		return
	}

	// 发送token 需要调用token验证方法
	token, err := common.ReleaseToken(user)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token异常：%v", err)
		return
	}

	// 返回结果
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登陆成功",
	//})
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(cxt *gin.Context) { // 从上下文中获取用户
	user, _ := cxt.Get("user")
	cxt.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)), // 通过ToUserDto 方法过滤掉password
		},
	})
}
