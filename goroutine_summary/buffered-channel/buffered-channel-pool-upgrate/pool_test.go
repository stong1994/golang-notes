package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// 初始化pool为容量为3，pool中的函数返回值为5
	p := NewPool(func() interface{} {
		return 5
	}, 4)

	var wg sync.WaitGroup
	var n int64
	var m int64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&n, 1)
				start := time.Now().UnixNano()

				/*p.Run(func(v interface{}) {
					time.Sleep(100 * time.Millisecond)
				})*/
				err := RunWithTimeout(func(v interface{}) {
					time.Sleep(100 * time.Millisecond)
				}, 200*time.Millisecond)
				if err != nil {
					fmt.Println(err)
					atomic.AddInt64(&m, 1)
				}
				now := time.Now().UnixNano()
				// 更新时间
				fmt.Println(atomic.LoadInt64(&n), "执行时间为", now-start)

			}
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
	fmt.Println("err num", m)
}
