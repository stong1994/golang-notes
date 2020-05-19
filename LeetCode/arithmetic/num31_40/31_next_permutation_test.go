package num31_40

import (
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
		{
			"test4",
			[]int{1, 3, 2},
			[]int{2, 1, 3},
		},
		{
			"test5",
			[]int{1, 3, 8, 5, 4},
			[]int{1, 4, 3, 5, 8},
		},
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

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
从数组的最右端开始扫描，找到从右向左递增的递增序列，158476531的数组序列，可以知道从右边往左边扫描的过程中发现13657是递增的序列，而到4的时候则不是递增的序列了，
因为4大于了7，所以这个时候循环结束，循环变量记录了4这个位置，
在后面递增的序列中从右往左找到第一个比4大的位置可以知道是13657中的5对应的位置，这个时候需要将4的位置与5的位置进行互换，
因为调换元素之后那么剩下来的从左到右是递增的，所以需要进行翻转，应该是从4这个位置后面进行翻转，这样形成的数字序列才是下一个更大的排列

作者：ba-xiang
链接：https://leetcode-cn.com/problems/next-permutation/solution/go-shuang-100-by-ba-xiang/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func nextPermutation(nums []int)  {
	var i = len(nums) - 2
	// 从右到左，找到第一个不符合递增规律的值
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		// 从右到左，找到第一个大于nums[i]的值
		var j = len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 颠倒i+1及其右边的值
	if i >= -1 {
		reverse(nums, i+1)
	}
}

func reverse(nums []int, start int)  {
	var i, j = start, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}