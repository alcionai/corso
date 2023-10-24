package debug

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case metadataFilesCommand:
		c, fs = utils.AddCommand(cmd, exchangeMetadataFilesCmd(), utils.MarkDebugCommand())

		c.Use = c.Use + " " + exchangeServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	exchangeServiceCommandDebugExamples = `# Display file contents for backup 1234abcd
corso debug metadata-files exchange --backup 1234abcd-12ab-cd34-56de-1234abcd`
)

// `corso debug metadata-files exchange [<flag>...] <destination>`
func exchangeMetadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     exchangeServiceCommand,
		Short:   "Display exchange metadata file content",
		RunE:    metadataFilesExchangeCmd,
		Args:    cobra.NoArgs,
		Example: exchangeServiceCommandDebugExamples,
	}
}

func metadataFilesExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// opts := utils.MakeExchangeOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewExchangeBackup([]string{"me"})

	return runMetadataFiles(ctx, cmd, args, sel.Selector, flags.BackupIDFV, "Exchange")
}
