package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
)

const (
	CollisionsFN  = "collisions"
	DestinationFN = "destination"
)

var (
	CollisionsFV  string
	DestinationFV string
)

// AddRestoreConfigFlags adds the restore config flag set.
func AddRestoreConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&CollisionsFV, CollisionsFN, string(control.Skip),
		"How to handle item collisions: "+string(control.Skip)+", "+string(control.Copy)+", or "+string(control.Replace))
	cobra.CheckErr(fs.MarkHidden(CollisionsFN))
	fs.StringVar(
		&DestinationFV, DestinationFN, "",
		"Overrides the destination where items get restored.  '/' places items back in their original location.")
	cobra.CheckErr(fs.MarkHidden(DestinationFN))
}
