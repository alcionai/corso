package options

import (
	"github.com/alcionai/corso/internal/operations"
	"github.com/spf13/cobra"
)

var (
	fastFail bool
)

// AddFlags adds the operation option flags
func AddOperationFlags(parent *cobra.Command) {
	fs := parent.Flags()
	fs.BoolVar(&fastFail, "fast-fail", false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	fs.MarkHidden("fast-fail")
}

// OperationOptions produces the operation options based on the user's flags.
func OperationOptions() operations.Options {
	return operations.NewOptions(fastFail)
}
