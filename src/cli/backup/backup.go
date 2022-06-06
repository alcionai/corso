package backup

import (
	"github.com/spf13/cobra"
)

var backupApplications = []func(parent *cobra.Command) *cobra.Command{
	addExchangeApp,
}

// AddCommands attaches all `corso backup * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(backupCmd)
	backupCmd.AddCommand(createCmd)

	for _, addBackupTo := range backupApplications {
		addBackupTo(createCmd)
	}
}

// The backup category of commands.
// `corso backup [<subcommand>] [<flag>...]`
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your application data.",
	Long:  `Backup the data stored in one of your M365 applications.`,
	Run:   handleBackupCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso backup`.
// Produces the same output as `corso backup --help`.
func handleBackupCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The backup create subcommand.
// `corso backup create <application> [<flag>...]`
var createCommand = "create"
var createCmd = &cobra.Command{
	Use:   createCommand,
	Short: "Create a new backup.",
	Long:  `Create a new backup of your application.`,
	Run:   handleCreateCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso backup create`.
// Produces the same output as `corso backup create --help`.
func handleCreateCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
