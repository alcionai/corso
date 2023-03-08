package cli

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/backup"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/help"
	"github.com/alcionai/corso/src/cli/options"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/restore"
	"github.com/alcionai/corso/src/cli/utils"
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
	PersistentPreRunE: preRun,
}

func preRun(cc *cobra.Command, args []string) error {
	if err := config.InitFunc(cc, args); err != nil {
		return err
	}

	ctx := cc.Context()
	log := logger.Ctx(ctx)

	flags := utils.GetPopulatedFlags(cc)
	flagSl := make([]string, 0, len(flags))

	// currently only tracking flag names to avoid pii leakage.
	for f := range flags {
		flagSl = append(flagSl, f)
	}

	avoidTheseCommands := []string{
		"corso", "env", "help", "backup", "details", "list", "restore", "delete", "repo", "init", "connect",
	}

	if len(logger.LogFile) > 0 && !slices.Contains(avoidTheseCommands, cc.Use) {
		print.Info(ctx, "Logging to file: "+logger.LogFile)
	}

	avoidTheseDescription := []string{
		"Initialize a repository.",
		"Initialize a S3 repository",
		"Help about any command",
	}

	avoidTheseCommands = []string{
		"corso", "help", "init",
	}

	if !(slices.Contains(avoidTheseCommands, cc.Use) || slices.Contains(avoidTheseDescription, cc.Short)) {
		cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
		if err != nil {
			log.Error("Error while getting config info to run command: ", cc.Use)
			return err
		}

		utils.SendStartCorsoEvent(
			ctx,
			cfg.Storage,
			cfg.Account.ID(),
			map[string]any{"command": cc.CommandPath()},
			cfg.RepoID,
			options.Control())
	}

	log.Infow("cli command", "command", cc.CommandPath(), "flags", flagSl, "version", version.CurrentVersion())

	return nil
}

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
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
	*c = *corsoCmd
	BuildCommandTree(c)

	return c
}

// BuildCommandTree builds out the command tree used by the Corso library.
func BuildCommandTree(cmd *cobra.Command) {
	// want to order flags explicitly
	cmd.PersistentFlags().SortFlags = false

	cmd.Flags().BoolP("version", "v", false, "current version info")
	cmd.PersistentPreRunE = preRun
	config.AddConfigFlags(cmd)
	logger.AddLoggingFlags(cmd)
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
	//nolint:forbidigo
	ctx := config.Seed(context.Background())
	ctx = print.SetRootCmd(ctx, corsoCmd)
	observe.SeedWriter(ctx, print.StderrWriter(ctx), observe.PreloadFlags())

	BuildCommandTree(corsoCmd)

	loglevel, logfile := logger.PreloadLoggingFlags()
	ctx, log := logger.Seed(ctx, loglevel, logfile)

	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := corsoCmd.ExecuteContext(ctx); err != nil {
		logger.Ctx(ctx).
			With("err", err).
			Errorw("cli execution", clues.InErr(err).Slice()...)
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
