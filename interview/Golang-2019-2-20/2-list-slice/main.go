package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := [3]int{1, 2, 3}
	changeSlice(a)
	changeList(b)
	fmt.Println(a)
	fmt.Println(b)
}

func changeList(data [3]int) {
	data[0] = 0
}

func changeSlice(data []int) {
	data[0] = 0
}
