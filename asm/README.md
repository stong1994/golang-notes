# 汇编

#### 学习资料
1. [夜读分享](https://github.com/cch123/asmshare/blob/master/layout.md)
2. [GO高级编程](https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-01-basic.html)
3. [官方教程](https://golang.org/doc/asm)

#### 命令
1. `DATA`  
    DATA命令用于初始化包变量，DATA命令的语法如下：
    ```
    DATA symbol+offset(SB)/width, value
    ```
    `symbol`为变量在汇编语言中对应的标识符，`offset`是符号开始地址的偏移量，`width`是要初始化内存的宽度大小，
    `value`是要初始化的值。其中当前包中Go语言定义的符号`symbol`，在汇编代码中对应`·symbol`，其中“`·`”中点符号为一个特殊的unicode符号。  
    
    采用以下命令可以给Id变量初始化为十六进制的0x2537，对应十进制的9527（常量需要以美元符号$开头表示）：
    ```
    DATA ·Id+0(SB)/1,$0x37
    DATA ·Id+1(SB)/1,$0x25
    ```
2. `GLOBL`  
    变量定义好之后需要导出以供其它代码引用。Go汇编语言提供了`GLOBL`命令用于将符号导出。  
    ```
    GLOBL symbol(SB), width
    ```
    其中`symbol`对应汇编中符号的名字，`width`为符号对应内存的大小。用以下命令将汇编中的·Id变量导出：
    ```
    GLOBL ·Id, $8
    ```
    
    分析下边代码
    ```
    GLOBL ·NameData(SB),$8
    DATA  ·NameData(SB)/8,$"gopher"
    
    GLOBL ·Name(SB),$16
    DATA  ·Name+0(SB)/8,$·NameData(SB)
    DATA  ·Name+8(SB)/8,$6
    ```
    第二行初始化NameData,内存宽度为8，内容为字符串"gopher"  
    第一行导出NameData, 内存大小为8  
    第五行初始化Name，内存宽度为8，值NameData  
    第六行初始化偏移量为8的Name，内存宽度为8，值为6  
    第四行导出Name，内存大小为16  
    上述代码表示，Name由两个属性，第一个是字符串的值的地址，第二个是长度。即
    ```
    type reflect.StringHeader struct {
        Data uintptr
        Len  int
    }
    ```
    前8个字节对应底层真实字符串数据的指针，也就是符号go.string."gopher"对应的地址。  
    后8个字节对应底层真实字符串数据的有效长度，这里是6个字节。
    
    上述代码中，当GC扫描到NameData变量时，无法知晓改变了内部是否包含指针，因此会报错：`pkgpath.NameData: missing Go //type information for global symbol: size 8`  
    需加上关键字，`NOPTR` 表示该数据不包含指针
    ```
    #include "textflag.h"
    
    GLOBL ·NameData(SB),NOPTR,$8
    ```
    上述代码中，可以将字符串数据和字符串头结构体定义在一起，这样可以避免引入NameData
    ```
    GLOBL ·Name(SB),$24
    
    DATA ·Name+0(SB)/8,$·Name+16(SB)
    DATA ·Name+8(SB)/8,$6
    DATA ·Name+16(SB)/8,$"gopher"
    ```
    


#### 其他
Go汇编为了简化汇编代码的编写，引入了PC、FP、SP、SB四个伪寄存器

在AMD64环境，`伪PC寄存器`其实是`IP指令计数器寄存器`的别名。  
`伪FP寄存器`对应的是函数的`帧指针`，一般用来访问函数的`参数`和`返回值`。  
`伪SP栈指针`对应的是当前`函数栈帧`的`底部`（不包括参数和返回值部分），一般用于`定位局部变量`。伪SP是一个比较特殊的寄存器，因为还存在一个同名的SP真寄存器。  
`真SP寄存器`对应的是`栈`的`顶部`，一般用于定位调用`其它函数`的`参数`和`返回值`。  
当需要区分伪寄存器和真寄存器的时候只需要记住一点：`伪寄存器`一般需要一个`标识符`和`偏移量`为前缀，如果没有标识符前缀则是真寄存器。比如(SP)、+8(SP)没有标识符前缀为真SP寄存器，而a(SP)、b+8(SP)有标识符为前缀表示伪寄存器。

`SRODATA`标志表示这个数据在只读内存段，`dupok`表示出现多个相同标识符的数据时只保留一个就可以了