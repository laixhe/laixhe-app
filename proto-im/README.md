### proto-im

#### 生成工具 - protoc 插件
```
# protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3
protoc --version

# protoc-gen-go
https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.26.0
protoc-gen-go --version

# protoc-gen-dart
dart pub global activate protoc_plugin
```

#### proto 文件编译 - golang
> 项目中使用的库为：google.golang.org/protobuf
```
./build_go.sh
```

#### proto 文件编译 - dart
> 项目中使用的库为：protobuf
```
./build_dart.sh
```

#### proto 文件编译 - js
> 项目中使用的库为：google-protobuf
```
./build_js.sh
```
