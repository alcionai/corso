package flags

import (
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/filters"
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

// ValidateExportConfigFlags ensures all export config flags that utilize
// enumerated values match a well-known value.
func ValidateExportConfigFlags() error {
	acceptedFormatTypes := []string{
		string(control.DefaultFormat),
		string(control.JSONFormat),
	}

	if !filters.Equal(acceptedFormatTypes).Compare(FormatFV) {
		return clues.New("unrecognized format type: " + FormatFV)
	}

	FormatFV = strings.ToLower(FormatFV)

	return nil
}
