package sharepoint

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
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
	Pages
)

var (
	_ data.Collection    = &Collection{}
	_ data.Stream        = &Item{}
	_ data.StreamInfo    = &Item{}
	_ data.StreamModTime = &Item{}
)

type numMetrics struct {
	attempts   int
	success    int
	totalBytes int64
}

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
	category      DataCategory
	service       graph.Servicer
	ctrl          control.Options
	betaService   *api.BetaService
	statusUpdater support.StatusUpdater
}

// NewCollection helper function for creating a Collection
func NewCollection(
	folderPath path.Path,
	service graph.Servicer,
	category DataCategory,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collection {
	c := &Collection{
		fullPath:      folderPath,
		jobs:          make([]string, 0),
		data:          make(chan data.Stream, collectionChannelBufferSize),
		service:       service,
		statusUpdater: statusUpdater,
		category:      category,
		ctrl:          ctrlOpts,
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
		1, // 1 folder
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
		metrics numMetrics
		errs    error
		writer  = kw.NewJsonSerializationWriter()
	)

	defer func() {
		sc.finishPopulation(ctx, metrics.attempts, metrics.success, int64(metrics.success), errs)
	}()
	// TODO: Insert correct ID for CollectionProgress
	colProgress, closer := observe.CollectionProgress(
		ctx,
		sc.fullPath.Category().String(),
		observe.Safe("name"),
		observe.PII(sc.fullPath.Folder()))
	go closer()

	defer func() {
		close(colProgress)
	}()

	// Switch retrieval function based on category
	switch sc.category {
	case List:
		// do the thing
		// ctx, service, writer
		metrics, errs = sc.retrieveLists(ctx, writer, colProgress)
	case Pages:
		// do the other thing
		metrics, errs = sc.retrievePages(ctx, writer, colProgress)
	}
}

// retrieveLists utility function for collection that downloads and serializes
// models.Listable objects based on M365 IDs from the jobs field.
func (sc *Collection) retrieveLists(
	ctx context.Context,
	wtr *kw.JsonSerializationWriter,
	progress chan<- struct{},
) (numMetrics, error) {
	var (
		errs    error
		metrics numMetrics
	)

	lists, err := loadSiteLists(ctx, sc.service, sc.fullPath.ResourceOwner(), sc.jobs)
	if err != nil {
		return metrics, errors.Wrap(err, sc.fullPath.ResourceOwner())
	}

	metrics.attempts += len(lists)
	// For each models.Listable, object is serialized and the metrics are collected.
	// The progress is objected via the passed in channel.
	for _, lst := range lists {
		byteArray, err := serializeContent(wtr, lst)
		if err != nil {
			errs = support.WrapAndAppend(*lst.GetId(), err, errs)
			if sc.ctrl.FailFast {
				return metrics, errs
			}

			continue
		}

		arrayLength := int64(len(byteArray))

		if arrayLength > 0 {
			t := time.Now()
			if t1 := lst.GetLastModifiedDateTime(); t1 != nil {
				t = *t1
			}

			metrics.totalBytes += arrayLength

			metrics.success++
			sc.data <- &Item{
				id:      *lst.GetId(),
				data:    io.NopCloser(bytes.NewReader(byteArray)),
				info:    sharePointListInfo(lst, arrayLength),
				modTime: t,
			}

			progress <- struct{}{}
		}
	}

	return metrics, nil
}

func (sc *Collection) retrievePages(
	ctx context.Context,
	wtr *kw.JsonSerializationWriter,
	progress chan<- struct{},
) (numMetrics, error) {
	var (
		errs    error
		metrics numMetrics
	)

	betaService := sc.betaService
	if betaService == nil {
		return metrics, fmt.Errorf("beta service not found in collection")
	}

	pages, err := GetSitePages(ctx, betaService, sc.fullPath.ResourceOwner(), sc.jobs)
	if err != nil {
		return metrics, errors.Wrap(err, sc.fullPath.ResourceOwner())
	}

	metrics.attempts = len(pages)
	// For each models.Pageable, object is serialize and the metrics are collected and returned.
	// Pageable objects are not supported in v1.0 of msgraph at this time.
	// TODO: Verify Parsable interface supported with modified-Pageable
	for _, pg := range pages {
		byteArray, err := serializeContent(wtr, pg)
		if err != nil {
			errs = support.WrapAndAppend(*pg.GetId(), err, errs)
			if sc.ctrl.FailFast {
				return metrics, errs
			}

			continue
		}

		arrayLength := int64(len(byteArray))

		if arrayLength > 0 {
			t := time.Now()
			if t1 := pg.GetLastModifiedDateTime(); t1 != nil {
				t = *t1
			}

			metrics.totalBytes += arrayLength
			metrics.success++
			sc.data <- &Item{
				id:      *pg.GetId(),
				data:    io.NopCloser(bytes.NewReader(byteArray)),
				info:    sharePointPageInfo(pg, arrayLength),
				modTime: t,
			}

			progress <- struct{}{}
		}
	}

	return numMetrics{}, nil
}

func serializeContent(writer *kw.JsonSerializationWriter, obj absser.Parsable) ([]byte, error) {
	defer writer.Close()

	err := writer.WriteObjectValue("", obj)
	if err != nil {
		return nil, err
	}

	byteArray, err := writer.GetSerializedContent()
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}
