package sharepoint

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
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
		success     int
		totalyBytes int64
		errs        error
		writer      = kw.NewJsonSerializationWriter()
	)

	// sc.jobs contains query = all of the site IDs.
	for _, identifier := range sc.jobs {
		query, err := sc.service.Client().SitesById(identifier).Lists().Get(ctx, nil)
		if err != nil {
			errs = support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}
		lists, err := loadLists(ctx, sc.service, identifier, query)
		for _, completeList := range lists {
			err = writer.WriteObjectValue("", completeList)
			if err != nil {
				errs = support.WrapAndAppend(*completeList.GetId(), err, errs)
				continue
			}

			byteArray, err := writer.GetSerializedContent()
			if err != nil {
				errs = support.WrapAndAppend(*completeList.GetId(), err, errs)
				continue
			}

			if len(byteArray) > 0 {
				success++
				sc.data <- &Stream{id: completeList.GetId()}
			}
		}

		// serialize the data
		// place in data stream
	}
}

func loadLists(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	resp models.ListCollectionResponseable,
) ([]models.Listable, error) {
	listing := make([]models.Listable, 0)
	prefix := gs.Client().SitesById(identifier)

	for _, entry := range resp.GetValue() {
		id := *entry.GetId()
		// get columns
		q1, _ := prefix.ListsById(id).Columns().Get(ctx, nil)
		cols, _ := loadColumns(ctx, gs, identifier, q1)
		entry.SetColumns(cols)
		// get contentTypes

		q2, _ := prefix.ListsById(id).ContentTypes().Get(ctx, nil)
		if q2 != nil {
			cTypes, err := loadContentTypes(ctx, gs, identifier, q2)
			if err != nil {
				return nil, err
			}

			entry.SetContentTypes(cTypes)
		}

		q3, err := prefix.ListsById(id).Items().Get(ctx, nil)

		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		if q3 != nil {
			items, _ := loadListItems(ctx, gs, identifier, id, q3)
			entry.SetItems(items)
		}

		listing = append(listing, entry)
	}

	return listing, nil
}

// Stream represents an individual SharePoint object retrieved from exchange
type Stream struct {
	id      string
	message []byte
	info    *details.SharePointInfo
}

func NewStream(streamID string, dataBytes []byte, detail details.SharePointInfo) Stream {
	return Stream{
		id:      streamID,
		message: dataBytes,
		info:    &detail,
	}

}

//==============================
// Interface Functions
//==============================

func (od *Stream) UUID() string {
	return od.id
}

func (od *Stream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.message))
}

func (od *Stream) Info() details.ItemInfo {
	return details.ItemInfo{SharePoint: od.info}
}
