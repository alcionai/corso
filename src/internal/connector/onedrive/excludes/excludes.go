package excludes

import "golang.org/x/exp/maps"

// Items that should be excluded when sourcing data from the base backup.
// Parent Path -> item ID -> {}
type ParentsItems map[string]map[string]struct{}

func NewParentsItems() ParentsItems {
	return make(ParentsItems)
}

// copies all items into the parent's bucket.
func (pi ParentsItems) Add(parent string, items map[string]struct{}) {
	p, ok := pi[parent]
	if !ok {
		p = map[string]struct{}{}
	}

	maps.Copy(p, items)
	pi[parent] = p
}
