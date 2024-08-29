package ratelimit

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity  int //the max capacity of the bucket 
	rate      int //filling rate of the bucket in t/sec
	tokens    int //the actual number of tokens in the bucket
	lastReset time.Time
	mu        sync.Mutex
}


/*
Create a new token Bucket
	- capacity : the max capacity of the bucket
	- the filling rate of the bucket in t/sec
*/
func NewTokenBucket(capacity int, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		rate:      rate,
		tokens:    capacity,
		lastReset: time.Now(),
	}
}


//Take n tokens out of the bucket
func (tb *TokenBucket) Take(n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastReset)
	tb.tokens = int(float64(tb.capacity) * (elapsed.Seconds() / float64(tb.rate)))
	tb.tokens = min(tb.tokens, tb.capacity)

	if tb.tokens >= n {
		tb.tokens -= n
		tb.lastReset = now
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Handle the rate limiting, work like a kind of gate: launcher -> rate-limiter -> webpage
	- tb : your token bucket
	- next the function you have to handle
*/
func RateLimitedHandler(tb *TokenBucket, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.Take(1) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			log.Printf("Rate limit exceeded")
			return
		}
		next(w, r)
	})
}