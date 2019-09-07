package main

import "fmt"

/*
模仿selectgo中的初始化pollorder和lockorder
pollorder := order1[:ncases:ncases]
lockorder := order1[ncases:][:ncases:ncases]
*/
func main() {
	order1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ncases := 4
	pollorder := order1[:ncases:ncases]
	fmt.Println("pollorder")
	fmt.Println(pollorder)
	fmt.Println("len", len(pollorder))
	fmt.Println("cap", cap(pollorder))
	lockorder := order1[ncases:][:ncases:ncases]
	fmt.Println("lockorder")
	fmt.Println(lockorder)
	fmt.Println("len", len(lockorder))
	fmt.Println("cap", cap(lockorder))
}

/* 输出：
pollorder
[1 2 3 4]
len 4
cap 4
lockorder
[5 6 7 8]
len 4
cap 4
*/
