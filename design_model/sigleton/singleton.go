package sigleton

import (
	"sync"
	"sync/atomic"
)

/*
单例模式

参考文章：http://marcio.io/2015/07/singleton-pattern-in-go/
 */

type singleton struct {
}

var instance *singleton

// 一个常见的错误使用
// 在并发情况下，多个协程同时进入（1），造成多次操作。 即并发不安全
func GetInstance_v1() *singleton {
	if instance == nil {
		instance = &singleton{} // （1）
	}
	return instance
}

// 通过加锁进行改进
// 通过枷锁能够保证线程安全，但是实际上只有第一次访问需要加锁，后边访问也要加锁，浪费了很多资源。
var m sync.Mutex
func GetInstance_v2() *singleton {
	m.Lock() // (2)
	defer m.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// 通过两次判断nil来改进 即 Check-Lock-Check Pattern
// 通过判断是否为nil，来判断是否需要加锁，这样能避免每次访问加锁。但是（3）部分不是原子操作
func GetInstance_v3() *singleton {
	if instance == nil { // (3)
		m.Lock()
		defer m.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

// 将上个版本的条件判断换为原子操作
var initialized uint32
func GetInstance_v4() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	m.Lock()
	defer m.Unlock()
	if initialized == 0 { // 因为上边有加锁，所以这里不用原子操作
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1) // 同理，这里是不是也不用原子操作？
	}
	return instance
}

// go中的习惯写法
// 用init()也可以达到同样效果，但是首先init()不能传参，其次，init()函数的执行顺序容易受其他init()函数影响，从而造成隐藏的BUG
var once sync.Once
func GetInstance_v5() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}