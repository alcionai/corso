package syncd

import (
	"sync"

	"github.com/puzpuzpuz/xsync/v3"
)

// ---------------------------------------------------------------------------
// string -> V
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// K -> V
// ---------------------------------------------------------------------------

// for laxy initialization
var mu sync.Mutex

// MapOf produces a threadsafe map[K]V
type MapOf[K comparable, V any] struct {
	xmo *xsync.MapOf[K, V]
}

// NewMapOf produces a new threadsafe mapOf[K]V
func NewMapOf[K comparable, V any]() MapOf[K, V] {
	return MapOf[K, V]{
		xmo: xsync.NewMapOf[K, V](),
	}
}

// LazyInit ensures the underlying map is populated.
// no-op if already initialized.
func (m *MapOf[K, V]) LazyInit() {
	mu.Lock()
	defer mu.Unlock()

	if m.xmo == nil {
		m.xmo = xsync.NewMapOf[K, V]()
	}
}

func (m MapOf[K, V]) Store(k K, v V) {
	m.xmo.Store(k, v)
}

func (m MapOf[K, V]) Load(k K) (V, bool) {
	return m.xmo.Load(k)
}

func (m MapOf[K, V]) Size() int {
	return m.xmo.Size()
}

func (m MapOf[K, V]) Values() map[K]V {
	vs := map[K]V{}

	if m.xmo == nil {
		return vs
	}

	m.xmo.Range(func(k K, v V) bool {
		vs[k] = v
		return true
	})

	return vs
}
