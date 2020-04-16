package num21_30

import (
	"testing"
)

/*
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		{
			"test1",
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
		},
		{
			"test2",
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 3, 0, 4},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			nums := v.nums
			l := removeElement(nums, v.val)
			if l != len(v.want) {
				t.Fatalf("want %d got %d", len(v.want), l)
			}

			//if !assert.IsEqual(v.want, nums) {
			//	t.Fatalf("wnat %v got %v", v.want, nums)
			//}
		})
	}
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
