package mock

import (
	"context"

	"github.com/pkg/errors"
)

type mockBus struct {
	TimesCalled map[string]int
	CalledWith  map[string][]map[string]any
	TimesClosed int
}

func NewBus() *mockBus {
	return &mockBus{
		TimesCalled: map[string]int{},
		CalledWith:  map[string][]map[string]any{},
	}
}

func (b *mockBus) Event(ctx context.Context, key string, data map[string]any) {
	b.TimesCalled[key] = b.TimesCalled[key] + 1

	cw := b.CalledWith[key]
	if len(cw) == 0 {
		cw = []map[string]any{}
	}

	cw = append(cw, data)
	b.CalledWith[key] = cw
}

func (b *mockBus) Close() error {
	b.TimesClosed++

	if b.TimesClosed > 1 {
		return errors.New("multiple closes on mockBus")
	}

	return nil
}
