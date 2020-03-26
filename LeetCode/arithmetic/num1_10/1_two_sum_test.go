package num_1_10

import (
	"testing"
)

func TestNum1(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			"test1",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			get := twoSum(v.nums, v.target)
			if v.want == nil || len(v.want) == 0 {
				if get != nil {
					t.Fatal("want nil, but get something")
				}
				return
			}
			if len(get) != 2 {
				t.Fatalf("get.len should be 2 but got %d", len(get))
			}
			if get[0] == v.want[0] && get[1] == v.want[1] {
				return
			}
			if get[1] == v.want[0] && get[0] == v.want[1] {
				return
			}
			t.Fatalf("want %v but got %v", v.want, get)
		})
	}
}

func foolish(nums []int, target int) []int {
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

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // key 元素 value 下标
	for i1, v := range nums {
		if i2, exist := m[target-v]; exist {
			return []int{i2, i1}
		}
		m[v] = i1
	}
	return nil
}
