package debug

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case metadataFilesCommand:
		c, _ = utils.AddCommand(cmd, exchangeMetadataFilesCmd(), utils.MarkDebugCommand())
		c.Use = c.Use + " " + exchangeServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
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

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewExchangeBackup([]string{"unused-placeholder"})
	sel.Include(sel.AllData())

	return genericMetadataFiles(
		ctx,
		cmd,
		args,
		sel.Selector,
		flags.BackupIDFV)
}
