package restore

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
)

// exchange bucket info from flags
var (
	folder         string
	mail           string
	restorePointID string
	user           string
)

// called by restore.go to map parent subcommands to provider-specific handling.
func addExchangeApp(parent *cobra.Command) *cobra.Command {
	parent.AddCommand(exchangeCmd)

	fs := exchangeCmd.Flags()
	fs.StringVar(&folder, "folder", "", "Name of the mail folder being restored")
	fs.StringVar(&mail, "mail", "", "ID of the mail message being restored")
	fs.StringVar(&restorePointID, "restore-point", "", "ID of the backup restore point")
	exchangeCmd.MarkFlagRequired("restore-point")
	fs.StringVar(&user, "user", "", "ID of the user whose exchange data will get restored")

	return exchangeCmd
}

// `corso restore create exchange [<flag>...]`
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Restore M365 Exchange service data",
	RunE:  createExchangeCmd,
	Args:  cobra.NoArgs,
}

// initializes a s3 repo.
func createExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateRestoreFlags(user, folder, mail, restorePointID); err != nil {
		return errors.Wrap(err, "Missing required flags")
	}

	s, cfgTenantID, err := config.MakeS3Config(true, nil)
	if err != nil {
		return err
	}

	m365 := credentials.GetM365()
	a := repository.Account{
		TenantID:     m365.TenantID,
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
	}
	if len(cfgTenantID) > 0 {
		a.TenantID = cfgTenantID
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"restorePointID", restorePointID,
		"tenantID", m365.TenantID,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0)

	r, err := repository.Connect(ctx, a, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	ro, err := r.NewRestore(ctx, restorePointID, []string{cfgTenantID, user, "mail", folder, mail})
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange restore")
	}

	if err := ro.Run(ctx); err != nil {
		return errors.Wrap(err, "Failed to run Exchange restore")
	}

	fmt.Printf("Restored Exchange in %s for user %s.\n", s.Provider, user)
	return nil
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
