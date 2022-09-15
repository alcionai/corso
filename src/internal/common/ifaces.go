package common

import "context"

type Eventer interface {
	Event(context.Context, string, map[string]any)
	Close() error
}
