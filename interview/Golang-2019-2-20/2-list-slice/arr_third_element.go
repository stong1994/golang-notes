package main

import "fmt"

// 切片中可以使用索引进行赋值，并且其余地方默认为零值
// 从切片或数组中 切割出元素时，也可以指定第三个元素，来指定新切片容量
func main() {
	s1 := []int{0, 1, 2, 3, 8: 100} // [0 1 2 3 0 0 0 0 100]
	fmt.Println(s1)
	s2 := []int{1: 1, 3: 3, 8: 8} // [0 1 0 3 0 0 0 0 8]
	fmt.Println(s2)
	s3 := s2[2:4:5] // [0 3], 索引到5，所以cap为3（2-5）
	fmt.Println(s3)
	s4 := s2[2:4:10] // panic: runtime error: slice bounds out of range
	fmt.Println(s4)
}
