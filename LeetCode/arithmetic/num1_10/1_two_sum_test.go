package num_1_10

import (
	"fmt"
	"testing"
)

func TestNum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, v := range nums {
		for j, m := range nums {
			if i != j {
				if v+m == target {
					return append(result, i, j)
				}
			}
		}
	}
	return result
}
