# errgroup


#### 场景
用于在一组`groutine`中,获取`groutine`执行过程中发生的错误


#### 1.结构体
```go
// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}
```
`Group`中内嵌了`sync.WaitGroup`用来管理多个`groutine`  
`cancel` 即`context`中的`cancel`,用来停止`groutine`  
`errOnce` 从名字可知道,只捕获第一个异常,`err`即为保存该异常的地方.

#### 2. WithContext

```go
// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}
```
`WithContext`很简单,获取`context`的`cancel`,并赋值给新创建的`Group`

#### 3. Wait

```go
// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
```

`Wait`也很简单,利用`sync.WaitGroup`中的`Wait`来实现等待所有的`groutine`执行完  
这里`wg.Wait`执行完后,如果是用`WithContext`创建的`Group`,则会调用`cancel`.  
但是既然所有的`groutine`都跑完了,那么还需要通过`cancel`去通知所有的协程吗?


#### 4. Go

```go
// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
```

`Go`中也很简单,`wg.Add(1)`,然后利用`go`关键字启动一个协程去执行函数,如果发生异常,
将异常赋给`Group.err`,如果`Group`是通过`WithContext`创建的,那么执行`cancel`.
并且只执行一次,也就是只保留最先发生的错误

### bilibili项目里的errgroup

bilibili中的`errgroup`在官方的`errgroup`基础上,并发数量控制,以及如果发生了`panic`,恢复,并增加了堆栈的错误信息
#### 1. 结构体

```go
// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan func() error
	chs        []func() error

	cancel func()
}
```
#### 2. WithContext

```go
// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}
```
和官方的一样
#### 3. do
```go
func (g *Group) do(f func() error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	err = f()
}
```
`do`执行函数,并捕获`panic`异常,`buf = buf[:runtime.Stack(buf, false)]`此代码能够捕获堆栈信息  
其他和官方一样.

#### 4. GOMAXPROCS
```go
// GOMAXPROCS set max goroutine to work.
func (g *Group) GOMAXPROCS(n int) {
	if n <= 0 {
		panic("errgroup: GOMAXPROCS must great than 0")
	}
	g.workerOnce.Do(func() {
		g.ch = make(chan func() error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch {
					g.do(f)
				}
			}()
		}
	})
}
```
设置最大的并发数量.启动n个协程来监听`g.ch`,该通道的类型是函数,一旦接收到函数,立即执行.  
那么初始化时,`g.ch`的容量是否就不一定要等于n?

#### 5. Go
```go
// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- f:
		default:
			g.chs = append(g.chs, f)
		}
		return
	}
	go g.do(f)
}
````
如果调用了`GOMAXPROCS`,那么`g.ch`就不会为`nil`,就将函数传给`g.ch`,如果`g.ch`已经关闭,
就将函数添加进数组`g.ch`.  
如果没有调用`GOMAXPROCS`,那么直接调用`do`来执行函数

#### 6. Wait
```go
// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}

	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
```
如果调用了`GOMAXPROCS`,那么`g.ch`就不会为`nil`,那么就将`g.chs`中所有的函数都传给协程执行.  
并且在所有协程执行完后,关闭`g.ch`