package drive

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// this file is used to separate the collections handling between the previous
// (list-based) design, and the in-progress (tree-based) redesign.
// see: https://github.com/alcionai/corso/issues/4688

func (c *Collections) getTree(
	ctx context.Context,
	prevMetadata []data.RestoreCollection,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	ctx = clues.AddTraceName(ctx, "GetTree")

	// extract the previous backup's metadata like: deltaToken urls and previousPath maps.
	// We'll need these to reconstruct / ensure the correct state of the world, after
	// enumerating through all the delta changes.
	deltasByDriveID, prevPathsByDriveID, canUsePrevBackup, err := deserializeAndValidateMetadata(
		ctx,
		prevMetadata,
		c.counter,
		errs)
	if err != nil {
		return nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePrevBackup)

	// in sharepoint, it's possible to delete an entire drive.
	// if we don't see a previously-existing drive in the drives enumeration,
	// we assume it was deleted and will remove it from storage using a tombstone.
	driveTombstones := map[string]struct{}{}
	for driveID := range prevPathsByDriveID {
		driveTombstones[driveID] = struct{}{}
	}

	pager := c.handler.NewDrivePager(c.protectedResource.ID(), nil)

	drives, err := api.GetAllDrives(ctx, pager)
	if err != nil {
		return nil, false, err
	}

	c.counter.Add(count.Drives, int64(len(drives)))
	c.counter.Add(count.PrevDeltas, int64(len(deltasByDriveID)))

	var (
		el                    = errs.Local()
		collections           = []data.BackupCollection{}
		driveIDToNewDeltaLink = map[string]string{}
		driveIDToNewPrevPaths = map[string]map[string]string{}
	)

	// each drive owns its own delta history.  We can't go more granular than that.
	// so our first order of business is to enumerate each drive's delta data, and
	// to use that as the basis for our backups.
	for _, drv := range drives {
		if el.Failure() != nil {
			break
		}

		var (
			driveID = ptr.Val(drv.GetId())
			cl      = c.counter.Local()
			ictx    = clues.Add(
				ctx,
				"drive_id", driveID,
				"drive_name", clues.Hide(ptr.Val(drv.GetName())))
		)

		ictx = clues.AddLabelCounter(ictx, cl.PlainAdder())

		// all the magic happens here.  expecations are that this process will:
		// - iterate over all data (new or delta, as needed) in the drive
		// - condense that data into a set of collections to backup
		// - stitch the new and previous path data into a new prevPaths map
		// - report the latest delta token details
		colls, newPrevPaths, du, err := c.makeDriveCollections(
			ictx,
			drv,
			prevPathsByDriveID[driveID],
			cl,
			el.Local())
		if err != nil {
			el.AddRecoverable(ictx, clues.Stack(err))
			continue
		}

		// add all the freshly aggregated data into our results
		collections = append(collections, colls...)
		driveIDToNewPrevPaths[driveID] = newPrevPaths
		driveIDToNewDeltaLink[driveID] = du.URL

		// this drive is still in use, so we'd better not delete it.
		delete(driveTombstones, driveID)
	}

	if el.Failure() != nil {
		return nil, false, clues.Stack(el.Failure())
	}

	alertIfPrevPathsHaveCollisions(ctx, driveIDToNewPrevPaths, c.counter, errs)

	// clean up any drives that have been deleted since the last backup.
	dts, err := c.makeDriveTombstones(ctx, driveTombstones, errs)
	if err != nil {
		return nil, false, clues.Stack(err)
	}

	collections = append(collections, dts...)

	// persist our updated metadata for use on the next backup
	colls := c.makeMetadataCollections(
		ctx,
		driveIDToNewDeltaLink,
		driveIDToNewPrevPaths)

	collections = append(collections, colls...)

	logger.Ctx(ctx).Infow("produced collections", "count_collections", len(collections))

	return collections, canUsePrevBackup, nil
}

func (c *Collections) makeDriveCollections(
	ctx context.Context,
	d models.Driveable,
	prevPaths map[string]string,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]string, pagers.DeltaUpdate, error) {
	cl := c.counter.Local()

	cl.Add(count.PrevPaths, int64(len(prevPaths)))
	logger.Ctx(ctx).Infow(
		"previous metadata for drive",
		"count_old_prev_paths", len(prevPaths))

	// TODO(keepers): leaving this code around for now as a guide
	// while implementation progresses.

	// --- pager aggregation

	// du, newPrevPaths, err := c.PopulateDriveCollections(
	// 	ctx,
	// 	d,
	// 	tree,
	// 	cl.Local(),
	// 	errs)
	// if err != nil {
	// 	return nil, false, clues.Stack(err)
	// }

	// numDriveItems := c.NumItems - numPrevItems
	// numPrevItems = c.NumItems

	// cl.Add(count.NewPrevPaths, int64(len(newPrevPaths)))

	// --- prev path incorporation

	// For both cases we don't need to do set difference on folder map if the
	// delta token was valid because we should see all the changes.
	// if !du.Reset {
	// 	if len(excludedItemIDs) == 0 {
	// 		continue
	// 	}

	// 	p, err := c.handler.CanonicalPath(odConsts.DriveFolderPrefixBuilder(driveID), c.tenantID)
	// 	if err != nil {
	// 		return nil, false, clues.WrapWC(ictx, err, "making exclude prefix")
	// 	}

	// 	ssmb.Add(p.String(), excludedItemIDs)

	// 	continue
	// }

	// Set all folders in previous backup but not in the current one with state
	// deleted. Need to compare by ID because it's possible to make new folders
	// with the same path as deleted old folders. We shouldn't merge items or
	// subtrees if that happens though.

	// --- post-processing

	// Attach an url cache to the drive if the number of discovered items is
	// below the threshold. Attaching cache to larger drives can cause
	// performance issues since cache delta queries start taking up majority of
	// the hour the refreshed URLs are valid for.

	// if numDriveItems < urlCacheDriveItemThreshold {
	// 	logger.Ctx(ictx).Infow(
	// 		"adding url cache for drive",
	// 		"num_drive_items", numDriveItems)

	// 	uc, err := newURLCache(
	// 		driveID,
	// 		prevDeltaLink,
	// 		urlCacheRefreshInterval,
	// 		c.handler,
	// 		cl,
	// 		errs)
	// 	if err != nil {
	// 		return nil, false, clues.Stack(err)
	// 	}

	// 	// Set the URL cache instance for all collections in this drive.
	// 	for id := range c.CollectionMap[driveID] {
	// 		c.CollectionMap[driveID][id].urlCache = uc
	// 	}
	// }

	return nil, nil, pagers.DeltaUpdate{}, clues.New("not yet implemented")
}

// quality-of-life wrapper that transforms each tombstone in the map
// into a backup collection that marks the backup as deleted.
func (c *Collections) makeDriveTombstones(
	ctx context.Context,
	driveTombstones map[string]struct{},
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	c.counter.Add(count.DriveTombstones, int64(len(driveTombstones)))

	var (
		colls = make([]data.BackupCollection, 0, len(driveTombstones))
		el    = errs.Local()
	)

	// generate tombstones for drives that were removed.
	for driveID := range driveTombstones {
		if el.Failure() != nil {
			break
		}

		prevDrivePath, err := c.handler.PathPrefix(c.tenantID, driveID)
		if err != nil {
			err = clues.WrapWC(ctx, err, "making drive tombstone for previous path").Label(count.BadPathPrefix)
			el.AddRecoverable(ctx, err)

			continue
		}

		coll, err := NewCollection(
			c.handler,
			c.protectedResource,
			nil, // delete the drive
			prevDrivePath,
			driveID,
			c.statusUpdater,
			c.ctrl,
			false,
			true,
			nil,
			c.counter.Local())
		if err != nil {
			err = clues.WrapWC(ctx, err, "making drive tombstone")
			el.AddRecoverable(ctx, err)

			continue
		}

		colls = append(colls, coll)
	}

	return colls, el.Failure()
}

// quality-of-life wrapper that transforms the delta tokens and previous paths
// into a backup collections for persitence.
func (c *Collections) makeMetadataCollections(
	ctx context.Context,
	deltaTokens map[string]string,
	prevPaths map[string]map[string]string,
) []data.BackupCollection {
	colls := []data.BackupCollection{}

	pathPrefix, err := c.handler.MetadataPathPrefix(c.tenantID)
	if err != nil {
		logger.CtxErr(ctx, err).Info("making metadata collection path prefixes")

		// It's safe to return here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		return colls
	}

	entries := []graph.MetadataCollectionEntry{
		graph.NewMetadataEntry(bupMD.DeltaURLsFileName, deltaTokens),
		graph.NewMetadataEntry(bupMD.PreviousPathFileName, prevPaths),
	}

	md, err := graph.MakeMetadataCollection(
		pathPrefix,
		entries,
		c.statusUpdater,
		c.counter.Local())
	if err != nil {
		logger.CtxErr(ctx, err).Info("making metadata collection for future incremental backups")

		// Technically it's safe to continue here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		return colls
	}

	return append(colls, md)
}
