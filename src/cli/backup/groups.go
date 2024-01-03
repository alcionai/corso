package backup

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	groupsServiceCommand                 = "groups"
	teamsServiceCommand                  = "teams"
	groupsServiceCommandCreateUseSuffix  = "--group <groupName> | '" + flags.Wildcard + "'"
	groupsServiceCommandDeleteUseSuffix  = "--backups <backupId>"
	groupsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	groupsServiceCommandCreateExamples = `# Backup all Groups and Teams data for the Marketing group
corso backup create groups --group Marketing

# Backup only Teams conversations messages
corso backup create groups --group Marketing --data messages

# Backup all Groups and Teams data for all groups
corso backup create groups --group '*'`

	groupsServiceCommandDeleteExamples = `# Delete Groups backup with ID 1234abcd-12ab-cd34-56de-1234abcd \
and 1234abcd-12ab-cd34-56de-1234abce
corso backup delete groups --backups 1234abcd-12ab-cd34-56de-1234abcd,1234abcd-12ab-cd34-56de-1234abce`

	groupsServiceCommandDetailsExamples = `# Explore items in Marketing's latest backup (1234abcd...)
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore Marketing messages posted after the start of 2022
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --last-message-reply-after 2022-01-01T00:00:00`
)

// called by backup.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, groupsCreateCmd(), utils.MarkPreviewCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandCreateUseSuffix
		c.Example = groupsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		flags.AddGroupFlag(c)
		flags.AddDataFlag(c, []string{flags.DataLibraries, flags.DataMessages, flags.DataConversations}, false)
		flags.AddFetchParallelismFlag(c)
		flags.AddDisableDeltaFlag(c)
		flags.AddGenericBackupFlags(c)
		flags.AddDisableLazyItemReader(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, groupsListCmd(), utils.MarkPreviewCommand())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddAllBackupListFlags(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, groupsDetailsCmd(), utils.MarkPreviewCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDetailsUseSuffix
		c.Example = groupsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddGroupDetailsAndRestoreFlags(c)
		flags.AddSharePointDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, groupsDeleteCmd(), utils.MarkPreviewCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDeleteUseSuffix
		c.Example = groupsServiceCommandDeleteExamples

		flags.AddMultipleBackupIDsFlag(c, false)
		flags.AddBackupIDFlag(c, false)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create groups [<flag>...]`
func groupsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     groupsServiceCommand,
		Aliases: []string{teamsServiceCommand},
		Short:   "Backup M365 Groups & Teams service data",
		RunE:    createGroupsCmd,
		Args:    cobra.NoArgs,
	}
}

// processes a groups service backup.
func createGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := validateGroupsBackupCreateFlags(flags.GroupFV, flags.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		path.GroupsService)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	svcCli, err := m365.NewM365Client(ctx, *acct)
	if err != nil {
		return Only(ctx, clues.Stack(err))
	}

	ins, err := svcCli.GroupsMap(ctx, errs)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 groups"))
	}

	sel := groupsBackupCreateSelectors(ctx, ins, flags.GroupFV, flags.CategoryDataFV)
	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return genericCreateCommand(
		ctx,
		r,
		"Group",
		selectorSet,
		ins)
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list groups [<flag>...]`
func groupsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "List the history of M365 Groups service backups",
		RunE:  listGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listGroupsCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.GroupsService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details groups [<flag>...]`
func groupsDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Shows the details of a M365 Groups service backup",
		RunE:  detailsGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a groups service backup.
func detailsGroupsCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	return runDetailsGroupsCmd(cmd)
}

func runDetailsGroupsCmd(cmd *cobra.Command) error {
	ctx := cmd.Context()
	opts := utils.MakeGroupsOpts(cmd)

	sel := utils.IncludeGroupsRestoreDataSelectors(ctx, opts)
	sel.Configure(selectors.Config{OnlyMatchItemNames: true})
	utils.FilterGroupsRestoreInfoSelectors(sel, opts)

	ds, err := genericDetailsCommand(cmd, flags.BackupIDFV, sel.Selector)
	if err != nil {
		return Only(ctx, err)
	}

	if len(ds.Entries) > 0 {
		ds.PrintEntries(ctx)
	} else {
		Info(ctx, selectors.ErrorNoMatchingItems)
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete groups [<flag>...]`
func groupsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Delete backed-up M365 Groups service data",
		RunE:  deleteGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an groups service backup.
func deleteGroupsCmd(cmd *cobra.Command, args []string) error {
	backupIDValue := []string{}

	if len(flags.BackupIDsFV) > 0 {
		backupIDValue = flags.BackupIDsFV
	} else if len(flags.BackupIDFV) > 0 {
		backupIDValue = append(backupIDValue, flags.BackupIDFV)
	} else {
		return clues.New("either --backup or --backups flag is required")
	}

	return genericDeleteCommand(cmd, path.GroupsService, "Groups", backupIDValue, args)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func validateGroupsBackupCreateFlags(groups, cats []string) error {
	if len(groups) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.GroupFN + " ids, or the wildcard --" +
				flags.GroupFN + " *")
	}

	// TODO(keepers): release conversations support

	msg := fmt.Sprintf(
		" is an unrecognized data type; only %s and %s are supported",
		flags.DataLibraries, flags.DataMessages)

	// msg := fmt.Sprintf(
	// 	" is an unrecognized data type; only %s, %s and %s are supported",
	// 	flags.DataLibraries, flags.DataMessages, flags.DataConversations)

	allowedCats := utils.GroupsAllowedCategories()

	for _, d := range cats {
		if _, ok := allowedCats[d]; !ok {
			return clues.New(d + msg)
		}
	}

	return nil
}

func groupsBackupCreateSelectors(
	ctx context.Context,
	ins idname.Cacher,
	group, cats []string,
) *selectors.GroupsBackup {
	if filters.PathContains(group).Compare(flags.Wildcard) {
		return includeAllGroupWithCategories(ins, cats)
	}

	sel := selectors.NewGroupsBackup(slices.Clone(group))

	return utils.AddGroupsCategories(sel, cats)
}

func includeAllGroupWithCategories(ins idname.Cacher, categories []string) *selectors.GroupsBackup {
	return utils.AddGroupsCategories(selectors.NewGroupsBackup(ins.IDs()), categories)
}
