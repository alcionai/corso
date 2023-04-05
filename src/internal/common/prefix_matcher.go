package common

import (
	"strings"

	"github.com/alcionai/clues"
)

type PrefixMatcher[T any] interface {
	Add(key string, value T) error
	Get(key string) (T, bool)
	LongestPrefix(key string) (string, T, bool)
	Map() map[string]T
	Empty() bool
}

type prefixMatcher[T any] struct {
	data map[string]T
}

func (m *prefixMatcher[T]) Add(key string, value T) error {
	if _, ok := m.data[key]; ok {
		return clues.New("entry already exists").With("key", key)
	}

	m.data[key] = value

	return nil
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

func (m prefixMatcher[T]) Map() map[string]T {
	return m.data
}

func (m prefixMatcher[T]) Empty() bool {
	return len(m.data) == 0
}

func NewPrefixMatcher[T any]() PrefixMatcher[T] {
	return &prefixMatcher[T]{data: map[string]T{}}
}
