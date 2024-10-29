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