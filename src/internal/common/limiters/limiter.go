package limiters

import "context"

type Limiter interface {
	Wait(ctx context.Context) error
	Shutdown()
	Reset()
}
