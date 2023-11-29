package drive

import "github.com/alcionai/corso/src/pkg/control"

// used to mark an unused variable while we transition handling.
const ignoreMe = -1

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

// atContainerPageLimit returns true if the limiter is enabled and the number of
// pages processed so far is beyond the limit for this backup.
func (l pagerLimiter) atPageLimit(stats *driveEnumerationStats) bool {
	return l.enabled() && stats.numPages >= l.limits.MaxPages
}

// atLimit returns true if the limiter is enabled and meets any of the
// conditions for max items, containers, etc for this backup.
func (l pagerLimiter) atLimit(
	stats *driveEnumerationStats,
	containerCount int,
) bool {
	nc := stats.numContainers
	if nc == 0 && containerCount > 0 {
		nc = containerCount
	}

	return l.enabled() &&
		(l.atItemLimit(stats) ||
			nc >= l.limits.MaxContainers ||
			stats.numPages >= l.limits.MaxPages)
}
