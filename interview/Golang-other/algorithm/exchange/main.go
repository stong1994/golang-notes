package main

import "fmt"

// 在不用临时变量的情况下交换两个变量
// 解法在数轴上更好理解

// 方法一
//func main()  {
//	a, b := 10, 11
//	b = b - a
//	a = b + a
//	b = a - b
//	fmt.Println(a, b)
//}

// 方法二
//func main()  {
//	a, b := 10, 11
//	a = a + b
//	b = a - b
//	a = a - b
//	fmt.Println(a, b)
//}

// 方法三 只适用于整数类型
func main() {
	a, b := 10, 11 // 1010 1011
	a = a ^ b      // 0001
	b = a ^ b      // 1010
	a = a ^ b      // 1011
	fmt.Println(a, b)
}
