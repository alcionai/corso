package cli

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/alcionai/canario/src/cli/backup"
	"github.com/alcionai/canario/src/cli/debug"
	"github.com/alcionai/canario/src/cli/export"
	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/cli/help"
	"github.com/alcionai/canario/src/cli/print"
	"github.com/alcionai/canario/src/cli/repo"
	"github.com/alcionai/canario/src/cli/restore"
	"github.com/alcionai/canario/src/internal/observe"
	"github.com/alcionai/canario/src/internal/version"
	"github.com/alcionai/canario/src/pkg/config"
	"github.com/alcionai/canario/src/pkg/logger"
)

// ------------------------------------------------------------------------------------------
// Corso Command
// ------------------------------------------------------------------------------------------

// The root-level command.
// `canario <command> [<subcommand>] [<service>] [<flag>...]`
var canarioCmd = &cobra.Command{
	Use:               "canario",
	Short:             "Free, Secure, Open-Source Backup for M365.",
	Long:              `Free, Secure, and Open-Source Backup for Microsoft 365.`,
	RunE:              handleCorsoCmd,
	PersistentPreRunE: preRun,
}

func preRun(cc *cobra.Command, args []string) error {
	if err := config.InitCmd(cc, args); err != nil {
		return err
	}

	ctx := cc.Context()
	log := logger.Ctx(ctx)

	fs := flags.GetPopulatedFlags(cc)
	flagSl := make([]string, 0, len(fs))

	// currently only tracking flag names to avoid pii leakage.
	for f := range fs {
		flagSl = append(flagSl, f)
	}

	avoidTheseCommands := []string{
		"canario", "env", "help", "backup", "details", "list", "restore", "export", "delete", "repo", "init", "connect",
	}

	if len(logger.ResolvedLogFile) > 0 && !slices.Contains(avoidTheseCommands, cc.Use) {
		print.Infof(ctx, "Logging to file: %s", logger.ResolvedLogFile)
	}

	// handle deprecated user flag in Backup exchange command
	if cc.CommandPath() == "canario backup create exchange" {
		handleMailBoxFlag(ctx, cc, flagSl)
	}

	log.Infow("cli command", "command", cc.CommandPath(), "flags", flagSl, "version", version.CurrentVersion())

	return nil
}

func handleMailBoxFlag(ctx context.Context, c *cobra.Command, flagNames []string) {
	if !slices.Contains(flagNames, "user") && !slices.Contains(flagNames, "mailbox") {
		print.Err(ctx, "either --user or --mailbox flag is required")
		os.Exit(1)
	}

	if slices.Contains(flagNames, "user") && slices.Contains(flagNames, "mailbox") {
		print.Err(ctx, "cannot use both [mailbox, user] flags in the same command")
		os.Exit(1)
	}
}

// Handler for flat calls to `canario`.
// Produces the same output as `canario --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) error {
	v, _ := cmd.Flags().GetBool("version")
	if v {
		print.Outf(cmd.Context(), "Corso version: "+version.CurrentVersion())
		return nil
	}

	return cmd.Help()
}

// CorsoCommand produces a copy of the cobra command used by Corso.
// The command tree is built and attached to the returned command.
func CorsoCommand() *cobra.Command {
	c := &cobra.Command{}
	*c = *canarioCmd
	BuildCommandTree(c)

	return c
}

// BuildCommandTree builds out the command tree used by the Corso library.
func BuildCommandTree(cmd *cobra.Command) {
	// want to order flags explicitly
	cmd.PersistentFlags().SortFlags = false
	flags.AddRunModeFlag(cmd, true)

	cmd.Flags().BoolP("version", "v", false, "current version info")
	cmd.PersistentPreRunE = preRun
	config.AddConfigFlags(cmd)
	logger.AddLoggingFlags(cmd)
	observe.AddProgressBarFlags(cmd)
	print.AddOutputFlag(cmd)
	flags.AddGlobalOperationFlags(cmd)
	cmd.SetUsageTemplate(indentExamplesTemplate(canarioCmd.UsageTemplate()))

	cmd.CompletionOptions.DisableDefaultCmd = true

	repo.AddCommands(cmd)
	backup.AddCommands(cmd)
	restore.AddCommands(cmd)
	export.AddCommands(cmd)
	debug.AddCommands(cmd)
	help.AddCommands(cmd)
}

// ------------------------------------------------------------------------------------------
// Running Corso
// ------------------------------------------------------------------------------------------

// Handle builds and executes the cli processor.
func Handle() {
	//nolint:forbidigo
	ctx := config.Seed(context.Background())
	ctx, log := logger.Seed(ctx, logger.PreloadLoggingFlags(os.Args[1:]))
	ctx = print.SetRootCmd(ctx, canarioCmd)
	ctx = observe.SeedObserver(ctx, print.StderrWriter(ctx), observe.PreloadFlags())

	BuildCommandTree(canarioCmd)

	defer func() {
		observe.Flush(ctx) // flush the progress bars

		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := canarioCmd.ExecuteContext(ctx); err != nil {
		logger.CtxErr(ctx, err).Error("cli execution")
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
