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

	rootCmd.PersistentFlags().StringVar(&user, "user", "", "m365 user id of M365 user")
	rootCmd.PersistentFlags().StringVar(&tenant, "tenant", "",
		"m365 Tenant: m365 identifier for the tenant, not required if active in OS Environment")

	exchange.AddCommands(rootCmd, user, tenant)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
