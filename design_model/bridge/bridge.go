package bridge

import "fmt"

/**
桥接模式
1. 定义：
	当一个系统中，包含一个组件，该组件是可变的，该系统是可变的。这个时候就需要一个桥连接抽象的系统和抽象的组件。
2. 实现步骤：
	定义组件的接口并实现
	定义一个系统并包含组件的接口
3. 使用场景：
	不希望在抽象和它的实现部分之间有一个固定的绑定关系。比如这种情况可能是因为在程序运行时刻实现部分应可以被选择或者切换。
	类的抽象以及它的实现应该可以通过生成子类的方法加以扩充。
	对一个抽象的实现部分的修改应对客户不产生影响，即客户的代码不必重新编译。
	一个类型存在两个独立变化的维度，且这两个维度都需要进行扩展。
4. 优点：
	分离接口及其实现部分
	提高可扩展性

参考文章：http://blog.ralch.com/tutorial/design-patterns/golang-bridge/
 */

// 构造图形页面可能使用OpenGL 或者Direct2D 两种渲染引擎，同时我们也需要渲染不同的形状。在这个例子中有两个独立的维度：引擎和图形
// 构造一个图形——圆形
type CircleDrawer struct {
	DrawContext Drawer
}

func NewCircleDrawer(DrawContext Drawer) *CircleDrawer {
	return &CircleDrawer{DrawContext}
}

func (c *CircleDrawer) Draw()  {
	c.DrawContext.DrawEclipse()
	return
}

// 构造组件——渲染工具
type Drawer interface {
	DrawEclipse()
}

type OpenGL struct {}

func (OpenGL) DrawEclipse() {
	fmt.Println("draw with OpenGL")
}

type Direct2D struct {}

func (Direct2D) DrawEclipse() {
	fmt.Println("draw with Direct2D")
}

func DrawCircle()  {
	openGLDrawer := NewCircleDrawer(OpenGL{})
	direct2DDrawer := NewCircleDrawer(Direct2D{})

	// 用openGL来渲染圆形
	openGLDrawer.Draw()
	// 用Direct2D来渲染圆形
	direct2DDrawer.Draw()
}