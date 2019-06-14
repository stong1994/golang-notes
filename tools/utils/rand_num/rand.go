package rand_num

import (
	"crypto/rand"
	"math/big"
)

// 获取随机值
func RandNum(max int64) int64 {
	index, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return RandNum(max)
	}
	i := index.Int64()
	if i < 0 {
		return RandNum(max)
	}
	return i
}
