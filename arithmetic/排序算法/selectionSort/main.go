package main

import "fmt"

func main() {
	arr := []int{6, 2, 4, 8, 9, 1, 4, 0, 10, 5, 2}
	selectionSort(arr, 0)
	fmt.Println(arr)
}

/**
选择排序法：在未排序的切片中选取最小的值放在首位，然后在未排序的切片中选取最小的值放在第二位，以此类推。。。
*/
func selectionSort(arr []int, start int) {
	if start == len(arr) {
		return
	}
	minIdx := start
	minVal := arr[start]
	for i := start + 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minIdx, minVal = i, arr[i]
		}
	}
	arr[start], arr[minIdx] = arr[minIdx], arr[start]
	selectionSort(arr, start+1)
}
