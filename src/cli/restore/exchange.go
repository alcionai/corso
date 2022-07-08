package restore

import (
	"fmt"

	"github.com/pkg/errors"
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
	folder   string
	mail     string
	backupID string
	user     string
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
		fs.StringVar(&folder, "folder", "", "Name of the mail folder being restored")
		fs.StringVar(&mail, "mail", "", "ID of the mail message being restored")
		fs.StringVar(&backupID, "backup", "", "ID of the backup to restore")
		cobra.CheckErr(c.MarkFlagRequired("backup"))
		fs.StringVar(&user, "user", "", "ID of the user whose exchange data will get restored")
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

	if err := validateRestoreFlags(user, folder, mail, backupID); err != nil {
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

	sel := restoreSelectors()
	// []string{m365.TenantID, user, "mail", folder, mail}

	ro, err := r.NewRestore(ctx, restorePointID, sel)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange restore")
	}

	if err := ro.Run(ctx); err != nil {
		return errors.Wrap(err, "Failed to run Exchange restore")
	}

	fmt.Printf("Restored Exchange in %s for user %s.\n", s.Provider, user)
	return nil
}

func restoreSelectors() selectors.Selector {
	sel := selectors.NewExchangeRestore()
	u := user
	if user == "*" {
		u = selectors.All
	}
	f := folder
	if folder == "*" {
		f = selectors.All
	}
	m := mail
	if mail == "*" {
		m = selectors.All
	}
	if len(m) > 0 {
		sel.Include(sel.Mails(u, f, m))
	}
	if len(f) > 0 && len(m) == 0 {
		sel.Include(sel.MailFolders(u, f))
	}
	if len(f) == 0 && len(m) == 0 {
		sel.Include(sel.Users(u))
	}
	return sel.Selector
}

func validateRestoreFlags(u, f, m, rpid string) error {
	if len(rpid) == 0 {
		return errors.New("a restore point ID is requried")
	}
	lu, lf, lm := len(u), len(f), len(m)
	if (lu == 0 || u == "*") && (lf+lm > 0) {
		return errors.New("a specific --user must be provided if --folder or --mail is specified")
	}
	if (lf == 0 || f == "*") && lm > 0 {
		return errors.New("a specific --folder must be provided if a --mail is specified")
	}
	return nil
}
