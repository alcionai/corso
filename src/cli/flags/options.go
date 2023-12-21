package flags

import (
	"github.com/spf13/cobra"
)

const (
	AlertsFN                      = "alerts"
	ConfigFileFN                  = "config-file"
	DeltaPageSizeFN               = "delta-page-size"
	DisableDeltaFN                = "disable-delta"
	DisableIncrementalsFN         = "disable-incrementals"
	DisableSlidingWindowLimiterFN = "disable-sliding-window-limiter"
	ForceItemDataDownloadFN       = "force-item-data-download"
	EnableImmutableIDFN           = "enable-immutable-id"
	FailFastFN                    = "fail-fast"
	FailedItemsFN                 = "failed-items"
	FetchParallelismFN            = "fetch-parallelism"
	NoStatsFN                     = "no-stats"
	RecoveredErrorsFN             = "recovered-errors"
	NoPermissionsFN               = "no-permissions"
	RunModeFN                     = "run-mode"
	SkippedItemsFN                = "skipped-items"
	SkipReduceFN                  = "skip-reduce"
)

var (
	ConfigFileFV                  string
	DeltaPageSizeFV               int
	DisableDeltaFV                bool
	DisableIncrementalsFV         bool
	DisableSlidingWindowLimiterFV bool
	ForceItemDataDownloadFV       bool
	EnableImmutableIDFV           bool
	FailFastFV                    bool
	FetchParallelismFV            int
	ListAlertsFV                  string
	ListFailedItemsFV             string
	ListSkippedItemsFV            string
	ListRecoveredErrorsFV         string
	NoStatsFV                     bool
	// RunMode describes the type of run, such as:
	// flagtest, dry, run.  Should default to 'run'.
	RunModeFV       string
	NoPermissionsFV bool
	SkipReduceFV    bool
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

// AddNoPermissionsFlag adds OneDrive flag for skipping restoring permissions
func AddNoPermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&NoPermissionsFV, NoPermissionsFN, false, "don't restore file and folder permissions")
}

// AddSkipReduceFlag adds a hidden flag that allows callers to skip the selector
// reduction step.  Currently only intended for details commands, not restore.
func AddSkipReduceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&SkipReduceFV, SkipReduceFN, false, "Skip the selector reduce filtering")
	cobra.CheckErr(fs.MarkHidden(SkipReduceFN))
}

// AddDeltaPageSizeFlag adds a hidden flag that allows callers to reduce delta
// query page sizes below 500.
func AddDeltaPageSizeFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.IntVar(
		&DeltaPageSizeFV,
		DeltaPageSizeFN,
		500,
		"Control quantity of items returned in paged queries. Valid range is [1-500]. Default: 500")
	cobra.CheckErr(fs.MarkHidden(DeltaPageSizeFN))
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

// Adds the hidden '--force-item-data-download' cli flag which, when set,
// disables kopia-assisted incremental backups.
func AddForceItemDataDownloadFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&ForceItemDataDownloadFV,
		ForceItemDataDownloadFN,
		false,
		"Disable cached data checks in backups to force item redownloads for "+
			"items changed since the last successful backup.")
	cobra.CheckErr(fs.MarkHidden(ForceItemDataDownloadFN))
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

// AddRunModeFlag adds the hidden --run-mode flag.
func AddRunModeFlag(cmd *cobra.Command, persistent bool) {
	fs := cmd.Flags()
	if persistent {
		fs = cmd.PersistentFlags()
	}

	fs.StringVar(&RunModeFV, RunModeFN, "run", "What mode to run: dry, test, run.  Defaults to run.")
	cobra.CheckErr(fs.MarkHidden(RunModeFN))
}

// AddDisableSlidingWindowLimiterFN disables the experimental sliding window rate
// limiter for graph API requests. This is only relevant for exchange backups.
// Exchange restores continue to use the default token bucket rate limiter.
// Setting this flag switches exchange backups to use the default token bucket
// rate limiter.
func AddDisableSlidingWindowLimiterFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&DisableSlidingWindowLimiterFV,
		DisableSlidingWindowLimiterFN,
		false,
		"Disable sliding window rate limiter.")
	cobra.CheckErr(fs.MarkHidden(DisableSlidingWindowLimiterFN))
}
