package repo

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var repoCommands = []func(parent *cobra.Command) *cobra.Command{
	addS3Commands,
}

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
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
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage your repositories.",
	Long:  `Initialize, configure, and connect to your account backup repositories.`,
	Run:   handleRepoCmd,
	Args:  cobra.NoArgs,
}

// Handler for flat calls to `corso repo`.
// Produces the same output as `corso repo --help`.
func handleRepoCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// The repo init subcommand.
// `corso repo init <repository> [<flag>...]`
var initCommand = "init"
var initCmd = &cobra.Command{
	Use:   initCommand,
	Short: "Initialize a repository.",
	Long:  `Create a new repository to store your backups.`,
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
	Short: "Connect to a repository.",
	Long:  `Connect to an existing repository.`,
	Run:   handleConnectCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso repo connect`.
func handleConnectCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// aggregates m365 details from flag and env_var values.
type m365Vars struct {
	clientID     string
	clientSecret string
	tenantID     string
}

// helper for aggregating m365 connection details.
func getM365Vars() m365Vars {
	return m365Vars{
		clientID:     os.Getenv("O365_CLIENT_ID"),
		clientSecret: os.Getenv("O356_SECRET"),
		tenantID:     "todo:tenantID",
	}
}

// validates the existence of the properties in the map.
// expects a map[propName]propVal.
func requireProps(props map[string]string) error {
	for name, val := range props {
		if len(val) == 0 {
			return errors.New(name + " is required to perform this command")
		}
	}
	return nil
}
