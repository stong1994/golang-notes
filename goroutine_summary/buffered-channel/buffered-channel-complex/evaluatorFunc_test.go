package main

import (
	"fmt"
	"testing"
)

func TestEvaluatorFunc(t *testing.T) {
	fun1 := EvaluatorFunc(func(inV interface{}) (interface{}, error) {
		return nil, nil
	})

	fmt.Println(Name())      // buffered-channel-complex.TestEvaluatorFunc.func1
	fmt.Println(Evaluate(1)) // <nil> <nil>
}
