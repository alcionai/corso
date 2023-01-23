package operations

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

type manifestFetcher interface {
	FetchPrevSnapshotManifests(
		ctx context.Context,
		reasons []kopia.Reason,
		tags map[string]string,
	) ([]*kopia.ManifestEntry, error)
}

type manifestRestorer interface {
	manifestFetcher
	restorer
}

type getDetailsIDer interface {
	GetDetailsIDFromBackupID(
		ctx context.Context,
		backupID model.StableID,
	) (string, *backup.Backup, error)
}

// calls kopia to retrieve prior backup manifests, metadata collections to supply backup heuristics.
func produceManifestsAndMetadata(
	ctx context.Context,
	mr manifestRestorer,
	gdi getDetailsIDer,
	reasons []kopia.Reason,
	tenantID string,
	getMetadata bool,
	errs fault.Adder,
) ([]*kopia.ManifestEntry, []data.Collection, bool, error) {
	var (
		metadataFiles = graph.AllMetadataFileNames()
		collections   []data.Collection
	)

	ms, err := mr.FetchPrevSnapshotManifests(
		ctx,
		reasons,
		map[string]string{kopia.TagBackupCategory: ""})
	if err != nil {
		return nil, nil, false, err
	}

	if !getMetadata {
		return ms, nil, false, nil
	}

	// We only need to check that we have 1:1 reason:base if we're doing an
	// incremental with associated metadata. This ensures that we're only sourcing
	// data from a single Point-In-Time (base) for each incremental backup.
	//
	// TODO(ashmrtn): This may need updating if we start sourcing item backup
	// details from previous snapshots when using kopia-assisted incrementals.
	if err := verifyDistinctBases(ctx, ms, errs); err != nil {
		logger.Ctx(ctx).With("error", err).Infow(
			"base snapshot collision, falling back to full backup",
			clues.Slice(ctx)...)

		return ms, nil, false, nil
	}

	for _, man := range ms {
		if len(man.IncompleteReason) > 0 {
			continue
		}

		mctx := clues.Add(ctx, "manifest_id", man.ID)

		bID, ok := man.GetTag(kopia.TagBackupID)
		if !ok {
			err = clues.New("snapshot manifest missing backup ID").WithMap(clues.Values(mctx))
			return nil, nil, false, err
		}

		mctx = clues.Add(mctx, "manifest_backup_id", man.ID)

		dID, _, err := gdi.GetDetailsIDFromBackupID(mctx, model.StableID(bID))
		if err != nil {
			// if no backup exists for any of the complete manifests, we want
			// to fall back to a complete backup.
			if errors.Is(err, kopia.ErrNotFound) {
				logger.Ctx(ctx).Infow("backup missing, falling back to full backup", clues.Slice(mctx)...)
				return ms, nil, false, nil
			}

			return nil, nil, false, errors.Wrap(err, "retrieving prior backup data")
		}

		mctx = clues.Add(mctx, "manifest_details_id", dID)

		// if no detailsID exists for any of the complete manifests, we want
		// to fall back to a complete backup.  This is a temporary prevention
		// mechanism to keep backups from falling into a perpetually bad state.
		// This makes an assumption that the ID points to a populated set of
		// details; we aren't doing the work to look them up.
		if len(dID) == 0 {
			logger.Ctx(ctx).Infow("backup missing details ID, falling back to full backup", clues.Slice(mctx)...)
			return ms, nil, false, nil
		}

		colls, err := collectMetadata(mctx, mr, man, metadataFiles, tenantID)
		if err != nil && !errors.Is(err, kopia.ErrNotFound) {
			// prior metadata isn't guaranteed to exist.
			// if it doesn't, we'll just have to do a
			// full backup for that data.
			return nil, nil, false, err
		}

		collections = append(collections, colls...)
	}

	return ms, collections, true, err
}

// verifyDistinctBases is a validation checker that ensures, for a given slice
// of manifests, that each manifest's Reason (owner, service, category) is only
// included once.  If a reason is duplicated by any two manifests, an error is
// returned.
func verifyDistinctBases(ctx context.Context, mans []*kopia.ManifestEntry, errs fault.Adder) error {
	var (
		failed  bool
		reasons = map[string]manifest.ID{}
	)

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
				failed = true
				errs.Add(clues.New("manifests have overlapping reasons").
					WithMap(clues.Values(ctx)).
					With("other_manifest_id", b))

				continue
			}

			reasons[reasonKey] = man.ID
		}
	}

	if failed {
		return clues.New("multiple base snapshots qualify").WithMap(clues.Values(ctx))
	}

	return nil
}

// collectMetadata retrieves all metadata files associated with the manifest.
func collectMetadata(
	ctx context.Context,
	r restorer,
	man *kopia.ManifestEntry,
	fileNames []string,
	tenantID string,
) ([]data.Collection, error) {
	paths := []path.Path{}

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
					WithAll("metadata_file", fn, "category", reason.Category)
			}

			paths = append(paths, p)
		}
	}

	dcs, err := r.RestoreMultipleItems(ctx, string(man.ID), paths, nil)
	if err != nil {
		// Restore is best-effort and we want to keep it that way since we want to
		// return as much metadata as we can to reduce the work we'll need to do.
		// Just wrap the error here for better reporting/debugging.
		return dcs, errors.Wrap(err, "collecting prior metadata")
	}

	return dcs, nil
}
