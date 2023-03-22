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

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-lazy/pkg/lazy"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO: This number needs to be tuned
	// Consider max open file limit `ulimit -n`, usually 1024 when setting this value
	collectionChannelBufferSize = 5

	// TODO: Tune this later along with collectionChannelBufferSize
	urlPrefetchChannelBufferSize = 5

	// maxDownloadRetires specifies the number of times a file download should
	// be retried
	maxDownloadRetires = 3

	MetaFileSuffix    = ".meta"
	DirMetaFileSuffix = ".dirmeta"
	DataFileSuffix    = ".data"
)

func IsMetaFile(name string) bool {
	return strings.HasSuffix(name, MetaFileSuffix) || strings.HasSuffix(name, DirMetaFileSuffix)
}

var (
	_ data.BackupCollection = &Collection{}
	_ data.Stream           = &Item{}
	_ data.StreamInfo       = &Item{}
	_ data.StreamModTime    = &Item{}
	_ data.Stream           = &metadataItem{}
	_ data.StreamModTime    = &metadataItem{}
)

type SharingMode int

const (
	SharingModeCustom = SharingMode(iota)
	SharingModeInherited
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

	// Primary M365 ID of the drive this collection was created from
	driveID string
	// Display Name of the associated drive
	driveName      string
	source         driveSource
	service        graph.Servicer
	statusUpdater  support.StatusUpdater
	itemGetter     itemGetterFunc
	itemReader     itemReaderFunc
	itemMetaReader itemMetaReaderFunc
	ctrl           control.Options

	// PrevPath is the previous hierarchical path used by this collection.
	// It may be the same as fullPath, if the folder was not renamed or
	// moved.  It will be empty on its first retrieval.
	prevPath path.Path

	// Specifies if it new, moved/rename or deleted
	state data.CollectionState

	// should only be true if the old delta token expired
	doNotMergeItems bool
}

// itemGetterFunc gets an specified item
type itemGetterFunc func(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error)

// itemReadFunc returns a reader for the specified item
type itemReaderFunc func(
	ctx context.Context,
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error)

// itemMetaReaderFunc returns a reader for the metadata of the
// specified item
type itemMetaReaderFunc func(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	item models.DriveItemable,
	fetchPermissions bool,
) (io.ReadCloser, int, error)

// NewCollection creates a Collection
func NewCollection(
	itemClient *http.Client,
	folderPath path.Path,
	prevPath path.Path,
	driveID string,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	source driveSource,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) *Collection {
	c := &Collection{
		itemClient:      itemClient,
		folderPath:      folderPath,
		prevPath:        prevPath,
		driveItems:      map[string]models.DriveItemable{},
		driveID:         driveID,
		source:          source,
		service:         service,
		data:            make(chan data.Stream, collectionChannelBufferSize),
		statusUpdater:   statusUpdater,
		ctrl:            ctrlOpts,
		state:           data.StateOf(prevPath, folderPath),
		doNotMergeItems: doNotMergeItems,
	}

	// Allows tests to set a mock populator
	switch source {
	case SharePointSource:
		c.itemGetter = getDriveItem
		c.itemReader = sharePointItemReader
		c.itemMetaReader = sharePointItemMetaReader
	default:
		c.itemGetter = getDriveItem
		c.itemReader = oneDriveItemReader
		c.itemMetaReader = oneDriveItemMetaReader
	}

	return c
}

// Adds an itemID to the collection.  This will make it eligible to be
// populated. The return values denotes if the item was previously
// present or is new one.
func (oc *Collection) Add(item models.DriveItemable) bool {
	_, found := oc.driveItems[ptr.Val(item.GetId())]
	oc.driveItems[ptr.Val(item.GetId())] = item

	return !found // !found = new
}

// Remove removes a item from the collection
func (oc *Collection) Remove(itemID string) bool {
	_, found := oc.driveItems[itemID]
	if !found {
		return false
	}

	delete(oc.driveItems, itemID)

	return true
}

// IsEmpty check if a collection does not contain any items
// TODO(meain): Should we just have function that returns driveItems?
func (oc *Collection) IsEmpty() bool {
	return len(oc.driveItems) == 0
}

// Items() returns the channel containing M365 Exchange objects
func (oc *Collection) Items(
	ctx context.Context,
	errs *fault.Bus, // TODO: currently unused while onedrive isn't up to date with clues/fault
) <-chan data.Stream {
	go oc.populateItems(ctx, errs)
	return oc.data
}

func (oc *Collection) FullPath() path.Path {
	return oc.folderPath
}

func (oc Collection) PreviousPath() path.Path {
	return oc.prevPath
}

func (oc *Collection) SetFullPath(curPath path.Path) {
	oc.folderPath = curPath
	oc.state = data.StateOf(oc.prevPath, curPath)
}

func (oc Collection) State() data.CollectionState {
	return oc.state
}

func (oc Collection) DoNotMergeItems() bool {
	return oc.doNotMergeItems
}

// FilePermission is used to store permissions of a specific user to a
// OneDrive item.
type UserPermission struct {
	ID         string     `json:"id,omitempty"`
	Roles      []string   `json:"role,omitempty"`
	Email      string     `json:"email,omitempty"` // DEPRECATED: Replaced with UserID in newer backups
	EntityID   string     `json:"entityId,omitempty"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

// ItemMeta contains metadata about the Item. It gets stored in a
// separate file in kopia
type Metadata struct {
	FileName string `json:"filename,omitempty"`
	// SharingMode denotes what the current mode of sharing is for the object.
	// - inherited: permissions same as parent permissions (no "shared" in delta)
	// - custom: use Permissions to set correct permissions ("shared" has value in delta)
	SharingMode SharingMode      `json:"permissionMode,omitempty"`
	Permissions []UserPermission `json:"permissions,omitempty"`
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id   string
	data io.ReadCloser
	info details.ItemInfo
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return od.data
}

// Deleted implements an interface function. However, OneDrive items are marked
// as deleted by adding them to the exclude list so this can always return
// false.
func (od Item) Deleted() bool {
	return false
}

func (od *Item) Info() details.ItemInfo {
	return od.info
}

func (od *Item) ModTime() time.Time {
	return od.info.Modified()
}

type metadataItem struct {
	id      string
	data    io.ReadCloser
	modTime time.Time
}

func (od *metadataItem) UUID() string {
	return od.id
}

func (od *metadataItem) ToReader() io.ReadCloser {
	return od.data
}

// Deleted implements an interface function. However, OneDrive items are marked
// as deleted by adding them to the exclude list so this can always return
// false.
func (od metadataItem) Deleted() bool {
	return false
}

func (od *metadataItem) ModTime() time.Time {
	return od.modTime
}

// getDriveItemContent fetch drive item's contents with retries
func (oc *Collection) getDriveItemContent(
	ctx context.Context,
	item models.DriveItemable,
	errs *fault.Bus,
) (io.ReadCloser, error) {
	var (
		itemID   = ptr.Val(item.GetId())
		itemName = ptr.Val(item.GetName())
		el       = errs.Local()

		itemData io.ReadCloser
		err      error
	)

	// Initial try with url from delta + 2 retries
	for i := 1; i <= maxDownloadRetires; i++ {
		_, itemData, err = oc.itemReader(ctx, oc.itemClient, item)
		if err == nil || !graph.IsErrUnauthorized(err) {
			break
		}

		// Assume unauthorized requests are a sign of an expired jwt
		// token, and that we've overrun the available window to
		// download the actual file.  Re-downloading the item will
		// refresh that download url.
		di, diErr := oc.itemGetter(ctx, oc.service, oc.driveID, itemID)
		if diErr != nil {
			err = errors.Wrap(diErr, "retrieving expired item")
			break
		}

		item = di
	}

	// check for errors following retries
	if err != nil {
		if clues.HasLabel(err, graph.LabelsMalware) || (item != nil && item.GetMalware() != nil) {
			logger.CtxErr(ctx, err).With("skipped_reason", fault.SkipMalware).Info("item flagged as malware")
			el.AddSkip(fault.FileSkip(fault.SkipMalware, itemID, itemName, graph.ItemInfo(item)))

			return nil, clues.Wrap(err, "downloading item").Label(graph.LabelsSkippable)
		}

		if clues.HasLabel(err, graph.LabelStatus(http.StatusNotFound)) || graph.IsErrDeletedInFlight(err) {
			logger.CtxErr(ctx, err).With("skipped_reason", fault.SkipNotFound).Error("item not found")
			el.AddSkip(fault.FileSkip(fault.SkipNotFound, itemID, itemName, graph.ItemInfo(item)))

			return nil, clues.Wrap(err, "downloading item").Label(graph.LabelsSkippable)
		}

		logger.CtxErr(ctx, err).Error("downloading item")
		el.AddRecoverable(clues.Stack(err).WithClues(ctx).Label(fault.LabelForceNoBackupCreation))

		// return err, not el.Err(), because the lazy reader needs to communicate to
		// the data consumer that this item is unreadable, regardless of the fault state.
		return nil, clues.Wrap(err, "downloading item")
	}

	return itemData, nil
}

// populateItems iterates through items added to the collection
// and uses the collection `itemReader` to read the item
func (oc *Collection) populateItems(ctx context.Context, errs *fault.Bus) {
	var (
		byteCount  int64
		itemsRead  int64
		dirsRead   int64
		itemsFound int64
		dirsFound  int64
		wg         sync.WaitGroup
		el         = errs.Local()
	)

	// Retrieve the OneDrive folder path to set later in
	// `details.OneDriveInfo`
	parentPathString, err := path.GetDriveFolderPath(oc.folderPath)
	if err != nil {
		oc.reportAsCompleted(ctx, 0, 0, 0)
		return
	}

	queuedPath := "/" + parentPathString
	if oc.source == SharePointSource && len(oc.driveName) > 0 {
		queuedPath = "/" + oc.driveName + queuedPath
	}

	folderProgress, colCloser := observe.ProgressWithCount(
		ctx,
		observe.ItemQueueMsg,
		observe.PII(queuedPath),
		int64(len(oc.driveItems)))
	defer colCloser()
	defer close(folderProgress)

	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	for _, item := range oc.driveItems {
		if el.Failure() != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(ctx context.Context, item models.DriveItemable) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			// Read the item
			var (
				itemID       = ptr.Val(item.GetId())
				itemName     = ptr.Val(item.GetName())
				itemSize     = ptr.Val(item.GetSize())
				itemInfo     details.ItemInfo
				itemMeta     io.ReadCloser
				itemMetaSize int
				metaFileName string
				metaSuffix   string
				err          error
			)

			ctx = clues.Add(
				ctx,
				"backup_item_id", itemID,
				"backup_item_name", itemName,
				"backup_item_size", itemSize)

			item.SetParentReference(setName(item.GetParentReference(), oc.driveName))

			isFile := item.GetFile() != nil

			if isFile {
				atomic.AddInt64(&itemsFound, 1)

				metaFileName = itemID
				metaSuffix = MetaFileSuffix
			} else {
				atomic.AddInt64(&dirsFound, 1)

				// metaFileName not set for directories so we get just ".dirmeta"
				metaSuffix = DirMetaFileSuffix
			}

			// Fetch metadata for the file
			itemMeta, itemMetaSize, err = oc.itemMetaReader(
				ctx,
				oc.service,
				oc.driveID,
				item,
				oc.ctrl.ToggleFeatures.EnablePermissionsBackup)

			if err != nil {
				el.AddRecoverable(clues.Wrap(err, "getting item metadata").Label(fault.LabelForceNoBackupCreation))
				return
			}

			switch oc.source {
			case SharePointSource:
				itemInfo.SharePoint = sharePointItemInfo(item, itemSize)
				itemInfo.SharePoint.ParentPath = parentPathString
			default:
				itemInfo.OneDrive = oneDriveItemInfo(item, itemSize)
				itemInfo.OneDrive.ParentPath = parentPathString
			}

			ctx = clues.Add(ctx, "backup_item_info", itemInfo)

			if isFile {
				dataSuffix := DataFileSuffix

				// Construct a new lazy readCloser to feed to the collection consumer.
				// This ensures that downloads won't be attempted unless that consumer
				// attempts to read bytes.  Assumption is that kopia will check things
				// like file modtimes before attempting to read.
				itemReader := lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
					itemData, err := oc.getDriveItemContent(ctx, item, errs)
					if err != nil {
						return nil, err
					}

					// display/log the item download
					progReader, closer := observe.ItemProgress(
						ctx,
						itemData,
						observe.ItemBackupMsg,
						observe.PII(itemID+dataSuffix),
						itemSize)
					go closer()

					return progReader, nil
				})

				oc.data <- &Item{
					id:   itemID + dataSuffix,
					data: itemReader,
					info: itemInfo,
				}
			}

			metaReader := lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
				progReader, closer := observe.ItemProgress(
					ctx, itemMeta, observe.ItemBackupMsg,
					observe.PII(metaFileName+metaSuffix), int64(itemMetaSize))
				go closer()
				return progReader, nil
			})

			oc.data <- &metadataItem{
				id:      metaFileName + metaSuffix,
				data:    metaReader,
				modTime: time.Now(),
			}

			// Item read successfully, add to collection
			if isFile {
				atomic.AddInt64(&itemsRead, 1)
			} else {
				atomic.AddInt64(&dirsRead, 1)
			}

			// byteCount iteration
			atomic.AddInt64(&byteCount, itemSize)

			folderProgress <- struct{}{}
		}(ctx, item)
	}

	wg.Wait()

	oc.reportAsCompleted(ctx, int(itemsFound), int(itemsRead), byteCount)
}

func (oc *Collection) reportAsCompleted(ctx context.Context, itemsFound, itemsRead int, byteCount int64) {
	close(oc.data)

	status := support.CreateStatus(ctx, support.Backup,
		1, // num folders (always 1)
		support.CollectionMetrics{
			Objects:   itemsFound,
			Successes: itemsRead,
			Bytes:     byteCount,
		},
		oc.folderPath.Folder(false))

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	oc.statusUpdater(status)
}
