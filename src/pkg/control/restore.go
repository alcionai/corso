package control

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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
	OnCollision CollisionPolicy `json:"onCollision"`

	// ProtectedResource specifies which resource the data will be restored to.
	// If empty, restores to the same resource that was backed up.
	// Defaults to empty.
	ProtectedResource string `json:"protectedResource"`

	// Location specifies the container into which the data will be restored.
	// Only accepts container names, does not accept IDs.
	// If empty or "/", data will get restored in place, beginning at the root.
	// Defaults to "Corso_Restore_<current_dttm>"
	Location string `json:"location"`

	// Drive specifies the name of the drive into which the data will be
	// restored. If empty, data is restored to the same drive that was backed
	// up.
	// Defaults to empty.
	Drive string `json:"drive"`
}

func DefaultRestoreConfig(timeFormat dttm.TimeFormat) RestoreConfig {
	return RestoreConfig{
		OnCollision: Skip,
		Location:    DefaultRestoreLocation + dttm.FormatNow(timeFormat),
	}
}

func DefaultRestoreContainerName(timeFormat dttm.TimeFormat) string {
	return DefaultRestoreLocation + dttm.FormatNow(timeFormat)
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

// ---------------------------------------------------------------------------
// pii control
// ---------------------------------------------------------------------------

var (
	// interface compliance required for handling PII
	_ clues.Concealer = &RestoreConfig{}
	_ fmt.Stringer    = &RestoreConfig{}

	// interface compliance for the observe package to display
	// values without concealing PII.
	_ clues.PlainStringer = &RestoreConfig{}
)

func (rc RestoreConfig) marshal() string {
	bs, err := json.Marshal(rc)
	if err != nil {
		return "err marshalling"
	}

	return string(bs)
}

func (rc RestoreConfig) concealed() RestoreConfig {
	return RestoreConfig{
		OnCollision:       rc.OnCollision,
		ProtectedResource: clues.Hide(rc.ProtectedResource).Conceal(),
		Location:          path.LoggableDir(rc.Location),
		Drive:             clues.Hide(rc.Drive).Conceal(),
	}
}

// Conceal produces a concealed representation of the config, suitable for
// logging, storing in errors, and other output.
func (rc RestoreConfig) Conceal() string {
	return rc.concealed().marshal()
}

// Format produces a concealed representation of the config, even when
// used within a PrintF, suitable for logging, storing in errors,
// and other output.
func (rc RestoreConfig) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, rc.concealed())
}

// String returns a plain text version of the restoreConfig.
func (rc RestoreConfig) String() string {
	return rc.PlainString()
}

// PlainString returns an unescaped, unmodified string of the restore configuration.
func (rc RestoreConfig) PlainString() string {
	return rc.marshal()
}
