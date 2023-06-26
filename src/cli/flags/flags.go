package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const Wildcard = "*"

type PopulatedFlags map[string]struct{}

func (fs PopulatedFlags) populate(pf *pflag.Flag) {
	if pf == nil {
		return
	}

	if pf.Changed {
		fs[pf.Name] = struct{}{}
	}
}

// GetPopulatedFlags returns a map of flags that have been
// populated by the user.  Entry keys match the flag's long
// name.  Values are empty.
func GetPopulatedFlags(cmd *cobra.Command) PopulatedFlags {
	pop := PopulatedFlags{}

	fs := cmd.Flags()
	if fs == nil {
		return pop
	}

	fs.VisitAll(pop.populate)

	return pop
}
