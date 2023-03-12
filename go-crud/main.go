package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"john/gin-curd/common"
	"log"
	"os"
)

func main() {
	// 初始化DB
	InitConfig()
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

	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run(":8080")
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	// 配置文件的路径
	viper.AddConfigPath(workDir + "/config")
	// 配置文件的名称
	viper.SetConfigName("application")
	// 配置文件的文件类型
	viper.SetConfigType("yml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("配置文件读取失败")
		return
	}
}
