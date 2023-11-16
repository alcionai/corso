package syncd

import (
	"github.com/puzpuzpuz/xsync/v3"
)

// MapTo produces a threadsafe map[string]V
type MapTo[V any] struct {
	xmo *xsync.MapOf[string, V]
}

// NewMapTo produces a new threadsafe mapOf[string]V
func NewMapTo[V any]() MapTo[V] {
	return MapTo[V]{
		xmo: xsync.NewMapOf[string, V](),
	}
}

func (m MapTo[V]) Store(k string, v V) {
	m.xmo.Store(k, v)
}

func (m MapTo[V]) Load(k string) (V, bool) {
	return m.xmo.Load(k)
}

func (m MapTo[V]) Size() int {
	return m.xmo.Size()
}
