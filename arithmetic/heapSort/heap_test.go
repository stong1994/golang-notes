package heapSort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type data []int

func (d *data) Len() int {
	return len(*d)
}

func (d *data) Less(i, j int) bool {
	return (*d)[i] < (*d)[j]
}

func (d *data) Swap(i, j int) {
	(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
}

func (d *data) Push(x Val) {
	*d = append((*d), x.(int))
}

func (d *data) Pop() Val {
	if d == nil || len(*d) == 0 {
		return nil
	}
	val := (*d)[len(*d)-1]
	*d = append((*d)[:len(*d)-1])
	return val
}

func init()  {
	rand.NewSource(time.Now().UnixNano())
}

/*
Pop(), 获取最小元素，并删除，由于是最小堆，那么获取的是首元素
Remove(),获取最后一个元素，由于是最大堆，那么获取的是尾元素
 */
func TestInit(t *testing.T) {
	var d data
	for i := 0; i < 20; i++ {
		d = append(d, rand.Intn(50))
	}
	t.Log("排序前", d)
	Init(&d)
	t.Log("排序后", d)
	for i := 0; i < 20; i++ {
		fmt.Println(Remove(&d, len(d)-1))
		fmt.Println(d)
	}
}
