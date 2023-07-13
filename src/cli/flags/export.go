package flags

import (
	"github.com/spf13/cobra"
)

const ArchiveFN = "archive"

var ArchiveFV bool

// AddExportConfigFlags adds the restore config flag set.
func AddExportConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&ArchiveFV, ArchiveFN, false, "Archive contents")
}
