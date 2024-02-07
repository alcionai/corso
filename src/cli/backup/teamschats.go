package backup

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/alcionai/canario/src/cli/flags"
	. "github.com/alcionai/canario/src/cli/print"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/filters"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
	"github.com/alcionai/canario/src/pkg/services/m365"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	teamschatsServiceCommand                 = "chats"
	teamschatsServiceCommandCreateUseSuffix  = "--user <userEmail> | '" + flags.Wildcard + "'"
	teamschatsServiceCommandDeleteUseSuffix  = "--backups <backupId>"
	teamschatsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	teamschatsServiceCommandCreateExamples = `# Backup all chats with bob@company.hr
corso backup create chats --user bob@company.hr

# Backup all chats for all users
corso backup create chats --user '*'`

	teamschatsServiceCommandDeleteExamples = `# Delete chats backup with ID 1234abcd-12ab-cd34-56de-1234abcd \
and 1234abcd-12ab-cd34-56de-1234abce
corso backup delete chats --backups 1234abcd-12ab-cd34-56de-1234abcd,1234abcd-12ab-cd34-56de-1234abce`

	teamschatsServiceCommandDetailsExamples = `# Explore chats in Bob's latest backup (1234abcd...)
corso backup details chats --backup 1234abcd-12ab-cd34-56de-1234abcd`
)

// called by backup.go to map subcommands to provider-specific handling.
func addTeamsChatsCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case createCommand:
		c, _ = utils.AddCommand(cmd, teamschatsCreateCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + teamschatsServiceCommandCreateUseSuffix
		c.Example = teamschatsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		flags.AddUserFlag(c)
		flags.AddDataFlag(c, []string{flags.DataChats}, false)
		flags.AddGenericBackupFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(cmd, teamschatsListCmd(), utils.MarkPreReleaseCommand())

		flags.AddBackupIDFlag(c, false)
		flags.AddAllBackupListFlags(c)

	case detailsCommand:
		c, _ = utils.AddCommand(cmd, teamschatsDetailsCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + teamschatsServiceCommandDetailsUseSuffix
		c.Example = teamschatsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddTeamsChatsDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, _ = utils.AddCommand(cmd, teamschatsDeleteCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + teamschatsServiceCommandDeleteUseSuffix
		c.Example = teamschatsServiceCommandDeleteExamples

		flags.AddMultipleBackupIDsFlag(c, false)
		flags.AddBackupIDFlag(c, false)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create chats [<flag>...]`
func teamschatsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     teamschatsServiceCommand,
		Aliases: []string{teamsServiceCommand},
		Short:   "Backup M365 Chats data",
		RunE:    createTeamsChatsCmd,
		Args:    cobra.NoArgs,
	}
}

// processes a teamschats backup.
func createTeamsChatsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := validateTeamsChatsBackupCreateFlags(flags.UserFV, flags.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		path.TeamsChatsService)
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

	ins, err := svcCli.AC.Users().GetAllIDsAndNames(ctx, errs)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 teamschats"))
	}

	sel := teamschatsBackupCreateSelectors(ctx, ins, flags.UserFV, flags.CategoryDataFV)
	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return genericCreateCommand(
		ctx,
		r,
		"Chats",
		selectorSet,
		ins)
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list teamschats [<flag>...]`
func teamschatsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamschatsServiceCommand,
		Short: "List the history of M365 Chats backups",
		RunE:  listTeamsChatsCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listTeamsChatsCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.TeamsChatsService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details teamschats [<flag>...]`
func teamschatsDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamschatsServiceCommand,
		Short: "Shows the details of a M365 Chats backup",
		RunE:  detailsTeamsChatsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a teamschats backup.
func detailsTeamsChatsCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	return runDetailsTeamsChatsCmd(cmd)
}

func runDetailsTeamsChatsCmd(cmd *cobra.Command) error {
	ctx := cmd.Context()
	opts := utils.MakeTeamsChatsOpts(cmd)

	sel := utils.IncludeTeamsChatsRestoreDataSelectors(ctx, opts)
	sel.Configure(selectors.Config{OnlyMatchItemNames: true})
	utils.FilterTeamsChatsRestoreInfoSelectors(sel, opts)

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

// `corso backup delete teamschats [<flag>...]`
func teamschatsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamschatsServiceCommand,
		Short: "Delete backed-up M365 Chats data",
		RunE:  deleteTeamsChatsCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an teamschats backup.
func deleteTeamsChatsCmd(cmd *cobra.Command, args []string) error {
	backupIDValue := []string{}

	if len(flags.BackupIDsFV) > 0 {
		backupIDValue = flags.BackupIDsFV
	} else if len(flags.BackupIDFV) > 0 {
		backupIDValue = append(backupIDValue, flags.BackupIDFV)
	} else {
		return clues.New("either --backup or --backups flag is required")
	}

	return genericDeleteCommand(cmd, path.TeamsChatsService, "TeamsChats", backupIDValue, args)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func validateTeamsChatsBackupCreateFlags(teamschats, cats []string) error {
	if len(teamschats) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.UserFN + " ids, or the wildcard --" +
				flags.UserFN + " *")
	}

	msg := fmt.Sprintf(
		" is an unrecognized data type; only %s is supported",
		flags.DataChats)

	allowedCats := utils.TeamsChatsAllowedCategories()

	for _, d := range cats {
		if _, ok := allowedCats[d]; !ok {
			return clues.New(d + msg)
		}
	}

	return nil
}

func teamschatsBackupCreateSelectors(
	ctx context.Context,
	ins idname.Cacher,
	users, cats []string,
) *selectors.TeamsChatsBackup {
	if filters.PathContains(users).Compare(flags.Wildcard) {
		return includeAllTeamsChatsWithCategories(ins, cats)
	}

	sel := selectors.NewTeamsChatsBackup(slices.Clone(users))

	return utils.AddTeamsChatsCategories(sel, cats)
}

func includeAllTeamsChatsWithCategories(ins idname.Cacher, categories []string) *selectors.TeamsChatsBackup {
	return utils.AddTeamsChatsCategories(selectors.NewTeamsChatsBackup(ins.IDs()), categories)
}
