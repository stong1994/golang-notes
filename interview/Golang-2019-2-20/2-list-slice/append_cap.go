package main

import "fmt"

// 错误的结论：当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。
// 在上边的结论后，还要进行 内存对齐，进行内存对齐后，新slice的容量是要大于之前分配的容量
// 正确的结论： 扩容策略并不是简单的扩为原切片容量的 2 倍或 1.25 倍，还有内存对齐的操作。扩容后的容量 >= 原容量的 2 倍或 1.25 倍。
func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

/*
[0 ->   -1] cap = 0     |  after append 0     cap = 1
[0 ->    0] cap = 1     |  after append 1     cap = 2
[0 ->    1] cap = 2     |  after append 2     cap = 4
[0 ->    3] cap = 4     |  after append 4     cap = 8
[0 ->    7] cap = 8     |  after append 8     cap = 16
[0 ->   15] cap = 16    |  after append 16    cap = 32
[0 ->   31] cap = 32    |  after append 32    cap = 64
[0 ->   63] cap = 64    |  after append 64    cap = 128
[0 ->  127] cap = 128   |  after append 128   cap = 256
[0 ->  255] cap = 256   |  after append 256   cap = 512
[0 ->  511] cap = 512   |  after append 512   cap = 1024
[0 -> 1023] cap = 1024  |  after append 1024  cap = 1280
[0 -> 1279] cap = 1280  |  after append 1280  cap = 1696
[0 -> 1695] cap = 1696  |  after append 1696  cap = 2304
*/
