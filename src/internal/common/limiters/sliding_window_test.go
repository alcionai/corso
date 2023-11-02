package limiters

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/corso/src/internal/tester"
)

func BenchmarkSlidingWindowLimiter(b *testing.B) {
	// 1 second window, 1 millisecond sliding interval, 1000 token capacity (1k per sec)
	limiter := NewLimiter(1*time.Second, 1*time.Millisecond, 1000)
	// If the allowed rate is 1k per sec, 4k goroutines should take 3.xx sec
	numGoroutines := 4000

	ctx, flush := tester.NewContext(b)
	defer flush()

	var wg sync.WaitGroup

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_ = limiter.Wait(ctx)
		}()
	}

	wg.Wait()
	b.StopTimer()

	totalDuration := b.Elapsed()

	fmt.Printf("Total time taken: %v\n", totalDuration)
}
