package cli

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/backup"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/help"
	"github.com/alcionai/corso/src/cli/options"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/restore"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ------------------------------------------------------------------------------------------
// Corso Command
// ------------------------------------------------------------------------------------------

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var corsoCmd = &cobra.Command{
	Use:               "corso",
	Short:             "Free, Secure, Open-Source Backup for M365.",
	Long:              `Free, Secure, and Open-Source Backup for Microsoft 365.`,
	RunE:              handleCorsoCmd,
	PersistentPreRunE: config.InitFunc(),
}

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) error {
	v, _ := cmd.Flags().GetBool("version")
	if v {
		print.Outf(cmd.Context(), "Corso\nversion: "+version.Version)
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

	cmd.Flags().BoolP("version", "v", false, "current version info")
	cmd.PersistentPostRunE = config.InitFunc()
	config.AddConfigFlags(cmd)
	logger.AddLogLevelFlag(cmd)
	observe.AddProgressBarFlags(cmd)
	print.AddOutputFlag(cmd)
	options.AddGlobalOperationFlags(cmd)

	cmd.SetUsageTemplate(indentExamplesTemplate(corsoCmd.UsageTemplate()))

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
	observe.SeedWriter(ctx, print.StderrWriter(ctx), observe.PreloadFlags())

	BuildCommandTree(corsoCmd)

	ctx, log := logger.Seed(ctx, logger.PreloadLogLevel())
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := corsoCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

// Adjust the default usage template which does not properly indent examples
func indentExamplesTemplate(template string) string {
	cobra.AddTemplateFunc("indent", func(spaces int, v string) string {
		pad := strings.Repeat(" ", spaces)
		return pad + strings.Replace(v, "\n", "\n"+pad, -1)
	})

	e := regexp.MustCompile(`{{\.Example}}`)

	return e.ReplaceAllString(template, "{{.Example | indent 2}}")
}
