package kopia

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/connector/graph/metadata"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type DetailsMergeInfoer interface {
	// ItemsToMerge returns the number of items that need to be merged.
	ItemsToMerge() int
	// TODO(keepers): remove once sharepoint fully supports metadata files.
	// ItemsToMerge returns the number of items that need to be merged,
	// excluding any item that ends in .meta or .dirmeta
	ItemsToMergeSansMeta() int
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
	) (path.Path, *path.Builder, error)
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

func (m *mergeDetails) ItemsToMergeSansMeta() int {
	if m == nil {
		return 0
	}

	var i int

	for _, rr := range m.repoRefs {
		if !metadata.IsMetadataFile(rr.repoRef) {
			i++
		}
	}

	return i
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
) (path.Path, *path.Builder, error) {
	pr, ok := m.repoRefs[oldRef.ShortRef()]
	if !ok {
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
	m prefixmatcher.Matcher[locRefs]
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
