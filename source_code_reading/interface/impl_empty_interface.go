package main

func main() {
	var es interface{} = 1
	var es2 interface{} = "hello"
	_, _ = es, es2
}

/*
go tool compile -S -N impl_empty_interface.go // -S 打印信息 -N 禁止优化

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
*/
