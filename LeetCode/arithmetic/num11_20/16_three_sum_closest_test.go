package main

import (
	"sort"
	"testing"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
*/
func TestThreeSumClosest(t *testing.T) {
	type arg struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		arg  arg
		want int
	}{
		{
			"test1",
			arg{[]int{-1, 2, 1, -4}, 1},
			2,
		},
		{
			"test2",
			arg{[]int{1, 1, 1, 0}, -100},
			2,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := threeSumClosest(v.arg.nums, v.arg.target)
			if got != v.want {
				t.Errorf("want %d got %d", v.want, got)
			}
		})
	}
}

// 笨法
func threeSumClosest2(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sum := make(map[int]struct{}, len(nums)*(len(nums)-1)*(len(nums)-2))
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum[nums[i]+nums[j]+nums[k]] = struct{}{}
			}
		}
	}
	close := 1 << 32
	total := 0
	for i := range sum {
		off := i - target
		if off == 0 {
			return i
		}
		if off < 0 {
			off = 0 - off
		}
		if off < close {
			total = i
			close = off
		}
	}
	return total
}

// 双指针
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	// 有关数组/切片，可以考虑排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	offset := 1 << 32
	closeSum := 0
	// 遍历两次，i从左到右遍历，同时i每移动一次，就从左右两边用双指针向中间逼近，得到最接近的值
	for i, v := range nums {
		start, end := 0, len(nums)-1
		tmp := 1 << 32
		tmpSum := 0
		for {
			if start == i {
				start++
			}
			if end == i {
				end--
			}
			if start >= end {
				break
			}
			sum := v + nums[start] + nums[end]
			if sum == target {
				return sum
			}
			off := sum - target
			if off < 0 {
				off = 0 - off
			}
			if off < tmp {
				tmp = off
				tmpSum = sum
			}
			if sum > target {
				end--
			} else {
				start++
			}
		}
		if tmp < offset {
			offset = tmp
			closeSum = tmpSum
		}
	}

	return closeSum
}
