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
	Type   MaintenanceType   `json:"type"`
	Safety MaintenanceSafety `json:"safety"`
	Force  bool              `json:"force"`
}

// ---------------------------------------------------------------------------
// Maintenance flags
// ---------------------------------------------------------------------------

type MaintenanceType int

// Can't be reordered as we rely on iota for numbering.
//
//go:generate stringer -type=MaintenanceType -linecomment
const (
	CompleteMaintenance MaintenanceType = iota // complete
	MetadataMaintenance                        // metadata
)

var StringToMaintenanceType = map[string]MaintenanceType{
	CompleteMaintenance.String(): CompleteMaintenance,
	MetadataMaintenance.String(): MetadataMaintenance,
}

type MaintenanceSafety int

// Can't be reordered as we rely on iota for numbering.
//
//go:generate stringer -type=MaintenanceSafety -linecomment
const (
	FullMaintenanceSafety MaintenanceSafety = iota
	//nolint:lll
	// Use only if there's no other kopia instances accessing the repo and the
	// storage backend is strongly consistent.
	// https://github.com/kopia/kopia/blob/f9de453efc198b6e993af8922f953a7e5322dc5f/repo/maintenance/maintenance_safety.go#L42
	NoMaintenanceSafety
)

type RetentionMode int

// Can't be reordered as we rely on iota for numbering.
//
//go:generate stringer -type=RetentionMode -linecomment
const (
	UnknownRetention    RetentionMode = 0
	NoRetention         RetentionMode = 1
	GovernanceRetention RetentionMode = 2
	ComplianceRetention RetentionMode = 3
)

// Retention contains various options for configuring the retention mode. Takes
// pointers instead of values so that we can tell the difference between an
// unset value and a set but invalid value. This allows for partial
// (re)configuration of things.
type Retention struct {
	Mode     *RetentionMode
	Duration *time.Duration
	Extend   *bool
}
