package onedrive

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
)

type driveSource int

const (
	unknownDriveSource driveSource = iota
	OneDriveSource
	SharePointSource
)

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

	// Update the collection map with items from each drive
	for _, d := range drives {
		err = collectItems(ctx, c.service, *d.GetId(), c.UpdateCollections)
		if err != nil {
			return nil, err
		}
	}

	observe.Message(ctx, fmt.Sprintf("Discovered %d items to backup", c.NumItems))

	collections := make([]data.Collection, 0, len(c.CollectionMap))
	for _, coll := range c.CollectionMap {
		collections = append(collections, coll)
	}

	return collections, nil
}

// UpdateCollections initializes and adds the provided drive items to Collections
// A new collection is created for every drive folder (or package)
func (c *Collections) UpdateCollections(ctx context.Context, driveID string, items []models.DriveItemable) error {
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
		case item.GetFolder() != nil, item.GetPackage() != nil, item.GetFile() != nil:
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
