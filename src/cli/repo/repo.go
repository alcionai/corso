package repo

import (
	"github.com/spf13/cobra"
)

var repoProviderInits = []func(cmd *cobra.Command){
	initRepoProviderS3,
}

// initialize all `corso repo * *` commands.
func InitCommands(cmd *cobra.Command) {
	cmd.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)

	// init per-provider subcommands and flags
	for _, initProviderCmd := range repoProviderInits {
		initProviderCmd(initCmd)
		initProviderCmd(connectCmd)
	}
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Init and connect to your repository.",
	Long: `TODO: wordsmithing something incredible about Corso
repositories and what you can do with them
(ie: back up all the data).`,
	Run:  handleCmd,
	Args: cobra.NoArgs,
}

// Handler for flat calls to `corso repo`.
// Produces the same output as `corso repo --help`.
func handleCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The repo init subcommand.
// `corso repo init <repository> [<flag>...]`
var initCommand = "init"
var initCmd = &cobra.Command{
	Use:   initCommand,
	Short: "Initialize a data storage repository.",
	Long:  `TODO: exhaustive details on initializing the repo.`,
	Run:   handleInitCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso repo init`.
func handleInitCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The repo connect subcommand.
// `corso repo connect <repository> [<flag>...]`
var connectCommand = "connect"
var connectCmd = &cobra.Command{
	Use:   connectCommand,
	Short: "Connect to an existing data storage repository.",
	Long:  `TODO: exhaustive details on connecting to the repo.`,
	Run:   handleInitCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso repo connect`.
func handleConnectCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
