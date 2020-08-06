# pprof实践

[学习博客](https://github.com/EDDYCJY/blog/blob/master/golang/2018-09-15-Golang%20%E5%A4%A7%E6%9D%80%E5%99%A8%E4%B9%8B%E6%80%A7%E8%83%BD%E5%89%96%E6%9E%90%20PProf.md)

*以下为实践+记录+摘抄。建议阅读上边的学习博客*

#### 一、通过WEB界面

在一个用gin框架做的web项目中，想要用`pprof`来剖析下性能。因为已经引入了`net/http`，于是在引入包的地方引入`_ "net/http/pprof"`。  
启动项目后，在网址上输入`http://localhost:8080/debug/pprof/`，结果报`404`...

好吧，果然还得动下脑子。

检测引入的`net/http`，发现这个`http`并没有检测端口，只是提供了一些常量，如`http.StatusOK`。项目中采用`gin`框架提供的引擎来监听端口，
所以即使引入了`_ "net/http/pprof"`，没有该`http`监听端口，所以`404`

于是在项目入口处添加以下代码后，即可访问。
```go
import (
	"net/http"
	_ "net/http/pprof"
)

go func() {
    http.ListenAndServe("0.0.0.0:8080", nil)
}()
```
`net/http/pprof.go`中的`init`函数如下：
```go
// net/http/pprof.go
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}
```
访问`http://localhost:8888/debug/pprof/`，得到许多子页面
- profile（CPU Profiling）: $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
- block（Block Profiling）：$HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
- goroutine：$HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
- heap（Memory Profiling）: $HOST/debug/pprof/heap，查看活动对象的内存分配情况
- mutex（Mutex Profiling）：$HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
- threadcreate：$HOST/debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪
- cmdline: $HOST/debug/pprof/cmdline, 查看项目的执行文件所在路径
- allocs: $HOST/debug/pprof/allocs?debug=1, 过去所有内存分配的样本
- trace: $HOST/debug/pprof/trace, 跟踪当前程序的执行，得到一个trace文件

#### 二、 通过交互式终端使用
几个参数说明：
- flat：给定函数上运行耗时
- flat%：同上的 CPU 运行耗时总比例
- sum%：给定函数累积使用 CPU 总比例
- cum：当前函数加上它之上的调用运行总耗时
- cum%：同上的 CPU 运行耗时总比例

1.go tool pprof http://localhost:8888/debug/pprof/profile?seconds=60
> 执行该命令后，需等待 60 秒（可调整 seconds 的值），pprof 会进行 CPU Profiling。结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。具体可执行 pprof help 查看命令说明
```
C:\Users\st598>go tool pprof http://localhost:8888/debug/pprof/profile?seconds=60
Fetching profile over HTTP from http://localhost:8888/debug/pprof/profile?seconds=60
Saved profile in C:\Users\st598\pprof\pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Jul 17, 2019 at 3:51pm (+07)
Duration: 1mins, Total samples = 10ms (0.017%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10ms, 100% of 10ms total
Showing top 10 nodes out of 15
      flat  flat%   sum%        cum   cum%
      10ms   100%   100%       10ms   100%  github.com/jinzhu/gorm.(*Scope).scan
         0     0%   100%       10ms   100%  CarrierBusinesss/app.handleCors.func1
         0     0%   100%       10ms   100%  CarrierBusinesss/app/route.(*agentRoute).GetAll
         0     0%   100%       10ms   100%  CarrierBusinesss/dao.(*AgentDao).GetAll
         0     0%   100%       10ms   100%  CarrierBusinesss/service.(*AgentService).Get
         0     0%   100%       10ms   100%  github.com/gin-gonic/gin.(*Context).Next
         0     0%   100%       10ms   100%  github.com/gin-gonic/gin.(*Engine).ServeHTTP
         0     0%   100%       10ms   100%  github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
         0     0%   100%       10ms   100%  github.com/gin-gonic/gin.LoggerWithConfig.func1
         0     0%   100%       10ms   100%  github.com/gin-gonic/gin.RecoveryWithWriter.func1
```
2.go tool pprof http://localhost:6060/debug/pprof/heap
```
C:\Users\st598>go tool pprof http://localhost:8888/debug/pprof/heap
Fetching profile over HTTP from http://localhost:8888/debug/pprof/heap
Saved profile in C:\Users\st598\pprof\pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
Type: inuse_space
Time: Jul 17, 2019 at 3:43pm (+07)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 14788.24kB, 100% of 14788.24kB total
Showing top 10 nodes out of 27
      flat  flat%   sum%        cum   cum%
13243.46kB 89.55% 89.55% 13243.46kB 89.55%  golang.org/x/net/webdav.(*memFile).Write
  519.03kB  3.51% 93.06%   519.03kB  3.51%  time.init.ializers
  513.69kB  3.47% 96.54%   513.69kB  3.47%  text/template/parse.(*Tree).newText
  512.05kB  3.46%   100%   512.05kB  3.46%  regexp/syntax.(*parser).newRegexp
         0     0%   100%   513.69kB  3.47%  CarrierBusinesss/app.(*App).route
         0     0%   100%   513.69kB  3.47%  CarrierBusinesss/app.NewApp
         0     0%   100%   512.05kB  3.46%  github.com/jinzhu/inflection.compile
         0     0%   100%   512.05kB  3.46%  github.com/jinzhu/inflection.init.0
         0     0%   100%   513.69kB  3.47%  github.com/swaggo/gin-swagger.CustomWrapHandler
         0     0%   100%   513.69kB  3.47%  github.com/swaggo/gin-swagger.WrapHandler
```
- inuse_space：分析应用程序的常驻内存占用情况
- alloc_objects：分析应用程序的内存临时分配情况

3.go tool pprof http://localhost:8888/debug/pprof/block
```
C:\Users\st598>go tool pprof http://localhost:8888/debug/pprof/block
Fetching profile over HTTP from http://localhost:8888/debug/pprof/block
Saved profile in C:\Users\st598\pprof\pprof.contentions.delay.001.pb.gz
Type: delay
Time: Jul 17, 2019 at 4:11pm (+07)
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
```
4.go tool pprof http://localhost:8888/debug/pprof/mutex
```
C:\Users\st598>go tool pprof http://localhost:8888/debug/pprof/block
Fetching profile over HTTP from http://localhost:8888/debug/pprof/block
Saved profile in C:\Users\st598\pprof\pprof.contentions.delay.002.pb.gz
Type: delay
Time: Jul 17, 2019 at 4:16pm (+07)
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
```

#### 三、PProf 可视化界面
创建一个`Benchmark`测试函数`data/data_test.go[BenchmarkAdd]`
```go
func BenchmarkAdd(b *testing.B) {
	var data []string
	for i := 0; i < b.N; i++ {
		data = append(data, fmt.Sprintf("hello %d", i))
	}
}
```
执行`go test -bench=. -cpuprofile=cpu.prof`  
结果输出如下，并生成`cpu.prof`文件
```
goos: windows
goarch: amd64
pkg: golang-learning/performance/pprof/data
BenchmarkAdd-16          3000000               583 ns/op
PASS
ok      golang-learning/performance/pprof/data  3.319s
```
启动pprof可视化界面有两种方式(需提前安装`graphviz`)
1.执行`go tool pprof cpu.prof`,然后输入`web`
```
D:\GO-SPACE\src\golang-learning\performance\pprof\data>go tool pprof cpu.prof
Type: cpu
Time: Jul 17, 2019 at 4:42pm (+07)
Duration: 2.51s, Total samples = 3.25s (129.25%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
```
2.执行`go tool pprof -http=:8888 cpu.prof`

方法2会自动打开浏览器，并访问`http://localhost:8888/ui/`，并有6个视图，分别是：
1. Top
2. Graph
3. Flame Graph
4. Peek
5. Source
6. Disassemble

方法1会生成一个图片，图片内容与方法2的第2个视图是一样的。即CPU在每个函数调用花费的时间。**线条越粗，占用比例越大**

其实在方法1中就已经有火焰图了，但是原文中，可能版本比较低，没有说明。并且给了第二种获取火焰图的方式。

1. 安装`PProf`：`$ go get -u github.com/google/pprof`
2. 启动PProf可视化界面 `$ pprof -http=:8888 cpu.prof`,然后发现和上述方法2的结果是一样的。

上边的可视化都是以cpu占用为例，如果要查看内存占用，可以使用`-memprofile`命令来生成对应的分析文件。如 `go test -bench=. -memprofile=mem.prof`