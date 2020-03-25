package data_structure

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue(t *testing.T) {
	obj := Constructor(6)
	param_1 := obj.EnQueue(6)
	param_2 := obj.Rear()
	param_3 := obj.Rear()
	param_4 := obj.DeQueue()
	param_5 := obj.EnQueue(4)
	param_6 := obj.Rear()
	param_7 := obj.DeQueue()
	param_8 := obj.Front()
	param_9 := obj.DeQueue()
	param_10 := obj.DeQueue()
	param_11 := obj.DeQueue()
	fmt.Println(param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9, param_10, param_11)
}