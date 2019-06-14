package sort

import "sort"

// 封装sort.Sort()方法，避免在需要对多个不同属性进行排序时，建立多个实现
type sorter struct {
	len  int
	swap func(i, j int)
	less func(i, j int) bool
}

func (s sorter) Len() int           { return s.len }
func (s sorter) Swap(i, j int)      { s.swap(i, j) }
func (s sorter) Less(i, j int) bool { return s.less(i, j) }

func Sort(n int, swap func(i, j int), less func(i, j int) bool) {
	sort.Sort(sorter{len: n, swap: swap, less: less})
}
