package Golang_2019_12_17

import "fmt"

// 注意两点
// 1. defer执行顺序为先进后出
// 2. 其参数如果为函数则直接执行,如果为值,则通过值复制的方式复制进去,即calc("1", a, b, calc(...))中的a,b为上边的1,2
func deferOrder()  {
	a := 1
	b := 2
	defer calc("1", a, b, calc("10", a, b, a + b)) //
	a = 0
	defer calc("2", a, b, calc("20", a, b, a + b))
	b = 1
}

func calc(index string, a, b, ret int) int {
	ret = ret*2
	fmt.Println(index, a, b, ret)
	return ret
}
/*
1. 执行 calc("10", a, b, a+b), a = 1, b = 2 输出: 10, 1, 2, 6
2. 执行 calc("20", a, b, 0), a = 0, b = 2 输出: 20, 0, 2, 4
3. 执行 calc("2", a, b, 2), a = 0, b = 2 输出: 2, 0, 2, 8
4. 执行 calc("1", a, b, 3), a = 1, b = 2 输出: 1, 1, 2, 12

输出:
10 1 2 6
20 0 2 4
2 0 2 8
1 1 2 12
*/