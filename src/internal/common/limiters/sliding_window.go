package limiters

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/clues"
)

type token struct{}

type fixedWindow struct {
	count []int
}

var _ Limiter = &slidingWindow{}

type slidingWindow struct {
	// capacity is the maximum number of requests allowed in a sliding window at
	// any given time.
	capacity int
	// windowSize is the total duration of the sliding window. Limiter will allow
	// at most capacity requests in this duration.
	windowSize time.Duration
	// slideInterval controls how frequently the window slides. Smaller interval
	// provides better accuracy at the cost of more frequent sliding & more
	// memory usage.
	slideInterval time.Duration

	// numIntervals is the number of intervals in the window. Calculated as
	// windowSize / slideInterval.
	numIntervals int
	// currentInterval tracks the current slide interval
	currentInterval int

	// Each request acquires a token from the permits channel. If the channel
	// is empty, the request is blocked until a permit is available or if the
	// context is cancelled.
	permits chan token

	// curr and prev are fixed windows of size windowSize. Each window contains
	// a slice of intervals which hold a count of the number of tokens granted
	// during that interval.
	curr fixedWindow
	prev fixedWindow

	// mu synchronizes access to the curr and prev windows
	mu sync.Mutex
	// stopTicker stops the recurring slide ticker
	stopTicker chan struct{}
	closeOnce  sync.Once
}

func NewSlidingWindowLimiter(
	windowSize, slideInterval time.Duration,
	capacity int,
) (Limiter, error) {
	if err := validate(windowSize, slideInterval, capacity); err != nil {
		return nil, err
	}

	ni := int(windowSize / slideInterval)

	s := &slidingWindow{
		windowSize:    windowSize,
		slideInterval: slideInterval,
		capacity:      capacity,
		permits:       make(chan token, capacity),
		numIntervals:  ni,
		prev: fixedWindow{
			count: make([]int, ni),
		},
		curr: fixedWindow{
			count: make([]int, ni),
		},
		currentInterval: -1,
		stopTicker:      make(chan struct{}),
	}

	s.initialize()

	return s, nil
}

// Wait blocks a request until a token is available or the context is cancelled.
// Equivalent to calling WaitN(ctx, 1).
func (s *slidingWindow) Wait(ctx context.Context) error {
	return s.WaitN(ctx, 1)
}

// Wait blocks a request until N tokens are available or the context gets
// cancelled.
func (s *slidingWindow) WaitN(ctx context.Context, N int) error {
	for i := 0; i < N; i++ {
		select {
		case <-ctx.Done():
			return clues.Stack(ctx.Err())
		case <-s.permits:
			s.mu.Lock()
			s.curr.count[s.currentInterval]++
			s.mu.Unlock()
		}
	}

	return nil
}

// Shutdown cleans up the slide goroutine. If shutdown is not called, the slide
// goroutine will continue to run until the program exits.
func (s *slidingWindow) Shutdown() {
	s.closeOnce.Do(func() {
		close(s.stopTicker)
	})
}

// initialize starts the slide goroutine and prefills tokens to full capacity.
func (s *slidingWindow) initialize() {
	// Ok to not hold the mutex here since nothing else is running yet.
	s.nextInterval()

	// Start a goroutine which runs every slideInterval. This goroutine will
	// continue to run until the program exits or until Shutdown is called.
	go func() {
		ticker := time.NewTicker(s.slideInterval)

		for {
			select {
			case <-ticker.C:
				s.slide()
			case <-s.stopTicker:
				ticker.Stop()
				return
			}
		}
	}()

	// Prefill permits to allow tokens to be granted immediately
	for i := 0; i < s.capacity; i++ {
		s.permits <- token{}
	}
}

// nextInterval increments the current interval and slides the fixed
// windows if needed. Should be called with the mutex held.
func (s *slidingWindow) nextInterval() {
	// Increment current interval
	s.currentInterval = (s.currentInterval + 1) % s.numIntervals

	// Slide the fixed windows if windowSize time has elapsed.
	if s.currentInterval == 0 {
		s.prev = s.curr
		s.curr = fixedWindow{
			count: make([]int, s.numIntervals),
		}
	}
}

// slide moves the window forward by one interval. It reclaims tokens from the
// interval that we slid past and adds them back to available permits.
func (s *slidingWindow) slide() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextInterval()

	for i := 0; i < s.prev.count[s.currentInterval]; i++ {
		s.permits <- token{}
	}
}

func validate(
	windowSize, slideInterval time.Duration,
	capacity int,
) error {
	if windowSize <= 0 {
		return clues.New("invalid window size")
	}

	if slideInterval <= 0 {
		return clues.New("invalid slide interval")
	}

	// Allow capacity to be 0 for testing purposes
	if capacity < 0 {
		return clues.New("invalid window capacity")
	}

	if windowSize < slideInterval {
		return clues.New("window too small to fit intervals")
	}

	if windowSize%slideInterval != 0 {
		return clues.New("window not divisible by slide interval")
	}

	return nil
}
