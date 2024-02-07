package debug

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	. "github.com/alcionai/canario/src/cli/print"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/selectors"
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

		for _, addTo := range debugCommands {
			servCmd := addTo(subCommand)
			flags.AddAllProviderFlags(servCmd)
			flags.AddAllStorageFlags(servCmd)
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
		Short: "debugging & troubleshooting utilities",
		Long:  `debug the data stored in corso.`,
		RunE:  handledebugCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso debug`.
// Produces the same output as `corso debug --help`.
func handledebugCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The debug metadataFiles subcommand.
// `corso debug metadata-files <service> [<flag>...]`
var metadataFilesCommand = "metadata-files"

func metadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   metadataFilesCommand,
		Short: "display all the metadata file contents stored by the service",
		RunE:  handleMetadataFilesCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso debug metadata-files`.
// Produces the same output as `corso debug metadata-files --help`.
func handleMetadataFilesCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ---------------------------------------------------------------------------
// runners
// ---------------------------------------------------------------------------

func genericMetadataFiles(
	ctx context.Context,
	cmd *cobra.Command,
	args []string,
	sel selectors.Selector,
	backupID string,
) error {
	ctx = clues.Add(ctx, "backup_id", backupID)

	r, _, err := utils.GetAccountAndConnect(ctx, cmd, sel.PathService())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// read metadata
	files, err := r.GetBackupMetadata(ctx, sel, backupID, fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "retrieving metadata files"))
	}

	for _, file := range files {
		Infof(ctx, "\n------------------------------")
		Info(ctx, file.Name)
		Info(ctx, file.Path)
		Pretty(ctx, file.Data)
	}

	return nil
}
