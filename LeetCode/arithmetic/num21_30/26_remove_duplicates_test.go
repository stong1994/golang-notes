package num21_30

import "testing"

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

func removeDuplicates(nums []int) int {
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
