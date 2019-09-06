package main

import "fmt"

// 只是声明的channel为nil，用make构造的channel不为nil，关闭后的channel不为nil
func main() {
	var c1 chan struct{}
	fmt.Println(c1 == nil) // true
	c2 := make(chan int, 2)
	fmt.Println(c2 == nil) // false
	close(c2)
	fmt.Println(c2 == nil) // false
	c1 <- struct{}{}       // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("unreachable ")
}
