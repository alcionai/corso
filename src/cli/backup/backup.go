package backup

import (
	"github.com/spf13/cobra"
)

// The backup category of commands.
// `corso backup [<subcommand>] [<flag>...]`
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Back up your application data.",
	Long:  `Back up the data stored in one of your M365 applications.`,
	Run:   handleBackupCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso backup`.
// Produces the same output as `corso backup --help`.
func handleBackupCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
