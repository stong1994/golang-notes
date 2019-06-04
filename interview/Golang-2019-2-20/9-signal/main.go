package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go signalListen()
	for {
		select {
		default:
			fmt.Println("emm")
			time.Sleep(time.Second)
		}
	}
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGKILL) // TODO 在windows杀死进程和没有打印预果
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}
