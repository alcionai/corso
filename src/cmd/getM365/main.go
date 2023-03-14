package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cmd/getM365/exchange"
	"github.com/alcionai/corso/src/pkg/logger"
)

var user, tenant string

var rootCmd = &cobra.Command{
	Use: "getM365",
}

func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)

	ctx = SetRootCmd(ctx, rootCmd)
	defer logger.Flush(ctx)

	exchange.AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
