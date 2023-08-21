package drive

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	urlCacheDriveItemThreshold = 300 * 1000
	urlCacheRefreshInterval    = 1 * time.Hour
)

type getItemPropertyer interface {
	getItemProperties(
		ctx context.Context,
		itemID string,
	) (itemProps, error)
}

type itemProps struct {
	downloadURL string
	isDeleted   bool
}

var _ getItemPropertyer = &urlCache{}

// urlCache caches download URLs for drive items
type urlCache struct {
	driveID         string
	prevDelta       string
	idToProps       map[string]itemProps
	lastRefreshTime time.Time
	refreshInterval time.Duration
	// cacheMu protects idToProps and lastRefreshTime
	cacheMu sync.RWMutex
	// refreshMu serializes cache refresh attempts by potential writers
	refreshMu       sync.Mutex
	deltaQueryCount int

	itemPager api.DriveItemDeltaEnumerator

	errs *fault.Bus
}

// newURLache creates a new URL cache for the specified drive ID
func newURLCache(
	driveID, prevDelta string,
	refreshInterval time.Duration,
	itemPager api.DriveItemDeltaEnumerator,
	errs *fault.Bus,
) (*urlCache, error) {
	err := validateCacheParams(
		driveID,
		refreshInterval,
		itemPager)
	if err != nil {
		return nil, clues.Wrap(err, "cache params")
	}

	return &urlCache{
			idToProps:       make(map[string]itemProps),
			lastRefreshTime: time.Time{},
			driveID:         driveID,
			prevDelta:       prevDelta,
			refreshInterval: refreshInterval,
			itemPager:       itemPager,
			errs:            errs,
		},
		nil
}

// validateCacheParams validates input params
func validateCacheParams(
	driveID string,
	refreshInterval time.Duration,
	itemPager api.DriveItemDeltaEnumerator,
) error {
	if len(driveID) == 0 {
		return clues.New("drive id is empty")
	}

	if refreshInterval < 1*time.Second {
		return clues.New("invalid refresh interval")
	}

	if itemPager == nil {
		return clues.New("nil item pager")
	}

	return nil
}

// getItemProps returns the item properties for the specified drive item ID
func (uc *urlCache) getItemProperties(
	ctx context.Context,
	itemID string,
) (itemProps, error) {
	if len(itemID) == 0 {
		return itemProps{}, clues.New("item id is empty")
	}

	ctx = clues.Add(ctx, "drive_id", uc.driveID)

	if uc.needsRefresh() {
		err := uc.RefreshCache(ctx)
		if err != nil {
			return itemProps{}, err
		}
	}

	props, err := uc.readCache(ctx, itemID)
	if err != nil {
		return itemProps{}, err
	}

	return props, nil
}

// needsRefresh returns true if the cache is empty or if refresh interval has
// elapsed
func (uc *urlCache) needsRefresh() bool {
	uc.cacheMu.RLock()
	defer uc.cacheMu.RUnlock()

	return len(uc.idToProps) == 0 ||
		time.Since(uc.lastRefreshTime) > uc.refreshInterval
}

// refreshCache refreshes the URL cache by performing a delta query.
func (uc *urlCache) RefreshCache(
	ctx context.Context,
) error {
	// Acquire mutex to prevent multiple threads from refreshing the
	// cache at the same time
	uc.refreshMu.Lock()
	defer uc.refreshMu.Unlock()

	// If the cache was refreshed by another thread while we were waiting
	// to acquire mutex, return
	if !uc.needsRefresh() {
		return nil
	}

	// Hold cache lock in write mode for the entire duration of the refresh.
	// This is to prevent other threads from reading the cache while it is
	// being updated page by page
	uc.cacheMu.Lock()
	defer uc.cacheMu.Unlock()

	// Issue a delta query to graph
	logger.Ctx(ctx).Info("refreshing url cache")

	err := uc.deltaQuery(ctx)
	if err != nil {
		// clear cache
		uc.idToProps = make(map[string]itemProps)

		return err
	}

	logger.Ctx(ctx).Info("url cache refreshed")

	// Update last refresh time
	uc.lastRefreshTime = time.Now()

	return nil
}

// deltaQuery performs a delta query on the drive and update the cache
func (uc *urlCache) deltaQuery(
	ctx context.Context,
) error {
	logger.Ctx(ctx).Debug("starting delta query")
	// Reset item pager to remove any previous state
	uc.itemPager.Reset()

	_, _, _, err := collectItems(
		ctx,
		uc.itemPager,
		uc.driveID,
		"",
		uc.updateCache,
		map[string]string{},
		uc.prevDelta,
		uc.errs)
	if err != nil {
		return clues.Wrap(err, "delta query")
	}

	uc.deltaQueryCount++

	return nil
}

// readCache returns the item properties for the specified item
func (uc *urlCache) readCache(
	ctx context.Context,
	itemID string,
) (itemProps, error) {
	uc.cacheMu.RLock()
	defer uc.cacheMu.RUnlock()

	ctx = clues.Add(ctx, "item_id", itemID)

	props, ok := uc.idToProps[itemID]
	if !ok {
		return itemProps{}, clues.New("item not found in cache").WithClues(ctx)
	}

	return props, nil
}

// updateCache consumes a slice of drive items and updates the url cache.
// It assumes that cacheMu is held by caller in write mode
func (uc *urlCache) updateCache(
	ctx context.Context,
	_, _ string,
	items []models.DriveItemable,
	_ map[string]string,
	_ map[string]string,
	_ map[string]struct{},
	_ map[string]map[string]string,
	_ bool,
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

		var (
			url string
			ad  = item.GetAdditionalData()
		)

		for _, key := range downloadURLKeys {
			if v, err := str.AnyValueToString(key, ad); err == nil {
				url = v
				break
			}
		}

		itemID := ptr.Val(item.GetId())

		uc.idToProps[itemID] = itemProps{
			downloadURL: url,
			isDeleted:   false,
		}

		// Mark deleted items in cache
		if item.GetDeleted() != nil {
			uc.idToProps[itemID] = itemProps{
				downloadURL: "",
				isDeleted:   true,
			}
		}
	}

	return el.Failure()
}
