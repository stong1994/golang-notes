package design_model

/*
享元模式
1. 定义：
	享元模式通过共享技术实现相同或相似对象的重用。
2. 实现步骤
	定义享元对象
	定义享元工厂
	使用享元工厂创建
3. 使用场景
	一个系统有大量相同或者相似的对象，由于这类对象的大量使用，造成内存的大量耗费。
	对象的大部分状态都可以外部化，可以将这些外部状态传入对象中。
	使用享元模式需要维护一个存储享元对象的享元池，而这需要耗费资源，因此，应当在多次重复使用享元对象时才值得使用享元模式。
4. 优点
	享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元，使多个对象共享，从而节省内存以及减少对象数量。
5. 享元模式有如下角色:
	Flyweight: 抽象享元类
	ConcreteFlyweight: 具体享元类
	UnsharedConcreteFlyweight: 非共享具体享元类
	FlyweightFactory: 享元工厂类
6. 参考文章: https://blog.csdn.net/lkysyzxz/article/details/79541798
 */

// 定义享元对象--一个经常被重复使用的组件
type IAnimal interface {
	GetWeight() int64
	GetAge() int64
}

type Animal struct {
	Weight int64
	Age int64
}

func (a *Animal) GetWeight() int64 {
	return a.Weight
}

func (a *Animal) GetAge() int64 {
	return a.Age
}

func NewAnimal() *Animal {
	return &Animal{
		Weight: 0,
		Age: 0,
	}
}

type GrowUp struct {
	animal IAnimal
	WeightAdd int64
	AgeAdd int64
}

func (g *GrowUp) GetWeight() int64 {
	return g.animal.GetWeight() + g.WeightAdd
}

func (g *GrowUp) GetAge() int64 {
	return g.animal.GetAge() + g.AgeAdd
}

func NewGrowUp(animal IAnimal, weightAdd, ageAdd int64) *GrowUp {
	return &GrowUp{
		animal:animal,
		WeightAdd:weightAdd,
		AgeAdd:ageAdd,
	}
}

// 定义享元对象工厂
type Element struct {
	Value interface{}
}

func newElement(value interface{})*Element{
	return &Element{value}
}

type FlyweightFactory struct {
	pool map[string]*Element
}

func (this *FlyweightFactory) GetElement(key string) interface{} {
	if val, ok := this.pool[key]; ok {
		return val.Value
	}
	return nil
}

func (this *FlyweightFactory)SetElement(key string,value interface{}){
	ne := newElement(value)
	this.pool[key]=ne
}

func NewFlyweight()*FlyweightFactory{
	flyweight := FlyweightFactory{}
	flyweight.pool=make(map[string]*Element)
	return &flyweight
}