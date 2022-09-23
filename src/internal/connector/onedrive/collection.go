// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"io"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO: This number needs to be tuned
	collectionChannelBufferSize = 50
)

var (
	_ data.Collection = &Collection{}
	_ data.Stream     = &Item{}
	_ data.StreamInfo = &Item{}
)

// Collection represents a set of OneDrive objects retreived from M365
type Collection struct {
	// data is used to share data streams with the collection consumer
	data chan data.Stream
	// folderPath indicates what level in the hierarchy this collection
	// represents
	folderPath path.Path
	// M365 IDs of file items within this collection
	driveItemIDs []string
	// M365 ID of the drive this collection was created from
	driveID       string
	service       graph.Service
	statusUpdater support.StatusUpdater
	itemReader    itemReaderFunc
}

// itemReadFunc returns a reader for the specified item
type itemReaderFunc func(
	ctx context.Context,
	service graph.Service,
	driveID, itemID string,
) (name string, itemData io.ReadCloser, err error)

// NewCollection creates a Collection
func NewCollection(
	folderPath path.Path,
	driveID string,
	service graph.Service,
	statusUpdater support.StatusUpdater,
) *Collection {
	c := &Collection{
		folderPath:    folderPath,
		driveItemIDs:  []string{},
		driveID:       driveID,
		service:       service,
		data:          make(chan data.Stream, collectionChannelBufferSize),
		statusUpdater: statusUpdater,
	}
	// Allows tests to set a mock populator
	c.itemReader = driveItemReader

	return c
}

// Adds an itemID to the collection
// This will make it eligible to be populated
func (oc *Collection) Add(itemID string) {
	oc.driveItemIDs = append(oc.driveItemIDs, itemID)
}

// Items() returns the channel containing M365 Exchange objects
func (oc *Collection) Items() <-chan data.Stream {
	go oc.populateItems(context.Background())
	return oc.data
}

func (oc *Collection) FullPath() path.Path {
	return oc.folderPath
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id   string
	data io.ReadCloser
	info *details.OneDriveInfo
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return od.data
}

func (od *Item) Info() details.ItemInfo {
	return details.ItemInfo{OneDrive: od.info}
}

// populateItems iterates through items added to the collection
// and uses the collection `itemReader` to read the item
func (oc *Collection) populateItems(ctx context.Context) {
	var (
		errs      error
		itemsRead = 0
	)

	for _, itemID := range oc.driveItemIDs {
		// Read the item
		itemName, itemData, err := oc.itemReader(ctx, oc.service, oc.driveID, itemID)
		if err != nil {
			errs = support.WrapAndAppendf(itemID, err, errs)

			if oc.service.ErrPolicy() {
				break
			}

			continue
		}
		// Item read successfully, add to collection
		itemsRead++
		oc.data <- &Item{
			id:   itemName,
			data: itemData,
			info: &details.OneDriveInfo{
				ItemType:   details.OneDriveItem,
				ItemName:   itemName,
				ParentPath: oc.folderPath.String(),
			},
		}
	}

	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		len(oc.driveItemIDs), // items to read
		itemsRead,            // items read successfully
		1,                    // num folders (always 1)
		errs)
	logger.Ctx(ctx).Debug(status.String())
	oc.statusUpdater(status)
}
