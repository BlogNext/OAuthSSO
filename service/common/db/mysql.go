package db

import (
	"github.com/OauthSSO/service/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	//注册db
	if db == nil {
		err := config.LoadConfig("db", "config", "yaml")
		if err != nil {
			panic("db配置加载失败")
		}
		dbConfig, _ := config.GetConfig("db")
		dbMap := dbConfig.GetStringMap("mysql")
		dsn := dbMap["sources"].(map[string]interface{})["dsn"].(string)
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}

//获取db实力
func GetDB() *gorm.DB {
	return db
}
