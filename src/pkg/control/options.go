package control

import (
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/extensions"
)

// Options holds the optional configurations for a process
type Options struct {
	// DeltaPageSize controls the quantity of items fetched in each page
	// during multi-page queries, such as graph api delta endpoints.
	DeltaPageSize        int32                              `json:"deltaPageSize"`
	DisableMetrics       bool                               `json:"disableMetrics"`
	FailureHandling      FailurePolicy                      `json:"failureHandling"`
	ItemExtensionFactory []extensions.CreateItemExtensioner `json:"-"`
	Parallelism          Parallelism                        `json:"parallelism"`
	Repo                 repository.Options                 `json:"repo"`
	SkipReduce           bool                               `json:"skipReduce"`
	ToggleFeatures       Toggles                            `json:"toggleFeatures"`
	// PreviewItemLimits defines the number of items and/or amount of data to
	// fetch on a best-effort basis. Right now it's used for preview backups.
	//
	// Since this is not split out by service or data categories these limits
	// apply independently to all data categories that appear in a single backup
	// where they are set. For example, if doing a teams backup and there's both a
	// SharePoint site and Messages available, both data categories would try to
	// backup data until the set limits without paying attention to what the other
	// had already backed up.
	PreviewLimits PreviewItemLimits `json:"previewItemLimits"`

	// specifying a resource tuple in this map allows that resource to produce
	// a Skip instead of a recoverable error in case of a failure due to 503 when
	// retrieving calendar event item data.
	SkipEventsOnInstance503ForResources map[string]struct{}
}

// RateLimiter is the set of options applied to any external service facing rate
// limiters Corso may use during backups or restores.
type RateLimiter struct {
	DisableSlidingWindowLimiter bool `json:"disableSlidingWindowLimiter"`
}

type FailurePolicy string

const (
	// fails and exits the run immediately
	FailFast FailurePolicy = "fail-fast"
	// recovers whenever possible, reports non-zero recoveries as a failure
	FailAfterRecovery FailurePolicy = "fail-after-recovery"
	// recovers whenever possible, does not report recovery as failure
	BestEffort FailurePolicy = "best-effort"
)

// DefaultOptions provides an Options with the default values set.
func DefaultOptions() Options {
	return Options{
		FailureHandling: FailAfterRecovery,
		DeltaPageSize:   500,
		ToggleFeatures:  Toggles{},
		Parallelism: Parallelism{
			CollectionBuffer: 4,
			ItemFetch:        4,
		},
	}
}

// ---------------------------------------------------------------------------
// Feature Flags and Toggles
// ---------------------------------------------------------------------------

// Toggles allows callers to force corso to behave in ways that deviate from
// the default expectations by turning on or shutting off certain features.
// The default state for every toggle is false; toggles are only turned on
// if specified by the caller.
type Toggles struct {
	// DisableIncrementals prevents backups from using incremental lookups,
	// forcing a new, complete backup of all data regardless of prior state.
	DisableIncrementals bool `json:"exchangeIncrementals,omitempty"`
	// ForceItemDataDownload disables finding cached items in previous failed
	// backups (i.e. kopia-assisted incrementals). Data dedupe will still occur
	// since that is based on content hashes. Items that have not changed since
	// the previous backup (i.e. in the merge base) will not be redownloaded. Use
	// DisableIncrementals to control that behavior.
	ForceItemDataDownload bool `json:"forceItemDataDownload,omitempty"`
	// DisableDelta prevents backups from using delta based lookups,
	// forcing a backup by enumerating all items. This is different
	// from DisableIncrementals in that this does not even makes use of
	// delta endpoints with or without a delta token. This is necessary
	// when the user has filled up the mailbox storage available to the
	// user as Microsoft prevents the API from being able to make calls
	// to delta endpoints.
	DisableDelta bool `json:"exchangeDelta,omitempty"`
	// ExchangeImmutableIDs denotes whether Corso should store items with
	// immutable Exchange IDs. This is only safe to set if the previous backup for
	// incremental backups used immutable IDs or if a full backup is being done.
	ExchangeImmutableIDs bool `json:"exchangeImmutableIDs,omitempty"`

	RunMigrations bool `json:"runMigrations"`

	// DisableSlidingWindowLimiter disables the experimental sliding window rate
	// limiter for graph API requests. This is only relevant for exchange backups.
	// Setting this flag switches exchange backups to fallback to the default token
	// bucket rate limiter.
	DisableSlidingWindowLimiter bool `json:"disableSlidingWindowLimiter"`

	// see: https://github.com/alcionai/corso/issues/4688
	UseDeltaTree       bool `json:"useDeltaTree"`
	UseOldDeltaProcess bool `json:"useOldDeltaProcess"`

	// AddDisableLazyItemReader disables lazy item reader, such that we fall
	// back to prefetch reader. This flag is currently only meant for groups
	// conversations backup. Although it can be utilized for other services
	// in future.
	//
	// This flag should only be used if lazy item reader is the default choice
	// and we want to fallback to prefetch reader.
	DisableLazyItemReader bool `json:"disableLazyItemReader"`
}
