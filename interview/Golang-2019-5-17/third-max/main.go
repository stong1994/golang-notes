package main

import "fmt"

// 找到第三大的元素,时间复杂度为O(n)
func main() {
	arr := []int{2, 5, 8, 1, 3, 5, 0, 11, 18, -1}

	biggst, second, third := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > biggst {
			third = second
			biggst, second = arr[i], biggst
			if third > second {
				third, second = second, third
			}
		} else if arr[i] > second {
			second, third = arr[i], second
		} else if arr[i] > third {
			third = arr[i]
		}
	}
	fmt.Println(biggst, second, third)

}
