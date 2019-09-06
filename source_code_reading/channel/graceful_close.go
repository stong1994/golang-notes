package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	ManySend2ManyRecvCloseBySend()
}

/*
1. 如果一个sender对应多个receiver，直接在sender处关闭channel
2. 如果多个sender对应一个receiver，额外声明一个channel ch1作为通信，receiver处关闭channel后，关闭ch1，sender接收到ch1关闭的信号后，停止发送
3. 如果多个sender对应多个receiver，上一个方法会receiver多次关闭channel，造成panic。声明一个缓存为1的channel ch2，当receiver关闭channel时，向ch2发送数据，ch2接收到数据后，关闭ch1，sender接收到ch1关闭的信号后，停止发送。
这种做法，由于缓冲的存在，只接收一次，但可以多次发送。但是传统上，我们在sender处进行关闭channel（实际上，需要根据使用场景判断在什么地方关闭channel，也有可能两处都有可能进行关闭）。
其实2,3并没有主动关闭我们要关闭的channel，而是通过ch1来停止goroutine对channel的发送和接收，这样GC会自动回收channel
*/

func ManySend2OneRecv() {
	ch := make(chan int, 10)
	closeCh := make(chan struct{})
	// sender
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-closeCh:
				return
			default:
				ch <- i
			}
		}(i)
	}
	// receiver
	go func() {
		for {
			select {
			case i := <-ch:
				if i == 5 {
					fmt.Println("close channel")
					closeCh <- struct{}{}
					return
				}
				fmt.Println("receiver", i)
			}
		}
	}()
	<-closeCh
	time.Sleep(time.Second)
}

/*
输出：
receiver 0
receiver 1
receiver 2
receiver 6
receiver 3
receiver 4
close channel
*/

func ManySend2ManyRecvCloseByRecv() {
	ch := make(chan int, 10)
	closeRecvCh := make(chan struct{}, 1)
	closeCh := make(chan struct{})

	go func() {
		<-closeRecvCh
		close(closeCh)
	}()
	// sender
	var (
		sendN int32
		recvN int32
	)
	for i := 0; i < 100; i++ {
		go func(i int) {
			select {
			case <-closeCh:
				return
			default:
				time.Sleep(time.Duration(MathRandNum(100)) * time.Millisecond)
				ch <- i
				atomic.AddInt32(&sendN, 1)
				fmt.Println("send", i)
			}
		}(i)
	}
	// receiver
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case i := <-ch:
					atomic.AddInt32(&recvN, 1)
					if i&1 == 1 {
						fmt.Println("close channel")
						select {
						case closeRecvCh <- struct{}{}: // 如果不用select的方式，这里可能会阻塞住。
						default:
						}
						return
					}
				case <-closeCh:
					return
				}
			}
		}()
	}

	<-closeCh
	time.Sleep(time.Second * 2)
	fmt.Println("send num", atomic.LoadInt32(&sendN))
	fmt.Println("recv num", atomic.LoadInt32(&recvN))
}

func ManySend2ManyRecvCloseBySend() {
	ch := make(chan int, 10)
	closeRecvCh := make(chan struct{}, 1)
	closeCh := make(chan struct{})

	go func() {
		<-closeRecvCh
		close(closeCh)
	}()
	// sender
	var (
		sendN int32
		recvN int32
	)
	for i := 0; i < 100; i++ {
		go func(i int) {
			if i&1 == 1 && i < 30 {
				for {
					select {
					case closeRecvCh <- struct{}{}: // 如果不用select的方式，这里可能会阻塞住。
					default:
					}
					return
				}
			}
			select { // 同理，如果不用select，可能channel已经停止接收了，但是我们还在发送
			case <-closeCh:
				return
			default:
				ch <- i
				atomic.AddInt32(&sendN, 1)
				fmt.Println("send", i)
			}

		}(i)
	}
	// receiver
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case _, ok := <-ch:
					if ok {
						atomic.AddInt32(&recvN, 1)
					}
				}
			}
		}()
	}

	<-closeCh
	time.Sleep(time.Second * 2)
	fmt.Println("send num", atomic.LoadInt32(&sendN))
	fmt.Println("recv num", atomic.LoadInt32(&recvN))
}

func MathRandNum(max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max)
}
