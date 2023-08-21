package kopia

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/store"
)

const (
	serviceCatTagPrefix = "sc-"
	kopiaPathLabel      = "path"
	tenantTag           = "tenant"
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
// yet created the details snapshot or the backup model it would result in this
// instance of corso marking the newly created item data snapshot for deletion
// because it appears orphaned.
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
		dataSnaps = map[manifest.ID]*manifest.EntryMetadata{}
		// toDelete is the set of objects to delete from kopia. It starts out with
		// all items and has ineligible items removed from it.
		toDelete = map[manifest.ID]struct{}{}
	)

	cutoff := nowFunc().Add(-gcBuffer)

	// Sort all the snapshots as either details snapshots or item data snapshots.
	for _, snap := range snaps {
		// Don't even try to see if this needs garbage collected because it's not
		// old enough and may correspond to an in-progress operation.
		if !cutoff.After(snap.ModTime) {
			continue
		}

		toDelete[snap.ID] = struct{}{}

		k, _ := makeTagKV(TagBackupCategory)
		if _, ok := snap.Labels[k]; ok {
			dataSnaps[snap.ID] = snap
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
		toDelete[d.ModelStoreID] = struct{}{}
	}

	// Get all backup models.
	bups, err := bs.GetIDsForType(ctx, model.BackupSchema, nil)
	if err != nil {
		return clues.Wrap(err, "getting all backup models")
	}

	// assistBackups is the set of backups that have a
	//   * a label denoting their an assist backup
	//   * item data snapshot
	//   * details snapshot
	var assistBackups []*backup.Backup

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

		d, dataOK := dataSnaps[manifest.ID(bm.SnapshotID)]
		_, deetsOK := deets[manifest.ID(ssid)]

		// All data is present, we shouldn't garbage collect this backup.
		if deetsOK && dataOK {
			delete(toDelete, bup.ModelStoreID)
			delete(toDelete, manifest.ID(bm.SnapshotID))
			delete(toDelete, manifest.ID(ssid))

			// Add to the assist backup set so that we can attempt to garbage collect
			// older assist backups below.
			if bup.Tags[model.BackupTypeTag] == model.AssistBackup {
				// This is a little messy to have, but can simplify the logic below.
				// The state of tagging in corso isn't all that great right now and we'd
				// really like to consolidate tags and clean them up. For now, we're
				// going to copy tags that are related to Reasons for a backup from the
				// item data snapshot to the backup model. This makes the function
				// checking if assist backups should be garbage collected a bit easier
				// because now they only have to source data from backup models.
				if err := transferTags(d, &bm); err != nil {
					logger.Ctx(ctx).Debugw(
						"transferring legacy tags to backup model",
						"err", err,
						"snapshot_id", d.ID,
						"backup_id", bup.ID)

					// Continuing here means the base won't be eligible for old assist
					// base garbage collection. We could add more logic to eventually
					// delete the base in question but I don't really expect to see
					// failures when transferring tags.
					continue
				}

				assistBackups = append(assistBackups, &bm)
			}
		}
	}

	logger.Ctx(ctx).Infow(
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

var skipKeys = []string{
	TagBackupID,
	TagBackupCategory,
}

func transferTags(snap *manifest.EntryMetadata, bup *backup.Backup) error {
	tenant, err := decodeElement(snap.Labels[kopiaPathLabel])
	if err != nil {
		return clues.Wrap(err, "decoding tenant from label")
	}

	bup.Tags[tenantTag] = tenant

	skipTags := map[string]struct{}{}

	for _, k := range skipKeys {
		key, _ := makeTagKV(k)
		skipTags[key] = struct{}{}
	}

	// Safe to check only this because the old field was deprecated prior to the
	// tagging of assist backups and this function only deals with assist
	// backups.
	roid := bup.ProtectedResourceID

	roidK, _ := makeTagKV(roid)
	skipTags[roidK] = struct{}{}

	// This is hacky, but right now we don't have a good way to get only the
	// Reason tags for something. We can however, find them by searching for all
	// the "normalized" tags and then discarding the ones we know aren't
	// reasons. Unfortunately this won't work if custom tags are added to the
	// backup that we don't know about.
	//
	// Convert them to the newer format that we'd like to have where the
	// service/category tags have the form "sc-<service><category>".
	for tag := range snap.Labels {
		if _, ok := skipTags[tag]; ok || !strings.HasPrefix(tag, userTagPrefix) {
			continue
		}

		bup.Tags[strings.Replace(tag, userTagPrefix, serviceCatTagPrefix, 1)] = "0"
	}

	return nil
}

func collectOldAssistBases(
	ctx context.Context,
	bups []*backup.Backup,
) []manifest.ID {
	// maybeDelete is the set of backups that could be deleted. It starts out as
	// the set of all backups and has ineligible backups removed from it.
	maybeDelete := map[manifest.ID]*backup.Backup{}
	// Figure out which backups have overlapping reasons. A single backup can
	// appear in multiple slices in the map, one for each Reason associated with
	// it.
	bupsByReason := map[string][]*backup.Backup{}

	for _, bup := range bups {
		// Safe to pull from this field since assist backups came after we switched
		// to using ProtectedResourceID.
		roid := bup.ProtectedResourceID

		tenant := bup.Tags[tenantTag]
		if len(tenant) == 0 {
			// We can skip this backup. It won't get garbage collected, but it also
			// won't result in incorrect behavior overall.
			logger.Ctx(ctx).Infow("missing tenant tag in backup", "backup_id", bup.ID)
			continue
		}

		maybeDelete[manifest.ID(bup.ModelStoreID)] = bup

		for tag := range bup.Tags {
			if strings.HasPrefix(tag, serviceCatTagPrefix) {
				// Precise way we concatenate all this info doesn't really matter as
				// long as it's consistent for all backups in the set and includes all
				// the pieces we need to ensure uniqueness across.
				fullTag := tenant + roid + tag
				bupsByReason[fullTag] = append(bupsByReason[fullTag], bup)
			}
		}
	}

	// For each set of backups we found, sort them by time. Mark all but the
	// youngest backup in each group as eligible for garbage collection.
	//
	// We implement this process as removing backups from the set of potential
	// backups to delete because it's possible for a backup to to not be the
	// youngest for one Reason but be the youngest for a different Reason (i.e.
	// most recent exchange mail backup but not the most recent exchange
	// contacts backup). A simple delete operation in the map is sufficient to
	// remove a backup even if it's only the youngest for a single Reason.
	// Otherwise we'd need to do another pass after this to determine the
	// isYoungest status for all Reasons in the backup.
	//
	// TODO(ashmrtn): Handle concurrent backups somehow? Right now backups that
	// have overlapping start and end times aren't explicitly handled.
	for _, bupSet := range bupsByReason {
		if len(bupSet) == 0 {
			continue
		}

		// Sort in reverse chronological order so that we can just remove the zeroth
		// item from the delete set instead of getting the slice length.
		// Unfortunately this could also put us in the pathologic case where almost
		// all items need swapped since in theory kopia returns results in
		// chronologic order and we're processing them in the order kopia returns
		// them.
		slices.SortStableFunc(bupSet, func(a, b *backup.Backup) int {
			return -a.CreationTime.Compare(b.CreationTime)
		})

		delete(maybeDelete, manifest.ID(bupSet[0].ModelStoreID))
	}

	res := make([]manifest.ID, 0, 3*len(maybeDelete))

	// For all items remaining in the delete set, generate the final set of items
	// to delete. This set includes the data snapshot ID, details snapshot ID, and
	// backup model ID to delete for each backup.
	for bupID, bup := range maybeDelete {
		// Don't need to check if we use StreamStoreID or DetailsID because
		// DetailsID was deprecated prior to tagging backups as assist backups.
		// Since the input set is only assist backups there's no overlap between the
		// two implementations.
		res = append(
			res,
			bupID,
			manifest.ID(bup.SnapshotID),
			manifest.ID(bup.StreamStoreID))
	}

	return res
}
