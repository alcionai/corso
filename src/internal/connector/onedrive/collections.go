package onedrive

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

		delta, paths, err := collectItems(ctx, c.service, driveID, c.UpdateCollections)
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
// A new collection is created for every drive folder (or package)
func (c *Collections) UpdateCollections(
	ctx context.Context,
	driveID string,
	items []models.DriveItemable,
	paths map[string]string,
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
		if !includePath(ctx, c.matcher, collectionPath) {
			logger.Ctx(ctx).Infof("Skipping path %s", collectionPath.String())
			continue
		}

		switch {
		case item.GetFolder() != nil, item.GetPackage() != nil:
			// Eventually, deletions of folders will be handled here so we may as well
			// start off by saving the path.Path of the item instead of just the
			// OneDrive parentRef or such.
			folderPath, err := collectionPath.Append(*item.GetName(), false)
			if err != nil {
				logger.Ctx(ctx).Errorw("failed building collection path", "error", err)
				return err
			}

			// TODO(ashmrtn): Handle deletions by removing this entry from the map.
			// TODO(ashmrtn): Handle moves by setting the collection state if the
			// collection doesn't already exist/have that state.
			paths[*item.GetId()] = folderPath.String()

			if c.source != OneDriveSource {
				continue
			}
			fallthrough

		case item.GetFile() != nil:
			col, found := c.CollectionMap[collectionPath.String()]
			if !found {
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
