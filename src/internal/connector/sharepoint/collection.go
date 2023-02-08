package sharepoint

import (
	"bytes"
	"context"
	"io"
	"time"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

type DataCategory int

//go:generate stringer -type=DataCategory
const (
	collectionChannelBufferSize              = 50
	Unknown                     DataCategory = iota
	List
	Drive
)

var (
	_ data.BackupCollection = &Collection{}
	_ data.Stream           = &Item{}
	_ data.StreamInfo       = &Item{}
	_ data.StreamModTime    = &Item{}
)

// Collection is the SharePoint.List implementation of data.Collection. SharePoint.Libraries collections are supported
// by the oneDrive.Collection as the calls are identical for populating the Collection
type Collection struct {
	// data is the container for each individual SharePoint.List
	data chan data.Stream
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// jobs contain the SharePoint.Site.ListIDs for the associated list(s).
	jobs []string
	// M365 IDs of the items of this collection
	service       graph.Servicer
	betaService   *api.BetaService
	statusUpdater support.StatusUpdater
}

// NewCollection helper function for creating a Collection
func NewCollection(
	folderPath path.Path,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
) *Collection {
	c := &Collection{
		fullPath:      folderPath,
		jobs:          make([]string, 0),
		data:          make(chan data.Stream, collectionChannelBufferSize),
		service:       service,
		statusUpdater: statusUpdater,
	}

	return c
}

// AddJob appends additional objectID to job field
func (sc *Collection) AddJob(objID string) {
	sc.jobs = append(sc.jobs, objID)
}

func (sc *Collection) FullPath() path.Path {
	return sc.fullPath
}

// TODO(ashmrtn): Fill in with previous path once GraphConnector compares old
// and new folder hierarchies.
func (sc Collection) PreviousPath() path.Path {
	return nil
}

func (sc Collection) State() data.CollectionState {
	return data.NewState
}

func (sc Collection) DoNotMergeItems() bool {
	return false
}

func (sc *Collection) Items() <-chan data.Stream {
	go sc.populate(context.TODO())
	return sc.data
}

type Item struct {
	id      string
	data    io.ReadCloser
	info    *details.SharePointInfo
	modTime time.Time

	// true if the item was marked by graph as deleted.
	deleted bool
}

func NewItem(name string, d io.ReadCloser) *Item {
	item := &Item{
		id:   name,
		data: d,
	}

	return item
}

func (sd *Item) UUID() string {
	return sd.id
}

func (sd *Item) ToReader() io.ReadCloser {
	return sd.data
}

func (sd Item) Deleted() bool {
	return sd.deleted
}

func (sd *Item) Info() details.ItemInfo {
	return details.ItemInfo{SharePoint: sd.info}
}

func (sd *Item) ModTime() time.Time {
	return sd.modTime
}

func (sc *Collection) finishPopulation(ctx context.Context, attempts, success int, totalBytes int64, errs error) {
	close(sc.data)

	attempted := attempts
	status := support.CreateStatus(
		ctx,
		support.Backup,
		len(sc.jobs),
		support.CollectionMetrics{
			Objects:    attempted,
			Successes:  success,
			TotalBytes: totalBytes,
		},
		errs,
		sc.fullPath.Folder())
	logger.Ctx(ctx).Debug(status.String())

	if sc.statusUpdater != nil {
		sc.statusUpdater(status)
	}
}

// populate utility function to retrieve data from back store for a given collection
func (sc *Collection) populate(ctx context.Context) {
	var (
		objects, success        int
		totalBytes, arrayLength int64
		errs                    error
		writer                  = kw.NewJsonSerializationWriter()
	)

	// TODO: Insert correct ID for CollectionProgress
	colProgress, closer := observe.CollectionProgress(
		ctx,
		sc.fullPath.Category().String(),
		observe.Safe("name"),
		observe.PII(sc.fullPath.Folder()))
	go closer()

	defer func() {
		close(colProgress)
		sc.finishPopulation(ctx, objects, success, totalBytes, errs)
	}()

	// Retrieve list data from M365
	lists, err := loadSiteLists(ctx, sc.service, sc.fullPath.ResourceOwner(), sc.jobs)
	if err != nil {
		errs = support.WrapAndAppend(sc.fullPath.ResourceOwner(), err, errs)
	}

	objects += len(lists)
	// Write Data and Send
	for _, lst := range lists {
		byteArray, err := serializeListContent(writer, lst)
		if err != nil {
			errs = support.WrapAndAppend(*lst.GetId(), err, errs)
			continue
		}

		arrayLength = int64(len(byteArray))

		if arrayLength > 0 {
			t := time.Now()
			if t1 := lst.GetLastModifiedDateTime(); t1 != nil {
				t = *t1
			}

			totalBytes += arrayLength

			success++
			sc.data <- &Item{
				id:      *lst.GetId(),
				data:    io.NopCloser(bytes.NewReader(byteArray)),
				info:    sharePointListInfo(lst, arrayLength),
				modTime: t,
			}

			colProgress <- struct{}{}
		}
	}
}

func serializeListContent(writer *kw.JsonSerializationWriter, lst models.Listable) ([]byte, error) {
	defer writer.Close()

	err := writer.WriteObjectValue("", lst)
	if err != nil {
		return nil, err
	}

	byteArray, err := writer.GetSerializedContent()
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}
