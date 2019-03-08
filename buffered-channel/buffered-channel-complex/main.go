package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type Evaluator interface {
	Evaluate(interface{}) (interface{}, error)
	Name() string
}

type EvaluatorFunc func(interface{}) (interface{}, error)

func (ef EvaluatorFunc) Evaluate(in interface{}) (interface{}, error) {
	return ef(in)
}

func (ef EvaluatorFunc) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(ef).Pointer()).Name()
}

type evaluatorInner struct {
	EvaluatorFunc
	name string
}

func (ei evaluatorInner) Name() string {
	return ei.name
}

func Name(ef EvaluatorFunc, name string) Evaluator {
	return evaluatorInner{
		EvaluatorFunc: ef,
		name:          name,
	}
}

func DivideAndConquer(data interface{}, evaluators []Evaluator, timeout time.Duration) ([]interface{}, []error) {
	gather := make(chan interface{}, len(evaluators))
	errors := make(chan error, len(evaluators))
	for _, v := range evaluators {
		go func(e Evaluator) {
			ch := make(chan interface{}, 1)
			ech := make(chan error, 1) // 无缓冲通道有可能会造成goroutine泄露
			go func() {
				result, err := e.Evaluate(data)
				if err != nil {
					ech <- err
				} else {
					ch <- result
				}
			}()
			// monitor result or err exist or timeout
			select {
			case r := <-ch:
				gather <- r
			case err := <-ech:
				errors <- err
			case <-time.After(timeout):
				errors <- fmt.Errorf("%s timed out after %v on %v", e.Name(), timeout, data)
			}
		}(v)
	}
	out := make([]interface{}, 0, len(evaluators))
	errs := make([]error, 0, len(evaluators))
	// 循环len(evaluators)次以接收到所有的值。否则的话，需要增加一个计数器来监控（如sync.WaitGroup）
	for range evaluators {
		select {
		case r := <-gather:
			out = append(out, r)
		case e := <-errors:
			errs = append(errs, e)
		}
	}
	return out, errs
}
