package dbx

import (
	"time"

	"github.com/laixhe/goutil/zaplog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"im-server/config"
)

var client *gorm.DB

// InitDB init
func InitDB() {
	var logMode = logger.Silent
	if config.IsAppDebug() {
		logMode = logger.Info
	}
	zaplog.Debug("▶▶▶数据库初始化...")
	var err error
	client, err = gorm.Open(mysql.Open(config.DBdsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		panic(err)
	}
	db, err := client.DB()
	if err != nil {
		panic(err)
	}
	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(config.DBMaxIdleConn())
	// 设置打开数据库连接的最大数量
	db.SetMaxOpenConns(config.DBMaxOpenConn())
	// 设置了连接可复用的最大时间
	db.SetConnMaxLifetime(config.DBConnMaxLifetime() * time.Second)
	// 验证数据库连接是否正常
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	zaplog.Debug("▶▶▶数据库初始化完毕...")
}
