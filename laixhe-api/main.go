package main

import (
	"github.com/laixhe/laixhe-app/laixhe-api/db"
	"github.com/laixhe/laixhe-app/laixhe-api/logs"
	"github.com/laixhe/laixhe-app/laixhe-api/routers"
)

// laixhe-api --config=./laixhe-api.ini
func main() {

	logs.Init()
	db.Init()

	// 启动项目
	err := routers.Run()
	if err != nil {
		panic(err)
	}

}
