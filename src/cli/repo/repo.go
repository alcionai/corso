package repo

import (
	"github.com/spf13/cobra"
)

const (
	initCommand    = "init"
	connectCommand = "connect"
)

var repoCommands = []func(parent *cobra.Command) *cobra.Command{
	addS3Commands,
}

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	var (
		// Get new instances so that setting the context during tests works
		// properly.
		repoCmd    = getRepoCmd()
		initCmd    = getInitCmd()
		connectCmd = getConnectCmd()
	)

	parent.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)

	for _, addRepoTo := range repoCommands {
		addRepoTo(initCmd)
		addRepoTo(connectCmd)
	}
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
func getRepoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "repo",
		Short: "Manage your repositories",
		Long:  `Initialize, configure, and connect to your account backup repositories.`,
		RunE:  handleRepoCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso repo`.
// Produces the same output as `corso repo --help`.
func handleRepoCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The repo init subcommand.
// `corso repo init <repository> [<flag>...]`
func getInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   initCommand,
		Short: "Initialize a repository.",
		Long:  `Create a new repository to store your backups.`,
		RunE:  handleInitCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso repo init`.
func handleInitCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The repo connect subcommand.
// `corso repo connect <repository> [<flag>...]`
func getConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   connectCommand,
		Short: "Connect to a repository.",
		Long:  `Connect to an existing repository.`,
		RunE:  handleConnectCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso repo connect`.
func handleConnectCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
