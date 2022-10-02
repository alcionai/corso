package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/backup"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/help"
	"github.com/alcionai/corso/src/cli/options"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/restore"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ------------------------------------------------------------------------------------------
// Corso Command
// ------------------------------------------------------------------------------------------

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var corsoCmd = &cobra.Command{
	Use:               "corso",
	Short:             "Protect your Microsoft 365 data.",
	Long:              `Reliable, secure, and efficient data protection for Microsoft 365.`,
	RunE:              handleCorsoCmd,
	PersistentPreRunE: config.InitFunc(),
}

// the root-level flags
var (
	version bool
)

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) error {
	if version {
		print.Infof(cmd.Context(), "Corso\nversion:\tpre-alpha\n")
		return nil
	}

	return cmd.Help()
}

// CorsoCommand produces a copy of the cobra command used by Corso.
// The command tree is built and attached to the returned command.
func CorsoCommand() *cobra.Command {
	c := &cobra.Command{}
	*c = *corsoCmd
	BuildCommandTree(c)

	return c
}

// BuildCommandTree builds out the command tree used by the Corso library.
func BuildCommandTree(cmd *cobra.Command) {
	// want to order flags explicitly
	cmd.PersistentFlags().SortFlags = false

	cmd.Flags().BoolP("version", "v", version, "current version info")
	cmd.PersistentPostRunE = config.InitFunc()
	config.AddConfigFlags(cmd)
	logger.AddLogLevelFlag(cmd)
	print.AddOutputFlag(cmd)
	options.AddGlobalOperationFlags(cmd)

	cmd.CompletionOptions.DisableDefaultCmd = true

	repo.AddCommands(cmd)
	backup.AddCommands(cmd)
	restore.AddCommands(cmd)
	help.AddCommands(cmd)
}

// ------------------------------------------------------------------------------------------
// Running Corso
// ------------------------------------------------------------------------------------------

// Handle builds and executes the cli processor.
func Handle() {
	ctx := config.Seed(context.Background())
	ctx = print.SetRootCmd(ctx, corsoCmd)

	BuildCommandTree(corsoCmd)

	ctx, log := logger.Seed(ctx)
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := corsoCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
