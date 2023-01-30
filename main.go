package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);no null"`
	Phone    string `gorm:"varchar(100;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	//defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
			name = RandomString(10)
		}

		log.Println(name, phone, password)

		if isPhoneExist(db, phone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "用户存在不允许创建用户"})
			return
		}
		// 创建用户
		newUser := User{
			Name:     name,
			Phone:    name,
			Password: password,
		}
		db.Create(&newUser)

		// 返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	log.Println("34")
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("sdafgertfsdfgssvfFGHHAASd")
	result := make([]byte, n)

	// 获取随机数 不添加随机种子，确保每一次的随机数都是不重复的
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := 3306
	Dbname := "go_gin_db"
	username := "root"
	password := "zyq4836.."
	timeout := "10s"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username,
		password,
		host,
		port,
		Dbname,
		timeout)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	//db.Create(&User{Name: "tom", Password: "123456", Phone: "12345678911"})
	return db
}
