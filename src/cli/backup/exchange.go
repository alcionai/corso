package backup

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// exchange bucket info from flags
var (
	user            []string
	backupDetailsID string
)

const (
	dataContacts = "contacts"
	dataEmail    = "email"
	dataEvents   = "events"
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
		fs.StringArrayVar(&user, "user", nil, "Back up Exchange data by user ID; accepts * to select all users")
		fs.BoolVar(&all, "all", false, "Back up all Exchange data for all users")
		fs.StringArrayVar(
			&data,
			"data",
			nil,
			"Select one or more types of data to backup: "+dataEmail+", "+dataContacts+", or "+dataEvents)
	case listCommand:
		c, _ = utils.AddCommand(parent, exchangeListCmd)
	case detailsCommand:
		c, fs = utils.AddCommand(parent, exchangeDetailsCmd)
		fs.StringVar(&backupDetailsID, "backup-details", "", "ID of the backup details to be shown")
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
	if err := validateBackupCreateFlags(all, user, data); err != nil {
		return err
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

	sel := exchangeBackupCreateSelectors(all, user, data)

	bo, err := r.NewBackup(ctx, sel)
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

func exchangeBackupCreateSelectors(all bool, users, data []string) selectors.Selector {
	sel := selectors.NewExchangeBackup()
	if all {
		sel.Include(sel.Users(selectors.All))
		return sel.Selector
	}
	if len(data) == 0 {
		for _, user := range users {
			if user == "*" {
				user = selectors.All
			}
			sel.Include(sel.ContactFolders(user, selectors.All))
			sel.Include(sel.MailFolders(user, selectors.All))
			sel.Include(sel.Events(user, selectors.All))
		}
	}
	for _, d := range data {
		switch d {
		case dataContacts:
			for _, user := range users {
				if user == "*" {
					user = selectors.All
				}
				sel.Include(sel.ContactFolders(user, selectors.All))
			}
		case dataEmail:
			for _, user := range users {
				if user == "*" {
					user = selectors.All
				}
				sel.Include(sel.MailFolders(user, selectors.All))
			}
		case dataEvents:
			for _, user := range users {
				if user == "*" {
					user = selectors.All
				}
				sel.Include(sel.Events(user, selectors.All))
			}
		}
	}
	return sel.Selector
}

func validateBackupCreateFlags(all bool, users, data []string) error {
	if len(users) == 0 && !all {
		return errors.New("requries one or more --user ids, the wildcard --user *, or the --all flag.")
	}
	if len(data) > 0 && all {
		return errors.New("--all backs up all data, and cannot be reduced with --data")
	}
	for _, d := range data {
		if d != dataContacts && d != dataEmail && d != dataEvents {
			return errors.New(d + " is an unrecognized data type; must be one of " + dataContacts + ", " + dataEmail + ", or " + dataEvents)
		}
	}
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

	print.Backups(rps)

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

	print.Entries(rpd.Entries)

	return nil
}
