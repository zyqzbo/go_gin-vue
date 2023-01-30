package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
)

func main() {
	db := common.InitDB()
	fmt.Println("db", db)
	//defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
