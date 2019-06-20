package rand_num

import "testing"

// BenchmarkCryptoRandNum-16    	 1000000	      1039 ns/op	      56 B/op	       4 allocs/op
func BenchmarkCryptoRandNum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i< b.N; i++ {
		CryptoRandNum(100)
	}
}

// BenchmarkMathRandNum-16    	  100000	     17044 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMathRandNum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i< b.N; i++ {
		MathRandNum(100)
	}
}

// BenchmarkRandNum-16    	 1000000	      1052 ns/op	      56 B/op	       4 allocs/op
func BenchmarkRandNum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i< b.N; i++ {
		RandNum(100)
	}
}