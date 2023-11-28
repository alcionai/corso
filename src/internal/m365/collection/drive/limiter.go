package drive

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/control"
)

var errHitLimit = clues.New("hit limiter limits")

type driveEnumerationStats struct {
	numPages      int
	numAddedFiles int
	numContainers int
	numBytes      int64
}

func newPagerLimiter(opts control.Options) *pagerLimiter {
	res := &pagerLimiter{limits: opts.PreviewLimits}

	if res.limits.MaxContainers == 0 {
		res.limits.MaxContainers = defaultPreviewMaxContainers
	}

	if res.limits.MaxItemsPerContainer == 0 {
		res.limits.MaxItemsPerContainer = defaultPreviewMaxItemsPerContainer
	}

	if res.limits.MaxItems == 0 {
		res.limits.MaxItems = defaultPreviewMaxItems
	}

	if res.limits.MaxBytes == 0 {
		res.limits.MaxBytes = defaultPreviewMaxBytes
	}

	if res.limits.MaxPages == 0 {
		res.limits.MaxPages = defaultPreviewMaxPages
	}

	return res
}

type pagerLimiter struct {
	limits control.PreviewItemLimits
}

func (l pagerLimiter) effectiveLimits() control.PreviewItemLimits {
	return l.limits
}

func (l pagerLimiter) enabled() bool {
	return l.limits.Enabled
}

// sizeLimit returns the total number of bytes this backup should try to
// contain.
func (l pagerLimiter) sizeLimit() int64 {
	return l.limits.MaxBytes
}

// atItemLimit returns true if the limiter is enabled and has reached the limit
// for individual items added to collections for this backup.
func (l pagerLimiter) atItemLimit(stats *driveEnumerationStats) bool {
	return l.enabled() &&
		(stats.numAddedFiles >= l.limits.MaxItems ||
			stats.numBytes >= l.limits.MaxBytes)
}

// atContainerItemsLimit returns true if the limiter is enabled and the current
// number of items is above the limit for the number of items for a container
// for this backup.
func (l pagerLimiter) atContainerItemsLimit(numItems int) bool {
	return l.enabled() && numItems >= l.limits.MaxItemsPerContainer
}

// atPageLimit returns true if the limiter is enabled and the number of
// pages processed so far is beyond the limit for this backup.
func (l pagerLimiter) atPageLimit(stats *driveEnumerationStats) bool {
	return l.enabled() && stats.numPages >= l.limits.MaxPages
}

// atLimit returns true if the limiter is enabled and meets any of the
// conditions for max items, containers, etc for this backup.
func (l pagerLimiter) atLimit(stats *driveEnumerationStats) bool {
	return l.enabled() &&
		(l.atItemLimit(stats) ||
			stats.numContainers >= l.limits.MaxContainers ||
			stats.numPages >= l.limits.MaxPages)
}

// ---------------------------------------------------------------------------
// Used by the tree version limit handling
// ---------------------------------------------------------------------------

// hitPageLimit returns true if the limiter is enabled and the number of
// pages processed so far is beyond the limit for this backup.
func (l pagerLimiter) hitPageLimit(pageCount int) bool {
	return l.enabled() && pageCount >= l.limits.MaxPages
}

// hitContainerLimit returns true if the limiter is enabled and the number of
// unique containers added so far is beyond the limit for this backup.
func (l pagerLimiter) hitContainerLimit(containerCount int) bool {
	return l.enabled() && containerCount >= l.limits.MaxContainers
}

// hitItemLimit returns true if the limiter is enabled and has reached the limit
// for unique items, or their accumulated size in bytes, added to collections for this backup.
func (l pagerLimiter) hitItemLimit(itemCount int) bool {
	return l.enabled() && itemCount >= l.limits.MaxItems
}

// hitTotalBytesLimit returns true if the limiter is enabled and has reached the limit
// for the accumulated byte size of all items (the file contents, not the item metadata)
// added to collections for this backup.
func (l pagerLimiter) hitTotalBytesLimit(i int64) bool {
	return l.enabled() && i >= l.limits.MaxBytes
}
