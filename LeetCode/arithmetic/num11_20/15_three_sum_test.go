package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := [...]struct {
		name string
		args []int
		want [][]int
	}{
		{
			"standard",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}},
		},
		{
			"four-zero",
			[]int{0, 0, 0, 0},
			[][]int{[]int{0, 0, 0}},
		},
		{
			"test3",
			[]int{-2,0,1,1,2},
			[][]int{[]int{-2, 0, 2}, []int{-2, 1, 1}},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := threeSum2(v.args)
			if len(got) != len(v.want) {
				t.Fatalf("want %v got %v", v.want, got)
			}
			for _, w := range v.want {
				findSame := false
				for _, g := range got {
					if judgeSame(w, g) {
						findSame = true
						break
					}
				}
				if !findSame {
					t.Fatalf("want %v got %v", v.want, got)
				}
			}
		})
	}
}

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
// 直接遍历的话会报超时错误
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			for l := j + 1; l < len(nums); l++ {
				if a+nums[j]+nums[l] == 0 {
					arr := []int{a, nums[j], nums[l]}
					sort.Ints(arr)
					res = append(res, arr)
				}
			}
		}
	}
	// 去重
	res2 := [][]int{}
	for _, m := range res {
		exist := false
		for _, n := range res2 {
			if judgeSame(m, n) {
				exist = true
			}
		}
		if !exist {
			res2 = append(res2, m)
		}
	}
	return res2
}

func judgeSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 先对原数据进行排序,然后遍历一个元素,和两个指针,一个从前往后遍历,另一个从后往前遍历,找到符合条件的就添加.
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := range nums {
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := len(nums) - 1
			for l < r {
				s := nums[i] + nums[l] + nums[r]
				if s == 0 {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					l += 1
					r -= 1
					for {
						if l < r && nums[l] == nums[l-1] {
							l += 1
						} else {
							break
						}
					}
					for {
						if l < r && nums[r] == nums[r+1] {
							r -= 1
						} else {
							break
						}
					}
				} else if s < 0 {
					l += 1
				} else if s > 0 {
					r -= 1
				}
			}
		}
	}
	return res
}

// 切片在遍历的时候,更改切片,会影响遍历吗
// 由结果可知,for循环开始时,会对arr进行复制,我们遍历的是复制后的arr,for循环中的与i相关的arr为复制后的arr,其他为复制前的arr.
// 结论,在遍历的时候更改切片,循环会按照初始的切片进行循环,而我们对切片做的操作会作用在for循环外的切片上
func TestChangeWhenItation(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range arr {
		fmt.Println("start:", arr)
		fmt.Println("i:", i, "len:", len(arr))
		arr = arr[:i]
		fmt.Println("end", arr)
	}
	fmt.Println(arr)
}

/*
start: [1 2 3 4 5 6 7 8 9 10]
i: 0 len: 10
end []
start: []
i: 1 len: 0
end [1]
start: [1]
i: 2 len: 1
end [1 2]
start: [1 2]
i: 3 len: 2
end [1 2 3]
start: [1 2 3]
i: 4 len: 3
end [1 2 3 4]
start: [1 2 3 4]
i: 5 len: 4
end [1 2 3 4 5]
start: [1 2 3 4 5]
i: 6 len: 5
end [1 2 3 4 5 6]
start: [1 2 3 4 5 6]
i: 7 len: 6
end [1 2 3 4 5 6 7]
start: [1 2 3 4 5 6 7]
i: 8 len: 7
end [1 2 3 4 5 6 7 8]
start: [1 2 3 4 5 6 7 8]
i: 9 len: 8
end [1 2 3 4 5 6 7 8 9]
[1 2 3 4 5 6 7 8 9]
 */
