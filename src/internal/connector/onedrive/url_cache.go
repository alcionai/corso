package onedrive

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	gapi "github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

type itemProperties struct {
	downloadURL string
	isDeleted   bool
	// temporary
	driveItem *models.DriveItem
}

// urlCache caches download URLs for drive items
type urlCache struct {
	// Item ID -> download URL map
	urlMap map[string]itemProperties
	// time of last cache  refresh
	lastRefreshTime time.Time
	// RW lock for the URL map and lastRefreshTime
	rwLock sync.RWMutex
	// semaphore for limiting the number of concurrent cache refreshes to 1
	refreshSemaphore chan struct{}

	driveID         string
	deltaQueryCount int
	Errors          *fault.Bus

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

// updateCacheFunc is a callback function that is called for each page of items
type collectorFunc func(
	ctx context.Context,
	items []models.DriveItemable,
	errs *fault.Bus,
) error

// NewUrlCache creates a new URL cache for the specified drive
func newURLCache(
	driveID, driveName string,
	driveEnumerator driveEnumeratorFunc,
	itemPagerFunc func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager,
) *urlCache {
	return &urlCache{
		urlMap: make(map[string]itemProperties),
		// TODO: use a mutex instead of a semaphore since it's
		// size is just 1?
		refreshSemaphore: make(chan struct{}, 1),
		driveID:          driveID,
		driveEnumerator:  driveEnumerator,
		itemPagerFunc:    itemPagerFunc,
	}
}

// getDownloadUrl returns the download URL for the specified item
// TODO: Any cache error should not be treated as a fatal error by client
// TODO: How to convey deleted item to client?
// TODO: bail on 3 consecutive errors
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

	url, err := uc.readCache(itemID)
	if err != nil {
		return "", err
	}

	return url, nil
}

// needsRefresh returns true if the cache is empty or if > 1 hr has elapsed since
// last refresh
// TODO: make it 55 mins to avoid 401s from possibly stale cache hits?
func (uc *urlCache) needsRefresh() bool {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()

	return len(uc.urlMap) == 0 || time.Since(uc.lastRefreshTime) > time.Hour
}

// refreshCache refreshes the URL cache by performing a delta query.
// It allows only one concurrent refresh at a time.
func (uc *urlCache) refreshCache(
	ctx context.Context,
	svc graph.Servicer,
) error {
	// semaphore to limit the number of concurrent cache refreshes to 1
	if uc.refreshSemaphore == nil {
		return clues.New("refresh semaphore is nil")
	}

	uc.refreshSemaphore <- struct{}{}
	defer func() { <-uc.refreshSemaphore }()

	// If the cache was refreshed by another thread while we were waiting
	// to acquire semaphore, return
	if !uc.needsRefresh() {
		return nil
	}

	err := uc.deltaQuery(ctx, svc)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Check if this function is adding any value?
// Remove it and use collectDriveItems directly
func (uc *urlCache) deltaQuery(
	ctx context.Context,
	svc graph.Servicer,
) error {
	ictx := clues.Add(ctx, "drive_id", uc.driveID)

	driveEnumerator := uc.driveEnumerator
	if driveEnumerator == nil {
		driveEnumerator = uc.collectDriveItems
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

	return nil
}

// collectDriveItems will enumerate all items in the specified drive and hand
// them to the provided `collector` method
func (uc *urlCache) collectDriveItems(
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

		nextLink, _ := gapi.NextAndDeltaLink(page)

		// Check if there are more items
		if len(nextLink) == 0 {
			break
		}

		logger.Ctx(ctx).Debugw("Found nextLink", "link", nextLink)
		pager.SetNext(nextLink)
	}

	return nil
}

// getFromCache returns the download URL for the specified item
func (uc *urlCache) readCache(itemID string) (string, error) {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()

	val, ok := uc.urlMap[itemID]
	if !ok {
		// TODO: improve clues
		return "", clues.New("item not found in cache")
	}

	if val.isDeleted {
		// TODO: standardize error
		return "", clues.New("item is deleted")
	}

	return val.downloadURL, nil
}

// updateCache is a callback function that is called for each page of items
// It will cache the download URL for each item
// TODO: Add debug logs and more error handling
func (uc *urlCache) updateCache(
	ctx context.Context,
	items []models.DriveItemable,
	errs *fault.Bus,
) error {
	uc.rwLock.Lock()
	defer uc.rwLock.Unlock()

	el := errs.Local()

	for _, item := range items {
		if el.Failure() != nil {
			break
		}

		// Skip if not a file
		if item.GetFile() != nil {
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
			driveItem:   item.(*models.DriveItem),
		}

		// Mark deleted items in cache
		if item.GetDeleted() != nil {
			uc.urlMap[itemID] = itemProperties{
				downloadURL: "",
				isDeleted:   true,
				driveItem:   item.(*models.DriveItem),
			}
		}
	}

	return el.Failure()
}
