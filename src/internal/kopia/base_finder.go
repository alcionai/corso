package kopia

import (
	"context"
	"sort"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup"
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

// TODO(ashmrtn): Move this into some inject package. Here to avoid import
// cycles.
type Reasoner interface {
	Tenant() string
	ProtectedResource() string
	Service() path.ServiceType
	Category() path.CategoryType
	// SubtreePath returns the path prefix for data in existing backups that have
	// parameters (tenant, protected resourced, etc) that match this Reasoner.
	SubtreePath() (path.Path, error)
}

func NewReason(
	tenant, resource string,
	service path.ServiceType,
	category path.CategoryType,
) Reasoner {
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
	p, err := path.ServicePrefix(
		r.Tenant(),
		r.ProtectedResource(),
		r.Service(),
		r.Category())

	return p, clues.Wrap(err, "building path").OrNil()
}

func tagKeys(r Reasoner) []string {
	return []string{
		r.ProtectedResource(),
		serviceCatString(r.Service(), r.Category()),
	}
}

// reasonKey returns the concatenation of the ProtectedResource, Service, and Category.
func reasonKey(r Reasoner) string {
	return r.ProtectedResource() + r.Service().String() + r.Category().String()
}

type BackupEntry struct {
	*backup.Backup
	Reasons []Reasoner
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
	Reasons []Reasoner
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
	LoadSnapshot(ctx context.Context, id manifest.ID) (*snapshot.Manifest, error)
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
	bg inject.GetBackuper
}

func newBaseFinder(
	sm snapshotManager,
	bg inject.GetBackuper,
) (*baseFinder, error) {
	if sm == nil {
		return nil, clues.New("nil snapshotManager")
	}

	if bg == nil {
		return nil, clues.New("nil GetBackuper")
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
	*BackupEntry
	*ManifestEntry
}

// findBasesInSet goes through manifest metadata entries and sees if they're
// incomplete or not. Manifests which don't have an associated backup
// are discarded as incomplete. Manifests are then checked to see if they
// are associated with an assist backup or merge backup.
func (b *baseFinder) findBasesInSet(
	ctx context.Context,
	reason Reasoner,
	metas []*manifest.EntryMetadata,
) (*backupBase, *backupBase, []ManifestEntry, error) {
	// Sort manifests by time so we can go through them sequentially. The code in
	// kopia appears to sort them already, but add sorting here just so we're not
	// reliant on undocumented behavior.
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].ModTime.Before(metas[j].ModTime)
	})

	var (
		mergeBackup      *backupBase
		assistBackup     *backupBase
		assistModel      *BackupEntry
		kopiaAssistSnaps []ManifestEntry
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
		// Check first if this is an assist backup. Criteria for selecting an
		// assist backup are:
		// 1. most recent assist backup for that reason.
		// 2. At most one assist backup per reason.
		// 3. It must be more recent than the merge backup for the reason, if
		// a merge backup exists.
		if assistModel == nil {
			assistModel = b.getAssistBackupModel(ictx, bup, reason)
			if assistModel != nil {
				assistSnap := ManifestEntry{
					Manifest: man,
					Reasons:  []Reasoner{reason},
				}
				kopiaAssistSnaps = append(kopiaAssistSnaps, assistSnap)

				assistBackup = &backupBase{
					BackupEntry:   assistModel,
					ManifestEntry: &assistSnap,
				}

				logger.Ctx(ictx).Infow(
					"found assist backup",
					"assist_backup_id", bup.ID,
					"ssid", ssid)

				continue
			}
		}

		// Found the most recent merge backup. We can stop here.
		//
		// TODO(pandeyabs): Ideally we would do some sanity checks here, e.g.
		// 1. Check for MergeBackup tag. Not safe, since we may have merge
		// backups with older versions of corso which won't have the tag.
		// 2. Check for ErrorCount == 0. We will enforce this after we remove
		// support for BestEffort mode.
		logger.Ctx(ictx).Infow("found complete backup", "base_backup_id", bup.ID)

		mergeSnap := ManifestEntry{
			Manifest: man,
			Reasons:  []Reasoner{reason},
		}

		kopiaAssistSnaps = append(kopiaAssistSnaps, mergeSnap)

		mergeModel := &BackupEntry{
			Backup:  bup,
			Reasons: []Reasoner{reason},
		}

		mergeBackup = &backupBase{
			BackupEntry:   mergeModel,
			ManifestEntry: &mergeSnap,
		}

		break
	}

	logger.Ctx(ctx).Info("no merge or assist backup for reason")

	return mergeBackup, assistBackup, kopiaAssistSnaps, nil
}

// getAssistBackupModel checks if the provided backup is an assist backup.
func (b *baseFinder) getAssistBackupModel(
	ctx context.Context,
	bup *backup.Backup,
	r Reasoner,
) *BackupEntry {
	if bup == nil {
		return nil
	}

	allTags := map[string]string{
		model.BackupTypeTag: model.AssistBackup,
	}

	for k, v := range allTags {
		if bup.Tags[k] != v {
			// This is not an assist backup so we can just exit here.
			logger.Ctx(ctx).Debugw(
				"assist backup does not have expected tag",
				"backup_id", bup.ID,
				"tag", k,
				"expected_value", v,
				"actual_value", bup.Tags[k])

			return nil
		}
	}

	// Check if it has a valid streamstore id and snapshot id.
	if len(bup.StreamStoreID) == 0 || len(bup.SnapshotID) == 0 {
		logger.Ctx(ctx).Infow(
			"nil ssid or snapshot id in assist backup",
			"ssid", bup.StreamStoreID,
			"snapshot_id", bup.SnapshotID)

		return nil
	}

	return &BackupEntry{
		Backup:  bup,
		Reasons: []Reasoner{r},
	}
}

func (b *baseFinder) getBase(
	ctx context.Context,
	r Reasoner,
	tags map[string]string,
) (*backupBase, *backupBase, []ManifestEntry, error) {
	allTags := map[string]string{}

	for _, k := range tagKeys(r) {
		allTags[k] = ""
	}

	maps.Copy(allTags, tags)
	allTags = normalizeTagKVs(allTags)

	metas, err := b.sm.FindManifests(ctx, allTags)
	if err != nil {
		return nil, nil, nil, clues.Wrap(err, "getting snapshots")
	}

	// No snapshots means no backups so we can just exit here.
	if len(metas) == 0 {
		return nil, nil, nil, nil
	}

	return b.findBasesInSet(ctx, r, metas)
}

func (b *baseFinder) FindBases(
	ctx context.Context,
	reasons []Reasoner,
	tags map[string]string,
) BackupBases {
	var (
		// All maps go from ID -> entry. We need to track by ID so we can coalesce
		// the reason for selecting something. Kopia assisted snapshots also use
		// ManifestEntry so we have the reasons for selecting them to aid in
		// debugging.
		mergeBups        = map[model.StableID]BackupEntry{}
		assistBups       = map[model.StableID]BackupEntry{}
		mergeSnaps       = map[manifest.ID]ManifestEntry{}
		kopiaAssistSnaps = map[manifest.ID]ManifestEntry{}
	)

	for _, searchReason := range reasons {
		ictx := clues.Add(
			ctx,
			"search_service", searchReason.Service().String(),
			"search_category", searchReason.Category().String())
		logger.Ctx(ictx).Info("searching for previous manifests")

		mergeBase, assistBase, assistSnaps, err := b.getBase(
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
			mergeBackup := mergeBase.BackupEntry
			mergeSnap := mergeBase.ManifestEntry

			if mergeBackup != nil {
				mBup, ok := mergeBups[mergeBackup.ID]
				if ok {
					mBup.Reasons = append(mBup.Reasons, mergeSnap.Reasons...)
				} else {
					mBup = *mergeBackup
				}

				// Reassign since it's structs not pointers to structs.
				mergeBups[mergeBackup.ID] = mBup
			}

			if mergeSnap != nil {
				bs, ok := mergeSnaps[mergeSnap.ID]
				if ok {
					bs.Reasons = append(bs.Reasons, mergeSnap.Reasons...)
				} else {
					bs = *mergeSnap
				}

				// Reassign since it's structs not pointers to structs.
				mergeSnaps[mergeSnap.ID] = bs
			}
		}

		if assistBase != nil {
			assistBackup := assistBase.BackupEntry
			assistSnap := assistBase.ManifestEntry

			aBup, ok := assistBups[assistBackup.ID]
			if ok {
				aBup.Reasons = append(aBup.Reasons, assistSnap.Reasons...)
			} else {
				aBup = *assistBackup
			}

			// Reassign since it's structs not pointers to structs.
			assistBups[assistBackup.ID] = aBup
		}

		for _, s := range assistSnaps {
			bs, ok := kopiaAssistSnaps[s.ID]
			if ok {
				bs.Reasons = append(bs.Reasons, s.Reasons...)
			} else {
				bs = s
			}

			// Reassign since it's structs not pointers to structs.
			kopiaAssistSnaps[s.ID] = bs
		}
	}

	// TODO(pandeyabs): Fix the terminology used in backupBases to go with
	// new definitions i.e. mergeSnaps instead of mergeBases, etc.
	//
	// Also, we lose the coupling between the backup model and the item data
	// snapshot here. This is fine for now, but we should consider refactoring
	// this to make the 1:1 mapping between the backup model and the snapshot
	// explicit for both merge & assist backups.
	res := &backupBases{
		backups:       maps.Values(mergeBups),
		assistBackups: maps.Values(assistBups),
		mergeBases:    maps.Values(mergeSnaps),
		assistBases:   maps.Values(kopiaAssistSnaps),
	}

	res.fixupAndVerify(ctx)

	return res
}
