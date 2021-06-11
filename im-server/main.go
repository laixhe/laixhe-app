package main

import (
	"im-server/routers"

	_ "im-server/database"
)

func main() {
	if err := routers.Run(); err != nil {
		panic(err)
	}
}
