package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

// TODO 面试有一题为写一段代码来对程序的性能进行监控，可从pprof和runtime分析。 原谅我太菜。。。
func main() {
	var wg sync.WaitGroup
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			fmt.Println("nil bytes")
		}
	}

	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	makeMem()
	wg.Add(1)
	wg.Wait()
}

func bigBytes() *[]byte {
	s := make([]byte, 1000000)
	return &s
}

func makeMem() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			fmt.Println("nil bytes")
		}
	}

	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)
}
