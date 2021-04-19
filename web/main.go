package main

import (
	"flag"
	"fmt"
	"net/http"
)

// 将命令行的参数解析到全局变量
var port = flag.String("port", "5053", "服务端监听端口")

// 静态(文件)服务器
func main() {
	// 解析命令行的参数
	flag.Parse()
	fmt.Println("▶▶▶启用端口:", *port)
	fmt.Println("▶▶▶可使用 --prot 修改端口，如：--prot=5053")

	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Println("✘✘✘服务启动失败:", err)
	}
}
