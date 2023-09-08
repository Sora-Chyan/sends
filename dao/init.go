package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sends/models"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=LocalTime"
	//dsn := "root:123456@tcp(localhost:3306)/graduation?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm.Init ERROR :", err)
	}
	//自动创建表
	err = db.AutoMigrate(
		models.Wait{})
	if err != nil {

		log.Println("AutoMigrate ERROR :", err)
	}

	//创建记录
	//u1 := models.User{OpenId: "123456", StuNum: "1711111031", StoryInit: false}
	//db.Create(&u1)

	return db

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
