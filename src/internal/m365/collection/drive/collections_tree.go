package drive

import (
	"context"

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
	globalExcludeItemIDs *prefixmatcher.StringSetMatchBuilder,
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

	collections, _, err := c.turnTreeIntoCollections(
		ctx,
		tree,
		driveID,
		prevDeltaLink,
		countPagesInDelta,
		errs)
	if err != nil {
		return nil, nil, pagers.DeltaUpdate{}, clues.Stack(err).Label(fault.LabelForceNoBackupCreation)
	}

	// this is a dumb hack to satisfy the linter.
	if ctx == nil {
		return nil, nil, du, nil
	}

	return collections, nil, du, errGetTreeNotImplemented
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
	ctx = clues.Add(ctx, "missing_prev_delta", len(prevDeltaLink) == 0)

	var (
		currDeltaLink = prevDeltaLink
		driveID       = ptr.Val(drv.GetId())
		el            = errs.Local()
		du            pagers.DeltaUpdate
		finished      bool
		hitLimit      bool
		// TODO: plug this into the limiter
		maxDeltas int64 = 100
		// this page counter is intentionally local and not
		// connected to the collections page counter.  It's
		// only used for tracking the enumerations in this
		// func, and we don't want it to cross contaminate
		// with other counters.
		pageCounter = count.New()
	)

	const (
		itemsInDelta  count.Key = "items-in-delta"
		totalDeltas   count.Key = "total-deltas"
		truePageCount count.Key = "true-page-count"
	)

	// enumerate through multiple deltas until we either:
	// 1. hit a consistent state (ie: no changes since last delta enum)
	// 2. hit the limit based on the limiter
	// 3. run 100 total delta enumerations without hitting 1. (no infinite loops)
	for !hitLimit && !finished && el.Failure() == nil {
		counter.Inc(count.TotalDeltasProcessed)

		var (
			iPageCounter = pageCounter.Local()
			err          error
		)

		pageCounter.Inc(totalDeltas)

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

			if reset {
				counter.Inc(count.PagerResets)
				tree.reset()
				c.resetStats()

				pageCounter = count.New()
				iPageCounter = pageCounter.Local()
			} else {
				counter.Inc(count.TotalPagesEnumerated)
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

			iPageCounter.Inc(truePageCount)
			iPageCounter.Add(itemsInDelta, int64(len(page)))

			// Stop enumeration early if we've reached the page limit. Keep this
			// at the end of the loop so we don't request another page (pager.NextPage)
			// before seeing we've passed the limit.
			if limiter.hitPageLimit(int(iPageCounter.Get(truePageCount))) {
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
		finished = iPageCounter.Get(truePageCount) < 2 &&
			iPageCounter.Get(itemsInDelta) == 0

		if pageCounter.Get(totalDeltas) >= maxDeltas {
			err := clues.NewWC(ctx, "unable to produce consistent delta after 100 queries")
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

	for i, item := range page {
		if el.Failure() != nil {
			break
		}

		var (
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

			el.AddRecoverable(ictx, clues.Wrap(err, "adding folder"))
		}
	}

	return clues.Stack(el.Failure()).OrNil()
}

func (c *Collections) addFolderToTree(
	ctx context.Context,
	tree *folderyMcFolderFace,
	drv models.Driveable,
	folder models.DriveItemable,
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
		parent      = folder.GetParentReference()
		parentID    string
		notSelected bool
	)

	// check container limits before adding the next new folder
	if !tree.containsFolder(folderID) && limiter.hitContainerLimit(tree.countLiveFolders()) {
		return nil, errHitLimit
	}

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
			graph.ItemInfo(custom.ToCustomDriveItem(folder)))

		logger.Ctx(ctx).Infow("malware folder detected")

		return skip, nil
	}

	if isDeleted {
		err := tree.setTombstone(ctx, folderID)
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

	err = tree.setFolder(ctx, parentID, folderID, folderName, isPkg)

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
	file models.DriveItemable,
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
			graph.ItemInfo(custom.ToCustomDriveItem(file)))

		logger.Ctx(ctx).Infow("malware file detected")

		return skip, nil
	}

	if isDeleted {
		tree.deleteFile(fileID)
		return nil, nil
	}

	_, alreadySeen := tree.fileIDToParentID[fileID]
	parentNode, parentNotNil := tree.folderIDToNode[parentID]

	if parentNotNil && !alreadySeen {
		countSize := tree.countLiveFilesAndSizes()

		// Don't add new items if the new collection has already reached it's limit.
		// item moves and updates are generally allowed through.
		if limiter.atContainerItemsLimit(len(parentNode.files)) || limiter.hitItemLimit(countSize.numFiles) {
			return nil, errHitLimit
		}

		// Skip large files that don't fit within the size limit.
		// unlike the other checks, which see if we're already at the limit, this check
		// needs to be forward-facing to ensure we don't go far over the limit.
		// Example case: a 1gb limit and a 25gb file.
		if limiter.hitTotalBytesLimit(fileSize + countSize.totalBytes) {
			return nil, errHitLimit
		}
	}

	err := tree.addFile(parentID, fileID, file)
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
	driveID string,
	prevDeltaLink string,
	countPagesInDelta int,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]string, error) {
	collectables, err := tree.generateCollectables()
	if err != nil {
		return nil, nil, clues.WrapWC(ctx, err, "generating backup collection data")
	}

	var (
		collections  = []data.BackupCollection{}
		newPrevPaths = map[string]string{}
		uc           *urlCache
		el           = errs.Local()
	)

	// Attach an url cache to the drive if the number of discovered items is
	// below the threshold. Attaching cache to larger drives can cause
	// performance issues since cache delta queries start taking up majority of
	// the hour the refreshed URLs are valid for.
	if countPagesInDelta < urlCacheDriveItemThreshold {
		logger.Ctx(ctx).Info("adding url cache for drive collections")

		uc, err = newURLCache(
			driveID,
			prevDeltaLink,
			urlCacheRefreshInterval,
			c.handler,
			c.counter.Local(),
			errs)
		if err != nil {
			return nil, nil, clues.StackWC(ctx, err)
		}
	}

	for id, cbl := range collectables {
		if el.Failure() != nil {
			break
		}

		if cbl.currPath != nil {
			newPrevPaths[id] = cbl.currPath.PlainString()
		}

		coll, err := NewCollection(
			c.handler,
			c.protectedResource,
			cbl.currPath,
			cbl.prevPath,
			driveID,
			c.statusUpdater,
			c.ctrl,
			cbl.isPackageOrChildOfPackage,
			tree.hadReset,
			uc,
			c.counter.Local())
		if err != nil {
			return nil, nil, clues.StackWC(ctx, err)
		}

		coll.driveItems = cbl.files

		collections = append(collections, coll)
	}

	return collections, newPrevPaths, el.Failure()
}
