package main

import (
	"fmt"
	"testing"
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
		func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a + i.b
			return out{"adder", r}, nil
		},
		func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a * i.b
			return out{"timer", r}, nil
		},
		func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a - i.b
			return out{"subber", r}, nil
		},
	}
	result, _ := DivideAndConquer(in{2, 3}, evaluators)
	fmt.Println(result)
}
