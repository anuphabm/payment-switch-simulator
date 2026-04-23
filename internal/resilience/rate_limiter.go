package resilience

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	tokens    int
	capacity  int
	refill    int
	interval  time.Duration
	lastCheck time.Time
}

func NewRateLimiter(capacity, refill int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:    capacity,
		capacity:  capacity,
		refill:    refill,
		interval:  interval,
		lastCheck: time.Now(),
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastCheck)

	// refill tokens
	if elapsed >= r.interval {
		refillCount := int(elapsed / r.interval)
		r.tokens += refillCount * r.refill

		if r.tokens > r.capacity {
			r.tokens = r.capacity
		}

		r.lastCheck = now
	}

	if r.tokens > 0 {
		r.tokens--
		return true
	}

	return false
}
