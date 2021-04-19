package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"im-server/config"
	"im-server/routers"
)

func main() {
	if config.AppPid() != "" {
		if err := ioutil.WriteFile(config.AppPid(), []byte(fmt.Sprintf("%d", os.Getpid())), 0666); err != nil {
			panic(err)
		}
	}
	if err := routers.Run(); err != nil {
		panic(err)
	}
}
