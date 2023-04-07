package backup

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
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
	exchangeServiceCommandCreateUseSuffix  = "--user <email> | '" + utils.Wildcard + "'"
	exchangeServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	exchangeServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	exchangeServiceCommandCreateExamples = `# Backup all Exchange data for Alice
corso backup create exchange --user alice@example.com

# Backup only Exchange contacts for Alice and Bob
corso backup create exchange --user alice@example.com,bob@example.com --data contacts

# Backup all Exchange data for all M365 users 
corso backup create exchange --user '*'`

	exchangeServiceCommandDeleteExamples = `# Delete Exchange backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete exchange --backup 1234abcd-12ab-cd34-56de-1234abcd`

	exchangeServiceCommandDetailsExamples = `# Explore Alice's items in backup 1234abcd-12ab-cd34-56de-1234abcd 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --user alice@example.com

# Explore Alice's emails with subject containing "Hello world" in folder "Inbox" from a specific backup 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --user alice@example.com --email-subject "Hello world" --email-folder Inbox

# Explore Bobs's events occurring after start of 2022 from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --user bob@example.com --event-starts-after 2022-01-01T00:00:00

# Explore Alice's contacts with name containing Andy from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --user alice@example.com --contact-name Andy`
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

		options.AddFeatureToggle(cmd, options.DisableIncrementals())

		c.Use = c.Use + " " + exchangeServiceCommandCreateUseSuffix
		c.Example = exchangeServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		utils.AddUserFlag(c)
		utils.AddDataFlag(c, []string{dataEmail, dataContacts, dataEvents}, false)
		options.AddFetchParallelismFlag(c)
		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, exchangeListCmd())
		fs.SortFlags = false

		utils.AddBackupIDFlag(c, false)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, exchangeDetailsCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + exchangeServiceCommandDetailsUseSuffix
		c.Example = exchangeServiceCommandDetailsExamples

		options.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		utils.AddBackupIDFlag(c, true)
		utils.AddExchangeDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, exchangeDeleteCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + exchangeServiceCommandDeleteUseSuffix
		c.Example = exchangeServiceCommandDeleteExamples

		utils.AddBackupIDFlag(c, true)
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

	if err := validateExchangeBackupCreateFlags(utils.UserFV, utils.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	sel := exchangeBackupCreateSelectors(utils.UserFV, utils.CategoryDataFV)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	ins, err := m365.UsersMap(ctx, *acct, errs)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 users"))
	}

	fmt.Printf("\n-----\nINS %+v\n-----\n", ins)

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return runBackups(
		ctx,
		r,
		"Exchange", "user",
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
		return clues.New("--user requires one or more email addresses or the wildcard '*'")
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
	return genericListCommand(cmd, utils.BackupIDFV, path.ExchangeService, args)
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

	ctx := cmd.Context()
	opts := utils.MakeExchangeOpts(cmd)

	r, _, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ctrlOpts := options.Control()

	ds, err := runDetailsExchangeCmd(ctx, r, utils.BackupIDFV, opts, ctrlOpts.SkipReduce)
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

// runDetailsExchangeCmd actually performs the lookup in backup details.
// the fault.Errors return is always non-nil.  Callers should check if
// errs.Failure() == nil.
func runDetailsExchangeCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.ExchangeOpts,
	skipReduce bool,
) (*details.Details, error) {
	if err := utils.ValidateExchangeRestoreFlags(backupID, opts); err != nil {
		return nil, err
	}

	ctx = clues.Add(ctx, "backup_id", backupID)

	d, _, errs := r.GetBackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Failure() != nil {
		if errors.Is(errs.Failure(), data.ErrNotFound) {
			return nil, clues.New("No backup exists with the id " + backupID)
		}

		return nil, clues.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	ctx = clues.Add(ctx, "details_entries", len(d.Entries))

	if !skipReduce {
		sel := utils.IncludeExchangeRestoreDataSelectors(opts)
		utils.FilterExchangeRestoreInfoSelectors(sel, opts)
		d = sel.Reduce(ctx, d, errs)
	}

	return d, nil
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
	return genericDeleteCommand(cmd, utils.BackupIDFV, "Exchange", args)
}
