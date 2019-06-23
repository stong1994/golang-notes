package main

import (
	"fmt"
	"time"
)

// 寻找最大子数组

// 1. 分治思想
//   如果最大子数组含有中间的元素,则只需要从中间元素开始向左右两边遍历即可找到最大子数组
//   如果最大子数组在左半边或者右半边,则继续按照上边的思路展开计算,直到子数组长度为2
func main() {
	arr := []int{3, 10, -8, -6, 12, -7, 9, 4, -10, 10}
	start := time.Now().UnixNano()
	fmt.Println(findMaxSubArr(arr))
	v1End := time.Now().UnixNano()
	fmt.Printf("v1 cost %d 纳秒\n", v1End - start)
	fmt.Println(findMaxSubArr_v2(arr))
	v2End := time.Now().UnixNano()
	fmt.Printf("v2 cost %d 纳秒\n", v2End - v1End)

	// result:
	// 18 [12 -7 9 4]
	// v1 cost 56567 纳秒
	// 18 [12 -7 9 4]
	// v2 cost 4616 纳秒
}

func findMaxSubArr(arr []int) (int, []int) {
	if len(arr) == 0 {
		return 0, arr
	}
	if len(arr) == 1 {
		return arr[0], arr
	}
	var (
		maxArr []int
		maxSum int
	)
	midSum, midArr := includeMidElement(arr)
	leftSum, leftArr := findMaxSubArr(arr[:len(arr)/2-1])
	rightSum, rightArr := findMaxSubArr(arr[len(arr)/2+1:])
	if midSum > leftSum {
		maxArr = midArr
		maxSum = midSum
	}else {
		maxArr = leftArr
		maxSum = leftSum
	}
	if maxSum < rightSum {
		maxSum  = rightSum
		maxArr = rightArr
	}
	return maxSum, maxArr
}

func includeMidElement(arr []int) (int, []int) {
	mid := len(arr) / 2
	sum := 0
	left, right := mid, mid
	sumMax := 0

	for i:= mid; i >= 0; i -- {
		sum += arr[i]
		if sum > sumMax { // todo 如果一个数组和子数组的元素之和相等,这里只有>号,那么取子数组,如果有等号,取长度长的数组,下边同理
			left = i
			sumMax = sum
		}
	}
	sum = sumMax
	for i:= mid+1; i < len(arr); i++ {
		sum += arr[i]
		if sum > sumMax { // todo
			right = i
			sumMax = sum
		}
	}
	return sumMax, arr[left:right+1]
}

// 2. 非递归,只需线性时间即可
// 先遍历一遍,找到最大的数组arr[0:i],即最大子数组的右边界,然后再按照这个逻辑找到最大子数组的左边界
func findMaxSubArr_v2(arr []int) (int, []int) {
	if len(arr) <= 0 {
		return 0, arr
	}
	if len(arr) == 1 {
		return arr[0], arr
	}

	sum, sumMax, left, right := 0, 0, 0, 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > sumMax {
			sumMax = sum
			right = i
		}
	}

	sum, sumMax = 0, 0
	for i := right; i >= 0; i-- {
		sum += arr[i]
		if sum > sumMax {
			sumMax = sum
			left = i
		}
	}
	return sumMax, arr[left: right+1]
}