package debug

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case metadataFilesCommand:
		c, fs = utils.AddCommand(cmd, groupsMetadataFilesCmd(), utils.MarkDebugCommand())

		c.Use = c.Use + " " + groupsServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
	}

	return c
}

// TODO: correct examples
const (
	groupsServiceCommand          = "groups"
	groupsServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	groupsServiceCommandDebugExamples = `# Display file contents for backup 1234abcd
	corso debug metadata-files groups --backup 1234abcd-12ab-cd34-56de-1234abcd`
)

// `corso debug metadata-files groups [<flag>...] <destination>`
func groupsMetadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     groupsServiceCommand,
		Short:   "Display groups metadata file content",
		RunE:    metadataFilesGroupsCmd,
		Args:    cobra.NoArgs,
		Example: groupsServiceCommandDebugExamples,
	}
}

func metadataFilesGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// opts := utils.MakeGroupsOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewGroupsBackup([]string{"me"})

	return runMetadataFiles(ctx, cmd, args, sel.Selector, flags.BackupIDFV, "Groups")
}
