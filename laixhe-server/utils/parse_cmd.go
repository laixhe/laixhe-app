package utils

import (
	"flag"
	"log"
)

// 解析命令行的参数

type ParseCmd struct {
	ConfigFile string
}

var parseCmd = &ParseCmd{}

// laixhe-server --config=./laixhe-server.ini
// 初始化-解析命令行的参数
func init() {

	flag.StringVar(&parseCmd.ConfigFile, "config", "", "配置文件路径")

	flag.Parse()

	if parseCmd.ConfigFile == "" {
		parseCmd.ConfigFile = "./laixhe-server.ini"

		log.Println("InitParseCmd: Default loading ./laixhe-server.ini")
	}

}

func GetParseCmd() *ParseCmd {
	return parseCmd
}
