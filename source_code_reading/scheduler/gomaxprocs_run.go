package main

import (
	"fmt"
	"runtime"
	"time"
)

// 以下程序用于不会输出 x
/*
对于 Go 语言中运行时间过长的 goroutine，Go scheduler 有一个后台线程在持续监控，
一旦发现 goroutine 运行超过 10 ms，会设置 goroutine 的“抢占标志位”，之后调度器会处理。
但是设置标志位的时机只有在函数“序言”部分，对于没有函数调用的就没有办法了。
 */
func main() {
	x := 0
	n := runtime.GOMAXPROCS(0)
	for i := 0; i < n; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
}

/*
解决办法
1. 将n减小
2. 协程中增加 time.Sleep(time.Second)让主协程有机会被调用
输出结果为0: 内存重排 (https://juejin.im/post/5d06e71df265da1bc64bc1c8)
 */