package main

import (
	"fmt"
	"sync"
)

func main() {
	m := &sync.Map{}

	// 循环增加key-value
	for i := 0; i < 10; i++ {
		m.Store(i, i*i)
	}

	// 打印输出
	m.Range(load)

	// 获取i为0的值
	if val, ok := m.Load(0); ok {
		fmt.Println("0对应的value为", val)
	} else {
		fmt.Println("没有获取到key为0的值")
	}

	// 删除key为0的键值对
	m.Delete(0)

	// 再次获取i为0的值
	if val, exit := m.LoadOrStore(0, 1); exit {
		fmt.Println("0对应的值为", val)
	} else {
		fmt.Println("0没有对应的值，设置为1")
	}

	// 再次设置key为0的键值对
	if val, exit := m.LoadOrStore(0, 2); exit {
		fmt.Println("0对应的值为", val)
	} else {
		fmt.Println("0没有对应的值，设置为2")
	}

	// 最后加载key为0的键值对
	if val, ok := m.Load(0); ok {
		fmt.Println("0对应的value为", val)
	} else {
		fmt.Println("没有获取到key为0的值")
	}
}

func load(key, value interface{}) bool {
	fmt.Printf("key: %d, value: %d\n", key, value)
	return true
}
