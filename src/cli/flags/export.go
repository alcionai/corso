package flags

import (
	"github.com/spf13/cobra"
)

const (
	ArchiveFN = "archive"
	RawFN     = "raw"
)

var (
	ArchiveFV bool
	RawFV     bool
)

// AddExportConfigFlags adds the restore config flag set.
func AddExportConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&ArchiveFV, ArchiveFN, false, "Export data as an archive instead of individual files")
	fs.BoolVar(&RawFV, RawFN, false, "Export data in its original format")
	cobra.CheckErr(fs.MarkHidden(RawFN))
}
