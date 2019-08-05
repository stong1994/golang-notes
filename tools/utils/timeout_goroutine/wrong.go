package main

import (
	"fmt"
	"time"
)

// 记录一次错误的使用

func main() {
	dataCh := make(chan int)
	done := make(chan struct{})
	go wrongFunc(dataCh, done)
	dataCh <- 1
	<-done
	close(dataCh)
}

// 以下执行后会发现超时无效，原因是select会随机选择满足条件的分支，既然data的信道接收到了值，那么就会选择该信道，从而略过了超时分支
// 该代码的作用是：如果data的信道延迟接收，那么最多等待1s，1s后data信道没有接收到数据，则timeout
// 因此，应该理解超时函数作用于要等待的函数接收值，而不是该函数本身
func wrongFunc(data chan int, done chan struct{}) {
	select {
	case d := <-data:
		time.Sleep(time.Second * 2) // 模拟函数运行
		fmt.Println(d)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
	close(done)
}
