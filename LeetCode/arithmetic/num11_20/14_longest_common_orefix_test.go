package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want string
	}{
		{
			"test1",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"test2",
			[]string{"dog", "racecar", "car"},
			"",
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := longestCommonPrefix(v.arg)
			if got != v.want {
				t.Errorf("want %s got %s", v.want, got)
			}
		})
	}
}

// 先找到最短的字符串，然后遍历，查看其他字符串在相同索引处的字符是否相同
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		prefixLen int
		short     = strs[0]
	)
	for _, v := range strs {
		if v == "" {
			return ""
		}
		if len(v) < len(short) {
			short = v
		}
	}

	for i := 0; i < len(short); i++ {
		for _, str := range strs {
			if short[i] != str[i] {
				return short[0:prefixLen]
			}
		}
		prefixLen++
	}
	return short[0:prefixLen]
}
