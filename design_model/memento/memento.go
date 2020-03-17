package memento

/**
备忘录模式
1. 定义
	在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。
2. 实现步骤
	1. 定义对象及状态
	2. 封装要保存的状态到对象中
	3. 恢复时，将保存的对象中的属性赋给当前对象
3. 使用场景
	可以用来创建程序某个时刻运行状态的快照，当程序异常崩溃或者因为其他原因导致退出后，可以使用备忘后的数据，恢复到原始状态，最常见的操作应该就是编辑器的撤销了，编辑器应用了备忘录模式，将编辑过程中的代码状态放在一个状态栈中，当使用ctrl+z 的时候，就从栈中弹出上一次保存的状态，来恢复到上一次的情况（即撤销）。
 */

type Circle struct {
	color string
}

func (c *Circle) DrawColor(color string) {
	c.color = color
}

func (c *Circle) Save() *Circle {
	return &Circle{color:c.color}
}

func (c *Circle) Redo(circle *Circle) {
	c.color = circle.color
}