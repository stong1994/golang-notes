package main

import "testing"

func TestSplit_Foo(t *testing.T) {
	n := foo()
	if n == 5 {
		t.Log("pass")
	} else {
		t.Error("no pass, num is", n)
	}
}

func TestSplit_Foo2(t *testing.T) {
	n := foo2()
	if n == 5 {
		t.Log("pass")
	} else {
		t.Error("no pass, num is", n)
	}
}
