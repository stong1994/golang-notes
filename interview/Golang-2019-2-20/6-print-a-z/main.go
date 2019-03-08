package main

import (
	"fmt"
)

func main() {
	min := int('a')                // a的ascii
	max := int('z')                // z的ascii
	endFalg := make(chan struct{}) // 结束标志
	var flag bool
	oddFlag := make(chan int)  // 奇数信号
	evenFlag := make(chan int) // 偶数信号

	end := func(i int) bool {
		if i > max {
			endFalg <- struct{}{}
			flag = true
			return true
		}
		return false
	}

	// 奇数
	odd := func(i int) {
		if !end(i) {
			fmt.Print(string(i))
			evenFlag <- i + 1
		}
	}

	// 偶数
	even := func(i int) {
		if !end(i) {
			fmt.Print(string(i))
			oddFlag <- i + 1
		}
	}

	//分配任务
	assignJob := func() {
		for {
			select {
			case o := <-oddFlag:
				go odd(o)
			case e := <-evenFlag:
				go even(e)
			}
		}
	}

	go assignJob()

	// 入口
	odd(min)

	// 等待退出
	for {
		select {
		case <-endFalg:
			return
		}
	}
}
