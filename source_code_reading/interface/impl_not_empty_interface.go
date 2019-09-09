package main

import "fmt"

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

/*
0x001d 00029 (impl_not_empty_interface.go:18)   PCDATA  $2, $1
0x001d 00029 (impl_not_empty_interface.go:18)   PCDATA  $0, $0
0x001d 00029 (impl_not_empty_interface.go:18)   LEAQ    type."".Cat(SB), AX
0x0024 00036 (impl_not_empty_interface.go:18)   PCDATA  $2, $0
0x0024 00036 (impl_not_empty_interface.go:18)   MOVQ    AX, (SP)
0x0028 00040 (impl_not_empty_interface.go:18)   CALL    runtime.newobject(SB)

 */

 // 发现汇编中没有类型转换函数
 // 将cat的值赋给接口就会有,即
//func main() {
//	var cat Animal = Cat{"cat"}
//	cat.Name()
//}

/*
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

 */

 // TODO : 遗留问题:在值类型实现接口时,为什么把指针赋给接口就不会调用类型转换函数.