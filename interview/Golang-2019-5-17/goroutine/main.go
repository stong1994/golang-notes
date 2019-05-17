package goroutine

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

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
