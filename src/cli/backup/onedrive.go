package backup

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	oneDriveServiceCommand                 = "onedrive"
	oneDriveServiceCommandCreateUseSuffix  = "--user <email> | '" + flags.Wildcard + "'"
	oneDriveServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	oneDriveServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	oneDriveServiceCommandCreateExamples = `# Backup OneDrive data for Alice
corso backup create onedrive --user alice@example.com

# Backup OneDrive for Alice and Bob
corso backup create onedrive --user alice@example.com,bob@example.com

# Backup all OneDrive data for all M365 users 
corso backup create onedrive --user '*'`

	oneDriveServiceCommandDeleteExamples = `# Delete OneDrive backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd`

	oneDriveServiceCommandDetailsExamples = `# Explore items in Bob's latest backup (1234abcd...)
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore files in the folder "Reports" named "Fiscal 22"
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file-name "Fiscal 22" --folder "Reports"

# Explore files created before the end of 2015
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file-created-before 2015-01-01T00:00:00`
)

// called by backup.go to map subcommands to provider-specific handling.
func addOneDriveCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, oneDriveCreateCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + oneDriveServiceCommandCreateUseSuffix
		c.Example = oneDriveServiceCommandCreateExamples

		flags.AddUserFlag(c)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)

		flags.AddFailFastFlag(c)
		flags.AddDisableIncrementalsFlag(c)
		flags.AddForceItemDataDownloadFlag(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, oneDriveListCmd())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, oneDriveDetailsCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + oneDriveServiceCommandDetailsUseSuffix
		c.Example = oneDriveServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)
		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		flags.AddOneDriveDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, oneDriveDeleteCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + oneDriveServiceCommandDeleteUseSuffix
		c.Example = oneDriveServiceCommandDeleteExamples

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

// `corso backup create onedrive [<flag>...]`
func oneDriveCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Backup M365 OneDrive service data",
		RunE:    createOneDriveCmd,
		Args:    cobra.NoArgs,
		Example: oneDriveServiceCommandCreateExamples,
	}
}

// processes an onedrive service backup.
func createOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateOneDriveBackupCreateFlags(flags.UserFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(ctx, path.OneDriveService, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	sel := oneDriveBackupCreateSelectors(flags.UserFV)

	ins, err := utils.UsersMap(ctx, *acct, utils.Control(), fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 users"))
	}

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return runBackups(
		ctx,
		r,
		"OneDrive",
		selectorSet,
		ins)
}

func validateOneDriveBackupCreateFlags(users []string) error {
	if len(users) == 0 {
		return clues.New("requires one or more --user ids or the wildcard --user *")
	}

	return nil
}

func oneDriveBackupCreateSelectors(users []string) *selectors.OneDriveBackup {
	sel := selectors.NewOneDriveBackup(users)
	sel.Include(sel.AllData())

	return sel
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list onedrive [<flag>...]`
func oneDriveListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   oneDriveServiceCommand,
		Short: "List the history of M365 OneDrive service backups",
		RunE:  listOneDriveCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listOneDriveCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.OneDriveService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details onedrive [<flag>...]`
func oneDriveDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Shows the details of a M365 OneDrive service backup",
		RunE:    detailsOneDriveCmd,
		Args:    cobra.NoArgs,
		Example: oneDriveServiceCommandDetailsExamples,
	}
}

// prints the item details for a given backup
func detailsOneDriveCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctx := cmd.Context()
	opts := utils.MakeOneDriveOpts(cmd)

	r, _, _, ctrlOpts, err := utils.GetAccountAndConnect(ctx, path.OneDriveService, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ds, err := runDetailsOneDriveCmd(ctx, r, flags.BackupIDFV, opts, ctrlOpts.SkipReduce)
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

// runDetailsOneDriveCmd actually performs the lookup in backup details.
// the fault.Errors return is always non-nil.  Callers should check if
// errs.Failure() == nil.
func runDetailsOneDriveCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.OneDriveOpts,
	skipReduce bool,
) (*details.Details, error) {
	if err := utils.ValidateOneDriveRestoreFlags(backupID, opts); err != nil {
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
		sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
		sel.Configure(selectors.Config{OnlyMatchItemNames: true})
		utils.FilterOneDriveRestoreInfoSelectors(sel, opts)
		d = sel.Reduce(ctx, d, errs)
	}

	return d, nil
}

// `corso backup delete onedrive [<flag>...]`
func oneDriveDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Delete backed-up M365 OneDrive service data",
		RunE:    deleteOneDriveCmd,
		Args:    cobra.NoArgs,
		Example: oneDriveServiceCommandDeleteExamples,
	}
}

// deletes a oneDrive service backup.
func deleteOneDriveCmd(cmd *cobra.Command, args []string) error {
	return genericDeleteCommand(cmd, path.OneDriveService, flags.BackupIDFV, "OneDrive", args)
}
