package main

import "fmt"

func main() {
	arr := []int{9, 4, 5, 1, 22, 8, 4, 7, 0, 2}
	bubbleSort(arr, len(arr))
	fmt.Println(arr)
}

/**
冒泡排序：第一次循环，确定最后一位为最大值；第二次循环，确定倒数第二位为倒数第二大的值；以此类推。。。
时间复杂度 O(n^2)
*/
func bubbleSort(arr []int, len int) {
	if len == 1 {
		return
	}
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	bubbleSort(arr, len-1)
}
