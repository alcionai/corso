// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO: This number needs to be tuned
	// Consider max open file limit `ulimit -n`, usually 1024 when setting this value
	collectionChannelBufferSize = 50

	// TODO: Tune this later along with collectionChannelBufferSize
	urlPrefetchChannelBufferSize = 25

	// Max number of retries to get doc from M365
	// Seems to timeout at times because of multiple requests
	maxRetries = 4 // 1 + 3 retries
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
) (itemInfo *details.OneDriveInfo, itemData io.ReadCloser, err error)

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
		byteCount int64
		itemsRead int64 = 0
	)

	// Retrieve the OneDrive folder path to set later in
	// `details.OneDriveInfo`
	parentPathString, err := getDriveFolderPath(oc.folderPath)
	if err != nil {
		oc.reportAsCompleted(ctx, 0, 0, err)
		return
	}

	folderProgress, colCloser := observe.ProgressWithCount(
		observe.ItemQueueMsg,
		"/"+parentPathString,
		int64(len(oc.driveItemIDs)),
	)
	defer colCloser()
	defer close(folderProgress)

	limitCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(limitCh)
	var wg sync.WaitGroup

	type uerr struct {
		itemID string
		err    error
	}
	errCh := make(chan uerr)
	defer close(errCh)

	ffail := int64(0) // handling fast-fail
	for _, itemID := range oc.driveItemIDs {
		if ffail > 0 {
			break
		}
		limitCh <- struct{}{}

		wg.Add(1)
		go func(itemID string) {
			defer func() { <-limitCh }()
			defer wg.Done()

			// Read the item
			var (
				itemInfo *details.OneDriveInfo
				itemData io.ReadCloser
				err      error
			)

			// Retrying as we were hitting timeouts when we have multiple requests
			// https://github.com/microsoftgraph/msgraph-sdk-go/issues/302
			for i := 0; i < maxRetries; i++ {
				itemInfo, itemData, err = oc.itemReader(ctx, oc.service, oc.driveID, itemID)
				if err == nil {
					break
				}
				// TODO: Tweak sleep times
				time.Sleep(time.Duration(3*(i+1)) * time.Second)
			}

			if err != nil {
				errCh <- uerr{itemID, err}

				if oc.service.ErrPolicy() {
					atomic.AddInt64(&ffail, 1)
					return
				}

				return
			}

			// Item read successfully, add to collection
			atomic.AddInt64(&itemsRead, 1)
			// byteCount iteration
			atomic.AddInt64(&byteCount, itemInfo.Size)

			itemInfo.ParentPath = parentPathString
			progReader, closer := observe.ItemProgress(itemData, observe.ItemBackupMsg, itemInfo.ItemName, itemInfo.Size)
			go closer()

			oc.data <- &Item{
				id:   itemInfo.ItemName,
				data: progReader,
				info: itemInfo,
			}
			folderProgress <- struct{}{}
		}(itemID)
	}
	wg.Wait()

	for i := 0; i < len(errCh); i++ {
		e, ok := <-errCh
		if !ok {
			break
		}
		errs = support.WrapAndAppendf(e.itemID, e.err, errs)
	}

	oc.reportAsCompleted(ctx, int(itemsRead), byteCount, errs)
}

func (oc *Collection) reportAsCompleted(ctx context.Context, itemsRead int, byteCount int64, errs error) {
	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		1, // num folders (always 1)
		support.CollectionMetrics{
			Objects:    len(oc.driveItemIDs), // items to read,
			Successes:  itemsRead,            // items read successfully,
			TotalBytes: byteCount,            // Number of bytes read in the operation,
		},
		errs,
		oc.folderPath.Folder(), // Additional details
	)
	logger.Ctx(ctx).Debug(status.String())
	oc.statusUpdater(status)
}
