package kopia

import (
	"context"
	"errors"
	"sort"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/logger"
)

type BackupBases struct {
	Backups     []BackupEntry
	MergeBases  []ManifestEntry
	AssistBases []ManifestEntry
}

type BackupEntry struct {
	*backup.Backup
	Reasons []Reason
}

type baseFinder struct {
	sm snapshotManager
	bg inject.GetBackuper
}

func NewBaseFinder(
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

// findBasesInSet goes through manifest metadata entries and sees if they're
// incomplete or not. If an entry is incomplete and we don't already have a
// complete or incomplete manifest add it to the set for kopia assisted
// incrementals. If it's complete, fetch the backup model and see if it
// corresponds to a successful backup. If it does, return it as we only need the
// most recent complete backup as the base.
func (b *baseFinder) findBasesInSet(
	ctx context.Context,
	reason Reason,
	metas []*manifest.EntryMetadata,
) (*BackupEntry, *ManifestEntry, []ManifestEntry, error) {
	// Sort manifests by time so we can go through them sequentially. The code in
	// kopia appears to sort them already, but add sorting here just so we're not
	// reliant on undocumented behavior.
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].ModTime.Before(metas[j].ModTime)
	})

	var (
		kopiaAssistSnaps []ManifestEntry
		foundIncomplete  bool
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
			if !foundIncomplete {
				foundIncomplete = true

				kopiaAssistSnaps = append(kopiaAssistSnaps, ManifestEntry{
					Manifest: man,
					Reasons:  []Reason{reason},
				})
			}

			continue
		}

		// This is a complete snapshot so see if we have a backup model for it.
		bup, err := b.getBackupModel(ictx, man)
		if err != nil {
			if errors.Is(err, data.ErrNotFound) {
				continue
			}

			// Safe to continue here as we'll just end up attempting to use an older
			// backup as the base.
			logger.CtxErr(ictx, err).Debug("searching for base backup")

			continue
		}

		ssid := bup.StreamStoreID
		if len(ssid) == 0 {
			ssid = bup.DetailsID
		}

		if len(ssid) == 0 {
			continue
		}

		// If we've made it to this point then we're considering the backup
		// complete as it has both an item data snapshot and a backup details
		// snapshot.
		me := ManifestEntry{
			Manifest: man,
			Reasons:  []Reason{reason},
		}
		kopiaAssistSnaps = append(kopiaAssistSnaps, me)

		return &BackupEntry{
			Backup:  bup,
			Reasons: []Reason{reason},
		}, &me, kopiaAssistSnaps, nil
	}

	logger.Ctx(ctx).Info("no base backups for reason")

	return nil, nil, kopiaAssistSnaps, nil
}

func (b *baseFinder) getBase(
	ctx context.Context,
	reason Reason,
	tags map[string]string,
) (*BackupEntry, *ManifestEntry, []ManifestEntry, error) {
	allTags := map[string]string{}

	for _, k := range reason.TagKeys() {
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

	return b.findBasesInSet(ctx, reason, metas)
}

func (b *baseFinder) FindBases(
	ctx context.Context,
	reasons []Reason,
	tags map[string]string,
) (BackupBases, error) {
	var (
		// All maps go from ID -> entry. We need to track by ID so we can coalesce
		// the reason for selecting something. Kopia assisted snapshots also use
		// ManifestEntry so we have the reasons for selecting them to aid in
		// debugging.
		baseBups         = map[model.StableID]BackupEntry{}
		baseSnaps        = map[manifest.ID]ManifestEntry{}
		kopiaAssistSnaps = map[manifest.ID]ManifestEntry{}
	)

	for _, reason := range reasons {
		ictx := clues.Add(
			ctx,
			"service", reason.Service.String(),
			"category", reason.Category.String())
		logger.Ctx(ictx).Info("searching for previous manifests")

		baseBackup, baseSnap, assistSnaps, err := b.getBase(ictx, reason, tags)
		if err != nil {
			logger.Ctx(ctx).Info("error getting base snapshots for reason. Will fallback to full backup for it")
			continue
		}

		if baseBackup != nil {
			bs, ok := baseBups[baseBackup.ID]
			if ok {
				bs.Reasons = append(bs.Reasons, baseSnap.Reasons...)
			} else {
				bs = *baseBackup
			}

			// Reassign since it's structs not pointers to structs.
			baseBups[baseBackup.ID] = bs
		}

		if baseSnap != nil {
			bs, ok := baseSnaps[baseSnap.ID]
			if ok {
				bs.Reasons = append(bs.Reasons, baseSnap.Reasons...)
			} else {
				bs = *baseSnap
			}

			// Reassign since it's structs not pointers to structs.
			baseSnaps[baseSnap.ID] = bs
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

	return BackupBases{
		Backups:     maps.Values(baseBups),
		MergeBases:  maps.Values(baseSnaps),
		AssistBases: maps.Values(kopiaAssistSnaps),
	}, nil
}
