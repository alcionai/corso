package excludes

import (
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
)

var _ prefixmatcher.StringSetBuilder = &ParentsItems{}

// Items that should be excluded when sourcing data from the base backup.
// Parent Path -> item ID -> {}
type ParentsItems struct {
	pmb prefixmatcher.StringSetBuilder
}

func NewParentsItems() *ParentsItems {
	return &ParentsItems{pmb: prefixmatcher.NewMatcher[map[string]struct{}]()}
}

// copies all items into the parent's bucket.
func (pi *ParentsItems) Add(parent string, items map[string]struct{}) {
	if pi == nil {
		return
	}

	vs, ok := pi.pmb.Get(parent)
	if !ok {
		pi.pmb.Add(parent, items)
		return
	}

	maps.Copy(vs, items)
	pi.pmb.Add(parent, vs)
}

func (pi *ParentsItems) LongestPrefix(parent string) (string, map[string]struct{}, bool) {
	if pi == nil {
		return "", nil, false
	}

	return pi.pmb.LongestPrefix(parent)
}

func (pi *ParentsItems) Empty() bool {
	return pi == nil || pi.pmb.Empty()
}

func (pi *ParentsItems) Get(parent string) (map[string]struct{}, bool) {
	if pi == nil {
		return nil, false
	}

	return pi.pmb.Get(parent)
}

func (pi *ParentsItems) Keys() []string {
	if pi == nil {
		return []string{}
	}

	return pi.pmb.Keys()
}
