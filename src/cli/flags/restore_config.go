package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
)

const (
	CollisionsFN  = "collisions"
	DestinationFN = "destination"
	ToResourceFN  = "to-resource"
)

var (
	CollisionsFV  string
	DestinationFV string
	ToResourceFV  string
)

// AddRestoreConfigFlags adds the restore config flag set.
func AddRestoreConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&CollisionsFV, CollisionsFN, string(control.Skip),
		//nolint:lll
		"Sets the behavior for existing item collisions: "+string(control.Skip)+", "+string(control.Copy)+", or "+string(control.Replace))
	fs.StringVar(
		&DestinationFV, DestinationFN, "",
		"Overrides the folder where items get restored; '/' places items into their original location")
	fs.StringVar(
		&ToResourceFV, ToResourceFN, "",
		"Overrides the protected resource (mailbox, site, user, etc) where data gets restored")
}
