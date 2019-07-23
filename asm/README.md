# 汇编

#### 学习资料
1. [夜读分享](https://github.com/cch123/asmshare/blob/master/layout.md)
2. [GO高级编程](https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-01-basic.html)
3. [官方教程](https://golang.org/doc/asm)

#### 知识点
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
    从上边对字符串和整数的表示中，我们可以看出，**汇编并没有这些具体的类型**，它只是通过不同的结构来表示这些类型。  
    `GLOBL file_private<>(SB),$1`其中`<>`表示私有类型
    
3. `textflag.h`
    Go汇编语言还在`textflag.h`文件定义了一些标志。其中用于变量的标志有`DUPOK`、`RODATA`和`NOPTR`几个。
    `DUPOK`表示该变量对应的标识符可能有多个，在链接时只选择其中一个即可（一般用于合并相同的常量字符串，减少重复数据占用的空间）  
    `RODATA`标志表示将变量定义在**只读**内存段，因此后续任何对此变量的修改操作将导致异常（`recover`也**无法捕获**）  
    `NOPTR`则表示此变量的内部`不含指针数据`，让垃圾回收器忽略对该变量的扫描。如果变量已经在Go代码中声明过的话，Go编译器会自动分析出该变量是否包含指针，这种时候可以不用手写NOPTR标志。  
    代码：  
    ```
    #include "textflag.h"
    
    GLOBL ·const_id(SB),NOPTR|RODATA,$8
    DATA  ·const_id+0(SB)/8,$9527
    ```
    该代码表示了一个int64类型的值为9527的整数，且该整数类型不包含指针数据且为只读。变量名为const_id
    
4. `TEXT`
    函数标识符通过TEXT汇编指令定义，表示该行开始的指令定义在TEXT内存段。  
    函数的名字后面是(SB)，表示是函数名符号相对于SB伪寄存器的偏移量，二者组合在一起最终是绝对地址。  
    标志部分用于指示函数的一些特殊行为，标志在textlags.h文件中定义，常见的`NOSPLIT`主要用于指示叶子函数不进行栈分裂。  
    framesize部分表示函数的局部变量需要多少栈空间，其中包含调用其它函数时准备调用参数的隐式栈空间。  
    最后是可以省略的参数大小，之所以可以省略是因为编译器可以从Go语言的函数声明中推导出函数参数的大小。   
    如用于交换数据的`Swap`函数:
    ```go
    package main
    
    //go:nosplit
    func Swap(a, b int) (int, int)
    ```
    zaIn汇编中有两种表示：
    ```
    // func Swap(a, b int) (int, int)
    TEXT ·Swap(SB), NOSPLIT, $0-32
    
    // func Swap(a, b int) (int, int)
    TEXT ·Swap(SB), NOSPLIT, $0
    ``` 
    其中`32`表示**参数**和**返回值**的`4`个`int`类型.  
    目前可能遇到的函数标志有`NOSPLIT`、`WRAPPER`和`NEEDCTXT`几个  
    `NOSPLIT`不会生成或包含栈分裂代码，这一般用于没有**任何其它函数调用**的叶子函数，这样可以适当提高性能。  
    `WRAPPER`标志则表示这个是一个包装函数，在`panic`或`runtime.caller`等某些处理函数帧的地方**不会增加函数帧计数**。  
    `NEEDCTXT`表示**需要一个上下文参数**，一般用于**闭包函数**。  
    需要注意的是函数也没有类型，上面定义的Swap函数签名可以下面任意一种格式:
    ```go
    func Swap(a, b, c int) int
    func Swap(a, b, c, d int)
    func Swap() (a, b, c, d int)
    func Swap() (a []int, d int)
    ```      
    下面的代码演示了如何在汇编函数中使用参数和返回值：   
    GO中的函数：`func Swap(a, b int) (ret0, ret1 int)`  
    汇编：
    ```
    TEXT ·Swap(SB), $0
        MOVQ a+0(FP), AX     // AX = a
        MOVQ b+8(FP), BX     // BX = b
        MOVQ BX, ret0+16(FP) // ret0 = BX
        MOVQ AX, ret1+24(FP) // ret1 = AX
        RET
    ```
    **复杂的参数和返回值的内存布局**  
    ```go
    func Foo(a bool, b int16) (c []byte)
    ```
    ```
    TEXT ·Foo(SB), $0
        MOVEQ a+0(FP),       AX // a
        MOVEQ b+2(FP),       BX // b
        MOVEQ c_dat+8*1(FP), CX // c.Data
        MOVEQ c_len+8*2(FP), DX // c.Len
        MOVEQ c_cap+8*3(FP), DI // c.Cap
        RET
    ```
    b和a之间空出一个字节，b和c之间空出4个字节。空出的原因是保证每个参数变量地址都要对其到相应的倍数。  
    为了便于访问**局部变量**，Go汇编语言引入了`伪SP`寄存器，对应**当前栈帧的底部**。    
    GO中的函数：
    ```go
    func Foo() {
        var c []byte
        var b int16
        var a bool
    }
    ```
    对应的汇编代码：
    ```
    TEXT ·Foo(SB), $32-0
        MOVQ a-32(SP),      AX // a
        MOVQ b-30(SP),      BX // b
        MOVQ c_data-24(SP), CX // c.Data
        MOVQ c_len-16(SP),  DX // c.Len
        MOVQ c_cap-8(SP),   DI // c.Cap
        RET
    ```
    可以计算出该函数的栈帧大小为32个字节（注意内存对齐所需要的空位）。  
    出现最后定义的`a`离伪`SP寄存器`**最近**的原因是：
    从Go语言函数角度理解，**先定义**的`c`变量地址要比**后定义**的变量的地址**更小**；另一个是伪SP寄存器对应栈帧的底部，而`X86`中`栈`是**从高向地生长**的，所以最先定义有着更小地址的c变量离栈的底部伪SP更远。  
    伪SP寄存器对应高地址，因此对应的局部变量的偏移量都是负数。  
    
    **汇编函数的参数是从哪里来的？**  
    被调用函数的参数是由调用方准备的：调用方在栈上设置好空间和数据后调用函数，被调用方在返回前将返回值放在对应的位置，函数通过`RET`指令返回调用方函数之后，调用方再从返回值对应的栈内存位置取出结果。
    Go语言函数的调用参数和返回值均是通过栈传输的，这样做的优点是函数调用栈比较清晰，缺点是函数调用有一定的性能损耗（Go编译器是通过**函数内联**来缓解这个问题的影响）。  
    
    GO代码：
    ```go
    func main() {
        printsum(1, 2)
    }
    
    func printsum(a, b int) {
        var ret = sum(a, b)
        println(ret)
    }
    
    func sum(a, b int) int {
        return a+b
    }

    ```
    偷张图片。。。
    ![](https://chai2010.cn/advanced-go-programming-book/images/ch3-12-func-call-frame-01.ditaa.png)

5. 控制流
    GO代码：
    ```go
    func main() {
        var a = 10
        println(a)
    
        var b = (a+a)*a
        println(b)
    }
    ```
    修改为伪汇编：
    ```go
    func main() {
        var a, b int
    
        a = 10
        runtime.printint(a)
        runtime.printnl()
    
        b = a
        b += b
        b *= a
        runtime.printint(b)
        runtime.printnl()
    }
    ```
    汇编代码：
    ```
    TEXT ·main(SB), $24-0
        MOVQ $0, a-8*2(SP) // a = 0
        MOVQ $0, b-8*1(SP) // b = 0
    
        // 将新的值写入a对应内存
        MOVQ $10, AX       // AX = 10
        MOVQ AX, a-8*2(SP) // a = AX
    
        // 以a为参数调用函数
        MOVQ AX, 0(SP)
        CALL runtime·printint(SB)
        CALL runtime·printnl(SB)
    
        // 函数调用后, AX/BX 寄存器可能被污染, 需要重新加载
        MOVQ a-8*2(SP), AX // AX = a
        MOVQ b-8*1(SP), BX // BX = b
    
        // 计算b值, 并写入内存
        MOVQ AX, BX        // BX = AX  // b = a
        ADDQ BX, BX        // BX += BX // b += a
        IMULQ AX, BX       // BX *= AX // b *= a
        MOVQ BX, b-8*1(SP) // b = BX
    
        // 以b为参数调用函数
        MOVQ BX, 0(SP)
        CALL runtime·printint(SB)
        CALL runtime·printnl(SB)
    
        RET
    ```
    for循环代码：
    ```go
    func LoopAdd(cnt, v0, step int) int {
        result := v0
        for i := 0; i < cnt; i++ {
            result += step
        }
        return result
    }
    ```
    转换成为汇编代码
    ```go
    func LoopAdd(cnt, v0, step int) int {
        var i = 0
        var result = 0
    
    LOOP_BEGIN:
        result = v0
    
    LOOP_IF:
        if i < cnt { goto LOOP_BODY }
        goto LOOP_END
    
    LOOP_BODY
        i = i+1
        result = result + step
        goto LOOP_IF
    
    LOOP_END:
    
        return result
    }
    ```
    汇编代码：
    ```
    #include "textflag.h"
    
    // func LoopAdd(cnt, v0, step int) int
    TEXT ·LoopAdd(SB), NOSPLIT,  $0-32
        MOVQ cnt+0(FP), AX   // cnt
        MOVQ v0+8(FP), BX    // v0/result
        MOVQ step+16(FP), CX // step
    
    LOOP_BEGIN:
        MOVQ $0, DX          // i
    
    LOOP_IF:
        CMPQ DX, AX          // compare i, cnt
        JL   LOOP_BODY       // if i < cnt: goto LOOP_BODY
        JMP LOOP_END
    
    LOOP_BODY:
        ADDQ $1, DX          // i++
        ADDQ CX, BX          // result += step
        JMP LOOP_IF
    
    LOOP_END:
    
        MOVQ BX, ret+24(FP)  // return result
        RET
    ```
6. 函数分析  
    先构造一个禁止栈分裂的printnl函数。printnl函数内部都通过调用runtime.printnl函数输出换行：  
    ```
    TEXT ·printnl_nosplit(SB), NOSPLIT, $8
        CALL runtime·printnl(SB)
        RET
    ```
    通过命令`go tool asm -S main_amd64.s`查看编译后的源码:
    ```
    "".printnl_nosplit STEXT nosplit size=29 args=0xffffffff80000000 locals=0x10
    0x0000 00000 (main_amd64.s:5) TEXT "".printnl_nosplit(SB), NOSPLIT    $16
    0x0000 00000 (main_amd64.s:5) SUBQ $16, SP
    
    0x0004 00004 (main_amd64.s:5) MOVQ BP, 8(SP)
    0x0009 00009 (main_amd64.s:5) LEAQ 8(SP), BP
    
    0x000e 00014 (main_amd64.s:6) CALL runtime.printnl(SB)
    
    0x0013 00019 (main_amd64.s:7) MOVQ 8(SP), BP
    0x0018 00024 (main_amd64.s:7) ADDQ $16, SP
    0x001c 00028 (main_amd64.s:7) RET
    ```
    加上缩进：
    ```
    TEXT "".printnl(SB), NOSPLIT, $16
        SUBQ $16, SP
            MOVQ BP, 8(SP)
            LEAQ 8(SP), BP
                CALL runtime.printnl(SB)
            MOVQ 8(SP), BP
        ADDQ $16, SP
    RET
    ```
    - 第一层`TEXT`表示指令开始，`RET`表示结束
    - 第二层`SUBQ`给`SP`分配了16个字节的空间，`ADDQ`收回这16个字节（发现多分配了8字节）
    - 第三层`MOVQ`将`BP`寄存器保存到了多分配的8字节中，`LEAQ`将`8(SP)`的地址再保存到BP中，最后`MOVQ`恢复之前备份的前`BP`寄存器中的值  
    
    去掉NOSPLIT
    ```
    TEXT "".printnl_nosplit(SB), $16
    L_BEGIN:
        MOVQ (TLS), CX
        CMPQ SP, 16(CX)
        JLS  L_MORE_STK
    
            SUBQ $16, SP
                MOVQ BP, 8(SP)
                LEAQ 8(SP), BP
                    CALL runtime.printnl(SB)
                MOVQ 8(SP), BP
            ADDQ $16, SP
    
    L_MORE_STK:
        CALL runtime.morestack_noctxt(SB)
        JMP  L_BEGIN
    RET
    ```
    发现增加了一些内容(**扩容**)：  
    `MOVQ (TLS), CX`用于加载`g`结构体指针  
    `CMPQ SP, 16(CX)` `SP`栈指针和`g`结构体中`stackguard0`成员比较，如果比较的结果小于0则跳转到结尾的`L_MORE_STK`部分  
    当获取到更多栈空间之后，通过`JMP L_BEGIN`指令跳转到函数的开始位置重新进行栈空间的检测。  
    > 在g结构体中的`stackguard0`成员是出现爆栈前的警戒线。`stackguard0`的偏移量是16个字节，因此上述代码中的`CMPQ SP, 16(AX)`表示将当前的真实`SP`和爆栈警戒线比较，如果超出警戒线则表示需要进行`栈扩容`，也就是跳转到`L_MORE_STK`。
    在`L_MORE_STK`标号处，先调用`runtime·morestack_noctxt`进行栈扩容，然后又跳回到函数的开始位置，此时此刻函数的栈已经调整了。然后再进行一次栈大小的检测，如果依然不足则继续扩容，直到栈足够大为止。
    
7. 方法函数
    GO代码：
    ```go
    package main
    
    type MyInt int
    
    func (v MyInt) Twice() int {
        return int(v)*2
    }
    
    func MyInt_Twice(v MyInt) int {
        return int(v)*2
    }
    ```
    汇编
    ```
    // func (v MyInt) Twice() int
    TEXT ·MyInt·Twice(SB), NOSPLIT, $0-16
        MOVQ a+0(FP), AX   // v
        ADDQ AX, AX        // AX *= 2
        MOVQ AX, ret+8(FP) // return v
        RET
    ```
8. 闭包
    GO代码
    ```go
    package main
    
    func NewTwiceFunClosure(x int) func() int {
        return func() int {
            x *= 2
            return x
        }
    }
    
    func main() {
        fnTwice := NewTwiceFunClosure(1)
    
        println(fnTwice()) // 1*2 => 2
        println(fnTwice()) // 2*2 => 4
        println(fnTwice()) // 4*2 => 8
    }
    ```
    手动构造闭包函数,`F`表示闭包函数的函数指令的地址，`X`表示闭包捕获的外部变量
    ```go
    type FunTwiceClosure struct {
        F uintptr
        X int
    }
    
    func NewTwiceFunClosure(x int) func() int {
        var p = &FunTwiceClosure{
            F: asmFunTwiceClosureAddr(),
            X: x,
        }
        return ptrToFunc(unsafe.Pointer(p))
    }
    ```
    汇编语言实现了以下三个辅助函数
    ```go
    func ptrToFunc(p unsafe.Pointer) func() int
    func asmFunTwiceClosureAddr() uintptr
    func asmFunTwiceClosureBody() int
    ```
    - `asmFunTwiceClosureAddr`用于获取闭包函数的函数指令的地址
    - `ptrToFunc`将结构体指针转为闭包函数对象
    - `asmFunTwiceClosureBody`是闭包函数对应的全局函数的实现
    用汇编实现以上三个辅助函数
    ```
    #include "textflag.h"
    
    TEXT ·ptrToFunc(SB), NOSPLIT, $0-16
        MOVQ ptr+0(FP), AX // AX = ptr
        MOVQ AX, ret+8(FP) // return AX
        RET
    
    TEXT ·asmFunTwiceClosureAddr(SB), NOSPLIT, $0-8
        LEAQ ·asmFunTwiceClosureBody(SB), AX // AX = ·asmFunTwiceClosureBody(SB)
        MOVQ AX, ret+0(FP)                   // return AX
        RET
    
    TEXT ·asmFunTwiceClosureBody(SB), NOSPLIT|NEEDCTXT, $0-8
        MOVQ 8(DX), AX
        ADDQ AX   , AX        // AX *= 2
        MOVQ AX   , 8(DX)     // ctx.X = AX
        MOVQ AX   , ret+0(FP) // return AX
        RET
    ```
    `NEEDCTXT`标志定义的汇编函数表示需要一个上下文环境，在AMD64环境下是通过`DX`寄存器来传递这个上下文环境指针，也就是对应`FunTwiceClosure`结构体的指针。  
    > 函数首先从`FunTwiceClosure`结构体对象取出之前捕获的`X`，将`X`乘以2之后写回内存，最后返回修改之后的`X`的值。  
    
    整个闭包的流程：
    > 1. 构建闭包对象，成员分别为闭包函数的指令地址和闭包函数的外部变量
    > 2. 调用闭包函数时，先拿到闭包对象，根据闭包函数的指令地址调用`CALL`来运行闭包函数，然后更新闭包对象中的外部变量。  
    
    闭包函数就是获取外层局部作用域的局部对象，于是闭包函数本身就有了状态，从这个角度讲，全局函数也是闭包函数，只是没有调用外层变量。
    
#### 其他
Go汇编为了简化汇编代码的编写，引入了PC、FP、SP、SB四个伪寄存器

在AMD64环境，`伪PC寄存器`其实是`IP指令计数器寄存器`的别名。  
`SB`: 全局静态基指针，一般用来声明函数或全局变量。
`伪FP寄存器`对应的是函数的`帧指针`，一般用来访问函数的`参数`和`返回值`。  
`伪SP栈指针`对应的是当前`函数栈帧`的`底部`（不包括参数和返回值部分），一般用于`定位局部变量`。伪SP是一个比较特殊的寄存器，因为还存在一个同名的SP真寄存器。  
`真SP寄存器`对应的是`栈`的`顶部`，一般用于定位调用`其它函数`的`参数`和`返回值`。  
当需要区分伪寄存器和真寄存器的时候只需要记住一点：`伪寄存器`一般需要一个`标识符`和`偏移量`为前缀，如果没有标识符前缀则是真寄存器。比如(SP)、+8(SP)没有标识符前缀为真SP寄存器，而a(SP)、b+8(SP)有标识符为前缀表示伪寄存器。

`SRODATA`标志表示这个数据在只读内存段，`dupok`表示出现多个相同标识符的数据时只保留一个就可以了


在add目录下做了一个简单的加法汇编，有几点感到很疑惑：
1. 可以用测试函数来调用Add方法，用main不可以
2. 汇编代码要在最后留一个空行。