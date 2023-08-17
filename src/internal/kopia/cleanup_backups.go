package kopia

import (
	"context"
	"errors"
	"time"

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

// cleanupOrphanedData uses bs and mf to lookup all models/snapshots for backups
// and deletes items that are older than nowFunc() - gcBuffer (cutoff) that are
// not "complete" backups with:
//   - a backup model
//   - an item data snapshot
//   - a details snapshot or details model
//
// We exclude all items younger than the cutoff to add some buffer so that even
// if this is run concurrently with a backup it's not likely to delete models
// just being created. For example, if there was no buffer period and this is
// run when another corso instance has created an item data snapshot but hasn't
// yet created the details snapshot or the backup model would result in this
// instance of corso marking the newly created item data snapshot for deletion
// because it appears orphaned.
//
// For simplicity, we exclude all items younger than the cutoff. It's possible
// to exclude only snapshots and details models since the backup model is the
// last thing persisted for a backup. However, If we selectively exclude things
// then changes to the order of persistence may require changes here too.
//
// The buffer duration should be longer than the difference in creation times
// between the first item data snapshot/details/backup model made during a
// backup operation and the last.
//
// We don't have hard numbers on the time right now, but if the order of
// persistence is (item data snapshot, details snapshot, backup model) it should
// be faster than creating the snapshot itself and probably happens O(minutes)
// or O(hours) instead of O(days). Of course, that assumes a non-adversarial
// setup where things such as machine hiberation, process freezing (i.e. paused
// at the OS level), etc. don't occur.
func cleanupOrphanedData(
	ctx context.Context,
	bs store.Storer,
	mf manifestFinder,
	gcBuffer time.Duration,
	nowFunc func() time.Time,
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

	cutoff := nowFunc().Add(-gcBuffer)

	// Sort all the snapshots as either details snapshots or item data snapshots.
	for _, snap := range snaps {
		// Don't even try to see if this needs garbage collected because it's not
		// old enough and may correspond to an in-progress operation.
		if !cutoff.After(snap.ModTime) {
			continue
		}

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
		// Don't even try to see if this needs garbage collected because it's not
		// old enough and may correspond to an in-progress operation.
		if !cutoff.After(d.ModTime) {
			continue
		}

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
		// Don't even try to see if this needs garbage collected because it's not
		// old enough and may correspond to an in-progress operation.
		if !cutoff.After(bup.ModTime) {
			continue
		}

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

			// Probably safe to continue if the model wasn't found because that means
			// that the possible item data and details for the backup are now
			// orphaned. They'll be deleted since we won't remove them from the delete
			// set.
			//
			// The fact that we exclude all items younger than the cutoff should
			// already exclude items that are from concurrent corso backup operations.
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

	logger.Ctx(ctx).Debugw(
		"garbage collecting orphaned items",
		"num_items", len(toDelete),
		"kopia_ids", maps.Keys(toDelete))

	// Use single atomic batch delete operation to cleanup to keep from making a
	// bunch of manifest content blobs.
	if err := bs.DeleteWithModelStoreIDs(ctx, maps.Keys(toDelete)...); err != nil {
		return clues.Wrap(err, "deleting orphaned data")
	}

	// TODO(ashmrtn): Do some pruning of assist backup models so we don't keep
	// them around forever.

	return nil
}
