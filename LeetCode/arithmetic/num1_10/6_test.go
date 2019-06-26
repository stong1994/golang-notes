package num_1_10

import (
	"fmt"
	"testing"
)

func TestNum6(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	result := convert(s, numRows)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	lenStr := len(s)
	if lenStr == 0 || numRows <= 1 {
		return s
	}
	batchSize := 2*numRows - 2
	// todo 漏掉最后一个batch
	result := make([][]uint8, numRows)
	batchNum := lenStr / batchSize
	for i := 0; i < batchNum; i++ {
		for j := 0; j < batchSize; j++ {
			mod := j % numRows
			index := i*batchSize + j
			//str := string(s[index])
			//_= str
			if j < numRows {
				result[mod] = append(result[mod], s[index])
			} else {
				row := numRows - (mod + 1) - 1
				result[row] = append(result[row], s[index])
			}
		}
	}
	// 最后一个batch
	mod := lenStr % batchSize
	for i := 0; i < mod; i++ {
		if i < numRows {
			result[i] = append(result[i], s[batchSize*batchNum+i])
		} else {
			mod2 := i % numRows
			row := numRows - (mod2 + 1) - 1
			result[row] = append(result[row], s[batchSize*batchNum+i])
		}

	}
	resultStr := ""
	for _, v := range result {
		for _, val := range v {
			resultStr += string(val)
		}
	}
	return resultStr
}
