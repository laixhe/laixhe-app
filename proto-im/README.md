### proto-im

#### 生成工具 - protoc 插件
```
# protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.8
protoc --version

# protoc-gen-go
https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.26.0
protoc-gen-go --version

# protoc-gen-go-grpc
https://github.com/grpc/grpc-go/releases/tag/cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0
protoc-gen-go-grpc --version
```

#### proto 文件编译
```
protoc --go-grpc_out=. --go_out=. *.proto

protoc 参数说明
--proto_path     proto 文件目录(简写 -I 如果没有指定参数，则在当前目录进行搜索)
--go-grpc_out=   生成的 go-grpc 源码保存目录
--go_out         生成的 go 源码保存目录
```
