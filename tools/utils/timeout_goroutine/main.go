package main

import (
	"fmt"
	"time"
)

// 以下执行后会发现超时无效，原因是select会随机选择满足条件的分支，既然data的信道接收到了值，那么就会选择该信道，从而略过了超时分支
// 该代码的作用是：如果data的信道延迟接收，那么最多等待1s，1s后data信道没有接收到数据，则timeout
// 因此，应该理解超时函数作用于要等待的函数接收值，而不是该函数本身
func timeoutFunc(data chan int, done chan struct{}, f func(data int)) {
	select {
	case d := <-data:
		f(d)
	case <-time.After(time.Second): // 补充：直接使用time.After存在内存泄露的问题，如果在1s内函数返回，Timer不会立刻被回收
		fmt.Println("timeout") // 如果访问很频繁，就会造成大量的内存泄露
	} // 使用time.NewTimer()
	close(done)
}

func main() {
	dataCh := make(chan int)
	doneCh := make(chan struct{})

	go waitFunc(dataCh) // 模仿函数延迟

	var handleFunc = func(data int) {
		fmt.Println(data)
	}
	timeoutFunc(dataCh, doneCh, handleFunc)
	<-doneCh
}

func waitFunc(dataCh chan int) {
	time.Sleep(time.Second * 2)
	dataCh <- 1
}

// 正确用法
func timeoutFunc2(data chan int, done chan struct{}, f func(data int)) {
	idleDelay := time.NewTimer(time.Second)
	select {
	case d := <-data:
		f(d)
	case <-idleDelay.C:
		fmt.Println("timeout")
	}
	close(done)
}
