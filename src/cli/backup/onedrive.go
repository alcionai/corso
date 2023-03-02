package backup

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
	"github.com/alcionai/corso/src/pkg/store"
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

	fileCreatedAfter   string
	fileCreatedBefore  string
	fileModifiedAfter  string
	fileModifiedBefore string
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
			&fileCreatedAfter,
			utils.FileCreatedAfterFN, "",
			"Select backup details for files created after this datetime.")
		fs.StringVar(
			&fileCreatedBefore,
			utils.FileCreatedBeforeFN, "",
			"Select backup details for files created before this datetime.")

		fs.StringVar(
			&fileModifiedAfter,
			utils.FileModifiedAfterFN, "",
			"Select backup details for files modified after this datetime.")
		fs.StringVar(
			&fileModifiedBefore,
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

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	sel := oneDriveBackupCreateSelectors(user)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	users, err := m365.UserPNs(ctx, cfg.Account, errs)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to retrieve M365 users"))
	}

	var (
		merrs *multierror.Error
		bIDs  []model.StableID
	)

	for _, discSel := range sel.SplitByResourceOwner(users) {
		bo, err := r.NewBackup(ctx, discSel.Selector)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to initialize OneDrive backup for user %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		err = bo.Run(ctx)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to run OneDrive backup for user %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		bIDs = append(bIDs, bo.Results.BackupID)
	}

	bups, ferrs := r.Backups(ctx, bIDs)
	// TODO: print/log recoverable errors
	if ferrs.Failure() != nil {
		return Only(ctx, errors.Wrap(ferrs.Failure(), "Unable to retrieve backup results from storage"))
	}

	backup.PrintAll(ctx, bups)

	if e := merrs.ErrorOrNil(); e != nil {
		return Only(ctx, e)
	}

	return nil
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
	ctx := cmd.Context()

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if len(backupID) > 0 {
		b, err := r.Backup(ctx, model.StableID(backupID))
		if err != nil {
			if errors.Is(err, data.ErrNotFound) {
				return Only(ctx, errors.Errorf("No backup exists with the id %s", backupID))
			}

			return Only(ctx, errors.Wrap(err, "Failed to find backup "+backupID))
		}

		b.Print(ctx)

		return nil
	}

	bs, err := r.BackupsByTag(ctx, store.Service(path.OneDriveService))
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to list backups in the repository"))
	}

	backup.PrintAll(ctx, bs)

	return nil
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
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}


	ctrlOpts := options.Control()

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, ctrlOpts)
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	opts := utils.OneDriveOpts{
		Users:              user,
		Paths:              folderPaths,
		Names:              fileNames,
		FileCreatedAfter:   fileCreatedAfter,
		FileCreatedBefore:  fileCreatedBefore,
		FileModifiedAfter:  fileModifiedAfter,
		FileModifiedBefore: fileModifiedBefore,

		Populated: utils.GetPopulatedFlags(cmd),
	}

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

	d, _, errs := r.BackupDetails(ctx, backupID)
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
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackup(ctx, model.StableID(backupID)); err != nil {
		return Only(ctx, errors.Wrapf(err, "Deleting backup %s", backupID))
	}

	Info(ctx, "Deleted OneDrive backup ", backupID)

	return nil
}
