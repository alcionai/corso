package onedrive

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type itemProperties struct {
	downloadUrl     string
	lastRefreshTime time.Time
	isDeleted       bool
}

type urlCache struct {
	// Item ID -> download URL map
	urlMap          map[string]itemProperties
	cacheLock       sync.RWMutex
	driveID         string
	driveName       string
	deltaQueryCount int
	Errors          *fault.Bus

	driveEnumerator driveEnumeratorFunc
	itemPagerFunc   func(
		servicer graph.Servicer,
		driveID, link string,
	) itemPager
}

// driveEnumeratorFunc enumerates all items in the specified drive and hands them to the
// provided `collector` method
type driveEnumeratorFunc func(
	ctx context.Context,
	pager itemPager,
	driveID, driveName string,
	collector itemCollector,
	oldPaths map[string]string,
	prevDelta string,
	errs *fault.Bus,
) (DeltaUpdate, map[string]string, map[string]struct{}, error)

func newUrlCache(driveEnumerator driveEnumeratorFunc) *urlCache {
	return &urlCache{
		urlMap:          map[string]itemProperties{},
		Errors:          fault.New(false),
		driveEnumerator: driveEnumerator,
		itemPagerFunc:   defaultItemPager,
	}
}

func (uc *urlCache) getDownloadUrl(
	ctx context.Context,
	svc graph.Servicer,
	itemID string) (string, bool) {
	// TODO: move map read/write to helpers
	// for safety
	uc.cacheLock.RLock()
	val, ok := uc.urlMap[itemID]
	uc.cacheLock.RUnlock()
	if !ok {
		// TODO: refresh cache locking
		uc.refreshCache(ctx, svc)
	}

	// TODO: handle deleted items
	return val.downloadUrl, ok
}

// refreshCache refreshes the URL cache by performing a delta query
func (uc *urlCache) refreshCache(
	ctx context.Context,
	svc graph.Servicer,
) error {
	uc.cacheLock.Lock()
	defer uc.cacheLock.Unlock()

	di, err := uc.deltaQuery(ctx, svc)
	if err != nil {
		return err
	}

	return nil
}

func (uc *urlCache) deltaQuery(
	ctx context.Context,
	svc graph.Servicer,
) ([]models.DriveItemable, error) {
	driveEnumerator := uc.driveEnumerator
	if driveEnumerator == nil {
		driveEnumerator = collectItems
	}

	_, _, _, err := driveEnumerator(
		ctx,
		uc.itemPagerFunc(svc, uc.driveID, uc.driveName),
		uc.driveID,
		uc.driveName,
		uc.cacheCollector,
		nil,
		"",
		uc.Errors,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (uc *urlCache) cacheCollector(
	ctx context.Context,
	driveID, driveName string,
	items []models.DriveItemable,
	oldPaths map[string]string,
	newPaths map[string]string,
	excluded map[string]struct{},
	itemCollection map[string]map[string]string,
	invalidPrevDelta bool,
	errs *fault.Bus,
) error {

	return nil
}
