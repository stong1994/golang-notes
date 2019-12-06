package main

import (
	"errors"
	"fmt"
)

// go 1.13中error新增了一些特性
// 1.13以前errors只有一个New()方法，现在新暴露了如下三个

// 0. Unwrap(err error) error  // 获取err的子error，及err所包装的error
// 1. Is(err, target error) bool // err及其error后代们是否包装了target
// 2. As(err error, target interface{}) bool // 将target转换为fmt.wrapErrortrue类型，并将err的内容赋值给target。 target必须是指针类型？

// 对error的包装
// fmt.Errorf() 新增了一个动词 %w, 用来接收error类型，构成对error的包装，值得注意的是，用%v同样可以接收error，但不构成对error的包装，只是产生一个新的error

// 发现了一个骚操作 假设a为interface类型，那么a.(interface{ As(interface{}) bool })能够查看a是否有As(interface{}) bool{}方法)

/*
if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
	return true
}
 */

var Err0 = errors.New("err0")
var Err1 = errors.New("err1")
var Err2 error = nil

func main()  {
	err := getErr()
	if errors.Is(err, Err0) {
		fmt.Println(1)
	}
	if errors.Is(err, Err1) {
		fmt.Println(2)
	}
	// As
	if errors.As(err, &Err0) {
		fmt.Println(3)
		fmt.Println(Err0)
	}
	if errors.As(err, &Err1) {
		fmt.Println(4)
		fmt.Println(Err1)
	}
	if errors.As(err, &Err2) {
		fmt.Println(5)
		fmt.Println(Err2)
	}
	v := interface{}(nil)
	if errors.As(err, &v) {
		fmt.Println(6)
		fmt.Printf("%T", v)
		fmt.Println(v)
	}

	var a interface{}
	a = aBool(true)
	if x, ok := a.(interface{ Get(interface{}) bool }); ok {
		fmt.Println(7)
		fmt.Println(x)
	}
}

func getErr() error {
	return fmt.Errorf("err is: %w", Err0)
}

type aBool bool

func (a aBool) Get(data interface{}) bool {
	return bool(a)
}

func (a aBool) Set() {
	a = false
}