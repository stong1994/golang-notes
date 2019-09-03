package main

import "fmt"

func main() {
	var c1 chan struct{}
	c1 <- struct{}{} // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("unreachable ")
}
