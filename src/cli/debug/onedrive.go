package debug

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addOneDriveCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case metadataFilesCommand:
		c, fs = utils.AddCommand(cmd, oneDriveMetadataFilesCmd(), utils.MarkDebugCommand())

		c.Use = c.Use + " " + oneDriveServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
	}

	return c
}

const (
	oneDriveServiceCommand          = "onedrive"
	oneDriveServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	oneDriveServiceCommandDebugExamples = `# Display file contents for backup 1234abcd
	corso debug metadata-files onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd`
)

// `corso debug metadata-files onedrive [<flag>...] <destination>`
func oneDriveMetadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Display onedrive metadata file content",
		RunE:    metadataFilesOneDriveCmd,
		Args:    cobra.NoArgs,
		Example: oneDriveServiceCommandDebugExamples,
	}
}

func metadataFilesOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// opts := utils.MakeOneDriveOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewOneDriveBackup([]string{"unused-placeholder"})

	return runMetadataFiles(ctx, cmd, args, sel.Selector, flags.BackupIDFV, "OneDrive")
}
