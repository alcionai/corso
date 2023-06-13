package operations

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

type manifestRestorer interface {
	inject.BaseFinder
	inject.RestoreProducer
}

type getBackuper interface {
	GetBackup(
		ctx context.Context,
		backupID model.StableID,
	) (*backup.Backup, error)
}

// calls kopia to retrieve prior backup manifests, metadata collections to supply backup heuristics.
func produceManifestsAndMetadata(
	ctx context.Context,
	mr manifestRestorer,
	gb getBackuper,
	reasons, fallbackReasons []kopia.Reason,
	tenantID string,
	getMetadata bool,
) ([]kopia.ManifestEntry, []data.RestoreCollection, bool, error) {
	var (
		tags          = map[string]string{kopia.TagBackupCategory: ""}
		metadataFiles = graph.AllMetadataFileNames()
		collections   []data.RestoreCollection
	)

	ms, err := mr.FindBases(ctx, reasons, tags)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "looking up prior snapshots")
	}

	// We only need to check that we have 1:1 reason:base if we're doing an
	// incremental with associated metadata. This ensures that we're only sourcing
	// data from a single Point-In-Time (base) for each incremental backup.
	//
	// TODO(ashmrtn): This may need updating if we start sourcing item backup
	// details from previous snapshots when using kopia-assisted incrementals.
	if err := verifyDistinctBases(ctx, ms); err != nil {
		logger.CtxErr(ctx, err).Info("base snapshot collision, falling back to full backup")
		return ms, nil, false, nil
	}

	fbms, err := mr.FindBases(ctx, fallbackReasons, tags)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "looking up prior snapshots under alternate id")
	}

	// Also check distinct bases for the fallback set.
	if err := verifyDistinctBases(ctx, fbms); err != nil {
		logger.CtxErr(ctx, err).Info("fallback snapshot collision, falling back to full backup")
		return ms, nil, false, nil
	}

	// one of three cases can occur when retrieving backups across reason migrations:
	// 1. the current reasons don't match any manifests, and we use the fallback to
	// look up the previous reason version.
	// 2. the current reasons only contain an incomplete manifest, and the fallback
	// can find a complete manifest.
	// 3. the current reasons contain all the necessary manifests.
	ms = unionManifests(reasons, ms, fbms)

	if !getMetadata {
		return ms, nil, false, nil
	}

	for _, man := range ms {
		if len(man.IncompleteReason) > 0 {
			continue
		}

		mctx := clues.Add(ctx, "manifest_id", man.ID)

		bID, ok := man.GetTag(kopia.TagBackupID)
		if !ok {
			err = clues.New("snapshot manifest missing backup ID").WithClues(ctx)
			return nil, nil, false, err
		}

		mctx = clues.Add(mctx, "manifest_backup_id", bID)

		bup, err := gb.GetBackup(mctx, model.StableID(bID))
		// if no backup exists for any of the complete manifests, we want
		// to fall back to a complete backup.
		if errors.Is(err, data.ErrNotFound) {
			logger.Ctx(mctx).Infow("backup missing, falling back to full backup", clues.In(mctx).Slice()...)
			return ms, nil, false, nil
		}

		if err != nil {
			return nil, nil, false, clues.Wrap(err, "retrieving prior backup data")
		}

		ssid := bup.StreamStoreID
		if len(ssid) == 0 {
			ssid = bup.DetailsID
		}

		mctx = clues.Add(mctx, "manifest_streamstore_id", ssid)

		// if no detailsID exists for any of the complete manifests, we want
		// to fall back to a complete backup.  This is a temporary prevention
		// mechanism to keep backups from falling into a perpetually bad state.
		// This makes an assumption that the ID points to a populated set of
		// details; we aren't doing the work to look them up.
		if len(ssid) == 0 {
			logger.Ctx(ctx).Infow("backup missing streamstore ID, falling back to full backup", clues.In(mctx).Slice()...)
			return ms, nil, false, nil
		}

		// a local fault.Bus intance is used to collect metadata files here.
		// we avoid the global fault.Bus because all failures here are ignorable,
		// and cascading errors up to the operation can cause a conflict that forces
		// the operation into a failure state unnecessarily.
		// TODO(keepers): this is not a pattern we want to
		// spread around.  Need to find more idiomatic handling.
		fb := fault.New(true)

		colls, err := collectMetadata(mctx, mr, man, metadataFiles, tenantID, fb)
		LogFaultErrors(ctx, fb.Errors(), "collecting metadata")

		if err != nil && !errors.Is(err, data.ErrNotFound) {
			// prior metadata isn't guaranteed to exist.
			// if it doesn't, we'll just have to do a
			// full backup for that data.
			return nil, nil, false, err
		}

		collections = append(collections, colls...)
	}

	if err != nil {
		return nil, nil, false, err
	}

	return ms, collections, true, nil
}

// unionManifests reduces the two manifest slices into a single slice.
// Assumes fallback represents a prior manifest version (across some migration
// that disrupts manifest lookup), and that mans contains the current version.
// Also assumes the mans slice will have, at most, one complete and one incomplete
// manifest per service+category tuple.
//
// Selection priority, for each reason, follows these rules:
// 1. If the mans manifest is complete, ignore fallback manifests for that reason.
// 2. If the mans manifest is only incomplete, look for a matching complete manifest in fallbacks.
// 3. If mans has no entry for a reason, look for both complete and incomplete fallbacks.
func unionManifests(
	reasons []kopia.Reason,
	mans []kopia.ManifestEntry,
	fallback []kopia.ManifestEntry,
) []kopia.ManifestEntry {
	if len(fallback) == 0 {
		return mans
	}

	if len(mans) == 0 {
		return fallback
	}

	type manTup struct {
		complete   *kopia.ManifestEntry
		incomplete *kopia.ManifestEntry
	}

	tups := map[string]manTup{}

	for _, r := range reasons {
		// no resource owner in the key.  Assume it's the same owner across all
		// manifests, but that the identifier is different due to migration.
		k := r.Service.String() + r.Category.String()
		tups[k] = manTup{}
	}

	// track the manifests that were collected with the current lookup
	for i := range mans {
		m := &mans[i]

		for _, r := range m.Reasons {
			k := r.Service.String() + r.Category.String()
			t := tups[k]
			// assume mans will have, at most, one complete and one incomplete per key
			if len(m.IncompleteReason) > 0 {
				t.incomplete = m
			} else {
				t.complete = m
			}

			tups[k] = t
		}
	}

	// backfill from the fallback where necessary
	for i := range fallback {
		m := &fallback[i]
		useReasons := []kopia.Reason{}

		for _, r := range m.Reasons {
			k := r.Service.String() + r.Category.String()
			t := tups[k]

			if t.complete != nil {
				// assume fallbacks contains prior manifest versions.
				// we don't want to stack a prior version incomplete onto
				// a current version's complete snapshot.
				continue
			}

			useReasons = append(useReasons, r)

			if len(m.IncompleteReason) > 0 && t.incomplete == nil {
				t.incomplete = m
			} else if len(m.IncompleteReason) == 0 {
				t.complete = m
			}

			tups[k] = t
		}

		if len(m.IncompleteReason) == 0 && len(useReasons) > 0 {
			m.Reasons = useReasons
		}
	}

	// collect the results into a single slice of manifests
	ms := map[string]kopia.ManifestEntry{}

	for _, m := range tups {
		if m.complete != nil {
			ms[string(m.complete.ID)] = *m.complete
		}

		if m.incomplete != nil {
			ms[string(m.incomplete.ID)] = *m.incomplete
		}
	}

	return maps.Values(ms)
}

// verifyDistinctBases is a validation checker that ensures, for a given slice
// of manifests, that each manifest's Reason (owner, service, category) is only
// included once.  If a reason is duplicated by any two manifests, an error is
// returned.
func verifyDistinctBases(ctx context.Context, mans []kopia.ManifestEntry) error {
	reasons := map[string]manifest.ID{}

	for _, man := range mans {
		// Incomplete snapshots are used only for kopia-assisted incrementals. The
		// fact that we need this check here makes it seem like this should live in
		// the kopia code. However, keeping it here allows for better debugging as
		// the kopia code only has access to a path builder which means it cannot
		// remove the resource owner from the error/log output. That is also below
		// the point where we decide if we should do a full backup or an incremental.
		if len(man.IncompleteReason) > 0 {
			continue
		}

		for _, reason := range man.Reasons {
			reasonKey := reason.ResourceOwner + reason.Service.String() + reason.Category.String()

			if b, ok := reasons[reasonKey]; ok {
				return clues.New("manifests have overlapping reasons").
					WithClues(ctx).
					With("other_manifest_id", b)
			}

			reasons[reasonKey] = man.ID
		}
	}

	return nil
}

// collectMetadata retrieves all metadata files associated with the manifest.
func collectMetadata(
	ctx context.Context,
	r inject.RestoreProducer,
	man kopia.ManifestEntry,
	fileNames []string,
	tenantID string,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	paths := []path.RestorePaths{}

	for _, fn := range fileNames {
		for _, reason := range man.Reasons {
			p, err := path.Builder{}.
				Append(fn).
				ToServiceCategoryMetadataPath(
					tenantID,
					reason.ResourceOwner,
					reason.Service,
					reason.Category,
					true)
			if err != nil {
				return nil, clues.
					Wrap(err, "building metadata path").
					With("metadata_file", fn, "category", reason.Category)
			}

			dir, err := p.Dir()
			if err != nil {
				return nil, clues.
					Wrap(err, "building metadata collection path").
					With("metadata_file", fn, "category", reason.Category)
			}

			paths = append(paths, path.RestorePaths{StoragePath: p, RestorePath: dir})
		}
	}

	dcs, err := r.ProduceRestoreCollections(ctx, string(man.ID), paths, nil, errs)
	if err != nil {
		// Restore is best-effort and we want to keep it that way since we want to
		// return as much metadata as we can to reduce the work we'll need to do.
		// Just wrap the error here for better reporting/debugging.
		return dcs, clues.Wrap(err, "collecting prior metadata")
	}

	return dcs, nil
}
