package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"john/gin-curd/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// username:password@(host:port)/database?charset=charset&parseTime=True&loc=loc
	// 注册驱动
	//driverName := "mysql"
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := "3309"
	database := "db1"
	charset := "utf8mb4"
	loc := "Local"

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 数据库初始化失败，直接panic
	if err != nil {
		panic("database connect failed")
	}
	// 迁移创建数据库表
	db.AutoMigrate(&models.User{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
