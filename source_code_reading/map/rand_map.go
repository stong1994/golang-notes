package main

import "fmt"

func iterator() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for v := range m {
		fmt.Printf("%d\t", v)
	}
}
// 0x0162 00354 (rand_map.go:7)    CALL    runtime.mapiterinit(SB)

//  0x01fc 00508 (rand_map.go:7)    CALL    runtime.mapiternext(SB)