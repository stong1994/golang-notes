package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("start")
	go func() {
		fmt.Println("hello")
	}()

	go func() {
		i := 0
		for {
			i++
		}
	}()
	select {

	}
}
