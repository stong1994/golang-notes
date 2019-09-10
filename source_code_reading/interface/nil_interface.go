package main

import "fmt"

func main() {
	var iface interface{}
	fmt.Println(iface == nil) // true

	type animal struct{}
	var i interface{} = new(animal)
	fmt.Println(i == nil) // false
}
