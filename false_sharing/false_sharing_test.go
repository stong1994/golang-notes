package false_sharing

import (
	"sync"
	"testing"
)

/*
>go test -bench=.
BenchmarkNoPad-16       2000000000               0.05 ns/op
BenchmarkPad-16         2000000000               0.02 ns/op
*/
func testAtomicIncrease(myatomic MyAtomic) {
	paraNum := 1000
	addTimes := 1000
	var wg sync.WaitGroup
	wg.Add(paraNum)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				myatomic.IncreaseAllEles()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// BenchmarkNoPad-16    	2000000000	         0.04 ns/op
func BenchmarkNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}

//BenchmarkPad-16    	1000000000	         0.05 ns/op
func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}
