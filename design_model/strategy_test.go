package design_model

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	add := &StrategyOperation{Addition{}}
	fmt.Println("add", add.Apply(1, 2))

	mul := &StrategyOperation{Multiplication{}}
	fmt.Println("mul", mul.Apply(2, 3))
}
