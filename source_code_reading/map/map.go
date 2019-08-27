package main

import (
	"fmt"
	"math"
)

// 如果map的键类型为float,且为math.Nan().那么将获取不到该键值.利用该特性可以做什么?

var floatMap = make(map[float64]int)

func float() {
	floatMap[math.NaN()] = 1
	floatMap[math.NaN()] = 2
	fmt.Println(floatMap)             // map[NaN:2 NaN:1]
	fmt.Println(floatMap[math.NaN()]) // 0
}
