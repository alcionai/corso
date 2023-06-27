package flags

import (
	"github.com/spf13/cobra"
)

const (
	DisableConcurrencyLimiterFN = "disable-concurrency-limiter"
	DisableDeltaFN              = "disable-delta"
	DisableIncrementalsFN       = "disable-incrementals"
	EnableImmutableIDFN         = "enable-immutable-id"
	FailFastFN                  = "fail-fast"
	FailedItemsFN               = "failed-items"
	FetchParallelismFN          = "fetch-parallelism"
	NoStatsFN                   = "no-stats"
	RecoveredErrorsFN           = "recovered-errors"
	RestorePermissionsFN        = "restore-permissions"
	RunModeFN                   = "run-mode"
	SkippedItemsFN              = "skipped-items"
	SkipReduceFN                = "skip-reduce"
)

var (
	DisableConcurrencyLimiterFV bool
	DisableDeltaFV              bool
	DisableIncrementalsFV       bool
	EnableImmutableIDFV         bool
	FailFastFV                  bool
	FetchParallelismFV          int
	ListFailedItemsFV           string
	ListSkippedItemsFV          string
	ListRecoveredErrorsFV       string
	NoStatsFV                   bool
	// RunMode describes the type of run, such as:
	// flagtest, dry, run.  Should default to 'run'.
	RunModeFV            string
	RestorePermissionsFV bool
	SkipReduceFV         bool
)

// well-known flag values
const (
	RunModeFlagTest = "flag-test"
	RunModeRun      = "run"
)

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&NoStatsFV, NoStatsFN, false, "disable anonymous usage statistics gathering")
}

// AddFailFastFlag adds a flag to toggle fail-fast error handling behavior.
func AddFailFastFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&FailFastFV, FailFastFN, false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden(FailFastFN))
}

// AddRestorePermissionsFlag adds OneDrive flag for restoring permissions
func AddRestorePermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&RestorePermissionsFV, RestorePermissionsFN, false, "Restore permissions for files and folders")
}

// AddSkipReduceFlag adds a hidden flag that allows callers to skip the selector
// reduction step.  Currently only intended for details commands, not restore.
func AddSkipReduceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&SkipReduceFV, SkipReduceFN, false, "Skip the selector reduce filtering")
	cobra.CheckErr(fs.MarkHidden(SkipReduceFN))
}

// AddFetchParallelismFlag adds a hidden flag that allows callers to reduce call
// paralellism (ie, the corso worker pool size) from 4 to as low as 1.
func AddFetchParallelismFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.IntVar(
		&FetchParallelismFV,
		FetchParallelismFN,
		4,
		"Control the number of concurrent data fetches for Exchange. Valid range is [1-4]. Default: 4")
	cobra.CheckErr(fs.MarkHidden(FetchParallelismFN))
}

// Adds the hidden '--disable-incrementals' cli flag which, when set, disables
// incremental backups.
func AddDisableIncrementalsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&DisableIncrementalsFV,
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
		&DisableDeltaFV,
		DisableDeltaFN,
		false,
		"Disable delta based data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(DisableDeltaFN))
}

// Adds the hidden '--enable-immutable-id' cli flag which, when set, enables
// immutable IDs for Exchange
func AddEnableImmutableIDFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&EnableImmutableIDFV,
		EnableImmutableIDFN,
		false,
		"Enable exchange immutable ID.")
	cobra.CheckErr(fs.MarkHidden(EnableImmutableIDFN))
}

// AddDisableConcurrencyLimiterFlag adds a hidden cli flag which, when set,
// removes concurrency limits when communicating with graph API. This
// flag is only relevant for exchange backups for now
func AddDisableConcurrencyLimiterFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&DisableConcurrencyLimiterFV,
		DisableConcurrencyLimiterFN,
		false,
		"Disable concurrency limiter middleware. Default: false")
	cobra.CheckErr(fs.MarkHidden(DisableConcurrencyLimiterFN))
}

// AddRunModeFlag adds the hidden --run-mode flag.
func AddRunModeFlag(cmd *cobra.Command, persistent bool) {
	fs := cmd.Flags()
	if persistent {
		fs = cmd.PersistentFlags()
	}

	fs.StringVar(&RunModeFV, RunModeFN, "run", "What mode to run: dry, test, run.  Defaults to run.")
	cobra.CheckErr(fs.MarkHidden(RunModeFN))
}
