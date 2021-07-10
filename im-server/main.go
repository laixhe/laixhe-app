package main

import (
	_ "im-server/app"

	"im-server/routers"
)

func main() {
	if err := routers.Run(); err != nil {
		panic(err)
	}
}
