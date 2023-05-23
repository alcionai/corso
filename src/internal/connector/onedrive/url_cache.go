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
	// urlMap stores Item ID -> download URL map
	urlMap          map[string]itemProperties
	lastRefreshTime time.Time
	refreshInterval time.Duration
	// rwLock protects urlMap and lastRefreshTime
	rwLock sync.RWMutex
	// refreshMutex serializes concurrent cache refreshes
	refreshMutex    sync.Mutex
	deltaQueryCount int
	// TODO: Handle error bus properly
	Errors *fault.Bus

	// TODO: Couple two below together
	driveEnumerator driveEnumeratorFunc
	itemPagerFunc   func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager
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

// newURLache creates a new URL cache for the specified drive
// TODO: move graph servicer to cache
func newURLCache(
	driveID string,
	refreshInterval time.Duration,
	driveEnumerator driveEnumeratorFunc,
	itemPagerFunc func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager,
) *urlCache {
	return &urlCache{
		urlMap:          make(map[string]itemProperties),
		lastRefreshTime: time.Time{},
		driveID:         driveID,
		refreshInterval: refreshInterval,
		driveEnumerator: driveEnumerator,
		itemPagerFunc:   itemPagerFunc,
		refreshMutex:    sync.Mutex{},
	}
}

// getDownloadUrl returns the download URL for the specified drive item
// TODO: Any cache error should not be treated as a fatal error by client
// TODO: How to convey deleted item to client?
// TODO: Move graph.Servicer to urlCache struct?
// TODO: Add info logs
func (uc *urlCache) getDownloadURL(
	ctx context.Context,
	svc graph.Servicer,
	itemID string,
) (string, error) {
	if len(itemID) == 0 {
		return "", clues.New("item id is empty")
	}

	if uc.needsRefresh() {
		err := uc.refreshCache(ctx, svc)
		if err != nil {
			return "", err
		}
	}

	url, err := uc.readCache(ctx, itemID)
	if err != nil {
		return "", err
	}

	return url, nil
}

// needsRefresh returns true if the cache is empty or if refresh interval has
// elapsed
func (uc *urlCache) needsRefresh() bool {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()

	return len(uc.urlMap) == 0 || time.Since(uc.lastRefreshTime) > uc.refreshInterval
}

// refreshCache refreshes the URL cache by performing a delta query.
func (uc *urlCache) refreshCache(
	ctx context.Context,
	svc graph.Servicer,
) error {
	// Acquire binary semaphore to prevent multiple threads from refreshing the
	// cache at the same time
	uc.refreshMutex.Lock()
	defer uc.refreshMutex.Unlock()

	// If the cache was refreshed by another thread while we were waiting
	// to acquire semaphore, return
	if !uc.needsRefresh() {
		return nil
	}

	// Hold lock in write mode for the entire duration of the refresh.
	// This is to prevent other threads from reading the cache while it is
	// being updated
	uc.rwLock.Lock()
	defer uc.rwLock.Unlock()

	// Issue a delta query to graph
	err := uc.deltaQuery(ctx, svc)
	if err != nil {
		return err
	}

	// Update last refresh time
	uc.lastRefreshTime = time.Now()

	return nil
}

// deltaQuery will perform a delta query on the drive and update the cache
// TODO: Check if this function is adding any value?
// Remove it and use collectDriveItems directly
func (uc *urlCache) deltaQuery(
	ctx context.Context,
	svc graph.Servicer,
) error {
	ictx := clues.Add(ctx, "drive_id", uc.driveID)

	driveEnumerator := uc.driveEnumerator
	if driveEnumerator == nil {
		driveEnumerator = collectDriveItems
	}

	err := driveEnumerator(
		ctx,
		uc.itemPagerFunc(svc, uc.driveID, ""),
		uc.updateCache,
		"",
		uc.Errors,
	)
	if err != nil {
		return clues.Wrap(err, "delta query failed").WithClues(ictx)
	}

	uc.deltaQueryCount++

	return nil
}

// collectDriveItems will enumerate all items in the specified drive and hand
// them to the provided `collector` method
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
) (string, error) {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()

	ictx := clues.Add(ctx, "item_id", itemID)

	val, ok := uc.urlMap[itemID]
	if !ok {
		return "", clues.New("item not found in cache").WithClues(ictx)
	}

	if val.isDeleted {
		// TODO: standardize error
		return "", clues.New("item is deleted").WithClues(ictx)
	}

	return val.downloadURL, nil
}

// updateCache will cache the download URL for each item
// Assumes that rwLock is held by caller in write mode
// TODO: Add debug logs and more error handling
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
