package main

import "fmt"

func main() {
	arr := []int{1, 0, 5, 7, 8, 5, 3, 6, 9, 2, 54, 33, 66}
	//newArr := []int{}
	//insertionSort(arr, &newArr)
	//fmt.Println(newArr)
	method2(arr)
	fmt.Println(arr)

}

/**
插入排序法：取原切片old中第一个值作为新切片中的第一个值，然后遍历old，将每个元素按照条件插入到新切片中
时间复杂度：O（n^2）
实现方式有两种：1，新建一个切片；2：在原切片中交换元素
*/
// 1，新建一个切片
func insertionSort(old []int, new *[]int) {
	if len(*new) == len(old) {
		return
	}
	current := len(*new)
	*new = append(*new, old[current])
	sort(*new)

	insertionSort(old, new)
}

func sort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
	}
}

// 2， 在原切片中交换元素 (由于不用创建新的切片，不用进行插入操作，只需要交换操作，所以要较方法一速度快些)
// 每次只在未排序的切片中拿一个值，然后与已经排序过的切片进行比较，并插入。
func method2(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
