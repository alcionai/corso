package site

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type DataCategory int

// channel sizes
const (
	collectionChannelBufferSize = 50
	fetchChannelSize            = 5
)

//go:generate stringer -type=DataCategory
const (
	Unknown DataCategory = 0
	List    DataCategory = 1
	Pages   DataCategory = 2
)

var _ data.BackupCollection = &Collection{}

// Collection is the SharePoint.List implementation of data.Collection. SharePoint.Libraries collections are supported
// by the oneDrive.Collection as the calls are identical for populating the Collection
type Collection struct {
	// data is the container for each individual SharePoint.List
	data chan data.Item
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// jobs contain the SharePoint.Site.ListIDs for the associated list(s).
	jobs []string
	// M365 IDs of the items of this collection
	category      path.CategoryType
	client        api.Sites
	ctrl          control.Options
	betaService   *betaAPI.BetaService
	statusUpdater support.StatusUpdater
}

// NewCollection helper function for creating a Collection
func NewCollection(
	folderPath path.Path,
	ac api.Client,
	scope selectors.SharePointScope,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collection {
	c := &Collection{
		fullPath:      folderPath,
		jobs:          make([]string, 0),
		data:          make(chan data.Item, collectionChannelBufferSize),
		client:        ac.Sites(),
		statusUpdater: statusUpdater,
		category:      scope.Category().PathType(),
		ctrl:          ctrlOpts,
	}

	return c
}

func (sc *Collection) SetBetaService(betaService *betaAPI.BetaService) {
	sc.betaService = betaService
}

// AddJob appends additional objectID to job field
func (sc *Collection) AddJob(objID string) {
	sc.jobs = append(sc.jobs, objID)
}

func (sc *Collection) FullPath() path.Path {
	return sc.fullPath
}

// TODO(ashmrtn): Fill in with previous path once the Controller compares old
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

func (sc *Collection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	go sc.populate(ctx, errs)
	return sc.data
}

func (sc *Collection) finishPopulation(
	ctx context.Context,
	metrics support.CollectionMetrics,
) {
	close(sc.data)

	status := support.CreateStatus(
		ctx,
		support.Backup,
		1, // 1 folder
		metrics,
		sc.fullPath.Folder(false))

	logger.Ctx(ctx).Debug(status.String())

	if sc.statusUpdater != nil {
		sc.statusUpdater(status)
	}
}

// populate utility function to retrieve data from back store for a given collection
func (sc *Collection) populate(ctx context.Context, errs *fault.Bus) {
	metrics, _ := sc.runPopulate(ctx, errs)
	sc.finishPopulation(ctx, metrics)
}

func (sc *Collection) runPopulate(
	ctx context.Context,
	errs *fault.Bus,
) (support.CollectionMetrics, error) {
	var (
		err     error
		metrics support.CollectionMetrics
		writer  = kjson.NewJsonSerializationWriter()
	)

	// TODO: Insert correct ID for CollectionProgress
	colProgress := observe.CollectionProgress(
		ctx,
		sc.fullPath.Category().HumanString(),
		sc.fullPath.Folders())
	defer close(colProgress)

	// Switch retrieval function based on category
	switch sc.category {
	case path.ListsCategory:
		metrics, err = sc.retrieveLists(ctx, writer, colProgress, errs)
	case path.PagesCategory:
		metrics, err = sc.retrievePages(ctx, sc.client, writer, colProgress, errs)
	}

	return metrics, err
}

// retrieveLists utility function for collection that downloads and serializes
// models.Listable objects based on M365 IDs from the jobs field.
func (sc *Collection) retrieveLists(
	ctx context.Context,
	wtr *kjson.JsonSerializationWriter,
	progress chan<- struct{},
	errs *fault.Bus,
) (support.CollectionMetrics, error) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
	)

	// TODO: Fetch lists via Lists client wrapper
	var lists = []models.Listable{}

	metrics.Objects += len(lists)
	// For each models.Listable, object is serialized and the metrics are collected.
	// The progress is objected via the passed in channel.
	for _, lst := range lists {
		if el.Failure() != nil {
			break
		}

		byteArray, err := serializeContent(ctx, wtr, lst)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "serializing list").Label(fault.LabelForceNoBackupCreation))
			continue
		}

		size := int64(len(byteArray))

		if size > 0 {
			metrics.Bytes += size

			metrics.Successes++

			item, err := data.NewPrefetchedItemWithInfo(
				io.NopCloser(bytes.NewReader(byteArray)),
				ptr.Val(lst.GetId()),
				details.ItemInfo{SharePoint: ListToSPInfo(lst, size)})
			if err != nil {
				el.AddRecoverable(ctx, clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			sc.data <- item
			progress <- struct{}{}
		}
	}

	return metrics, el.Failure()
}

func (sc *Collection) retrievePages(
	ctx context.Context,
	as api.Sites,
	wtr *kjson.JsonSerializationWriter,
	progress chan<- struct{},
	errs *fault.Bus,
) (support.CollectionMetrics, error) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
	)

	betaService := sc.betaService
	if betaService == nil {
		return metrics, clues.NewWC(ctx, "beta service required")
	}

	parent, err := as.GetByID(ctx, sc.fullPath.ProtectedResource(), api.CallConfig{})
	if err != nil {
		return metrics, err
	}

	root := ptr.Val(parent.GetWebUrl())

	pages, err := betaAPI.GetSitePages(ctx, betaService, sc.fullPath.ProtectedResource(), sc.jobs, errs)
	if err != nil {
		return metrics, err
	}

	metrics.Objects = len(pages)
	// For each models.Pageable, object is serialize and the metrics are collected and returned.
	// Pageable objects are not supported in v1.0 of msgraph at this time.
	// TODO: Verify Parsable interface supported with modified-Pageable
	for _, pg := range pages {
		if el.Failure() != nil {
			break
		}

		byteArray, err := serializeContent(ctx, wtr, pg)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "serializing page").Label(fault.LabelForceNoBackupCreation))
			continue
		}

		size := int64(len(byteArray))

		if size > 0 {
			metrics.Bytes += size
			metrics.Successes++

			item, err := data.NewPrefetchedItemWithInfo(
				io.NopCloser(bytes.NewReader(byteArray)),
				ptr.Val(pg.GetId()),
				details.ItemInfo{SharePoint: pageToSPInfo(pg, root, size)})
			if err != nil {
				el.AddRecoverable(ctx, clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			sc.data <- item
			progress <- struct{}{}
		}
	}

	return metrics, el.Failure()
}

func serializeContent(
	ctx context.Context,
	writer *kjson.JsonSerializationWriter,
	obj serialization.Parsable,
) ([]byte, error) {
	defer writer.Close()

	err := writer.WriteObjectValue("", obj)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "writing object")
	}

	byteArray, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting content from writer")
	}

	return byteArray, nil
}
