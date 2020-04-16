package iterator

import (
	"fmt"
	"testing"
)

func TestClosureIterator(t *testing.T) {
	var ints Integers = []int{1,2,3,4,5,6,7}
	iterator := ints.ClosureIterator()
	for {
		if val, ok := iterator(); ok {
			fmt.Println(val)
		}else {
			return
		}
	}
}

func TestIIterator(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	container := &ConcreteContainer{arr}
	iterator := container.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Current())
		iterator.Next()
	}
}