package main

import (
	"fmt"
	"sync"
)

type FanState string            // 状态
type FanEvent string            // 事件
type FanHandler func() FanState // 处理方法，并返回新的状态

// 有限状态机
type Fan struct {
	mu       sync.Mutex                           // 排他锁
	state    FanState                             // 当前状态
	handlers map[FanState]map[FanEvent]FanHandler // 处理地图集，每一个状态都可以触发有限个事件，执行有限个处理
}

// 获取当前状态
func (f *Fan) getState() FanState {
	return f.state
}

// 设置当前状态
func (f *Fan) setState(newState FanState) {
	f.state = newState
}

// 某状态添加事件处理方法
func (f *Fan) AddHandler(state FanState, event FanEvent, handler FanHandler) *Fan {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FanEvent]FanHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("[警告] 状态(%s)事件(%s)已定义过", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

// 事件处理
func (f *Fan) Call(event FanEvent) FanState {
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.handlers[f.getState()]
	if events == nil {
		return f.getState()
	}
	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn())
		newState := f.getState()
		fmt.Println("状态从 [", oldState, "] 变成 [", newState, "]")
	}
	return f.getState()
}

// 实例化Fan
func NewFan(initState FanState) *Fan {
	return &Fan{
		state:    initState,
		handlers: make(map[FanState]map[FanEvent]FanHandler),
	}
}
