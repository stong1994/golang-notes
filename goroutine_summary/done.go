package main

import (
	"fmt"
	"time"
)

// 用channel来通知一个消息，如果仅仅是通知而不用传递信息的话，通常会done<-val，
// 然后用for-select来接受信息，但实际上并不用done<-val,
// 直接关闭channel，for-select中的<-done会接收到默认值，从而达到传递消息的效果
// 并且能够造成一处关闭,多个地方接收到通知的效果.
// TODO 思考与context是否有相同的作用?
func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(time.Second)
		// Do not need do it, after close() executed, <-done will get a default value
		//done <- struct{}{}
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("go1 closed")
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			fmt.Println("go2 closed")
			done = nil // set to nil, select不会再进入此分支
		case <-time.After(time.Second*2):
			fmt.Println("ended")
			return
		}
	}
}