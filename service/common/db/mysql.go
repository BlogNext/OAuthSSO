package db

import (
	"github.com/OauthSSO/service/common/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func LoadDBConfig() {
	//注册db
	if db == nil {
		err := config.LoadConfig("db", "config", "yaml")
		if err != nil {
			panic("db配置加载失败")
		}
		dbConfig, _ := config.GetConfig("db")
		dbMap := dbConfig.GetStringMap("mysql")
		dsn := dbMap["sources"].(map[string]interface{})["dsn"].(string)

		//mysql的日志配置
		var newLogger logger.Interface
		if gin.Mode() == "release" {
			//正式环境
			//mysql慢查询日志打印
			newLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: 4 * time.Second, // 慢 SQL 阈值
					LogLevel:      logger.Error,    // Log level, 生产环境下，超过阈值的sql开启慢查询
					Colorful:      true,            // 禁用彩色打印
				},
			)
		} else {
			//非正式环境
			//mysql慢查询日志打印
			newLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Nanosecond,   // 慢 SQL 阈值
					LogLevel:      logger.Info, // Log level  warn和error级别下，slowthreshold才生效，非正式环境下开启sql打印
					Colorful:      true,          // 禁用彩色打印
				},
			)
		}

		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger:      newLogger,
			PrepareStmt: true,
		})
	}
}

//获取db实力
func GetDB() *gorm.DB {
	return db
}
