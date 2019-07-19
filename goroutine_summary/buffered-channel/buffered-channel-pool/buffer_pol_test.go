package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_buffer_pool(t *testing.T) {
	noCh := make(chan struct{})
	defer close(noCh)
	go server()
	go client()
	for {
		select {
		case <-noCh:
			fmt.Println("not possible access there")
		case <-time.After(time.Second * 10):
			return
		}
	}
}
