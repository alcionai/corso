package repo

import (
	"github.com/spf13/cobra"
)

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)

	for _, addF := range repoProviders {
		addPerProviderCommands(initCmd, addF)
		addPerProviderCommands(connectCmd, addF)
	}
}

// All of the supported repository providers.
var repoProviders = []func(*cobra.Command) *cobra.Command{
	addS3Commands,
}

// m365 account info from flags
var (
	m365Tenant       string
	m365ClientID     string
	m365ClientSecret string
)

// initialze the per-provider commands for each
func addPerProviderCommands(parent *cobra.Command, addCommands func(*cobra.Command) *cobra.Command) {
	c := addCommands(parent)

	// m365 flags are available independent of the provider
	fs := c.Flags()
	fs.StringVar(&m365Tenant, "tenant", "", "Your m365 account ID.")
	fs.StringVar(&m365ClientID, "client-id", "", "Your m365 application client ID.")

	if parent.Use == initUse {
		c.MarkFlagRequired("m365-tenant")
		c.MarkFlagRequired("m365-client-id")
	}
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage your repositories",
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
var initUse = "init"
var initCmd = &cobra.Command{
	Use:   initUse,
	Short: "Initialize a repository",
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
var connectUse = "connect"
var connectCmd = &cobra.Command{
	Use:   connectUse,
	Short: "Connect to a repository",
	Long:  `Connect to an existing repository.`,
	Run:   handleInitCmd,
	Args:  cobra.NoArgs,
}

// Handler for calls to `corso repo connect`.
func handleConnectCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

type repoConnector interface {
	Initialize() error
	Connect() error
}

func initializeRepo(rc repoConnector) error {
	return rc.Initialize()
}

func connectRepo(rc repoConnector) error {
	return rc.Connect()
}
