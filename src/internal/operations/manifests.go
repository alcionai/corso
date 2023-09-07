package operations

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/kopia/inject"
	oinject "github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

func produceManifestsAndMetadata(
	ctx context.Context,
	bf inject.BaseFinder,
	bp oinject.BackupProducer,
	rp inject.RestoreProducer,
	reasons, fallbackReasons []identity.Reasoner,
	tenantID string,
	getMetadata, dropAssistBases bool,
) (kopia.BackupBases, []data.RestoreCollection, bool, error) {
	bb, meta, useMergeBases, err := getManifestsAndMetadata(
		ctx,
		bf,
		bp,
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

		bb.DisableMergeBases()
	}

	if dropAssistBases {
		logger.Ctx(ctx).Debug("no caching requested, dropping assist bases")

		bb.DisableAssistBases()
	}

	return bb, meta, useMergeBases, nil
}

// getManifestsAndMetadata calls kopia to retrieve prior backup manifests,
// metadata collections to supply backup heuristics.
func getManifestsAndMetadata(
	ctx context.Context,
	bf inject.BaseFinder,
	bp oinject.BackupProducer,
	rp inject.RestoreProducer,
	reasons, fallbackReasons []identity.Reasoner,
	tenantID string,
	getMetadata bool,
) (kopia.BackupBases, []data.RestoreCollection, bool, error) {
	var (
		tags        = map[string]string{kopia.TagBackupCategory: ""}
		collections []data.RestoreCollection
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
	bb = bb.MergeBackupBases(
		ctx,
		fbb,
		func(r identity.Reasoner) string {
			return r.Service().String() + r.Category().String()
		})

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

		colls, err := bp.CollectMetadata(mctx, rp, man, fb)
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
