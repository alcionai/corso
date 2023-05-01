package prefixmatcher

import "golang.org/x/exp/maps"

// StringSetReader is a reader designed specifially to contain a set
// of string values (ie: Reader[map[string]struct{}]).
// This is a quality-of-life typecast for the generic Reader.
type StringSetReader interface {
	Reader[map[string]struct{}]
}

// StringSetReader is a builder designed specifially to contain a set
// of string values (ie: Builder[map[string]struct{}]).
// This is a quality-of-life typecast for the generic Builder.
type StringSetBuilder interface {
	Builder[map[string]struct{}]
}

// ---------------------------------------------------------------------------
// Implementation
// ---------------------------------------------------------------------------

var (
	_ StringSetReader  = &StringSetMatcher{}
	_ StringSetBuilder = &StringSetMatchBuilder{}
)

// Items that should be excluded when sourcing data from the base backup.
// Parent Path -> item ID -> {}
type StringSetMatcher struct {
	ssb StringSetBuilder
}

func (m *StringSetMatcher) LongestPrefix(parent string) (string, map[string]struct{}, bool) {
	if m == nil {
		return "", nil, false
	}

	return m.ssb.LongestPrefix(parent)
}

func (m *StringSetMatcher) Empty() bool {
	return m == nil || m.ssb.Empty()
}

func (m *StringSetMatcher) Get(parent string) (map[string]struct{}, bool) {
	if m == nil {
		return nil, false
	}

	return m.ssb.Get(parent)
}

func (m *StringSetMatcher) Keys() []string {
	if m == nil {
		return []string{}
	}

	return m.ssb.Keys()
}

func (m *StringSetMatchBuilder) ToReader() *StringSetMatcher {
	if m == nil {
		return nil
	}

	return m.ssm
}

// Items that should be excluded when sourcing data from the base backup.
// Parent Path -> item ID -> {}
type StringSetMatchBuilder struct {
	ssm *StringSetMatcher
}

func NewStringSetBuilder() *StringSetMatchBuilder {
	return &StringSetMatchBuilder{
		ssm: &StringSetMatcher{
			ssb: NewMatcher[map[string]struct{}](),
		},
	}
}

// copies all items into the key's bucket.
func (m *StringSetMatchBuilder) Add(key string, items map[string]struct{}) {
	if m == nil {
		return
	}

	vs, ok := m.ssm.Get(key)
	if !ok {
		m.ssm.ssb.Add(key, items)
		return
	}

	maps.Copy(vs, items)
	m.ssm.ssb.Add(key, vs)
}

func (m *StringSetMatchBuilder) LongestPrefix(parent string) (string, map[string]struct{}, bool) {
	return m.ssm.LongestPrefix(parent)
}

func (m *StringSetMatchBuilder) Empty() bool {
	return m == nil || m.ssm.Empty()
}

func (m *StringSetMatchBuilder) Get(parent string) (map[string]struct{}, bool) {
	if m == nil {
		return nil, false
	}

	return m.ssm.Get(parent)
}

func (m *StringSetMatchBuilder) Keys() []string {
	if m == nil {
		return []string{}
	}

	return m.ssm.Keys()
}
