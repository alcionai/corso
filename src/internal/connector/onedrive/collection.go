// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
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
	driveID       string
	source        driveSource
	service       graph.Servicer
	statusUpdater support.StatusUpdater
	ctrl          control.Options

	// TODO: these should be interfaces, not funcs
	itemReader     itemReaderFunc
	itemMetaReader itemMetaReaderFunc
	itemGetter     itemGetterFunc

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

type itemGetterFunc func(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error)

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
		ctrl:          ctrlOpts,
		data:          make(chan data.Stream, collectionChannelBufferSize),
		driveID:       driveID,
		driveItems:    map[string]models.DriveItemable{},
		folderPath:    folderPath,
		itemClient:    itemClient,
		itemGetter:    getDriveItem,
		service:       service,
		source:        source,
		statusUpdater: statusUpdater,
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

// populateItems iterates through items added to the collection
// and uses the collection `itemReader` to read the item
func (oc *Collection) populateItems(ctx context.Context) {
	// Retrieve the OneDrive folder path to set later in `details.OneDriveInfo`
	parentPathString, err := path.GetDriveFolderPath(oc.folderPath)
	if err != nil {
		oc.reportAsCompleted(ctx, 0, 0, 0, 0, 0, err)
		return
	}

	var (
		errs       error
		byteCount  int64
		itemsRead  int64
		dirsRead   int64
		itemsFound int64
		dirsFound  int64

		wg sync.WaitGroup
		m  sync.Mutex

		errUpdater = func(id string, err error) {
			m.Lock()
			defer m.Unlock()
			errs = support.WrapAndAppend(id, err, errs)
		}
		countUpdater = func(size, dirs, items, dReads, iReads int64) {
			atomic.AddInt64(&byteCount, size)
			atomic.AddInt64(&dirsFound, dirs)
			atomic.AddInt64(&itemsFound, items)
			atomic.AddInt64(&dirsRead, dReads)
			atomic.AddInt64(&itemsRead, iReads)
		}
	)

	folderProgress, colCloser := observe.ProgressWithCount(
		ctx,
		observe.ItemQueueMsg,
		observe.PII("/"+parentPathString),
		int64(len(oc.driveItems)))
	defer colCloser()
	defer close(folderProgress)

	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	for _, item := range oc.driveItems {
		if oc.ctrl.FailFast && errs != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		// fetch the item's metadata (if it has any, will no-op otherwise),
		// and stream it into the collection's data channel
		go oc.streamItemMeta(
			ctx,
			&wg,
			semaphoreCh,
			folderProgress,
			errUpdater,
			countUpdater,
			item,
			parentPathString)

		semaphoreCh <- struct{}{}

		wg.Add(1)

		// fetch the item, and stream it into the collection's data channel
		go oc.streamItem(
			ctx,
			&wg,
			semaphoreCh,
			folderProgress,
			errUpdater,
			countUpdater,
			item,
			parentPathString)
	}

	wg.Wait()

	oc.reportAsCompleted(ctx, dirsFound, dirsRead, itemsFound, itemsRead, byteCount, errs)
}

func (oc *Collection) streamItemMeta(
	ctx context.Context,
	wg *sync.WaitGroup,
	semaphore <-chan struct{},
	progress chan<- struct{},
	errUpdater func(string, error),
	countUpdater func(int64, int64, int64, int64, int64),
	item models.DriveItemable,
	parentPath string,
) {
	defer wg.Done()
	defer func() { <-semaphore }()

	var (
		id     = *item.GetId()
		name   = *item.GetName()
		size   = *item.GetSize()
		isFile = item.GetFile() != nil

		lazyRC   io.ReadCloser
		info     details.ItemInfo
		metadata io.ReadCloser

		suffix     string
		dirsFound  int64
		dirsRead   int64
		itemsFound int64
		itemsRead  int64

		err error
	)

	// we only handle metadata for onedrive right now
	if oc.source != OneDriveSource {
		return
	}

	if isFile {
		itemsFound, itemsRead = 1, 1
		suffix = MetaFileSuffix
	} else {
		dirsFound, dirsRead = 1, 1
		suffix = DirMetaFileSuffix
	}

	info.OneDrive = oneDriveItemInfo(item, size)
	info.OneDrive.ParentPath = parentPath

	// Construct a new lazy readCloser to feed to the collection consumer.
	// This ensures that downloads won't be attempted unless that consumer
	// attempts to read bytes.  Assumption is that kopia will check things
	// like file modtimes before attempting to read.
	lazyRC = lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		var msize int
		metadata, msize, err = getItemMeta(
			ctx,
			oc.service,
			oc.driveID,
			item,
			maxRetries,
			oc.ctrl.ToggleFeatures.EnablePermissionsBackup,
			oc.itemMetaReader)
		if err != nil {
			errUpdater(id, err)
			return nil, err
		}

		name += suffix
		size = int64(msize)

		progReader, closer := observe.ItemProgress(
			ctx,
			metadata,
			observe.ItemBackupMsg,
			observe.PII(name),
			size)
		go closer()
		return progReader, nil
	})

	// Item read successfully, record its addition.
	//
	// Note: this can cause inaccurate counts.  Right now it counts all
	// the items we intend to read.  Errors within the lazy readCloser
	// will create a conflict: an item is both successful and erroneous.
	// But the async control to fix that is more error-prone than helpful.
	//
	// TODO: transform this into a stats bus so that async control of stats
	// aggregation is handled at the backup level, not at the item iteration
	// level.
	countUpdater(size, dirsFound, itemsFound, dirsRead, itemsRead)

	// stream the item to the data consumer.
	oc.data <- &Item{
		id:   name,
		data: lazyRC,
		info: info,
	}

	progress <- struct{}{}
}

func (oc *Collection) streamItem(
	ctx context.Context,
	wg *sync.WaitGroup,
	semaphore <-chan struct{},
	progress chan<- struct{},
	errUpdater func(string, error),
	countUpdater func(int64, int64, int64, int64, int64),
	item models.DriveItemable,
	parentPath string,
) {
	defer wg.Done()
	defer func() { <-semaphore }()

	var (
		id     = *item.GetId()
		name   = *item.GetName()
		size   = *item.GetSize()
		isFile = item.GetFile() != nil

		lazyRC io.ReadCloser
		info   details.ItemInfo

		dirsFound  int64
		itemsFound int64
		dirsRead   int64
		itemsRead  int64
	)

	if isFile {
		itemsFound++
		itemsRead++
	} else {
		dirsFound++
		dirsRead++
	}

	switch oc.source {
	case SharePointSource:
		info.SharePoint = sharePointItemInfo(item, size)
		info.SharePoint.ParentPath = parentPath
	default:
		info.OneDrive = oneDriveItemInfo(item, size)
		info.OneDrive.ParentPath = parentPath
	}

	// Construct a new lazy readCloser to feed to the collection consumer.
	// This ensures that downloads won't be attempted unless that consumer
	// attempts to read bytes.  Assumption is that kopia will check things
	// like file modtimes before attempting to read.
	lazyRC = lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		itemData, _, err := readDriveItem(
			ctx,
			oc.service,
			oc.itemClient,
			oc.driveID, id,
			item,
			oc.itemReader,
			oc.itemGetter)
		if err != nil {
			errUpdater(id, err)
			return nil, err
		}

		// this separates onedrive items from their metadata
		if oc.source == OneDriveSource && isFile {
			name += DataFileSuffix
		}

		progReader, closer := observe.ItemProgress(
			ctx,
			itemData,
			observe.ItemBackupMsg,
			observe.PII(name),
			size)
		go closer()
		return progReader, nil
	})

	// Item read successfully, record its addition.
	//
	// Note: this can cause inaccurate counts.  Right now it counts all
	// the items we intend to read.  Errors within the lazy readCloser
	// will create a conflict: an item is both successful and erroneous.
	// But the async control to fix that is more error-prone than helpful.
	//
	// TODO: transform this into a stats bus so that async control of stats
	// aggregation is handled at the backup level, not at the item iteration
	// level.
	countUpdater(size, dirsFound, itemsFound, dirsRead, itemsRead)

	// stream the item to the data consumer.
	oc.data <- &Item{
		id:   name,
		data: lazyRC,
		info: info,
	}

	progress <- struct{}{}
}

func getItemMeta(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	item models.DriveItemable,
	maxRetries int,
	enablePermissionsBackup bool,
	read itemMetaReaderFunc,
) (io.ReadCloser, int, error) {
	if !enablePermissionsBackup {
		// We are still writing the metadata file but with
		// empty permissions as we don't have a way to
		// signify that the permissions was explicitly
		// not added.
		return io.NopCloser(strings.NewReader("{}")), 2, nil
	}

	var (
		rc   io.ReadCloser
		size int
		err  error
	)

	err = graph.RunWithRetry(func() error {
		rc, size, err = read(ctx, service, driveID, item)
		return err
	})

	return rc, size, nil
}

func readDriveItem(
	ctx context.Context,
	service graph.Servicer,
	itemClient *http.Client,
	driveID, itemID string,
	original models.DriveItemable,
	read itemReaderFunc,
	get itemGetterFunc,
) (io.ReadCloser, models.DriveItemable, error) {
	var (
		err  error
		rc   io.ReadCloser
		item = original
	)

	err = graph.RunWithRetry(func() error {
		_, rc, err = read(itemClient, item)
		return err
	})

	return rc, item, err
}

func (oc *Collection) reportAsCompleted(
	ctx context.Context,
	dirsFound, dirsRead, itemsFound, itemsRead, byteCount int64,
	errs error,
) {
	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		1, // num folders (always 1)
		support.CollectionMetrics{
			Objects:    int(itemsFound) + int(dirsFound), // items to read,
			Successes:  int(itemsRead) + int(dirsRead),   // items read successfully,
			TotalBytes: byteCount,                        // Number of bytes read in the operation,
		},
		errs,
		oc.folderPath.Folder(), // Additional details
	)
	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())
	oc.statusUpdater(status)
}
