package restore

import (
	"github.com/spf13/cobra"
)

var restoreCommands = []func(parent *cobra.Command) *cobra.Command{
	addExchangeCommands,
}

// AddCommands attaches all `corso restore * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(restoreCmd)

	for _, addRestoreTo := range restoreCommands {
		addRestoreTo(restoreCmd)
	}
}

const restoreCommand = "restore"

// The restore category of commands.
// `corso restore [<subcommand>] [<flag>...]`
var restoreCmd = &cobra.Command{
	Use:   restoreCommand,
	Short: "Restore your service data",
	Long:  `Restore the data stored in one of your M365 services.`,
	RunE:  handleRestoreCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso restore`.
// Produces the same output as `corso restore --help`.
func handleRestoreCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
