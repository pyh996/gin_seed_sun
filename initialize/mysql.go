package initialize

import (
	"fmt"
	"go_gin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() {
	mysqlInfo := global.Settings.Mysqlinfo
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
		mysqlInfo.Port, mysqlInfo.DBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
}
