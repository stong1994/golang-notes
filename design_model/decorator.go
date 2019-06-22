package design_model

/*
装饰模式
 */
type Component interface {
	Sum(a, b int) int
}

type ConcreteComponent struct  {
	
}

func (c *ConcreteComponent) Sum(a, b int) int {
	return a + b
}

// 装饰对象
type DecorateComponent struct {
	Component
	num int
}

func WrapDecorateComponent()