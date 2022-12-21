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

	if exchangeIncrementals {
		opt.EnabledFeatures.ExchangeIncrementals = true
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

var exchangeIncrementals bool

type exposeFeatureFlag func(*pflag.FlagSet)

// AddFeatureFlags adds CLI flags for each exposed feature flags to the
// persistent flag set within the command.
func AddFeatureFlags(cmd *cobra.Command, effs ...exposeFeatureFlag) {
	fs := cmd.PersistentFlags()
	for _, fflag := range effs {
		fflag(fs)
	}
}

// Adds the '--exchange-incrementals' cli flag which, when set, enables
// incrementals data retrieval for exchange backups.
func ExchangeIncrementals() func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.BoolVar(
			&exchangeIncrementals,
			"exchange-incrementals",
			false,
			"Enable incremental data retrieval in Exchange backups.")
		cobra.CheckErr(fs.MarkHidden("exchange-incrementals"))
	}
}
