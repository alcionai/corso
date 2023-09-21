package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type collectionScope int

const (
	// CollectionScopeUnknown is used when we don't know and don't need
	// to know the kind, like in the case of deletes
	CollectionScopeUnknown collectionScope = 0

	// CollectionScopeFolder is used for regular folder collections
	CollectionScopeFolder collectionScope = 1

	// CollectionScopePackage is used to represent OneNote items
	CollectionScopePackage collectionScope = 2
)

const restrictedDirectory = "Site Pages"

// Collections is used to retrieve drive data for a
// resource owner, which can be either a user or a sharepoint site.
type Collections struct {
	handler BackupHandler

	tenantID      string
	resourceOwner string

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
}

func NewCollections(
	bh BackupHandler,
	tenantID string,
	resourceOwner string,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collections {
	return &Collections{
		handler:       bh,
		tenantID:      tenantID,
		resourceOwner: resourceOwner,
		CollectionMap: map[string]map[string]*Collection{},
		statusUpdater: statusUpdater,
		ctrl:          ctrlOpts,
	}
}

func deserializeMetadata(
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
				return nil, nil, false, clues.Wrap(ctx.Err(), "deserialzing previous backup metadata").WithClues(ctx)

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

				if err == nil {
					// Successful decode.
					continue
				}

				// This is conservative, but report an error if either any of the items
				// for any of the deserialized maps have duplicate drive IDs or there's
				// some other problem deserializing things. This will cause the entire
				// backup to fail, but it's not clear if higher layers would have caught
				// these cases. We can make the logic for deciding when to continue vs.
				// when to fail less strict in the future if needed.
				if err != nil {
					return nil, nil, false, clues.Stack(err).WithClues(ictx)
				}
			}
		}

		// Go through and remove delta tokens if we didn't have any paths for them
		// or one or more paths are empty (incorrect somehow). This will ensure we
		// don't accidentally try to pull in delta results when we should have
		// enumerated everything instead.
		//
		// Loop over the set of previous deltas because it's alright to have paths
		// without a delta but not to have a delta without paths. This way ensures
		// we check at least all the path sets for the deltas we have.
		for drive := range prevDeltas {
			paths := prevFolders[drive]
			if len(paths) == 0 {
				delete(prevDeltas, drive)
			}

			// Drives have only a single delta token. If we find any folder that
			// seems like the path is bad we need to drop the entire token and start
			// fresh. Since we know the token will be gone we can also stop checking
			// for other possibly incorrect folder paths.
			for _, prevPath := range paths {
				if len(prevPath) == 0 {
					delete(prevDeltas, drive)
					break
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

// Retrieves drive data as set of `data.Collections` and a set of item names to
// be excluded from the upcoming backup.
func (c *Collections) Get(
	ctx context.Context,
	prevMetadata []data.RestoreCollection,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	prevDriveIDToDelta, oldPrevPathsByDriveID, canUsePrevBackup, err := deserializeMetadata(ctx, prevMetadata)
	if err != nil {
		return nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePrevBackup)

	driveTombstones := map[string]struct{}{}

	for driveID := range oldPrevPathsByDriveID {
		driveTombstones[driveID] = struct{}{}
	}

	progressBar := observe.MessageWithCompletion(
		ctx,
		observe.Bulletf(path.FilesCategory.HumanString()))
	defer close(progressBar)

	// Enumerate drives for the specified resourceOwner
	pager := c.handler.NewDrivePager(c.resourceOwner, nil)

	drives, err := api.GetAllDrives(ctx, pager)
	if err != nil {
		return nil, false, err
	}

	var (
		driveIDToDeltaLink = map[string]string{}
		driveIDToPrevPaths = map[string]map[string]string{}
		numPrevItems       = 0
	)

	for _, d := range drives {
		var (
			driveID   = ptr.Val(d.GetId())
			driveName = ptr.Val(d.GetName())
			ictx      = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)

			excludedItemIDs = map[string]struct{}{}
			oldPrevPaths    = oldPrevPathsByDriveID[driveID]
			prevDeltaLink   = prevDriveIDToDelta[driveID]
		)

		delete(driveTombstones, driveID)

		if _, ok := driveIDToPrevPaths[driveID]; !ok {
			driveIDToPrevPaths[driveID] = map[string]string{}
		}

		if _, ok := c.CollectionMap[driveID]; !ok {
			c.CollectionMap[driveID] = map[string]*Collection{}
		}

		logger.Ctx(ictx).Infow(
			"previous metadata for drive",
			"num_paths_entries", len(oldPrevPaths))

		items, du, err := c.handler.EnumerateDriveItemsDelta(
			ictx,
			driveID,
			prevDeltaLink)
		if err != nil {
			return nil, false, err
		}

		// It's alright to have an empty folders map (i.e. no folders found) but not
		// an empty delta token. This is because when deserializing the metadata we
		// remove entries for which there is no corresponding delta token/folder. If
		// we leave empty delta tokens then we may end up setting the State field
		// for collections when not actually getting delta results.
		if len(du.URL) > 0 {
			driveIDToDeltaLink[driveID] = du.URL
		}

		newPrevPaths, err := c.UpdateCollections(
			ctx,
			driveID,
			driveName,
			items,
			oldPrevPaths,
			driveIDToPrevPaths[driveID],
			excludedItemIDs,
			du.Reset,
			errs)
		if err != nil {
			return nil, false, clues.Stack(err)
		}

		// Avoid the edge case where there's no paths but we do have a valid delta
		// token. We can accomplish this by adding an empty paths map for this
		// drive. If we don't have this then the next backup won't use the delta
		// token because it thinks the folder paths weren't persisted.
		driveIDToPrevPaths[driveID] = map[string]string{}
		maps.Copy(driveIDToPrevPaths[driveID], newPrevPaths)

		logger.Ctx(ictx).Infow(
			"persisted metadata for drive",
			"num_new_paths_entries", len(newPrevPaths),
			"delta_reset", du.Reset)

		numDriveItems := c.NumItems - numPrevItems
		numPrevItems = c.NumItems

		// Attach an url cache
		if numDriveItems < urlCacheDriveItemThreshold {
			logger.Ctx(ictx).Info("adding url cache for drive")

			err = c.addURLCacheToDriveCollections(
				ictx,
				driveID,
				prevDeltaLink,
				errs)
			if err != nil {
				return nil, false, err
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
				return nil, false, clues.Wrap(err, "making exclude prefix").WithClues(ictx)
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
				err = clues.Wrap(err, "invalid previous path").WithClues(ictx).With("deleted_path", p)
				return nil, false, err
			}

			col, err := NewCollection(
				c.handler,
				nil, // delete the folder
				prevPath,
				driveID,
				c.statusUpdater,
				c.ctrl,
				CollectionScopeUnknown,
				true,
				nil)
			if err != nil {
				return nil, false, clues.Wrap(err, "making collection").WithClues(ictx)
			}

			c.CollectionMap[driveID][fldID] = col
		}
	}

	observe.Message(ctx, fmt.Sprintf("Discovered %d items to backup", c.NumItems))

	collections := []data.BackupCollection{}

	// add all the drives we found
	for _, driveColls := range c.CollectionMap {
		for _, coll := range driveColls {
			collections = append(collections, coll)
		}
	}

	// generate tombstones for drives that were removed.
	for driveID := range driveTombstones {
		prevDrivePath, err := c.handler.PathPrefix(c.tenantID, driveID)
		if err != nil {
			return nil, false, clues.Wrap(err, "making drive tombstone for previous path").WithClues(ctx)
		}

		coll, err := NewCollection(
			c.handler,
			nil, // delete the drive
			prevDrivePath,
			driveID,
			c.statusUpdater,
			c.ctrl,
			CollectionScopeUnknown,
			true,
			nil)
		if err != nil {
			return nil, false, clues.Wrap(err, "making drive tombstone").WithClues(ctx)
		}

		collections = append(collections, coll)
	}

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
		c.statusUpdater)

	if err != nil {
		// Technically it's safe to continue here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		logger.CtxErr(ctx, err).Info("making metadata collection for future incremental backups")
	} else {
		collections = append(collections, md)
	}

	return collections, canUsePrevBackup, nil
}

// addURLCacheToDriveCollections adds an URL cache to all collections belonging to
// a drive.
func (c *Collections) addURLCacheToDriveCollections(
	ctx context.Context,
	driveID, prevDelta string,
	errs *fault.Bus,
) error {
	uc, err := newURLCache(
		driveID,
		prevDelta,
		urlCacheRefreshInterval,
		c.handler,
		errs)
	if err != nil {
		return err
	}

	// Set the URL cache for all collections in this drive
	for _, driveColls := range c.CollectionMap {
		for _, coll := range driveColls {
			coll.urlCache = uc
		}
	}

	return nil
}

func updateCollectionPaths(
	driveID, itemID string,
	cmap map[string]map[string]*Collection,
	curPath path.Path,
) (bool, error) {
	var initialCurPath path.Path

	col, found := cmap[driveID][itemID]
	if found {
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
	itemID, driveID string,
	oldPrevPaths, currPrevPaths, newPrevPaths map[string]string,
	isFolder bool,
	excluded map[string]struct{},
	invalidPrevDelta bool,
) error {
	if !isFolder {
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

	var prevPath path.Path

	prevPathStr, ok := oldPrevPaths[itemID]
	if ok {
		var err error

		prevPath, err = path.FromDataLayerPath(prevPathStr, false)
		if err != nil {
			return clues.Wrap(err, "invalid previous path").
				With(
					"drive_id", driveID,
					"item_id", itemID,
					"path_string", prevPathStr)
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
		nil, // deletes the collection
		prevPath,
		driveID,
		c.statusUpdater,
		c.ctrl,
		CollectionScopeUnknown,
		// DoNotMerge is not checked for deleted items.
		false,
		nil)
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

// UpdateCollections initializes and adds the provided drive items to Collections
// A new collection is created for every drive folder (or package).
// oldPrevPaths is the unchanged data that was loaded from the metadata file.
// This map is not modified during the call.
// currPrevPaths starts as a copy of oldPaths and is updated as changes are found in
// the returned results.  Items are added to this collection throughout the call.
// newPrevPaths, ie: the items added during this call, get returned as a map.
func (c *Collections) UpdateCollections(
	ctx context.Context,
	driveID, driveName string,
	items []models.DriveItemable,
	oldPrevPaths map[string]string,
	currPrevPaths map[string]string,
	excluded map[string]struct{},
	invalidPrevDelta bool,
	errs *fault.Bus,
) (map[string]string, error) {
	var (
		el           = errs.Local()
		newPrevPaths = map[string]string{}
	)

	if !invalidPrevDelta {
		maps.Copy(newPrevPaths, oldPrevPaths)
	}

	for _, item := range items {
		if el.Failure() != nil {
			break
		}

		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			isFolder = item.GetFolder() != nil || item.GetPackageEscaped() != nil
			ictx     = clues.Add(
				ctx,
				"item_id", itemID,
				"item_name", clues.Hide(itemName),
				"item_is_folder", isFolder)
		)

		if item.GetMalware() != nil {
			addtl := graph.ItemInfo(item)
			skip := fault.FileSkip(fault.SkipMalware, driveID, itemID, itemName, addtl)

			if isFolder {
				skip = fault.ContainerSkip(fault.SkipMalware, driveID, itemID, itemName, addtl)
			}

			errs.AddSkip(ctx, skip)
			logger.Ctx(ctx).Infow("malware detected", "item_details", addtl)

			continue
		}

		// Deleted file or folder.
		if item.GetDeleted() != nil {
			if err := c.handleDelete(
				itemID,
				driveID,
				oldPrevPaths,
				currPrevPaths,
				newPrevPaths,
				isFolder,
				excluded,
				invalidPrevDelta); err != nil {
				return nil, clues.Stack(err).WithClues(ictx)
			}

			continue
		}

		collectionPath, err := c.getCollectionPath(driveID, item)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err).
				WithClues(ictx).
				Label(fault.LabelForceNoBackupCreation))

			continue
		}

		// Skip items that don't match the folder selectors we were given.
		if shouldSkip(ctx, collectionPath, c.handler, driveName) {
			logger.Ctx(ictx).Debugw("path not selected", "skipped_path", collectionPath.String())
			continue
		}

		switch {
		case isFolder:
			// Deletions are handled above so this is just moves/renames.
			var prevPath path.Path

			prevPathStr, ok := oldPrevPaths[itemID]
			if ok {
				prevPath, err = path.FromDataLayerPath(prevPathStr, false)
				if err != nil {
					el.AddRecoverable(ctx, clues.Wrap(err, "invalid previous path").
						WithClues(ictx).
						With("prev_path_string", prevPathStr))
				}
			} else if item.GetRoot() != nil {
				// Root doesn't move or get renamed.
				prevPath = collectionPath
			}

			// Moved folders don't cause delta results for any subfolders nested in
			// them. We need to go through and update paths to handle that. We only
			// update newPaths so we don't accidentally clobber previous deletes.
			updatePath(newPrevPaths, itemID, collectionPath.String())

			found, err := updateCollectionPaths(driveID, itemID, c.CollectionMap, collectionPath)
			if err != nil {
				return nil, clues.Stack(err).WithClues(ictx)
			}

			if found {
				continue
			}

			colScope := CollectionScopeFolder
			if item.GetPackageEscaped() != nil {
				colScope = CollectionScopePackage
			}

			col, err := NewCollection(
				c.handler,
				collectionPath,
				prevPath,
				driveID,
				c.statusUpdater,
				c.ctrl,
				colScope,
				invalidPrevDelta,
				nil)
			if err != nil {
				return nil, clues.Stack(err).WithClues(ictx)
			}

			col.driveName = driveName

			c.CollectionMap[driveID][itemID] = col
			c.NumContainers++

			if item.GetRoot() != nil {
				continue
			}

			// Add an entry to fetch permissions into this collection. This assumes
			// that OneDrive always returns all folders on the path of an item
			// before the item. This seems to hold true for now at least.
			if col.Add(item) {
				c.NumItems++
			}

		case item.GetFile() != nil:
			// Deletions are handled above so this is just moves/renames.
			if len(ptr.Val(item.GetParentReference().GetId())) == 0 {
				return nil, clues.New("file without parent ID").WithClues(ictx)
			}

			// Get the collection for this item.
			parentID := ptr.Val(item.GetParentReference().GetId())
			ictx = clues.Add(ictx, "parent_id", parentID)

			collection, ok := c.CollectionMap[driveID][parentID]
			if !ok {
				return nil, clues.New("item seen before parent folder").WithClues(ictx)
			}

			// This will only kick in if the file was moved multiple times
			// within a single delta query.  We delete the file from the previous
			// collection so that it doesn't appear in two places.
			prevParentContainerID, ok := currPrevPaths[itemID]
			if ok {
				prevColl, found := c.CollectionMap[driveID][prevParentContainerID]
				if !found {
					return nil, clues.New("previous collection not found").
						With("prev_parent_container_id", prevParentContainerID).
						WithClues(ictx)
				}

				if ok := prevColl.Remove(itemID); !ok {
					return nil, clues.New("removing item from prev collection").
						With("prev_parent_container_id", prevParentContainerID).
						WithClues(ictx)
				}
			}

			currPrevPaths[itemID] = parentID

			if collection.Add(item) {
				c.NumItems++
				c.NumFiles++
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
				excluded[itemID+metadata.DataFileSuffix] = struct{}{}
				excluded[itemID+metadata.MetaFileSuffix] = struct{}{}
			}

		default:
			el.AddRecoverable(ictx, clues.New("item is neither folder nor file").
				WithClues(ictx).
				Label(fault.LabelForceNoBackupCreation))
		}
	}

	return newPrevPaths, el.Failure()
}

type dirScopeChecker interface {
	IsAllPass() bool
	IncludesDir(dir string) bool
}

func shouldSkip(ctx context.Context, drivePath path.Path, dsc dirScopeChecker, driveName string) bool {
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
	oldPath := paths[id]
	if len(oldPath) == 0 {
		paths[id] = newPath
		return
	}

	if oldPath == newPath {
		return
	}

	// We need to do a prefix search on the rest of the map to update the subtree.
	// We don't need to make collections for all of these, as hierarchy merging in
	// other components should take care of that. We do need to ensure that the
	// resulting map contains all folders though so we know the next time around.
	for folderID, p := range paths {
		if !strings.HasPrefix(p, oldPath) {
			continue
		}

		paths[folderID] = strings.Replace(p, oldPath, newPath, 1)
	}
}
