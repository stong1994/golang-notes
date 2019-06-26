# golang中关于noCopy的使用
[原文地址](https://medium.com/@bronzesword/what-does-nocopy-after-first-use-mean-in-golang-and-how-12396c31de47)

![](https://cdn-images-1.medium.com/max/800/1*Q9_TOm_PD_tjKNsM7Gszyg.png)

当我查看sync包下的源码时，经常会看到这样的文字：`A WaitGroup must not be copied after first use.` 就是纸面上的意思——**不允许复制**

sync包下其他地方如cond、mutex、rwmutex等也有相关说明，在strings包下的builder也有说`Do not copy a non-zero Builder.` 但是实现方式不同。

为什么会不允许复制呢？

***如果一个结构体中有指针类型或者引用类型，如果我们只是对结构体进行浅拷贝，那么我们改变拷贝后的结构体中的指针类型或引用类型就会影响到原结构体。*** 
所以我们需要一个方式来保证结构体的不允许复制。

### 1. 运行时检查
将结构体的指针作为结构体的一个属性，并且在进行操作前进行检查。如`strings.Builder`:
```go
type Builder struct {
    addr *Builder
    buf []byte
}
func (b *Builder) copyCheck() {
    if b.addr == nil {
       b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
    } else if b.addr != b {
       panic("strings: illegal use of non-zero Builder copied by value")
    }
}
func (b *Builder) Write(p []byte) (int, error) {
    b.copyCheck()
    ...
}
// test case
var a strings.Builder
a.Write([]byte("testa"))
var b = a
b.Write([]byte("testb"))   // will panic here
```
当我们声明了`a`，并且调用了`Write()`方法，`a.addr`就保存为它自身的指针。当我们将`a`分配给`b`时，这时进行了浅层拷贝，`b.addr`与`a.addr`是相等的，
但是`b`的指针地址和`b.addr`是不相等的。因为进行`b.Write()`操作时，会`panic`

再举另外一个例子`sync.Cond`:
```go
type Cond struct {
    noCopy  noCopy
    L       Locker
    notify  notifyList
    checker copyChecker
}
type copyChecker uintptr
func (c *copyChecker) check() {
    if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
       !atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
       uintptr(*c) != uintptr(unsafe.Pointer(c)) {
           panic("sync.Cond is copied")
    }
}
func (c *Cond) Wait() {
    c.checker.check()
    ...
}
```
`check()`函数有些复杂，让我们通过一个简单的结构体来测试下
```go
type copyChecker uintptr

type cond struct {
	checker copyChecker
}
func (c *copyChecker) check() {
	fmt.Printf("Before: c: %v, *c: %v, uintptr(*c): %v, uintptr(unsafe.Pointer(c)): %v\n", c, *c, uintptr(*c), uintptr(unsafe.Pointer(c)))
	atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c)))
	fmt.Printf("After: c: %v, *c: %v, uintptr(*c): %v, uintptr(unsafe.Pointer(c)): %v\n", c, *c, uintptr(*c), uintptr(unsafe.Pointer(c)))
}

/** result
Before: c: 0xc0000682d0, *c: 0, 		   uintptr(*c): 0, 			  uintptr(unsafe.Pointer(c)): 824634147536
After: c: 0xc0000682d0,  *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147536
Before: c: 0xc0000682f8, *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147576
After: c: 0xc0000682f8,  *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147576
 */
func TestNocopyCond(t *testing.T) {
	var a cond
	a.checker.check()
	b := a
	b.checker.check()
}
```
当我们声明了一个`a`，它的`checker`字段为`0`，地址为`0xc0000682d0`或`824634147536`。经过`CompareAndSwapUintptr()`后，它的地址没有变。

当我们将`a`赋值给`b`，`a`的`checker`字段也复制给了`b`，可以看到`b.checker`的`*c`,`uintptr(*c)`和`a.checker`是一样的.这样就能理解`sync.Cond`中的`check()`了。

总结：**运行时检查通过自己的指针，在运行时去检查是否复制**

2. 命令检查`Go vet copylocks`
通过`go ver`命令的标签`-copylocks`来检查一个`locker`类型是否被复制了。`locker`类型包含了两个方法：`Lock()`与`Unlock()`

```go
// src/sync/cond.go
type noCopy struct{}
func (*noCopy) Lock() {}
func (*noCopy) Unlock() {}
// sync.Pool
type Pool struct {
   noCopy noCopy  
   ...  
}
// sync.WaitGroup 
type WaitGroup struct {
   noCopy noCopy  
   ...
}
```
`go vet`命令会检查一个`locker`类型，这样能够在运行前就能检测是否有复制操作。所以，如果你想要一个类型不要被复制，那么最简单的方式就是
在包里定义一个`noCopy`结构体，并且将它包含进那个类型，如：
```go
type noCopy struct{}
func (*noCopy) Lock() {}
func (*noCopy) Unlock() {}
type MyType struct {
   noCopy noCopy
   ...
}
```

通过一系列测试，结构体中包含`noCopy`或者结构体实现`locker`都能在命令`go vet -copylocks`中检测到结构体是否被复制。测试文件在`noCopy.go`


**最后。虽然第二种方式比较简便且高效，但是有些时候我们会忘记用命令检查，所以两种方式各有各的优点。**