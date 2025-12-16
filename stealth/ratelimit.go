package stealth

import "time"

type RateLimiter struct {
	Count     int
	MaxPerDay int
	LastReset time.Time
}

func NewRateLimiter(max int) *RateLimiter {
	return &RateLimiter{
		Count:     0,
		MaxPerDay: max,
		LastReset: time.Now(),
	}
}

func (r *RateLimiter) Allow() bool {
	if time.Since(r.LastReset).Hours() >= 24 {
		r.Count = 0
		r.LastReset = time.Now()
	}

	if r.Count >= r.MaxPerDay {
		return false
	}

	r.Count++
	return true
}
