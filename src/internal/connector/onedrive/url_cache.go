package onedrive

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type itemProperties struct {
	downloadURL string
	isDeleted   bool
}

// urlCache caches download URLs for drive items
type urlCache struct {
	driveID string
	// urlMap stores Item ID -> item property map
	urlMap          map[string]itemProperties
	lastRefreshTime time.Time
	refreshInterval time.Duration
	// cacheLock protects urlMap and lastRefreshTime
	cacheLock sync.RWMutex
	// refreshMutex serializes cache refresh attempts
	refreshMutex    sync.Mutex
	deltaQueryCount int

	driveEnumerator driveEnumeratorFunc
	svc             graph.Servicer
	itemPagerFunc   func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager

	errors *fault.Bus
}

// driveEnumeratorFunc enumerates all items in the specified drive and hands
// them to the provided collector method
type driveEnumeratorFunc func(
	ctx context.Context,
	pager itemPager,
	collector collectorFunc,
	prevDelta string,
	errs *fault.Bus,
) error

// collectorFunc is a callback function that is called by driveEnumeratorFunc
// for each page of items
type collectorFunc func(
	ctx context.Context,
	items []models.DriveItemable,
	errs *fault.Bus,
) error

// newURLache creates a new URL cache for the specified drive ID
func newURLCache(
	driveID string,
	refreshInterval time.Duration,
	driveEnumerator driveEnumeratorFunc,
	svc graph.Servicer,
	itemPagerFunc func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager,
) (*urlCache, error) {
	err := validateCacheParams(driveID, refreshInterval, driveEnumerator, svc, itemPagerFunc)
	if err != nil {
		return nil, clues.Wrap(err, "invalid cache parameters")
	}

	return &urlCache{
			urlMap:          make(map[string]itemProperties),
			lastRefreshTime: time.Time{},
			driveID:         driveID,
			refreshInterval: refreshInterval,
			driveEnumerator: driveEnumerator,
			svc:             svc,
			itemPagerFunc:   itemPagerFunc,
			errors:          fault.New(false),
		},
		nil
}

// validateCacheParams validates the parameters passed to newURLCache
func validateCacheParams(
	driveID string,
	refreshInterval time.Duration,
	driveEnumerator driveEnumeratorFunc,
	svc graph.Servicer,
	itemPagerFunc func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager,
) error {
	if len(driveID) == 0 {
		return clues.New("drive id is empty")
	}

	if refreshInterval < 0 {
		return clues.New("invalid refresh interval")
	}

	if driveEnumerator == nil {
		return clues.New("nil drive enumerator")
	}

	if svc == nil {
		return clues.New("nil graph servicer")
	}

	if itemPagerFunc == nil {
		return clues.New("nil item pager")
	}

	return nil
}

// getItemProps returns the item properties for the specified drive item ID
func (uc *urlCache) getItemProperties(
	ctx context.Context,
	itemID string,
) (*itemProperties, error) {
	if len(itemID) == 0 {
		return nil, clues.New("item id is empty")
	}

	ctx = clues.Add(ctx, "drive_id", uc.driveID)

	// Lazy refresh
	if uc.needsRefresh() {
		err := uc.refreshCache(ctx)
		if err != nil {
			return nil, err
		}
	}

	url, err := uc.readCache(ctx, itemID)
	if err != nil {
		return nil, err
	}

	return url, nil
}

// needsRefresh returns true if the cache is empty or if refresh interval has
// elapsed
func (uc *urlCache) needsRefresh() bool {
	uc.cacheLock.RLock()
	defer uc.cacheLock.RUnlock()

	return len(uc.urlMap) == 0 || time.Since(uc.lastRefreshTime) > uc.refreshInterval
}

// refreshCache refreshes the URL cache by performing a delta query.
func (uc *urlCache) refreshCache(
	ctx context.Context,
) error {
	// Acquire mutex to prevent multiple threads from refreshing the
	// cache at the same time
	uc.refreshMutex.Lock()
	defer uc.refreshMutex.Unlock()

	// If the cache was refreshed by another thread while we were waiting
	// to acquire mutex, return
	if !uc.needsRefresh() {
		return nil
	}

	// Hold cache lock in write mode for the entire duration of the refresh.
	// This is to prevent other threads from reading the cache while it is
	// being updated page by page
	uc.cacheLock.Lock()
	defer uc.cacheLock.Unlock()

	// Issue a delta query to graph
	logger.Ctx(ctx).Debugw("refreshing url cache")

	err := uc.deltaQuery(ctx)
	if err != nil {
		return err
	}

	logger.Ctx(ctx).Debugw("url cache refresh complete")
	// Update last refresh time
	uc.lastRefreshTime = time.Now()

	return nil
}

// deltaQuery will perform a delta query on the drive. updateCache will be
// called for each page of items returned by the delta query and the cache will
// be updated with the results
func (uc *urlCache) deltaQuery(
	ctx context.Context,
) error {
	driveEnumerator := uc.driveEnumerator

	logger.Ctx(ctx).Debugw("Starting delta query")

	err := driveEnumerator(
		ctx,
		uc.itemPagerFunc(uc.svc, uc.driveID, ""),
		uc.updateCache,
		"",
		uc.errors,
	)
	if err != nil {
		return clues.Wrap(err, "delta query failed").WithClues(ctx)
	}

	uc.deltaQueryCount++

	return nil
}

// collectDriveItems will enumerate all items in the specified drive and hand
// them to the provided collector method
// TODO: This is a clone of collectItems call. Refactor collectItems to remove
// duplication
func collectDriveItems(
	ctx context.Context,
	pager itemPager,
	collector collectorFunc,
	prevDelta string,
	errs *fault.Bus,
) error {
	invalidPrevDelta := len(prevDelta) == 0

	if !invalidPrevDelta {
		pager.SetNext(prevDelta)
	}

	for {
		// assume delta urls here, which allows single-token consumption
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))

		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

			pager.Reset()

			continue
		}

		if err != nil {
			return graph.Wrap(ctx, err, "getting page")
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return graph.Wrap(ctx, err, "extracting items from response")
		}

		err = collector(ctx, vals, errs)
		if err != nil {
			return graph.Wrap(ctx, err, "collecting items")
		}

		nextLink, _ := api.NextAndDeltaLink(page)

		// Check if there are more items
		if len(nextLink) == 0 {
			break
		}

		logger.Ctx(ctx).Debugw("Found nextLink", "link", nextLink)
		pager.SetNext(nextLink)
	}

	return nil
}

// readCache returns the download URL for the specified item
func (uc *urlCache) readCache(
	ctx context.Context,
	itemID string,
) (*itemProperties, error) {
	uc.cacheLock.RLock()
	defer uc.cacheLock.RUnlock()

	ctx = clues.Add(ctx, "item_id", itemID)

	itemProps, ok := uc.urlMap[itemID]
	if !ok {
		return nil, clues.New("item not found in cache").WithClues(ctx)
	}

	return &itemProps, nil
}

// updateCache consumes a slice of drive items and updates the url cache.
// It assumes that cacheLock is held by caller in write mode
func (uc *urlCache) updateCache(
	ctx context.Context,
	items []models.DriveItemable,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for _, item := range items {
		if el.Failure() != nil {
			break
		}

		// Skip if not a file
		if item.GetFile() == nil {
			continue
		}

		var url string

		for _, key := range downloadURLKeys {
			tmp, ok := item.GetAdditionalData()[key].(*string)
			if ok {
				url = ptr.Val(tmp)
				break
			}
		}

		itemID := ptr.Val(item.GetId())

		uc.urlMap[itemID] = itemProperties{
			downloadURL: url,
			isDeleted:   false,
		}

		// Mark deleted items in cache
		if item.GetDeleted() != nil {
			uc.urlMap[itemID] = itemProperties{
				downloadURL: "",
				isDeleted:   true,
			}
		}
	}

	return el.Failure()
}
