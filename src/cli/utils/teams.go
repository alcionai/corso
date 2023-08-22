package utils

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
)

type TeamsOpts struct {
	Teams []string

	RestoreCfg RestoreCfgOpts
	ExportCfg  ExportCfgOpts

	Populated flags.PopulatedFlags
}

func MakeTeamsOpts(cmd *cobra.Command) TeamsOpts {
	return TeamsOpts{
		Teams: flags.UserFV,

		RestoreCfg: makeRestoreCfgOpts(cmd),
		ExportCfg:  makeExportCfgOpts(cmd),

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}
