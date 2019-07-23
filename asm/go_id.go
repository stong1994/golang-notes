package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	fmt.Println(id)
}

// 汇编实现：https://github.com/cch123/goroutineid
