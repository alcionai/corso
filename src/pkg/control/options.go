package control

import (
	"github.com/alcionai/corso/src/internal/common"
)

// Options holds the optional configurations for a process
type Options struct {
	Collision       CollisionPolicy `json:"-"`
	DisableMetrics  bool            `json:"disableMetrics"`
	FailFast        bool            `json:"failFast"`
	EnabledFeatures FeatureFlags    `json:"enabledFeatures"`
}

// Defaults provides an Options with the default values set.
func Defaults() Options {
	return Options{
		FailFast:        true,
		EnabledFeatures: FeatureFlags{},
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
// Feature Flags
// ---------------------------------------------------------------------------

type FeatureFlags struct {
	// ExchangeIncrementals allow for re-use of delta links when backing up
	// exchange data, reducing the amount of data pulled from graph.
	ExchangeIncrementals bool `json:"incrementals,omitempty"`
}
