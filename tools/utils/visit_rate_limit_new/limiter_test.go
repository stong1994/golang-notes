package limiter

import (
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	const (
		id = "dianmi.com"
		limit = 10
		burst = 10
	)
	limiter := NewRateLimiter(limit, burst)
	for i := 0; i < limit; i++ {
		if !limiter.Allow(id) {
			t.Fatal("should allow，index is", i)
		}
	}

	for i := 0; i < limit; i++ {
		if limiter.Allow(id) {
			t.Fatal("should not allow，index is", i)
		}
	}

	time.Sleep(time.Second)

	for i := 0; i < limit; i++ {
		if !limiter.Allow(id) {
			t.Fatal("should allow，index is", i)
		}
	}

	for i := 0; i < limit; i++ {
		if limiter.Allow(id) {
			t.Fatal("should not allow，index is", i)
		}
	}
}

func BenchmarkLimiter(b *testing.B) {
	const (
		id = "dianmi.com"
		limit = 10
		burst = 10
	)
	limiter := NewRateLimiter(limit, burst)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		limiter.Allow(id)
	}
}

func BenchmarkVisitLimiter(b *testing.B) {
	const (
		id = "dianmi.com"
		limit = 10
	)
	limiter := NewVisitLimit(limit, 1000, nodeCount)
	allow := func(id string) bool {
		limiter.UpdateIp(id)
		return limiter.CheckIP(id)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		allow(id)
	}
}