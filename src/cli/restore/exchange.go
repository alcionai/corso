package restore

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/options"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// exchange bucket info from flags
var (
	emailFolder string
	email       string
	backupID    string
	user        string
)

// called by restore.go to map parent subcommands to provider-specific handling.
func addExchangeCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(parent, exchangeRestoreCmd)
		fs.StringVar(&emailFolder, "email-folder", "", "Name of the email folder being restored")
		fs.StringVar(&email, "email", "", "ID of the email being restored")
		fs.StringVar(&backupID, "backup", "", "ID of the backup to restore")
		cobra.CheckErr(c.MarkFlagRequired("backup"))
		fs.StringVar(&user, "user", "", "ID of the user whose exchange data will get restored")
		options.AddOperationFlags(c)
	}
	return c
}

const exchangeServiceCommand = "exchange"

// `corso restore exchange [<flag>...]`
var exchangeRestoreCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "Restore M365 Exchange service data",
	RunE:  restoreExchangeCmd,
	Args:  cobra.NoArgs,
}

// processes an exchange service restore.
func restoreExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateRestoreFlags(user, emailFolder, email, backupID); err != nil {
		return errors.Wrap(err, "Missing required flags")
	}

	s, a, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := a.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"backupID", backupID,
		"tenantID", m365.TenantID,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0)

	r, err := repository.Connect(ctx, a, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	ro, err := r.NewRestore(ctx, backupID, exchangeRestoreSelectors(user, emailFolder, email), options.OperationOptions())
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange restore")
	}

	if err := ro.Run(ctx); err != nil {
		return errors.Wrap(err, "Failed to run Exchange restore")
	}

	fmt.Printf("Restored Exchange in %s for user %s.\n", s.Provider, user)
	return nil
}

func exchangeRestoreSelectors(u, f, m string) selectors.Selector {
	sel := selectors.NewExchangeRestore()
	if len(m) > 0 {
		sel.Include(sel.Mails(
			[]string{u}, []string{f}, []string{m},
		))
	}
	if len(f) > 0 && len(m) == 0 {
		sel.Include(sel.MailFolders(
			[]string{u}, []string{f},
		))
	}
	if len(f) == 0 && len(m) == 0 {
		sel.Include(sel.Users([]string{u}))
	}
	return sel.Selector
}

func validateRestoreFlags(u, f, m, rpid string) error {
	if len(rpid) == 0 {
		return errors.New("a restore point ID is requried")
	}
	lu, lf, lm := len(u), len(f), len(m)
	if (lu == 0 || u == "*") && (lf+lm > 0) {
		return errors.New("a specific --user must be provided if --email-folder or --email is specified")
	}
	if (lf == 0 || f == "*") && lm > 0 {
		return errors.New("a specific --email-folder must be provided if a --email is specified")
	}
	return nil
}
