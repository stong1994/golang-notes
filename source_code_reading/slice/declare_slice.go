package main

var declare []int
var declare_make = make([]int, 5, 10)
var slice_init = []int{1}

/*
0x003a 00058 (declare_slice.go:4)       CALL    runtime.makeslice(SB)
0x005a 00090 (declare_slice.go:4)       CMPL    runtime.writeBarrier(SB), $0
0x007b 00123 (declare_slice.go:4)       CALL    runtime.gcWriteBarrier(SB)
"".declare SBSS size=24
"".declare_make SBSS size=24
"".slice_init SDATA size=24
 */

var declare_make_64 = make([]int, int64(5), int64(10))
// 0x003a 00058 (declare_slice.go:16)      CALL    runtime.makeslice(SB)