package backup

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/source"
)

// exchange bucket info from flags
var (
	user string
)

// called by backup.go to map parent subcommands to provider-specific handling.
func addExchangeApp(parent *cobra.Command) *cobra.Command {
	var c *cobra.Command
	switch parent.Use {
	case createCommand:
		c = exchangeCreateCmd
	}
	parent.AddCommand(c)
	fs := c.Flags()
	fs.StringVar(&user, "user", "", "ID of the user whose Exchange data is to be backed up.")
	return c
}

// `corso backup create exchange [<flag>...]`
var exchangeCreateCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Backup M365 Exchange",
	RunE:  createExchangeCmd,
	Args:  cobra.NoArgs,
}

// initializes a s3 repo.
func createExchangeCmd(cmd *cobra.Command, args []string) error {
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

	fmt.Printf(
		"Called - %s\n\t365TenantID:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n",
		cmd.CommandPath(),
		m365.TenantID,
		m365.ClientID,
		len(m365.ClientSecret) > 0)

	r, err := repository.Connect(cmd.Context(), a, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(cmd.Context(), r)

	src := source.NewSource(source.ServiceExchange)
	err = src.AddUsers(user)
	if err != nil {
		return errors.Wrap(err, "Failed to configure backup data source")
	}
	bo, err := r.NewBackup(cmd.Context(), src)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange backup")
	}

	if err := bo.Run(cmd.Context()); err != nil {
		return errors.Wrap(err, "Failed to run Exchange backup")
	}

	fmt.Printf("Backed up Exchange in %s for user %s.\n", s.Provider, user)
	return nil
}
