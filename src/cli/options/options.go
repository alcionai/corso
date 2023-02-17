package options

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/pkg/control"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	opt.FailFast = strings.ToLower(onError) != "continue"
	opt.DisableMetrics = noStats
	opt.RestorePermissions = restorePermissions
	opt.ToggleFeatures.DisableIncrementals = disableIncrementals
	opt.ToggleFeatures.EnablePermissionsBackup = enablePermissionsBackup

	return opt
}

// ---------------------------------------------------------------------------
// Operations Flags
// ---------------------------------------------------------------------------

var (
	onError            string
	noStats            bool
	restorePermissions bool
)

// AddOperationFlags adds command-local operation flags
func AddOperationFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&onError, "on-error",
		"continue",
		"Configure how Corso handles recoverable errors.  Accepts 'fail' or 'continue' (default).")
}

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&noStats, "no-stats", false, "disable anonymous usage statistics gathering")
}

// AddRestorePermissionsFlag adds OneDrive flag for restoring permissions
func AddRestorePermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&restorePermissions, "restore-permissions", false, "Restore permissions for files and folders")
	// TODO: reveal this flag once backing up permissions becomes default
	cobra.CheckErr(fs.MarkHidden("restore-permissions"))
}

// ---------------------------------------------------------------------------
// Feature Flags
// ---------------------------------------------------------------------------

var (
	disableIncrementals     bool
	enablePermissionsBackup bool
)

type exposeFeatureFlag func(*pflag.FlagSet)

// AddFeatureToggle adds CLI flags for each exposed feature toggle to the
// persistent flag set within the command.
func AddFeatureToggle(cmd *cobra.Command, effs ...exposeFeatureFlag) {
	fs := cmd.PersistentFlags()
	for _, fflag := range effs {
		fflag(fs)
	}
}

// Adds the hidden '--no-incrementals' cli flag which, when set, disables
// incremental backups.
func DisableIncrementals() func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.BoolVar(
			&disableIncrementals,
			"disable-incrementals",
			false,
			"Disable incremental data retrieval in backups.")
		cobra.CheckErr(fs.MarkHidden("disable-incrementals"))
	}
}

// Adds the hidden '--enable-permissions-backup' cli flag which, when
// set, enables backing up permissions.
func EnablePermissionsBackup() func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.BoolVar(
			&enablePermissionsBackup,
			"enable-permissions-backup",
			false,
			"Enable backing up item permissions for OneDrive")
		cobra.CheckErr(fs.MarkHidden("enable-permissions-backup"))
	}
}
