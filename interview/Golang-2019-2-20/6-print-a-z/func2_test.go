package main

import (
	"context"
	"fmt"
	"testing"
)

// 安全的关闭goroutine
func TestFun2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := generate(ctx)
	for n := range gen {
		fmt.Println(n)
		if n >= 10 {
			break
		}
	}
}

// 生成1-10
func generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
