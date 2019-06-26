package num_1_10

import (
	"fmt"
	"strings"
	"testing"
)

func TestNum5(t *testing.T) {
	str := "abcda"
	result := method2(str)
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	maxDuration := 0
	maxStr := ""
	tmpMap := make(map[int32][]int)
	for i, v := range s {
		if val, ok := tmpMap[v]; ok {
			for _, k := range val {
				duration := i - k
				if duration > maxDuration {
					tmpStr := string([]byte(s)[k : i+1])
					if isTrue(tmpStr) {
						maxDuration = duration
						maxStr = tmpStr
					}
				}
			}
			val = append(val, i)
			tmpMap[v] = val
		} else {
			tmpMap[v] = []int{i}
		}
	}

	if maxStr == "" {
		maxStr = string(s[0])
	}
	return maxStr
}

func isTrue(s string) bool {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func method2(s string) string {
	if len(s) < 2 {
		return s
	}
	s2 := ""
	for i := 0; i < len(s)-1; i++ {
		s2 = s2 + string(s[i]) + "-"
	}
	s2 += string(s[len(s)-1])
	result := make(map[int]string, 0)
	for start := 1; start < len(s2)-1; start++ {
		res := symmetry(s2, start)
		l := len(res)
		if l > 1 {
			trimL := strings.Replace(res, "-", "", -1)
			result[len(trimL)] = trimL
		}

	}
	maxIdx := 0
	for i := range result {
		if i > maxIdx {
			maxIdx = i
		}
	}
	if maxIdx <= 1 {
		return string(s[0])
	}
	return result[maxIdx]
}

func symmetry(s string, index int) string {
	result := string(s[index])
	tag := false // true 为index所处索引位于字符串的左侧
	if index < len(s)/2 {
		tag = true
	}
	if tag {
		for i := index - 1; i >= 0; i-- {
			j := 2*index - i
			if res, tag := judge(s, i, j); tag {
				result = string(res) + result + string(res)
			} else {
				return result
			}
		}
	} else {
		for i := index + 1; i <= len(s)-1; i++ {
			j := 2*index - i
			if res, tag := judge(s, i, j); tag {
				result = string(res) + result + string(res)
			} else {
				return result
			}
		}
	}
	return result
}

func judge(s string, left, right int) (byte, bool) {
	if s[left] == s[right] {
		return s[left], true
	}
	return 0, false
}
