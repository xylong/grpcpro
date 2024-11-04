grpc

### 安装grpc
```shell
go get -u google.golang.org/grpc
```

### 下载protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v28.3
加环境变量

### 安装protoc-gen-go
```shell
# 弃用
go get github.com/golang/protobuf/protoc-gen-go
# 替换
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
```


```shell
# 安装
go get github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
```

### 生成
```shell
protoc --go_out=pbfiles/ proto/Prod.proto

# --go_out=plugins=grpc 不支持
protoc --go_out=./services --go-grpc_out=./services proto/Prod.proto
```