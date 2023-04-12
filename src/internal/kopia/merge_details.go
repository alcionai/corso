package kopia

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type DetailsMergeInfoer interface {
	// ItemsToMerge returns the number of items that need to be merged.
	ItemsToMerge() int
	// GetNewPathRefs takes the old RepoRef and old LocationRef of an item and
	// returns the new RepoRef, a prefix of the old LocationRef to replace, and
	// the new LocationRefPrefix of the item if the item should be merged. If the
	// item shouldn't be merged nils are returned.
	//
	// If the returned old LocationRef prefix is equal to the old LocationRef then
	// the entire LocationRef should be replaced with the returned value.
	GetNewPathRefs(
		oldRef *path.Builder,
		oldLoc details.LocationIDer,
	) (path.Path, *path.Builder, *path.Builder)
}

type prevRef struct {
	repoRef path.Path
	locRef  *path.Builder
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
	}

	m.repoRefs[oldRef.ShortRef()] = pr

	return nil
}

func (m *mergeDetails) GetNewPathRefs(
	oldRef *path.Builder,
	oldLoc details.LocationIDer,
) (path.Path, *path.Builder, *path.Builder) {
	pr, ok := m.repoRefs[oldRef.ShortRef()]
	if !ok {
		return nil, nil, nil
	}

	// This was a location specified directly by a collection. Say the prefix is
	// the whole oldLoc so other code will replace everything.
	//
	// TODO(ashmrtn): Should be able to remove the nil check later as we'll be
	// able to ensure that old locations actually exist in backup details.
	if oldLoc == nil {
		return pr.repoRef, nil, pr.locRef
	} else if pr.locRef != nil {
		return pr.repoRef, oldLoc.InDetails(), pr.locRef
	}

	// This is a location that we need to do prefix matching on because we didn't
	// see the new location of it in a collection. For example, it's a subfolder
	// whose parent folder was moved.
	prefix, newPrefix := m.locations.longestPrefix(oldLoc.ID())

	return pr.repoRef, prefix, newPrefix
}

func (m *mergeDetails) addLocation(
	oldRef details.LocationIDer,
	newLoc *path.Builder,
) error {
	return m.locations.add(oldRef.ID(), newLoc)
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
	m prefixmatcher.Matcher[locRefs]
}

func (m *locationPrefixMatcher) add(oldRef, newLoc *path.Builder) error {
	key := oldRef.String()

	if _, ok := m.m.Get(key); ok {
		return clues.New("RepoRef already in matcher").With("repo_ref", oldRef)
	}

	m.m.Add(key, locRefs{oldLoc: oldRef, newLoc: newLoc})

	return nil
}

func (m *locationPrefixMatcher) longestPrefix(
	oldRef *path.Builder,
) (*path.Builder, *path.Builder) {
	_, v, _ := m.m.LongestPrefix(oldRef.String())
	return v.oldLoc, v.newLoc
}

func newLocationPrefixMatcher() *locationPrefixMatcher {
	return &locationPrefixMatcher{m: prefixmatcher.NewMatcher[locRefs]()}
}
