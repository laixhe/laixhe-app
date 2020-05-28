package config

import (
	"errors"

	"github.com/laixhe/goutil"

	"github.com/laixhe/laixhe-app/laixhe-api/utils"
)

// 初始化配置文件
func init() {

	// 加载配置文件
	err := goutil.OpenConfig(utils.GetParseCmd().ConfigFile)
	if err != nil {
		panic(err)
	}

	err = setApp()
	if err != nil {
		panic(err)
	}

	err = setLog()
	if err != nil {
		panic(err)
	}

	err = setHttp()
	if err != nil {
		panic(err)
	}

	err = setDB()
	if err != nil {
		panic(err)
	}

}

// 项目配置
type AppConfig struct {
	AppMode string // 项目运行模式
	Token   string
}

var ac = &AppConfig{}

func setApp() error {

	ac.AppMode = goutil.ConfigGetMustString("app.mode", "debug")

	token := goutil.ConfigGetString("app.token")
	if token == "" {
		token = goutil.GetRandomString(16)
	}
	ac.Token = token

	return nil
}

func GetApp() *AppConfig {
	return ac
}

// 日志 相关配置
type LogConfig struct {
	Path      string // 日志文件路径
	Level     string // 设置日志级别
	MaxSize   int    // 每个日志文件保存大小，为M
	MaxBackup int    // 保留 N 个备份
	MaxAge    int    // 保留 N 天
}

var lc = &LogConfig{}

func setLog() error {

	logPath := goutil.ConfigGetString("log.path")
	if logPath == "" {
		return errors.New("get log path err")
	}

	lc.Path = logPath
	lc.Level = goutil.ConfigGetString("log.level")
	lc.MaxSize = goutil.ConfigGetMustInt("log.max_size", 0)
	lc.MaxBackup = goutil.ConfigGetMustInt("log.max_backup", 0)
	lc.MaxAge = goutil.ConfigGetMustInt("log.max_age", 0)

	return nil
}

func GetLog() *LogConfig {
	return lc
}

// http 相关配置
type httpConfig struct {
	Addr string // http 运行的端口
}

var hc = &httpConfig{}

func setHttp() error {

	addr := goutil.ConfigGetString("http.addr")
	if addr == "" {
		return errors.New("get http addr err")
	}

	hc.Addr = addr

	return nil
}

func GetHttp() *httpConfig {
	return hc
}

// DB 相关配置
type dbConfig struct {
	DbType  string // 数据库类型
	Dsn     string // 数据库连接信息
	OpenMax int    // 设置最大打开的连接数
	IdleMax int    // 设置闲置的连接数
}

var dc = &dbConfig{}

func setDB() error {

	// 数据库类型
	dbType := goutil.ConfigGetString("db.db_type")
	if dbType == "" {
		return errors.New("get db db_type err")
	}

	// 数据库连接信息
	dsn := goutil.ConfigGetString("db.dsn")
	if dbType == "" {
		return errors.New("get db dsn err")
	}

	// 设置最大打开的连接数
	openMax, err := goutil.ConfigGetInt("db.open_max")
	if err != nil {
		openMax = 100
	}

	// 设置闲置的连接数
	idleMax, err := goutil.ConfigGetInt("db.idle_max")
	if err != nil {
		idleMax = 10
	}

	dc.DbType = dbType
	dc.Dsn = dsn
	dc.OpenMax = openMax
	dc.IdleMax = idleMax

	return nil
}

func GetDB() *dbConfig {
	return dc
}
