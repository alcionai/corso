package utils

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/filters"
)

type ExportCfgOpts struct {
	Archive bool
	Format  string

	Populated flags.PopulatedFlags
}

func makeExportCfgOpts(cmd *cobra.Command) ExportCfgOpts {
	return ExportCfgOpts{
		Archive: flags.ArchiveFV,
		Format:  flags.FormatFV,

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

func MakeExportConfig(
	ctx context.Context,
	opts ExportCfgOpts,
) control.ExportConfig {
	exportCfg := control.DefaultExportConfig()

	exportCfg.Archive = opts.Archive
	exportCfg.Format = control.FormatType(opts.Format)

	return exportCfg
}

// ValidateExportConfigFlags ensures all export config flags that utilize
// enumerated values match a well-known value.
func ValidateExportConfigFlags(opts *ExportCfgOpts, acceptedFormatTypes []string) error {
	if _, populated := opts.Populated[flags.FormatFN]; !populated {
		opts.Format = string(control.DefaultFormat)
	} else if !filters.Equal(acceptedFormatTypes).Compare(opts.Format) {
		opts.Format = string(control.DefaultFormat)
		return clues.New("unrecognized format type: " + opts.Format)
	}

	opts.Format = strings.ToLower(opts.Format)

	return nil
}
