package control

import (
	"context"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/logger"
)

// Options holds the optional configurations for a process
type Options struct {
	DisableMetrics     bool               `json:"disableMetrics"`
	FailureHandling    FailurePolicy      `json:"failureHandling"`
	RestorePermissions bool               `json:"restorePermissions"`
	SkipReduce         bool               `json:"skipReduce"`
	ToggleFeatures     Toggles            `json:"toggleFeatures"`
	Parallelism        Parallelism        `json:"parallelism"`
	Repo               repository.Options `json:"repo"`
}

type Parallelism struct {
	// sets the collection buffer size before blocking.
	CollectionBuffer int
	// sets the parallelism of item population within a collection.
	ItemFetch int
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
// Restore Configuration
// ---------------------------------------------------------------------------

const (
	defaultRestoreLocation = "Corso_Restore_"
)

// CollisionPolicy describes how the datalayer behaves in case of a collision.
type CollisionPolicy string

const (
	Unknown CollisionPolicy = ""
	Skip    CollisionPolicy = "skip"
	Copy    CollisionPolicy = "copy"
	Replace CollisionPolicy = "replace"
)

const RootLocation = "/"

// RestoreConfig contains
type RestoreConfig struct {
	// Defines the per-item collision handling policy.
	// Defaults to Skip.
	OnCollision CollisionPolicy

	// ProtectedResource specifies which resource the data will be restored to.
	// If empty, restores to the same resource that was backed up.
	// Defaults to empty.
	ProtectedResource string

	// Location specifies the container into which the data will be restored.
	// Only accepts container names, does not accept IDs.
	// If empty or "/", data will get restored in place, beginning at the root.
	// Defaults to "Corso_Restore_<current_dttm>"
	Location string

	// Drive specifies the drive into which the data will be restored.
	// If empty, data is restored to the same drive that was backed up.
	// Defaults to empty.
	Drive string
}

func DefaultRestoreConfig(timeFormat dttm.TimeFormat) RestoreConfig {
	return RestoreConfig{
		OnCollision: Skip,
		Location:    defaultRestoreLocation + dttm.FormatNow(timeFormat),
	}
}

// EnsureRestoreConfigDefaults sets all non-supported values in the config
// struct to the default value.
func EnsureRestoreConfigDefaults(
	ctx context.Context,
	rc RestoreConfig,
) RestoreConfig {
	if !slices.Contains([]CollisionPolicy{Skip, Copy, Replace}, rc.OnCollision) {
		logger.Ctx(ctx).
			With(
				"bad_collision_policy", rc.OnCollision,
				"default_collision_policy", Skip).
			Info("setting collision policy to default")

		rc.OnCollision = Skip
	}

	if strings.TrimSpace(rc.Location) == RootLocation {
		rc.Location = ""
	}

	return rc
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

	// DisableConcurrencyLimiter removes concurrency limits when communicating with
	// graph API. This flag is only relevant for exchange backups for now
	DisableConcurrencyLimiter bool `json:"disableConcurrencyLimiter,omitempty"`
}
