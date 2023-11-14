package limiters

import (
	"context"

	"golang.org/x/time/rate"
)

var _ Limiter = &TokenBucket{}

// Thin wrapper around the golang.org/x/time/rate token bucket rate limiter.
type TokenBucket struct {
	*rate.Limiter
}

func NewTokenBucketLimiter(r int, burst int) Limiter {
	return &TokenBucket{
		Limiter: rate.NewLimiter(rate.Limit(r), burst),
	}
}

func (tb *TokenBucket) Wait(ctx context.Context) error {
	return tb.Limiter.Wait(ctx)
}

func (tb *TokenBucket) WaitN(ctx context.Context, n int) error {
	return tb.Limiter.WaitN(ctx, n)
}

// Reset and shutdown are no-ops for the token bucket limiter.
func (tb *TokenBucket) Reset()    {}
func (tb *TokenBucket) Shutdown() {}
