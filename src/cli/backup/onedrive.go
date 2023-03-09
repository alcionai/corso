package backup

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/clues"
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
	oneDriveServiceCommand                 = "onedrive"
	oneDriveServiceCommandCreateUseSuffix  = "--user <email> | '" + utils.Wildcard + "'"
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

	oneDriveServiceCommandDetailsExamples = `# Explore Alice's files from backup 1234abcd-12ab-cd34-56de-1234abcd 
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --user alice@example.com

# Explore Alice or Bob's files with name containing "Fiscal 22" in folder "Reports"
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com,bob@example.com  --file-name "Fiscal 22" --folder "Reports"

# Explore Alice's files created before end of 2015 from a specific backup
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --file-created-before 2015-01-01T00:00:00`
)

var (
	folderPaths []string
	fileNames   []string
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
		options.AddFeatureToggle(cmd, options.EnablePermissionsBackup())

		c.Use = c.Use + " " + oneDriveServiceCommandCreateUseSuffix
		c.Example = oneDriveServiceCommandCreateExamples

		fs.StringSliceVar(&user,
			utils.UserFN, nil,
			"Backup OneDrive data by user's email address; accepts '"+utils.Wildcard+"' to select all users. (required)")
		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, oneDriveListCmd())

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to retrieve.")

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, oneDriveDetailsCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandDetailsUseSuffix
		c.Example = oneDriveServiceCommandDetailsExamples

		options.AddSkipReduceFlag(c)

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to explore. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))

		// onedrive hierarchy flags

		fs.StringSliceVar(
			&folderPaths,
			utils.FolderFN, nil,
			"Select backup details by OneDrive folder; defaults to root.")

		fs.StringSliceVar(
			&fileNames,
			utils.FileFN, nil,
			"Select backup details by file name or ID.")

		// onedrive info flags

		fs.StringVar(
			&utils.FileCreatedAfter,
			utils.FileCreatedAfterFN, "",
			"Select backup details for files created after this datetime.")
		fs.StringVar(
			&utils.FileCreatedBefore,
			utils.FileCreatedBeforeFN, "",
			"Select backup details for files created before this datetime.")

		fs.StringVar(
			&utils.FileModifiedAfter,
			utils.FileModifiedAfterFN, "",
			"Select backup details for files modified after this datetime.")
		fs.StringVar(
			&utils.FileModifiedBefore,
			utils.FileModifiedBeforeFN, "",
			"Select backup details for files modified before this datetime.")

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, oneDriveDeleteCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandDeleteUseSuffix
		c.Example = oneDriveServiceCommandDeleteExamples

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to delete. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))
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

	if err := validateOneDriveBackupCreateFlags(user); err != nil {
		return err
	}

	r, acct, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	sel := oneDriveBackupCreateSelectors(user)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	users, err := m365.UserPNs(ctx, *acct, errs)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to retrieve M365 users"))
	}

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(users) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return runBackups(
		ctx,
		r,
		"OneDrive", "user",
		selectorSet,
	)
}

func validateOneDriveBackupCreateFlags(users []string) error {
	if len(users) == 0 {
		return errors.New("requires one or more --user ids or the wildcard --user *")
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
	return genericListCommand(cmd, backupID, path.OneDriveService, args)
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
	opts := utils.OneDriveOpts{
		Users:              user,
		Names:              fileNames,
		Paths:              folderPaths,
		FileCreatedAfter:   utils.FileCreatedAfter,
		FileCreatedBefore:  utils.FileCreatedBefore,
		FileModifiedAfter:  utils.FileModifiedAfter,
		FileModifiedBefore: utils.FileModifiedBefore,

		Populated: utils.GetPopulatedFlags(cmd),
	}

	r, _, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ctrlOpts := options.Control()

	ds, err := runDetailsOneDriveCmd(ctx, r, backupID, opts, ctrlOpts.SkipReduce)
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
			return nil, errors.Errorf("no backup exists with the id %s", backupID)
		}

		return nil, errors.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	if !skipReduce {
		sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
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
	return genericDeleteCommand(cmd, backupID, "OneDrive", args)
}
