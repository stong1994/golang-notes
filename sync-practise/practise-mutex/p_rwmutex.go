package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁并不是一个加锁操作和一个解锁操作，而是读加锁操作，读解锁操作和写加锁操作和写解锁操作，也就是锁定的不是一块代码，也不是一个对象/变量，而是锁本身
// 一个锁既用到了读操作，也用到了写操作，那么在写操作时，其他的读操作和写操作是不允许的；在读操作时，不允许写操作，但允许读操作。

type rw bool // 读为true，写为false
var r rw = true
var w rw = false

func main() {
	l := sync.RWMutex{}
	ch := make(chan rw, 10)
	start := time.Now().Unix()
	defer func() {
		end := time.Now().Unix()
		dur := end - start
		fmt.Printf("程序总执行时间为 %ds", dur)
	}()

	for i := 0; i < 5; i++ {
		go func(i int) {
			l.RLock()
			defer l.RUnlock()
			fmt.Printf("读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为%d。。。\n", i)
			time.Sleep(time.Second * 2)
			ch <- r
			fmt.Printf("读操作完成，index为%d\n", i)
			fmt.Println("")
		}(i)
	}

	for i := 5; i < 9; i++ {
		go func(i int) {
			l.Lock()
			defer l.Unlock()
			fmt.Printf("写操作中，执行时间为3s，不允许其他的读操作和写操作, index为%d。。。\n", i)
			time.Sleep(time.Second * 3)
			ch <- w
			fmt.Printf("写操作完成, index为%d\n", i)
			fmt.Println("")
		}(i)
	}

	for i := 0; i < 9; i++ {
		if res := <-ch; res {
			// fmt.Println("正在进行读操作。。。")
		} else {
			// fmt.Println("正在进行写操作。。。")
		}
	}
}

// 结果
/*
读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为0。。。
读操作完成，index为0

写操作中，执行时间为3s，不允许其他的读操作和写操作, index为8。。。
写操作完成, index为8

读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为1。。。
读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为3。。。
读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为4。。。
读操作中，执行时间为2s，允许其他的读操作，但不允许写操作，index为2。。。
读操作完成，index为1

读操作完成，index为3

读操作完成，index为4

读操作完成，index为2

写操作中，执行时间为3s，不允许其他的读操作和写操作, index为5。。。
写操作完成, index为5

写操作中，执行时间为3s，不允许其他的读操作和写操作, index为6。。。
写操作完成, index为6

写操作中，执行时间为3s，不允许其他的读操作和写操作, index为7。。。
写操作完成, index为7

程序总执行时间为 16s
*/
