package repository

import (
	"time"
)

// Repo represents options that are specific to the repo storing backed up data.
type Options struct {
	User string `json:"user"`
	Host string `json:"host"`
	// ViewTimestamp is the time at which the repo should be opened at if
	// immutable backups are being used. If nil then the current time is used.
	ViewTimestamp *time.Time `json:"viewTimestamp"`
}

type Maintenance struct {
	Type          MaintenanceType   `json:"type"`
	Safety        MaintenanceSafety `json:"safety"`
	Force         bool              `json:"force"`
	CleanupBuffer *time.Duration
}

// ---------------------------------------------------------------------------
// Maintenance flags
// ---------------------------------------------------------------------------

type MaintenanceType int

//go:generate stringer -type=MaintenanceType -linecomment
const (
	CompleteMaintenance MaintenanceType = 0 // complete
	MetadataMaintenance MaintenanceType = 1 // metadata
	// Adding this here so we can test failed backup cleanup more without forcing
	// SDK consumers to use it right away. We can remove this and add the
	// additional cleanup functionality to CompleteMaintenance when we feel more
	// comfortable with it.
	CompletePlusMaintenance MaintenanceType = 2 // dontuse
)

var StringToMaintenanceType = map[string]MaintenanceType{
	CompleteMaintenance.String(): CompleteMaintenance,
	MetadataMaintenance.String(): MetadataMaintenance,
}

type MaintenanceSafety int

//go:generate stringer -type=MaintenanceSafety -linecomment
const (
	FullMaintenanceSafety MaintenanceSafety = 0
	//nolint:lll
	// Use only if there's no other kopia instances accessing the repo and the
	// storage backend is strongly consistent.
	// https://github.com/kopia/kopia/blob/f9de453efc198b6e993af8922f953a7e5322dc5f/repo/maintenance/maintenance_safety.go#L42
	NoMaintenanceSafety MaintenanceSafety = 1
)

type RetentionMode int

//go:generate stringer -type=RetentionMode -linecomment
const (
	UnknownRetention    RetentionMode = 0
	NoRetention         RetentionMode = 1 // none
	GovernanceRetention RetentionMode = 2 // governance
	ComplianceRetention RetentionMode = 3 // compliance
)

func ValidRetentionModeNames() map[string]RetentionMode {
	return map[string]RetentionMode{
		NoRetention.String():         NoRetention,
		GovernanceRetention.String(): GovernanceRetention,
		ComplianceRetention.String(): ComplianceRetention,
	}
}

// Retention contains various options for configuring the retention mode. Takes
// pointers instead of values so that we can tell the difference between an
// unset value and a set but invalid value. This allows for partial
// (re)configuration of things.
type Retention struct {
	Mode     *RetentionMode
	Duration *time.Duration
	Extend   *bool
}
