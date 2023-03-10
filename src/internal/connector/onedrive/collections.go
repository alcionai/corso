package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
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
	itemClient *http.Client

	tenant        string
	resourceOwner string
	source        driveSource
	matcher       folderMatcher
	service       graph.Servicer
	statusUpdater support.StatusUpdater

	ctrl control.Options

	// collectionMap allows lookup of the data.BackupCollection
	// for a OneDrive folder
	CollectionMap map[string]*Collection

	// Not the most ideal, but allows us to change the pager function for testing
	// as needed. This will allow us to mock out some scenarios during testing.
	drivePagerFunc func(
		source driveSource,
		servicer graph.Servicer,
		resourceOwner string,
		fields []string,
	) (drivePager, error)
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
	itemClient *http.Client,
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
		CollectionMap:  map[string]*Collection{},
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
				logger.Ctx(ictx).
					With("err", err).
					Errorw("deserializing base backup metadata", clues.InErr(err).Slice()...)
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
		return errors.Wrap(err, "deserializing file contents")
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
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	prevDeltas, oldPathsByDriveID, err := deserializeMetadata(ctx, prevMetadata, errs)
	if err != nil {
		return nil, nil, err
	}

	// Enumerate drives for the specified resourceOwner
	pager, err := c.drivePagerFunc(c.source, c.service, c.resourceOwner, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	retry := c.source == OneDriveSource

	drives, err := drives(ctx, pager, retry)
	if err != nil {
		return nil, nil, err
	}

	var (
		// Drive ID -> delta URL for drive
		deltaURLs = map[string]string{}
		// Drive ID -> folder ID -> folder path
		folderPaths = map[string]map[string]string{}
		// Items that should be excluded when sourcing data from the base backup.
		excludedItems = map[string]map[string]struct{}{}
	)

	for _, d := range drives {
		var (
			driveID     = ptr.Val(d.GetId())
			driveName   = ptr.Val(d.GetName())
			prevDelta   = prevDeltas[driveID]
			oldPaths    = oldPathsByDriveID[driveID]
			numOldDelta = 0
		)

		if len(prevDelta) > 0 {
			numOldDelta++
		}

		logger.Ctx(ctx).Infow(
			"previous metadata for drive",
			"num_paths_entries", len(oldPaths),
			"num_deltas_entries", numOldDelta)

		delta, paths, excluded, err := collectItems(
			ctx,
			c.itemPagerFunc(c.service, driveID, ""),
			driveID,
			driveName,
			c.UpdateCollections,
			oldPaths,
			prevDelta,
			errs)
		if err != nil {
			return nil, nil, err
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

		logger.Ctx(ctx).Infow(
			"persisted metadata for drive",
			"num_paths_entries",
			len(paths),
			"num_deltas_entries",
			numDeltas)

		if !delta.Reset {
			excludedItems[driveID] = excluded
			continue
		}

		// Set all folders in previous backup but not in the current
		// one with state deleted
		modifiedPaths := map[string]struct{}{}
		for _, p := range c.CollectionMap {
			modifiedPaths[p.FullPath().String()] = struct{}{}
		}

		for i, p := range oldPaths {
			_, found := paths[i]
			if found {
				continue
			}

			_, found = modifiedPaths[p]
			if found {
				// Original folder was deleted and new folder with the
				// same name/path was created in its place
				continue
			}

			delete(paths, i)

			prevPath, err := path.FromDataLayerPath(p, false)
			if err != nil {
				return nil, nil,
					clues.Wrap(err, "invalid previous path").WithClues(ctx).With("deleted_path", p)
			}

			col := NewCollection(
				c.itemClient,
				nil,
				prevPath,
				driveID,
				c.service,
				c.statusUpdater,
				c.source,
				c.ctrl,
				true,
			)
			c.CollectionMap[i] = col
		}
	}

	observe.Message(ctx, observe.Safe(fmt.Sprintf("Discovered %d items to backup", c.NumItems)))

	// Add an extra for the metadata collection.
	collections := make([]data.BackupCollection, 0, len(c.CollectionMap)+1)
	for _, coll := range c.CollectionMap {
		collections = append(collections, coll)
	}

	service, category := c.source.toPathServiceCat()
	metadata, err := graph.MakeMetadataCollection(
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
		logger.Ctx(ctx).
			With("err", err).
			Infow("making metadata collection for future incremental backups", clues.InErr(err).Slice()...)
	} else {
		collections = append(collections, metadata)
	}

	// TODO(ashmrtn): Track and return the set of items to exclude.
	return collections, excludedItems, nil
}

func updateCollectionPaths(
	id string,
	cmap map[string]*Collection,
	curPath path.Path,
) (bool, error) {
	var initialCurPath path.Path

	col, found := cmap[id]
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

	for i, c := range cmap {
		if i == id {
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
) error {
	if !isFolder {
		excluded[itemID+DataFileSuffix] = struct{}{}
		excluded[itemID+MetaFileSuffix] = struct{}{}
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
				With("collection_id", itemID, "path_string", prevPathStr)
		}
	}

	// Nested folders also return deleted delta results so we don't have to
	// worry about doing a prefix search in the map to remove the subtree of
	// the deleted folder/package.
	delete(newPaths, itemID)

	if prevPath == nil {
		// It is possible that an item was created and
		// deleted between two delta invocations. In
		// that case, it will only produce a single
		// delete entry in the delta response.
		return nil
	}

	col := NewCollection(
		c.itemClient,
		nil,
		prevPath,
		driveID,
		c.service,
		c.statusUpdater,
		c.source,
		c.ctrl,
		// DoNotMerge is not checked for deleted items.
		false,
	)

	c.CollectionMap[itemID] = col

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
	itemCollection map[string]string,
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
			addtl := graph.MalwareInfo(item)
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
		}

		// Skip items that don't match the folder selectors we were given.
		if shouldSkipDrive(ctx, collectionPath, c.matcher, driveName) {
			logger.Ctx(ictx).Infow("Skipping path", "skipped_path", collectionPath.String())
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

			found, err := updateCollectionPaths(itemID, c.CollectionMap, collectionPath)
			if err != nil {
				return clues.Stack(err).WithClues(ictx)
			}

			if found {
				continue
			}

			col := NewCollection(
				c.itemClient,
				collectionPath,
				prevPath,
				driveID,
				c.service,
				c.statusUpdater,
				c.source,
				c.ctrl,
				invalidPrevDelta,
			)
			c.CollectionMap[itemID] = col
			c.NumContainers++

			if c.source != OneDriveSource || item.GetRoot() != nil {
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
			collectionID := ptr.Val(item.GetParentReference().GetId())
			ictx = clues.Add(ictx, "collection_id", collectionID)

			collection, found := c.CollectionMap[collectionID]
			if !found {
				return clues.New("item seen before parent folder").WithClues(ictx)
			}

			// Delete the file from previous collection. This will
			// only kick in if the file was moved multiple times
			// within a single delta query
			itemColID, found := itemCollection[itemID]
			if found {
				pcollection, found := c.CollectionMap[itemColID]
				if !found {
					return clues.New("previous collection not found").WithClues(ictx)
				}

				removed := pcollection.Remove(item)
				if !removed {
					return clues.New("removing from prev collection").WithClues(ictx)
				}
			}

			itemCollection[itemID] = collectionID

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
				excluded[itemID+DataFileSuffix] = struct{}{}
				excluded[itemID+MetaFileSuffix] = struct{}{}
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
		return nil, errors.Wrap(err, "converting to canonical path")
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
