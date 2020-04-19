package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			"test1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"test2",
			[]int{9, 100, 6, 2, 5, 4, 8, 3, 7, 9},
			72,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := maxArea(v.arg)
			if got != v.want {
				t.Errorf("want %d got %d", v.want, got)
			}
		})
	}
}

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
*/
// 观察题目，面积的大小取决于两个索引所在的值的较小的那个。于是在符合条件的情况下移动较小的指针，遍历所有元素，即可找到最大的面积
// 双指针法：从两端开始设置两个指针，移动较短的那个指针，直到找到面积更大的，然后再比较两个指针，再移动较短的那个指针。直到两个指针重合则代表遍历完，返回最大的面积
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	i, j := 0, len(height)-1

	maxA := area(i, j, height)
	for i < j {
		if height[i] < height[j] {
			i++
			nowArea, ok := biggerArea(maxA, i, j, height)
			if ok {
				maxA = nowArea
			}
			continue
		}
		j--
		nowArea, ok := biggerArea(maxA, i, j, height)
		if ok {
			maxA = nowArea
		}
	}
	return maxA
}

func biggerArea(rawArea int, i, j int, height []int) (int, bool) {
	nowArea := area(i, j, height)
	if nowArea > rawArea {
		return nowArea, true
	}
	return rawArea, false
}

func area(i, j int, height []int) int {
	if height[i] < height[j] {
		return height[i] * (j - i)
	}
	return height[j] * (j - i)
}