package onedrive

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

type driveSource int

const (
	unknownDriveSource driveSource = iota
	OneDriveSource
	SharePointSource
)
const restrictedDirectory = "Site Pages"

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
	tenant        string
	resourceOwner string
	source        driveSource
	matcher       folderMatcher
	service       graph.Servicer
	statusUpdater support.StatusUpdater

	ctrl control.Options

	// collectionMap allows lookup of the data.Collection
	// for a OneDrive folder
	CollectionMap map[string]data.Collection

	// Track stats from drive enumeration. Represents the items backed up.
	NumItems      int
	NumFiles      int
	NumContainers int
}

func NewCollections(
	tenant string,
	resourceOwner string,
	source driveSource,
	matcher folderMatcher,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collections {
	return &Collections{
		tenant:        tenant,
		resourceOwner: resourceOwner,
		source:        source,
		matcher:       matcher,
		CollectionMap: map[string]data.Collection{},
		service:       service,
		statusUpdater: statusUpdater,
		ctrl:          ctrlOpts,
	}
}

// Retrieves drive data as set of `data.Collections`
func (c *Collections) Get(ctx context.Context) ([]data.Collection, error) {
	// Enumerate drives for the specified resourceOwner
	drives, err := drives(ctx, c.service, c.resourceOwner, c.source)
	if err != nil {
		return nil, err
	}

	var (
		// Drive ID -> delta URL for drive
		deltaURLs = map[string]string{}
		// Drive ID -> folder ID -> folder path
		folderPaths = map[string]map[string]string{}
	)

	// Update the collection map with items from each drive
	for _, d := range drives {
		driveID := *d.GetId()
		driveName := *d.GetName()

		delta, paths, err := collectItems(ctx, c.service, driveID, driveName, c.UpdateCollections)
		if err != nil {
			return nil, err
		}

		if len(delta) > 0 {
			deltaURLs[driveID] = delta
		}

		if len(paths) > 0 {
			folderPaths[driveID] = map[string]string{}

			for id, p := range paths {
				folderPaths[driveID][id] = p
			}
		}
	}

	observe.Message(ctx, fmt.Sprintf("Discovered %d items to backup", c.NumItems))

	// Add an extra for the metadata collection.
	collections := make([]data.Collection, 0, len(c.CollectionMap)+1)
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
		c.statusUpdater,
	)

	if err != nil {
		// Technically it's safe to continue here because the logic for starting an
		// incremental backup should eventually find that the metadata files are
		// empty/missing and default to a full backup.
		logger.Ctx(ctx).Warnw(
			"making metadata collection for future incremental backups",
			"error",
			err,
		)
	} else {
		collections = append(collections, metadata)
	}

	return collections, nil
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
) error {
	for _, item := range items {
		if item.GetRoot() != nil {
			// Skip the root item
			continue
		}

		if item.GetParentReference() == nil || item.GetParentReference().GetPath() == nil {
			return errors.Errorf("item does not have a parent reference. item name : %s", *item.GetName())
		}

		// Create a collection for the parent of this item
		collectionPath, err := GetCanonicalPath(
			*item.GetParentReference().GetPath(),
			c.tenant,
			c.resourceOwner,
			c.source,
		)
		if err != nil {
			return err
		}

		// Skip items that don't match the folder selectors we were given.
		if shouldSkipDrive(ctx, collectionPath, c.matcher, driveName) {
			logger.Ctx(ctx).Infof("Skipping path %s", collectionPath.String())
			continue
		}

		switch {
		case item.GetFolder() != nil, item.GetPackage() != nil:
			if item.GetDeleted() != nil {
				// Nested folders also return deleted delta results so we don't have to
				// worry about doing a prefix search in the map to remove the subtree of
				// the deleted folder/package.
				delete(newPaths, *item.GetId())

				// TODO(ashmrtn): Create a collection with state Deleted.

				break
			}

			// Deletions of folders are handled in this case so we may as well start
			// off by saving the path.Path of the item instead of just the OneDrive
			// parentRef or such.
			folderPath, err := collectionPath.Append(*item.GetName(), false)
			if err != nil {
				logger.Ctx(ctx).Errorw("failed building collection path", "error", err)
				return err
			}

			// Moved folders don't cause delta results for any subfolders nested in
			// them. We need to go through and update paths to handle that. We only
			// update newPaths so we don't accidentally clobber previous deletes.
			//
			// TODO(ashmrtn): Since we're also getting notifications about folder
			// moves we may need to handle updates to a path of a collection we've
			// already created and partially populated.
			updatePath(newPaths, *item.GetId(), folderPath.String())

		case item.GetFile() != nil:
			col, found := c.CollectionMap[collectionPath.String()]
			if !found {
				// TODO(ashmrtn): Compare old and new path and set collection state
				// accordingly.
				col = NewCollection(
					collectionPath,
					driveID,
					c.service,
					c.statusUpdater,
					c.source,
					c.ctrl,
				)

				c.CollectionMap[collectionPath.String()] = col
				c.NumContainers++
				c.NumItems++
			}

			collection := col.(*Collection)
			collection.Add(item)
			c.NumFiles++
			c.NumItems++

		default:
			return errors.Errorf("item type not supported. item name : %s", *item.GetName())
		}
	}

	return nil
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
		return nil, errors.Errorf("unrecognized drive data source")
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
		logger.Ctx(ctx).Error(err)
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
