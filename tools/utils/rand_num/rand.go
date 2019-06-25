package rand_num

import (
	cRand "crypto/rand"
	"fmt"
	"math/big"
	mRand "math/rand"
	"time"
)

// 获取随机值
func CryptoRandNum(max int64) int64 {
	index, err := cRand.Int(cRand.Reader, big.NewInt(max))
	if err != nil {
		return CryptoRandNum(max)
	}
	return index.Int64()
}

func MathRandNum(max int64) int64 {
	mRand.Seed(time.Now().Unix())
	return mRand.Int63n(max)
}

func RandNum(max int64) int64 {
	index, err := cRand.Int(cRand.Reader, big.NewInt(max))
	if err != nil {
		return MathRandNum(max)
	}
	return index.Int64()
}
