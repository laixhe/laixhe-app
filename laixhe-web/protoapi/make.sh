
# 编译生成访问类文件
protoc --js_out=import_style=commonjs,binary:. *.proto
