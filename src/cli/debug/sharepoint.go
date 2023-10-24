package debug

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case metadataFilesCommand:
		c, fs = utils.AddCommand(cmd, sharePointMetadataFilesCmd(), utils.MarkDebugCommand())

		c.Use = c.Use + " " + sharePointServiceCommandUseSuffix

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

	// opts := utils.MakeSharePointOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewSharePointBackup([]string{"me"})

	return runMetadataFiles(ctx, cmd, args, sel.Selector, flags.BackupIDFV, "SharePoint")
}
