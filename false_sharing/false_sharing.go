package false_sharing

import "sync/atomic"

type MyAtomic interface {
	IncreaseAllEles()
}

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (myatomic *NoPad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

type Pad struct {
	a   uint64
	_p1 [6]uint64
	b   uint64
	_p2 [6]uint64
	c   uint64
	_p3 [6]uint64
}

func (myatomic *Pad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}
