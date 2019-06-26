package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"unsafe"
)

type copyChecker uintptr

type cond struct {
	checker copyChecker
}
func (c *copyChecker) check() {
	fmt.Printf("Before: c: %v, *c: %v, uintptr(*c): %v, uintptr(unsafe.Pointer(c)): %v\n", c, *c, uintptr(*c), uintptr(unsafe.Pointer(c)))
	atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c)))
	fmt.Printf("After: c: %v, *c: %v, uintptr(*c): %v, uintptr(unsafe.Pointer(c)): %v\n", c, *c, uintptr(*c), uintptr(unsafe.Pointer(c)))
}

/** result
Before: c: 0xc0000682d0, *c: 0, 		   uintptr(*c): 0, 			  uintptr(unsafe.Pointer(c)): 824634147536
After: c: 0xc0000682d0,  *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147536
Before: c: 0xc0000682f8, *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147576
After: c: 0xc0000682f8,  *c: 824634147536, uintptr(*c): 824634147536, uintptr(unsafe.Pointer(c)): 824634147576
 */
func TestNocopyCond(t *testing.T) {
	var a cond
	a.checker.check()
	b := a
	b.checker.check()
}
