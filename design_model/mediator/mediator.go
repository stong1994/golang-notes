package mediator

import "fmt"

/**
中介者模式
1. 定义
	用一个中介对象来封装一系列的对象交互，中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互
2. 实现步骤
	1. 定义不同的对象。
	2. 定义中介，在中介控制不同对象的交互。
3. 使用场景
	对于一个模块，可能由很多对象构成，而且这些对象之间可能存在相互的引用，为了减少对象两两之间复杂的引用关系，使之成为一个松耦合的系统，我们需要使用中介者模式，这就是中介者模式的模式动机。
4. 优点
	解耦合，减少对象间复杂的引用
5. 缺点
	有可能使系统更加复杂

参考文章：
	https://sourcemaking.com/design_patterns/mediator
	https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/mediator.html
	https://github.com/godsarmy/golang_design_pattern/blob/master/src/mediator.go
 */

type User1 struct {
	mediator *Mediator
	name string
}

func (u *User1) SendMsg(msg string) {
	fmt.Printf("i'm %s, i sent: %s\n", u.name, msg)
	u.mediator.ReceiveMsg(u, msg)
}
func (u *User1) ReceiveMsg(msg string) {
	fmt.Printf("i'm %s, i received: %s\n", u.name, msg)
}

type User2 struct {
	mediator *Mediator
	name string
}

func (u *User2) SendMsg(msg string) {
	fmt.Printf("i'm %s, i sent: %s\n", u.name, msg)
	u.mediator.ReceiveMsg(u, msg)
}
func (u *User2) ReceiveMsg(msg string) {
	fmt.Printf("i'm %s, i received: %s\n", u.name, msg)
}

type Mediator struct {
	u1 *User1
	u2 *User2
}

func (m *Mediator) SendMsg(u interface{}, msg string)  {
	if _, ok := u.(*User1); ok{
		m.u1.ReceiveMsg(msg)
	}else {
		m.u2.ReceiveMsg(msg)
	}
}

func (m *Mediator) ReceiveMsg(u interface{}, msg string) {
	if _, ok := u.(*User1); ok{
		m.SendMsg(m.u2, msg)
	}else {
		m.SendMsg(m.u1, msg)
	}
}