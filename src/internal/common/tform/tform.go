package tform

import (
	"fmt"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
)

func AnyValueToT[T any](k string, m map[string]any) (T, error) {
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

func AnyToT[T any](a any) (T, error) {
	if a == nil {
		return *new(T), clues.New("missing value")
	}

	pt, ok := a.(*T)
	if ok {
		return ptr.Val(pt), nil
	}

	t, ok := a.(T)
	if ok {
		return t, nil
	}

	return *new(T), clues.New(fmt.Sprintf("unexpected type: %T", a))
}
