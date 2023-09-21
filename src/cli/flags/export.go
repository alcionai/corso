package flags

import (
	"github.com/spf13/cobra"
)

const (
	ArchiveFN = "archive"
	FormatFN  = "format"
)

var (
	ArchiveFV bool
	FormatFV  string
)

// AddExportConfigFlags adds the restore config flag set.
func AddExportConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&ArchiveFV, ArchiveFN, false, "Export data as an archive instead of individual files")
	fs.StringVar(&FormatFV, FormatFN, "", "Specify the export file format")
	cobra.CheckErr(fs.MarkHidden(FormatFN))
}
