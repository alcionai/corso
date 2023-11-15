package drive

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

const (
	restrictedDirectory = "Site Pages"

	defaultPreviewNumContainers              = 5
	defaultPreviewNumItemsPerContainer       = 10
	defaultPreviewNumItems                   = defaultPreviewNumContainers * defaultPreviewNumItemsPerContainer
	defaultPreviewNumBytes             int64 = 100 * 1024 * 1024
	defaultPreviewNumPages                   = 50
)

// Collections is used to retrieve drive data for a
// resource owner, which can be either a user or a sharepoint site.
type Collections struct {
	handler BackupHandler

	tenantID          string
	protectedResource idname.Provider

	statusUpdater support.StatusUpdater

	ctrl control.Options

	// collectionMap allows lookup of the data.BackupCollection
	// for a OneDrive folder.
	// driveID -> itemID -> collection
	CollectionMap map[string]map[string]*Collection

	// Track stats from drive enumeration. Represents the items backed up.
	NumItems      int
	NumFiles      int
	NumContainers int

	counter *count.Bus
}

func NewCollections(
	bh BackupHandler,
	tenantID string,
	protectedResource idname.Provider,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
	counter *count.Bus,
) *Collections {
	return &Collections{
		handler:           bh,
		tenantID:          tenantID,
		protectedResource: protectedResource,
		CollectionMap:     map[string]map[string]*Collection{},
		statusUpdater:     statusUpdater,
		ctrl:              ctrlOpts,
		counter:           counter,
	}
}

func (c *Collections) resetStats() {
	c.NumItems = 0
	c.NumFiles = 0
	c.NumContainers = 0
}

func deserializeAndValidateMetadata(
	ctx context.Context,
	cols []data.RestoreCollection,
	counter *count.Bus,
	fb *fault.Bus,
) (map[string]string, map[string]map[string]string, bool, error) {
	deltas, prevs, canUse, err := DeserializeMetadata(ctx, cols)
	if err != nil || !canUse {
		return deltas, prevs, false, clues.Stack(err).OrNil()
	}

	// Go through and remove delta tokens if we didn't have any paths for them
	// or one or more paths are empty (incorrect somehow). This will ensure we
	// don't accidentally try to pull in delta results when we should have
	// enumerated everything instead.
	//
	// Loop over the set of previous deltas because it's alright to have paths
	// without a delta but not to have a delta without paths. This way ensures
	// we check at least all the path sets for the deltas we have.
	for drive := range deltas {
		ictx := clues.Add(ctx, "drive_id", drive)

		paths := prevs[drive]
		if len(paths) == 0 {
			logger.Ctx(ictx).Info("dropping drive delta due to 0 prev paths")
			delete(deltas, drive)
		}

		// Drives have only a single delta token. If we find any folder that
		// seems like the path is bad we need to drop the entire token and start
		// fresh. Since we know the token will be gone we can also stop checking
		// for other possibly incorrect folder paths.
		for _, prevPath := range paths {
			if len(prevPath) == 0 {
				logger.Ctx(ictx).Info("dropping drive delta due to 0 len path")
				delete(deltas, drive)

				break
			}
		}
	}

	alertIfPrevPathsHaveCollisions(ctx, prevs, counter, fb)

	return deltas, prevs, canUse, nil
}

func alertIfPrevPathsHaveCollisions(
	ctx context.Context,
	prevs map[string]map[string]string,
	counter *count.Bus,
	fb *fault.Bus,
) {
	for driveID, folders := range prevs {
		prevPathCollisions := map[string]string{}

		for fid, prev := range folders {
			if otherID, collision := prevPathCollisions[prev]; collision {
				ctx = clues.Add(
					ctx,
					"collision_folder_id_1", fid,
					"collision_folder_id_2", otherID,
					"collision_drive_id", driveID,
					"collision_prev_path", path.LoggableDir(prev))

				fb.AddAlert(ctx, fault.NewAlert(
					fault.AlertPreviousPathCollision,
					"", // no namespace
					"", // no item id
					"previousPaths",
					map[string]any{
						"collision_folder_id_1": fid,
						"collision_folder_id_2": otherID,
						"collision_drive_id":    driveID,
						"collision_prev_path":   prev,
					}))

				counter.Inc(count.PreviousPathMetadataCollision)
			}

			prevPathCollisions[prev] = fid
		}
	}
}

func DeserializeMetadata(
	ctx context.Context,
	cols []data.RestoreCollection,
) (map[string]string, map[string]map[string]string, bool, error) {
	logger.Ctx(ctx).Infow(
		"deserialzing previous backup metadata",
		"num_collections", len(cols))

	var (
		prevDeltas  = map[string]string{}
		prevFolders = map[string]map[string]string{}
		errs        = fault.New(true) // metadata item reads should not fail backup
	)

	for _, col := range cols {
		if errs.Failure() != nil {
			break
		}

		items := col.Items(ctx, errs)

		for breakLoop := false; !breakLoop; {
			select {
			case <-ctx.Done():
				return nil, nil, false, clues.WrapWC(ctx, ctx.Err(), "deserializing previous backup metadata")

			case item, ok := <-items:
				if !ok {
					breakLoop = true
					break
				}

				var (
					err  error
					ictx = clues.Add(ctx, "item_uuid", item.ID())
				)

				switch item.ID() {
				case bupMD.PreviousPathFileName:
					err = DeserializeMap(item.ToReader(), prevFolders)

				case bupMD.DeltaURLsFileName:
					err = DeserializeMap(item.ToReader(), prevDeltas)

				default:
					logger.Ctx(ictx).Infow(
						"skipping unknown metadata file",
						"file_name", item.ID())

					continue
				}

				// This is conservative, but report an error if either any of the items
				// for any of the deserialized maps have duplicate drive IDs or there's
				// some other problem deserializing things. This will cause the entire
				// backup to fail, but it's not clear if higher layers would have caught
				// these cases. We can make the logic for deciding when to continue vs.
				// when to fail less strict in the future if needed.
				if err != nil {
					errs.Fail(clues.StackWC(ictx, err))

					return map[string]string{}, map[string]map[string]string{}, false, nil
				}
			}
		}
	}

	// if reads from items failed, return empty but no error
	if errs.Failure() != nil {
		logger.CtxErr(ctx, errs.Failure()).Info("reading metadata collection items")

		return map[string]string{}, map[string]map[string]string{}, false, nil
	}

	return prevDeltas, prevFolders, true, nil
}

var errExistingMapping = clues.New("mapping already exists for same drive ID")

// DeserializeMap takes an reader and a map of already deserialized items and
// adds the newly deserialized items to alreadyFound. Items are only added to
// alreadyFound if none of the keys in the freshly deserialized map already
// exist in alreadyFound. reader is closed at the end of this function.
func DeserializeMap[T any](reader io.ReadCloser, alreadyFound map[string]T) error {
	defer reader.Close()

	tmp := map[string]T{}

	if err := json.NewDecoder(reader).Decode(&tmp); err != nil {
		return clues.Wrap(err, "deserializing file contents")
	}

	var duplicate bool

	for k := range tmp {
		if _, ok := alreadyFound[k]; ok {
			duplicate = true
			break
		}
	}

	if duplicate {
		return clues.Stack(errExistingMapping)
	}

	maps.Copy(alreadyFound, tmp)

	return nil
}

// Retrieves drive data as set of `data.Collections`.
func (c *Collections) Get(
	ctx context.Context,
	prevMetadata []data.RestoreCollection,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	deltasByDriveID, prevPathsByDriveID, canUsePrevBackup, err := deserializeAndValidateMetadata(
		ctx,
		prevMetadata,
		c.counter,
		errs)
	if err != nil {
		return nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePrevBackup)

	driveTombstones := map[string]struct{}{}

	for driveID := range prevPathsByDriveID {
		driveTombstones[driveID] = struct{}{}
	}

	// Enumerate drives for the specified resourceOwner
	pager := c.handler.NewDrivePager(c.protectedResource.ID(), nil)

	drives, err := api.GetAllDrives(ctx, pager)
	if err != nil {
		return nil, false, err
	}

	c.counter.Add(count.Drives, int64(len(drives)))
	c.counter.Add(count.PrevDeltas, int64(len(deltasByDriveID)))

	var (
		driveIDToDeltaLink = map[string]string{}
		driveIDToPrevPaths = map[string]map[string]string{}
		numPrevItems       = 0
	)

	for _, d := range drives {
		var (
			cl        = c.counter.Local()
			driveID   = ptr.Val(d.GetId())
			driveName = ptr.Val(d.GetName())
			ictx      = clues.Add(
				ctx,
				"drive_id", driveID,
				"drive_name", clues.Hide(driveName))

			excludedItemIDs = map[string]struct{}{}
			oldPrevPaths    = prevPathsByDriveID[driveID]
			prevDeltaLink   = deltasByDriveID[driveID]

			// packagePaths is keyed by folder paths to a parent directory
			// which is marked as a package by its driveItem GetPackage
			// property.  Packages are only marked at the top level folder,
			// so we need this map to identify and mark all subdirs as also
			// being package cased.
			packagePaths = map[string]struct{}{}
		)

		ictx = clues.AddLabelCounter(ictx, cl.PlainAdder())

		delete(driveTombstones, driveID)

		if _, ok := driveIDToPrevPaths[driveID]; !ok {
			driveIDToPrevPaths[driveID] = map[string]string{}
		}

		if _, ok := c.CollectionMap[driveID]; !ok {
			c.CollectionMap[driveID] = map[string]*Collection{}
		}

		cl.Add(count.PrevPaths, int64(len(oldPrevPaths)))
		logger.Ctx(ictx).Infow(
			"previous metadata for drive",
			"count_old_prev_paths", len(oldPrevPaths))

		du, newPrevPaths, err := c.PopulateDriveCollections(
			ctx,
			driveID,
			driveName,
			oldPrevPaths,
			excludedItemIDs,
			packagePaths,
			prevDeltaLink,
			cl.Local(),
			errs)
		if err != nil {
			return nil, false, clues.Stack(err)
		}

		// It's alright to have an empty folders map (i.e. no folders found) but not
		// an empty delta token. This is because when deserializing the metadata we
		// remove entries for which there is no corresponding delta token/folder. If
		// we leave empty delta tokens then we may end up setting the State field
		// for collections when not actually getting delta results.
		if len(du.URL) > 0 {
			driveIDToDeltaLink[driveID] = du.URL
		}

		// Avoid the edge case where there's no paths but we do have a valid delta
		// token. We can accomplish this by adding an empty paths map for this
		// drive. If we don't have this then the next backup won't use the delta
		// token because it thinks the folder paths weren't persisted.
		driveIDToPrevPaths[driveID] = map[string]string{}
		maps.Copy(driveIDToPrevPaths[driveID], newPrevPaths)

		logger.Ctx(ictx).Infow(
			"persisted metadata for drive",
			"count_new_prev_paths", len(newPrevPaths),
			"delta_reset", du.Reset)

		numDriveItems := c.NumItems - numPrevItems
		numPrevItems = c.NumItems

		cl.Add(count.NewPrevPaths, int64(len(newPrevPaths)))

		// Attach an url cache to the drive if the number of discovered items is
		// below the threshold. Attaching cache to larger drives can cause
		// performance issues since cache delta queries start taking up majority of
		// the hour the refreshed URLs are valid for.
		if numDriveItems < urlCacheDriveItemThreshold {
			logger.Ctx(ictx).Infow(
				"adding url cache for drive",
				"num_drive_items", numDriveItems)

			uc, err := newURLCache(
				driveID,
				prevDeltaLink,
				urlCacheRefreshInterval,
				c.handler,
				cl,
				errs)
			if err != nil {
				return nil, false, clues.Stack(err)
			}

			// Set the URL cache instance for all collections in this drive.
			for id := range c.CollectionMap[driveID] {
				c.CollectionMap[driveID][id].urlCache = uc
			}
		}

		// For both cases we don't need to do set difference on folder map if the
		// delta token was valid because we should see all the changes.
		if !du.Reset {
			if len(excludedItemIDs) == 0 {
				continue
			}

			p, err := c.handler.CanonicalPath(odConsts.DriveFolderPrefixBuilder(driveID), c.tenantID)
			if err != nil {
				return nil, false, clues.WrapWC(ictx, err, "making exclude prefix")
			}

			ssmb.Add(p.String(), excludedItemIDs)

			continue
		}

		// Set all folders in previous backup but not in the current one with state
		// deleted. Need to compare by ID because it's possible to make new folders
		// with the same path as deleted old folders. We shouldn't merge items or
		// subtrees if that happens though.
		foundFolders := map[string]struct{}{}

		for id := range c.CollectionMap[driveID] {
			foundFolders[id] = struct{}{}
		}

		for fldID, p := range oldPrevPaths {
			if _, ok := foundFolders[fldID]; ok {
				continue
			}

			prevPath, err := path.FromDataLayerPath(p, false)
			if err != nil {
				err = clues.WrapWC(ictx, err, "invalid previous path").With("deleted_path", p)
				return nil, false, err
			}

			col, err := NewCollection(
				c.handler,
				c.protectedResource,
				nil, // delete the folder
				prevPath,
				driveID,
				c.statusUpdater,
				c.ctrl,
				false,
				true,
				nil,
				cl.Local())
			if err != nil {
				return nil, false, clues.WrapWC(ictx, err, "making collection")
			}

			c.CollectionMap[driveID][fldID] = col
		}
	}

	collections := []data.BackupCollection{}

	// add all the drives we found
	for _, driveColls := range c.CollectionMap {
		for _, coll := range driveColls {
			collections = append(collections, coll)
		}
	}

	c.counter.Add(count.DriveTombstones, int64(len(driveTombstones)))

	// generate tombstones for drives that were removed.
	for driveID := range driveTombstones {
		prevDrivePath, err := c.handler.PathPrefix(c.tenantID, driveID)
		if err != nil {
			return nil, false, clues.WrapWC(ctx, err, "making drive tombstone for previous path").Label(count.BadPathPrefix)
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
			return nil, false, clues.WrapWC(ctx, err, "making drive tombstone")
		}

		collections = append(collections, coll)
	}

	alertIfPrevPathsHaveCollisions(ctx, driveIDToPrevPaths, c.counter, errs)

	// add metadata collections
	pathPrefix, err := c.handler.MetadataPathPrefix(c.tenantID)
	if err != nil {
		// It's safe to return here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		logger.CtxErr(ctx, err).Info("making metadata collection path prefixes")

		return collections, canUsePrevBackup, nil
	}

	md, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(bupMD.PreviousPathFileName, driveIDToPrevPaths),
			graph.NewMetadataEntry(bupMD.DeltaURLsFileName, driveIDToDeltaLink),
		},
		c.statusUpdater,
		count.New())

	if err != nil {
		// Technically it's safe to continue here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		logger.CtxErr(ctx, err).Info("making metadata collection for future incremental backups")
	} else {
		collections = append(collections, md)
	}

	logger.Ctx(ctx).Infow("produced collections", "count_collections", len(collections))

	return collections, canUsePrevBackup, nil
}

func updateCollectionPaths(
	driveID, itemID string,
	cmap map[string]map[string]*Collection,
	curPath path.Path,
) (bool, error) {
	var initialCurPath path.Path

	col, found := cmap[driveID][itemID]
	if found && col.FullPath() != nil {
		initialCurPath = col.FullPath()
		if initialCurPath.String() == curPath.String() {
			return found, nil
		}

		col.SetFullPath(curPath)
	}

	if initialCurPath == nil {
		return found, nil
	}

	for iID, c := range cmap[driveID] {
		if iID == itemID {
			continue
		}

		colPath := c.FullPath()

		// Only updates if initialCurPath parent of colPath
		updated := colPath.UpdateParent(initialCurPath, curPath)
		if updated {
			c.SetFullPath(colPath)
		}
	}

	return found, nil
}

func (c *Collections) handleDelete(
	ctx context.Context,
	itemID, driveID string,
	oldPrevPaths, currPrevPaths, newPrevPaths map[string]string,
	isFolder bool,
	excluded map[string]struct{},
	invalidPrevDelta bool,
	counter *count.Bus,
) error {
	if !isFolder {
		counter.Inc(count.DeleteItemMarker)

		// Try to remove the item from the Collection if an entry exists for this
		// item. This handles cases where an item was created and deleted during the
		// same delta query.
		if parentID, ok := currPrevPaths[itemID]; ok {
			if col := c.CollectionMap[driveID][parentID]; col != nil {
				col.Remove(itemID)
			}

			delete(currPrevPaths, itemID)
		}

		// Don't need to add to exclude list if the delta is invalid since the
		// exclude list only matters if we're merging with a base.
		if invalidPrevDelta {
			return nil
		}

		excluded[itemID+metadata.DataFileSuffix] = struct{}{}
		excluded[itemID+metadata.MetaFileSuffix] = struct{}{}
		// Exchange counts items streamed through it which includes deletions so
		// add that here too.
		c.NumFiles++
		c.NumItems++

		return nil
	}

	counter.Inc(count.DeleteFolderMarker)

	var prevPath path.Path

	prevPathStr, ok := oldPrevPaths[itemID]
	if ok {
		var err error

		prevPath, err = path.FromDataLayerPath(prevPathStr, false)
		if err != nil {
			return clues.WrapWC(ctx, err, "invalid previous path").
				With(
					"drive_id", driveID,
					"item_id", itemID,
					"path_string", prevPathStr).
				Label(count.BadPrevPath)
		}
	}

	// Nested folders also return deleted delta results so we don't have to
	// worry about doing a prefix search in the map to remove the subtree of
	// the deleted folder/package.
	delete(newPrevPaths, itemID)

	if prevPath == nil || invalidPrevDelta {
		// It is possible that an item was created and deleted between two delta
		// invocations. In that case, it will only produce a single delete entry in
		// the delta response.
		//
		// It's also possible the item was made and deleted while getting the delta
		// results or our delta token expired and the folder was seen and now is
		// marked as deleted. If either of those is the case we should try to delete
		// the collection with this ID so it doesn't show up with items. For the
		// latter case, we rely on the set difference in the Get() function to find
		// folders that need to be marked as deleted and make collections for them.
		delete(c.CollectionMap[driveID], itemID)
		return nil
	}

	col, err := NewCollection(
		c.handler,
		c.protectedResource,
		nil, // deletes the collection
		prevPath,
		driveID,
		c.statusUpdater,
		c.ctrl,
		false,
		// DoNotMerge is not checked for deleted items.
		false,
		nil,
		counter.Local())
	if err != nil {
		return clues.Wrap(err, "making collection").With(
			"drive_id", driveID,
			"item_id", itemID,
			"path_string", prevPathStr)
	}

	c.CollectionMap[driveID][itemID] = col

	return nil
}

func (c *Collections) getCollectionPath(
	driveID string,
	item models.DriveItemable,
) (path.Path, error) {
	var (
		pb     = odConsts.DriveFolderPrefixBuilder(driveID)
		isRoot = item.GetRoot() != nil
		isFile = item.GetFile() != nil
	)

	if !isRoot {
		if item.GetParentReference() == nil ||
			item.GetParentReference().GetPath() == nil {
			err := clues.New("no parent reference").
				With("item_name", clues.Hide(ptr.Val(item.GetName())))

			return nil, err
		}

		pb = path.Builder{}.Append(path.Split(ptr.Val(item.GetParentReference().GetPath()))...)
	}

	collectionPath, err := c.handler.CanonicalPath(pb, c.tenantID)
	if err != nil {
		return nil, clues.Wrap(err, "making item path")
	}

	if isRoot || isFile {
		return collectionPath, nil
	}

	// Append folder name to path since we want the path for the collection, not
	// the path for the parent of the collection. The root and files don't need
	// to append an extra element because the root already refers to itself and
	// the collection containing the item is the parent path.
	name := ptr.Val(item.GetName())
	if len(name) == 0 {
		return nil, clues.New("folder with empty name")
	}

	collectionPath, err = collectionPath.Append(false, name)
	if err != nil {
		return nil, clues.Wrap(err, "making non-root folder path")
	}

	return collectionPath, nil
}

type driveEnumerationStats struct {
	numPages      int
	numAddedFiles int
	numContainers int
	numBytes      int64
}

func newPagerLimiter(opts control.Options) *pagerLimiter {
	res := &pagerLimiter{
		isPreview:            opts.ToggleFeatures.PreviewBackup,
		maxContainers:        opts.ItemLimits.MaxContainers,
		maxItemsPerContainer: opts.ItemLimits.MaxItemsPerContainer,
		maxItems:             opts.ItemLimits.MaxItems,
		maxBytes:             opts.ItemLimits.MaxBytes,
		maxPages:             opts.ItemLimits.MaxPages,
	}

	if res.maxContainers == 0 {
		res.maxContainers = defaultPreviewNumContainers
	}

	if res.maxItemsPerContainer == 0 {
		res.maxItemsPerContainer = defaultPreviewNumItemsPerContainer
	}

	if res.maxItems == 0 {
		res.maxItems = defaultPreviewNumItems
	}

	if res.maxBytes == 0 {
		res.maxBytes = defaultPreviewNumBytes
	}

	if res.maxPages == 0 {
		res.maxPages = defaultPreviewNumPages
	}

	return res
}

type pagerLimiter struct {
	isPreview            bool
	maxContainers        int
	maxItemsPerContainer int
	maxItems             int
	maxBytes             int64
	maxPages             int
}

func (l pagerLimiter) enabled() bool {
	return l.isPreview
}

// sizeLimit returns the total number of bytes this backup should try to
// contain.
func (l pagerLimiter) sizeLimit() int64 {
	return l.maxBytes
}

// atItemLimit returns true if the limiter is enabled and has reached the limit
// for individual items added to collections for this backup.
func (l pagerLimiter) atItemLimit(stats *driveEnumerationStats) bool {
	return l.isPreview &&
		(stats.numAddedFiles >= l.maxItems ||
			stats.numBytes >= l.maxBytes)
}

// atContainerItemsLimit returns true if the limiter is enabled and the current
// number of items is above the limit for the number of items for a container
// for this backup.
func (l pagerLimiter) atContainerItemsLimit(numItems int) bool {
	return l.isPreview && numItems >= l.maxItemsPerContainer
}

// atContainerPageLimit returns true if the limiter is enabled and the number of
// pages processed so far is beyond the limit for this backup.
func (l pagerLimiter) atPageLimit(stats *driveEnumerationStats) bool {
	return l.isPreview && stats.numPages >= l.maxPages
}

// atLimit returns true if the limiter is enabled and meets any of the
// conditions for max items, containers, etc for this backup.
func (l pagerLimiter) atLimit(stats *driveEnumerationStats) bool {
	return l.isPreview &&
		(l.atItemLimit(stats) ||
			stats.numContainers >= l.maxContainers ||
			stats.numPages >= l.maxPages)
}

// PopulateDriveCollections initializes and adds the provided drive items to Collections
// A new collection is created for every drive folder.
// Along with populating the collection items and updating the excluded item IDs, this func
// returns the current DeltaUpdate and PreviousPaths for metadata records.
func (c *Collections) PopulateDriveCollections(
	ctx context.Context,
	driveID, driveName string,
	oldPrevPaths map[string]string,
	excludedItemIDs map[string]struct{},
	topLevelPackages map[string]struct{},
	prevDeltaLink string,
	counter *count.Bus,
	errs *fault.Bus,
) (pagers.DeltaUpdate, map[string]string, error) {
	var (
		el               = errs.Local()
		newPrevPaths     = map[string]string{}
		invalidPrevDelta = len(prevDeltaLink) == 0

		// currPrevPaths is used to identify which collection a
		// file belongs to. This is useful to delete a file from the
		// collection it was previously in, in case it was moved to a
		// different collection within the same delta query
		// item ID -> item ID
		currPrevPaths = map[string]string{}

		// seenFolders is used to track the folders that we have
		// already seen. This will help us track in case a folder was
		// recreated multiple times in between a run.
		seenFolders = map[string]string{}

		limiter = newPagerLimiter(c.ctrl)
		stats   = &driveEnumerationStats{}
	)

	ctx = clues.Add(ctx, "invalid_prev_delta", invalidPrevDelta)
	logger.Ctx(ctx).Infow("running backup with limiter", "limiter", limiter)

	if !invalidPrevDelta {
		maps.Copy(newPrevPaths, oldPrevPaths)
	}

	pager := c.handler.EnumerateDriveItemsDelta(
		ctx,
		driveID,
		prevDeltaLink,
		api.CallConfig{
			Select: api.DefaultDriveItemProps(),
		})

	// Needed since folders are mixed in with items. This allows us to handle
	// hitting the maxContainer limit while (hopefully) still adding items to the
	// container we reached the limit on. It may not behave as expected across
	// page page boundaries if items in other folders have also changed.
	var lastFolderPath string

	for page, reset, done := pager.NextPage(); !done; page, reset, done = pager.NextPage() {
		if el.Failure() != nil {
			break
		}

		counter.Inc(count.PagesEnumerated)

		if reset {
			counter.Inc(count.PagerResets)

			ctx = clues.Add(ctx, "delta_reset_occurred", true)
			newPrevPaths = map[string]string{}
			currPrevPaths = map[string]string{}
			seenFolders = map[string]string{}
			c.CollectionMap[driveID] = map[string]*Collection{}
			invalidPrevDelta = true

			// Reset collections and stats counts since we're starting over.
			c.resetStats()

			stats = &driveEnumerationStats{}
		}

		for _, item := range page {
			if el.Failure() != nil {
				break
			}

			// Check if we got the max number of containers we're looking for and also
			// processed items for the final container.
			if limiter.enabled() {
				if item.GetFolder() != nil || item.GetPackageEscaped() != nil {
					// Don't check for containers we've already seen.
					if _, ok := c.CollectionMap[driveID][ptr.Val(item.GetId())]; !ok {
						cp, err := c.getCollectionPath(driveID, item)
						if err != nil {
							el.AddRecoverable(ctx, clues.Stack(err).
								WithClues(ctx).
								Label(fault.LabelForceNoBackupCreation))

							continue
						}

						if cp.String() != lastFolderPath {
							if limiter.atLimit(stats) {
								break
							}

							lastFolderPath = cp.String()
							stats.numContainers++
						}
					}
				}
			}

			err := c.processItem(
				ctx,
				item,
				driveID,
				driveName,
				oldPrevPaths,
				currPrevPaths,
				newPrevPaths,
				seenFolders,
				excludedItemIDs,
				topLevelPackages,
				invalidPrevDelta,
				counter,
				stats,
				limiter,
				el)
			if err != nil {
				el.AddRecoverable(ctx, clues.Stack(err))
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
		return du, nil, clues.Stack(err)
	}

	logger.Ctx(ctx).Infow("populated collection", "stats", counter.Values())

	return du, newPrevPaths, el.Failure()
}

func (c *Collections) processItem(
	ctx context.Context,
	item models.DriveItemable,
	driveID, driveName string,
	oldPrevPaths, currPrevPaths, newPrevPaths map[string]string,
	seenFolders map[string]string,
	excludedItemIDs map[string]struct{},
	topLevelPackages map[string]struct{},
	invalidPrevDelta bool,
	counter *count.Bus,
	stats *driveEnumerationStats,
	limiter *pagerLimiter,
	skipper fault.AddSkipper,
) error {
	var (
		itemID   = ptr.Val(item.GetId())
		itemName = ptr.Val(item.GetName())
		isFolder = item.GetFolder() != nil || item.GetPackageEscaped() != nil
	)

	ctx = clues.Add(
		ctx,
		"item_id", itemID,
		"item_name", clues.Hide(itemName),
		"item_is_folder", isFolder)

	if item.GetMalware() != nil {
		addtl := graph.ItemInfo(item)
		skip := fault.FileSkip(fault.SkipMalware, driveID, itemID, itemName, addtl)

		if isFolder {
			skip = fault.ContainerSkip(fault.SkipMalware, driveID, itemID, itemName, addtl)
		}

		skipper.AddSkip(ctx, skip)
		logger.Ctx(ctx).Infow("malware detected", "item_details", addtl)
		counter.Inc(count.Malware)

		return nil
	}

	// Deleted file or folder.
	if item.GetDeleted() != nil {
		err := c.handleDelete(
			ctx,
			itemID,
			driveID,
			oldPrevPaths,
			currPrevPaths,
			newPrevPaths,
			isFolder,
			excludedItemIDs,
			invalidPrevDelta,
			counter)

		return clues.StackWC(ctx, err).OrNil()
	}

	collectionPath, err := c.getCollectionPath(driveID, item)
	if err != nil {
		return clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation, count.BadCollPath)
	}

	// Skip items that don't match the folder selectors we were given.
	if shouldSkip(ctx, collectionPath, c.handler, driveName) {
		counter.Inc(count.SkippedContainers)
		logger.Ctx(ctx).Debugw("path not selected", "skipped_path", collectionPath.String())

		return nil
	}

	switch {
	case isFolder:
		// Deletions are handled above so this is just moves/renames.
		var prevPath path.Path

		prevPathStr, ok := oldPrevPaths[itemID]
		if ok {
			prevPath, err = path.FromDataLayerPath(prevPathStr, false)
			if err != nil {
				return clues.WrapWC(ctx, err, "invalid previous path").
					With("prev_path_string", path.LoggableDir(prevPathStr)).
					Label(count.BadPrevPath)
			}
		} else if item.GetRoot() != nil {
			// Root doesn't move or get renamed.
			prevPath = collectionPath
		}

		// Moved folders don't cause delta results for any subfolders nested in
		// them. We need to go through and update paths to handle that. We only
		// update newPaths so we don't accidentally clobber previous deletes.
		updatePath(newPrevPaths, itemID, collectionPath.String())

		found, err := updateCollectionPaths(
			driveID,
			itemID,
			c.CollectionMap,
			collectionPath)
		if err != nil {
			return clues.StackWC(ctx, err)
		}

		if found {
			return nil
		}

		isPackage := item.GetPackageEscaped() != nil
		if isPackage {
			counter.Inc(count.Packages)
			// mark this path as a package type for all other collections.
			// any subfolder should get marked as a childOfPackage below.
			topLevelPackages[collectionPath.String()] = struct{}{}
		} else {
			counter.Inc(count.Folders)
		}

		childOfPackage := filters.
			PathPrefix(maps.Keys(topLevelPackages)).
			Compare(collectionPath.String())

		// This check is to ensure that if a folder was deleted and
		// recreated multiple times between a backup, we only use the
		// final one.
		alreadyHandledFolderID, collPathAlreadyExists := seenFolders[collectionPath.String()]
		collPathAlreadyExists = collPathAlreadyExists && alreadyHandledFolderID != itemID

		if collPathAlreadyExists {
			// we don't have a good way of juggling multiple previous paths
			// at this time.  If a path was declared twice, it's a bit ambiguous
			// which prior data the current folder now contains.  Safest thing to
			// do is to call it a new folder and ingest items fresh.
			prevPath = nil

			c.NumContainers--
			c.NumItems--

			delete(c.CollectionMap[driveID], alreadyHandledFolderID)
			delete(newPrevPaths, alreadyHandledFolderID)
		}

		if invalidPrevDelta {
			prevPath = nil
		}

		seenFolders[collectionPath.String()] = itemID

		col, err := NewCollection(
			c.handler,
			c.protectedResource,
			collectionPath,
			prevPath,
			driveID,
			c.statusUpdater,
			c.ctrl,
			isPackage || childOfPackage,
			invalidPrevDelta || collPathAlreadyExists,
			nil,
			counter.Local())
		if err != nil {
			return clues.StackWC(ctx, err)
		}

		col.driveName = driveName

		c.CollectionMap[driveID][itemID] = col
		c.NumContainers++

		if item.GetRoot() != nil {
			return nil
		}

		// Add an entry to fetch permissions into this collection. This assumes
		// that OneDrive always returns all folders on the path of an item
		// before the item. This seems to hold true for now at least.
		if col.Add(item) {
			c.NumItems++
		}

	case item.GetFile() != nil:
		counter.Inc(count.Files)

		// Deletions are handled above so this is just moves/renames.
		if len(ptr.Val(item.GetParentReference().GetId())) == 0 {
			return clues.NewWC(ctx, "file without parent ID").Label(count.MissingParent)
		}

		// Get the collection for this item.
		parentID := ptr.Val(item.GetParentReference().GetId())
		ctx = clues.Add(ctx, "parent_id", parentID)

		collection, ok := c.CollectionMap[driveID][parentID]
		if !ok {
			return clues.NewWC(ctx, "item seen before parent folder").Label(count.ItemBeforeParent)
		}

		// Don't move items if the new collection's already reached it's limit. This
		// helps ensure we don't get some pathological case where we end up dropping
		// a bunch of items that got moved.
		//
		// We need to check if the collection already contains the item though since
		// it could be an item update instead of a move.
		if !collection.ContainsItem(item) &&
			limiter.atContainerItemsLimit(collection.AddedItems()) {
			return nil
		}

		// Skip large files that don't fit within the size limit.
		if limiter.enabled() &&
			limiter.sizeLimit() < ptr.Val(item.GetSize())+stats.numBytes {
			return nil
		}

		// This will only kick in if the file was moved multiple times
		// within a single delta query.  We delete the file from the previous
		// collection so that it doesn't appear in two places.
		prevParentContainerID, alreadyAdded := currPrevPaths[itemID]
		if alreadyAdded {
			prevColl, found := c.CollectionMap[driveID][prevParentContainerID]
			if !found {
				return clues.NewWC(ctx, "previous collection not found").
					With("prev_parent_container_id", prevParentContainerID)
			}

			if ok := prevColl.Remove(itemID); !ok {
				return clues.NewWC(ctx, "removing item from prev collection").
					With("prev_parent_container_id", prevParentContainerID)
			}
		}

		currPrevPaths[itemID] = parentID

		// Only increment counters if the file didn't already get counted (i.e. it's
		// not an item that was either updated or moved during the delta query).
		if collection.Add(item) && !alreadyAdded {
			c.NumItems++
			c.NumFiles++
			stats.numAddedFiles++
			stats.numBytes += ptr.Val(item.GetSize())
		}

		// Do this after adding the file to the collection so if we fail to add
		// the item to the collection for some reason and we're using best effort
		// we don't just end up deleting the item in the resulting backup. The
		// resulting backup will be slightly incorrect, but it will have the most
		// data that we were able to preserve.
		if !invalidPrevDelta {
			// Always add a file to the excluded list. The file may have been
			// renamed/moved/modified, so we still have to drop the
			// original one and download a fresh copy.
			excludedItemIDs[itemID+metadata.DataFileSuffix] = struct{}{}
			excludedItemIDs[itemID+metadata.MetaFileSuffix] = struct{}{}
		}

	default:
		return clues.NewWC(ctx, "item is neither folder nor file").
			Label(fault.LabelForceNoBackupCreation, count.UnknownItemType)
	}

	return nil
}

type dirScopeChecker interface {
	IsAllPass() bool
	IncludesDir(dir string) bool
}

func shouldSkip(
	ctx context.Context,
	drivePath path.Path,
	dsc dirScopeChecker,
	driveName string,
) bool {
	return !includePath(ctx, dsc, drivePath) ||
		(drivePath.Category() == path.LibrariesCategory && restrictedDirectory == driveName)
}

func includePath(ctx context.Context, dsc dirScopeChecker, folderPath path.Path) bool {
	// Check if the folder is allowed by the scope.
	pb, err := path.GetDriveFolderPath(folderPath)
	if err != nil {
		logger.Ctx(ctx).With("err", err).Error("getting drive folder path")
		return true
	}

	// Hack for the edge case where we're looking at the root folder and can
	// select any folder. Right now the root folder has an empty folder path.
	if len(pb.Elements()) == 0 && dsc.IsAllPass() {
		return true
	}

	return dsc.IncludesDir(pb.String())
}

func updatePath(paths map[string]string, id, newPath string) {
	currPath := paths[id]
	if len(currPath) == 0 {
		paths[id] = newPath
		return
	}

	if currPath == newPath {
		return
	}

	// We need to do a prefix search on the rest of the map to update the subtree.
	// We don't need to make collections for all of these, as hierarchy merging in
	// other components should take care of that. We do need to ensure that the
	// resulting map contains all folders though so we know the next time around.
	for folderID, p := range paths {
		if !strings.HasPrefix(p, currPath) {
			continue
		}

		paths[folderID] = strings.Replace(p, currPath, newPath, 1)
	}
}
