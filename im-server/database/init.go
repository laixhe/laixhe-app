package database

import (
	"im-server/database/dbx"
	"im-server/database/redisx"
)

func init() {
	dbx.InitDB()
	redisx.InitRedis()
}
