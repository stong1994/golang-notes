package num21_30

import (
	"math"
	"testing"
)

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
*/
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{
			"test1",
			10,
			3,
			3,
		},
		{
			"test2",
			7,
			-3,
			-2,
		},
		{
			"test3",
			-1,
			1,
			-1,
		},
		{
			"test4",
			-1,
			-1,
			1,
		},
		{
			"test5",
			-2,
			-1,
			2,
		},
		{
			"test6",
			-1,
			7,
			0,
		},
		{
			"test7",
			-2147483648,
			1,
			-2147483648,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := divide(v.dividend, v.divisor)
			if got != v.want {
				t.Fatalf("want %d got %d", v.want, got)
			}
		})
	}
}

func divide(dividend int, divisor int) int {
	f := false
	if divisor == 0 {
		panic("divisor can not be 0")
	}
	if dividend == 0 {
		return 0
	}
	if dividend > 0 && divisor > 0 {
		f = true
	} else if dividend < 0 && divisor < 0 {
		divisor = 0 - divisor
		dividend = 0 - dividend
		f = true
	} else if dividend < 0 && divisor > 0 {
		dividend = 0 - dividend
	} else if dividend > 0 && divisor < 0 {
		divisor = 0 - divisor
	}

	nums := addTimes(dividend, divisor)

	if !f {
		if nums > math.MaxInt32+1 {
			nums = math.MaxInt32 + 1
		}
		nums = 0 - nums
	} else {
		if nums > math.MaxInt32 {
			nums = math.MaxInt32
		}
	}
	return nums
}

func addTimes(dividend, divisor int) int {
	times := 1

	if dividend < divisor {
		return 0
	} else if dividend < divisor+divisor {
		return 1
	}

	lastDivisor := divisor
	divisors := divisor + divisor
	for divisors <= dividend {
		times = times + times
		lastDivisor = divisors
		divisors = divisors + divisors
	}
	return times + addTimes(dividend-lastDivisor, divisor)
}
