package main

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct{
		name string
		arg int
		want string
	}{
		{
			"test1",
			3,
			"III",
		},
		{
			"test2",
			4,
			"IV",
		},
		{
			"test3",
			9,
			"IX",
		},
		{
			"test4",
			58,
			"LVIII",
		},
		{
			"test5",
			1994,
			"MCMXCIV",
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := intToRoman(v.arg)
			if got != v.want {
				t.Errorf("want %s but got %s", v.want, got)
			}
		})
	}
}

func intToRoman(num int) string {
	var res string
	romanRole := map[int]string{1:"I", 5: "V", 10: "X", 50:"L", 100:"C", 500:"D", 1000:"M", 4:"IV", 9:"IX", 40:"XL", 90:"XC", 400: "CD", 900:"CM"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for num > 0 {
		for _, n := range nums {
			if num >= n {
				res += romanRole[n]
				num -= n
				goto next
			}
		}
		next:
	}
	return res
}