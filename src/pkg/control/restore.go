package control

import (
	"context"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	DefaultRestoreLocation = "Corso_Restore_"
)

// CollisionPolicy describes how the datalayer behaves in case of a collision.
type CollisionPolicy string

const (
	Unknown CollisionPolicy = ""
	Skip    CollisionPolicy = "skip"
	Copy    CollisionPolicy = "copy"
	Replace CollisionPolicy = "replace"
)

func ValidCollisionPolicies() map[CollisionPolicy]struct{} {
	return map[CollisionPolicy]struct{}{
		Skip:    {},
		Copy:    {},
		Replace: {},
	}
}

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
		Location:    DefaultRestoreLocation + dttm.FormatNow(timeFormat),
	}
}

// EnsureRestoreConfigDefaults sets all non-supported values in the config
// struct to the default value.
func EnsureRestoreConfigDefaults(
	ctx context.Context,
	rc RestoreConfig,
) RestoreConfig {
	if !slices.Contains(maps.Keys(ValidCollisionPolicies()), rc.OnCollision) {
		logger.Ctx(ctx).
			With(
				"bad_collision_policy", rc.OnCollision,
				"default_collision_policy", Skip).
			Info("setting collision policy to default")

		rc.OnCollision = Skip
	}

	rc.Location = strings.TrimPrefix(strings.TrimSpace(rc.Location), "/")

	return rc
}
