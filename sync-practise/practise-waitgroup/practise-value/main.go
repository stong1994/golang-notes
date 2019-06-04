package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	v := atomic.Value{}
	ch := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func(i int) {
			v.Store(i)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < 10; i++ {
		if _, ok := <-ch; ok {
			val := v.Load().(int)
			fmt.Println("val is ", val)
		}
	}
}
