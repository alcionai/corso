package backup

import (
	"os"

	"github.com/pkg/errors"
	"github.com/segmentio/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// exchange bucket info from flags
var (
	user            string
	backupDetailsID string
)

// called by backup.go to map parent subcommands to provider-specific handling.
func addExchangeCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)
	switch parent.Use {
	case createCommand:
		c, fs = utils.AddCommand(parent, exchangeCreateCmd)
		fs.StringVar(&user, "user", "", "ID of the user whose Exchange data is to be backed up.")
	case listCommand:
		c, _ = utils.AddCommand(parent, exchangeListCmd)
	case detailsCommand:
		c, fs = utils.AddCommand(parent, exchangeDetailsCmd)
		fs.StringVar(&backupDetailsID, "backup-details", "", "ID of the backup details to be shown.")
		cobra.CheckErr(c.MarkFlagRequired("backup-details"))
	}
	return c
}

const exchangeServiceCommand = "exchange"

// `corso backup create exchange [<flag>...]`
var exchangeCreateCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "Backup M365 Exchange service data",
	RunE:  createExchangeCmd,
	Args:  cobra.NoArgs,
}

// processes an exchange service backup.
func createExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
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

	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users(user))

	bo, err := r.NewBackup(ctx, sel.Selector)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange backup")
	}

	err = bo.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to run Exchange backup")
	}

	// todo: revive when backups are hooked up to backupOperation results
	// fmt.Printf("Created backup %s in %s for Exchange user %s.\n", result.SnapshotID, s.Provider, user)
	return nil
}

// `corso backup list exchange [<flag>...]`
var exchangeListCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "List the history of M365 Exchange service backups",
	RunE:  listExchangeCmd,
	Args:  cobra.NoArgs,
}

// lists the history of backup operations
func listExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	rps, err := r.Backups(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to list backups in the repository")
	}

	// TODO: Can be used to print in alternative forms (e.g. json)
	p, err := cli.Format("text", os.Stdout)
	if err != nil {
		return err
	}
	defer p.Flush()
	for _, rp := range rps {
		p.Print(*rp)
	}
	return nil
}

// `corso backup details exchange [<flag>...]`
var exchangeDetailsCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "Shows the details of a M365 Exchange service backup",
	RunE:  detailsExchangeCmd,
	Args:  cobra.NoArgs,
}

// lists the history of backup operations
func detailsExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	rpd, err := r.BackupDetails(ctx, backupDetailsID)
	if err != nil {
		return errors.Wrap(err, "Failed to get backup details in the repository")
	}

	// TODO: Can be used to print in alternative forms
	p, err := cli.Format("json", os.Stdout)
	if err != nil {
		return err
	}
	defer p.Flush()
	p.Print(*rpd)

	return nil
}
