package main

import (
	"fmt"
	"sync"
	"time"
)

//    c.L.Lock()
//    for !condition() {
//        c.Wait()
//    }
//    ... make use of condition ...
//    c.L.Unlock()

/* 注意事项：
1. 注意上方Wait()的使用条件是不满足条件，也就是符合条件的会被阻塞，直到它不符合条件才会通过
2. 判断条件用for来一直等待，而不是if直接判断
3. 修改全局条件变量和广播时，也应该加锁和解锁。
*/
func main() {
	l := sync.Mutex{}
	c := sync.NewCond(&l)
	flag := 5

	for i := 0; i < 10; i++ {
		go func(i int) {
			c.L.Lock()
			fmt.Printf("current i = %d, flag = %d, waitting... \n", i, flag)
			// flag为5，那么5到9会阻塞，逐次增大flag。。。
			for i >= flag {
				c.Wait()
				fmt.Printf("current i = %d, flag = %d, running... \n", i, flag)
			}
			fmt.Printf("current i = %d, flag = %d, ending \n", i, flag)
			c.L.Unlock()
		}(i)
	}

	time.Sleep(time.Second) // 等待协程跑完
	fmt.Println("wait goroutines running")

	// 先广播来查看所有阻塞的协程
	fmt.Println("broadcast first, all blocked goroutines")
	c.L.Lock()
	c.Broadcast()
	c.L.Unlock()
	time.Sleep(time.Second)

	// 修改条件，并通知一个符合条件的协程
	fmt.Println("signal...")
	c.L.Lock()
	flag += 2 // flag为7， 5和6会不满足条件，得到通知，然后随机运行一个协程
	c.Signal()
	c.L.Unlock()
	time.Sleep(time.Second)

	// 通知剩下不符合条件的协程运行
	fmt.Println("broadcast second")
	c.L.Lock()
	c.Broadcast()
	c.L.Unlock()
	time.Sleep(time.Second)

	// 修改条件，跑完剩下所有的协程
	fmt.Println("broadcast third")
	c.L.Lock()
	flag = 10
	c.Broadcast()
	c.L.Unlock()
	time.Sleep(time.Second * 10)
}
