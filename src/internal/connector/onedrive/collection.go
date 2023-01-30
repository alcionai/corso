// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-lazy/pkg/lazy"

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

	MetaFileSuffix    = ".meta"
	DirMetaFileSuffix = ".dirmeta"
	DataFileSuffix    = ".data"
)

var (
	_ data.Collection    = &Collection{}
	_ data.Stream        = &Item{}
	_ data.StreamInfo    = &Item{}
	_ data.StreamModTime = &Item{}
)

// Collection represents a set of OneDrive objects retrieved from M365
type Collection struct {
	// configured to handle large item downloads
	itemClient *http.Client

	// data is used to share data streams with the collection consumer
	data chan data.Stream
	// folderPath indicates what level in the hierarchy this collection
	// represents
	folderPath path.Path
	// M365 IDs of file items within this collection
	driveItems map[string]models.DriveItemable
	// M365 ID of the drive this collection was created from
	driveID        string
	source         driveSource
	service        graph.Servicer
	statusUpdater  support.StatusUpdater
	itemReader     itemReaderFunc
	itemMetaReader itemMetaReaderFunc
	ctrl           control.Options

	// should only be true if the old delta token expired
	doNotMergeItems bool
}

// itemReadFunc returns a reader for the specified item
type itemReaderFunc func(
	hc *http.Client,
	item models.DriveItemable,
) (itemInfo details.ItemInfo, itemData io.ReadCloser, err error)

// itemMetaReaderFunc returns a reader for the metadata of the
// specified item
type itemMetaReaderFunc func(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	item models.DriveItemable,
) (io.ReadCloser, int, error)

// NewCollection creates a Collection
func NewCollection(
	itemClient *http.Client,
	folderPath path.Path,
	driveID string,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	source driveSource,
	ctrlOpts control.Options,
) *Collection {
	c := &Collection{
		itemClient:    itemClient,
		folderPath:    folderPath,
		driveItems:    map[string]models.DriveItemable{},
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
		c.itemMetaReader = oneDriveItemMetaReader
	}

	return c
}

// Adds an itemID to the collection
// This will make it eligible to be populated
func (oc *Collection) Add(item models.DriveItemable) {
	oc.driveItems[*item.GetId()] = item
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
type Metadata struct {
	Permissions []UserPermission `json:"permissions,omitempty"`
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id   string
	data io.ReadCloser
	info details.ItemInfo

	// true if the item was marked by graph as deleted.
	deleted bool
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return od.data
}

// TODO(ashmrtn): Fill in once delta tokens return deleted items.
func (od Item) Deleted() bool {
	return od.deleted
}

func (od *Item) Info() details.ItemInfo {
	return od.info
}

func (od *Item) ModTime() time.Time {
	return od.info.Modified()
}

func newLazyReader(
	ctx context.Context,
	oc *Collection,
	itemName string,
	item models.DriveItemable,
	errUpdater func(id string, err error),
) io.ReadCloser {
	var (
		itemID   = *item.GetId()
		itemSize = *item.GetSize()
	)

	return lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		// Read the item
		var (
			itemData io.ReadCloser
			err      error
		)

		for i := 1; i <= maxRetries; i++ {
			_, itemData, err = oc.itemReader(oc.itemClient, item)
			if err == nil {
				break
			}

			if graph.IsErrUnauthorized(err) {
				// assume unauthorized requests are a sign of an expired
				// jwt token, and that we've overrun the available window
				// to download the actual file.  Re-downloading the item
				// will refresh that download url.
				di, diErr := getDriveItem(ctx, oc.service, oc.driveID, itemID)
				if diErr != nil {
					err = errors.Wrap(diErr, "retrieving expired item")
					break
				}

				item = di

				continue

			} else if !graph.IsErrTimeout(err) && !graph.IsErrThrottled(err) && !graph.IsSericeUnavailable(err) {
				// TODO: graphAPI will provides headers that state the duration to wait
				// in order to succeed again.  The one second sleep won't cut it here.
				//
				// for all non-timeout, non-unauth, non-throttling errors, do not retry
				break
			}

			if i < maxRetries {
				time.Sleep(1 * time.Second)
			}
		}

		// check for errors following retries
		if err != nil {
			errUpdater(itemID, err)
			return nil, err
		}

		// display/log the item download
		progReader, closer := observe.ItemProgress(ctx, itemData, observe.ItemBackupMsg, observe.PII(itemName), itemSize)
		go closer()

		return progReader, nil
	})
}

// populateItems iterates through items added to the collection
// and uses the collection `itemReader` to read the item
func (oc *Collection) populateItems(ctx context.Context) {
	var (
		errs       error
		byteCount  int64
		itemsRead  int64
		dirsRead   int64
		itemsFound int64
		dirsFound  int64
		wg         sync.WaitGroup
		m          sync.Mutex
	)

	// Retrieve the OneDrive folder path to set later in
	// `details.OneDriveInfo`
	parentPathString, err := path.GetDriveFolderPath(oc.folderPath)
	if err != nil {
		oc.reportAsCompleted(ctx, 0, 0, 0, err)
		return
	}

	folderProgress, colCloser := observe.ProgressWithCount(
		ctx,
		observe.ItemQueueMsg,
		observe.PII("/"+parentPathString),
		int64(len(oc.driveItems)))
	defer colCloser()
	defer close(folderProgress)

	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	errUpdater := func(id string, err error) {
		m.Lock()
		errs = support.WrapAndAppend(id, err, errs)
		m.Unlock()
	}

	for id, item := range oc.driveItems {
		if oc.ctrl.FailFast && errs != nil {
			break
		}

		if item == nil {
			errUpdater(id, errors.New("nil item"))
			continue
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(item models.DriveItemable) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			var (
				itemName   = *item.GetName()
				itemSize   = *item.GetSize()
				isFile     = item.GetFile() != nil
				itemInfo   details.ItemInfo
				metaSuffix string
			)

			if isFile {
				atomic.AddInt64(&itemsFound, 1)
				metaSuffix = MetaFileSuffix
			} else {
				atomic.AddInt64(&dirsFound, 1)
				metaSuffix = DirMetaFileSuffix
			}

			switch oc.source {
			case SharePointSource:
				itemInfo.SharePoint = sharePointItemInfo(item, itemSize)
				itemInfo.SharePoint.ParentPath = parentPathString
			default:
				itemInfo.OneDrive = oneDriveItemInfo(item, itemSize)
				itemInfo.OneDrive.ParentPath = parentPathString
			}

			// Construct a new lazy readCloser to feed to the collection consumer.
			// This ensures that downloads won't be attempted unless that consumer
			// attempts to read bytes.  Assumption is that kopia will check things
			// like file modtimes before attempting to read.
			itemReader := newLazyReader(ctx, oc, itemName+DataFileSuffix, item, errUpdater)
			fmt.Println(metaSuffix)

			// This can cause inaccurate counts.  Right now it counts all the items
			// we intend to read.  Errors within the lazy readCloser will create a
			// conflict: an item is both successful and erroneous.  But the async
			// control to fix that is more error-prone than helpful.
			//
			// TODO: transform this into a stats bus so that async control of stats
			// aggregation is handled at the backup level, not at the item iteration
			// level.
			//
			// Item read successfully, add to collection
			if isFile {
				atomic.AddInt64(&itemsRead, 1)
			} else {
				atomic.AddInt64(&dirsRead, 1)
			}
			// byteCount iteration
			atomic.AddInt64(&byteCount, itemSize)

			oc.data <- &Item{
				id:   itemName,
				data: itemReader,
				info: itemInfo,
			}
			folderProgress <- struct{}{}
		}(item)
	}

	wg.Wait()

	oc.reportAsCompleted(ctx, int(itemsFound), int(itemsRead), byteCount, errs)
}

func (oc *Collection) reportAsCompleted(ctx context.Context, itemsFound, itemsRead int, byteCount int64, errs error) {
	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		1, // num folders (always 1)
		support.CollectionMetrics{
			Objects:    itemsFound, // items to read,
			Successes:  itemsRead,  // items read successfully,
			TotalBytes: byteCount,  // Number of bytes read in the operation,
		},
		errs,
		oc.folderPath.Folder(), // Additional details
	)
	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())
	oc.statusUpdater(status)
}
