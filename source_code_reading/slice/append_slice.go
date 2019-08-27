package main

func appendSlice()  {
	var arr []int
	arr = append(arr, 1)
}

/*
 0x0042 00066 (append_slice.go:5)        CALL    runtime.growslice(SB)
 */