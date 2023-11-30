package limiters

import "context"

type Limiter interface {
	Wait(ctx context.Context) error
	WaitN(ctx context.Context, n int) error
	Shutdown()
	Reset()
}
