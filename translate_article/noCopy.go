package main

import (
	"fmt"
	"sync"
)

func main() {
	InNoCopy()
}

/*
直接运行没有问题
通过命令检测
	go vet -copylocks noCopy.go
	.\noCopy.go:14:8: assignment copies lock value to cL: sync.Mutex

*/
func CopyLocker()  {
	l := sync.Mutex{}
	cL := l
	cL.Lock()
	fmt.Println("检测 mutex 是否真的不能复制")
	cL.Unlock()
}
/*
直接运行没有问题
通过命令检测
	go vet -copylocks noCopy.go
	.\noCopy.go:29:9: assignment copies lock value to cWg: sync.WaitGroup contains sync.noCopy
*/
func CopyWaitGroup()  {
	wg := sync.WaitGroup{}
	cWg := wg
	cWg.Add(1)
	go func() {
		fmt.Println("检测 waitgroup 被复制")
		cWg.Done()
	}()
	cWg.Wait()
}

type noCopy struct{}
func (*noCopy) Lock() {}
func (*noCopy) Unlock() {}

type noCopyS struct {
	noCopy noCopy
	num int
}

// .\noCopy.go:54:7: assignment copies lock value to m: command-line-arguments.noCopyS contains command-line-arguments.noCopy
func InNoCopy()  {
	n := noCopyS{noCopy{}, 2}
	m := n
	m.num++
}

type Integer struct {
	n int
}

func (*Integer) Lock() {
	panic("implement me")
}

func (*Integer) Unlock() {
	panic("implement me")
}

// .\noCopy.go:72:8: assignment copies lock value to cL: command-line-arguments.Integer
func intNoCopy() {
	l := Integer{}
	cL := l
	cL.n++
}