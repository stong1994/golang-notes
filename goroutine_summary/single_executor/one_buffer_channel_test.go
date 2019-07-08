package single_executor

import (
	"sync"
	"testing"
)

var counter int

func TestOneBufferChannel(t *testing.T) {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock() // todo 这里要注释掉才能实现 单一任务执行者
		}()
	}
	wg.Wait()
}
