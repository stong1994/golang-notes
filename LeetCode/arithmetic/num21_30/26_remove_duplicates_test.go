package num21_30

import (
	"testing"
)

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"test1",
			[]int{1, 1, 2},
			[]int{1, 2},
		},
		{
			"test1",
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			l := removeDuplicates(v.nums)
			if len(v.want) != l {
				t.Fatalf("want len is %d got %d", len(v.want), l)
			}
			for i := range v.want {
				if v.want[i] != v.nums[i] {
					t.Fatalf("want len is %v got %v", v.want, v.nums)
				}
			}
		})
	}
}

func foolish(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numsMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		} else {
			numsMap[nums[i]] = struct{}{}
		}
	}
	return len(nums)
}

func mine(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tmpV := nums[0]
	tmpI := 1
	count := 0
	tmpL := len(nums)
	for {
		if count >= tmpL-1 {
			break
		}
		if nums[tmpI] == tmpV {
			nums = append(nums[:tmpI-1], nums[tmpI:]...)
		}else {
			tmpV = nums[tmpI]
			tmpI++
		}
		count++
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2{
		return n
	}
	l, r := 0, 1
	for r < n{
		if nums[l] < nums[r]{
			l++
			nums[l] = nums[r]
		}
		r++
	}
	nums = nums[:l+1]
	return l + 1
}

//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n < 2{
//		return n
//	}
//	i := 1
//	tmpV := nums[0]
//	for _, v := range nums {
//		if tmpV != v {
//			tmpV = v
//			i++
//		}
//	}
//	return i
//}