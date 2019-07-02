package design_model

import "fmt"

/**
访问者模式
1. 定义
	表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素类的前提下定义作用于这些元素的新操作。
2. 实现步骤
	1. 定义访问者接口，并根据情况（比如有两种数据库Pgsql与Mysql）实现该接口
	2. 定义结构体中元素的接口，并将访问者作为参数传入接收函数
	3. 实现元素的接口
	4. 在接收者中调用访问者
3. 使用场景
	1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
	2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。
4. 优点
	1、符合单一职责原则。 2、优秀的扩展性。 3、灵活性。
5. 缺点
	1、具体元素对访问者公布细节，违反了迪米特原则。 2、具体元素变更比较困难。 3、违反了依赖倒置原则，依赖了具体类，没有依赖抽象。
参考文章：https://juejin.im/post/5c00be786fb9a049c64392ad
 */

// 定义访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("这是测试环境")
}

// 定义元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

type VisitorElement struct {
}

func (el VisitorElement) Accept(visitor IVisitor) {
	visitor.Visit()
}

type EnvExample struct {
	VisitorElement
}

func (e EnvExample) Print(visitor IVisitor) {
	e.VisitorElement.Accept(visitor)
}