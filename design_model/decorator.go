package design_model

/*
装饰模式
1. 定义：
	动态地给一个对象增加一些额外的职责。
2. 实现步骤:
	定义组件
	定义装饰对象
	使用装饰对象生成组件
3. 使用场景:
	在不影响其他对象的情况下，以动态、透明的方式给单个对象添加职责。
4. 优点:
	可以通过一种动态的方式来扩展一个对象的功能，通过配置文件可以在运行时选择不同的装饰器，从而实现不同的行为。
 */
type DecoratorComponent interface {
	Cal(a, b int) int
}

type ConcreteComponent struct  {}

func (c *ConcreteComponent) Cal(a, b int) int {
	return a + b
}

// 装饰对象--增加减运算
type MulDecorate struct {
	DecoratorComponent
	num int
}

func WrapMulDecorate(component DecoratorComponent, num int) DecoratorComponent {
	return &MulDecorate{
		DecoratorComponent: component,
		num: num,
	}
}

func (m *MulDecorate) Cal(a, b int) int {
	return m.DecoratorComponent.Cal(a, b) + m.num
}