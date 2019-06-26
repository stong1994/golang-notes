package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// 初始化pool为容量为3，pool中的函数返回值为5
	p := NewPool(func() interface{} {
		return 5
	}, 3) // 改变count的值，发现程序运行时间会不一样。

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				fmt.Println("borrowing num, waiting...")
				v := Borrow().(int) // 初始化的时候定义返回值为int类型，所以这里可以进行类型强制转换
				fmt.Println("borrowed num", v)
				time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
				fmt.Println("returnint num", v)
				Return(v)
				fmt.Println("returned num", v)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
