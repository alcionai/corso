package backup

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	dataContacts = "contacts"
	dataEmail    = "email"
	dataEvents   = "events"
)

const (
	exchangeServiceCommand                 = "exchange"
	exchangeServiceCommandCreateUseSuffix  = "--mailbox <email> | '" + flags.Wildcard + "'"
	exchangeServiceCommandDeleteUseSuffix  = "--backups <backupId>"
	exchangeServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	exchangeServiceCommandCreateExamples = `# Backup all Exchange data for Alice
corso backup create exchange --mailbox alice@example.com

# Backup only Exchange contacts for Alice and Bob
corso backup create exchange --mailbox alice@example.com,bob@example.com --data contacts

# Backup all Exchange data for all M365 users 
corso backup create exchange --mailbox '*'`

	exchangeServiceCommandDeleteExamples = `# Delete Exchange backup with IDs 1234abcd-12ab-cd34-56de-1234abcd \
and 1234abcd-12ab-cd34-56de-1234abce
corso backup delete exchange --backups 1234abcd-12ab-cd34-56de-1234abcd,1234abcd-12ab-cd34-56de-1234abce`

	exchangeServiceCommandDetailsExamples = `# Explore items in Alice's latest backup (1234abcd...)
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore emails in the folder "Inbox" with subject containing "Hello world"
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --email-subject "Hello world" --email-folder Inbox

# Explore calendar events occurring after start of 2022
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-starts-after 2022-01-01T00:00:00

# Explore contacts named Andy
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --contact-name Andy`
)

// called by backup.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, exchangeCreateCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + exchangeServiceCommandCreateUseSuffix
		c.Example = exchangeServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddMailBoxFlag(c)
		flags.AddDataFlag(c, []string{dataEmail, dataContacts, dataEvents}, false)
		flags.AddFetchParallelismFlag(c)
		flags.AddDisableDeltaFlag(c)
		flags.AddEnableImmutableIDFlag(c)
		flags.AddDisableConcurrencyLimiterFlag(c)
		flags.AddDeltaPageSizeFlag(c)
		flags.AddGenericBackupFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, exchangeListCmd())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddAllBackupListFlags(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, exchangeDetailsCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + exchangeServiceCommandDetailsUseSuffix
		c.Example = exchangeServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddExchangeDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, exchangeDeleteCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + exchangeServiceCommandDeleteUseSuffix
		c.Example = exchangeServiceCommandDeleteExamples

		flags.AddMultipleBackupIDsFlag(c, false)
		flags.AddBackupIDFlag(c, false)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create exchange [<flag>...]`
func exchangeCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Backup M365 Exchange service data",
		RunE:  createExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// processes an exchange service backup.
func createExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := validateExchangeBackupCreateFlags(flags.UserFV, flags.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		path.ExchangeService)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	sel := exchangeBackupCreateSelectors(flags.UserFV, flags.CategoryDataFV)

	ins, err := utils.UsersMap(
		ctx,
		*acct,
		utils.Control(),
		r.Counter(),
		fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 users"))
	}

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return genericCreateCommand(
		ctx,
		r,
		"Exchange",
		selectorSet,
		ins)
}

func exchangeBackupCreateSelectors(userIDs, cats []string) *selectors.ExchangeBackup {
	sel := selectors.NewExchangeBackup(userIDs)

	if len(cats) == 0 {
		sel.Include(sel.ContactFolders(selectors.Any()))
		sel.Include(sel.MailFolders(selectors.Any()))
		sel.Include(sel.EventCalendars(selectors.Any()))
	}

	for _, d := range cats {
		switch d {
		case dataContacts:
			sel.Include(sel.ContactFolders(selectors.Any()))
		case dataEmail:
			sel.Include(sel.MailFolders(selectors.Any()))
		case dataEvents:
			sel.Include(sel.EventCalendars(selectors.Any()))
		}
	}

	return sel
}

func validateExchangeBackupCreateFlags(userIDs, cats []string) error {
	if len(userIDs) == 0 {
		return clues.New("--user/--mailbox requires one or more email addresses or the wildcard '*'")
	}

	for _, d := range cats {
		if d != dataContacts && d != dataEmail && d != dataEvents {
			return clues.New(
				d + " is an unrecognized data type; must be one of " + dataContacts + ", " + dataEmail + ", or " + dataEvents)
		}
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list exchange [<flag>...]`
func exchangeListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "List the history of M365 Exchange service backups",
		RunE:  listExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listExchangeCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.ExchangeService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details exchange [<flag>...]`
func exchangeDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Shows the details of a M365 Exchange service backup",
		RunE:  detailsExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// lists all items in the backup, running the results first through
// selector reduction as a filtering step.
func detailsExchangeCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	return runDetailsExchangeCmd(cmd)
}

func runDetailsExchangeCmd(cmd *cobra.Command) error {
	ctx := cmd.Context()
	opts := utils.MakeExchangeOpts(cmd)

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	sel.Configure(selectors.Config{OnlyMatchItemNames: true})
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

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

// `corso backup delete exchange [<flag>...]`
func exchangeDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Delete backed-up M365 Exchange service data",
		RunE:  deleteExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an exchange service backup.
func deleteExchangeCmd(cmd *cobra.Command, args []string) error {
	var backupIDValue []string

	if len(flags.BackupIDsFV) > 0 {
		backupIDValue = flags.BackupIDsFV
	} else if len(flags.BackupIDFV) > 0 {
		backupIDValue = append(backupIDValue, flags.BackupIDFV)
	} else {
		return clues.New("either --backup or --backups flag is required")
	}

	return genericDeleteCommand(cmd, path.ExchangeService, "Exchange", backupIDValue, args)
}
