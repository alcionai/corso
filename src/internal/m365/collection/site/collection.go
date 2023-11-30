package site

import (
	"bytes"
	"context"
	"io"
	"sync"
	"sync/atomic"

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

// Collection is the SharePoint.List or SharePoint.Page implementation of data.Collection.

// SharePoint.Libraries collections are supported by the oneDrive.Collection
// as the calls are identical for populating the Collection
type Collection struct {
	// stream is the container for each individual SharePoint item of (page/list)
	stream chan data.Item

	// fullPath indicates the hierarchy within the collection
	fullPath path.Path

	// jobs contain the SharePoint.List.IDs or SharePoint.Page.IDs
	jobs []string

	// M365 IDs of the items of this collection
	category      path.CategoryType
	client        api.Sites
	ctrl          control.Options
	betaService   *betaAPI.BetaService
	statusUpdater support.StatusUpdater
	getter        getItemByIDer
}

// NewCollection helper function for creating a Collection
func NewCollection(
	getter getItemByIDer,
	folderPath path.Path,
	ac api.Client,
	scope selectors.SharePointScope,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *Collection {
	c := &Collection{
		fullPath:      folderPath,
		jobs:          make([]string, 0),
		getter:        getter,
		stream:        make(chan data.Item, collectionChannelBufferSize),
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

func (sc *Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	go sc.streamItems(ctx, errs)
	return sc.stream
}

func (sc *Collection) finishPopulation(ctx context.Context, metrics support.CollectionMetrics) {
	close(sc.stream)

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

// streamItems utility function to retrieve data from back store for a given collection
func (sc *Collection) streamItems(ctx context.Context, errs *fault.Bus) {
	// Switch retrieval function based on category
	switch sc.category {
	case path.ListsCategory:
		sc.streamLists(ctx, errs)
	case path.PagesCategory:
		sc.retrievePages(ctx, sc.client, errs)
	}
}

// streamLists utility function for collection that fetches and serializes
// models.Listable objects based on M365 IDs from the jobs field.
func (sc *Collection) streamLists(ctx context.Context, errs *fault.Bus) {
	var (
		metrics support.CollectionMetrics
		objects int64
		wg      sync.WaitGroup
		el      = errs.Local()
	)

	defer sc.finishPopulation(ctx, metrics)

	// TODO: Insert correct ID for CollectionProgress
	colProgress := observe.CollectionProgress(ctx, sc.fullPath.Category().HumanString(), sc.fullPath.Folders())
	defer close(colProgress)

	semaphoreCh := make(chan struct{}, fetchChannelSize)
	defer close(semaphoreCh)

	for _, listID := range sc.jobs {
		if el.Failure() != nil {
			break
		}

		wg.Add(1)
		semaphoreCh <- struct{}{}

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			var (
				entry models.Listable
				err   error
			)

			entry, err = sc.getter.GetItemByID(ctx, id)
			if err != nil {
				err = clues.Wrap(err, "getting list data").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			atomic.AddInt64(&objects, 1)

			if err := writer.WriteObjectValue("", entry); err != nil {
				err = clues.Wrap(err, "writing list to serializer").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			entryBytes, err := writer.GetSerializedContent()
			if err != nil {
				err = clues.Wrap(err, "serializing list").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			size := int64(len(entryBytes))

			if size > 0 {
				metrics.Bytes += size
				metrics.Successes++

				rc := io.NopCloser(bytes.NewReader(entryBytes))
				itemID := ptr.Val(entry.GetId())
				itemInfo := details.ItemInfo{SharePoint: ListToSPInfo(entry, size)}

				item, err := data.NewPrefetchedItemWithInfo(rc, itemID, itemInfo)
				if err != nil {
					el.AddRecoverable(ctx, clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation))

					return
				}

				sc.stream <- item
				colProgress <- struct{}{}
			}
		}(listID)
	}

	wg.Wait()

	metrics.Objects += int(objects)
}

func (sc *Collection) retrievePages(ctx context.Context, as api.Sites, errs *fault.Bus) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
	)

	defer sc.finishPopulation(ctx, metrics)

	// TODO: Insert correct ID for CollectionProgress
	colProgress := observe.CollectionProgress(ctx, sc.fullPath.Category().HumanString(), sc.fullPath.Folders())
	defer close(colProgress)

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	betaService := sc.betaService
	if betaService == nil {
		logger.Ctx(ctx).Error(clues.New("beta service required"))

		return
	}

	parent, err := as.GetByID(ctx, sc.fullPath.ProtectedResource(), api.CallConfig{})
	if err != nil {
		logger.Ctx(ctx).Error(err)

		return
	}

	root := ptr.Val(parent.GetWebUrl())

	pages, err := betaAPI.GetSitePages(ctx, betaService, sc.fullPath.ProtectedResource(), sc.jobs, errs)
	if err != nil {
		logger.Ctx(ctx).Error(err)

		return
	}

	metrics.Objects = len(pages)
	// For each models.Pageable, object is serialize and the metrics are collected and returned.
	// Pageable objects are not supported in v1.0 of msgraph at this time.
	// TODO: Verify Parsable interface supported with modified-Pageable
	for _, pg := range pages {
		if el.Failure() != nil {
			break
		}

		byteArray, err := serializeContent(ctx, writer, pg)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "serializing page").Label(fault.LabelForceNoBackupCreation))
			continue
		}

		size := int64(len(byteArray))

		if size > 0 {
			metrics.Bytes += size
			metrics.Successes++

			rc := io.NopCloser(bytes.NewReader(byteArray))
			itemID := ptr.Val(pg.GetId())
			itemInfo := details.ItemInfo{SharePoint: pageToSPInfo(pg, root, size)}

			item, err := data.NewPrefetchedItemWithInfo(rc, itemID, itemInfo)
			if err != nil {
				el.AddRecoverable(ctx, clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation))

				continue
			}

			sc.stream <- item
			colProgress <- struct{}{}
		}
	}
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
