1. for循环中协程引用i,v。
首先，关于闭包：闭包是在词法上下文中引用了自由变量的函数。
go中，被闭包引用的变量会逃逸到堆上。所以i会通过runtime.newobject来创建，且for循环中i只初始化一次，即循环中都是引用的同一个变量。
协程在引用的时候，引用的是那个是不确定的，只是for循环的速度太快，所以一般都是最后一个。

一般编译的时候会将函数内部的变量放到栈或者堆中，但是对于for循环这种表达式，编译器会做优化，将i放到寄存器中，
且只放到一个寄存器中（减少寄存器的占用与指令执行），所以for循环全局的i的地址是一样的，对于v的值来说，依赖于i，
因为i的地址只有一个，那么v的地址也可以理解为只有一个（其实是通过i来找到当前元素相对于切片首地址的偏移量），
