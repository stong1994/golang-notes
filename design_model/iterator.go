package design_model

/**
迭代模式
1. 定义
	能够顺序访问容器中的每个元素，而外部无需知道底层实现。
2. 实现步骤
	1. 定义迭代器接口
	2. 创建具体容器
	3. 创建迭代器（迭代器中包含具体容器）并实现迭代器接口
3. 使用场景
	顺序访问容器中的元素
4. 优点
	不关注内部实现细节
5。 缺点
	迭代模式的意义在于不关注内部细节就能遍历。但是golang中的slice、map遍历都很简单，不需要再次封装。使用场景上较少
参考文章：https://juejin.im/post/5c0a03f56fb9a049c2323d54
 */

// 闭包的迭代模式
type Integers []int

func (i Integers) ClosureIterator() func() (int, bool) {
	index := 0
	return func() (int, bool) {
		if index >= len(i) {
			return 0, false
		}
		v := i[index]
		index++
		return v, true
	}
}

// 正常的迭代模式
// 迭代器接口
type IIterator interface {
	HasNext() bool
	Current() int
	Next() bool
}

type Iterator struct {
	cursor int // 游标
	container *ConcreteContainer
}

func (i *Iterator) HasNext() bool {
	if i.cursor >= len(i.container.container) {
		return false
	}
	return true
}

func (i *Iterator) Current() int {
	return i.container.container[i.cursor]
}

func (i *Iterator) Next() bool {
	if i.cursor >= len(i.container.container) {
		return false
	}
	i.cursor++
	return true
}

// 具体容器
type ConcreteContainer struct {
	container []int
}

func (c *ConcreteContainer) Iterator() IIterator {
	i := new(Iterator)
	i.container = c
	return i
}