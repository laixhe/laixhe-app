package db

import (
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

// 初始化数据库
func Init(dbType, dsn string, openMax, idleMax int) error {

	var err error
	dbx, err = sqlx.Connect(dbType, dsn)
	if err != nil {
		return err
	}

	err = dbx.Ping()
	if err != nil {
		return err
	}

	// 设置闲置的连接数
	dbx.SetMaxIdleConns(idleMax)
	// 设置最大打开的连接数
	dbx.SetMaxOpenConns(openMax)

	return nil
}

func SQLX() *sqlx.DB {
	return dbx
}
