package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			var count int
			for i := 0; i < 1e10; i++ {
				count++
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
