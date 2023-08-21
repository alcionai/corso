package kopia

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/store"
)

func cleanupOrphanedData(
	ctx context.Context,
	bs store.Storer,
	mf manifestFinder,
) error {
	// Get all snapshot manifests.
	snaps, err := mf.FindManifests(
		ctx,
		map[string]string{
			manifest.TypeLabelKey: snapshot.ManifestType,
		})
	if err != nil {
		return clues.Wrap(err, "getting snapshots")
	}

	var (
		// deets is a hash set of the ModelStoreID or snapshot IDs for backup
		// details. It contains the IDs for both legacy details stored in the model
		// store and newer details stored as a snapshot because it doesn't matter
		// what the storage format is. We only need to know the ID so we can:
		//   1. check if there's a corresponding backup for them
		//   2. delete the details if they're orphaned
		deets = map[manifest.ID]struct{}{}
		// dataSnaps is a hash set of the snapshot IDs for item data snapshots.
		dataSnaps = map[manifest.ID]struct{}{}
	)

	// TODO(ashmrtn): Exclude all snapshots and details younger than X <duration>.
	// Doing so adds some buffer so that even if this is run concurrently with a
	// backup it's not likely to delete models just being created. For example,
	// running this when another corso instance has created an item data snapshot
	// but hasn't yet created the details snapshot or the backup model would
	// result in this instance of corso marking the newly created item data
	// snapshot for deletion because it appears orphaned.
	//
	// Excluding only snapshots and details models works for now since the backup
	// model is the last thing persisted out of them. If we switch the order of
	// persistence then this will need updated as well.
	//
	// The buffer duration should be longer than the time it would take to do
	// details merging and backup model creation. We don't have hard numbers on
	// that, but it should be faster than creating the snapshot itself and
	// probably happens O(minutes) or O(hours) instead of O(days). Of course, that
	// assumes a non-adversarial setup where things such as machine hiberation,
	// process freezing (i.e. paused at the OS level), etc. don't occur.

	// Sort all the snapshots as either details snapshots or item data snapshots.
	for _, snap := range snaps {
		k, _ := makeTagKV(TagBackupCategory)
		if _, ok := snap.Labels[k]; ok {
			dataSnaps[snap.ID] = struct{}{}
			continue
		}

		deets[snap.ID] = struct{}{}
	}

	// Get all legacy backup details models. The initial version of backup delete
	// didn't seem to delete them so they may also be orphaned if the repo is old
	// enough.
	deetsModels, err := bs.GetIDsForType(ctx, model.BackupDetailsSchema, nil)
	if err != nil {
		return clues.Wrap(err, "getting legacy backup details")
	}

	for _, d := range deetsModels {
		deets[d.ModelStoreID] = struct{}{}
	}

	// Get all backup models.
	bups, err := bs.GetIDsForType(ctx, model.BackupSchema, nil)
	if err != nil {
		return clues.Wrap(err, "getting all backup models")
	}

	toDelete := maps.Clone(deets)
	maps.Copy(toDelete, dataSnaps)

	for _, bup := range bups {
		toDelete[manifest.ID(bup.ModelStoreID)] = struct{}{}

		bm := backup.Backup{}

		if err := bs.GetWithModelStoreID(
			ctx,
			model.BackupSchema,
			bup.ModelStoreID,
			&bm,
		); err != nil {
			if !errors.Is(err, data.ErrNotFound) {
				return clues.Wrap(err, "getting backup model").
					With("search_backup_id", bup.ID)
			}

			// TODO(ashmrtn): This actually needs revised, see above TODO. Leaving it
			// here for the moment to get the basic logic in.
			//
			// Safe to continue if the model wasn't found because that means that the
			// possible item data and details for the backup are now orphaned. They'll
			// be deleted since we won't remove them from the delete set.
			//
			// This isn't expected to really pop up, but it's possible if this
			// function is run concurrently with either a backup delete or another
			// instance of this function.
			logger.Ctx(ctx).Debugw(
				"backup model not found",
				"search_backup_id", bup.ModelStoreID)

			continue
		}

		ssid := bm.StreamStoreID
		if len(ssid) == 0 {
			ssid = bm.DetailsID
		}

		_, dataOK := dataSnaps[manifest.ID(bm.SnapshotID)]
		_, deetsOK := deets[manifest.ID(ssid)]

		// All data is present, we shouldn't garbage collect this backup.
		if deetsOK && dataOK {
			delete(toDelete, bup.ModelStoreID)
			delete(toDelete, manifest.ID(bm.SnapshotID))
			delete(toDelete, manifest.ID(ssid))
		}
	}

	// Use single atomic batch delete operation to cleanup to keep from making a
	// bunch of manifest content blobs.
	if err := bs.DeleteWithModelStoreIDs(ctx, maps.Keys(toDelete)...); err != nil {
		return clues.Wrap(err, "deleting orphaned data")
	}

	// TODO(ashmrtn): Do some pruning of assist backup models so we don't keep
	// them around forever.

	return nil
}
