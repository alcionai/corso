package prefixmatcher

import (
	"strings"
)

type View[T any] interface {
	Get(key string) (T, bool)
	LongestPrefix(key string) (string, T, bool)
	Empty() bool
}

type Matcher[T any] interface {
	// Add adds or updates the item with key to have value value.
	Add(key string, value T)
	View[T]
}

type prefixMatcher[T any] struct {
	data map[string]T
}

func (m *prefixMatcher[T]) Add(key string, value T) {
	m.data[key] = value
}

func (m *prefixMatcher[T]) Get(key string) (T, bool) {
	if m == nil {
		return *new(T), false
	}

	res, ok := m.data[key]

	return res, ok
}

func (m *prefixMatcher[T]) LongestPrefix(key string) (string, T, bool) {
	if m == nil {
		return "", *new(T), false
	}

	var (
		rk    string
		rv    T
		found bool
		// Set to -1 so if there's "" as a prefix ("all match") we still select it.
		longest = -1
	)

	for k, v := range m.data {
		if strings.HasPrefix(key, k) && len(k) > longest {
			found = true
			longest = len(k)
			rk = k
			rv = v
		}
	}

	return rk, rv, found
}

func (m prefixMatcher[T]) Empty() bool {
	return len(m.data) == 0
}

func NewMatcher[T any]() Matcher[T] {
	return &prefixMatcher[T]{data: map[string]T{}}
}
