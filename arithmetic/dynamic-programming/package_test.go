package dynamic_programming

import (
	"fmt"
	"testing"
)

/*
https://time.geekbang.org/column/article/74788
背包问题
背包能承受的最大重量为15KG
现在有5个物品，分别是1kg、1kg、6kg、6kg、10kg
求背包中最大的重量
*/
const totalN = 5
const maxW = 15

var things = [5]int{1, 1, 6, 6, 10}

// 回溯算法实现
func TestNormal(t *testing.T) {
	result := 0
	var f func(currW int, currI int)
	f = func(currW int, currI int) {
		if currW >= maxW || currI >= totalN {
			if currW > result {
				result = currW
			}
			return
		}
		// 选择跳过当前物品
		f(currW, currI+1)
		// 选择当前物品
		if currW+things[currI] <= maxW {
			f(currW+things[currI], currI+1)
		}
	}

	f(0, 0)
	fmt.Println(result)
}

// 时间复杂度是 O(n*w)。n 表示物品个数，w 表示背包可以承载的总重量。
func TestDynamic(t *testing.T) {
	// 二维数组，选择当前物品的状态  arr[i][w] = bool  i表示第几个物品，w表示当前状态的重量，bool表示是否存在
	/*
		  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
		0 1 1
		1 1 1 1
		2 1 1 1       1 1 1
		3 1 1 1       1 1 1          1  1  1
		4 1 1 1       1 1 1    1  1  1  1  1
	*/
	arr := [totalN][maxW + 1]bool{}
	arr[0][0] = true
	if things[0] < maxW {
		arr[0][things[0]] = true
	}

	for i := 1; i < totalN; i++ {
		// 根据上层来更新当前层的状态
		for j := 0; j < maxW+1; j++ {
			// 将上层状态延续下来
			if arr[i-1][j] {
				arr[i][j] = true
				// 更新当前物品加入后的状态
				if j+things[i] <= maxW {
					arr[i][j+things[i]] = true
				}
			}
		}
	}
	//打印结果
	//fmt.Println("  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14")
	//for i := 0; i < totalN; i ++ {
	//	fmt.Print(i)
	//	fmt.Print(" ")
	//	for j := 0; j < maxW+1; j++ {
	//		if !arr[i][j] {
	//			fmt.Print("  ")
	//			continue
	//		}
	//		fmt.Print("1")
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}

	// 获取最大重量
	for j := maxW; j >= 0; j-- {
		if arr[totalN-1][j] {
			fmt.Println(j)
			break
		}
	}
}

// 上边使用了二维数组，其实只使用一维数组即可
func TestDynamic2(t *testing.T) {
	arr := [maxW + 1]bool{}

	arr[0] = true
	if things[0] <= maxW {
		arr[things[0]] = true
	}
	for i := 1; i < totalN; i++ {
		// 在已有的基础上增加该物品（必须逆序，否则前边改变数组，后边会获取到刚改变的值）
		for j := maxW - things[i]; j >= 0; j-- {
			if arr[j] {
				arr[j+things[i]] = true
			}
		}
	}
	fmt.Println("arr", arr)
	for i := maxW; i >= 0; i-- {
		if arr[i] {
			fmt.Println(i)
			break
		}
	}
}

// 对每个物品加入价值
var prices = [5]int{5, 3, 4, 6, 20}

func TestDynamic3(t *testing.T) {
	var w, p int
	var f func(currW, currI, currP int)
	f = func(currW, currI, currP int) {
		if currW >= maxW || currI >= totalN {
			if currP > p {
				p = currP
				w = currW
			}
			return
		}
		// 不把这个物品放入背包
		f(currW, currI+1, currP)
		// 将这个物品放入背包
		if currW+things[currI] <= maxW {
			f(currW+things[currI], currI+1, currP+prices[currI])
		}
	}
	f(0, 0, 0)
	fmt.Println(w, p)
}

// 使用动态规划实现上边
func TestDynamic4(t *testing.T) {
	// todo
}
