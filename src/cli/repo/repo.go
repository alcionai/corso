package repo

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/flags"
)

// All of the supported repository providers.
var repoProviders = []struct {
	cmd *cobra.Command
	fs  []flags.CliFlag
}{
	{providerS3Cmd, providerS3Flags},
}

var (
	m365Tenant       string
	m365ClientID     string
	m365ClientSecret string

	m365AccountFlags = []flags.CliFlag{
		{
			Name:        "m365-tenant",
			Description: "Your m365 account ID.",
			VarType:     flags.StringType,
			Var:         &m365Tenant,
		},
		{
			Name:        "m365-client-id",
			Description: "Your m365 application client id.",
			VarType:     flags.StringType,
			Var:         &m365ClientID,
		},
		{
			Name:        "m365-client-id",
			Description: "Your m365 application client id.",
			VarType:     flags.StringType,
			Var:         &m365ClientID,
		},
	}
)

// initialize all `corso repo * *` commands.
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)

	for _, cf := range repoProviders {
		addPerProviderCommands(initCmd, cf.cmd, cf.fs)
		addPerProviderCommands(connectCmd, cf.cmd, cf.fs)
	}
}

// initialze the per-provider commands.
func addPerProviderCommands(parent, child *cobra.Command, fs []flags.CliFlag) {
	c := &cobra.Command{}
	*c = *child

	switch parent.Use {
	case initCommand:
		c.Run = initS3Cmd
		m365AccountFlags[0].Required = true
		m365AccountFlags[1].Required = true
	case connectCommand:
		c.Run = connectS3Cmd
		m365AccountFlags[0].Required = false
		m365AccountFlags[1].Required = false
	}

	parent.AddCommand(c)
	flags.AddAll(providerS3Flags, c)
	flags.AddAll(fs, c)
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Initialize and connect to your repository.",
	Long: `TODO: wordsmithing something incredible about Corso
repositories and what you can do with them
(ie: back up all the data).`,
	Run:  handleRepoCmd,
	Args: cobra.NoArgs,
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
