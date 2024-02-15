package utils

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type TeamsChatsOpts struct {
	Users []string

	ExportCfg ExportCfgOpts

	Populated flags.PopulatedFlags
}

func TeamsChatsAllowedCategories() map[string]struct{} {
	return map[string]struct{}{
		flags.DataChats: {},
	}
}

func AddTeamsChatsCategories(sel *selectors.TeamsChatsBackup, cats []string) *selectors.TeamsChatsBackup {
	if len(cats) == 0 {
		sel.Include(sel.AllData())
	}

	for _, d := range cats {
		switch d {
		case flags.DataChats:
			sel.Include(sel.Chats(selectors.Any()))
		}
	}

	return sel
}

func MakeTeamsChatsOpts(cmd *cobra.Command) TeamsChatsOpts {
	return TeamsChatsOpts{
		Users: flags.UserFV,

		ExportCfg: makeExportCfgOpts(cmd),

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// ValidateTeamsChatsRestoreFlags checks common flags for correctness and interdependencies
func ValidateTeamsChatsRestoreFlags(backupID string, opts TeamsChatsOpts, isRestore bool) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	// restore isn't currently supported
	if isRestore {
		return clues.New("restore not supported")
	}

	return nil
}

// AddTeamsChatsFilter adds the scope of the provided values to the selector's
// filter set
func AddTeamsChatsFilter(
	sel *selectors.TeamsChatsRestore,
	v string,
	f func(string) []selectors.TeamsChatsScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// IncludeTeamsChatsRestoreDataSelectors builds the common data-selector
// inclusions for teamschats commands.
func IncludeTeamsChatsRestoreDataSelectors(ctx context.Context, opts TeamsChatsOpts) *selectors.TeamsChatsRestore {
	users := opts.Users

	if len(opts.Users) == 0 {
		users = selectors.Any()
	}

	sel := selectors.NewTeamsChatsRestore(users)
	sel.Include(sel.Chats(selectors.Any()))

	return sel
}

// FilterTeamsChatsRestoreInfoSelectors builds the common info-selector filters.
func FilterTeamsChatsRestoreInfoSelectors(
	sel *selectors.TeamsChatsRestore,
	opts TeamsChatsOpts,
) {
	// TODO: populate when adding filters
}
