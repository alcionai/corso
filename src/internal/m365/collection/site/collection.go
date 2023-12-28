package site

import (
	"bytes"
	"context"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

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

var _ data.BackupCollection = &prefetchCollection{}

// Collection is the SharePoint.List or SharePoint.Page implementation of data.Collection.

// SharePoint.Libraries collections are supported by the oneDrive.Collection
// as the calls are identical for populating the Collection
type prefetchCollection struct {
	// stream is the container for each individual SharePoint item of (page/list)
	stream map[path.CategoryType]chan data.Item
	// fullPath indicates the hierarchy within the collection
	fullPath path.Path
	// items contains the SharePoint.List.IDs or SharePoint.Page.IDs
	// and their corresponding last modified time
	items map[string]time.Time
	// M365 IDs of the items of this collection
	category      path.CategoryType
	client        api.Sites
	ctrl          control.Options
	betaService   *betaAPI.BetaService
	statusUpdater support.StatusUpdater
	getter        getItemByIDer
}

// NewPrefetchCollection constructor function for creating a prefetchCollection
func NewPrefetchCollection(
	getter getItemByIDer,
	folderPath path.Path,
	ac api.Client,
	scope selectors.SharePointScope,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
) *prefetchCollection {
	c := &prefetchCollection{
		fullPath:      folderPath,
		items:         make(map[string]time.Time),
		getter:        getter,
		stream:        make(map[path.CategoryType]chan data.Item),
		client:        ac.Sites(),
		statusUpdater: statusUpdater,
		category:      scope.Category().PathType(),
		ctrl:          ctrlOpts,
	}

	return c
}

func (pc *prefetchCollection) SetBetaService(betaService *betaAPI.BetaService) {
	pc.betaService = betaService
}

// AddItem appends additional itemID to items field
func (pc *prefetchCollection) AddItem(itemID string, lastModifedTime time.Time) {
	pc.items[itemID] = lastModifedTime
}

func (pc *prefetchCollection) FullPath() path.Path {
	return pc.fullPath
}

// TODO(ashmrtn): Fill in with previous path once the Controller compares old
// and new folder hierarchies.
func (pc prefetchCollection) PreviousPath() path.Path {
	return nil
}

func (pc prefetchCollection) LocationPath() *path.Builder {
	return path.Builder{}.Append(pc.fullPath.Folders()...)
}

func (pc prefetchCollection) State() data.CollectionState {
	return data.NewState
}

func (pc prefetchCollection) DoNotMergeItems() bool {
	return false
}

func (pc *prefetchCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	if _, ok := pc.stream[pc.category]; !ok {
		pc.stream[pc.category] = make(chan data.Item, collectionChannelBufferSize)
	}

	go pc.streamItems(ctx, errs)

	return pc.stream[pc.category]
}

// streamItems utility function to retrieve data from back store for a given collection
func (pc *prefetchCollection) streamItems(
	ctx context.Context,
	errs *fault.Bus,
) {
	// Switch retrieval function based on category
	switch pc.category {
	case path.ListsCategory:
		pc.streamLists(ctx, errs)
	case path.PagesCategory:
		pc.streamPages(ctx, pc.client, errs)
	}
}

// streamLists utility function for collection that downloads and serializes
// models.Listable objects based on M365 IDs from the jobs field.
func (pc *prefetchCollection) streamLists(
	ctx context.Context,
	errs *fault.Bus,
) {
	var (
		metrics         support.CollectionMetrics
		el              = errs.Local()
		wg              sync.WaitGroup
		objects         int64
		objectBytes     int64
		objectSuccesses int64
	)

	defer updateStatus(
		ctx,
		pc.stream[path.ListsCategory],
		pc.statusUpdater,
		pc.fullPath,
		&metrics)

	// TODO: Insert correct ID for CollectionProgress
	progress := observe.CollectionProgress(ctx, pc.fullPath.Category().HumanString(), pc.fullPath.Folders())
	defer close(progress)

	semaphoreCh := make(chan struct{}, fetchChannelSize)
	defer close(semaphoreCh)

	// For each models.Listable, object is serialized and the metrics are collected.
	// The progress is objected via the passed in channel.
	for listID := range pc.items {
		if el.Failure() != nil {
			break
		}

		wg.Add(1)
		semaphoreCh <- struct{}{}

		go pc.handleListItems(
			ctx,
			semaphoreCh,
			progress,
			&wg,
			listID,
			&objects,
			&objectBytes,
			&objectSuccesses,
			el)
	}

	wg.Wait()

	metrics.Objects = int(objects)
	metrics.Bytes = objectBytes
	metrics.Successes = int(objectSuccesses)
}

func (pc *prefetchCollection) streamPages(
	ctx context.Context,
	as api.Sites,
	errs *fault.Bus,
) {
	var (
		metrics support.CollectionMetrics
		el      = errs.Local()
	)

	defer updateStatus(
		ctx,
		pc.stream[path.PagesCategory],
		pc.statusUpdater,
		pc.fullPath,
		&metrics)

	// TODO: Insert correct ID for CollectionProgress
	progress := observe.CollectionProgress(ctx, pc.fullPath.Category().HumanString(), pc.fullPath.Folders())
	defer close(progress)

	betaService := pc.betaService
	if betaService == nil {
		logger.Ctx(ctx).Error(clues.New("beta service required"))
		return
	}

	parent, err := as.GetByID(ctx, pc.fullPath.ProtectedResource(), api.CallConfig{})
	if err != nil {
		logger.Ctx(ctx).Error(err)

		return
	}

	root := ptr.Val(parent.GetWebUrl())

	pageIDs := maps.Keys(pc.items)

	pages, err := betaAPI.GetSitePages(ctx, betaService, pc.fullPath.ProtectedResource(), pageIDs, errs)
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

		byteArray, err := serializeContent(ctx, pg)
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

		pc.stream[path.PagesCategory] <- item
		progress <- struct{}{}
	}
}

func (pc *prefetchCollection) handleListItems(
	ctx context.Context,
	semaphoreCh chan struct{},
	progress chan<- struct{},
	wg *sync.WaitGroup,
	listID string,
	objects *int64,
	objectBytes *int64,
	objectSuccesses *int64,
	el *fault.Bus,
) {
	defer wg.Done()
	defer func() { <-semaphoreCh }()

	var (
		list models.Listable
		info *details.SharePointInfo
		err  error
	)

	list, info, err = pc.getter.GetItemByID(ctx, listID)
	if err != nil {
		err = clues.WrapWC(ctx, err, "getting list data").Label(fault.LabelForceNoBackupCreation)
		el.AddRecoverable(ctx, err)

		return
	}

	atomic.AddInt64(objects, 1)

	entryBytes, err := serializeContent(ctx, list)
	if err != nil {
		el.AddRecoverable(ctx, err)
		return
	}

	size := int64(len(entryBytes))

	if size == 0 {
		return
	}

	atomic.AddInt64(objectBytes, size)
	atomic.AddInt64(objectSuccesses, 1)

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

	pc.stream[path.ListsCategory] <- item
	progress <- struct{}{}
}

func serializeContent(
	ctx context.Context,
	obj serialization.Parsable,
) ([]byte, error) {
	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	err := writer.WriteObjectValue("", obj)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "writing to serializer").Label(fault.LabelForceNoBackupCreation)
	}

	byteArray, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting content from writer").Label(fault.LabelForceNoBackupCreation)
	}

	return byteArray, nil
}

func updateStatus(
	ctx context.Context,
	stream chan data.Item,
	su support.StatusUpdater,
	fullPath path.Path,
	metrics *support.CollectionMetrics,
) {
	close(stream)

	status := support.CreateStatus(
		ctx,
		support.Backup,
		1, // 1 folder
		*metrics,
		fullPath.Folder(false))

	logger.Ctx(ctx).Debug(status.String())

	if su != nil {
		su(status)
	}
}
