package main

import (
	"fmt"
	"testing"
)

func TestEvaluatorFunc(t *testing.T) {
	fun1 := EvaluatorFunc(func(inV interface{}) (interface{}, error) {
		return nil, nil
	})

	fmt.Println(fun1.Name())      // buffered-channel-complex.TestEvaluatorFunc.func1
	fmt.Println(fun1.Evaluate(1)) // <nil> <nil>
}
