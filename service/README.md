```shell
goctl rpc -o video.proto
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```