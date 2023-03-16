package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"john/gin-curd/models"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// username:password@(host:port)/database?charset=charset&parseTime=True&loc=loc
	// 注册驱动
	//driverName := "mysql"
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connect failed")
	}
	// 数据库初始化失败，直接panic

	// 迁移创建数据库表
	db.AutoMigrate(&models.User{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
