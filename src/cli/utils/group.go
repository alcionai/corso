package utils

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GroupOpts struct {
	GroupID []string

	RestoreCfg RestoreCfgOpts
	ExportCfg  ExportCfgOpts

	Populated flags.PopulatedFlags
}

func MakeGroupOpts(cmd *cobra.Command) GroupOpts {
	return GroupOpts{
		GroupID: flags.GroupFV,

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// ValidateGroupRestoreFlags checks common flags for correctness and interdependencies
func ValidateGroupRestoreFlags(backupID string, opts GroupOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	// TODO(meain): selectors (refer sharepoint)

	return validateRestoreConfigFlags(flags.CollisionsFV, opts.RestoreCfg)
}

// AddGroupInfo adds the scope of the provided values to the selector's
// filter set
func AddGroupInfo(
	sel *selectors.GroupsRestore,
	v string,
	f func(string) []selectors.GroupsScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// IncludeGroupRestoreDataSelectors builds the common data-selector
// inclusions for Group commands.
func IncludeGroupRestoreDataSelectors(ctx context.Context, opts GroupOpts) *selectors.GroupsRestore {
	groups := opts.GroupID

	ls := len(opts.GroupID)

	if ls == 0 {
		groups = selectors.Any()
	}

	sel := selectors.NewGroupsRestore(groups)

	// TODO(meain): add selectors
	sel.Include(sel.AllData())

	return sel
}

// FilterGroupRestoreInfoSelectors builds the common info-selector filters.
func FilterGroupRestoreInfoSelectors(
	sel *selectors.GroupsRestore,
	opts GroupOpts,
) {
	// TODO(meain)
	// AddGroupInfo(sel, opts.GroupID, sel.Library)
}
