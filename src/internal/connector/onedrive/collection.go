// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO: This number needs to be tuned
	// Consider max open file limit `ulimit -n`, usually 1024 when setting this value
	collectionChannelBufferSize = 5

	// TODO: Tune this later along with collectionChannelBufferSize
	urlPrefetchChannelBufferSize = 5

	// Max number of retries to get doc from M365
	// Seems to timeout at times because of multiple requests
	maxRetries = 4 // 1 + 3 retries
)

var (
	_ data.Collection = &Collection{}
	_ data.Stream     = &Item{}
	_ data.StreamInfo = &Item{}
	// TODO(ashmrtn): Uncomment when #1702 is resolved.
	//_ data.StreamModTime = &Item{}
)

// Collection represents a set of OneDrive objects retrieved from M365
type Collection struct {
	// data is used to share data streams with the collection consumer
	data chan data.Stream
	// folderPath indicates what level in the hierarchy this collection
	// represents
	folderPath path.Path
	// M365 IDs of file items within this collection
	driveItems []models.DriveItemable
	// M365 ID of the drive this collection was created from
	driveID       string
	source        DriveSource
	service       graph.Servicer
	statusUpdater support.StatusUpdater
	itemReader    itemReaderFunc
	ctrl          control.Options

	// should only be true if the old delta token expired
	doNotMergeItems bool
}

// itemReadFunc returns a reader for the specified item
type itemReaderFunc func(
	ctx context.Context,
	item models.DriveItemable,
) (itemInfo details.ItemInfo, itemData io.ReadCloser, err error)

// NewCollection creates a Collection
func NewCollection(
	folderPath path.Path,
	driveID string,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	source DriveSource,
	ctrlOpts control.Options,
) *Collection {
	c := &Collection{
		folderPath:    folderPath,
		driveID:       driveID,
		source:        source,
		service:       service,
		data:          make(chan data.Stream, collectionChannelBufferSize),
		statusUpdater: statusUpdater,
		ctrl:          ctrlOpts,
	}

	// Allows tests to set a mock populator
	switch source {
	case SharePointSource:
		c.itemReader = sharePointItemReader
	default:
		c.itemReader = oneDriveItemReader
	}

	return c
}

// Adds an itemID to the collection
// This will make it eligible to be populated
func (oc *Collection) Add(item models.DriveItemable) {
	oc.driveItems = append(oc.driveItems, item)
}

// Items() returns the channel containing M365 Exchange objects
func (oc *Collection) Items() <-chan data.Stream {
	go oc.populateItems(context.Background())
	return oc.data
}

func (oc *Collection) FullPath() path.Path {
	return oc.folderPath
}

// TODO(ashmrtn): Fill in with previous path once GraphConnector compares old
// and new folder hierarchies.
func (oc Collection) PreviousPath() path.Path {
	return nil
}

// TODO(ashmrtn): Fill in once GraphConnector compares old and new folder
// hierarchies.
func (oc Collection) State() data.CollectionState {
	return data.NewState
}

func (oc Collection) DoNotMergeItems() bool {
	return oc.doNotMergeItems
}

// FilePermission is used to store permissions of a specific user to a
// OneDrive item.
type UserPermission struct {
	ID         string     `json:"id,omitempty"`
	Roles      []string   `json:"role,omitempty"`
	Email      string     `json:"email,omitempty"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

// ItemMeta contains metadata about the Item. It gets stored in a
// separate file in kopia
type ItemMeta struct {
	Permissions []UserPermission `json:"permissions,omitempty"`
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id   string
	data io.ReadCloser
	info details.ItemInfo
	meta ItemMeta

	// true if the item was marked by graph as deleted.
	deleted bool
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return od.data
}

func (od *Item) ToMetaReader() (io.ReadCloser, error) {
	c, err := json.Marshal(od.meta)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(c)), nil
}

// TODO(ashmrtn): Fill in once delta tokens return deleted items.
func (od Item) Deleted() bool {
	return od.deleted
}

func (od *Item) Info() details.ItemInfo {
	return od.info
}

// TODO(ashmrtn): Uncomment when #1702 is resolved.
//func (od *Item) ModTime() time.Time {
//	return od.info.Modified
//}

// populateItems iterates through items added to the collection
// and uses the collection `itemReader` to read the item
func (oc *Collection) populateItems(ctx context.Context) {
	var (
		errs      error
		byteCount int64
		itemsRead int64
		wg        sync.WaitGroup
		m         sync.Mutex
	)

	// Retrieve the OneDrive folder path to set later in
	// `details.OneDriveInfo`
	parentPathString, err := path.GetDriveFolderPath(oc.folderPath)
	if err != nil {
		oc.reportAsCompleted(ctx, 0, 0, err)
		return
	}

	folderProgress, colCloser := observe.ProgressWithCount(
		observe.ItemQueueMsg,
		"/"+parentPathString,
		int64(len(oc.driveItems)),
	)
	defer colCloser()
	defer close(folderProgress)

	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	errUpdater := func(id string, err error) {
		m.Lock()
		errs = support.WrapAndAppend(id, err, errs)
		m.Unlock()
	}

	for _, item := range oc.driveItems {
		if oc.ctrl.FailFast && errs != nil {
			break
		}

		// Don't process folders for non-onedrive(sharepoint).
		isFile := item.GetFile() != nil
		if !isFile && oc.source != OneDriveSource {
			continue
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(item models.DriveItemable) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			// Read the item
			var (
				itemInfo details.ItemInfo
				itemData io.ReadCloser
				itemMeta ItemMeta
				err      error
			)

			isFile := item.GetFile() != nil
			for i := 1; i <= maxRetries; i++ {
				if isFile {
					itemInfo, itemData, err = oc.itemReader(ctx, item)
				} else {
					itemInfo = details.ItemInfo{OneDrive: oneDriveItemInfo(item, *item.GetSize())}
					itemData = nil
					err = nil
				}

				if err == nil {
					itemMeta, err = oneDriveItemMetaInfo(ctx, oc.driveID, item, oc.service)
				}

				// retry on Timeout type errors, break otherwise.
				if err == nil || graph.IsErrTimeout(err) == nil {
					break
				}

				if i < maxRetries {
					time.Sleep(1 * time.Second)
				}
			}

			if err != nil {
				errUpdater(*item.GetId(), err)
				return
			}

			var (
				itemName string
				itemSize int64
			)

			switch oc.source {
			case SharePointSource:
				itemInfo.SharePoint.ParentPath = parentPathString
				itemName = itemInfo.SharePoint.ItemName
				itemSize = itemInfo.SharePoint.Size
			default:
				itemInfo.OneDrive.ParentPath = parentPathString
				itemName = itemInfo.OneDrive.ItemName
				itemSize = itemInfo.OneDrive.Size
			}

			progReader, closer := observe.ItemProgress(itemData, observe.ItemBackupMsg, itemName, itemSize)
			go closer()

			// Item read successfully, add to collection
			atomic.AddInt64(&itemsRead, 1)
			// byteCount iteration
			atomic.AddInt64(&byteCount, itemSize)

			oc.data <- &Item{
				id:   itemName,
				data: progReader,
				info: itemInfo,
				meta: itemMeta,
			}
			folderProgress <- struct{}{}
		}(item)
	}

	wg.Wait()

	oc.reportAsCompleted(ctx, int(itemsRead), byteCount, errs)
}

func (oc *Collection) reportAsCompleted(ctx context.Context, itemsRead int, byteCount int64, errs error) {
	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		1, // num folders (always 1)
		support.CollectionMetrics{
			Objects:    len(oc.driveItems), // items to read,
			Successes:  itemsRead,          // items read successfully,
			TotalBytes: byteCount,          // Number of bytes read in the operation,
		},
		errs,
		oc.folderPath.Folder(), // Additional details
	)
	logger.Ctx(ctx).Debug(status.String())
	oc.statusUpdater(status)
}
