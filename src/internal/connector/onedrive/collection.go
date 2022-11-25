// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"io"
	"io/ioutil"
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

	// keep in max open file limit `ulimit -n`, usually 1024
	// https://superuser.com/questions/1356320/what-is-the-number-of-open-files-limits
	// M365 seems to throttle with 100
	const maxConcurrency = 10
	const maxRetries = 3
	limitCh := make(chan struct{}, maxConcurrency)
	defer close(limitCh)
	var wg sync.WaitGroup
	wg.Add(len(oc.driveItemIDs))

	ffail := false // handling fast-fail
	for _, itemID := range oc.driveItemIDs {
		if ffail {
			break
		}
		limitCh <- struct{}{}

		go func() {
			defer func() { <-limitCh }()
			defer wg.Done()

			// Read the item
			var (
				itemInfo *details.OneDriveInfo
				itemData io.ReadCloser
				err      error
			)
			// https://github.com/microsoftgraph/msgraph-sdk-go/issues/302
			// TODO(meain): should we be retrying here or inside the driveItemReader
			for i := 0; i < maxRetries; i++ {
				itemInfo, itemData, err = oc.itemReader(ctx, oc.service, oc.driveID, itemID)
				if err == nil {
					break
				}
				time.Sleep(time.Duration(5*(i+1)) * time.Second)
			}
			if err != nil {
				errs = support.WrapAndAppendf(itemID, err, errs)
				// TODO(meain): errs needs to be returned

				ioutil.WriteFile("/tmp/corso-err"+itemID, []byte(err.Error()), 0644)

				if oc.service.ErrPolicy() {
					ffail = true
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
		}()
	}
	wg.Wait()

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
