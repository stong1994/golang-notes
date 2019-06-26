package num_1_10

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNum9(t *testing.T) {
	i := 121
	target := true
	result := isPalindrome(i)
	if target == result {
		fmt.Println("true")
	}
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	byteX := []byte(str)
	maxLen := len(byteX)
	for i := 0; i < maxLen/2; i++ {
		if byteX[i] != byteX[maxLen-i-1] {
			return false
		}
	}
	return true
}
