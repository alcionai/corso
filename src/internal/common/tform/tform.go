package tform

import (
	"fmt"

	"github.com/alcionai/clues"
)

func FromMapToAny[T any](k string, m map[string]any) (T, error) {
	v, ok := m[k]
	if !ok {
		return *new(T), clues.New("entry not found")
	}

	if v == nil {
		return *new(T), clues.New("nil entry")
	}

	vt, ok := v.(T)
	if !ok {
		return *new(T), clues.New(fmt.Sprintf("unexpected type: %T", v))
	}

	return vt, nil
}
