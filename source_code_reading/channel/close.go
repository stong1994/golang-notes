package main

import (
	"fmt"
	"time"
)

func main() {
	recvqWhenClose()
}

// 关闭channel，会导致等待发送的队列中的goroutine panic
func sendqWhenClose() {
	ch := make(chan int)
	go func() {
		ch <- 1 // panic: send on closed channel
	}()
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
}

// 关闭channel，在等待接收的队列中，如果channel还有缓冲，返回缓存值和true，如果没有缓冲，返回零值和false
/*
输出
1 true
0 false
0 false
*/
func recvqWhenClose() {
	ch := make(chan int, 1)
	go func() {
		for {
			time.Sleep(time.Second)
			v, ok := <-ch
			fmt.Println(v, ok)
		}
	}()
	ch <- 1
	close(ch)
	fmt.Println(ch == nil)
	time.Sleep(time.Second * 3)
}
