package main

// 关键字：return xxx
// 执行顺序为
// 返回值 = xxx
// 调用defer
// 空的return命令

// 分解defer
// 第一个例子
func foo() (r int) {
	t := 5
	defer func() {
		t += 5
	}()

	return t
}

/* 拆解
t := 5
r = t // 赋值
func () {t+=5} // defer被插入到赋值与返回值之间执行，r没有被修改过
return t // 空的return命令
*/

// 第二个列子
func foo2() (r int) {
	defer func(r int) {
		r += 10
	}(r)
	return 5
}

/*
拆解：
r = 5 //赋值
func(r int) {
	r+=5
}(r)
return // 空的return命令
*/
