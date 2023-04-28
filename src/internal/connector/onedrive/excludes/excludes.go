package excludes

import (
	"strings"

	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
)

var _ prefixmatcher.MapBuilder = &ParentsItems{}

// Items that should be excluded when sourcing data from the base backup.
// Parent Path -> item ID -> {}
type ParentsItems struct {
	m map[string]map[string]struct{}
}

func NewParentsItems() *ParentsItems {
	return &ParentsItems{m: make(map[string]map[string]struct{})}
}

// copies all items into the parent's bucket.
func (pi *ParentsItems) Add(parent string, items map[string]struct{}) {
	if pi == nil {
		return
	}

	p, ok := pi.m[parent]
	if !ok {
		p = map[string]struct{}{}
	}

	maps.Copy(p, items)
	pi.m[parent] = p
}

func (pi *ParentsItems) LongestPrefix(parent string) (string, map[string]struct{}, bool) {
	if pi == nil {
		return "", nil, false
	}

	var (
		found bool
		rk    string
		rv    map[string]struct{}
	)

	for k, v := range pi.m {
		if strings.HasPrefix(parent, k) && (len(rk) == 0 || len(k) > len(rk)) {
			found = true
			rk = k
			rv = v
		}
	}

	return rk, rv, found
}

func (pi *ParentsItems) Empty() bool {
	return pi == nil || len(pi.m) == 0
}

func (pi *ParentsItems) Get(parent string) (map[string]struct{}, bool) {
	if pi == nil {
		return nil, false
	}

	m, ok := pi.m[parent]

	return m, ok
}

func (pi *ParentsItems) Keys() []string {
	if pi == nil {
		return []string{}
	}

	return maps.Keys(pi.m)
}
