package main

import "fmt"

/*
1.Given an array of ints = [6, 4, -3, 5, -2, -1, 0, 1, -9],
implement a function in any language to move all positive numbers to the left,
and move all negative numbers to the right.
Try your best to make its time complexity to O(n), and space complexity to O(1).
 */

func main() {
	data := []int{6, 4, -3, 5, -2, -1, 0, 1, -9}
	fmt.Println(move(data))
}

func move(data []int) []int {
	posArr := []int{}
	negArr := []int{}
	zeroArr := []int{}

	for _, v := range data {
		if v == 0 {
			zeroArr = append(zeroArr, v)
			continue
		}
		if v > 0 {
			posArr = append(posArr, v)
			continue
		}
		if v < 0 {
			negArr = append(negArr, v)
		}
	}
	data = append(posArr, zeroArr...)
	data = append(data, negArr...)
	return data
}
