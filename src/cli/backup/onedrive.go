package backup

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const oneDriveServiceCommand = "onedrive"

// called by backup.go to map parent subcommands to provider-specific handling.
func addOneDriveCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case createCommand:
		c, fs = utils.AddCommand(parent, oneDriveCreateCmd())

		fs.StringArrayVar(&user, "user", nil,
			"Backup OneDrive data by user ID; accepts "+utils.Wildcard+" to select all users")
		options.AddOperationFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(parent, oneDriveListCmd())

	case detailsCommand:
		c, fs = utils.AddCommand(parent, oneDriveDetailsCmd())
		fs.StringVar(&backupID, "backup", "", "ID of the backup containing the details to be shown")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

	case deleteCommand:
		c, fs = utils.AddCommand(parent, oneDriveDeleteCmd())
		fs.StringVar(&backupID, "backup", "", "ID of the backup containing the details to be shown")
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
		Use:   oneDriveServiceCommand,
		Short: "Backup M365 OneDrive service data",
		RunE:  createOneDriveCmd,
		Args:  cobra.NoArgs,
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
		Use:   oneDriveServiceCommand,
		Short: "Shows the details of a M365 OneDrive service backup",
		RunE:  detailsOneDriveCmd,
		Args:  cobra.NoArgs,
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

	ds, _, err := r.BackupDetails(ctx, backupID)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to get backup details in the repository"))
	}

	// TODO: Support selectors and filters

	ds.PrintEntries(ctx)

	return nil
}

// `corso backup delete onedrive [<flag>...]`
func oneDriveDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   oneDriveServiceCommand,
		Short: "Delete backed-up M365 OneDrive service data",
		RunE:  deleteOneDriveCmd,
		Args:  cobra.NoArgs,
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
