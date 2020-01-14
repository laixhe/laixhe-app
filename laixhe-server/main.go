package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/laixhe/laixhe-app/laixhe-server/utils/logs"

	"github.com/laixhe/laixhe-app/laixhe-server/config"
	"github.com/laixhe/laixhe-app/laixhe-server/db"
	"github.com/laixhe/laixhe-app/laixhe-server/routers"
	"github.com/laixhe/laixhe-app/laixhe-server/utils"
)

func main() {

	// laixhe-server --config=./config.ini
	// 初始化-解析命令行的参数
	utils.InitParseCmd()

	// 初始化配置文件
	err := config.InitConfig(utils.GetParseCmd().ConfigFile)
	if err != nil {
		panic(err)
	}

	// 初始 zap 日志
	logs.ZapLogInit(
		config.GetLog().Path,
		config.GetLog().Level,
		config.GetLog().MaxSize,
		config.GetLog().MaxBackup,
		config.GetLog().MaxAge)

	// 初始化数据库
	err = db.Init(
		config.GetDB().DbType,
		config.GetDB().Dsn,
		config.GetDB().OpenMax,
		config.GetDB().IdleMax)
	if err != nil {
		panic(err)
	}

	// 启动项目
	err = routers.Run(config.GetHttp().Addr, config.GetApp().AppMode)
	if err != nil {
		panic(err)
	}

}
