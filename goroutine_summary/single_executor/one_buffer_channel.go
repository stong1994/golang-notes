package single_executor

// 在并发的场景下，希望有一个单一的任务执行者，而不是所有协程都执行。

// 步骤：创建一个缓冲为1的通道，并在初始化的时候填充通道，然后在第一获取的时候获取到通道中的值，这样其他的协程就不能拿到这个值，从而保证了单一的任务执行者。
// 思考：直接用sync.Once?
// Lock try lock
type Lock struct {
	c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock , Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

// 缺点：大量的goroutine抢锁可能会导致CPU无意义的资源浪费。有一个专有名词用来描述这种抢锁的场景：活锁
// 参考文章：https://github.com/chai2010/advanced-go-programming-book/blob/master/ch6-cloud/ch6-02-lock.md
