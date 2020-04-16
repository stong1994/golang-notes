package main

/*
文章（https://todebug.com/defer-in-go/）的结论为：
1. return对返回变量赋值，如果是匿名返回值就先声明再赋值；
2. 执行defer函数；
3. return携带返回值返回。
*/
func main() {
	d()
}

func d() {
	var i int
	defer func() {
		i++
	}()
	return
}

/*
go tool compile -S main.go
编译后部分代码如下，可以看到执行了两次RET，第一次执行后调用了defer，然后调用了第二次RET,证明上边的结论是正确的。

0x0057 00087 (main.go:18)       CALL    "".d.func1(SB)
0x005c 00092 (main.go:18)       MOVQ    40(SP), BP
0x0061 00097 (main.go:18)       ADDQ    $48, SP
0x0065 00101 (main.go:18)       RET
0x0066 00102 (main.go:18)       CALL    runtime.deferreturn(SB)
0x006b 00107 (main.go:18)       MOVQ    40(SP), BP
0x0070 00112 (main.go:18)       ADDQ    $48, SP
0x0074 00116 (main.go:18)       RET
*/
