package main

import (
	"fmt"
	"time"
)

type number int

func (n number) print() {
	fmt.Println("print", n)
}

func (n *number) pprint() {
	fmt.Println("pprint", *n)
}

func main() {
	var n number
	defer n.print()
	defer n.pprint()
	defer func(name string) {
		fmt.Print(name)
		n.print()
	}("匿名函数1 ")
	defer func(name string) {
		fmt.Print(name)
		n.pprint()
	}("匿名函数2 ")
	defer func(name string) func() {
		fmt.Println(name)
		return func() {
			fmt.Print("匿名函数4 ")
			n.print()
		}
	}("匿名函数3 ")()

	n = 3
	time.Sleep(time.Second)
	fmt.Println("休眠1s")
}

/*
执行结果：
匿名函数3
休眠1s
匿名函数4 print 3
匿名函数2 pprint 3
匿名函数1 print 3
pprint 3
print 0
*/

/*
分析：
defer特性1：函数执行完后执行，因此会先打印"休眠1s", 但是注意最后一个defer函数，是一个“立即执行”函数，该函数内主体立即执行，而返回的函数成为了defer的执行函数。
因此会先打印“匿名函数3 ”，再打印"休眠1s"。（当然，我是从结论分析原因，有可能是错的，欢迎大佬指正）
defer特性2：后进先出。于是再打印"匿名函数4 print 3"，“匿名函数2 pprint 3”，“匿名函数1 print 3”， “pprint 3”，“print 0”
至于最后一个输出为0，是因为n初始化为0，第一个defer直接对n取值。而第二个defer是取n的指针对应的值，所以会受后边的影响。

利用最后一个defer的特殊性，我们可以封装一个函数来计算某个函数的执行时间——在elapsed.go文件中
*/
