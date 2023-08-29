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
	"github.com/alcionai/corso/src/cli/repo"
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
	teamsServiceCommand                 = "teams"
	teamsServiceCommandCreateUseSuffix  = "--team <teamsName> | '" + flags.Wildcard + "'"
	teamsServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	teamsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

// TODO: correct examples
const (
	teamsServiceCommandCreateExamples = `# Backup all Teams data for Alice
corso backup create teams --team alice@example.com 

# Backup only Teams contacts for Alice and Bob
corso backup create teams --team engineering,sales --data contacts

# Backup all Teams data for all M365 users 
corso backup create teams --team '*'`

	teamsServiceCommandDeleteExamples = `# Delete Teams backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete teams --backup 1234abcd-12ab-cd34-56de-1234abcd`

	teamsServiceCommandDetailsExamples = `# Explore items in Alice's latest backup (1234abcd...)
corso backup details teams --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore calendar events occurring after start of 2022
corso backup details teams --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-starts-after 2022-01-01T00:00:00`
)

// called by backup.go to map subcommands to provider-specific handling.
func addTeamsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, teamsCreateCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandCreateUseSuffix
		c.Example = teamsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		flags.AddTeamFlag(c)
		flags.AddDataFlag(c, []string{dataEmail, dataContacts, dataEvents}, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		flags.AddFetchParallelismFlag(c)
		flags.AddFailFastFlag(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, teamsListCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, teamsDetailsCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandDetailsUseSuffix
		c.Example = teamsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, teamsDeleteCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandDeleteUseSuffix
		c.Example = teamsServiceCommandDeleteExamples

		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create teams [<flag>...]`
func teamsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Backup M365 Team service data",
		RunE:  createTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a teams service backup.
func createTeamsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateTeamsBackupCreateFlags(flags.TeamFV, flags.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(ctx, path.GroupsService, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	ins, err := m365.GroupsMap(ctx, *acct, errs)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 teams"))
	}

	sel := teamsBackupCreateSelectors(ctx, ins, flags.TeamFV, flags.CategoryDataFV)
	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return runBackups(
		ctx,
		r,
		"Group",
		selectorSet,
		ins)
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list teams [<flag>...]`
func teamsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "List the history of M365 Teams service backups",
		RunE:  listTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listTeamsCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.TeamsService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details teams [<flag>...]`
func teamsDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Shows the details of a M365 Teams service backup",
		RunE:  detailsTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a teams service backup.
func detailsTeamsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	return Only(ctx, utils.ErrNotYetImplemented)
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete teams [<flag>...]`
func teamsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Delete backed-up M365 Teams service data",
		RunE:  deleteTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an teams service backup.
func deleteTeamsCmd(cmd *cobra.Command, args []string) error {
	return genericDeleteCommand(cmd, path.TeamsService, flags.BackupIDFV, "Teams", args)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func validateTeamsBackupCreateFlags(teams, cats []string) error {
	if len(teams) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.TeamFN + " ids, or the wildcard --" +
				flags.TeamFN + " *",
		)
	}

	msg := fmt.Sprintf(
		" is an unrecognized data type; only %s and %s are supported",
		dataLibraries, dataMessages)

	allowedCats := map[string]struct{}{
		dataLibraries: {},
		dataMessages:  {},
	}

	for _, d := range cats {
		if _, ok := allowedCats[d]; !ok {
			return clues.New(d + msg)
		}
	}

	return nil
}

func teamsBackupCreateSelectors(
	ctx context.Context,
	ins idname.Cacher,
	team, cats []string,
) *selectors.GroupsBackup {
	if filters.PathContains(team).Compare(flags.Wildcard) {
		return includeAllTeamWithCategories(ins, cats)
	}

	sel := selectors.NewGroupsBackup(slices.Clone(team))

	return addTeamsCategories(sel, cats)
}

func includeAllTeamWithCategories(ins idname.Cacher, categories []string) *selectors.GroupsBackup {
	return addTeamsCategories(selectors.NewGroupsBackup(ins.IDs()), categories)
}

func addTeamsCategories(sel *selectors.GroupsBackup, cats []string) *selectors.GroupsBackup {
	if len(cats) == 0 {
		sel.Include(sel.AllData())
	}

	for _, d := range cats {
		switch d {
		case dataLibraries:
			sel.Include(sel.LibraryFolders(selectors.Any()))
		case dataMessages:
			sel.Include(sel.ChannelMessages(selectors.Any(), selectors.Any()))
		}
	}

	return sel
}
