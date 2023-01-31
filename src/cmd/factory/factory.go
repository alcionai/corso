package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cmd/factory/impl"
	"github.com/alcionai/corso/src/pkg/logger"
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

// ------------------------------------------------------------------------------------------
// CLI command handlers
// ------------------------------------------------------------------------------------------

func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)
	ctx = SetRootCmd(ctx, factoryCmd)

	defer logger.Flush(ctx)

	// persistent flags that are common to all use cases
	fs := factoryCmd.PersistentFlags()
	fs.StringVar(&impl.Tenant, "tenant", "", "m365 tenant containing the user")
	fs.StringVar(&impl.User, "user", "", "m365 user owning the new data")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("user"))
	fs.IntVar(&impl.Count, "count", 0, "count of items to produce")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("count"))
	fs.StringVar(&impl.Destination, "destination", "", "destination of the new data (will create as needed)")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("destination"))

	factoryCmd.AddCommand(exchangeCmd)
	impl.AddExchangeCommands(exchangeCmd)
	factoryCmd.AddCommand(oneDriveCmd)
	impl.AddOneDriveCommands(oneDriveCmd)

	if err := factoryCmd.ExecuteContext(ctx); err != nil {
		logger.Flush(ctx)
		os.Exit(1)
	}
}

func handleFactoryRoot(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), impl.ErrNotYetImplemeted)
	return cmd.Help()
}

func handleExchangeFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), impl.ErrNotYetImplemeted)
	return cmd.Help()
}

func handleOneDriveFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), impl.ErrNotYetImplemeted)
	return cmd.Help()
}
