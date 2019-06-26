package main

import (
	"fmt"
)

// 结论：无缓冲的通道与缓冲为0的通道效果是一样的。
func main() {
	buffered()
}

func buffered() {
	in := make(chan int, 1)
	done := make(chan struct{})
	times := 10

	go func() {
		for {
			select {
			case i := <-in:
				fmt.Println("接收", i)
				if i >= times-1 {
					done <- struct{}{}
				}
			}
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			fmt.Println("准备发送", i)
			in <- i
			fmt.Println("已经发送", i)
		}
	}()

	<-done
	close(done)
	close(in)
}
