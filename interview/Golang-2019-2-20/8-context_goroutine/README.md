* main.go: 利用context和goroutine来爬取网站
* gen-num: 用goroutine来顺序生成数字，并用context来取消goroutine，防止goroutine泄露
* same-data-with-diff-url: 是在之前遇到过一个需求：一个数据有多个数据源，但是只取一个返回速度最快的。完成了需求实现。


# context的使用规范

* Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first
parameter, typically named ctx；不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx；

* Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to
use；即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO；

* Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to
functions；使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数；

* The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple
goroutines；同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全的；

