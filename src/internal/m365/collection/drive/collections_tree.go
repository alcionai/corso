package drive

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

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
	globalExcludeItemIDsByDrivePrefix *prefixmatcher.StringSetMatchBuilder,
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
			globalExcludeItemIDsByDrivePrefix,
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

	// hack to satisfy the linter since we're returning an error
	if ctx == nil {
		return nil, false, nil
	}

	return collections, canUsePrevBackup, errGetTreeNotImplemented
}

func (c *Collections) makeDriveCollections(
	ctx context.Context,
	drv models.Driveable,
	prevPaths map[string]string,
	prevDeltaLink string,
	limiter *pagerLimiter,
	globalExcludeItemIDsByDrivePrefix *prefixmatcher.StringSetMatchBuilder,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]string, pagers.DeltaUpdate, error) {
	driveID := ptr.Val(drv.GetId())

	ppfx, err := c.handler.PathPrefix(c.tenantID, driveID)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Wrap(err, "generating backup tree prefix")
	}

	root, err := c.handler.GetRootFolder(ctx, driveID)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Wrap(err, "getting root folder")
	}

	tree := newFolderyMcFolderFace(ppfx, ptr.Val(root.GetId()))

	counter.Add(count.PrevPaths, int64(len(prevPaths)))

	// --- delta item aggregation

	du, countPagesInDelta, err := c.populateTree(
		ctx,
		tree,
		drv,
		prevDeltaLink,
		limiter,
		counter,
		errs)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Stack(err)
	}

	// --- prev path incorporation

	err = addPrevPathsToTree(
		ctx,
		tree,
		prevPaths,
		errs)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Stack(err).Label(fault.LabelForceNoBackupCreation)
	}

	// --- post-processing

	collections, newPrevs, excludedItemIDs, err := c.turnTreeIntoCollections(
		ctx,
		tree,
		drv,
		prevPaths,
		prevDeltaLink,
		countPagesInDelta,
		errs)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Stack(err).Label(fault.LabelForceNoBackupCreation)
	}

	// only populate the global excluded items if no delta reset occurred.
	// if a reset did occur, the collections should already be marked as
	// "do not merge", therefore everything will get processed as a new addition.
	if !tree.hadReset {
		p, err := c.handler.CanonicalPath(odConsts.DriveFolderPrefixBuilder(driveID), c.tenantID)
		if err != nil {
			err = clues.WrapWC(ctx, err, "making canonical path for item exclusions")
			return nil, nil, pagers.DeltaUpdate{}, err
		}

		globalExcludeItemIDsByDrivePrefix.Add(p.String(), excludedItemIDs)
	}

	return collections, newPrevs, du, nil
}

// populateTree constructs a new tree and populates it with items
// retrieved by enumerating the delta query for the drive.
func (c *Collections) populateTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	prevDeltaLink string,
	limiter *pagerLimiter,
	counter *count.Bus,
	errs *fault.Bus,
) (pagers.DeltaUpdate, int, error) {
	ctx = clues.Add(ctx, "has_prev_delta", len(prevDeltaLink) > 0)

	var (
		currDeltaLink = prevDeltaLink
		driveID       = ptr.Val(drv.GetId())
		el            = errs.Local()
		du            pagers.DeltaUpdate
		finished      bool
		hitLimit      bool
		// TODO: plug this into the limiter
		maximumTotalDeltasAllowed int64 = 100
		// pageCounter is intended as a separate local instance
		// compared to the counter we use for other item tracking.
		// IE: don't pass it around into other funcs.
		//
		// This allows us to reset pageCounter on a reset without
		// cross-contaminating other counts.
		//
		// We use this to track three keys: 1. the total number of
		// deltas enumerated (so that we don't hit an infinite
		// loop); 2. the number of pages in each delta (for the
		// limiter, but also for the URL cache so that it knows
		// if we have too many pages for it to efficiently operate);
		// and 3. the number of items in each delta (to know if we're
		// done enumerating delta queries).
		pageCounter = counter.Local()
	)

	const (
		// track the exact number of pages across all deltas (correct across resets)
		// so that the url cache knows if it can operate within performance bounds.
		truePageCount count.Key = "pages-with-items-across-all-deltas"
	)

	// enumerate through multiple deltas until we either:
	// 1. hit a consistent state (ie: no changes since last delta enum)
	// 2. hit the limit based on the limiter
	// 3. run 100 total delta enumerations without hitting 1. (no infinite loops)
	for !hitLimit && !finished && el.Failure() == nil {
		counter.Inc(count.TotalDeltasProcessed)

		var (
			// this is used to track stats the total number of items
			// processed in each delta.  Since delta queries don't give
			// us a plain flag for "no changes occurred", we check for
			// 0 items in the delta as the "no changes occurred" state.
			// The final page of any delta query may also return 0 items,
			// so we need to combine both the item count and the deltaPageCount
			// to get a correct flag.
			iPageCounter = pageCounter.Local()
			err          error
		)

		pager := c.handler.EnumerateDriveItemsDelta(
			ctx,
			driveID,
			currDeltaLink,
			api.CallConfig{
				Select: api.DefaultDriveItemProps(),
			})

		for page, reset, done := pager.NextPage(); !done; page, reset, done = pager.NextPage() {
			if el.Failure() != nil {
				return du, 0, el.Failure()
			}

			// track the exact number of pages within a single delta (correct across resets)
			// so that we can check for "no changes occurred" results.
			// Note: don't inc `count.TotalPagesEnumerated` outside of this (ie, for the
			// truePageCount), or else we'll double up on the inc.
			iPageCounter.Inc(count.TotalPagesEnumerated)

			if reset {
				counter.Inc(count.PagerResets)
				tree.reset()
				c.resetStats()

				pageCounter = counter.Local()
				iPageCounter = pageCounter.Local()
			}

			err = c.enumeratePageOfItems(
				ctx,
				tree,
				drv,
				page,
				limiter,
				counter,
				errs)
			if err != nil {
				if errors.Is(err, errHitLimit) {
					hitLimit = true
					break
				}

				el.AddRecoverable(ctx, clues.Stack(err))
			}

			itemCount := int64(len(page))
			iPageCounter.Add(count.TotalItemsProcessed, itemCount)

			if itemCount > 0 {
				pageCounter.Inc(truePageCount)
			}

			// Stop enumeration early if we've reached the total page limit. Keep this
			// at the end of the loop so we don't request another page (pager.NextPage)
			// before seeing we've passed the limit.
			if limiter.hitPageLimit(int(pageCounter.Get(truePageCount))) {
				hitLimit = true
				break
			}
		}

		// Always cancel the pager so that even if we exit early from the loop above
		// we don't deadlock. Cancelling a pager that's already completed is
		// essentially a noop.
		pager.Cancel()

		du, err = pager.Results()
		if err != nil {
			return du, 0, clues.Stack(err)
		}

		currDeltaLink = du.URL

		// 0 pages is never expected.  We should at least have one (empty) page to
		// consume.  But checking pageCount == 1 is brittle in a non-helpful way.
		finished = iPageCounter.Get(count.TotalPagesEnumerated) < 2 &&
			iPageCounter.Get(count.TotalItemsProcessed) == 0

		// ensure we don't enumerate more than the maximum allotted count of deltas.
		if counter.Get(count.TotalDeltasProcessed) >= maximumTotalDeltasAllowed {
			err := clues.NewWC(
				ctx,
				fmt.Sprintf("unable to produce consistent delta after %d queries", maximumTotalDeltasAllowed))

			return pagers.DeltaUpdate{}, 0, err
		}
	}

	logger.Ctx(ctx).Infow(
		"enumerated collection delta",
		"stats", counter.Values(),
		"delta_stats", pageCounter.Values())

	return du, int(pageCounter.Get(truePageCount)), el.Failure()
}

func (c *Collections) enumeratePageOfItems(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	page []models.DriveItemable,
	limiter *pagerLimiter,
	counter *count.Bus,
	errs *fault.Bus,
) error {
	ctx = clues.Add(ctx, "page_lenth", len(page))
	el := errs.Local()

	for i, driveItem := range page {
		if el.Failure() != nil {
			break
		}

		var (
			item     = custom.ToCustomDriveItem(driveItem)
			isFolder = item.GetFolder() != nil || item.GetPackageEscaped() != nil
			isFile   = item.GetFile() != nil
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

		switch {
		case isFolder:
			skipped, err = c.addFolderToTree(ictx, tree, drv, item, limiter, counter)
		case isFile:
			skipped, err = c.addFileToTree(ictx, tree, drv, item, limiter, counter)
		default:
			err = clues.NewWC(ictx, "item is neither folder nor file").
				Label(fault.LabelForceNoBackupCreation, count.UnknownItemType)
		}

		if skipped != nil {
			el.AddSkip(ctx, skipped)
		}

		if err != nil {
			if errors.Is(err, errHitLimit) {
				return err
			}

			// special case: we only want to add a limited number of files
			// to each collection.  But if one collection fills up, we don't
			// want to break out of the whole backup.  That allows us to preview
			// many folders with a small selection of files in each.
			if errors.Is(err, errHitCollectionLimit) {
				continue
			}

			el.AddRecoverable(ictx, clues.Wrap(err, "adding folder"))
		}
	}

	return clues.Stack(el.Failure()).OrNil()
}

func (c *Collections) addFolderToTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	folder *custom.DriveItem,
	limiter *pagerLimiter,
	counter *count.Bus,
) (*fault.Skipped, error) {
	var (
		driveID     = ptr.Val(drv.GetId())
		folderID    = ptr.Val(folder.GetId())
		folderName  = ptr.Val(folder.GetName())
		isDeleted   = folder.GetDeleted() != nil
		isMalware   = folder.GetMalware() != nil
		isPkg       = folder.GetPackageEscaped() != nil
		notSelected bool
	)

	// check container limits before adding the next new folder
	if !tree.containsFolder(folderID) && limiter.hitContainerLimit(tree.countLiveFolders()) {
		return nil, errHitLimit
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
			graph.ItemInfo(folder))

		logger.Ctx(ctx).Infow("malware folder detected")

		return skip, nil
	}

	if isDeleted {
		err := tree.setTombstone(ctx, folder)
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

	err = tree.setFolder(ctx, folder)

	return nil, clues.Stack(err).OrNil()
}

func (c *Collections) makeFolderCollectionPath(
	ctx context.Context,
	driveID string,
	folder *custom.DriveItem,
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
	file *custom.DriveItem,
	limiter *pagerLimiter,
	counter *count.Bus,
) (*fault.Skipped, error) {
	var (
		driveID   = ptr.Val(drv.GetId())
		fileID    = ptr.Val(file.GetId())
		fileName  = ptr.Val(file.GetName())
		fileSize  = ptr.Val(file.GetSize())
		isDeleted = file.GetDeleted() != nil
		isMalware = file.GetMalware() != nil
		parent    = file.GetParentReference()
		parentID  string
	)

	if parent != nil {
		parentID = ptr.Val(parent.GetId())
	}

	defer func() {
		switch {
		case isMalware:
			counter.Inc(count.TotalMalwareProcessed)
		case isDeleted:
			counter.Inc(count.TotalDeleteFilesProcessed)
		default:
			counter.Inc(count.TotalFilesProcessed)
		}
	}()

	if isMalware {
		skip := fault.FileSkip(
			fault.SkipMalware,
			driveID,
			fileID,
			fileName,
			graph.ItemInfo(file))

		logger.Ctx(ctx).Infow("malware file detected")

		return skip, nil
	}

	if isDeleted {
		tree.deleteFile(fileID)
		return nil, nil
	}

	alreadySeen := tree.hasFile(fileID)
	parentNode, parentNotNil := tree.folderIDToNode[parentID]

	if parentNotNil && !alreadySeen {
		countSize := tree.countLiveFilesAndSizes()

		// Don't add new items if the new collection has already reached it's limit.
		// item moves and updates are generally allowed through.
		if limiter.atContainerItemsLimit(len(parentNode.files)) || limiter.hitItemLimit(countSize.numFiles) {
			return nil, errHitCollectionLimit
		}

		// Don't include large files that don't fit within the size limit.
		// Unlike the other checks, which see if we're already at the limit, this check
		// needs to be forward-facing to ensure we don't go far over the limit.
		// Example case: a 1gb limit and a 25gb file.
		if limiter.hitTotalBytesLimit(fileSize + countSize.totalBytes) {
			// don't return errHitLimit here; we only want to skip the
			// current file.  We may not want to skip files after it.
			return nil, nil
		}
	}

	err := tree.addFile(file)
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	return nil, nil
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

		coll := data.NewTombstoneCollection(
			prevDrivePath,
			c.ctrl,
			c.counter.Local())

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

func addPrevPathsToTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	prevPaths map[string]string,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for folderID, p := range prevPaths {
		if el.Failure() != nil {
			break
		}

		prevPath, err := path.FromDataLayerPath(p, false)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "invalid previous path").
				With("folderID", folderID, "prev_path", p).
				Label(count.BadPrevPath))

			continue
		}

		err = tree.setPreviousPath(folderID, prevPath)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "setting previous path").
				With("folderID", folderID, "prev_path", p))

			continue
		}
	}

	return el.Failure()
}

func (c *Collections) turnTreeIntoCollections(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	prevPaths map[string]string,
	prevDeltaLink string,
	countPagesInDelta int,
	errs *fault.Bus,
) (
	[]data.BackupCollection,
	map[string]string,
	map[string]struct{},
	error,
) {
	collectables, err := tree.generateCollectables()
	if err != nil {
		err = clues.WrapWC(ctx, err, "generating backup collection data")
		return nil, nil, nil, err
	}

	var (
		collections = []data.BackupCollection{}
		uc          *urlCache
		el          = errs.Local()
		driveID     = ptr.Val(drv.GetId())
		driveName   = ptr.Val(drv.GetName())
	)

	// Attach an url cache to the drive if the number of discovered items is
	// below the threshold. Attaching cache to larger drives can cause
	// performance issues since cache delta queries start taking up majority of
	// the hour the refreshed URLs are valid for.
	if countPagesInDelta < urlCacheDriveItemThreshold {
		logger.Ctx(ctx).Info("adding url cache for drive collections")

		uc, err = newURLCache(
			driveID,
			// we need the original prevDeltaLink here; a cache update will need
			// to process all changes since the start of the backup.  On the bright
			// side, instead of running multiple delta enumerations, all changes
			// in the backup should get compressed into the single delta query, which
			// ensures the two states are sufficiently consistent with just the
			// original delta token.
			prevDeltaLink,
			urlCacheRefreshInterval,
			c.handler,
			c.counter.Local(),
			errs)
		if err != nil {
			return nil, nil, nil, clues.StackWC(ctx, err)
		}
	}

	for _, cbl := range collectables {
		if el.Failure() != nil {
			break
		}

		coll, err := NewCollection(
			c.handler,
			c.protectedResource,
			cbl.currPath,
			cbl.prevPath,
			driveID,
			driveName,
			c.statusUpdater,
			c.ctrl,
			cbl.isPackageOrChildOfPackage,
			tree.hadReset,
			uc,
			c.counter.Local())
		if err != nil {
			return nil, nil, nil, clues.StackWC(ctx, err)
		}

		coll.driveItems = cbl.files

		collections = append(collections, coll)
	}

	if el.Failure() != nil {
		return nil, nil, nil, el.Failure()
	}

	// use the collectables and old previous paths
	// to generate new previous paths
	newPrevPaths, err := tree.generateNewPreviousPaths(collectables, prevPaths)
	if err != nil {
		return nil, nil, nil, clues.WrapWC(ctx, err, "generating new previous paths")
	}

	return collections, newPrevPaths, tree.generateExcludeItemIDs(), nil
}
