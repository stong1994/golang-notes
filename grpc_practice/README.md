# 一个简单grpc练习

### 注意事项
0. 安装依赖库和插件：`go get -u github.com/golang/protobuf/{protoc-gen-go,proto} // 前者是 plugin；后者是 go 的依赖库`
1. go文件生成命令：`protoc --go_out=plugins=grpc:. proto/login.proto`  
注意输出目录的参数`plugins=grpc:.`,如果没有的话生成的文件不全。
2. 生成的.pb.go文件中有两个接口{serverName}Client与{serverName}Server,
这两个接口中的方法就是我们要实现的grpc接口。  
server端需要实现接口{serverName}Server并注册服务，  
client端需要调用生成的New{serverName}Client()函数，即可调用server端的接口

### 主要流程
> 服务端
```go
lis, err := net.Listen("tcp", addr) // 创建监听套接字
grpc.NewServer() // 创建服务端
pb.RegisterxxxxxServer() // 注册服务
s.Serve(lis) // 启动服务端
```
> 客户端
```go
conn, err := grpc.Dial(addr, opt...) // 创建连接
client := pb.NewxxxxxxClient(conn) // 创建客户端
```

### grpc参考资料：
[官方中文文档](https://doc.oschina.net/grpc?t=60133)\
[protobuf语法博客](https://segmentfault.com/a/1190000007917576#articleHeader11)