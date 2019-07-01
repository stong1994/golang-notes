package design_model

import "fmt"

/**
状态模式
1. 定义
	当一个对象的内部状态发生改变时，会导致其行为的改变，对象看起来似乎修改了它的类
2. 使用步骤
	1. 定义状态
	2. 根据不同状态构建不同的函数
	3. 用户调用时，根据不同状态调用不同的函数
3. 使用场景
	根据状态来调整行为的场景，如工作流或游戏等
4. 优点
	允许状态转换逻辑与状态对象合成一体，而不是提供一个巨大的条件语句块，状态模式可以避免使用庞大的条件语句来将业务方法和状态转换代码交织在一起
5. 缺点
	状态模式对“开闭原则”的支持并不太好，增加新的状态类需要修改那些负责状态转换的源代码，否则无法转换到新增状态；而且修改某个状态类的行为也需修改对应类的源代码。
 */

type State int
const (
	Start State = iota
	First
	Second
	Third
	End
)

type Events struct {
	state State
}

func (e *Events) Alloc()  {
	switch e.state {
	case Start:
		e.Start()
	case First:
		e.First()
	case Second:
		e.Second()
	case Third:
		e.Third()
	case End:
		e.End()
	default:
		fmt.Println("all done")
		return
	}
	e.state++
}

func (e Events) Start() {
	fmt.Println("start")
}

func (e Events) First() {
	fmt.Println("first")
}

func (e Events) Second() {
	fmt.Println("second")
}

func (e Events) Third() {
	fmt.Println("third")
}

func (e Events) End() {
	fmt.Println("end")
}