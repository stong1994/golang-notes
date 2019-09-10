# interface源码

*go version go1.12.5 linux/amd64*

按是否有方法可将接口分为**空接口**和**非空接口**两种

#### 结构体
```go
// 空interface// 
type eface struct {
	_type *_type // 接口类型
	data  unsafe.Pointer // 指向接口具体的值
}

// 非空interface
type iface struct {
	tab  *itab // 接口的类型以及实体类型
	data unsafe.Pointer // 指向接口具体的值
}

// 接口的类型及实体类型等
type itab struct {
	inter *interfacetype // 接口类型内容 // 8 byte
	_type *_type // 实体类型 // 8 bite
	hash  uint32 // 用于类型转换 // 4 byte
	_     [4]byte // 用于填充填充,具体可看 false_sharing 目录下的 README.md
	fun   [1]uintptr // 接口对应的具体方法的地址 variable sized. fun[0]==0 means _type does not implement inter.
}

type interfacetype struct {
	typ     _type // 接口类型
	pkgpath name // 包路径
	mhdr    []imethod // 接口方法的声明数组,不包含具体的方法
}

type imethod struct {
	name nameOff // 方法名的偏移量
	ityp typeOff // 方法类型的偏移量
}

// 数据类型的基础结构体,每种类型的结构体都包含该元素
type _type struct {
	size       uintptr // 类型大小
	ptrdata    uintptr // 存储所有指针的内存前缀的大小
	hash       uint32 // 类型的哈希值
	tflag      tflag // 和反射有关
	// 和对齐相关
	align      uint8 
	fieldalign uint8 
	// 类型的编号
	kind       uint8
	alg        *typeAlg
	// gc相关
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
```
对比空接口和非空接口,发现非空接口比空接口少了几乎一个`itab`类型的元素,原因在于空接口没有方法,因此字段会少一些.

#### 转换为非空接口
go代码
```go
type Animal interface {
	Name()
}

type Cat struct {
	name string
}

func (c Cat) Name() {
	fmt.Printf(c.name)
}

func main() {
	var cat Animal = &Cat{"cat"}
	cat.Name()
}
```
汇编结果
```
0x0025 00037 (impl_not_empty_interface.go:41)   LEAQ    go.string."cat"(SB), AX
0x002c 00044 (impl_not_empty_interface.go:41)   PCDATA  $2, $0
0x002c 00044 (impl_not_empty_interface.go:41)   MOVQ    AX, ""..autotmp_1+32(SP)
0x0031 00049 (impl_not_empty_interface.go:41)   MOVQ    $3, ""..autotmp_1+40(SP)
0x003a 00058 (impl_not_empty_interface.go:41)   PCDATA  $2, $1
0x003a 00058 (impl_not_empty_interface.go:41)   LEAQ    go.itab."".Cat,"".Animal(SB), AX
0x0041 00065 (impl_not_empty_interface.go:41)   PCDATA  $2, $0
0x0041 00065 (impl_not_empty_interface.go:41)   MOVQ    AX, (SP)
0x0045 00069 (impl_not_empty_interface.go:41)   PCDATA  $2, $1
0x0045 00069 (impl_not_empty_interface.go:41)   PCDATA  $0, $0
0x0045 00069 (impl_not_empty_interface.go:41)   LEAQ    ""..autotmp_1+32(SP), AX
0x004a 00074 (impl_not_empty_interface.go:41)   PCDATA  $2, $0
0x004a 00074 (impl_not_empty_interface.go:41)   MOVQ    AX, 8(SP)
0x004f 00079 (impl_not_empty_interface.go:41)   CALL    runtime.convT2I(SB)
```
查看`runtime.convT2I(SB)`函数
```go
// runtime/iface.go:386
func convT2I(tab *itab, elem unsafe.Pointer) (i iface) {
	t := tab._type // 实体类型
	x := mallocgc(t.size, t, true) // 按照实体的大小分配内存
	typedmemmove(t, x, elem) // 赋值elem给新分配的x
	i.tab = tab // 
	i.data = x
	return
}
```
**结论**:值类型转换为非空接口时,将原数据复制一份,然后将地址给`iface.data`.

修改`main`函数,将指针类型转换为非空接口

```go
func main() {
	var cat Animal = &Cat{"cat"}
	cat.Name()
}
```
汇编结果
```
0x001d 00029 (impl_not_empty_interface.go:18)   PCDATA  $2, $1
0x001d 00029 (impl_not_empty_interface.go:18)   PCDATA  $0, $0
0x001d 00029 (impl_not_empty_interface.go:18)   LEAQ    type."".Cat(SB), AX
0x0024 00036 (impl_not_empty_interface.go:18)   PCDATA  $2, $0
0x0024 00036 (impl_not_empty_interface.go:18)   MOVQ    AX, (SP)
0x0028 00040 (impl_not_empty_interface.go:18)   CALL    runtime.newobject(SB)

``` 
查看`runtime.newobject(SB)`
```go
// runtime/malloc.go:1067
func newobject(typ *_type) unsafe.Pointer {
	return mallocgc(typ.size, typ, true)
}
```
**结论**: 指针类型转化为接口时,数据没有被复制,而是直接将地址赋给`iface.data`

#### 转换为空接口
go代码
```go
var es interface{} = 1
var es2 interface{} = "hello"
_, _ = es, es2
```
汇编结果

`go tool compile -S -N impl_empty_interface.go // -S 打印信息 -N 禁止优化`
```
0x000e 00014 (impl_empty_interface.go:4)        LEAQ    type.int(SB), AX
0x0015 00021 (impl_empty_interface.go:4)        PCDATA  $2, $0
0x0015 00021 (impl_empty_interface.go:4)        MOVQ    AX, "".es+16(SP)
0x001a 00026 (impl_empty_interface.go:4)        PCDATA  $2, $1
0x001a 00026 (impl_empty_interface.go:4)        LEAQ    "".statictmp_0(SB), AX
0x0021 00033 (impl_empty_interface.go:4)        PCDATA  $2, $0
0x0021 00033 (impl_empty_interface.go:4)        MOVQ    AX, "".es+24(SP)
0x0026 00038 (impl_empty_interface.go:5)        PCDATA  $2, $1
0x0026 00038 (impl_empty_interface.go:5)        LEAQ    type.string(SB), AX
0x002d 00045 (impl_empty_interface.go:5)        PCDATA  $2, $0
0x002d 00045 (impl_empty_interface.go:5)        MOVQ    AX, "".es2(SP)
0x0031 00049 (impl_empty_interface.go:5)        PCDATA  $2, $1
0x0031 00049 (impl_empty_interface.go:5)        LEAQ    "".statictmp_1(SB), AX
0x0038 00056 (impl_empty_interface.go:5)        PCDATA  $2, $0
0x0038 00056 (impl_empty_interface.go:5)        MOVQ    AX, "".es2+8(SP)
0x003d 00061 (impl_empty_interface.go:7)        MOVQ    32(SP), BP
0x0042 00066 (impl_empty_interface.go:7)        ADDQ    $40, SP
0x0046 00070 (impl_empty_interface.go:7)        RET
```
发现只是对类型进行了赋值
#### 利用编译器检查接口实现
```go
type myWriter struct {

}

/*func (w myWriter) Write(p []byte) (n int, err error) {
    return
}*/

func main() {
    // 检查 *myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = (*myWriter)(nil)

    // 检查 myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = myWriter{}
}
```

#### 知识点
1. 接口最少有两部分：接口类型和数据地址。因此只有接口类型和数据都为空时，接口才会等于nil
2. 如果一个接口的实体类型为值类型,那么会赋值数据并将地址传给`iface.data`,因此改变转为接口后的值不会改变原值,而引用类型和指针类型相反.

### 相关资料
1. [golang中interface底层分析](https://www.jianshu.com/p/ce91ca87fef1?utm_campaign=haruki&utm_content=note&utm_medium=reader_share&utm_source=weixin)
2. [深度解密Go语言之关于 interface 的10个问题--饶大](https://www.cnblogs.com/qcrao-2018/p/10766091.html)

### TODO
1. gdb调试
2. 接口转换
3. 阅读汇编源码
