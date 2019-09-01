package main

import "fmt"

func main()  {
	ch1 := make(chan bool)
	go func() {
		ch1 <- true
	}()
	select {
	case <-ch1:
		fmt.Println("hello")
	}

}
