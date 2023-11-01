package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var cmd = &cobra.Command{
	Use:               "device",
	Short:             "device token POC",
	RunE:              help,
	PersistentPreRunE: cli.PreRun,
}

var request = &cobra.Command{
	Use:   "request",
	Short: "request device token POC",
	RunE:  requestToken,
}

var get = &cobra.Command{
	Use:   "get",
	Short: "get device token POC",
	RunE:  getToken,
}

func main() {
	cli.AddSupportFlags(cmd)
	ctx, log := cli.SeedCtx()

	defer func() {
		observe.Flush(ctx) // flush the progress bars

		_ = log.Sync() // flush all logs in the buffer
	}()

	cmd.AddCommand(request)
	cmd.AddCommand(get)

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func help(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func requestToken(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	_, details, err := utils.GetAccountAndConnect(
		ctx,
		cmd,
		path.ExchangeService)
	if err != nil {
		return Only(ctx, err)
	}

	creds, err := details.Repo.Account.M365Config()
	if err != nil {
		return Only(ctx, err)
	}

	ac, err := api.NewClient(
		creds,
		details.Opts,
		count.New())
	if err != nil {
		return Only(ctx, err)
	}

	da, err := ac.Access().RequestDeviceToken(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	PrettyJSON(ctx, da)

	return nil
}

func getToken(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	_, details, err := utils.GetAccountAndConnect(
		ctx,
		cmd,
		path.ExchangeService)
	if err != nil {
		return Only(ctx, err)
	}

	creds, err := details.Repo.Account.M365Config()
	if err != nil {
		return Only(ctx, err)
	}

	ac, err := api.NewClient(
		creds,
		details.Opts,
		count.New())
	if err != nil {
		return Only(ctx, err)
	}

	da, err := ac.Access().GetDeviceToken(ctx, args[0])
	if err != nil {
		return Only(ctx, err)
	}

	PrettyJSON(ctx, da)

	return nil
}
