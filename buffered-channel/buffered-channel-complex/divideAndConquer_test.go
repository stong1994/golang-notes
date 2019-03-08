package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDivideAndConquer(t *testing.T) {
	type in struct {
		a int
		b int
	}
	type out struct {
		source string
		result int
	}
	evaluators := []Evaluator{
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a + i.b
			return out{"adder", r}, nil
		}),
		Name(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a * i.b
			time.Sleep(50 * time.Millisecond)
			return out{"timer", r}, nil
		}, "name"),
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a - i.b
			return out{"subber", r}, nil
		}),
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a / i.b
			time.Sleep(30 * time.Millisecond)
			return out{"divider", r}, nil
		}),
	}
	result, err := DivideAndConquer(in{2, 3}, evaluators, 10*time.Millisecond)
	fmt.Println(result)
	fmt.Println(err)
}
