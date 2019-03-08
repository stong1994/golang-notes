package main

import (
	"fmt"
	"math/rand"
)

func getRandom() *int {
	result := rand.Intn(100)
	return &result
}

func main() {
	randNum := getRandom()
	fmt.Println(*randNum)
}

/*
Go逃逸分析最基本的原则是：如果一个函数返回对一个变量的引用，那么它就会发生逃逸。
go build -gcflags -m main.go

# command-line-arguments
.\main.go:10:9: &result escapes to heap
.\main.go:9:2: moved to heap: result
.\main.go:15:14: *randNum escapes to heap
.\main.go:15:13: main ... argument does not escape

关于逃逸问题，请参考：https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653372198&idx=1&sn=7f52768337e4ee77e7276b39afcdc516&chksm=bce4df3c8b93562ad0f288ef54bbc15fb75bb2ac6b6ba38a42d5230b15d4ad0f582aa9ba1872&mpshare=1&scene=1&srcid=&key=e3738c51d3aaafb50afb3eaad4c8ea733b7f44b398bf0453c5497cc0f34b412a0124c68b917f359ff808f1183a9d3fd745870d78d8d87715c70cf4fcb98ae3327f32af5c51ff3dd0af6d43e4f7c4c76c&ascene=1&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060728&lang=zh_CN&pass_ticket=b7LZnV49n4%2FGYAMBmnhW1xjJfbSpfbVAwbNv06eB9T8Oi%2F37FB6AsR3giUor%2B5RD
*/
