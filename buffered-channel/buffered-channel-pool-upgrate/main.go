package main

import (
	"errors"
	"fmt"
	"time"
)

type Factory func() interface{}  // 创建一座图书馆，填满书记
type Processor func(interface{}) // 仅供人在读书馆阅读书籍

type Pool interface {
	Run(Processor)
	RunWithTimeout(Processor, time.Duration) error
}

type PoolInner struct {
	items chan interface{}
}

func NewPool(f Factory, count int) Pool {
	pi := &PoolInner{items: make(chan interface{}, count)}
	for i := 0; i < count; i++ {
		pi.items <- f()
	}
	return pi
}

// 阅读，读完书籍自动归还
func (pi *PoolInner) Run(p Processor) {
	item := <-pi.items
	defer func() {
		pi.items <- item
	}()
	p(item)
}

// 限时阅读，读完书籍自动归还
func (pi *PoolInner) RunWithTimeout(p Processor, timeout time.Duration) error {
	select {
	case item := <-pi.items:
		defer func() {
			pi.items <- item
		}()
		p(item)
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("pool timeout after %s", timeout.String())
	}
	return errors.New("should never get here")
}
