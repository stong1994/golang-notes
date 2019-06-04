package main

import (
	"fmt"
	"sync"
	"time"
)

// 在高并发下同时读写map会报错
func main() {
	m := &sync.Map{}
	for i := 0; i < 100000; i++ {
		go write(m, i, i*i)
		go read(m, i)
	}
	time.Sleep(time.Second)
	// 打印输出
	m.Range(load)
}

func read(m *sync.Map, key int) (interface{}, bool) {
	return m.Load(key)
}

func write(m *sync.Map, key, val int) {
	m.Store(key, val)
}

func load(key, value interface{}) bool {
	fmt.Printf("key: %d, value: %v\n", key, value)
	return true
}
