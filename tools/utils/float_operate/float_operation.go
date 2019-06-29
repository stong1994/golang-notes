package float

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// 直接计算(工厂模式胡里花哨)
func Add(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Float64()
	return val
}

func Sub(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Float64()
	return val
}

func Mul(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Float64()
	return val
}

func Div(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Float64()
	return val
}

// 复杂计算(按照顺序计算，乘和除优先级与加减一样)
func MultiCalc(params ...interface{}) float64 {
	if len(params) < 3 {
		panic("multiCalc must have three params at least")
	}
	// 如果长度为偶数，报错
	if len(params)&1 == 0 {
		panic(fmt.Sprintf("params length can not be even, but got length is %d", len(params)))
	}

	var (
		secondParam float64
		opName      string
		val         = decimal.NewFromFloat(params[0].(float64))
	)

	for index := 1; index < len(params); index++ {
		if index&1 == 0 { // index为偶数，为float64
			secondParam = params[index].(float64)
			val = calc(val, secondParam, opName)
		} else { // index为奇数，为计算符号
			opName = params[index].(string)
		}
	}
	result, _ := val.Float64()
	return result
}

func calc(a decimal.Decimal, b float64, opName string) decimal.Decimal {
	switch opName {
	case "+":
		return a.Add(decimal.NewFromFloat(b))
	case "-":
		return a.Sub(decimal.NewFromFloat(b))
	case "*":
		return a.Mul(decimal.NewFromFloat(b))
	case "/":
		return a.Div(decimal.NewFromFloat(b))
	default:
		panic(fmt.Sprintf("opName is not valid: %s ", opName))
	}
}

// 工厂模式——胡里花哨。。。
type IFloatOperator interface {
	Calc(a, b float64) float64
}

type addOperator struct{}

func (addOperator) Calc(a, b float64) float64 {
	sum, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Float64()
	return sum
}

type subOperator struct{}

func (subOperator) Calc(a, b float64) float64 {
	sub, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Float64()
	return sub
}

type mulOperator struct{}

func (mulOperator) Calc(a, b float64) float64 {
	mul, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Float64()
	return mul
}

type divOperator struct{}

func (divOperator) Calc(a, b float64) float64 {
	div, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Float64()
	return div
}

var factories = map[string]IFloatOperator{"+": addOperator{}, "-": subOperator{}, "*": mulOperator{}, "/": divOperator{}}

func CreateOpFactory(opName string) IFloatOperator {
	if op, ok := factories[opName]; ok {
		return op
	}
	panic(fmt.Sprintf("operation is not valid: %s", opName))
}
