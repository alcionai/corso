package backup

import (
	"github.com/spf13/cobra"
)

var backupCommands = []func(parent *cobra.Command) *cobra.Command{
	addExchangeCommands,
}

// AddCommands attaches all `corso backup * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(backupCmd)
	backupCmd.AddCommand(createCmd)
	backupCmd.AddCommand(listCmd)
	backupCmd.AddCommand(detailsCmd)

	for _, addBackupTo := range backupCommands {
		addBackupTo(createCmd)
		addBackupTo(listCmd)
		addBackupTo(detailsCmd)
	}
}

// The backup category of commands.
// `corso backup [<subcommand>] [<flag>...]`
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your service data",
	Long:  `Backup the data stored in one of your M365 services.`,
	Run:   handleBackupCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso backup`.
// Produces the same output as `corso backup --help`.
func handleBackupCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The backup create subcommand.
// `corso backup create <service> [<flag>...]`
var createCommand = "create"
var createCmd = &cobra.Command{
	Use:   createCommand,
	Short: "Backup an M365 Service",
	Run:   handleCreateCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso backup create`.
// Produces the same output as `corso backup create --help`.
func handleCreateCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The backup list subcommand.
// `corso backup list <service> [<flag>...]`
var listCommand = "list"
var listCmd = &cobra.Command{
	Use:   listCommand,
	Short: "List the history of restore points for a service",
	Run:   handleListCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso backup list`.
// Produces the same output as `corso backup list --help`.
func handleListCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The backup details subcommand.
// `corso backup list <service> [<flag>...]`
var detailsCommand = "details"
var detailsCmd = &cobra.Command{
	Use:   detailsCommand,
	Short: "Shows the details of a restore point for a service",
	Run:   handleDetailsCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso backup list`.
// Produces the same output as `corso backup list --help`.
func handleDetailsCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
