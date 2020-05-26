package main

import (
	"github.com/laixhe/laixhe-app/laixhe-server/routers"
)

func main() {

	// 启动项目
	err := routers.Run()
	if err != nil {
		panic(err)
	}

}
