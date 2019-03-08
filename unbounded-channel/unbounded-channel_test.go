package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnboundedChannel(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for v := range out {
			vi := v.(int)
			fmt.Println("readed", vi)
			if vi != lastVal+1 {
				t.Errorf("expected %d but get %d", lastVal+1, vi)
			}

			lastVal = vi
		}
		wg.Done()
		fmt.Println("finished reading")
	}()

	for i := 0; i < 100; i++ {
		fmt.Println("writting ", i)
		in <- i
	}
	close(in)
	wg.Wait()
	if lastVal != 99 {
		t.Errorf("did not get all data, latest data id %d", lastVal)
	}
}
