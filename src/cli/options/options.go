package options

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/pkg/control"
)

var fastFail bool

// AddFlags adds the operation option flags
func AddOperationFlags(parent *cobra.Command) {
	fs := parent.Flags()
	fs.BoolVar(&fastFail, "fast-fail", false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden("fast-fail"))
}

// Control produces the control options based on the user's flags.
func Control() control.Options {
	return control.NewOptions(fastFail)
}
