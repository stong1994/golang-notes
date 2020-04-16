package num31_40

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name  string
		param []int
		want  []int
	}{
		//{
		//	"test1",
		//	[]int{1, 2, 3},
		//	[]int{1, 3, 2},
		//},
		//{
		//	"test2",
		//	[]int{3, 2, 1},
		//	[]int{1, 2, 3},
		//},
		//{
		//	"test3",
		//	[]int{1, 1, 5},
		//	[]int{1, 5, 1},
		//},
		//{
		//	"test4",
		//	[]int{1, 3, 2},
		//	[]int{2, 1, 3},
		//},
		//{
		//	"test5",
		//	[]int{1, 3, 8, 5, 4},
		//	[]int{1, 4, 3, 5, 8},
		//},
		{
				"test6",
				[]int{4,2,0,2,3,2,0},
				[]int{4,2,0,3,0,2,2},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			nextPermutation(v.param)
			if !assert.IsEqual(v.param, v.want) {
				t.Fatalf("want %v got %v", v.want, v.param)
			}
		})
	}
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	l := len(nums)
	if nums[l-1] > nums[l-2] {
		nums[l-1], nums[l-2] = nums[l-2], nums[l-1]
		return
	}
	idxLeft, idxRight := findExchangeIndex(nums, l)
	if idxLeft == -1 {
		quickSort(nums, 0, l-1)
		return
	}
	fmt.Println(idxLeft, idxRight)
	nums[idxLeft], nums[idxRight] = nums[idxRight], nums[idxLeft]
	// sort left elem
	quickSort(nums, idxLeft+1, l-1)
}

// 1, 3,8, 5, 4    1, 4, 3, 5, 8

// 找到要替换的index
func findExchangeIndex(nums []int, l, r int) (int, int) {
	if l >= l {
		return -1, -1
	}

	for i := 
	for i := l - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				return j, i
			}
		}
	}
	return -1, -1
}

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
