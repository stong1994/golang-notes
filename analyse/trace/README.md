# trace实践

[学习博客](https://github.com/EDDYCJY/blog/blob/master/golang/2019-07-11-go-tool-trace.md)  
[go tool trace](https://making.pusher.com/go-tool-trace/)

### 使用方式：
在主程序前，增加trace
```go
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	// 主程序代码
}

```
 程序运行完才会产生`trace.out`文件，使用命令`go tool trace trace.out`即可打开UI界面。有9个子页面
- View trace：查看跟踪
- Goroutine analysis：Goroutine 分析
- Network blocking profile：网络阻塞概况
- Synchronization blocking profile：同步阻塞概况
- Syscall blocking profile：系统调用阻塞概况
- Scheduler latency profile：调度延迟概况
- User defined tasks：用户自定义任务
- User defined regions：用户自定义区域
- Minimum mutator utilization：最低 Mutator 利用率

先查看概括`Syscall blocking profile`,然后是协程分析`Goroutine analysis`,最后查看追踪`View trace`。  
追踪图可以用`WASD`四个按键来进行简单的缩放和移动，按`Shift+?`或者点右上角的问号来查看其他快捷键。