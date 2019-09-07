package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"test1",
			"III",
			3,
		},
		{
			"test2",
			"IV",
			4,
		},
		{
			"test3",
			"IX",
			9,
		},
		{
			"test4",
			"LVIII",
			58,
		},
		{
			"test5",
			"MCMXCIV",
			1994,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := romanToInt(v.arg)
			if got != v.want {
				t.Errorf("want %d but got %d", v.want, got)
			}
		})
	}
}

func romanToInt(s string) int {
	res := 0
	romanNum := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	romanRole := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	for i := 0; i < len(s); i++ {
		now := romanNum[s[i]]
		if i+1 < len(s) && romanNum[s[i+1]] > now {
			res += romanRole[s[i:i+2]]
			i++
			continue
		}
		res += now
	}
	return res
}
