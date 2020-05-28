module github.com/laixhe/laixhe-app/laixhe-server

go 1.14

replace (
	go.uber.org/zap => github.com/uber-go/zap v1.15.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.6
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/laixhe/goutil v0.0.0-20200528092253-861ac326df0e
)
