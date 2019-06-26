package num_1_10

import (
	"fmt"
	"sort"
	"testing"
)

func TestNum4(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2}
	target := float64(2)
	result := findMedianSortedArrays(num1, num2)
	if target == result {
		fmt.Println("true")
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > 0 {
		for _, v := range nums1 {
			nums2 = append(nums2, v)
		}
	}
	sort.Ints(nums2)
	len := len(nums2)
	if len%2 != 0 {
		return float64(nums2[len/2])
	} else {
		mid := len / 2
		return (float64(nums2[mid-1]) + float64(nums2[mid])) / 2
	}
	return 0.0
}
