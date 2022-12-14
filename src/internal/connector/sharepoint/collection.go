package sharepoint

import (
	"bytes"
	"context"
	"io"
	"time"

	kw "github.com/microsoft/kiota-serialization-json-go"

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
	_ data.Collection    = &Collection{}
	_ data.Stream        = &Item{}
	_ data.StreamInfo    = &Item{}
	_ data.StreamModTime = &Item{}
)

type Collection struct {
	data chan data.Stream
	jobs []string
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// M365 IDs of the items of this collection
	service       graph.Servicer
	statusUpdater support.StatusUpdater
}

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

// TODO(ashmrtn): Fill in once GraphConnector compares old and new folder
// hierarchies.
func (sc Collection) State() data.CollectionState {
	return data.NewState
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
}

func (sd *Item) UUID() string {
	return sd.id
}

func (sd *Item) ToReader() io.ReadCloser {
	return sd.data
}

func (sd Item) Deleted() bool {
	return false
}

func (sd *Item) Info() details.ItemInfo {
	return details.ItemInfo{SharePoint: sd.info}
}

func (sd *Item) ModTime() time.Time {
	return sd.modTime
}

func (sc *Collection) finishPopulation(ctx context.Context, success int, totalBytes int64, errs error) {
	close(sc.data)
	attempted := len(sc.jobs)
	status := support.CreateStatus(
		ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:    attempted,
			Successes:  success,
			TotalBytes: totalBytes,
		},
		errs,
		sc.fullPath.Folder())
	logger.Ctx(ctx).Debug(status.String())
}

// populate utility function to retrieve data from back store for a given collection
func (sc *Collection) populate(ctx context.Context) {
	var (
		success                 int
		totalBytes, arrayLength int64
		errs                    error
		writer                  = kw.NewJsonSerializationWriter()
	)

	// TODO: Insert correct ID for CollectionProgress
	colProgress, closer := observe.CollectionProgress("name", sc.fullPath.Category().String(), sc.fullPath.Folder())
	go closer()

	defer func() {
		close(colProgress)
		sc.finishPopulation(ctx, success, totalBytes, errs)
	}()

	// sc.jobs contains query = all of the site IDs.
	for _, id := range sc.jobs {
		// Retrieve list data from M365
		lists, err := loadLists(ctx, sc.service, id)
		if err != nil {
			errs = support.WrapAndAppend(id, err, errs)
		}
		// Write Data and Send
		for _, lst := range lists {
			err = writer.WriteObjectValue("", lst)
			if err != nil {
				errs = support.WrapAndAppend(*lst.GetId(), err, errs)
				continue
			}

			byteArray, err := writer.GetSerializedContent()
			if err != nil {
				errs = support.WrapAndAppend(*lst.GetId(), err, errs)
				continue
			}

			writer.Close()

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
}
