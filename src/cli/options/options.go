package options

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	if failFastFV {
		opt.FailureHandling = control.FailFast
	}

	opt.DisableMetrics = noStatsFV
	opt.RestorePermissions = restorePermissionsFV
	opt.SkipReduce = skipReduceFV
	opt.ToggleFeatures.DisableIncrementals = disableIncrementalsFV
	opt.ToggleFeatures.DisableDelta = disableDeltaFV
	opt.ToggleFeatures.ExchangeImmutableIDs = enableImmutableID
	opt.ToggleFeatures.DisableConcurrencyLimiter = disableConcurrencyLimiterFV
	opt.Parallelism.ItemFetch = fetchParallelismFV

	opt.Repo.ViewTimestamp = viewTimestampFV.Get()

	return opt
}

// ---------------------------------------------------------------------------
// Operations Flags
// ---------------------------------------------------------------------------

const (
	FailFastFN                  = "fail-fast"
	FetchParallelismFN          = "fetch-parallelism"
	NoStatsFN                   = "no-stats"
	RestorePermissionsFN        = "restore-permissions"
	SkipReduceFN                = "skip-reduce"
	DisableDeltaFN              = "disable-delta"
	DisableIncrementalsFN       = "disable-incrementals"
	EnableImmutableIDFN         = "enable-immutable-id"
	DisableConcurrencyLimiterFN = "disable-concurrency-limiter"
	ViewTimestampFN             = "point-in-time"
)

var (
	failFastFV           bool
	fetchParallelismFV   int
	noStatsFV            bool
	restorePermissionsFV bool
	skipReduceFV         bool
	viewTimestampFV      timestamp
)

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&noStatsFV, NoStatsFN, false, "disable anonymous usage statistics gathering")
}

// AddFailFastFlag adds a flag to toggle fail-fast error handling behavior.
func AddFailFastFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&failFastFV, FailFastFN, false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden(FailFastFN))
}

// AddRestorePermissionsFlag adds OneDrive flag for restoring permissions
func AddRestorePermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&restorePermissionsFV, RestorePermissionsFN, false, "Restore permissions for files and folders")
}

// AddSkipReduceFlag adds a hidden flag that allows callers to skip the selector
// reduction step.  Currently only intended for details commands, not restore.
func AddSkipReduceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&skipReduceFV, SkipReduceFN, false, "Skip the selector reduce filtering")
	cobra.CheckErr(fs.MarkHidden(SkipReduceFN))
}

// AddFetchParallelismFlag adds a hidden flag that allows callers to reduce call
// paralellism (ie, the corso worker pool size) from 4 to as low as 1.
func AddFetchParallelismFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.IntVar(
		&fetchParallelismFV,
		FetchParallelismFN,
		4,
		"Control the number of concurrent data fetches for Exchange. Valid range is [1-4]. Default: 4")
	cobra.CheckErr(fs.MarkHidden(FetchParallelismFN))
}

// AddViewTimestampFlag adds a hidden flag that allows callers to pass a
// timestamp to view the corso repo at if immutable backups are enabled.
func AddViewTimestampFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.Var(
		&viewTimestampFV,
		ViewTimestampFN,
		"View the repository at a previous datetime by passing an RFC3339 timestamp")
	cobra.CheckErr(fs.MarkHidden(ViewTimestampFN))
}

// ---------------------------------------------------------------------------
// Feature Flags
// ---------------------------------------------------------------------------

var (
	disableIncrementalsFV bool
	disableDeltaFV        bool
)

// Adds the hidden '--disable-incrementals' cli flag which, when set, disables
// incremental backups.
func AddDisableIncrementalsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableIncrementalsFV,
		DisableIncrementalsFN,
		false,
		"Disable incremental data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(DisableIncrementalsFN))
}

// Adds the hidden '--disable-delta' cli flag which, when set, disables
// delta based backups.
func AddDisableDeltaFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableDeltaFV,
		DisableDeltaFN,
		false,
		"Disable delta based data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(DisableDeltaFN))
}

var enableImmutableID bool

// Adds the hidden '--enable-immutable-id' cli flag which, when set, enables
// immutable IDs for Exchange
func AddEnableImmutableIDFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&enableImmutableID,
		EnableImmutableIDFN,
		false,
		"Enable exchange immutable ID.")
	cobra.CheckErr(fs.MarkHidden(EnableImmutableIDFN))
}

var disableConcurrencyLimiterFV bool

// AddDisableConcurrencyLimiterFlag adds a hidden cli flag which, when set,
// removes concurrency limits when communicating with graph API. This
// flag is only relevant for exchange backups for now
func AddDisableConcurrencyLimiterFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableConcurrencyLimiterFV,
		DisableConcurrencyLimiterFN,
		false,
		"Disable concurrency limiter middleware. Default: false")
	cobra.CheckErr(fs.MarkHidden(DisableConcurrencyLimiterFN))
}
