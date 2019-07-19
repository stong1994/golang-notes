package main

import (
	"errors"
	"fmt"
	"time"
)

type Factory func() interface{}

type Pool interface {
	Borrow() interface{}
	Return(interface{})
	BorrowWithTimeout(time time.Duration) (interface{}, error)
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

func (pi *PoolInner) Borrow() interface{} {
	item := <-pi.items
	return item
}

func (pi *PoolInner) Return(in interface{}) {
	pi.items <- in
}

func (pi *PoolInner) BorrowWithTimeout(timeout time.Duration) (interface{}, error) {
	select {
	case item := <-pi.items:
		return item, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("pool timeout after %s", timeout.String())
	}
	return nil, errors.New("should never get here")
}
