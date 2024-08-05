package database

import (
	"fmt"
	"gofiber/app/config"
	_ "gofiber/app/module/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

// InitGorm 初始化Gorm
func init() {
	databaseConf := config.Conf.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", databaseConf.User, databaseConf.Pass, databaseConf.Server, databaseConf.Port, databaseConf.DbName)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: new(GormZapLogger).LogMode(logger.Info),
	})

	Db = db
}
