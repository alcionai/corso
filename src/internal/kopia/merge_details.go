package kopia

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/pkg/path"
)

type LocationPrefixMatcher struct {
	m prefixmatcher.Matcher[*path.Builder]
}

func (m *LocationPrefixMatcher) Add(oldRef path.Path, newLoc *path.Builder) error {
	if _, ok := m.m.Get(oldRef.String()); ok {
		return clues.New("RepoRef already in matcher").With("repo_ref", oldRef)
	}

	m.m.Add(oldRef.String(), newLoc)

	return nil
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
	return &LocationPrefixMatcher{m: prefixmatcher.NewMatcher[*path.Builder]()}
}
