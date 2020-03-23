package adapter
/*
适配器模式
1. 定义： 将一个接口转换成客户希望的另一个接口。
2. 实现步骤
	定义被适配的接口和实现
	定义目标接口和实现，并且实现接口由被适配接口创建
3. 使用场景
	系统需要使用现有的类，而这些类的接口不符合系统的需要。
4. 优点
	将目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，而无须修改原有代码。
5. 局限
	例子中原接口有返回值，需要的新的接口没有返回值，因此直接调用不返回即可，但是如果反过来呢？因此，在接口定义时，一定要尽量合理。
 */


// 定义被适配的接口(旧的接口)
type Adapter interface {
	Request() string
}

type adapter struct {
}

func (adapter) Request() string {
	return ""
}

func NewAdapter() Adapter {
	return &adapter{}
}

// 定义目标接口(新的接口)
type Target interface {
	TargetRequest()
}

func New(adapter Adapter) Target {
	return &target{adapter}
}

type target struct {
	Adapter
}

func (t *target) TargetRequest() {
	t.Request()
}
