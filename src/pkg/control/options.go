package control

import (
	"github.com/alcionai/corso/src/internal/common"
)

// Options holds the optional configurations for a process
type Options struct {
	Collision          CollisionPolicy `json:"-"`
	DisableMetrics     bool            `json:"disableMetrics"`
	FailureHandling    FailureBehavior `json:"failureHandling"`
	RestorePermissions bool            `json:"restorePermissions"`
	SkipReduce         bool            `json:"skipReduce"`
	ToggleFeatures     Toggles         `json:"toggleFeatures"`
	Parallelism        Parallelism     `json:"parallelism"`
	Repo               RepoOptions     `json:"repo"`
}

type FailureBehavior string

type Parallelism struct {
	// sets the collection buffer size before blocking.
	CollectionBuffer int
	// sets the parallelism of item population within a collection.
	ItemFetch int
}

const (
	// fails and exits the run immediately
	FailFast FailureBehavior = "fail-fast"
	// recovers whenever possible, reports non-zero recoveries as a failure
	FailAfterRecovery FailureBehavior = "fail-after-recovery"
	// recovers whenever possible, does not report recovery as failure
	BestEffort FailureBehavior = "best-effort"
)

// Repo represents options that are specific to the repo storing backed up data.
type RepoOptions struct {
	User string `json:"user"`
	Host string `json:"host"`
}

// Defaults provides an Options with the default values set.
func Defaults() Options {
	return Options{
		FailureHandling: FailAfterRecovery,
		ToggleFeatures:  Toggles{},
		Parallelism: Parallelism{
			CollectionBuffer: 4,
			ItemFetch:        4,
		},
	}
}

// ---------------------------------------------------------------------------
// Restore Item Collision Policy
// ---------------------------------------------------------------------------

// CollisionPolicy describes how the datalayer behaves in case of a collision.
type CollisionPolicy int

//go:generate stringer -type=CollisionPolicy
const (
	Unknown CollisionPolicy = iota
	Copy
	Skip
	Replace
)

// ---------------------------------------------------------------------------
// Restore Destination
// ---------------------------------------------------------------------------

const (
	defaultRestoreLocation = "Corso_Restore_"
)

// RestoreDestination is a POD that contains an override of the resource owner
// to restore data under and the name of the root of the restored container
// hierarchy.
type RestoreDestination struct {
	// ResourceOwnerOverride overrides the default resource owner to restore to.
	// If it is not populated items should be restored under the previous resource
	// owner of the item.
	ResourceOwnerOverride string
	// ContainerName is the name of the root of the restored container hierarchy.
	// This field must be populated for a restore.
	ContainerName string
}

func DefaultRestoreDestination(timeFormat common.TimeFormat) RestoreDestination {
	return RestoreDestination{
		ContainerName: defaultRestoreLocation + common.FormatNow(timeFormat),
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
	// ExchangeImmutableIDs denotes whether Corso should store items with
	// immutable Exchange IDs. This is only safe to set if the previous backup for
	// incremental backups used immutable IDs or if a full backup is being done.
	ExchangeImmutableIDs bool `json:"exchangeImmutableIDs,omitempty"`

	RunMigrations bool `json:"runMigrations"`

	// DisableConcurrencyLimiter removes concurrency limits when communicating with
	// graph API. This flag is only relevant for exchange backups for now
	DisableConcurrencyLimiter bool `json:"disableConcurrencyLimiter,omitempty"`
}

// ---------------------------------------------------------------------------
// Repo Maintenance flags
// ---------------------------------------------------------------------------

type Safety string

const (
	FullSafety Safety = "full"
	NoSafety   Safety = "none"
)

var SafetyValues = map[Safety]struct{}{
	FullSafety: {},
	NoSafety:   {},
}
