package main

import (
	"github.com/gin-gonic/gin"
	"john/gin-curd/common"
)

func main() {
	// 初始化DB
	db := common.InitDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		defer sqlDB.Close()
	}()

	// gin路由拦截
	r := gin.Default()
	CollectRoute(r)
	r.Run(":8080")
}
