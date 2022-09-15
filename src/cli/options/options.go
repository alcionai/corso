package options

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
)

var (
	fastFail  bool
	noMetrics bool
)

// AddOperationFlags adds command-local operation flags
func AddOperationFlags(parent *cobra.Command) {
	fs := parent.Flags()
	fs.BoolVar(&fastFail, "fast-fail", false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden("fast-fail"))
}

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.BoolVar(&noMetrics, "noMetrics", false, "disable anonymous information gathering")
}

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	if fastFail {
		opt.FailFast = true
	}

	if noMetrics {
		opt.DisableMetrics = true
	}

	return opt
}
