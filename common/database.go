package common

import (
	"fmt"
	"go_gin+vue/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	db.AutoMigrate(&model.User{})
	//db.Create(&User{Name: "tom", Password: "123456", Phone: "12345678911"})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
