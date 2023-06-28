package dao

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var sqlDb *sql.DB
var db *gorm.DB

func InitDataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	// 配置
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("mysql.tablePrefix"),
			SingularTable: true,
		},
		Logger: logger.Default, //默认不输出sql
	}

	if viper.GetBool("mysql.logMode") {
		gcf.Logger = logger.Default.LogMode(logger.Info) // logger.info 输出 sql
	}
	db, _ = gorm.Open(mysql.Open(dsn), gcf)
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(1000)
	sqlDb.SetMaxOpenConns(5000)
	sqlDb.SetConnMaxLifetime(-1)
	db.AutoMigrate()
}
