package design_model

import (
	"testing"
)

func TestState(t *testing.T) {
	event := &Events{}
	for i := 0; i <6; i ++ {
		event.Alloc()
	}
}
