package export

import (
	"github.com/spf13/cobra"
)

var exportCommands = []func(cmd *cobra.Command) *cobra.Command{
	addOneDriveCommands,
	// addExchangeCommands,
	// addSharePointCommands,
}

// AddCommands attaches all `corso export * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	exportC := exportCmd()
	cmd.AddCommand(exportC)

	for _, addExportTo := range exportCommands {
		addExportTo(exportC)
	}
}

const exportCommand = "export"

// The export category of commands.
// `corso export [<subcommand>] [<flag>...]`
func exportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exportCommand,
		Short: "Export your service data",
		Long:  `Export the data stored in one of your M365 services.`,
		RunE:  handleExportCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso export`.
// Produces the same output as `corso export --help`.
func handleExportCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
