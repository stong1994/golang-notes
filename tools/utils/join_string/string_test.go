package join_string

import (
	"fmt"
	"testing"
)


func getArr() []string {
	var arr []string
	for i := 0; i <10000 ; i++ {
		arr = append(arr, fmt.Sprintf("%d", i))
	}
	return arr
}

// BenchmarkStringAdd-16    	      20	  61680145 ns/op	204470574 B/op	   10001 allocs/op
func BenchmarkStringAdd(b *testing.B) {
	arr := getArr()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringAdd(arr)
	}
}

// BenchmarkStringJoin-16    	   10000	    226626 ns/op	   40960 B/op	       1 allocs/op
func BenchmarkStringJoin(b *testing.B) {
	arr := getArr()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringJoin(arr)
	}
}

// BenchmarkStringJoinWithAppend-16    	    5000	    355088 ns/op	  204803 B/op	       2 allocs/op
func BenchmarkStringJoinWithAppend(b *testing.B) {
	arr := getArr()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringJoinWithAppend(arr)
	}
}

// BenchmarkSprintfAdd-16    	 1000000	      1442 ns/op	     176 B/op	      11 allocs/op
func BenchmarkSprintfAdd(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SprintfAdd("a","a","a","a","a","a","a","a","a","a")
	}
}

// BenchmarkBufferWith-16    	    5000	    249919 ns/op	  192944 B/op	      12 allocs/op
func BenchmarkBufferWith(b *testing.B) {
	arr := getArr()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BufferWith(arr)
	}
}

// BenchmarkBufferWithGrow-16    	    5000	    240722 ns/op	  112640 B/op	       4 allocs/op
func BenchmarkBufferWithGrow(b *testing.B) {
	arr := getArr()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BufferWithGrow(arr)
	}
}

var s string
var x = []byte{1023:'x'}
var y = []byte{1023:'y'}

func fc()  {
	s = (" " + string(x) + string(y))[1:]
}

func fd()  {
	s = string(x) + string(y)
}

func TestStringGo101(t *testing.T) {
	fmt.Println(testing.AllocsPerRun(1, fc)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd)) // 3
}

/*
BenchmarkFC-8            5000000               339 ns/op            2304 B/op          1 allocs/op
BenchmarkFD-8            2000000               623 ns/op            4096 B/op          3 allocs/op
*/
func BenchmarkFC(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		fc()
	}
}

func BenchmarkFD(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		fd()
	}
}

var x1 = string(x)
var y1 = string(y)

func fc1()  {
	s = (" " + x1 + y1)[1:]
}

func fd1()  {
	s = x1 + y1
}

func TestStringGo101_2(t *testing.T) {
	fmt.Println(testing.AllocsPerRun(1, fc1)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd1)) // 1
}

/*
BenchmarkFC-8            5000000               337 ns/op            2304 B/op          1 allocs/op
BenchmarkFD-8            5000000               329 ns/op            2048 B/op          1 allocs/op
*/