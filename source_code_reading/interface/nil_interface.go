package main

import (
	"fmt"
	"reflect"
)

func main() {
	var iface interface{}
	fmt.Println(iface == nil) // true

	type animal struct{}
	var i interface{} = new(animal)
	fmt.Println(i == nil) // false
	v := reflect.ValueOf(i)
	if v.IsValid() {
		fmt.Println(v.IsNil())
	}
}
