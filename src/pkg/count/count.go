package count

import (
	"context"

	"github.com/puzpuzpuz/xsync/v2"
)

// Bus handles threadsafe counting of arbitrarily keyed metrics.
type Bus struct {
	parent *Bus
	stats  *xsync.MapOf[string, *xsync.Counter]
}

func New() *Bus {
	return &Bus{
		stats: xsync.NewMapOf[*xsync.Counter](),
	}
}

// Local generates a bus with a parent link.  Any value added to
// the local instance also updates the parent by the same increment.
// This allows you to maintain an isolated set of counts for a
// bounded context while automatically tallying the global total.
func (b *Bus) Local() *Bus {
	bus := New()
	bus.parent = b

	return bus
}

func (b *Bus) getCounter(k key) *xsync.Counter {
	xc, _ := b.stats.LoadOrStore(string(k), xsync.NewCounter())
	return xc
}

// Inc increases the count by 1.
func (b *Bus) Inc(k key) {
	if b == nil {
		return
	}

	b.Add(k, 1)
}

// Inc increases the count by n.
func (b *Bus) Add(k key, n int64) {
	if b == nil {
		return
	}

	b.getCounter(k).Add(n)

	if b.parent != nil {
		b.parent.Add(k, n)
	}
}

// AdderFor returns a func that adds any value of i
// to the bus using the given key.
func (b *Bus) AdderFor(k key) func(i int64) {
	return func(i int64) {
		b.Add(k, i)
	}
}

// Get returns the local count.
func (b *Bus) Get(k key) int64 {
	if b == nil {
		return -1
	}

	return b.getCounter(k).Value()
}

// Total returns the global count.
func (b *Bus) Total(k key) int64 {
	if b == nil {
		return -1
	}

	if b.parent != nil {
		return b.parent.Total(k)
	}

	return b.Get(k)
}

// Values returns a map of all local values.
// Not a snapshot, and therefore not threadsafe.
func (b *Bus) Values() map[string]int64 {
	if b == nil {
		return map[string]int64{}
	}

	m := make(map[string]int64, b.stats.Size())

	b.stats.Range(func(k string, v *xsync.Counter) bool {
		m[k] = v.Value()
		return true
	})

	return m
}

// TotalValues returns a map of all global values.
// Not a snapshot, and therefore not threadsafe.
func (b *Bus) TotalValues() map[string]int64 {
	if b == nil {
		return map[string]int64{}
	}

	if b.parent != nil {
		return b.parent.TotalValues()
	}

	return b.Values()
}

// ---------------------------------------------------------------------------
// context embedding
// ---------------------------------------------------------------------------

type countKey string

const ctxKey countKey = "corsoCounter"

// Ctx retrieves the count.Bus embedded in the context.
func Ctx(ctx context.Context) *Bus {
	l := ctx.Value(ctxKey)
	if l == nil {
		return New()
	}

	return l.(*Bus)
}

// Embed allows users to embed their own count.Bus within the context.
func Embed(
	ctx context.Context,
	bus *Bus,
) context.Context {
	if bus == nil {
		bus = New()
	}

	return context.WithValue(ctx, ctxKey, bus)
}
