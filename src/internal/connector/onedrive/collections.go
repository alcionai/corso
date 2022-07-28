package onedrive

import (
	"context"
	"fmt"
	"path"

	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/delta"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/logger"
)

type Collections struct {
	user string
	// collectionMap allows lookup of the data.Collection
	// for a OneDrive folder
	collectionMap map[string]data.Collection
	service       graph.Service

	// Track stats from drive enumeration
	totalItems    int
	totalDirs     int
	totalEntries  int
	totalPackages int
}

func NewCollections(user string, service graph.Service) *Collections {
	return &Collections{
		user:          user,
		collectionMap: map[string]data.Collection{},
		service:       service,
	}
}

func (c *Collections) Get(ctx context.Context) ([]data.Collection, error) {

	// For the specified user
	// 1) Enumerate Drives
	// 2) Enumerate Folders -> Create a DataCollection for each
	// 3) For each DataCollection -> Initialize it with files to pull

	err := c.enumerateDrives(ctx)
	if err != nil {
		return nil, err
	}

	collections := make([]data.Collection, 0, len(c.collectionMap))
	for _, c := range c.collectionMap {
		collections = append(collections, c)
	}

	return collections, nil
}

// Enumerates drives and the items within them for a specific user
func (c *Collections) enumerateDrives(ctx context.Context) error {
	logger.Ctx(ctx).Debug("Enumerating drives")
	// Get Drives
	r, err := c.service.Client().UsersById(c.user).Drives().Get()
	if err != nil {
		return errors.Wrapf(err, "failed to retrieve user drives. user: %s, details: %s", c.user, support.ConnectorStackErrorTrace(err))
	}
	logger.Ctx(ctx).Debugf("Found %d drives", len(r.GetValue()))

	for _, d := range r.GetValue() {
		logger.Ctx(ctx).Debugf("Found Drive: ID: %s Type: %s", *d.GetId(), *d.GetDriveType())
		err := c.enumerateDrive(ctx, *d.GetId())
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Collections) enumerateDrive(ctx context.Context, driveID string) error {
	// TODO: Specify a timestamp in the delta query
	// https://docs.microsoft.com/en-us/graph/api/driveitem-delta?view=graph-rest-1.0&tabs=http#example-4-retrieving-delta-results-using-a-timestamp
	builder := c.service.Client().DrivesById(driveID).Root().Delta()
	for {
		r, err := builder.Get()
		if err != nil {
			return errors.Wrapf(err, "failed to query drive items. details: %s", support.ConnectorStackErrorTrace(err))
		}
		items := len(r.GetValue())
		logger.Ctx(ctx).Debugf("Found %d items", items)
		c.totalItems += len(r.GetValue())

		c.updateCollections(r.GetValue())

		if _, found := r.GetAdditionalData()["@odata.nextLink"]; !found {
			logger.Ctx(ctx).Debugf("Done enumerating")
			break
		}
		nextLink := r.GetAdditionalData()["@odata.nextLink"].(*string)
		logger.Ctx(ctx).Debugf("Found %s nextLink", *nextLink)
		builder = delta.NewDeltaRequestBuilder(*nextLink, c.service.Adapter())
	}
	logger.Ctx(ctx).Debugf("Found %d total, %d directories, %d entries,  %d packages, %d collections", c.totalItems, c.totalDirs, c.totalEntries, c.totalPackages, len(c.collectionMap))

	return nil
}

func (c *Collections) updateCollections(items []models.DriveItemable) {
	for _, item := range items {
		c.stats(item)
		if item.GetParentReference() == nil || item.GetParentReference().GetPath() == nil {
			continue
		}
		// TODO: Validate paths
		switch {
		case item.GetFolder() != nil, item.GetPackage() != nil:
			itemPath := path.Join(*item.GetParentReference().GetPath(), *item.GetName())
			if _, found := c.collectionMap[itemPath]; !found {
				c.collectionMap[itemPath] = NewCollection(itemPath)
			}
		case item.GetFile() != nil:
			collectionPath := *item.GetParentReference().GetPath()
			if _, found := c.collectionMap[collectionPath]; !found {
				c.collectionMap[collectionPath] = NewCollection(collectionPath)
			}
			collection := c.collectionMap[collectionPath].(*Collection)
			collection.driveItems = append(collection.driveItems, *item.GetId())
		default:
			// TODO: Handle this
			fmt.Printf("%s\n", *item.GetName())
			panic("should not get here")
		}
	}
}

func (c *Collections) stats(item models.DriveItemable) {
	switch {
	case item.GetFolder() != nil:
		c.totalDirs++
	case item.GetPackage() != nil:
		c.totalPackages++
	case item.GetFile() != nil:
		c.totalEntries++
	default:
		// TODO: Handle this
		fmt.Printf("%s\n", *item.GetName())
		panic("should not get here")
	}
}
