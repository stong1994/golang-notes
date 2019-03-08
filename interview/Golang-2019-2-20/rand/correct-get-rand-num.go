package main

import (
	"crypto/rand"
	"math/big"
)

// 获取随机值的正确方法
func NewRandInt(max int64) int64 {
	index, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return NewRandInt(max)
	}
	i := index.Int64()
	if i < 0 {
		return NewRandInt(max)
	}
	return i
}
