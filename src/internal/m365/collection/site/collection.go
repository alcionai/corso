package site

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

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

var (
	_ data.BackupCollection = &prefetchCollection{}
	_ data.BackupCollection = &lazyFetchCollection{}
)

// Collection is the SharePoint.List or SharePoint.Page implementation of data.Collection.

// SharePoint.Libraries collections are supported by the oneDrive.Collection
// as the calls are identical for populating the Collection
type prefetchCollection struct {
	// stream is the container for each individual SharePoint item of (page/list)
	stream chan data.Item
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// jobs contain the SharePoint.List.IDs or SharePoint.Page.IDs
	items []string
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
) *prefetchCollection {
	c := &prefetchCollection{
		fullPath:      folderPath,
		items:         make([]string, 0),
		getter:        getter,
		stream:        make(chan data.Item, collectionChannelBufferSize),
		client:        ac.Sites(),
		statusUpdater: statusUpdater,
		category:      scope.Category().PathType(),
		ctrl:          ctrlOpts,
	}

	return c
}

func (sc *prefetchCollection) SetBetaService(betaService *betaAPI.BetaService) {
	sc.betaService = betaService
}

// AddItem appends additional itemID to items field
func (sc *prefetchCollection) AddItem(itemID string) {
	sc.items = append(sc.items, itemID)
}

func (sc *prefetchCollection) FullPath() path.Path {
	return sc.fullPath
}

func (sc prefetchCollection) LocationPath() *path.Builder {
	return path.Builder{}.Append(sc.fullPath.Folders()...)
}

// TODO(hitesh): Implement PreviousPath, State, DoNotMergeItems
// once the Controller compares old and new folder hierarchies.
func (sc prefetchCollection) PreviousPath() path.Path {
	return nil
}

func (sc prefetchCollection) State() data.CollectionState {
	return data.NewState
}

func (sc prefetchCollection) DoNotMergeItems() bool {
	return false
}

func (sc *prefetchCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	go sc.streamItems(ctx, errs)
	return sc.stream
}

// streamItems utility function to retrieve data from back store for a given collection
func (sc *prefetchCollection) streamItems(
	ctx context.Context,
	errs *fault.Bus,
) {
	// Switch retrieval function based on category
	switch sc.category {
	case path.ListsCategory:
		sc.streamLists(ctx, errs)
	case path.PagesCategory:
		sc.retrievePages(ctx, sc.client, errs)
	}
}

// streamLists utility function for collection that downloads and serializes
// models.Listable objects based on M365 IDs from the jobs field.
func (sc *prefetchCollection) streamLists(
	ctx context.Context,
	errs *fault.Bus,
) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
		wg      sync.WaitGroup
	)

	defer finishPopulation(
		ctx,
		sc.stream,
		sc.statusUpdater,
		sc.fullPath,
		metrics,
	)

	// TODO: Insert correct ID for CollectionProgress
	progress := observe.CollectionProgress(ctx, sc.fullPath.Category().HumanString(), sc.fullPath.Folders())
	defer close(progress)

	semaphoreCh := make(chan struct{}, fetchChannelSize)
	defer close(semaphoreCh)

	// For each models.Listable, object is serialized and the metrics are collected.
	// The progress is objected via the passed in channel.
	for _, listID := range sc.items {
		if el.Failure() != nil {
			break
		}

		wg.Add(1)
		semaphoreCh <- struct{}{}

		sc.handleListItems(ctx, semaphoreCh, progress, listID, el, &metrics)

		wg.Done()
	}

	wg.Wait()
}

func (sc *prefetchCollection) retrievePages(
	ctx context.Context,
	as api.Sites,
	errs *fault.Bus,
) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
	)

	defer finishPopulation(
		ctx,
		sc.stream,
		sc.statusUpdater,
		sc.fullPath,
		metrics,
	)

	// TODO: Insert correct ID for CollectionProgress
	progress := observe.CollectionProgress(ctx, sc.fullPath.Category().HumanString(), sc.fullPath.Folders())
	defer close(progress)

	wtr := kjson.NewJsonSerializationWriter()
	defer wtr.Close()

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

	pages, err := betaAPI.GetSitePages(ctx, betaService, sc.fullPath.ProtectedResource(), sc.items, errs)
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

		byteArray, err := serializeContent(ctx, wtr, pg)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "serializing page").Label(fault.LabelForceNoBackupCreation))
			continue
		}

		size := int64(len(byteArray))

		if size == 0 {
			return
		}

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

		sc.stream <- item
		progress <- struct{}{}
	}
}

func (sc *prefetchCollection) handleListItems(
	ctx context.Context,
	semaphoreCh chan struct{},
	progress chan<- struct{},
	listID string,
	el *fault.Bus,
	metrics *support.CollectionMetrics,
) {
	defer func() { <-semaphoreCh }()

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	var (
		list models.Listable
		info *details.SharePointInfo
		err  error
	)

	list, info, err = sc.getter.GetItemByID(ctx, listID)
	if err != nil {
		err = clues.WrapWC(ctx, err, "getting list data").Label(fault.LabelForceNoBackupCreation)
		el.AddRecoverable(ctx, err)

		return
	}

	metrics.Objects++

	if err := writer.WriteObjectValue("", list); err != nil {
		err = clues.WrapWC(ctx, err, "writing list to serializer").Label(fault.LabelForceNoBackupCreation)
		el.AddRecoverable(ctx, err)

		return
	}

	entryBytes, err := writer.GetSerializedContent()
	if err != nil {
		err = clues.WrapWC(ctx, err, "serializing list").Label(fault.LabelForceNoBackupCreation)
		el.AddRecoverable(ctx, err)

		return
	}

	size := int64(len(entryBytes))

	if size == 0 {
		return
	}

	metrics.Bytes += size
	metrics.Successes++

	rc := io.NopCloser(bytes.NewReader(entryBytes))
	itemInfo := details.ItemInfo{
		SharePoint: info,
	}

	item, err := data.NewPrefetchedItemWithInfo(rc, listID, itemInfo)
	if err != nil {
		err = clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation)
		el.AddRecoverable(ctx, err)

		return
	}

	sc.stream <- item
	progress <- struct{}{}
}

type lazyFetchCollection struct {
	// stream is the container for each individual SharePoint item of list
	stream chan data.Item
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// jobs contain the SharePoint.List.IDs
	items         []string
	statusUpdater support.StatusUpdater
	getter        getItemByIDer
}

func (lc *lazyFetchCollection) AddItem(itemID string) {
	lc.items = append(lc.items, itemID)
}

func (lc *lazyFetchCollection) FullPath() path.Path {
	return lc.fullPath
}

func (lc lazyFetchCollection) LocationPath() *path.Builder {
	return path.Builder{}.Append(lc.fullPath.Folders()...)
}

// TODO(hitesh): Implement PreviousPath, State, DoNotMergeItems
// once the Controller compares old and new folder hierarchies.
func (lc lazyFetchCollection) PreviousPath() path.Path {
	return nil
}

func (lc lazyFetchCollection) State() data.CollectionState {
	return data.NewState
}

func (lc lazyFetchCollection) DoNotMergeItems() bool {
	return false
}

func (lc lazyFetchCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	go lc.streamItems(ctx, errs)
	return lc.stream
}

func (lc *lazyFetchCollection) streamItems(
	ctx context.Context,
	errs *fault.Bus,
) {
	_ = lc.getter
	_ = lc.handleListItems
}

func (lc *lazyFetchCollection) handleListItems(
	ctx context.Context,
	semaphoreCh chan struct{},
	progress chan<- struct{},
	numLists int64,
	listID string,
	el *fault.Bus,
	metrics support.CollectionMetrics,
) {
	_ = lc.statusUpdater
	lig := &lazyItemGetter{}
	_, _, _, _ = lig.GetData(ctx, el)
}

type lazyItemGetter struct {
	getter       getItemByIDer
	userID       string
	itemID       string
	parentPath   string
	modTime      time.Time
	immutableIDs bool
}

func (lig *lazyItemGetter) GetData(
	ctx context.Context,
	el *fault.Bus,
) (io.ReadCloser, *details.ItemInfo, bool, error) {
	_ = lig.getter
	_ = lig.userID
	_ = lig.itemID
	_ = lig.parentPath
	_ = lig.modTime
	_ = lig.immutableIDs

	return nil, nil, false, nil
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

func finishPopulation(
	ctx context.Context,
	stream chan data.Item,
	su support.StatusUpdater,
	fullPath path.Path,
	metrics support.CollectionMetrics,
) {
	close(stream)

	status := support.CreateStatus(
		ctx,
		support.Backup,
		1, // 1 folder
		metrics,
		fullPath.Folder(false))

	logger.Ctx(ctx).Debug(status.String())

	if su != nil {
		su(status)
	}
}
