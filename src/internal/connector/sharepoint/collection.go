package sharepoint

import (
	"bytes"
	"context"
	"io"

	kw "github.com/microsoft/kiota-serialization-json-go"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
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
	_ data.Collection = &Collection{}
	_ data.Stream     = &Item{}
)

type Collection struct {
	data chan data.Stream
	jobs []string
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// M365 IDs of the items of this collection
	service       graph.Service
	statusUpdater support.StatusUpdater
}

func NewCollection(
	folderPath path.Path,
	service graph.Service,
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

func (sc *Collection) Items() <-chan data.Stream {
	return sc.data
}

type Item struct {
	id   string
	data io.ReadCloser
	info *details.SharePointInfo
}

func (sd *Item) UUID() string {
	return sd.id
}

func (sd *Item) ToReader() io.ReadCloser {
	return sd.data
}

func (sd *Item) Info() details.ItemInfo {
	return details.ItemInfo{SharePoint: sd.info}
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

func (sc *Collection) populate(ctx context.Context) {
	var (
		success                 int
		totalBytes, arrayLength int64
		errs                    error
		writer                  = kw.NewJsonSerializationWriter()
	)

	// sc.jobs contains query = all of the site IDs.
	for _, identifier := range sc.jobs {
		// Retrieve list data from M365
		lists, err := loadLists(ctx, sc.service, identifier)
		// Transfer list data to
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

			arrayLength = int64(len(byteArray))

			if arrayLength > 0 {
				success++
				totalBytes += arrayLength
				sc.data <- &Item{id: *lst.GetId(), data: io.NopCloser(bytes.NewReader(byteArray)), info: sharePointListInfo(lst, arrayLength)}
			}
		}

	}
}
