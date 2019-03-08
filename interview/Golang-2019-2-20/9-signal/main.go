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
	signal.Notify(c, syscall.SIGKILL) // TODO 在windows杀死进程和没有打印预期结果。。。，需要在linux上实验下。另外信号相较于通道的好处待研究。
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}
