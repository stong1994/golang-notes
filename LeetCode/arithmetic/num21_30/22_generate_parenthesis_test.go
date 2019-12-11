package num21_30

import (
	"fmt"
	"testing"
)

// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
func TestGenerateParenthesis(t *testing.T) {
	arr := generateParenthesis(5)
	fmt.Println(len(arr))
	fmt.Println(arr)
}


//func generateParenthesis(n int) []string {
//	if n == 0 {
//		return []string{""}
//	}
//	var res []string
//	for c := 0; c < n; c++ {
//		left := generateParenthesis(c)
//		for _, l := range left {
//			right := generateParenthesis(n - 1 - c)
//			for _, r := range right {
//				res = append(res, "("+l+")"+r)
//			}
//		}
//	}
//	return res
//}

func generateParenthesis(n int) []string {
	ants := make([]string, 0, n*2)

	backtrack(&ants, "(", 1, 0, n)
	return ants
}

// 回溯法
func backtrack(ants *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ants = append(*ants, cur)
		return
	}
	if open < max {
		backtrack(ants, cur+"(", open+1, close, max)
	}
	if close < max && close < open {
		backtrack(ants, cur+")", open, close+1, max)
	}
	return
}