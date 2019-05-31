package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	l := sync.Once{}
	for i := 0; i < 10; i ++ {
		go l.Do(print)
	}
	time.Sleep(time.Second)
}

func print()  {
	fmt.Println(rand.Intn(10))
}
