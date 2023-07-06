package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
)

const (
	ArchiveFN    = "archive"
	ExportDestFN = "destination"
)

var (
	ArchiveFV    bool
	ExportDestFV string
)

// AddExportConfigFlags adds the restore config flag set.
func AddExportConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&ArchiveFV, ArchiveFN, false, "Archive contents")

	// TODO(meain): convert to mandatory positional argument
	fs.StringVar(
		&ExportDestFV,
		ExportDestFN,
		"",
		"Folder to export to (default:"+control.DefaultRestoreLocation+"<timestamp>)",
	)
}
