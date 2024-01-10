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
	UseDeltaTree bool `json:"useDeltaTree"`

	// AddDisableLazyItemReader disables lazy item reader, such that we fall
	// back to prefetch reader. This flag is currently only meant for groups
	// conversations backup. Although it can be utilized for other services
	// in future.
	//
	// This flag should only be used if lazy item reader is the default choice
	// and we want to fallback to prefetch reader.
	DisableLazyItemReader bool `json:"disableLazyItemReader"`
}
