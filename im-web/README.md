#### 初始化
```
npm install
```

#### 将pd文件转为浏览器能用的js文件
```
npm run build
```

#### 使用 golang 打包静态文件为二进制文件
> 需要 go1.16 以上的版本

```
# 生成名为 im-web 的二进制文件
go build

# 本地调试
./im-web -debug=true
```
