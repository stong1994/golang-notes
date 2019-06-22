package main

import (
	"fmt"
)

func main() {
	testArr := []int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}
	quickSort(testArr, 0, len(testArr)-1)
	fmt.Println(testArr)
}

/**
快速排序：分治法+递归实现
随意去一个值A，将比A大的放在A的右边，比A小的放在A的左边
*/

func quickSort(arr []int, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}

	quickSort(arr, left, flag-1)
	quickSort(arr, flag+1, right)
}
