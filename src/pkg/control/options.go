package control

import (
	"github.com/alcionai/corso/src/internal/common"
)

// Options holds the optional configurations for a process
type Options struct {
	Collision            CollisionPolicy `json:"-"`
	DisableMetrics       bool            `json:"disableMetrics"`
	FailureHandling      FailureBehavior `json:"failureHandling"`
	ItemFetchParallelism int             `json:"itemFetchParallelism"`
	RestorePermissions   bool            `json:"restorePermissions"`
	SkipReduce           bool            `json:"skipReduce"`
	ToggleFeatures       Toggles         `json:"ToggleFeatures"`
}

type FailureBehavior string

const (
	// fails and exits the run immediately
	FailFast FailureBehavior = "fail-fast"
	// recovers whenever possible, reports non-zero recoveries as a failure
	FailAfterRecovery FailureBehavior = "fail-after-recovery"
	// recovers whenever possible, does not report recovery as failure
	BestEffort FailureBehavior = "best-effort"
)

// Defaults provides an Options with the default values set.
func Defaults() Options {
	return Options{
		FailureHandling: FailAfterRecovery,
		ToggleFeatures:  Toggles{},
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

	// EnablePermissionsBackup is used to enable backups of item
	// permissions. Permission metadata increases graph api call count,
	// so disabling their retrieval when not needed is advised.
	EnablePermissionsBackup bool `json:"enablePermissionsBackup,omitempty"`
}
