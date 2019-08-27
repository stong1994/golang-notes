package main

import "runtime"

var m = map[int]struct{}{1: struct{}{}}

func del() {
	delete(m, 1)
	runtime.LockOSThread()
}
