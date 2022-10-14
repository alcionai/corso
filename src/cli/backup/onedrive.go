package backup

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	oneDriveServiceCommand                 = "onedrive"
	oneDriveServiceCommandCreateUseSuffix  = "--user <userId or email> | '" + utils.Wildcard + "'"
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

// called by backup.go to map parent subcommands to provider-specific handling.
func addOneDriveCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case createCommand:
		c, fs = utils.AddCommand(parent, oneDriveCreateCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandCreateUseSuffix
		c.Example = oneDriveServiceCommandCreateExamples

		fs.StringArrayVar(&user, "user", nil,
			"Backup OneDrive data by user ID; accepts '"+utils.Wildcard+"' to select all users. (required)")
		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(parent, oneDriveListCmd())

		fs.StringVar(&backupID,
			"backup", "",
			"ID of the backup to retrieve.")

	case detailsCommand:
		c, fs = utils.AddCommand(parent, oneDriveDetailsCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandDetailsUseSuffix
		c.Example = oneDriveServiceCommandDetailsExamples

		fs.StringVar(&backupID, "backup", "", "ID of the backup to explore. (required)")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

		// onedrive hierarchy flags

		fs.StringSliceVar(
			&folderPaths,
			"folder", nil,
			"Select backup details by OneDrive folder; defaults to root.")

		fs.StringSliceVar(
			&fileNames,
			"file", nil,
			"Select backup details by file name or ID.")

		// onedrive info flags

		fs.StringVar(
			&fileCreatedAfter,
			"file-created-after", "",
			"Select backup details for files created after this datetime.")
		fs.StringVar(
			&fileCreatedBefore,
			"file-created-before", "",
			"Select backup details for files created before this datetime.")

		fs.StringVar(
			&fileModifiedAfter,
			"file-modified-after", "",
			"Select backup details for files modified after this datetime.")
		fs.StringVar(
			&fileModifiedBefore,
			"file-modified-before", "",
			"Select backup details for files modified before this datetime.")

	case deleteCommand:
		c, fs = utils.AddCommand(parent, oneDriveDeleteCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandDeleteUseSuffix
		c.Example = oneDriveServiceCommandDeleteExamples

		fs.StringVar(&backupID, "backup", "", "ID of the backup to delete. (required)")
		cobra.CheckErr(c.MarkFlagRequired("backup"))
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

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	sel := oneDriveBackupCreateSelectors(user)

	bo, err := r.NewBackup(ctx, sel)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize OneDrive backup"))
	}

	err = bo.Run(ctx)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to run OneDrive backup"))
	}

	bu, err := r.Backup(ctx, bo.Results.BackupID)
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve backup results from storage")
	}

	bu.Print(ctx)

	return nil
}

func validateOneDriveBackupCreateFlags(users []string) error {
	if len(users) == 0 {
		return errors.New("requires one or more --user ids or the wildcard --user *")
	}

	return nil
}

func oneDriveBackupCreateSelectors(users []string) selectors.Selector {
	sel := selectors.NewOneDriveBackup()
	sel.Include(sel.Users(users))

	return sel.Selector
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

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if len(backupID) > 0 {
		b, err := r.Backup(ctx, model.StableID(backupID))
		if err != nil {
			if errors.Is(err, kopia.ErrNotFound) {
				return Only(ctx, errors.Errorf("No backup exists with the id %s", backupID))
			}

			return Only(ctx, errors.Wrap(err, "Failed to find backup "+backupID))
		}

		b.Print(ctx)

		return nil
	}

	bs, err := r.Backups(ctx)
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

// lists the history of backup operations
func detailsOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	opts := utils.OneDriveOpts{
		Users:          user,
		Paths:          folderPaths,
		Names:          fileNames,
		CreatedAfter:   fileCreatedAfter,
		CreatedBefore:  fileCreatedBefore,
		ModifiedAfter:  fileModifiedAfter,
		ModifiedBefore: fileModifiedBefore,
	}

	ds, err := runDetailsOneDriveCmd(ctx, r, backupID, opts)
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

// runDetailsOneDriveCmd actually performs the lookup in backup details. Assumes
// len(backupID) > 0.
func runDetailsOneDriveCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.OneDriveOpts,
) (*details.Details, error) {
	d, _, err := r.BackupDetails(ctx, backupID)
	if err != nil {
		if errors.Is(err, kopia.ErrNotFound) {
			return nil, errors.Errorf("no backup exists with the id %s", backupID)
		}

		return nil, errors.Wrap(err, "Failed to get backup details in the repository")
	}

	sel := selectors.NewOneDriveRestore()
	utils.IncludeOneDriveRestoreDataSelectors(sel, opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	// if no selector flags were specified, get all data in the service.
	if len(sel.Scopes()) == 0 {
		sel.Include(sel.Users(selectors.Any()))
	}

	return sel.Reduce(ctx, d), nil
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

// deletes an exchange service backup.
func deleteOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackup(ctx, model.StableID(backupID)); err != nil {
		return Only(ctx, errors.Wrapf(err, "Deleting backup %s", backupID))
	}

	return nil
}
