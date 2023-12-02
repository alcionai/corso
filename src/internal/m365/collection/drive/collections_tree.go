package drive

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// ---------------------------------------------------------------------------
// Processing
// ---------------------------------------------------------------------------

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
	limiter := newPagerLimiter(c.ctrl)

	logger.Ctx(ctx).Infow(
		"running backup: getting collection data using tree structure",
		"limits", c.ctrl.PreviewLimits,
		"effective_limits", limiter.effectiveLimits(),
		"preview_mode", limiter.enabled())

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
			deltasByDriveID[driveID],
			limiter,
			cl,
			el)
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

var errTreeNotImplemented = clues.New("backup tree not implemented")

func (c *Collections) makeDriveCollections(
	ctx context.Context,
	drv models.Driveable,
	prevPaths map[string]string,
	prevDeltaLink string,
	limiter *pagerLimiter,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]string, pagers.DeltaUpdate, error) {
	ppfx, err := c.handler.PathPrefix(c.tenantID, ptr.Val(drv.GetId()))
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Wrap(err, "generating backup tree prefix")
	}

	var (
		tree  = newFolderyMcFolderFace(ppfx)
		stats = &driveEnumerationStats{}
	)

	counter.Add(count.PrevPaths, int64(len(prevPaths)))

	// --- delta item aggregation

	du, err := c.populateTree(
		ctx,
		tree,
		limiter,
		stats,
		drv,
		prevDeltaLink,
		counter,
		errs)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Stack(err)
	}

	// numDriveItems := c.NumItems - numPrevItems
	// numPrevItems = c.NumItems

	// cl.Add(count.NewPrevPaths, int64(len(newPrevPaths)))

	// TODO(keepers): leaving this code around for now as a guide
	// while implementation progresses.

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

	// this is a dumb hack to satisfy the linter.
	if ctx == nil {
		return nil, nil, du, nil
	}

	return nil, nil, du, errTreeNotImplemented
}

// populateTree constructs a new tree and populates it with items
// retrieved by enumerating the delta query for the drive.
func (c *Collections) populateTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	limiter *pagerLimiter,
	stats *driveEnumerationStats,
	drv models.Driveable,
	prevDeltaLink string,
	counter *count.Bus,
	errs *fault.Bus,
) (pagers.DeltaUpdate, error) {
	ctx = clues.Add(ctx, "invalid_prev_delta", len(prevDeltaLink) == 0)

	var (
		driveID = ptr.Val(drv.GetId())
		el      = errs.Local()
	)

	// TODO(keepers): to end in a correct state, we'll eventually need to run this
	// query multiple times over, until it ends in an empty change set.
	pager := c.handler.EnumerateDriveItemsDelta(
		ctx,
		driveID,
		prevDeltaLink,
		api.CallConfig{
			Select: api.DefaultDriveItemProps(),
		})

	for page, reset, done := pager.NextPage(); !done; page, reset, done = pager.NextPage() {
		if el.Failure() != nil {
			break
		}

		counter.Inc(count.PagesEnumerated)

		if reset {
			counter.Inc(count.PagerResets)
			tree.Reset()
			c.resetStats()

			*stats = driveEnumerationStats{}
		}

		err := c.enumeratePageOfItems(
			ctx,
			tree,
			limiter,
			stats,
			drv,
			page,
			counter,
			errs)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
		}

		// Stop enumeration early if we've reached the item or page limit. Do this
		// at the end of the loop so we don't request another page in the
		// background.
		//
		// We don't want to break on just the container limit here because it's
		// possible that there's more items in the current (final) container that
		// we're processing. We need to see the next page to determine if we've
		// reached the end of the container. Note that this doesn't take into
		// account the number of items in the current container, so it's possible it
		// will fetch more data when it doesn't really need to.
		if limiter.atPageLimit(stats) || limiter.atItemLimit(stats) {
			break
		}
	}

	// Always cancel the pager so that even if we exit early from the loop above
	// we don't deadlock. Cancelling a pager that's already completed is
	// essentially a noop.
	pager.Cancel()

	du, err := pager.Results()
	if err != nil {
		return du, clues.Stack(err)
	}

	logger.Ctx(ctx).Infow("enumerated collection delta", "stats", counter.Values())

	return du, el.Failure()
}

func (c *Collections) enumeratePageOfItems(
	ctx context.Context,
	tree *folderyMcFolderFace,
	limiter *pagerLimiter,
	stats *driveEnumerationStats,
	drv models.Driveable,
	page []models.DriveItemable,
	counter *count.Bus,
	errs *fault.Bus,
) error {
	ctx = clues.Add(ctx, "page_lenth", len(page))
	el := errs.Local()

	for i, item := range page {
		if el.Failure() != nil {
			break
		}

		var (
			isFolder = item.GetFolder() != nil || item.GetPackageEscaped() != nil
			itemID   = ptr.Val(item.GetId())
			err      error
			skipped  *fault.Skipped
		)

		ictx := clues.Add(
			ctx,
			"item_id", itemID,
			"item_name", clues.Hide(ptr.Val(item.GetName())),
			"item_index", i,
			"item_is_folder", isFolder,
			"item_is_package", item.GetPackageEscaped() != nil)

		if isFolder {
			// check if the preview needs to exit before adding each folder
			if !tree.ContainsFolder(itemID) && limiter.atLimit(stats, len(tree.folderIDToNode)) {
				break
			}

			skipped, err = c.addFolderToTree(ictx, tree, drv, item, stats, counter)
		} else {
			skipped, err = c.addFileToTree(ictx, tree, drv, item, stats, counter)
		}

		if skipped != nil {
			el.AddSkip(ctx, skipped)
		}

		if err != nil {
			el.AddRecoverable(ictx, clues.Wrap(err, "adding folder"))
		}

		// Check if we reached the item or size limit while processing this page.
		// The check after this loop will get us out of the pager.
		// We don't want to check all limits because it's possible we've reached
		// the container limit but haven't reached the item limit or really added
		// items to the last container we found.
		if limiter.atItemLimit(stats) {
			break
		}
	}

	stats.numPages++

	return clues.Stack(el.Failure()).OrNil()
}

func (c *Collections) addFolderToTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	folder models.DriveItemable,
	stats *driveEnumerationStats,
	counter *count.Bus,
) (*fault.Skipped, error) {
	var (
		driveID     = ptr.Val(drv.GetId())
		folderID    = ptr.Val(folder.GetId())
		folderName  = ptr.Val(folder.GetName())
		isDeleted   = folder.GetDeleted() != nil
		isMalware   = folder.GetMalware() != nil
		isPkg       = folder.GetPackageEscaped() != nil
		parent      = folder.GetParentReference()
		parentID    string
		notSelected bool
	)

	if parent != nil {
		parentID = ptr.Val(parent.GetId())
	}

	defer func() {
		switch {
		case notSelected:
			counter.Inc(count.TotalContainersSkipped)
		case isMalware:
			counter.Inc(count.TotalMalwareProcessed)
		case isDeleted:
			counter.Inc(count.TotalDeleteFoldersProcessed)
		case isPkg:
			counter.Inc(count.TotalPackagesProcessed)
		default:
			counter.Inc(count.TotalFoldersProcessed)
		}
	}()

	// FIXME(keepers): if we don't track this as previously visited,
	// we could add a skip multiple times, every time we visit the
	// folder again at the top of the page.
	if isMalware {
		skip := fault.ContainerSkip(
			fault.SkipMalware,
			driveID,
			folderID,
			folderName,
			graph.ItemInfo(custom.ToLiteDriveItemable(folder)))

		logger.Ctx(ctx).Infow("malware detected")

		return skip, nil
	}

	if isDeleted {
		err := tree.SetTombstone(ctx, folderID)
		return nil, clues.Stack(err).OrNil()
	}

	collectionPath, err := c.makeFolderCollectionPath(ctx, driveID, folder)
	if err != nil {
		return nil, clues.Stack(err).Label(fault.LabelForceNoBackupCreation, count.BadCollPath)
	}

	// Skip items that don't match the folder selectors we were given.
	notSelected = shouldSkip(ctx, collectionPath, c.handler, ptr.Val(drv.GetName()))
	if notSelected {
		logger.Ctx(ctx).Debugw("path not selected", "skipped_path", collectionPath.String())
		return nil, nil
	}

	err = tree.SetFolder(ctx, parentID, folderID, folderName, isPkg)

	return nil, clues.Stack(err).OrNil()
}

func (c *Collections) makeFolderCollectionPath(
	ctx context.Context,
	driveID string,
	folder models.DriveItemable,
) (path.Path, error) {
	if folder.GetRoot() != nil {
		pb := odConsts.DriveFolderPrefixBuilder(driveID)
		collectionPath, err := c.handler.CanonicalPath(pb, c.tenantID)

		return collectionPath, clues.WrapWC(ctx, err, "making canonical root path").OrNil()
	}

	if folder.GetParentReference() == nil || folder.GetParentReference().GetPath() == nil {
		return nil, clues.NewWC(ctx, "no parent reference in folder").Label(count.MissingParent)
	}

	// Append folder name to path since we want the path for the collection, not
	// the path for the parent of the collection.
	name := ptr.Val(folder.GetName())
	if len(name) == 0 {
		return nil, clues.NewWC(ctx, "missing folder name")
	}

	folderPath := path.Split(ptr.Val(folder.GetParentReference().GetPath()))
	folderPath = append(folderPath, name)
	pb := path.Builder{}.Append(folderPath...)
	collectionPath, err := c.handler.CanonicalPath(pb, c.tenantID)

	return collectionPath, clues.WrapWC(ctx, err, "making folder collection path").OrNil()
}

func (c *Collections) addFileToTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	item models.DriveItemable,
	stats *driveEnumerationStats,
	counter *count.Bus,
) (*fault.Skipped, error) {
	return nil, clues.New("not yet implemented")
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

		// TODO: call NewTombstoneCollection
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
