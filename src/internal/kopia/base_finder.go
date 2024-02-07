package kopia

import (
	"context"
	"sort"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/canario/src/internal/model"
	"github.com/alcionai/canario/src/pkg/backup"
	"github.com/alcionai/canario/src/pkg/backup/identity"
	"github.com/alcionai/canario/src/pkg/logger"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/store"
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

func tagKeys(r identity.Reasoner) []string {
	return []string{
		r.ProtectedResource(),
		serviceCatString(r.Service(), r.Category()),
	}
}

// reasonKey returns the concatenation of the ProtectedResource, Service, and Category.
func reasonKey(r identity.Reasoner) string {
	return r.ProtectedResource() + r.Service().String() + r.Category().String()
}

type BackupBase struct {
	Backup           *backup.Backup
	ItemDataSnapshot *snapshot.Manifest
	// Reasons contains the tenant, protected resource and service/categories that
	// caused this snapshot to be selected as a base. It's possible some
	// (tenant, protected resources) will have a subset of the categories as
	// the reason for selecting a snapshot. For example:
	// 1. backup user1 email,contacts -> B1
	// 2. backup user1 contacts -> B2 (uses B1 as base)
	// 3. backup user1 email,contacts,events (uses B1 for email, B2 for contacts)
	Reasons []identity.Reasoner
}

func (bb BackupBase) GetReasons() []identity.Reasoner {
	return bb.Reasons
}

func (bb BackupBase) GetSnapshotID() manifest.ID {
	return bb.ItemDataSnapshot.ID
}

func (bb BackupBase) GetSnapshotTag(key string) (string, bool) {
	k, _ := makeTagKV(key)
	v, ok := bb.ItemDataSnapshot.Tags[k]

	return v, ok
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

type baseFinder struct {
	sm snapshotManager
	bg store.BackupGetter
}

func newBaseFinder(
	sm snapshotManager,
	bg store.BackupGetter,
) (*baseFinder, error) {
	if sm == nil {
		return nil, clues.New("nil snapshotManager")
	}

	if bg == nil {
		return nil, clues.New("nil BackupGetter")
	}

	return &baseFinder{
		sm: sm,
		bg: bg,
	}, nil
}

func (b *baseFinder) getBackupModel(
	ctx context.Context,
	man *snapshot.Manifest,
) (*backup.Backup, error) {
	k, _ := makeTagKV(TagBackupID)
	bID := man.Tags[k]

	ctx = clues.Add(ctx, "search_backup_id", bID)

	bup, err := b.bg.GetBackup(ctx, model.StableID(bID))
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	return bup, nil
}

// findBasesInSet goes through manifest metadata entries and sees if they're
// incomplete or not. Manifests which don't have an associated backup
// are discarded as incomplete. Manifests are then checked to see if they
// are associated with an assist backup or merge backup.
func (b *baseFinder) findBasesInSet(
	ctx context.Context,
	reason identity.Reasoner,
	metas []*manifest.EntryMetadata,
) (*BackupBase, *BackupBase, error) {
	// Sort manifests by time so we can go through them sequentially. The code in
	// kopia appears to sort them already, but add sorting here just so we're not
	// reliant on undocumented behavior.
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].ModTime.Before(metas[j].ModTime)
	})

	var (
		mergeBase  *BackupBase
		assistBase *BackupBase
	)

	for i := len(metas) - 1; i >= 0; i-- {
		meta := metas[i]
		ictx := clues.Add(ctx, "search_snapshot_id", meta.ID)

		man, err := b.sm.LoadSnapshot(ictx, meta.ID)
		if err != nil {
			// Safe to continue here as we'll just end up attempting to use an older
			// backup as the base.
			logger.CtxErr(ictx, err).Info("attempting to get snapshot")
			continue
		}

		if len(man.IncompleteReason) > 0 {
			// Skip here since this snapshot cannot be considered an assist base.
			logger.Ctx(ictx).Debugw(
				"Incomplete snapshot",
				"incomplete_reason", man.IncompleteReason)

			continue
		}

		// This is a complete snapshot so see if we have a backup model for it.
		bup, err := b.getBackupModel(ictx, man)
		if err != nil {
			// Safe to continue here as we'll just end up attempting to use an older
			// backup as the base.
			logger.CtxErr(ictx, err).Debug("searching for backup model")
			continue
		}

		ictx = clues.Add(ictx, "search_backup_id", bup.ID)

		ssid := bup.StreamStoreID
		if len(ssid) == 0 {
			ssid = bup.DetailsID
		}

		if len(ssid) == 0 {
			logger.Ctx(ictx).Debug("empty backup stream store ID")
			continue
		}

		ictx = clues.Add(ictx, "ssid", ssid)

		if bup.SnapshotID != string(man.ID) {
			logger.Ctx(ictx).Infow(
				"retrieved backup has empty or different snapshot ID from provided manifest",
				"backup_snapshot_id", bup.SnapshotID)

			continue
		}

		// If we've made it to this point then we're considering the backup
		// complete as it has both an item data snapshot and a backup details
		// snapshot.
		//
		// Check first if this is an assist base. Criteria for selecting an
		// assist base are:
		// 1. most recent assist base for the reason.
		// 2. at most one assist base per reason.
		// 3. it must be more recent than the merge backup for the reason, if
		// a merge backup exists.
		switch bup.Type() {
		case model.AssistBackup:
			// Only add an assist base if we haven't already found one.
			if assistBase == nil {
				logger.Ctx(ictx).Info("found assist base")

				assistBase = &BackupBase{
					Backup:           bup,
					ItemDataSnapshot: man,
					Reasons:          []identity.Reasoner{reason},
				}
			}

		case model.MergeBackup:
			logger.Ctx(ictx).Info("found merge base")

			mergeBase = &BackupBase{
				Backup:           bup,
				ItemDataSnapshot: man,
				Reasons:          []identity.Reasoner{reason},
			}

		case model.PreviewBackup:
			// Preview backups are listed here for clarity though they use the same
			// handling as the default case because they can't be used as bases.
			fallthrough
		default:
			logger.Ctx(ictx).Infow(
				"skipping backup with empty or invalid type for incremental backups",
				"backup_type", bup.Type())
		}

		// Need to check here if we found a merge base because adding a break in the
		// case-statement will just leave the case not the for-loop.
		if mergeBase != nil {
			break
		}
	}

	if mergeBase == nil && assistBase == nil {
		logger.Ctx(ctx).Info("no merge or assist base found for reason")
	}

	return mergeBase, assistBase, nil
}

func (b *baseFinder) getBase(
	ctx context.Context,
	r identity.Reasoner,
	tags map[string]string,
) (*BackupBase, *BackupBase, error) {
	allTags := map[string]string{}

	for _, k := range tagKeys(r) {
		allTags[k] = ""
	}

	maps.Copy(allTags, tags)
	allTags = normalizeTagKVs(allTags)

	metas, err := b.sm.FindManifests(ctx, allTags)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting snapshots")
	}

	// No snapshots means no backups so we can just exit here.
	if len(metas) == 0 {
		return nil, nil, nil
	}

	return b.findBasesInSet(ctx, r, metas)
}

func (b *baseFinder) FindBases(
	ctx context.Context,
	reasons []identity.Reasoner,
	tags map[string]string,
) BackupBases {
	var (
		// Backup models and item data snapshot manifests are 1:1 for bases so just
		// track things by the backup ID. We need to track by ID so we can coalesce
		// the reason for selecting something.
		mergeBases  = map[model.StableID]BackupBase{}
		assistBases = map[model.StableID]BackupBase{}
	)

	for _, searchReason := range reasons {
		ictx := clues.Add(
			ctx,
			"search_service", searchReason.Service().String(),
			"search_category", searchReason.Category().String())
		logger.Ctx(ictx).Info("searching for previous manifests")

		mergeBase, assistBase, err := b.getBase(ictx, searchReason, tags)
		if err != nil {
			logger.Ctx(ctx).Info(
				"getting base, falling back to full backup for reason",
				"error", err)

			continue
		}

		if mergeBase != nil {
			mb, ok := mergeBases[mergeBase.Backup.ID]
			if ok {
				mb.Reasons = append(mb.Reasons, mergeBase.Reasons...)
			} else {
				mb = *mergeBase
			}

			mergeBases[mergeBase.Backup.ID] = mb
		}

		if assistBase != nil {
			ab, ok := assistBases[assistBase.Backup.ID]
			if ok {
				ab.Reasons = append(ab.Reasons, assistBase.Reasons...)
			} else {
				ab = *assistBase
			}

			assistBases[assistBase.Backup.ID] = ab
		}
	}

	res := &backupBases{
		mergeBases:  maps.Values(mergeBases),
		assistBases: maps.Values(assistBases),
	}

	res.fixupAndVerify(ctx)

	return res
}
