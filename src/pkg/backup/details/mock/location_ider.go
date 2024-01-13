package mock

import "github.com/alcionai/corso/src/pkg/path"

type LocationIDer struct {
	Unique  *path.Builder
	Details *path.Builder
}

func (li LocationIDer) ID() *path.Builder {
	return li.Unique
}

func (li LocationIDer) InDetails() *path.Builder {
	return li.Details
}
