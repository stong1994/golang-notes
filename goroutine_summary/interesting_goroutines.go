package main

import (
	"fmt"
	"time"
)

// [有趣的goroutines&可视化goroutines](https://divan.dev/posts/go_concurrency_visualize/)

// 利用无缓冲的通道的阻塞性质,可以来打一场网球/羽毛球/乒乓球..
// 为了增加趣味性,可以通过获取随机数的奇偶来定义接球失败或者发球失败. todo
// 如果再加上一个player呢? todo
//Go runtime holds waiting FIFO queue for receivers (goroutines ready to receive on the particular channel)
func main() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	fmt.Println("the first ball")
	table <- Ball
	time.Sleep(time.Second)
	<- table
	fmt.Println("the last ball")
}

func player(table chan int) {
	for {
		ball := <- table
		fmt.Printf("received the ball, and is the %dst ball\n", ball)
		ball++
		time.Sleep(time.Millisecond*100)
		fmt.Println("hit it back")
		table <- ball
	}
}
