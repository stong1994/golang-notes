package main

import (
	"fmt"
	"testing"
)

func TestPrintA2Z_Fun2(t *testing.T) {
	a := 'a'
	z := 'z'
	oddChan := make(chan int)
	evenChan := make(chan int)
	doneChan := make(chan struct{})
	go func() {
		for v := range oddChan {
			fmt.Print(string(v))
			if v == int(z) {
				doneChan <- struct{}{}
				return
			}
			evenChan <- v + 1
		}
	}()

	go func() {
		for v := range evenChan {
			fmt.Print(string(v))
			if v == int(z) {
				doneChan <- struct{}{}
				return
			}
			oddChan <- v + 1
		}
	}()
	oddChan <- int(a)
	<-doneChan
	close(oddChan)
	close(evenChan)
	close(doneChan)
}
