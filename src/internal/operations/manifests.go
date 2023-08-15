package operations

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

func produceManifestsAndMetadata(
	ctx context.Context,
	bf inject.BaseFinder,
	rp inject.RestoreProducer,
	reasons, fallbackReasons []identity.Reasoner,
	tenantID string,
	getMetadata, dropAssistBases bool,
) (kopia.BackupBases, []data.RestoreCollection, bool, error) {
	bb, meta, useMergeBases, err := getManifestsAndMetadata(
		ctx,
		bf,
		rp,
		reasons,
		fallbackReasons,
		tenantID,
		getMetadata)
	if err != nil {
		return nil, nil, false, clues.Stack(err)
	}

	if !useMergeBases || !getMetadata {
		logger.Ctx(ctx).Debug("full backup requested, dropping merge bases")

		bb.ClearMergeBases()
	}

	if dropAssistBases {
		logger.Ctx(ctx).Debug("no caching requested, dropping assist bases")

		bb.ClearAssistBases()
	}

	return bb, meta, useMergeBases, nil
}

// getManifestsAndMetadata calls kopia to retrieve prior backup manifests,
// metadata collections to supply backup heuristics.
func getManifestsAndMetadata(
	ctx context.Context,
	bf inject.BaseFinder,
	rp inject.RestoreProducer,
	reasons, fallbackReasons []identity.Reasoner,
	tenantID string,
	getMetadata bool,
) (kopia.BackupBases, []data.RestoreCollection, bool, error) {
	var (
		tags          = map[string]string{kopia.TagBackupCategory: ""}
		metadataFiles = graph.AllMetadataFileNames()
		collections   []data.RestoreCollection
	)

	bb := bf.FindBases(ctx, reasons, tags)
	// TODO(ashmrtn): Only fetch these if we haven't already covered all the
	// reasons for this backup.
	fbb := bf.FindBases(ctx, fallbackReasons, tags)

	// one of three cases can occur when retrieving backups across reason migrations:
	// 1. the current reasons don't match any manifests, and we use the fallback to
	// look up the previous reason version.
	// 2. the current reasons only contain an incomplete manifest, and the fallback
	// can find a complete manifest.
	// 3. the current reasons contain all the necessary manifests.
	// Note: This is not relevant for assist backups, since they are newly introduced
	// and they don't exist with fallback reasons.
	bb = bb.MergeBackupBases(ctx, fbb, kopia.BaseKeyServiceCategory)

	if !getMetadata {
		return bb, nil, false, nil
	}

	for _, man := range bb.MergeBases() {
		mctx := clues.Add(ctx, "manifest_id", man.ID)

		// a local fault.Bus intance is used to collect metadata files here.
		// we avoid the global fault.Bus because all failures here are ignorable,
		// and cascading errors up to the operation can cause a conflict that forces
		// the operation into a failure state unnecessarily.
		// TODO(keepers): this is not a pattern we want to
		// spread around.  Need to find more idiomatic handling.
		fb := fault.New(true)

		colls, err := collectMetadata(mctx, rp, man, metadataFiles, tenantID, fb)
		LogFaultErrors(ctx, fb.Errors(), "collecting metadata")

		// TODO(ashmrtn): It should be alright to relax this condition a little. We
		// should be able to just remove the offending manifest and backup from the
		// set of bases. Since we're looking at manifests in this loop, it should be
		// possible to find the backup by either checking the reasons or extracting
		// the backup ID from the manifests tags.
		//
		// Assuming that only the corso metadata is corrupted for the manifest, it
		// should be safe to leave this manifest in the AssistBases set, though we
		// could remove it there too if we want to be conservative. That can be done
		// by finding the manifest ID.
		if err != nil && !errors.Is(err, data.ErrNotFound) {
			// prior metadata isn't guaranteed to exist.
			// if it doesn't, we'll just have to do a
			// full backup for that data.
			return nil, nil, false, err
		}

		collections = append(collections, colls...)
	}

	return bb, collections, true, nil
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
					reason.ProtectedResource(),
					reason.Service(),
					reason.Category(),
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
