package num_1_10

import (
	"fmt"
	"testing"
)

func TestNum3(t *testing.T) {
	str := "abcabcbb"
	result := lengthOfLongestSubstring(str)
	target := 3
	if result == target {
		fmt.Println("true")
	}
}

func lengthOfLongestSubstring(s string) int {
	var (
		arr []map[int32]int
		tmp = make(map[int32]int)
	)

	for i, v := range s {
		if i == 0 {
			tmp[v] = i
			continue
		}
		if k, ok := tmp[v]; ok {
			arr = append(arr, tmp)
			tmp = make(map[int32]int)
			for j := k; j <= i; j++ {
				n := int32(s[j])
				tmp[n] = j
			}
		} else {
			tmp[v] = i
		}
	}
	//fmt.Println(arr)
	max := len(tmp)
	for _, v := range arr {
		l := len(v)
		if l > max {
			max = l
		}
	}
	return max
}
