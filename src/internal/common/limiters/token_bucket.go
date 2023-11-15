package limiters

import (
	"context"

	"golang.org/x/time/rate"
)

var _ Limiter = &tokenBucket{}

// Thin wrapper around the golang.org/x/time/rate token bucket rate limiter.
type tokenBucket struct {
	*rate.Limiter
}

func NewTokenBucketLimiter(r int, burst int) Limiter {
	return &tokenBucket{
		Limiter: rate.NewLimiter(rate.Limit(r), burst),
	}
}

func (tb *tokenBucket) Wait(ctx context.Context) error {
	return tb.Limiter.Wait(ctx)
}

func (tb *tokenBucket) WaitN(ctx context.Context, n int) error {
	return tb.Limiter.WaitN(ctx, n)
}

// Reset and shutdown are no-ops for the token bucket limiter.
func (tb *tokenBucket) Reset()    {}
func (tb *tokenBucket) Shutdown() {}
