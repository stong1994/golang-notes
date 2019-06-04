package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := sync.Pool{New: random}
	for i := 5; i < 10; i++ {
		p.Put(i)
	}
	for i := 0; i < 20; i++ {
		fmt.Println("random num is ", p.Get())
	}
}

// 当p.Get()获取到的值为nil时，获取此数据
func random() interface{} {
	return rand.Intn(5)
}
