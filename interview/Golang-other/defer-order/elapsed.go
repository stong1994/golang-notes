package main

import (
	"fmt"
	"time"
)

/**
利用延迟函数defer来封装访问方法的时间函数
使用方法: defer Elapsed("funcName")()
funName: 方法名称
*/
func Elapsed(funName string) func() {
	start := time.Now()
	fmt.Println("entered ", funName)
	return func() {
		duration := time.Since(start)
		fmt.Println(funName, " elapsed ", duration)
	}
}
