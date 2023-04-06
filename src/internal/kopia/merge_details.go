package kopia

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/pkg/path"
)

type DetailsMergeInfoer interface {
	// Count returns the number of items that need to be merged.
	Count() int
	// GetNewRepoRef takes the path of the old location of the item and returns
	// its new RepoRef if the item needs merged. If the item doesn't need merged
	// returns nil.
	GetNewRepoRef(oldRef *path.Builder) path.Path
	// GetNewLocation takes the path of the folder containing the item and returns
	// the location of the folder containing the item if it was updated. Otherwise
	// returns nil.
	GetNewLocation(oldRef *path.Builder) *path.Builder
}

type mergeDetails struct {
	repoRefs  map[string]path.Path
	locations *locationPrefixMatcher
}

func (m *mergeDetails) Count() int {
	if m == nil {
		return 0
	}

	return len(m.repoRefs)
}

func (m *mergeDetails) addRepoRef(oldRef *path.Builder, newRef path.Path) error {
	if _, ok := m.repoRefs[oldRef.ShortRef()]; ok {
		return clues.New("duplicate RepoRef").With("repo_ref", oldRef.String())
	}

	m.repoRefs[oldRef.ShortRef()] = newRef

	return nil
}

func (m *mergeDetails) GetNewRepoRef(oldRef *path.Builder) path.Path {
	return m.repoRefs[oldRef.ShortRef()]
}

func (m *mergeDetails) addLocation(oldRef, newLoc *path.Builder) error {
	return m.locations.add(oldRef, newLoc)
}

func (m *mergeDetails) GetNewLocation(oldRef *path.Builder) *path.Builder {
	return m.locations.longestPrefix(oldRef.String())
}

func newMergeDetails() *mergeDetails {
	return &mergeDetails{
		repoRefs:  map[string]path.Path{},
		locations: newLocationPrefixMatcher(),
	}
}

type locationPrefixMatcher struct {
	m prefixmatcher.Matcher[*path.Builder]
}

func (m *locationPrefixMatcher) add(oldRef, newLoc *path.Builder) error {
	if _, ok := m.m.Get(oldRef.String()); ok {
		return clues.New("RepoRef already in matcher").With("repo_ref", oldRef)
	}

	m.m.Add(oldRef.String(), newLoc)

	return nil
}

func (m *locationPrefixMatcher) longestPrefix(oldRef string) *path.Builder {
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

func newLocationPrefixMatcher() *locationPrefixMatcher {
	return &locationPrefixMatcher{m: prefixmatcher.NewMatcher[*path.Builder]()}
}
