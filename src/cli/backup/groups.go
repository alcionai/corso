package backup

import (
	"context"
	"errors"
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
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const dataMessages = "messages"

const (
	groupsServiceCommand                 = "groups"
	groupsServiceCommandCreateUseSuffix  = "--group <groupsName> | '" + flags.Wildcard + "'"
	groupsServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	groupsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

// TODO: correct examples
const (
	groupsServiceCommandCreateExamples = `# Backup all Groups data for Alice
corso backup create groups --group alice@example.com 

# Backup only Groups contacts for Alice and Bob
corso backup create groups --group engineering,sales --data contacts

# Backup all Groups data for all M365 users 
corso backup create groups --group '*'`

	groupsServiceCommandDeleteExamples = `# Delete Groups backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete groups --backup 1234abcd-12ab-cd34-56de-1234abcd`

	groupsServiceCommandDetailsExamples = `# Explore items in Alice's latest backup (1234abcd...)
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore calendar events occurring after start of 2022
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-starts-after 2022-01-01T00:00:00`
)

// called by backup.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, groupsCreateCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandCreateUseSuffix
		c.Example = groupsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		flags.AddGroupFlag(c)
		flags.AddDataFlag(c, []string{dataLibraries, dataMessages}, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		flags.AddFetchParallelismFlag(c)
		flags.AddFailFastFlag(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, groupsListCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, groupsDetailsCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDetailsUseSuffix
		c.Example = groupsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, groupsDeleteCmd(), utils.MarkPreReleaseCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDeleteUseSuffix
		c.Example = groupsServiceCommandDeleteExamples

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

// `corso backup create groups [<flag>...]`
func groupsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Backup M365 Group service data",
		RunE:  createGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a groups service backup.
func createGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateGroupsBackupCreateFlags(flags.GroupFV, flags.CategoryDataFV); err != nil {
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
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 groups"))
	}

	sel := groupsBackupCreateSelectors(ctx, ins, flags.GroupFV, flags.CategoryDataFV)
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

	ctx := cmd.Context()
	opts := utils.MakeGroupsOpts(cmd)

	r, _, _, ctrlOpts, err := utils.GetAccountAndConnect(ctx, path.GroupsService, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ds, err := runDetailsGroupsCmd(ctx, r, flags.BackupIDFV, opts, ctrlOpts.SkipReduce)
	if err != nil {
		return Only(ctx, err)
	}

	if len(ds.Entries) == 0 {
		Info(ctx, selectors.ErrorNoMatchingItems)
		return nil
	}

	ds.PrintEntries(ctx)

	return nil
}

// runDetailsGroupsCmd actually performs the lookup in backup details.
// the fault.Errors return is always non-nil.  Callers should check if
// errs.Failure() == nil.
func runDetailsGroupsCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.GroupsOpts,
	skipReduce bool,
) (*details.Details, error) {
	if err := utils.ValidateGroupsRestoreFlags(backupID, opts); err != nil {
		return nil, err
	}

	ctx = clues.Add(ctx, "backup_id", backupID)

	d, _, errs := r.GetBackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Failure() != nil {
		if errors.Is(errs.Failure(), data.ErrNotFound) {
			return nil, clues.New("no backup exists with the id " + backupID)
		}

		return nil, clues.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	ctx = clues.Add(ctx, "details_entries", len(d.Entries))

	if !skipReduce {
		sel := utils.IncludeGroupsRestoreDataSelectors(ctx, opts)
		sel.Configure(selectors.Config{OnlyMatchItemNames: true})
		utils.FilterGroupsRestoreInfoSelectors(sel, opts)
		d = sel.Reduce(ctx, d, errs)
	}

	return d, nil
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
	return genericDeleteCommand(cmd, path.GroupsService, flags.BackupIDFV, "Groups", args)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func validateGroupsBackupCreateFlags(groups, cats []string) error {
	if len(groups) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.GroupFN + " ids, or the wildcard --" +
				flags.GroupFN + " *",
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

func groupsBackupCreateSelectors(
	ctx context.Context,
	ins idname.Cacher,
	group, cats []string,
) *selectors.GroupsBackup {
	if filters.PathContains(group).Compare(flags.Wildcard) {
		return includeAllGroupWithCategories(ins, cats)
	}

	sel := selectors.NewGroupsBackup(slices.Clone(group))

	return addGroupsCategories(sel, cats)
}

func includeAllGroupWithCategories(ins idname.Cacher, categories []string) *selectors.GroupsBackup {
	return addGroupsCategories(selectors.NewGroupsBackup(ins.IDs()), categories)
}

func addGroupsCategories(sel *selectors.GroupsBackup, cats []string) *selectors.GroupsBackup {
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
