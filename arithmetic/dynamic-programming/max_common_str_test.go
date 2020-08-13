package dynamic_programming

import (
	"fmt"
	"testing"
)

// 求最长的公共字符串
const str1 = "abcdefghijklmn"
const str2 = "bcdeihjklmn"

func TestMaxCommonStr(t *testing.T) {
	fun1(str1, str2)
}

func fun1(s1, s2 string) string {
	m := make(map[int]string, 0)
	maxL := 0
	var f func(s1, s2, common string, lastMatch bool)
	f = func(leftS1, leftS2, common string, lastMatch bool) {
		if leftS1 == "" || leftS2 == "" {
			if len(common) > maxL {
				maxL = len(common)
				m[len(common)] = common
			}
			return
		}

		// 不考虑leftS1的第一个字符
		f(leftS1[1:], leftS2, "", false)
		// 不考虑leftS2的第一个字符
		f(leftS1, leftS2[1:], "", false)
		// 如果leftS1和leftS2的第一个字符相等
		if leftS1[0] == leftS2[0] {
			if lastMatch {
				f(leftS1[1:], leftS2[1:], common+string(leftS1[0]), true)
			} else {
				f(leftS1[1:], leftS2[1:], string(leftS1[0]), true)
			}

		}
	}
	f(s1, s2, "", false)
	fmt.Println(m)
	fmt.Println(m[maxL])
	return m[maxL]
}

// todo 动态规划
