package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1e9; i++ {
		wg.Add(1)
		go func() {
			c := closure()
			for i := 0; i < 1e5; i++ {
				c()
			}
		}()
	}
}

func closure() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a = b
		b = c + b
		return a + b
	}
}
