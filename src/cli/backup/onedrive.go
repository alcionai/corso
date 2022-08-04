package backup

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/options"
	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const onedriveServiceCommand = "onedrive"

// called by backup.go to map parent subcommands to provider-specific handling.
func addOneDriveCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {

	case createCommand:
		c, fs = utils.AddCommand(parent, onedriveCreateCmd)
		fs.StringArrayVar(&user, "user", nil, "Backup OneDrive data by user ID; accepts "+utils.Wildcard+" to select all users")
		options.AddOperationFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(parent, onedriveListCmd)

	case detailsCommand:
		c, fs = utils.AddCommand(parent, onedriveDetailsCmd)
		fs.StringVar(&backupID, "backup", "", "ID of the backup containing the details to be shown")
		cobra.CheckErr(c.MarkFlagRequired("backup"))
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create onedrive [<flag>...]`
var onedriveCreateCmd = &cobra.Command{
	Use:   onedriveServiceCommand,
	Short: "Backup M365 OneDrive service data",
	RunE:  createOneDriveCmd,
	Args:  cobra.NoArgs,
}

// processes an onedrive service backup.
func createOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// if err := validateExchangeBackupCreateFlags(exchangeAll, user, exchangeData); err != nil {
	// 	return err
	// }

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	sel := onedriveBackupCreateSelectors(user)

	bo, err := r.NewBackup(ctx, sel, options.Control())
	if err != nil {
		return Only(errors.Wrap(err, "Failed to initialize OneDrive backup"))
	}

	err = bo.Run(ctx)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to run OneDrive backup"))
	}

	bu, err := r.Backup(ctx, bo.Results.BackupID)
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve backup results from storage")
	}

	OutputBackup(*bu)
	return nil
}

func onedriveBackupCreateSelectors(users []string) selectors.Selector {
	sel := selectors.NewOneDriveBackup()
	sel.Include(sel.Users(users))
	return sel.Selector
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list onedrive [<flag>...]`
var onedriveListCmd = &cobra.Command{
	Use:   onedriveServiceCommand,
	Short: "List the history of M365 OneDrive service backups",
	RunE:  listOneDriveCmd,
	Args:  cobra.NoArgs,
}

// lists the history of backup operations
func listOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return Only(errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}
	defer utils.CloseRepo(ctx, r)

	rps, err := r.Backups(ctx)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to list backups in the repository"))
	}

	OutputBackups(rps)
	return nil
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details onedrive [<flag>...]`
var onedriveDetailsCmd = &cobra.Command{
	Use:   onedriveServiceCommand,
	Short: "Shows the details of a M365 OneDrive service backup",
	RunE:  detailsOneDriveCmd,
	Args:  cobra.NoArgs,
}

// lists the history of backup operations
func detailsOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return Only(errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}
	defer utils.CloseRepo(ctx, r)

	ds, _, err := r.BackupDetails(ctx, backupID)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to get backup details in the repository"))
	}

	OutputEntries(ds.Entries)
	return nil
}
