package main

import "fmt"

// 寻找最大子数组

// 1. 分治思想
//   如果最大子数组含有中间的元素,则只需要从中间元素开始向左右两边遍历即可找到最大子数组
//   如果最大子数组在左半边或者右半边,则继续按照上边的思路展开计算,直到子数组长度为2
func main() {
	arr := []int{3, 10, -8, -6, 12, -7, 9, 4, -10, 8}
	fmt.Println(findMaxSubArr(arr))
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
	leftSum, leftArr := maxInSide(arr[:len(arr)/2-1])
	rightSum, rightArr := maxInSide(arr[:len(arr)/2+1])
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
	fmt.Println("mid arr", arr)
	mid := len(arr) / 2
	sum := 0
	left, right := mid, mid+1
	sumMax := 0

	for i:= mid; i >= 0; i -- {
		sum += arr[i]
		if sum >= sumMax { // todo
			left = i
			sumMax = sum
		}
	}
	sum = sumMax
	for i:= mid+1; i < len(arr); i++ {
		sum += arr[i]
		if sum > sumMax {
			right = i
			sumMax = sum
		}
	}
	fmt.Println("mid result", sumMax, arr[left:right+1])
	return sumMax, arr[left:right+1]
}

func maxInSide(arr []int) (int, []int) {
	fmt.Println("aside", arr)
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
	leftSum, leftArr := maxInSide(arr[:len(arr)/2-1])
	rightSum, rightArr := maxInSide(arr[len(arr)/2+1:])
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
	fmt.Println("aside result", maxSum, maxArr)
	return maxSum, maxArr
}