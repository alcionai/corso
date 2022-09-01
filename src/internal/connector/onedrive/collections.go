package onedrive

import (
	"context"
	"path"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
)

// Collections is used to retrieve OneDrive data for a
// specified user
type Collections struct {
	user string
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
	user string,
	service graph.Service,
	statusUpdater support.StatusUpdater,
) *Collections {
	return &Collections{
		user:          user,
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
		collectionPath := *item.GetParentReference().GetPath()
		if _, found := c.collectionMap[collectionPath]; !found {
			c.collectionMap[collectionPath] = NewCollection(collectionPath, driveID, c.service, c.statusUpdater)
		}
		switch {
		case item.GetFolder() != nil, item.GetPackage() != nil:
			// For folders and packages we also create a collection to represent those
			// TODO: This is where we might create a "special file" to represent these in the backup repository
			// e.g. a ".folderMetadataFile"
			itemPath := path.Join(*item.GetParentReference().GetPath(), *item.GetName())
			if _, found := c.collectionMap[itemPath]; !found {
				c.collectionMap[itemPath] = NewCollection(itemPath, driveID, c.service, c.statusUpdater)
			}
		case item.GetFile() != nil:
			collection := c.collectionMap[collectionPath].(*Collection)
			collection.Add(*item.GetId())
		default:
			return errors.Errorf("item type not supported. item name : %s", *item.GetName())
		}
	}
	return nil
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
