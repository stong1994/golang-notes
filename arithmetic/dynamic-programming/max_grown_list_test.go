package dynamic_programming

import (
	"fmt"
	"testing"
)

/*
我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？
比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。
*/
var arr = []int{2, 9, 3, 6, 5, 1, 7}

const arrLen = 7

func TestMaxSubStr(t *testing.T) {
	hs()
}

// 回溯
func hs() {
	m := make(map[int][]int) // 长度-对应切片
	var f func(i int, curr []int)
	f = func(i int, curr []int) {
		if i >= arrLen {
			m[len(curr)] = curr
			return
		}
		//if i == 0 {
		//	// 不用第一个
		//	f(1, curr)
		//	// 用第一个
		//	f(1, append(curr, arr[0]))
		//	return
		//}
		// 不用这个
		f(i+1, curr)
		// 用这个
		//fmt.Println(i)
		if len(curr) == 0 || arr[i] >= curr[len(curr)-1] {
			f(i+1, append(curr, arr[i]))
		}
	}
	f(0, []int{})
	fmt.Println(m)
	for i := arrLen; i >= 0; i-- {
		if a, ok := m[i]; ok {
			fmt.Println(a)
			fmt.Println(i)
			break
		}
	}

}
