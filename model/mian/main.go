package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mang_srv/model"
	"os"
	"time"
)

func main() {
	dsn := "root:tang5230@tcp(127.0.0.1:3306)/mango?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注） 商品管理 管理商品
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger})
	if err != nil {
		panic(err)
	}
	//_ = db.AutoMigrate(&model.User{})
	//_ = db.AutoMigrate(&model.Directory{}, &model.UserCategoryDirectory{}, &model.HomeDirectoryCategoryDirectory{}, &model.HomeDirectory{})
	//_ = db.AutoMigrate(&model.HomeParameters{})
	//_ = db.AutoMigrate(&model.HomeUpdateLog{})
	_ = db.AutoMigrate(&model.Image{})

}
