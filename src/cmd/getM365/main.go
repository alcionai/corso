package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cmd/getM365/exchange"
	"github.com/alcionai/corso/src/cmd/getM365/onedrive"
	"github.com/alcionai/corso/src/pkg/logger"
)

var rootCmd = &cobra.Command{
	Use: "getM365",
}

func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)

	ctx = SetRootCmd(ctx, rootCmd)
	defer logger.Flush(ctx)

	exchange.AddCommands(rootCmd)
	onedrive.AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		Err(ctx, err)
		os.Exit(1)
	}
}
