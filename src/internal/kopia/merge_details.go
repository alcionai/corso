package kopia

import (
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type DetailsMergeInfoer interface {
	// ItemsToMerge returns the number of items that need to be merged.
	ItemsToMerge() int
	// GetNewPathRefs takes the old RepoRef and old LocationRef of an item and
	// returns the new RepoRef, the new location of the item, and the mod time of
	// the item if available. If the item shouldn't be merged nils are returned.
	GetNewPathRefs(
		oldRef *path.Builder,
		modTime time.Time,
		oldLoc details.LocationIDer,
	) (path.Path, *path.Builder, error)
}

type prevRef struct {
	repoRef path.Path
	locRef  *path.Builder
	modTime *time.Time
}

type mergeDetails struct {
	repoRefs  map[string]prevRef
	locations *locationPrefixMatcher
}

func (m *mergeDetails) ItemsToMerge() int {
	if m == nil {
		return 0
	}

	return len(m.repoRefs)
}

func (m *mergeDetails) addRepoRef(
	oldRef *path.Builder,
	modTime *time.Time,
	newRef path.Path,
	newLocRef *path.Builder,
) error {
	if newRef == nil {
		return clues.New("nil RepoRef")
	}

	if _, ok := m.repoRefs[oldRef.ShortRef()]; ok {
		return clues.New("duplicate RepoRef").With("repo_ref", oldRef.String())
	}

	pr := prevRef{
		repoRef: newRef,
		locRef:  newLocRef,
		modTime: modTime,
	}

	m.repoRefs[oldRef.ShortRef()] = pr

	return nil
}

func (m *mergeDetails) GetNewPathRefs(
	oldRef *path.Builder,
	modTime time.Time,
	oldLoc details.LocationIDer,
) (path.Path, *path.Builder, error) {
	pr, ok := m.repoRefs[oldRef.ShortRef()]
	if !ok {
		return nil, nil, nil
	}

	// ModTimes don't match which means we're attempting to merge a different
	// version of the item (i.e. an older version from an assist base). We
	// shouldn't return a match because it could cause us to source out-of-date
	// details for the item.
	if pr.modTime != nil && !pr.modTime.Equal(modTime) {
		return nil, nil, nil
	}

	// This was a location specified directly by a collection.
	if pr.locRef != nil {
		return pr.repoRef, pr.locRef, nil
	} else if oldLoc == nil || oldLoc.ID() == nil || len(oldLoc.ID().Elements()) == 0 {
		return nil, nil, clues.New("empty location key")
	}

	// This is a location that we need to do prefix matching on because we didn't
	// see the new location of it in a collection. For example, it's a subfolder
	// whose parent folder was moved.
	prefixes := m.locations.longestPrefix(oldLoc)
	newLoc := oldLoc.InDetails()

	// Noop if prefix or newPrefix are nil. Them being nil means that the
	// LocationRef hasn't changed.
	newLoc.UpdateParent(prefixes.oldLoc, prefixes.newLoc)

	return pr.repoRef, newLoc, nil
}

func (m *mergeDetails) addLocation(
	oldRef details.LocationIDer,
	newLoc *path.Builder,
) error {
	return m.locations.add(oldRef, newLoc)
}

func newMergeDetails() *mergeDetails {
	return &mergeDetails{
		repoRefs:  map[string]prevRef{},
		locations: newLocationPrefixMatcher(),
	}
}

type locRefs struct {
	oldLoc *path.Builder
	newLoc *path.Builder
}

type locationPrefixMatcher struct {
	m prefixmatcher.Builder[locRefs]
}

func (m *locationPrefixMatcher) add(
	oldRef details.LocationIDer,
	newLoc *path.Builder,
) error {
	key := oldRef.ID().String()

	if _, ok := m.m.Get(key); ok {
		return clues.New("RepoRef already in matcher").With("repo_ref", oldRef)
	}

	m.m.Add(key, locRefs{oldLoc: oldRef.InDetails(), newLoc: newLoc})

	return nil
}

func (m *locationPrefixMatcher) longestPrefix(
	oldRef details.LocationIDer,
) locRefs {
	_, v, _ := m.m.LongestPrefix(oldRef.ID().String())
	return v
}

func newLocationPrefixMatcher() *locationPrefixMatcher {
	return &locationPrefixMatcher{m: prefixmatcher.NewMatcher[locRefs]()}
}
