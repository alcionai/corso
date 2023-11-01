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
	Use:               "delegated",
	Short:             "delegated token POC",
	RunE:              getToken,
	PersistentPreRunE: cli.PreRun,
}

func main() {
	cli.AddSupportFlags(cmd)
	ctx, log := cli.SeedCtx()

	defer func() {
		observe.Flush(ctx) // flush the progress bars

		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
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

	da, err := ac.Access().GetDelegatedToken(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	PrettyJSON(ctx, da)

	return nil
}
