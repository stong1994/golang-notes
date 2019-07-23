package main

import "testing"

func TestAdd(t *testing.T) {
	n := Add(1, 2)
	if n != 3 {
		t.Errorf("expectd %d, but got %d", 3, n)
	}
}
