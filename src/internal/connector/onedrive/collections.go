package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/excludes"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

type driveSource int

const (
	unknownDriveSource driveSource = iota
	OneDriveSource
	SharePointSource
)

type collectionScope int

const (
	// CollectionScopeUnknown is used when we don't know and don't need
	// to know the kind, like in the case of deletes
	CollectionScopeUnknown collectionScope = iota

	// CollectionScopeFolder is used for regular folder collections
	CollectionScopeFolder

	// CollectionScopePackage is used to represent OneNote items
	CollectionScopePackage
)

const (
	restrictedDirectory = "Site Pages"
	rootDrivePattern    = "/drives/%s/root:"
)

func (ds driveSource) toPathServiceCat() (path.ServiceType, path.CategoryType) {
	switch ds {
	case OneDriveSource:
		return path.OneDriveService, path.FilesCategory
	case SharePointSource:
		return path.SharePointService, path.LibrariesCategory
	default:
		return path.UnknownService, path.UnknownCategory
	}
}

type folderMatcher interface {
	IsAny() bool
	Matches(string) bool
}

// Collections is used to retrieve drive data for a
// resource owner, which can be either a user or a sharepoint site.
type Collections struct {
	// configured to handle large item downloads
	itemClient graph.Requester

	tenant        string
	resourceOwner string
	source        driveSource
	matcher       folderMatcher
	service       graph.Servicer
	statusUpdater support.StatusUpdater

	ctrl control.Options

	// collectionMap allows lookup of the data.BackupCollection
	// for a OneDrive folder.
	// driveID -> itemID -> collection
	CollectionMap map[string]map[string]*Collection

	// Not the most ideal, but allows us to change the pager function for testing
	// as needed. This will allow us to mock out some scenarios during testing.
	drivePagerFunc func(
		source driveSource,
		servicer graph.Servicer,
		resourceOwner string,
		fields []string,
	) (api.DrivePager, error)
	itemPagerFunc func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager

	// Track stats from drive enumeration. Represents the items backed up.
	NumItems      int
	NumFiles      int
	NumContainers int
}

func NewCollections(
	itemClient graph.Requester,
	tenant string,
	resourceOwner string,
	source driveSource,
	matcher folderMatcher,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collections {
	return &Collections{
		itemClient:     itemClient,
		tenant:         tenant,
		resourceOwner:  resourceOwner,
		source:         source,
		matcher:        matcher,
		CollectionMap:  map[string]map[string]*Collection{},
		drivePagerFunc: PagerForSource,
		itemPagerFunc:  defaultItemPager,
		service:        service,
		statusUpdater:  statusUpdater,
		ctrl:           ctrlOpts,
	}
}

func deserializeMetadata(
	ctx context.Context,
	cols []data.RestoreCollection,
	errs *fault.Bus,
) (map[string]string, map[string]map[string]string, error) {
	logger.Ctx(ctx).Infow(
		"deserialzing previous backup metadata",
		"num_collections", len(cols))

	var (
		prevDeltas  = map[string]string{}
		prevFolders = map[string]map[string]string{}
		el          = errs.Local()
	)

	for _, col := range cols {
		if el.Failure() != nil {
			break
		}

		items := col.Items(ctx, errs)

		for breakLoop := false; !breakLoop; {
			select {
			case <-ctx.Done():
				return nil, nil, clues.Wrap(ctx.Err(), "deserialzing previous backup metadata").WithClues(ctx)

			case item, ok := <-items:
				if !ok {
					breakLoop = true
					break
				}

				var (
					err  error
					ictx = clues.Add(ctx, "item_uuid", item.UUID())
				)

				switch item.UUID() {
				case graph.PreviousPathFileName:
					err = deserializeMap(item.ToReader(), prevFolders)

				case graph.DeltaURLsFileName:
					err = deserializeMap(item.ToReader(), prevDeltas)

				default:
					logger.Ctx(ictx).Infow(
						"skipping unknown metadata file",
						"file_name", item.UUID())

					continue
				}

				if err == nil {
					// Successful decode.
					continue
				}

				// This is conservative, but report an error if any of the items for
				// any of the deserialized maps have duplicate drive IDs. This will
				// cause the entire backup to fail, but it's not clear if higher
				// layers would have caught this. Worst case if we don't handle this
				// we end up in a situation where we're sourcing items from the wrong
				// base in kopia wrapper.
				if errors.Is(err, errExistingMapping) {
					return nil, nil, clues.Wrap(err, "deserializing metadata file").WithClues(ictx)
				}

				err = clues.Stack(err).WithClues(ictx)

				el.AddRecoverable(err)
				logger.CtxErr(ictx, err).Error("deserializing base backup metadata")
			}
		}

		// Go through and remove partial results (i.e. path mapping but no delta URL
		// or vice-versa).
		for k, v := range prevDeltas {
			// Remove entries with an empty delta token as it's not useful.
			if len(v) == 0 {
				delete(prevDeltas, k)
				delete(prevFolders, k)
			}

			// Remove entries without a folders map as we can't tell kopia the
			// hierarchy changes.
			if _, ok := prevFolders[k]; !ok {
				delete(prevDeltas, k)
			}
		}

		for k := range prevFolders {
			if _, ok := prevDeltas[k]; !ok {
				delete(prevFolders, k)
			}
		}
	}

	return prevDeltas, prevFolders, el.Failure()
}

var errExistingMapping = clues.New("mapping already exists for same drive ID")

// deserializeMap takes an reader and a map of already deserialized items and
// adds the newly deserialized items to alreadyFound. Items are only added to
// alreadyFound if none of the keys in the freshly deserialized map already
// exist in alreadyFound. reader is closed at the end of this function.
func deserializeMap[T any](reader io.ReadCloser, alreadyFound map[string]T) error {
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
	epi *excludes.ParentsItems,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	prevDeltas, oldPathsByDriveID, err := deserializeMetadata(ctx, prevMetadata, errs)
	if err != nil {
		return nil, err
	}

	driveComplete, closer := observe.MessageWithCompletion(ctx, observe.Bulletf("files"))
	defer closer()
	defer close(driveComplete)

	// Enumerate drives for the specified resourceOwner
	pager, err := c.drivePagerFunc(c.source, c.service, c.resourceOwner, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	drives, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	if err != nil {
		return nil, err
	}

	var (
		// Drive ID -> delta URL for drive
		deltaURLs = map[string]string{}
		// Drive ID -> folder ID -> folder path
		folderPaths = map[string]map[string]string{}
	)

	for _, d := range drives {
		var (
			driveID     = ptr.Val(d.GetId())
			driveName   = ptr.Val(d.GetName())
			prevDelta   = prevDeltas[driveID]
			oldPaths    = oldPathsByDriveID[driveID]
			numOldDelta = 0
			ictx        = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)
		)

		if _, ok := c.CollectionMap[driveID]; !ok {
			c.CollectionMap[driveID] = map[string]*Collection{}
		}

		if len(prevDelta) > 0 {
			numOldDelta++
		}

		logger.Ctx(ictx).Infow(
			"previous metadata for drive",
			"num_paths_entries", len(oldPaths),
			"num_deltas_entries", numOldDelta)

		delta, paths, excluded, err := collectItems(
			ictx,
			c.itemPagerFunc(c.service, driveID, ""),
			driveID,
			driveName,
			c.UpdateCollections,
			oldPaths,
			prevDelta,
			errs)
		if err != nil {
			return nil, err
		}

		// Used for logging below.
		numDeltas := 0

		// It's alright to have an empty folders map (i.e. no folders found) but not
		// an empty delta token. This is because when deserializing the metadata we
		// remove entries for which there is no corresponding delta token/folder. If
		// we leave empty delta tokens then we may end up setting the State field
		// for collections when not actually getting delta results.
		if len(delta.URL) > 0 {
			deltaURLs[driveID] = delta.URL
			numDeltas++
		}

		// Avoid the edge case where there's no paths but we do have a valid delta
		// token. We can accomplish this by adding an empty paths map for this
		// drive. If we don't have this then the next backup won't use the delta
		// token because it thinks the folder paths weren't persisted.
		folderPaths[driveID] = map[string]string{}
		maps.Copy(folderPaths[driveID], paths)

		logger.Ctx(ictx).Infow(
			"persisted metadata for drive",
			"num_paths_entries", len(paths),
			"num_deltas_entries", numDeltas,
			"delta_reset", delta.Reset)

		// For both cases we don't need to do set difference on folder map if the
		// delta token was valid because we should see all the changes.
		if !delta.Reset && len(excluded) == 0 {
			continue
		} else if !delta.Reset {
			p, err := GetCanonicalPath(
				fmt.Sprintf(rootDrivePattern, driveID),
				c.tenant,
				c.resourceOwner,
				c.source)
			if err != nil {
				return nil, clues.Wrap(err, "making exclude prefix").WithClues(ictx)
			}

			epi.Add(p.String(), excluded)

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

		for fldID, p := range oldPaths {
			if _, ok := foundFolders[fldID]; ok {
				continue
			}

			delete(paths, fldID)

			prevPath, err := path.FromDataLayerPath(p, false)
			if err != nil {
				err = clues.Wrap(err, "invalid previous path").WithClues(ictx).With("deleted_path", p)
				return nil, err
			}

			col, err := NewCollection(
				c.itemClient,
				nil,
				prevPath,
				driveID,
				c.service,
				c.statusUpdater,
				c.source,
				c.ctrl,
				CollectionScopeUnknown,
				true)
			if err != nil {
				return nil, clues.Wrap(err, "making collection").WithClues(ictx)
			}

			c.CollectionMap[driveID][fldID] = col
		}
	}

	observe.Message(ctx, fmt.Sprintf("Discovered %d items to backup", c.NumItems))

	// Add an extra for the metadata collection.
	collections := []data.BackupCollection{}

	for _, driveColls := range c.CollectionMap {
		for _, coll := range driveColls {
			collections = append(collections, coll)
		}
	}

	service, category := c.source.toPathServiceCat()
	md, err := graph.MakeMetadataCollection(
		c.tenant,
		c.resourceOwner,
		service,
		category,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(graph.PreviousPathFileName, folderPaths),
			graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs),
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

	// TODO(ashmrtn): Track and return the set of items to exclude.
	return collections, nil
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
	oldPaths, newPaths map[string]string,
	isFolder bool,
	excluded map[string]struct{},
	itemCollection map[string]map[string]string,
	invalidPrevDelta bool,
) error {
	if !isFolder {
		// Try to remove the item from the Collection if an entry exists for this
		// item. This handles cases where an item was created and deleted during the
		// same delta query.
		if parentID, ok := itemCollection[driveID][itemID]; ok {
			if col := c.CollectionMap[driveID][parentID]; col != nil {
				col.Remove(itemID)
			}

			delete(itemCollection[driveID], itemID)
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

	prevPathStr, ok := oldPaths[itemID]
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
	delete(newPaths, itemID)

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
		c.itemClient,
		nil,
		prevPath,
		driveID,
		c.service,
		c.statusUpdater,
		c.source,
		c.ctrl,
		CollectionScopeUnknown,
		// DoNotMerge is not checked for deleted items.
		false)
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
		collectionPathStr string
		isRoot            = item.GetRoot() != nil
		isFile            = item.GetFile() != nil
	)

	if isRoot {
		collectionPathStr = fmt.Sprintf(rootDrivePattern, driveID)
	} else {
		if item.GetParentReference() == nil ||
			item.GetParentReference().GetPath() == nil {
			err := clues.New("no parent reference").
				With("item_name", ptr.Val(item.GetName()))

			return nil, err
		}

		collectionPathStr = ptr.Val(item.GetParentReference().GetPath())
	}

	collectionPath, err := GetCanonicalPath(
		collectionPathStr,
		c.tenant,
		c.resourceOwner,
		c.source,
	)
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

	collectionPath, err = collectionPath.Append(name, false)
	if err != nil {
		return nil, clues.Wrap(err, "making non-root folder path")
	}

	return collectionPath, nil
}

// UpdateCollections initializes and adds the provided drive items to Collections
// A new collection is created for every drive folder (or package).
// oldPaths is the unchanged data that was loaded from the metadata file.
// newPaths starts as a copy of oldPaths and is updated as changes are found in
// the returned results.
func (c *Collections) UpdateCollections(
	ctx context.Context,
	driveID, driveName string,
	items []models.DriveItemable,
	oldPaths map[string]string,
	newPaths map[string]string,
	excluded map[string]struct{},
	itemCollection map[string]map[string]string,
	invalidPrevDelta bool,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for _, item := range items {
		if el.Failure() != nil {
			break
		}

		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			ictx     = clues.Add(ctx, "item_id", itemID, "item_name", itemName)
			isFolder = item.GetFolder() != nil || item.GetPackage() != nil
		)

		if item.GetMalware() != nil {
			addtl := graph.ItemInfo(item)
			skip := fault.FileSkip(fault.SkipMalware, itemID, itemName, addtl)

			if isFolder {
				skip = fault.ContainerSkip(fault.SkipMalware, itemID, itemName, addtl)
			}

			errs.AddSkip(skip)
			logger.Ctx(ctx).Infow("malware detected", "item_details", addtl)

			continue
		}

		// Deleted file or folder.
		if item.GetDeleted() != nil {
			if err := c.handleDelete(
				itemID,
				driveID,
				oldPaths,
				newPaths,
				isFolder,
				excluded,
				itemCollection,
				invalidPrevDelta,
			); err != nil {
				return clues.Stack(err).WithClues(ictx)
			}

			continue
		}

		collectionPath, err := c.getCollectionPath(driveID, item)
		if err != nil {
			el.AddRecoverable(clues.Stack(err).
				WithClues(ictx).
				Label(fault.LabelForceNoBackupCreation))

			continue
		}

		// Skip items that don't match the folder selectors we were given.
		if shouldSkipDrive(ctx, collectionPath, c.matcher, driveName) {
			logger.Ctx(ictx).Debugw("Skipping drive path", "skipped_path", collectionPath.String())
			continue
		}

		switch {
		case isFolder:
			// Deletions are handled above so this is just moves/renames.
			var prevPath path.Path

			prevPathStr, ok := oldPaths[itemID]
			if ok {
				prevPath, err = path.FromDataLayerPath(prevPathStr, false)
				if err != nil {
					el.AddRecoverable(clues.Wrap(err, "invalid previous path").
						WithClues(ictx).
						With("path_string", prevPathStr))
				}
			} else if item.GetRoot() != nil {
				// Root doesn't move or get renamed.
				prevPath = collectionPath
			}

			// Moved folders don't cause delta results for any subfolders nested in
			// them. We need to go through and update paths to handle that. We only
			// update newPaths so we don't accidentally clobber previous deletes.
			updatePath(newPaths, itemID, collectionPath.String())

			found, err := updateCollectionPaths(driveID, itemID, c.CollectionMap, collectionPath)
			if err != nil {
				return clues.Stack(err).WithClues(ictx)
			}

			if found {
				continue
			}

			colScope := CollectionScopeFolder
			if item.GetPackage() != nil {
				colScope = CollectionScopePackage
			}

			col, err := NewCollection(
				c.itemClient,
				collectionPath,
				prevPath,
				driveID,
				c.service,
				c.statusUpdater,
				c.source,
				c.ctrl,
				colScope,
				invalidPrevDelta,
			)
			if err != nil {
				return clues.Stack(err).WithClues(ictx)
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
				return clues.New("file without parent ID").WithClues(ictx)
			}

			// Get the collection for this item.
			parentID := ptr.Val(item.GetParentReference().GetId())
			ictx = clues.Add(ictx, "parent_id", parentID)

			collection, found := c.CollectionMap[driveID][parentID]
			if !found {
				return clues.New("item seen before parent folder").WithClues(ictx)
			}

			// Delete the file from previous collection. This will
			// only kick in if the file was moved multiple times
			// within a single delta query
			icID, found := itemCollection[driveID][itemID]
			if found {
				pcollection, found := c.CollectionMap[driveID][icID]
				if !found {
					return clues.New("previous collection not found").WithClues(ictx)
				}

				removed := pcollection.Remove(itemID)
				if !removed {
					return clues.New("removing from prev collection").WithClues(ictx)
				}
			}

			itemCollection[driveID][itemID] = parentID

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
			return clues.New("item type not supported").WithClues(ictx)
		}
	}

	return el.Failure()
}

func shouldSkipDrive(ctx context.Context, drivePath path.Path, m folderMatcher, driveName string) bool {
	return !includePath(ctx, m, drivePath) ||
		(drivePath.Category() == path.LibrariesCategory && restrictedDirectory == driveName)
}

// GetCanonicalPath constructs the standard path for the given source.
func GetCanonicalPath(p, tenant, resourceOwner string, source driveSource) (path.Path, error) {
	var (
		pathBuilder = path.Builder{}.Append(strings.Split(p, "/")...)
		result      path.Path
		err         error
	)

	switch source {
	case OneDriveSource:
		result, err = pathBuilder.ToDataLayerOneDrivePath(tenant, resourceOwner, false)
	case SharePointSource:
		result, err = pathBuilder.ToDataLayerSharePointPath(tenant, resourceOwner, path.LibrariesCategory, false)
	default:
		return nil, clues.New("unrecognized data source")
	}

	if err != nil {
		return nil, clues.Wrap(err, "converting to canonical path")
	}

	return result, nil
}

func includePath(ctx context.Context, m folderMatcher, folderPath path.Path) bool {
	// Check if the folder is allowed by the scope.
	folderPathString, err := path.GetDriveFolderPath(folderPath)
	if err != nil {
		logger.Ctx(ctx).With("err", err).Error("getting drive folder path")
		return true
	}

	// Hack for the edge case where we're looking at the root folder and can
	// select any folder. Right now the root folder has an empty folder path.
	if len(folderPathString) == 0 && m.IsAny() {
		return true
	}

	return m.Matches(folderPathString)
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
