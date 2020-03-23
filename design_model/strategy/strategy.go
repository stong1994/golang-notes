package strategy

/**
策略模式
1. 定义
	策略行为设计模式允许在运行时选择算法的行为。
2. 实现步骤
	1. 定义策略
	2. 使用不同算法实现策略
	3. 构建策略实体时，将算法策略作为参数传入
3. 使用场景
	有不同的算法
4. 优点
	定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。
*/

// 定义策略
type StrategyOperator interface {
	Apply(a, b int) int
}

type StrategyOperation struct {
	operator StrategyOperator
}

func (so *StrategyOperation) Apply(a, b int) int {
	return so.operator.Apply(a, b)
}

// 实现不同策略
type Addition struct {}

func (Addition) Apply(a, b int) int {
	return a + b
}

type Multiplication struct{}

func (Multiplication) Apply(a, b int) int {
	return a * b
}