package dynamic_programming

import (
	"fmt"
	"math"
	"testing"
)

// 双11购物有优惠：满200减50
// 假设女朋友的购物车中有n个物品，她想要从中选几个，使商品的价格总和最大程度上接近200元，来达到撸羊毛的效果。
// 求应该购买哪些物品

// 分析：关键点是满200，和减50没有任何关系。也就是说购买的商品价格要不小于200且最接近200
const thresholdPrice = 200

var thingPrices = []int{55, 65, 46, 110, 95}

// 回溯法
func TestDouble11_1(t *testing.T) {
	set := make(map[int][]int) // 价格-商品列表
	p := math.MaxInt64
	var f func(currP, currI int, ts []int) // 当前价格，当前索引, 商品列表
	f = func(currP, currI int, ts []int) {
		if currP >= thresholdPrice || currI >= len(thingPrices) {
			if currP > thresholdPrice {
				set[currP] = ts
				if currP < p {
					p = currP
				}
			}
			return
		}
		// 跳过这个商品
		f(currP, currI+1, ts)
		// 选择这个商品
		f(thingPrices[currI]+currP, currI+1, append(ts, currI))
	}
	f(0, 0, []int{})
	fmt.Println(set)
	fmt.Println(p)
	fmt.Println(set[p])
}

// 动态规划
// 动态规划需要定义数组，所以需要控制最大长度
// 规定购物花费的金额不能超过满减的3倍，即600
func TestDouble11_2(t *testing.T) {
	m := make(map[int][]int) // 价格-商品列表
	const totalNum = 5
	const maxLen = thresholdPrice * 3
	arr := [totalNum][]int{}

}
