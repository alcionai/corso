package kopia

import (
	"context"
	"sort"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/messaging"
	"github.com/alcionai/corso/src/pkg/logger"
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

type ManifestEntry struct {
	*snapshot.Manifest
	// Reason contains the ResourceOwners and Service/Categories that caused this
	// snapshot to be selected as a base. We can't reuse OwnersCats here because
	// it's possible some ResourceOwners will have a subset of the Categories as
	// the reason for selecting a snapshot. For example:
	// 1. backup user1 email,contacts -> B1
	// 2. backup user1 contacts -> B2 (uses B1 as base)
	// 3. backup user1 email,contacts,events (uses B1 for email, B2 for contacts)
	Reasons []messaging.Reason
}

func (me ManifestEntry) GetTag(key string) (string, bool) {
	k, _ := makeTagKV(key)
	v, ok := me.Tags[k]

	return v, ok
}

type snapshotManager interface {
	FindManifests(
		ctx context.Context,
		tags map[string]string,
	) ([]*manifest.EntryMetadata, error)
	LoadSnapshots(ctx context.Context, ids []manifest.ID) ([]*snapshot.Manifest, error)
}

type OwnersCats struct {
	ResourceOwners map[string]struct{}
	ServiceCats    map[string]ServiceCat
}

type ServiceCat struct {
	Service  path.ServiceType
	Category path.CategoryType
}

// MakeServiceCat produces the expected OwnersCats.ServiceCats key from a
// path service and path category, as well as the ServiceCat value.
func MakeServiceCat(s path.ServiceType, c path.CategoryType) (string, ServiceCat) {
	return serviceCatString(s, c), ServiceCat{s, c}
}

// TODO(ashmrtn): Remove in a future PR.
//
//nolint:unused
//lint:ignore U1000 will be removed in future PR.
func serviceCatTag(p path.Path) string {
	return serviceCatString(p.Service(), p.Category())
}

func serviceCatString(s path.ServiceType, c path.CategoryType) string {
	return s.String() + c.String()
}

// MakeTagKV normalizes the provided key to protect it from clobbering
// similarly named tags from non-user input (user inputs are still open
// to collisions amongst eachother).
// Returns the normalized Key plus a default value.  If you're embedding a
// key-only tag, the returned default value msut be used instead of an
// empty string.
func makeTagKV(k string) (string, string) {
	return userTagPrefix + k, defaultTagValue
}

// tagsFromStrings returns a map[string]string with tags for all ownersCats
// passed in. Currently uses placeholder values for each tag because there can
// be multiple instances of resource owners and categories in a single snapshot.
// TODO(ashmrtn): Remove in future PR.
//
//nolint:unused
//lint:ignore U1000 will be removed in future PR.
func tagsFromStrings(oc *OwnersCats) map[string]string {
	if oc == nil {
		return map[string]string{}
	}

	res := make(map[string]string, len(oc.ServiceCats)+len(oc.ResourceOwners))

	for k := range oc.ServiceCats {
		tk, tv := makeTagKV(k)
		res[tk] = tv
	}

	for k := range oc.ResourceOwners {
		tk, tv := makeTagKV(k)
		res[tk] = tv
	}

	return res
}

// getLastIdx searches for manifests contained in both foundMans and metas
// and returns the most recent complete manifest index and the manifest it
// corresponds to. If no complete manifest is in both lists returns nil, -1.
func getLastIdx(
	foundMans map[manifest.ID]*ManifestEntry,
	metas []*manifest.EntryMetadata,
) (*ManifestEntry, int) {
	// Minor optimization: the current code seems to return the entries from
	// earliest timestamp to latest (this is undocumented). Sort in the same
	// fashion so that we don't incur a bunch of swaps.
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].ModTime.Before(metas[j].ModTime)
	})

	// Search newest to oldest.
	for i := len(metas) - 1; i >= 0; i-- {
		m := foundMans[metas[i].ID]
		if m == nil || len(m.IncompleteReason) > 0 {
			continue
		}

		return m, i
	}

	return nil, -1
}

// manifestsSinceLastComplete searches through mans and returns the most recent
// complete manifest (if one exists), maybe the most recent incomplete
// manifest, and a bool denoting if a complete manifest was found. If the newest
// incomplete manifest is more recent than the newest complete manifest then
// adds it to the returned list. Otherwise no incomplete manifest is returned.
// Returns nil if there are no complete or incomplete manifests in mans.
func manifestsSinceLastComplete(
	mans []*snapshot.Manifest,
) ([]*snapshot.Manifest, bool) {
	var (
		res             []*snapshot.Manifest
		foundIncomplete bool
		foundComplete   bool
	)

	// Manifests should maintain the sort order of the original IDs that were used
	// to fetch the data, but just in case sort oldest to newest.
	mans = snapshot.SortByTime(mans, false)

	for i := len(mans) - 1; i >= 0; i-- {
		m := mans[i]

		if len(m.IncompleteReason) > 0 {
			if !foundIncomplete {
				foundIncomplete = true

				res = append(res, m)
			}

			continue
		}

		// Once we find a complete snapshot we're done, even if we haven't
		// found an incomplete one yet.
		res = append(res, m)
		foundComplete = true

		break
	}

	return res, foundComplete
}

// fetchPrevManifests returns the most recent, as-of-yet unfound complete and
// (maybe) incomplete manifests in metas. If the most recent incomplete manifest
// is older than the most recent complete manifest no incomplete manifest is
// returned. If only incomplete manifests exists, returns the most recent one.
// Returns no manifests if an error occurs.
func fetchPrevManifests(
	ctx context.Context,
	sm snapshotManager,
	foundMans map[manifest.ID]*ManifestEntry,
	reason messaging.Reason,
	tags map[string]string,
) ([]*ManifestEntry, error) {
	allTags := map[string]string{}

	for _, k := range reason.TagKeys() {
		allTags[k] = ""
	}

	maps.Copy(allTags, tags)
	allTags = normalizeTagKVs(allTags)

	metas, err := sm.FindManifests(ctx, allTags)
	if err != nil {
		return nil, errors.Wrap(err, "fetching manifest metas by tag")
	}

	if len(metas) == 0 {
		return nil, nil
	}

	man, lastCompleteIdx := getLastIdx(foundMans, metas)

	// We have a complete cached snapshot and it's the most recent. No need
	// to do anything else.
	if lastCompleteIdx == len(metas)-1 {
		man.Reasons = append(man.Reasons, reason)
		return nil, nil
	}

	// TODO(ashmrtn): Remainder of the function can be simplified if we can inject
	// different tags to the snapshot checkpoints than the complete snapshot.

	// Fetch all manifests newer than the oldest complete snapshot. A little
	// wasteful as we may also re-fetch the most recent incomplete manifest, but
	// it reduces the complexity of returning the most recent incomplete manifest
	// if it is newer than the most recent complete manifest.
	ids := make([]manifest.ID, 0, len(metas)-(lastCompleteIdx+1))
	for i := lastCompleteIdx + 1; i < len(metas); i++ {
		ids = append(ids, metas[i].ID)
	}

	mans, err := sm.LoadSnapshots(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "fetching previous manifests")
	}

	found, hasCompleted := manifestsSinceLastComplete(mans)
	res := make([]*ManifestEntry, 0, len(found))

	for _, m := range found {
		res = append(res, &ManifestEntry{
			Manifest: m,
			Reasons:  []messaging.Reason{reason},
		})
	}

	// If we didn't find another complete manifest then we need to mark the
	// previous complete manifest as having this ResourceOwner, Service, Category
	// as the reason as well.
	if !hasCompleted && man != nil {
		man.Reasons = append(man.Reasons, reason)
	}

	return res, nil
}

// fetchPrevSnapshotManifests returns a set of manifests for complete and maybe
// incomplete snapshots for the given (resource owner, service, category)
// tuples. Up to two manifests can be returned per tuple: one complete and one
// incomplete. An incomplete manifest may be returned if it is newer than the
// newest complete manifest for the tuple. Manifests are deduped such that if
// multiple tuples match the same manifest it will only be returned once.
// External callers can access this via wrapper.FetchPrevSnapshotManifests().
// If tags are provided, manifests must include a superset of the k:v pairs
// specified by those tags.  Tags should pass their raw values, and will be
// normalized inside the func using MakeTagKV.
func fetchPrevSnapshotManifests(
	ctx context.Context,
	sm snapshotManager,
	reasons []messaging.Reason,
	tags map[string]string,
) []*ManifestEntry {
	mans := map[manifest.ID]*ManifestEntry{}

	// For each serviceCat/resource owner pair that we will be backing up, see if
	// there's a previous incomplete snapshot and/or a previous complete snapshot
	// we can pass in. Can be expanded to return more than the most recent
	// snapshots, but may require more memory at runtime.
	for _, reason := range reasons {
		found, err := fetchPrevManifests(
			ctx,
			sm,
			mans,
			reason,
			tags,
		)
		if err != nil {
			logger.Ctx(ctx).Warnw(
				"fetching previous snapshot manifests for service/category/resource owner",
				"error",
				err,
				"service",
				reason.Service.String(),
				"category",
				reason.Category.String(),
			)

			// Snapshot can still complete fine, just not as efficient.
			continue
		}

		// If we found more recent snapshots then add them.
		for _, m := range found {
			man := mans[m.ID]
			if man == nil {
				mans[m.ID] = m
				continue
			}

			// If the manifest already exists and it's incomplete then we should
			// merge the reasons for consistency. This will become easier to handle
			// once we update how checkpoint manifests are tagged.
			if len(man.IncompleteReason) == 0 {
				continue
			}

			man.Reasons = append(man.Reasons, m.Reasons...)
		}
	}

	res := make([]*ManifestEntry, 0, len(mans))
	for _, m := range mans {
		res = append(res, m)
	}

	return res
}

func normalizeTagKVs(tags map[string]string) map[string]string {
	t2 := make(map[string]string, len(tags))

	for k, v := range tags {
		mk, mv := makeTagKV(k)

		if len(v) == 0 {
			v = mv
		}

		t2[mk] = v
	}

	return t2
}
