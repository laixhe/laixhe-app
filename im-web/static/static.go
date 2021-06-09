package static

import (
	"embed"
	"net/http"
)

//go:embed index.html img js
var index embed.FS

// IndexFS 静态服务
var IndexFS = http.FileServer(http.FS(index))
