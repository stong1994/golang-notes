# go mock

## gomock
[参考资料](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/gomock)
### 步骤
1. 安装
```cassandraql
$ go get -u github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen
```
2. 生产mock文件
```cassandraql
mockgen -source=gomock/person.go -destination=./gomock/mock/male_mock.go -package=mock
```

- 全局变量可通过 GoStub 框架打桩
- 过程可通过 GoStub 框架打桩
- 函数可通过 GoStub 框架打桩
- interface 可通过 GoMock 框架打桩
- 方法可通过 monkey/gomonkey 来打桩

测试过程中发现monkey/gomonkey对于有些函数失效，原因在于这些函数由于太简单被编译器内联。因此在执行时需要带上禁止内联的参数。
`go test -gcflags=-l -v user_test.go user.go  -test.run TestUser_Get`

[参考资料](https://zhuanlan.zhihu.com/p/267341653)