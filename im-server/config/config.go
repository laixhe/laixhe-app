package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/laixhe/goutil/zaplog"
	"github.com/spf13/viper"
)

// AppConfig App项目配置
type AppConfig struct {
	Name string // 项目名
	Mode string // 运行模式
	IP   string // 运行IP
	Port int    // 运行端口
	Pid  string // PID 存放文件
}

// LogsConfig 日志配置
type LogsConfig struct {
	Path      string // 日志文件路径: 【当为空时是 控制台 输出】【当不为空(./logs.log)时是 文件 输出】
	LogLevel  string // 日志级别 debug  info  error
	MaxSize   int    // 每个日志文件保存大小 ?M
	MaxBackup int    // 保留 N 个备份
	MaxAge    int    // 保留 N 天
}

// WebsocketConfig websocket配置
type WebsocketConfig struct {
	ReadTime  time.Duration // 读超时(秒)，同时也是心跳超时
	WriteTime time.Duration // 写超时(秒)
}

// MysqlConfig Mysql配置
type MysqlConfig struct {
	Dsn             string        // 连接地址
	MaxIdleConn     int           // 设置空闲连接池中连接的最大数量
	MaxOpenConn     int           // 设置打开数据库连接的最大数量
	ConnMaxLifetime time.Duration // 设置了连接可复用的最大时间(单位：分)
}

// RedisConfig Redis配置
type RedisConfig struct {
	DBNum    int      // 选择数据库
	Password string   // 密码
	Nodes    []string // 当节点只有一个视为单机
}

// JwtConfig jwt配置
type JwtConfig struct {
	SecretKey string
	ExpTime   int // 过期时长(秒)
}

// Config 总配置
type Config struct {
	App       AppConfig
	Logs      LogsConfig
	Websocket WebsocketConfig
	Mysql     MysqlConfig
	Redis     RedisConfig
	Jwt       JwtConfig
}

var conf *Config

func AppName() string {
	if conf.App.Name == "" {
		if len(os.Args) == 0 {
			conf.App.Name = "app-name"
		} else {
			conf.App.Name = os.Args[0]
		}
	}
	return conf.App.Name
}

// IsAppDebug 是否是调试模式
func IsAppDebug() bool {
	if conf.App.Mode == "" {
		conf.App.Mode = ModeDebug
	}
	return conf.App.Mode == ModeDebug
}

// AppIP 获取 http 运行IP
func AppIP() string {
	if conf.App.IP == "" {
		conf.App.IP = "0.0.0.0"
	}
	return conf.App.IP
}

// AppAddr 获取 http 运行端口
func AppAddr() int {
	if conf.App.Port <= 0 || conf.App.Port > 65535 {
		conf.App.Port = 5050
	}
	return conf.App.Port
}

// AppIPAddr 获取 http 运行IP:端口
func AppIPAddr() string {
	return fmt.Sprintf("%s:%d", AppIP(), AppAddr())
}

// AppPid PID 存放文件
func AppPid() string {
	if conf.App.Pid == "" {
		conf.App.Pid = AppName() + ".pid"
	}
	return conf.App.Pid
}

// WebsocketReadTime 读超时(秒)，同时也是心跳超时
func WebsocketReadTime() time.Duration {
	if conf.Websocket.ReadTime <= 0 {
		conf.Websocket.ReadTime = 20
	}
	return conf.Websocket.ReadTime
}

func WebsocketWriteTime() time.Duration {
	if conf.Websocket.WriteTime <= 0 {
		conf.Websocket.WriteTime = 3
	}
	return conf.Websocket.WriteTime
}

// DBDsn 获取数据库连接地址
func DBDsn() string {
	if conf.Mysql.Dsn == "" {
		panic(errors.New("获取数据库连接地址失败：为空"))
	}
	return conf.Mysql.Dsn
}

// DBMaxIdleConn 获取数据库-空闲连接池中连接的最大数量
func DBMaxIdleConn() int {
	if conf.Mysql.MaxIdleConn <= 0 {
		conf.Mysql.MaxIdleConn = 5
	}
	return conf.Mysql.MaxIdleConn
}

// DBMaxOpenConn 获取数据库-打开数据库连接的最大数量
func DBMaxOpenConn() int {
	if conf.Mysql.MaxOpenConn <= 0 {
		conf.Mysql.MaxOpenConn = 100
	}
	return conf.Mysql.MaxOpenConn
}

// DBConnMaxLifetime 设置了连接可复用的最大时间(单位：分)
func DBConnMaxLifetime() time.Duration {
	if conf.Mysql.ConnMaxLifetime <= 0 {
		conf.Mysql.ConnMaxLifetime = 30
	}
	return conf.Mysql.ConnMaxLifetime
}

// RedisDBNum 选择数据库
func RedisDBNum() int {
	if conf.Redis.DBNum < 0 {
		conf.Redis.DBNum = 0
	}
	return conf.Redis.DBNum
}

// RedisPassword 密码
func RedisPassword() string {
	return conf.Redis.Password
}

// RedisNodes 节点
func RedisNodes() []string {
	if len(conf.Redis.Nodes) == 0 {
		panic(errors.New("获取Redis连接地址失败：为空"))
	}
	if len(conf.Redis.Nodes) == 1 {
		if conf.Redis.Nodes[0] == "" {
			panic(fmt.Errorf("获取Redis连接地址失败：%#v", conf.Redis.Nodes))
		}
	} else {
		for _, v := range conf.Redis.Nodes {
			if v == "" {
				panic(fmt.Errorf("获取Redis连接地址失败：%#v", conf.Redis.Nodes))
			}
		}
	}
	return conf.Redis.Nodes
}

// IsConsoleLog 日志是否控制台输出
func IsConsoleLog() bool {
	return conf.Logs.Path == ""
}

// initLog 初始化日志
func initLog() {
	switch conf.Logs.LogLevel {
	case LogDebug:
	case LogInfo:
	case LogError:
	default:
		conf.Logs.LogLevel = LogDebug
	}
	if conf.Logs.MaxSize < 0 {
		conf.Logs.MaxSize = 0
	}
	if conf.Logs.MaxBackup < 0 {
		conf.Logs.MaxBackup = 0
	}
	if conf.Logs.MaxAge < 0 {
		conf.Logs.MaxAge = 0
	}
	if IsConsoleLog() {
		zaplog.InitConsole(zaplog.Config{
			ServiceName: conf.App.Name,
			LogLevel:    conf.Logs.LogLevel,
		})
	} else {
		zaplog.InitSizeFile(zaplog.Config{
			ServiceName: AppName(),
			LogPath:     conf.Logs.Path,
			LogLevel:    conf.Logs.LogLevel,
			MaxSize:     conf.Logs.MaxSize,
			MaxBackups:  conf.Logs.MaxBackup,
			MaxAge:      conf.Logs.MaxAge,
		})
	}
}
func init() {
	v := viper.New()
	// 设置配置文件的名字
	v.SetConfigName("config")
	// 添加配置文件所在的路径
	v.AddConfigPath("./")
	// 设置配置文件类型
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = v.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	// 一些必要检查
	DBDsn()
	RedisNodes()

	initLog()
	fmt.Printf("%#v\n", conf.Websocket)
	zaplog.Debug("▶▶▶配置与日志初始化完毕...")
}
