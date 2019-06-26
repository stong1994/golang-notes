package main

import (
	"fmt"
	"testing"
	"time"
)

/*
	执行 TestGoroutineLeak1不会报错，空的结构体发送到通道中，因为通道有缓冲，所以没有被接收也可以等待被GC回收
	执行 TestGoroutineLeak2会报错，fatal error: all goroutines are asleep - deadlock! 通道没有缓冲，需要有对象来接收，但是没有。隐私报错
	执行 TestGoroutineLeak3会报错，fatal error: all goroutines are asleep - deadlock! 通道只有一个缓冲，而发送了两个值，通道只能接收一个。
	执行 TestGoroutineLeak4不会报错，打印结果为1，2. 通道只有一个缓冲，而发送了两个值，通道只能接收一个,第二个由于没有接收者而一直阻塞，导致goroutine泄露。
*/
func TestGoroutineLeak1(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
}

func TestGoroutineLeak2(t *testing.T) {
	ch := make(chan struct{})
	ch <- struct{}{}
}

func TestGoroutineLeak3(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	ch <- struct{}{}
}

func TestGoroutineLeak4(t *testing.T) {
	ch := make(chan struct{}, 1)
	go func() {
		fmt.Println(1)
		ch <- struct{}{}
		fmt.Println(2)
		ch <- struct{}{}
		fmt.Println(3)
	}()
	time.Sleep(time.Second)
}
