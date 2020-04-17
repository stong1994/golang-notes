package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// RateLimiter r: 1秒内允许访问的最大次数，b：允许同时访问的数量（如没有特殊要求，r和b设置成一致即可）
type RateLimiter struct {
	ids map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b int
}

// NewRateLimiter .
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	i := &RateLimiter{
		ids: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b: b,
	}

	return i
}

// Add creates a new rate limiter and adds it to the ips map,
// using the id as the key
func (i *RateLimiter) add(id string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ids[id] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided ID if it exists.
// Otherwise calls add to add ID to the map
func (i *RateLimiter) getLimiter(id string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ids[id]

	if !exists {
		i.mu.Unlock()
		return i.add(id)
	}

	i.mu.Unlock()

	return limiter
}

func (i *RateLimiter) Allow(id string) bool {
	return i.getLimiter(id).Allow()
}