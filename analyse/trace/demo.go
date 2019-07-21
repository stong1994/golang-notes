package main

// 代码来源：https://gist.github.com/WillSewell/3246161e67f897a530a8120db8917bee
import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func qSort(data []int) {
	if len(data) < 2 {
		return
	}
	pivot := data[0]
	left, right := 1, len(data)-1
	for right >= left {
		if data[left] <= pivot {
			left++
		} else {
			data[right], data[left] = data[left], data[right]
			right--
		}
	}
	// swap pivot into middle
	data[left-1], data[0] = data[0], data[left-1]
	qSort(data[:left-1])
	qSort(data[left:])
}

const thresh = 1e3

func qSortPar(data []int, wg *sync.WaitGroup) {
	if len(data) < 2 {
		// should have bailed to qSort by now but still
		wg.Done()
		return
	}
	pivot := data[0]
	left, right := 1, len(data)-1
	for right >= left {
		if data[left] <= pivot {
			left++
		} else {
			data[right], data[left] = data[left], data[right]
			right--
		}
	}
	// swap pivot into middle
	data[left-1], data[0] = data[0], data[left-1]

	// launch tasks for big subsorts
	if left-1 > thresh {
		wg.Add(1)
		go qSortPar(data[:left-1], wg)
	}
	if len(data)-right > thresh {
		wg.Add(1)
		go qSortPar(data[left:], wg)
	}

	// do small subsorts now
	if left-1 <= thresh {
		qSort(data[:left-1])
	}
	if len(data)-right <= thresh {
		qSort(data[left:])
	}

	// we're done
	wg.Done()
}

func quicksort(data []int) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	qSortPar(data, wg)
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(4)
	data := make([]int, 1e5)

	for i := range data {
		data[i] = int(rand.Uint32() >> 1)
	}
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	trace.Start(f)
	quicksort(data)
	for i := range data[1:] {
		if data[i] > data[i+1] {
			fmt.Println("not sorted at index", i)
			panic("not sorted")
		}
	}
	trace.Stop()
}
