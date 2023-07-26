package utils

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/control"
)

type ExportCfgOpts struct {
	Archive bool

	Populated flags.PopulatedFlags
}

func makeExportCfgOpts(cmd *cobra.Command) ExportCfgOpts {
	return ExportCfgOpts{
		Archive: flags.ArchiveFV,

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

	return exportCfg
}
