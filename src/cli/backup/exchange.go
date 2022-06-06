package backup

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
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
	Short: "Start backing up Exchange data",
	Long:  `Creates a new backup of your microsoft Exchange data.`,
	Run:   createExchangeCmd,
	Args:  cobra.NoArgs,
}

// initializes a s3 repo.
func createExchangeCmd(cmd *cobra.Command, args []string) {
	mv := utils.GetM365Vars()
	fmt.Printf(
		"Called - %s\n\t365TenantID:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n",
		cmd.CommandPath(),
		mv.TenantID,
		mv.ClientID,
		len(mv.ClientSecret) > 0)

	a := repository.Account{
		TenantID:     mv.TenantID,
		ClientID:     mv.ClientID,
		ClientSecret: mv.ClientSecret,
	}
	// todo (rkeepers) - retrieve storage details from corso config
	s, err := storage.NewStorage(storage.ProviderUnknown)
	if err != nil {
		utils.Fatalf("Failed to configure storage provider: %v", err)
	}

	r, err := repository.Connect(cmd.Context(), a, s)
	if err != nil {
		utils.Fatalf("Failed to connect to the %s repository: %v", s.Provider, err)
	}
	defer utils.CloseRepo(cmd.Context(), r)

	bo, err := r.NewBackup(cmd.Context(), []string{user})
	if err != nil {
		utils.Fatalf("Failed to initialize Exchange backup: %v", err)
	}

	if err := bo.Run(cmd.Context()); err != nil {
		utils.Fatalf("Failed to run Exchange backup: %v", err)
	}

	fmt.Printf("Backed up Exchange in %s for user %s.\n", s.Provider, user)
}
