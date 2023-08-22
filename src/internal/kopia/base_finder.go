package kopia

import (
	"context"
	"sort"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/store"
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

func NewReason(
	tenant, resource string,
	service path.ServiceType,
	category path.CategoryType,
) identity.Reasoner {
	return reason{
		tenant:   tenant,
		resource: resource,
		service:  service,
		category: category,
	}
}

type reason struct {
	// tenant appears here so that when this is moved to an inject package nothing
	// needs changed. However, kopia itself is blind to the fields in the reason
	// struct and relies on helper functions to get the information it needs.
	tenant   string
	resource string
	service  path.ServiceType
	category path.CategoryType
}

func (r reason) Tenant() string {
	return r.tenant
}

func (r reason) ProtectedResource() string {
	return r.resource
}

func (r reason) Service() path.ServiceType {
	return r.service
}

func (r reason) Category() path.CategoryType {
	return r.category
}

func (r reason) SubtreePath() (path.Path, error) {
	p, err := path.BuildPrefix(
		r.Tenant(),
		r.ProtectedResource(),
		r.Service(),
		r.Category())

	return p, clues.Wrap(err, "building path").OrNil()
}

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

type BackupEntry struct {
	*backup.Backup
	Reasons []identity.Reasoner
}

type ManifestEntry struct {
	*snapshot.Manifest
	// Reasons contains the ResourceOwners and Service/Categories that caused this
	// snapshot to be selected as a base. We can't reuse OwnersCats here because
	// it's possible some ResourceOwners will have a subset of the Categories as
	// the reason for selecting a snapshot. For example:
	// 1. backup user1 email,contacts -> B1
	// 2. backup user1 contacts -> B2 (uses B1 as base)
	// 3. backup user1 email,contacts,events (uses B1 for email, B2 for contacts)
	Reasons []identity.Reasoner
}

func (me ManifestEntry) GetTag(key string) (string, bool) {
	k, _ := makeTagKV(key)
	v, ok := me.Tags[k]

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
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return bup, nil
}

type backupBase struct {
	backup   BackupEntry
	manifest ManifestEntry
}

// findBasesInSet goes through manifest metadata entries and sees if they're
// incomplete or not. Manifests which don't have an associated backup
// are discarded as incomplete. Manifests are then checked to see if they
// are associated with an assist backup or merge backup.
func (b *baseFinder) findBasesInSet(
	ctx context.Context,
	reason identity.Reasoner,
	metas []*manifest.EntryMetadata,
) (*backupBase, *backupBase, error) {
	// Sort manifests by time so we can go through them sequentially. The code in
	// kopia appears to sort them already, but add sorting here just so we're not
	// reliant on undocumented behavior.
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].ModTime.Before(metas[j].ModTime)
	})

	var (
		mergeBase  *backupBase
		assistBase *backupBase
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

		ssid := bup.StreamStoreID
		if len(ssid) == 0 {
			ssid = bup.DetailsID
		}

		if len(ssid) == 0 {
			logger.Ctx(ictx).Debugw(
				"empty backup stream store ID",
				"search_backup_id", bup.ID)

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

		if b.isAssistBackupModel(ictx, bup) {
			if assistBase == nil {
				assistModel := BackupEntry{
					Backup:  bup,
					Reasons: []identity.Reasoner{reason},
				}
				assistSnap := ManifestEntry{
					Manifest: man,
					Reasons:  []identity.Reasoner{reason},
				}

				assistBase = &backupBase{
					backup:   assistModel,
					manifest: assistSnap,
				}

				logger.Ctx(ictx).Infow(
					"found assist base",
					"search_backup_id", bup.ID,
					"search_snapshot_id", meta.ID,
					"ssid", ssid)
			}

			// Skip if an assist base has already been selected.
			continue
		}

		logger.Ctx(ictx).Infow("found merge base",
			"search_backup_id", bup.ID,
			"search_snapshot_id", meta.ID,
			"ssid", ssid)

		mergeSnap := ManifestEntry{
			Manifest: man,
			Reasons:  []identity.Reasoner{reason},
		}

		mergeModel := BackupEntry{
			Backup:  bup,
			Reasons: []identity.Reasoner{reason},
		}

		mergeBase = &backupBase{
			backup:   mergeModel,
			manifest: mergeSnap,
		}

		break
	}

	if mergeBase == nil && assistBase == nil {
		logger.Ctx(ctx).Info("no merge or assist base found for reason")
	}

	return mergeBase, assistBase, nil
}

// isAssistBackupModel checks if the provided backup is an assist backup.
func (b *baseFinder) isAssistBackupModel(
	ctx context.Context,
	bup *backup.Backup,
) bool {
	allTags := map[string]string{
		model.BackupTypeTag: model.AssistBackup,
	}

	for k, v := range allTags {
		if bup.Tags[k] != v {
			// This is not an assist backup so we can just exit here.
			logger.Ctx(ctx).Debugw(
				"assist backup model missing tags",
				"backup_id", bup.ID,
				"tag", k,
				"expected_value", v,
				"actual_value", bup.Tags[k])

			return false
		}
	}

	// Check if it has a valid streamstore id and snapshot id.
	if len(bup.StreamStoreID) == 0 || len(bup.SnapshotID) == 0 {
		logger.Ctx(ctx).Infow(
			"nil ssid or snapshot id in assist base",
			"ssid", bup.StreamStoreID,
			"snapshot_id", bup.SnapshotID)

		return false
	}

	return true
}

func (b *baseFinder) getBase(
	ctx context.Context,
	r identity.Reasoner,
	tags map[string]string,
) (*backupBase, *backupBase, error) {
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
		// All maps go from ID -> entry. We need to track by ID so we can coalesce
		// the reason for selecting something. Kopia assisted snapshots also use
		// ManifestEntry so we have the reasons for selecting them to aid in
		// debugging.
		mergeBups   = map[model.StableID]BackupEntry{}
		assistBups  = map[model.StableID]BackupEntry{}
		mergeSnaps  = map[manifest.ID]ManifestEntry{}
		assistSnaps = map[manifest.ID]ManifestEntry{}
	)

	for _, searchReason := range reasons {
		ictx := clues.Add(
			ctx,
			"search_service", searchReason.Service().String(),
			"search_category", searchReason.Category().String())
		logger.Ctx(ictx).Info("searching for previous manifests")

		mergeBase, assistBase, err := b.getBase(
			ictx,
			searchReason,
			tags)
		if err != nil {
			logger.Ctx(ctx).Info(
				"getting base, falling back to full backup for reason",
				"error", err)

			continue
		}

		if mergeBase != nil {
			mergeSnap := mergeBase.manifest
			mergeBackup := mergeBase.backup

			ms, ok := mergeSnaps[mergeSnap.ID]
			if ok {
				ms.Reasons = append(ms.Reasons, mergeSnap.Reasons...)
			} else {
				ms = mergeSnap
			}

			mergeSnaps[mergeSnap.ID] = ms

			mb, ok := mergeBups[mergeBackup.ID]
			if ok {
				mb.Reasons = append(mb.Reasons, mergeSnap.Reasons...)
			} else {
				mb = mergeBackup
			}

			mergeBups[mergeBackup.ID] = mb
		}

		if assistBase != nil {
			assistSnap := assistBase.manifest
			assistBackup := assistBase.backup

			as, ok := assistSnaps[assistSnap.ID]
			if ok {
				as.Reasons = append(as.Reasons, assistSnap.Reasons...)
			} else {
				as = assistSnap
			}

			assistSnaps[assistSnap.ID] = as

			ab, ok := assistBups[assistBackup.ID]
			if ok {
				ab.Reasons = append(ab.Reasons, assistBackup.Reasons...)
			} else {
				ab = assistBackup
			}

			assistBups[assistBackup.ID] = ab
		}
	}

	// TODO(pandeyabs): Fix the terminology used in backupBases to go with
	// new definitions i.e. mergeSnaps instead of mergeBases, etc.
	res := &backupBases{
		backups:       maps.Values(mergeBups),
		assistBackups: maps.Values(assistBups),
		mergeBases:    maps.Values(mergeSnaps),
		assistBases:   maps.Values(assistSnaps),
	}

	res.fixupAndVerify(ctx)

	return res
}
