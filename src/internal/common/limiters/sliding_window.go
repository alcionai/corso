package limiters

import (
	"context"
	"sync"
	"time"
)

type (
	token   struct{}
	Limiter interface {
		Wait(ctx context.Context) error
	}
)

// TODO: Expose interfaces for limiter and window
type window struct {
	// TODO: See if we need to store start time. Without it there is no way
	// to tell if the ticker is lagging behind ( due to contention from consumers or otherwise).
	// Although with our use cases, at max we'd have 10k requests contending with the ticker which
	// should be easily doable in fraction of 1 sec. Although we should benchmark this.
	// start time.Time
	count []int64
}

var _ Limiter = &slidingWindow{}

type slidingWindow struct {
	w               time.Duration
	slidingInterval time.Duration
	capacity        int64
	currentInterval int64
	numIntervals    int64
	permits         chan token
	mu              sync.Mutex
	curr            window
	prev            window
}

// slidingInterval controls degree of movement of the sliding window from left to right
// Smaller slidingInterval means more frequent movement of the sliding window.
// TODO: Introduce an option to control token refresh frequency. Otherwise, if the sliding interval is
// large, it may slow down the token refresh rate. Not implementing this for simplicity, since for our
// use cases we are going to have a sliding interval of 1 sec which is good enough.
func NewLimiter(w time.Duration, slidingInterval time.Duration, capacity int64) Limiter {
	ni := int64(w / slidingInterval)

	sw := &slidingWindow{
		w:               w,
		slidingInterval: slidingInterval,
		capacity:        capacity,
		permits:         make(chan token, capacity),
		numIntervals:    ni,
		prev: window{
			count: make([]int64, ni),
		},
		curr: window{
			count: make([]int64, ni),
		},
		currentInterval: -1,
	}

	// Initialize
	sw.nextInterval()

	// Move the sliding window forward every slidingInterval
	// TODO: fix leaking goroutine
	go sw.run()

	// Prefill permits
	for i := int64(0); i < capacity; i++ {
		sw.permits <- token{}
	}

	return sw
}

// TODO: Implement stopping the ticker
func (s *slidingWindow) run() {
	ticker := time.NewTicker(s.slidingInterval)

	for range ticker.C {
		s.slide()
	}
}

func (s *slidingWindow) slide() {
	// Slide into the next interval
	s.nextInterval()

	// Remove permits from the previous window
	for i := int64(0); i < s.prev.count[s.currentInterval]; i++ {
		select {
		case s.permits <- token{}:
		default:
			// Skip if permits are at capacity
			return
		}
	}
}

// next increments the current interval and resets the current window if needed
func (s *slidingWindow) nextInterval() {
	s.mu.Lock()
	// Increment current interval
	s.currentInterval = (s.currentInterval + 1) % s.numIntervals

	// If it's the first interval, move curr window to prev window and reset curr window.
	if s.currentInterval == 0 {
		s.prev = s.curr
		s.curr = window{
			count: make([]int64, s.numIntervals),
		}
	}

	s.mu.Unlock()
}

// TODO: Implement WaitN
func (s *slidingWindow) Wait(ctx context.Context) error {
	<-s.permits

	// Acquire mutex and increment current interval's count
	s.mu.Lock()
	defer s.mu.Unlock()

	s.curr.count[s.currentInterval]++

	return nil
}
