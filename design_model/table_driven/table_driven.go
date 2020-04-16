package table_driven

import (
	"fmt"
	"github.com/shopspring/decimal"
)
/*
表驱动模式
1. 定义：
	用一个表（map或者数组）来替代if-else或者switch-case
2. 优点
	逻辑和代码更简单、清晰
3. 缺点
	如果是map的话，因为需要对输入key计算哈希，在性能敏感的场合，需要多加斟酌。
*/
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

// 表驱动
var factories = map[string]IFloatOperator{
	"+": addOperator{},
	"-": subOperator{},
	"*": mulOperator{},
	"/": divOperator{},
}

func CreateOpFactory(opName string) IFloatOperator {
	if op, ok := factories[opName]; ok {
		return op
	}
	panic(fmt.Sprintf("operation is not valid: %s", opName))
}
