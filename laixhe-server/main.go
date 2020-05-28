package main

import (
	"github.com/laixhe/laixhe-app/laixhe-server/db"
	"github.com/laixhe/laixhe-app/laixhe-server/logs"
	"github.com/laixhe/laixhe-app/laixhe-server/routers"
)

// laixhe-server --config=./laixhe-server.ini
func main() {

	logs.Init()
	db.Init()

	// 启动项目
	err := routers.Run()
	if err != nil {
		panic(err)
	}

}
