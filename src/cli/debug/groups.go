package debug

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/pkg/selectors"
)

// called by debug.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case metadataFilesCommand:
		c, _ = utils.AddCommand(cmd, groupsMetadataFilesCmd(), utils.MarkDebugCommand())

		c.Use = c.Use + " " + groupsServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
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

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	sel := selectors.NewGroupsBackup([]string{"unused-placeholder"})
	sel.Include(sel.AllData())

	return genericMetadataFiles(
		ctx,
		cmd,
		args,
		sel.Selector,
		flags.BackupIDFV)
}
