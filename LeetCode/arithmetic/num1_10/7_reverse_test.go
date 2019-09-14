package num_1_10

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			"test1",
			123,
			321,
		}, {
			"test2",
			-123,
			-321,
		}, {
			"test3",
			120,
			21,
		},
		{
			"test4",
			1534236469,
			0,
		}, {
			"test5",
			-2147483648,
			0,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := reverse(v.arg)
			if got != v.want {
				t.Errorf("want %d, got %d", v.want, got)
			}
		})
	}
}

func reverse(x int) int {
	if x >= 1<<31 || x < -(1<<31) {
		return 0
	}

	positive := true
	if x < 0 {
		positive = false
		x = 0 - x
	}
	// 利用数组
	arr := []int{}
	for {
		n := x / 10
		mod := x - n*10
		arr = append(arr, mod)
		if n == 0 {
			break
		}
		x = n
	}

	res := 0
	for _, v := range arr {
		res = res*10 + v
	}
	if !positive {
		res = 0 - res
	}
	if res >= 1<<31 || res < -(1<<31) {
		return 0
	}
	return res
}
