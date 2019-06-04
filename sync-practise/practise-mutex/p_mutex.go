package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁定的不是某个参数或者对象,而是锁本身

func main() {
	l := sync.Mutex{}
	start := time.Now().Second()

	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("锁定两秒")
		time.Sleep(time.Second * 2)
	}()

	fmt.Println("等待0.5秒,确保协程运行")
	time.Sleep(time.Second / 2)

	// 尝试获取锁
	l.Lock()
	defer l.Unlock()
	dur := time.Now().Second() - start
	fmt.Println("花费时间为", dur, "s")
}
