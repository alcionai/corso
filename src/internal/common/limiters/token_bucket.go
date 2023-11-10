package limiters

import (
	"context"

	"golang.org/x/time/rate"
)

var _ Limiter = &TokenBucket{}

// Wrapper around the golang.org/x/time/rate token bucket rate limiter.
type TokenBucket struct {
	lim *rate.Limiter
}

func NewTokenBucketLimiter(r int, burst int) Limiter {
	lim := rate.NewLimiter(rate.Limit(r), burst)
	return &TokenBucket{lim: lim}
}

func (tb *TokenBucket) Wait(ctx context.Context) error {
	return tb.lim.Wait(ctx)
}

func (tb *TokenBucket) WaitN(ctx context.Context, n int) error {
	return tb.lim.WaitN(ctx, n)
}

func (tb *TokenBucket) Shutdown() {}
