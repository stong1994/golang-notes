package main

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	arr := make([]int, 3, 3)
	//arr = []int{1, 2, 3}

	arr2 := append(arr, arr[3:]...)
	fmt.Println(arr2)
}
