package onedrive

import (
	"context"
	stdpath "path"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// Collections is used to retrieve OneDrive data for a
// specified user
type Collections struct {
	tenant string
	user   string
	scope  selectors.OneDriveScope
	// collectionMap allows lookup of the data.Collection
	// for a OneDrive folder
	collectionMap map[string]data.Collection
	service       graph.Service
	statusUpdater support.StatusUpdater

	// Track stats from drive enumeration
	numItems    int
	numDirs     int
	numFiles    int
	numPackages int
}

func NewCollections(
	tenant string,
	user string,
	scope selectors.OneDriveScope,
	service graph.Service,
	statusUpdater support.StatusUpdater,
) *Collections {
	return &Collections{
		tenant:        tenant,
		user:          user,
		scope:         scope,
		collectionMap: map[string]data.Collection{},
		service:       service,
		statusUpdater: statusUpdater,
	}
}

// Retrieves OneDrive data as set of `data.Collections`
func (c *Collections) Get(ctx context.Context) ([]data.Collection, error) {
	// Enumerate drives for the specified user
	drives, err := drives(ctx, c.service, c.user)
	if err != nil {
		return nil, err
	}

	// Update the collection map with items from each drive
	for _, d := range drives {
		err = collectItems(ctx, c.service, *d.GetId(), c.updateCollections)
		if err != nil {
			return nil, err
		}
	}

	collections := make([]data.Collection, 0, len(c.collectionMap))
	for _, coll := range c.collectionMap {
		collections = append(collections, coll)
	}

	return collections, nil
}

func getCanonicalPath(p, tenant, user string) (path.Path, error) {
	pathBuilder := path.Builder{}.Append(strings.Split(p, "/")...)

	res, err := pathBuilder.ToDataLayerOneDrivePath(tenant, user, false)
	if err != nil {
		return nil, errors.Wrap(err, "converting to canonical path")
	}

	return res, nil
}

// Returns the path to the folder within the drive (i.e. under `root:`)
func getDriveFolderPath(p path.Path) (string, error) {
	drivePath, err := toOneDrivePath(p)
	if err != nil {
		return "", err
	}

	return path.Builder{}.Append(drivePath.folders...).String(), nil
}

// updateCollections initializes and adds the provided OneDrive items to Collections
// A new collection is created for every OneDrive folder (or package)
func (c *Collections) updateCollections(ctx context.Context, driveID string, items []models.DriveItemable) error {
	for _, item := range items {
		err := c.stats(item)
		if err != nil {
			return err
		}

		if item.GetRoot() != nil {
			// Skip the root item
			continue
		}

		if item.GetParentReference() == nil || item.GetParentReference().GetPath() == nil {
			return errors.Errorf("item does not have a parent reference. item name : %s", *item.GetName())
		}

		// Create a collection for the parent of this item
		collectionPath, err := getCanonicalPath(
			*item.GetParentReference().GetPath(),
			c.tenant,
			c.user,
		)
		if err != nil {
			return err
		}
		if _, found := c.collectionMap[collectionPath.String()]; !found {
			if !includePath(context.Background(), c.scope, collectionPath) {
				logger.Ctx(ctx).Infof("Skipping path %s", collectionPath.String())
				continue
			}

			c.collectionMap[collectionPath.String()] = NewCollection(
				collectionPath,
				driveID,
				c.service,
				c.statusUpdater,
			)
		}

		switch {
		case item.GetFolder() != nil, item.GetPackage() != nil:
			// For folders and packages we also create a collection to represent those
			// TODO: This is where we might create a "special file" to represent these in the backup repository
			// e.g. a ".folderMetadataFile"
			itemPath, err := getCanonicalPath(
				stdpath.Join(
					*item.GetParentReference().GetPath(),
					*item.GetName(),
				),
				c.tenant,
				c.user,
			)
			if err != nil {
				return err
			}

			if _, found := c.collectionMap[itemPath.String()]; !found {
				c.collectionMap[itemPath.String()] = NewCollection(
					itemPath,
					driveID,
					c.service,
					c.statusUpdater,
				)
			}
		case item.GetFile() != nil:
			if !includePath(ctx, c.scope, collectionPath) {
				continue
			}
			collection := c.collectionMap[collectionPath.String()].(*Collection)
			collection.Add(*item.GetId())
		default:
			return errors.Errorf("item type not supported. item name : %s", *item.GetName())
		}
	}

	return nil
}

func includePath(ctx context.Context, scope selectors.OneDriveScope, folderPath path.Path) bool {
	// Check if the folder is allowed by the scope.
	folderPathString, err := getDriveFolderPath(folderPath)
	if err != nil {
		logger.Ctx(ctx).Error(err)
		return true
	}

	// Hack for the edge case where we're looking at the root folder and can
	// select any folder. Right now the root folder has an empty folder path.
	if len(folderPathString) == 0 && scope.IsAny(selectors.OneDriveFolder) {
		return true
	}

	logger.Ctx(ctx).Infof("Checking path %q", folderPathString)

	return scope.Matches(selectors.OneDriveFolder, folderPathString)
}

func (c *Collections) stats(item models.DriveItemable) error {
	switch {
	case item.GetFolder() != nil:
		c.numDirs++
	case item.GetPackage() != nil:
		c.numPackages++
	case item.GetFile() != nil:
		c.numFiles++
	default:
		return errors.Errorf("item type not supported. item name : %s", *item.GetName())
	}

	c.numItems++

	return nil
}
