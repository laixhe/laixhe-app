#!/bin/bash

dist_code="
package dist

import (
	\"embed\"
	\"net/http\"
)

//go:embed index.html favicon.ico assets
var dist embed.FS

// DistFS 静态服务
var DistFS = http.FileServer(http.FS(dist))
"

main_code="
package main

import (
	\"im-web/dist\"
	\"net/http\"
)

func main() {
	http.Handle(\"/\", dist.DistFS)
	http.ListenAndServe(\":5053\", nil)
}
"

echo "${dist_code}" > dist/dist.go

echo "${main_code}" > main.go

go build

rm main.go
rm dist/dist.go
