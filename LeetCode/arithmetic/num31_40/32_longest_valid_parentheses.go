package num31_40

import "testing"

/*
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

输入："((()())()("
输出：8
解释: (()())()
 */

func TestLongestValidParentheses(t *testing.T) {
	test := []struct {
		Name string
		Str string
		Want int
	}{
		{
			Name: "test1",
			Str: "(()",
			Want: 2,
		},
		{
			Name: "test2",
			Str: ")()())",
			Want: 4,
		},
		{
			Name: "test3",
			Str: "((()())()(",
			Want: 8,
		},
	}
}

func longestValidParentheses(s string) int {
	// 先找到最内层的括号组()的位置
	// 依次向外遍历，得到符合条件的括号对数
	var idx []int
	for i := 0; i < len(s)-1; i+=2 {
		if s[i] == '(' && s[i+1] == ')' {
			idx = append(idx, i)
		}
	}
	if len(idx) == 0 {
		return -1
	}
	for _, v := range idx {

	}
}

func matchPair(s string, i int) int {
	num := 1
	j := i+1
	for i > 0  {
		if len(s)<= j+1 {
			return num
		}
		if s[i-1] == '(' && s[j+1] == ')'
	}
}