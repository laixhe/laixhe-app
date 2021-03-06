package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/laixhe/goutil/db_sqlx"

	"github.com/laixhe/laixhe-app/laixhe-api/config"
)

// 初始化数据库
func Init() {

	err := db_sqlx.Init(
		config.GetDB().DbType,
		config.GetDB().Dsn,
		config.GetDB().OpenMax,
		config.GetDB().IdleMax)

	if err != nil {
		panic(err)
	}

}