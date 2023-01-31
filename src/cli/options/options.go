package options

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/pkg/control"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	if fastFail {
		opt.FailFast = true
	}

	if noStats {
		opt.DisableMetrics = true
	}

	if disableIncrementals {
		opt.ToggleFeatures.DisableIncrementals = true
	}

	return opt
}

// ---------------------------------------------------------------------------
// Operations Flags
// ---------------------------------------------------------------------------

var (
	fastFail bool
	noStats  bool
)

// AddOperationFlags adds command-local operation flags
func AddOperationFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&fastFail, "fast-fail", false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden("fast-fail"))
}

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&noStats, "no-stats", false, "disable anonymous usage statistics gathering")
}

// ---------------------------------------------------------------------------
// Feature Flags
// ---------------------------------------------------------------------------

var disableIncrementals bool

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
