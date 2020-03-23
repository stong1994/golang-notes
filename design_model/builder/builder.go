package builder

import "fmt"

/**
创建者模式
1. 定义：
	将内部流程一步一步实现，同一个创建者可以构造不同的产品
2. 实现步骤：
	定义产品结构
	定义创建者（过程加结果）
	构造创建者
	构造不同的产品
3. 使用场景：
	相似的产品具有相似的流程，并且流程较复杂，可以用创建者模式
4. 优点：
	创建者模式将复杂的内部结构和创建分离开，用户不用在意内部细节就可创造出产品
	同一个构造流程能够构造不同产品
5. 缺点：
	要求这些产品有相同的或者相似的结构

相比较工厂模式，创建者模式更关注于一步一步创建。

参考文章：https://medium.com/@haluan/golang-builder-design-pattern-a8b7c92969a7
 */

// 背景：生产两款电子产品——笔记本电脑和手机，其中电脑需要一个显示器和一个摄像头，手机需要一个显示器和两个摄像头
// 所以我们有三个组件：成品，显示器，摄像头
// 组装电子产品分三步：1，定义成品，2，组装显示器，3，组装摄像头


// 定义电子产品
type ElectronicProduct struct {
	Structure string
	Monitor int
	Camera int
}

// 创建`创建者`过程
type BuildProcess interface {
	SetStructure() BuildProcess
	SetMonitor() BuildProcess
	SetCamera() BuildProcess
	GetGadget() ElectronicProduct
}

// 创建制造者
type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() ElectronicProduct {
	m.builder.SetStructure().SetCamera().SetMonitor()
	return m.builder.GetGadget()
}

func (m *ManufacturingDirector) PrintProduct() {
	gadget := m.builder.GetGadget()
	fmt.Println("structure", gadget.Structure)
	fmt.Println("monitor", gadget.Monitor)
	fmt.Println("camera", gadget.Camera)
	fmt.Println()
}

// 平板电脑
type Laptop struct {
	electronicProduct ElectronicProduct
}

func (l *Laptop) SetStructure() BuildProcess {
	l.electronicProduct.Structure = "Laptop"
	return l
}

func (l *Laptop) SetMonitor() BuildProcess {
	l.electronicProduct.Monitor = 1
	return l
}

func (l *Laptop) SetCamera() BuildProcess {
	l.electronicProduct.Camera = 1
	return l
}

func (l *Laptop) GetGadget() ElectronicProduct {
	return l.electronicProduct
}

// 手机
type SmartPhone struct {
	electronicProduct ElectronicProduct
}

func (s *SmartPhone) SetStructure() BuildProcess {
	s.electronicProduct.Structure = "Laptop"
	return s
}

func (s *SmartPhone) SetMonitor() BuildProcess {
	s.electronicProduct.Monitor = 1
	return s
}

func (s *SmartPhone) SetCamera() BuildProcess {
	s.electronicProduct.Camera = 2
	return s
}

func (s *SmartPhone) GetGadget() ElectronicProduct {
	return s.electronicProduct
}

/*
1. 给制造者提供一个产品类型
2. 让制造者开始制造，而不关心细节
3. 制造完后打印结果
 */
func BuilderModel()  {
	manufacturingDirector := ManufacturingDirector{}
	laptop := &Laptop{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()

	phone := &SmartPhone{}
	manufacturingDirector.SetBuilder(phone)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
}