package kopia

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"

	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// Kopia does not do comparisons properly for empty tags right now so add some
	// placeholder value to them.
	defaultTagValue = "0"
	// Kopia CLI prefixes all user tags with "tag:"[1]. Maintaining this will
	// ensure we don't accidentally take reserved tags and that tags can be
	// displayed with kopia CLI.
	// (permalinks)
	// [1] https://github.com/kopia/kopia/blob/05e729a7858a6e86cb48ba29fb53cb6045efce2b/cli/command_snapshot_create.go#L169
	userTagPrefix = "tag:"
)

type snapshotManager interface {
	FindManifests(
		ctx context.Context,
		tags map[string]string,
	) ([]*manifest.EntryMetadata, error)
	LoadSnapshots(ctx context.Context, ids []manifest.ID) ([]*snapshot.Manifest, error)
}

type ownersCats struct {
	resourceOwners map[string]struct{}
	serviceCats    map[string]struct{}
}

func serviceCatTag(p path.Path) string {
	return p.Service().String() + p.Category().String()
}

func makeTagPair(k string) (string, string) {
	return userTagPrefix + k, defaultTagValue
}

// tagsFromStrings returns a map[string]string with the union of both maps
// passed in. Currently uses placeholder values for each tag because there can be
// multiple instances of resource owners and categories in a single snapshot.
func tagsFromStrings(oc *ownersCats) map[string]string {
	res := make(map[string]string, len(oc.serviceCats)+len(oc.resourceOwners))

	for k := range oc.serviceCats {
		tk, tv := makeTagPair(k)
		res[tk] = tv
	}

	for k := range oc.resourceOwners {
		tk, tv := makeTagPair(k)
		res[tk] = tv
	}

	return res
}
