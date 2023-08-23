package utils

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GroupsOpts struct {
	Groups []string

	RestoreCfg RestoreCfgOpts
	ExportCfg  ExportCfgOpts

	Populated flags.PopulatedFlags
}

func MakeGroupsOpts(cmd *cobra.Command) GroupsOpts {
	return GroupsOpts{
		Groups: flags.UserFV,

		RestoreCfg: makeRestoreCfgOpts(cmd),
		ExportCfg:  makeExportCfgOpts(cmd),

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// ValidateGroupsRestoreFlags checks common flags for correctness and interdependencies
func ValidateGroupsRestoreFlags(backupID string, opts GroupsOpts) error {
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

// IncludeGroupsRestoreDataSelectors builds the common data-selector
// inclusions for Group commands.
func IncludeGroupsRestoreDataSelectors(ctx context.Context, opts GroupsOpts) *selectors.GroupsRestore {
	groups := opts.Groups

	ls := len(opts.Groups)

	if ls == 0 {
		groups = selectors.Any()
	}

	sel := selectors.NewGroupsRestore(groups)

	// TODO(meain): add selectors
	sel.Include(sel.AllData())

	return sel
}

// FilterGroupsRestoreInfoSelectors builds the common info-selector filters.
func FilterGroupsRestoreInfoSelectors(
	sel *selectors.GroupsRestore,
	opts GroupsOpts,
) {
	// TODO(meain)
	// AddGroupInfo(sel, opts.GroupID, sel.Library)
}
