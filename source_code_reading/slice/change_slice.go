package main

import "fmt"

func copyThrouthFunc() {
	arr := []int{1,2,3}
	fmt.Println(arr)
	fmt.Printf("%p\n", arr)

	myAppend(arr)
	fmt.Println(arr)
	fmt.Printf("%p\n", arr)

	myAppendPtr(&arr)
	fmt.Println(arr)
	fmt.Printf("外部%p\n", arr)
}

func myAppend(arr []int)  {
	arr = append(arr, 4)
}

func myAppendPtr(arr *[]int)  {
	fmt.Printf("appendPtr前 %p\n", arr)
	*arr = append(*arr, 5)
	fmt.Printf("appendPtr后 %p\n", arr)
}
/*
结果:
[1 2 3]
0xc0000a4020
[1 2 3]
0xc0000a4020
appendPtr前 0xc000090020
appendPtr后 0xc000090020
[1 2 3 5]
外部0xc000096060
 */

/*
 解释:
 第一次append传递的是slice的值,即普通意义上的值拷贝,因此函数内部的改变不会影响函数外部
 第二次append传递的是slice的指针地址,指针地址发生拷贝,所以指针地址改变了,但是函数内部修改的slice的值会影响外部.
	TODO 但是为什么外部的指针地址会改变??
  */

func shareBottomArr()  {
	s1 := []int{1, 2, 3}
	s2 := s1[:]

	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 0
	fmt.Println(s1)
	fmt.Println(s2)
	s1 = append(s1, 4)
	fmt.Println(s1)
	fmt.Println(s2)
}
/*
输出:
[1 2 3]
[1 2 3]
[0 2 3]
[0 2 3]
[0 2 3 4]
[0 2 3]
 */
 /*
 解释:
 s1和s2共用同一个底层数组,所以修改s1,s2也会对应得到修改.
 但是增加元素则不同,因为s1的len会加1,而s2不会,所以s2还是老样子.
  */
