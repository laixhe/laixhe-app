package main

import (
	"net/http"
)

func main() {
	// 设置文件服务
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":5053", nil)
	if err != nil {
		panic(err)
	}
}
