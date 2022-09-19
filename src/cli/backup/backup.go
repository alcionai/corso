package backup

import (
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	createCmd,
	listCmd,
	detailsCmd,
	deleteCmd,
}

var serviceCommands = []func(parent *cobra.Command) *cobra.Command{
	addExchangeCommands,
	addOneDriveCommands,
}

// AddCommands attaches all `corso backup * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(backupCmd)

	for _, sc := range subCommands {
		backupCmd.AddCommand(sc)
	}

	for _, addBackupTo := range serviceCommands {
		for _, sc := range subCommands {
			addBackupTo(sc)
		}
	}
}

// The backup category of commands.
// `corso backup [<subcommand>] [<flag>...]`
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your service data",
	Long:  `Backup the data stored in one of your M365 services.`,
	RunE:  handleBackupCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso backup`.
// Produces the same output as `corso backup --help`.
func handleBackupCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup create subcommand.
// `corso backup create <service> [<flag>...]`
var (
	createCommand = "create"
	createCmd     = &cobra.Command{
		Use:   createCommand,
		Short: "Backup an M365 Service",
		RunE:  handleCreateCmd,
		Args:  cobra.NoArgs,
	}
)

// Handler for calls to `corso backup create`.
// Produces the same output as `corso backup create --help`.
func handleCreateCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup list subcommand.
// `corso backup list <service> [<flag>...]`
var (
	listCommand = "list"
	listCmd     = &cobra.Command{
		Use:   listCommand,
		Short: "List the history of backups for a service",
		RunE:  handleListCmd,
		Args:  cobra.NoArgs,
	}
)

// Handler for calls to `corso backup list`.
// Produces the same output as `corso backup list --help`.
func handleListCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup details subcommand.
// `corso backup details <service> [<flag>...]`
var (
	detailsCommand = "details"
	detailsCmd     = &cobra.Command{
		Use:   detailsCommand,
		Short: "Shows the details of a backup for a service",
		RunE:  handleDetailsCmd,
		Args:  cobra.NoArgs,
	}
)

// Handler for calls to `corso backup details`.
// Produces the same output as `corso backup details --help`.
func handleDetailsCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup delete subcommand.
// `corso backup delete <service> [<flag>...]`
var (
	deleteCommand = "delete"
	deleteCmd     = &cobra.Command{
		Use:   deleteCommand,
		Short: "Deletes a backup for a service",
		RunE:  handleDeleteCmd,
		Args:  cobra.NoArgs,
	}
)

// Handler for calls to `corso backup delete`.
// Produces the same output as `corso backup delete --help`.
func handleDeleteCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
