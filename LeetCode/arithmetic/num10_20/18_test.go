package num10_20

import (
	"fmt"
	"sort"
	"testing"
)

func TestNum18(t *testing.T) {
	tests := [...]struct{
		name string
		nums []int
		target int
		want [][]int
	}{
		{
			"standard",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{[]int{-1,  0, 0, 1}, []int{-2, -1, 1, 2}, []int{-2,  0, 0, 2}},
		},

	}


	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := fourSum(v.nums, v.target)
			fmt.Println(got)
			for _, g := range v.want {
				var sameN bool
				for _, w := range got {
					same := judge(g, w)
					if same {
						sameN = true
						break
					}
				}
				if !sameN {
					t.Errorf("want %v got %v", v.want, got)
					t.FailNow()
				}
			}
		})
	}


}

var judge = func(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, subA := range a {
		if subA != b[i] {
			return false
		}
	}
	return true
}

/*
1. 对nums中的元素两两组合,并记录其和
2. 遍历map,根据target来得出另一个值,并判断出是否有重复的index,得到正确的结果
 */
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	type data struct {
		sum int
		n1 int
		n2 int
	}
	combine := make([]*data, 0, (len(nums)*2-1)/2)
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			combine = append(combine, &data{nums[i]+nums[j], i, j})
		}
	}


	var result [][]int

	for _, v := range combine {
		another := target - v.sum
		for _, n := range combine {
			if n.sum == another {
				if v.n1 != n.n1 && v.n1 != n.n2 && v.n2 != n.n1 && v.n2 != n.n2 {
					arr := sortArr(nums[v.n1], nums[v.n2], nums[n.n1], nums[n.n2])
					if len(result) == 0 {
						result = append(result, arr)
						continue
					}

					for _, m := range result {
						if judge(arr, m) {
							goto next
						}
					}
						result = append(result, arr)
				}
			}
		next:
		}
	}

	return result
}

func sortArr(a, b, c, d int) []int {
	arr := []int{a, b, c, d}
	sort.Ints(arr)
	return arr
}