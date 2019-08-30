package main

func appendSlice()  {
	arr:=  make([]int, 0)
	arr = append(arr, 1)
}

/*
 0x0042 00066 (append_slice.go:5)        CALL    runtime.growslice(SB)
 */