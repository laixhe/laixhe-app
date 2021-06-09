package main

import (
	"flag"
	"fmt"
	"net/http"

	"im-web/static"
)

var (
	debug = flag.Bool("debug", false, "运行模式")
	port  = flag.Int("port", 5053, "服务端监听端口")
)

func main() {
	flag.Parse()

	show := `
-debug bool
	运行模式 %v
-port int
	服务端监听端口 %v
	`
	fmt.Printf(show+"\n", *debug, *port)

	// 设置文件服务
	if *debug {
		http.Handle("/", http.FileServer(http.Dir("./static")))
	} else {
		http.Handle("/", static.IndexFS)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic(err)
	}
}
