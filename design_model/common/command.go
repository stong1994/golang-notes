package common

import "fmt"

/**
命令模式
1. 定义
	命令模式将请求封装为对象，从而允许我们使用不同的请求，队列或日志请求参数化其他对象，并支持可撤销操作。
2. 实现步骤
	1. 定义命令接口
	2. 将每个对象的每个动作都封装为结构体，并实现接口命令
3. 使用场景
	批处理、任务队列、undo、redo等把具体命令封装到对象中使用的场合
4. 优点
	使我们的代码可扩展，因为我们可以添加新命令而无需更改现有代码。
	减少命令的调用者和接收者的耦合。
5. 缺点
	每个命令都要单独出来作为一个结构体，并实现接口

参考文章：https://www.geeksforgeeks.org/command-pattern/
 */

type Command interface {
	execute()
}

type Control struct{
	command Command
}

func (c *Control) Execute() {
	c.command.execute()
}

func (c *Control) SetCommand(command Command) {
	c.command = command
}

// 各种命令
type Light struct{}

func (l *Light) On() {
	fmt.Println("light is on")
}

func (l *Light) Off() {
	fmt.Println("light is off")
}

type LightOnCommand struct {
	light *Light
}

func (l *LightOnCommand) execute() {
	l.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (l *LightOffCommand) execute() {
	l.light.Off()
}

type Video struct{}

func (v *Video) On() {
	fmt.Println("video is on")
}

func (v *Video) Off() {
	fmt.Println("video is off")
}

type VideoOnCommand struct {
	video *Video
}

func (v *VideoOnCommand) execute() {
	v.video.On()
}

type VideoOffCommand struct {
	video *Video
}

func (v *VideoOffCommand) execute() {
	v.video.Off()
}


