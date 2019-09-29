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

// 初始化-解析命令行的参数
func InitParseCmd() {

	flag.StringVar(&parseCmd.ConfigFile, "config", "", "配置文件路径")

	flag.Parse()

	if parseCmd.ConfigFile == "" {
		parseCmd.ConfigFile = "./config.ini"

		log.Println("InitParseCmd: Default loading ./config.ini")
	}

}

func GetParseCmd() *ParseCmd {
	return parseCmd
}
