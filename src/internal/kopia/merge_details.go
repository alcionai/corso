package kopia

import (
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/path"
)

type LocationPrefixMatcher struct {
	m common.PrefixMatcher[*path.Builder]
}

func (m *LocationPrefixMatcher) Add(oldRef path.Path, newLoc *path.Builder) error {
	return m.m.Add(oldRef.String(), newLoc)
}

func (m *LocationPrefixMatcher) LongestPrefix(oldRef string) *path.Builder {
	if m == nil {
		return nil
	}

	k, v, _ := m.m.LongestPrefix(oldRef)
	if k != oldRef {
		// For now we only want to allow exact matches because this is only enabled
		// for Exchange at the moment.
		return nil
	}

	return v
}

func NewLocationPrefixMatcher() *LocationPrefixMatcher {
	return &LocationPrefixMatcher{m: common.NewPrefixMatcher[*path.Builder]()}
}
