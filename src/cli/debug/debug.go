package debug

import (
	"context"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var subCommandFuncs = []func() *cobra.Command{
	metadataFilesCmd,
}

var debugCommands = []func(cmd *cobra.Command) *cobra.Command{
	addOneDriveCommands,
	addSharePointCommands,
	addGroupsCommands,
	addExchangeCommands,
}

// AddCommands attaches all `corso debug * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	debugC, _ := utils.AddCommand(cmd, debugCmd(), utils.MarkDebugCommand())

	for _, sc := range subCommandFuncs {
		subCommand := sc()
		utils.AddCommand(debugC, subCommand, utils.MarkDebugCommand())

		for _, addBackupTo := range debugCommands {
			addBackupTo(subCommand)
		}
	}
}

// ---------------------------------------------------------------------------
// Commands
// ---------------------------------------------------------------------------

const debugCommand = "debug"

// The debug category of commands.
// `corso debug [<subcommand>] [<flag>...]`
func debugCmd() *cobra.Command {
	return &cobra.Command{
		Use:   debugCommand,
		Short: "debug your service data",
		Long:  `debug the data stored in one of your M365 services.`,
		RunE:  handledebugCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso debug`.
// Produces the same output as `corso debug --help`.
func handledebugCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup metadataFiles subcommand.
// `corso backup metadata-files <service> [<flag>...]`
var metadataFilesCommand = "metadata-files"

func metadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   metadataFilesCommand,
		Short: "Display all the metadata file contents stored by the service",
		RunE:  handleMetadataFilesCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso backup metadata-files`.
// Produces the same output as `corso backup metadata-files --help`.
func handleMetadataFilesCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ---------------------------------------------------------------------------
// runners
// ---------------------------------------------------------------------------

func runMetadataFiles(
	ctx context.Context,
	cmd *cobra.Command,
	args []string,
	sel selectors.Selector,
	backupID, serviceName string,
) error {
	r, _, err := utils.GetAccountAndConnect(ctx, cmd, sel.PathService())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// TODO: read and print out all metadata files in the backup

	return nil
}
