package utils

import (
	"strconv"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

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

// IsValidTimeFormat returns true if the input is recognized as a
// supported format by the common time parser.
func IsValidTimeFormat(in string) bool {
	_, err := common.ParseTime(in)
	return err == nil
}

// IsValidTimeFormat returns true if the input is recognized as a
// boolean.
func IsValidBool(in string) bool {
	_, err := strconv.ParseBool(in)
	return err == nil
}
