package design_model

import "fmt"

/*
工厂模式
参考文章：https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
 */

// 简单工厂模式 Simple Factory
type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Println("Hi! My name is ", p.Name)
}

// 新建对象用 NewPerson()，而不用&Person{}。这样能避免忘记对name或者age进行赋值
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

// 接口工厂 Interface factories
type IPerson interface {
	Greet()
}

type person struct {
	name string
	age int
}

func (p person) Greet() {
	fmt.Println("Hi! My name is ", p.name)
}

// 返回接口，能够保护内部变量
func NewIPerson(name string, age int) IPerson {
	return person{
		name: name,
		age: age,
	}
}

// 工厂方法 Factory methods
type Animall struct {
	species string
	age int
}

type AnimalHouse struct {
	name string
	sizeInMeters int
}

type AnimalFactory struct {
	species string
	houseName string
}

func NewAnimalFactory(species, houseName string) *AnimalFactory {
	return &AnimalFactory{
		species: species,
		houseName: houseName,
	}
}

func (af AnimalFactory) NewAnimal(age int) *Animall {
	return &Animall{
		species: af.species,
		age: age,
	}
}

func (af AnimalFactory) NewHouse(sizeInMeters int) *AnimalHouse {
	return &AnimalHouse{
		name: af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

func DoIt()  {
	dogFactory := NewAnimalFactory("dog", "kennel")
	dog := dogFactory.NewAnimal(1)
	kennel := dogFactory.NewHouse(3)

	catFactory := NewAnimalFactory("cat", "cattery")
	cat := catFactory.NewAnimal(2)
	cattery := catFactory.NewHouse(2)
	_, _, _, _ = dog, kennel, cat, cattery
}

// 工厂函数 Factory functions
type Pen struct {
	name string
	length int
}

func NewPenFactory(name string) func(length int) Pen {
	return func(length int) Pen {
		return Pen{
			name: name,
			length: length,
		}
	}
}

func DoIt_v2()  {
	pencilFactory := NewPenFactory("pencil")
	bigPencil := pencilFactory(10)

	chalkFactory := NewPenFactory("chalk")
	bigChalk := chalkFactory(5)

	_, _ = bigPencil, bigChalk
}

// ============= 国内资料的工厂模式 ============== //
type Operator interface {
	Operate(int, int) int
}

type AddOperate struct {}

func (AddOperate) Operate(a, b int) int {
	return a + b
}

type MulOperate struct {}

func (MulOperate) Operate(a, b int) int {
	return a * b
}

type OperatorFactory struct {}

func CreateOperator(operatorName string) Operator {
	switch operatorName {
	case "+":
		return AddOperate{}
	case "-":
		return MulOperate{}
	default:
		panic("operatorName is not valid")
	}
}

func DoIt_v3()  {
	addOperator := CreateOperator("+")
	sum := addOperator.Operate(1, 2)

	mulOperator := CreateOperator("-")
	mul := mulOperator.Operate(1, 2)

	_, _ = sum, mul
}