package count

import (
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

func (b *Bus) getCounter(k Key) *xsync.Counter {
	xc, _ := b.stats.LoadOrStore(string(k), xsync.NewCounter())
	return xc
}

// Inc increases the count by 1.
func (b *Bus) Inc(k Key) int64 {
	if b == nil {
		return -1
	}

	return b.Add(k, 1)
}

// Add increases the count by n.
func (b *Bus) Add(k Key, n int64) int64 {
	if b == nil {
		return -1
	}

	b.getCounter(k).Add(n)

	if b.parent != nil {
		b.parent.Add(k, n)
	}

	return b.Get(k)
}

// Get returns the local count.
func (b *Bus) Get(k Key) int64 {
	if b == nil {
		return -1
	}

	return b.getCounter(k).Value()
}

// Total returns the global count.
func (b *Bus) Total(k Key) int64 {
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
// compliance with callbacks and external packages
// ---------------------------------------------------------------------------

// AdderFor returns a func that adds any value of i
// to the bus using the given key.
func (b *Bus) AdderFor(k Key) func(i int64) {
	return func(i int64) {
		b.Add(k, i)
	}
}

type plainAdder struct {
	bus *Bus
}

func (pa plainAdder) Add(k string, n int64) {
	if pa.bus == nil {
		return
	}

	pa.bus.Add(Key(k), n)
}

// PlainAdder provides support to external packages that could take in a count.Bus
// but don't recognize the `Key` type, and would prefer a string type key.
func (b *Bus) PlainAdder() *plainAdder {
	return &plainAdder{b}
}
