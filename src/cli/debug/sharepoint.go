package debug

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case metadataFilesCommand:
		c, _ = utils.AddCommand(cmd, sharePointMetadataFilesCmd(), utils.MarkDebugCommand())
		c.Use = c.Use + " " + sharePointServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
	}

	return c
}

const (
	sharePointServiceCommand          = "sharepoint"
	sharePointServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	sharePointServiceCommandDebugExamples = `# Display file contents for backup 1234abcd
	corso debug metadata-files sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd`
)

// `corso debug metadata-files sharepoint [<flag>...] <destination>`
func sharePointMetadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     sharePointServiceCommand,
		Short:   "Display sharepoint metadata file content",
		RunE:    metadataFilesSharePointCmd,
		Args:    cobra.NoArgs,
		Example: sharePointServiceCommandDebugExamples,
	}
}

func metadataFilesSharePointCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewSharePointBackup([]string{"unused-placeholder"})
	sel.Include(sel.LibraryFolders(selectors.Any()))

	return genericMetadataFiles(
		ctx,
		cmd,
		args,
		sel.Selector,
		flags.BackupIDFV)
}
