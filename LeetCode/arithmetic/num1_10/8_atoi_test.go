package num_1_10

import (
	"fmt"
	"strings"
	"testing"
)

func TestNum8(t *testing.T) {
	str := "4193 with words"
	target := 4193
	result := myAtoi(str)
	if target == result {
		fmt.Println("true")
	}
}

func myAtoi(str string) int {
	//notSpanStr := strings.Replace(str, " ","", -1)
	notSpanStr := strings.TrimLeft(str, "! ")
	if len(notSpanStr) == 0 {
		return 0
	}
	head := string(notSpanStr[0])
	if head == "+" && len(str) > 1 && isNum(notSpanStr[1]) {
		n := handStr(notSpanStr[1:], 1)
		a := +n
		return a
	} else if head == "-" && len(str) > 1 && isNum(notSpanStr[1]) {
		n := handStr(notSpanStr[1:], -1)
		a := -n
		return a
	} else if isNum(byte(notSpanStr[0])) {
		n := handStr(notSpanStr, 0)
		return n
	}
	return 0
}

func isNum(str byte) bool {
	str -= '0'
	if str < 0 || str > 9 {
		return false
	}
	return true
}

func handStr(str string, flag int) int {
	n := 0
	symbol := byte('.')
	for _, v := range []byte(str) {
		if v == symbol {
			return n
		}
		if !isNum(v) {
			return n
		}
		v -= '0'
		if v <= 9 {
			m := 10*n + int(v)
			if flag >= 0 && m >= 2147483647 {
				return 2147483647
			}
			if flag < 0 && m >= 2147483648 {
				return 2147483648
			}
			n = m
		}
	}
	return n
}
