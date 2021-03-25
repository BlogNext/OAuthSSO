package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	//注册db
	if db == nil {
		dsn := "root:admin123@tcp(154.8.142.48:3306)/blog_xiaochen?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}

//获取db实力
func GetDB() *gorm.DB {
	return db
}
