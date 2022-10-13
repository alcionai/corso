package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
)

var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "Generate all types of m365 folders",
	RunE:  handleFactoryRoot,
}

var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Generate exchange data",
	RunE:  handleExchangeFactory,
}

var oneDriveCmd = &cobra.Command{
	Use:   "onedrive",
	Short: "Generate onedrive data",
	RunE:  handleOneDriveFactory,
}

var (
	count     int
	container string
	tenant    string
	user      string
)

// TODO: ErrGenerating       = errors.New("not all items were successfully generated")

var ErrNotYetImplemeted = errors.New("not yet implemented")

// ------------------------------------------------------------------------------------------
// CLI command handlers
// ------------------------------------------------------------------------------------------

func main() {
	ctx := SetRootCmd(context.Background(), factoryCmd)

	// persistent flags that are common to all use cases
	fs := factoryCmd.PersistentFlags()
	fs.StringVar(&tenant, "tenant", "", "m365 tenant containing the user")
	fs.StringVar(&user, "user", "", "m365 user owning the new data")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("user"))
	fs.IntVar(&count, "count", 0, "count of items to produce")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("count"))
	fs.StringVar(&container, "container", "", "container location of the new data (will create as needed)")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("container"))

	factoryCmd.AddCommand(exchangeCmd)
	addExchangeCommands(exchangeCmd)
	factoryCmd.AddCommand(oneDriveCmd)
	addOneDriveCommands(oneDriveCmd)

	if err := factoryCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func handleFactoryRoot(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}

func handleExchangeFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}

func handleOneDriveFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}
