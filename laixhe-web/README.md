## web

#### 安装 js 所需要的库
> 前置条件：安装 node.js

```
npm install -g require      # 对库文件的引用库
npm install -g browserify   # 用来打包成前端使用的 js 文件
npm install google-protobuf
```

#### 编译 proto 生成 js 文件
```
protoc --js_out=import_style=commonjs,binary:. xxx.proto

# 可使用 protoapi/make.sh 写好的脚本执行
```

#### 编译 proto 生成 js 的文件为前端 html使用的 js
```
# 可使用 protoapi/make_js.sh 写好的脚本执行
```

#### proto 的 js 相关 api
```
.serializeBinary()      // 序列化 (返回 uint8 数组)
.deserializeBinary(bin) // 反序列化 (静态方法)
.toObject()             // 打印数据 
```

##### nginx 配置

```
server {
    listen        80;
    server_name   localhost;

    location / {
        root    "/xxx/laixhe-web";
        index   index.html;
    }
}
```
