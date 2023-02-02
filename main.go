package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
)

func main() {
	//InitConfig()
	db := common.InitDB()
	fmt.Println("db", db)
	//defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	//port := viper.GetString("datasource.port")
	//if port != "" {
	//	panic(r.Run(":" + port))
	//}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

// 因用viper获取数据库失败暂时不用了
//func InitConfig() {
//	workDir, _ := os.Getwd()                 // 获取工作目录
//	viper.SetConfigName("application")       // 要读取的文件名
//	viper.SetConfigType("yml")               // 设置要读取的文件类型
//	viper.AddConfigPath(workDir + "/config") // 设置要读取的路径
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic("")
//	}
//}
