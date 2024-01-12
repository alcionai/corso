package kopia

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/store"
)

// verifyBackups uses bs and mf to lookup all models/snapshots for backups
// and outputs summary information for backups that are not "complete" backups
// with:
//   - a backup model
//   - an item data snapshot
//   - a details snapshot or details model
//
// Output summary information has the form:
//
//	 BackupID: <corso backup ID>
//	   ItemDataSnapshotID: <kopia manifest ID>
//	   DetailsSnapshotID: <kopia manifest ID>
//
//	Items that are missing will have a (missing) note appended to them.
func verifyBackups(
	ctx context.Context,
	bs store.Storer,
	mf multiSnapshotLoader,
) error {
	logger.Ctx(ctx).Infow("scanning for incomplete backups")

	// Get all snapshot manifests.
	snapMetas, err := mf.FindManifests(
		ctx,
		map[string]string{
			manifest.TypeLabelKey: snapshot.ManifestType,
		})
	if err != nil {
		return clues.Wrap(err, "getting snapshot metadata")
	}

	snapIDs := make([]manifest.ID, 0, len(snapMetas))
	for _, m := range snapMetas {
		snapIDs = append(snapIDs, m.ID)
	}

	snaps, err := mf.LoadSnapshots(ctx, snapIDs)
	if err != nil {
		return clues.Wrap(err, "getting snapshots")
	}

	var (
		// deets is a hash set of the ModelStoreID or snapshot IDs for backup
		// details. It contains the IDs for both legacy details stored in the model
		// store and newer details stored as a snapshot because it doesn't matter
		// what the storage format is. We only need to know the ID so we can:
		//   1. check if there's a corresponding backup for them
		deets = map[manifest.ID]struct{}{}
		// dataSnaps is a hash set of the snapshot IDs for item data snapshots.
		dataSnaps = map[manifest.ID]struct{}{}
	)

	// Sort all the snapshots as either details snapshots or item data snapshots.
	for _, snap := range snaps {
		// Filter out checkpoint snapshots as they aren't expected to have a backup
		// associated with them.
		if snap.IncompleteReason == "checkpoint" {
			continue
		}

		k, _ := makeTagKV(TagBackupCategory)
		if _, ok := snap.Tags[k]; ok {
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

	fmt.Println("Incomplete backups:")

	for _, bup := range bups {
		bm := backup.Backup{}

		if err := bs.GetWithModelStoreID(
			ctx,
			model.BackupSchema,
			bup.ModelStoreID,
			&bm); err != nil {
			logger.CtxErr(ctx, err).Infow(
				"backup model not found",
				"search_backup_id", bup.ModelStoreID)

			continue
		}

		ssid := bm.StreamStoreID
		if len(ssid) == 0 {
			ssid = bm.DetailsID
		}

		dataMissing := ""
		deetsMissing := ""

		if _, dataOK := dataSnaps[manifest.ID(bm.SnapshotID)]; !dataOK {
			dataMissing = " (missing)"
		}

		if _, deetsOK := deets[manifest.ID(ssid)]; !deetsOK {
			deetsMissing = " (missing)"
		}

		// Remove from the set so we can mention items that don't seem to have
		// backup models referring to them.
		delete(dataSnaps, manifest.ID(bm.SnapshotID))
		delete(deets, manifest.ID(ssid))

		// Output info about the state of the backup if needed.
		if len(dataMissing) > 0 || len(deetsMissing) > 0 {
			fmt.Printf(
				"\tBackupID: %s\n\t\tItemDataSnapshotID: %s%s\n\t\tDetailsSnapshotID: %s%s\n",
				bm.ID,
				bm.SnapshotID,
				dataMissing,
				ssid,
				deetsMissing)
		}
	}

	fmt.Println("Additional ItemDataSnapshotIDs missing backup models:")
	printIDs(maps.Keys(dataSnaps))

	fmt.Println("Additional DetailsSnapshotIDs missing backup models:")
	printIDs(maps.Keys(deets))

	return nil
}

func printIDs(ids []manifest.ID) {
	for _, id := range ids {
		fmt.Printf("\t%s\n", id)
	}
}
