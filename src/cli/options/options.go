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
	opt.ItemFetchParallelism = fetchParallelismFV

	return opt
}

// ---------------------------------------------------------------------------
// Operations Flags
// ---------------------------------------------------------------------------

const (
	failFastFN           = "fail-fast"
	fetchParallelismFN   = "fetch-parallelism"
	noStatsFN            = "no-stats"
	restorePermissionsFN = "restore-permissions"
	skipReduceFN         = "skip-reduce"
)

var (
	failFastFV           bool
	fetchParallelismFV   int
	noStatsFV            bool
	restorePermissionsFV bool
	skipReduceFV         bool
)

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&noStatsFV, noStatsFN, false, "disable anonymous usage statistics gathering")
}

// AddFailFastFlag adds a flag to toggle fail-fast error handling behavior.
func AddFailFastFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&failFastFV, failFastFN, false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden(failFastFN))
}

// AddRestorePermissionsFlag adds OneDrive flag for restoring permissions
func AddRestorePermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&restorePermissionsFV, restorePermissionsFN, false, "Restore permissions for files and folders")
}

// AddSkipReduceFlag adds a hidden flag that allows callers to skip the selector
// reduction step.  Currently only intended for details commands, not restore.
func AddSkipReduceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&skipReduceFV, skipReduceFN, false, "Skip the selector reduce filtering")
	cobra.CheckErr(fs.MarkHidden(skipReduceFN))
}

// AddFetchParallelismFlag adds a hidden flag that allows callers to reduce call
// paralellism (ie, the corso worker pool size) from 4 to as low as 1.
func AddFetchParallelismFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.IntVar(
		&fetchParallelismFV,
		fetchParallelismFN,
		4,
		"Control the number of concurrent data fetches for Exchange. Valid range is [1-4]. Default: 4")
	cobra.CheckErr(fs.MarkHidden(fetchParallelismFN))
}

// ---------------------------------------------------------------------------
// Feature Flags
// ---------------------------------------------------------------------------

const disableIncrementalsFN = "disable-incrementals"

var disableIncrementalsFV bool

// Adds the hidden '--disable-incrementals' cli flag which, when set, disables
// incremental backups.
func AddDisableIncrementalsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableIncrementalsFV,
		disableIncrementalsFN,
		false,
		"Disable incremental data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(disableIncrementalsFN))
}
