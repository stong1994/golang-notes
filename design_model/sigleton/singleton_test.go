package sigleton

import "testing"

// BenchmarkGetInstance_v3-8       317467203                4.87 ns/op
//func BenchmarkGetInstance_v3(b *testing.B) {
//	for i := 0 ; i < b.N; i++ {
//		GetInstance_v3()
//	}
//}

// BenchmarkGetInstance_v5-8       583165237                1.94 ns/op

func BenchmarkGetInstance_v5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstance_v5()
	}
}
