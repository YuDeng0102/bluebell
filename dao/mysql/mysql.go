package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web-app/model"
	"web-app/settings"
)

var db *gorm.DB
var sqlDB *sql.DB // 保存底层 *sql.DB 对象，用于关闭连接池

func Init(config *settings.MySQLConfig) (err error) {
	username := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	database := config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	//dsn := "root:041135sz@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Error("failed to connect mysql database,err:" + err.Error())
		return err
	}
	// 连接池
	sqlDB, err = db.DB()
	if err != nil {
		zap.L().Error("connect db server failed, err:" + err.Error())
		return
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns) // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(config.MaxOpenConns) //设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)        //设置了连接可复用的最大时间

	//数据库迁移
	_ = db.AutoMigrate(&model.User{})

	zap.L().Info("mysql init success")
	return
}
func Close() {
	if sqlDB != nil {
		_ = sqlDB.Close()
		zap.L().Info("MySQL connection pool closed")
	}
}
