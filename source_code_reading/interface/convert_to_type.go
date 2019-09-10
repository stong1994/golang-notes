package main

import "fmt"

func main() {
	type animal interface{}
	type cat struct{}
	var c animal = cat{}
	switch c.(type) {
	case cat:
		fmt.Println(c.(cat))
	}
	//cl, ok := c.(cat)
	//fmt.Println(cl, ok) // {} true
}
