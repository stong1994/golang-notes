package main

import (
	"fmt"
	"time"
)

func main() {
	dataCh := make(chan int)
	done := make(chan struct{})
	go rightFunc1(dataCh, done)
	dataCh <- 1
	<-done
	close(dataCh)
}

// 以下是错误代码，执行后会发现超时无效，原因是select会随机选择满足条件的分支，既然data的信道接收到了值，那么就会选择该信道，从而略过了超时分支
func wrongFunc(data chan int, done chan struct{}) {
	select {
	case d := <-data:
		time.Sleep(time.Second * 2)
		fmt.Println("get data, data is", d)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
	close(done)
}

func rightFunc1(data chan int, done chan struct{}) {
	outCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		close(outCh)
	}()
	for {
		select {
		case d := <-data:
			time.Sleep(time.Second * 2)
			fmt.Println("get data, data is", d)
			goto end
		case <-outCh:
			fmt.Println("timeout")
			goto end
		}
	}
end:
	close(done)
}

//
func rightFunc2(data chan int, done chan struct{}) {
	outCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		outCh <- struct{}{}
	}()
	for {
		select {
		case d := <-data:
			time.Sleep(time.Second * 2)
			fmt.Println("get data, data is", d)
			goto end
		case <-outCh:
			fmt.Println("timeout")
			goto end
		}
	}
end:
	close(done)
}
