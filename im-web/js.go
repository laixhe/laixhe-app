package main

import (
	"embed"
	"net/http"
)

//go:embed js
var dist embed.FS

// DistFS 静态服务
var DistFS = http.FileServer(http.FS(dist))
