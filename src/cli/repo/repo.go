package repo

import (
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/path"
	repo "github.com/alcionai/corso/src/pkg/repository"
)

const (
	initCommand             = "init"
	connectCommand          = "connect"
	updatePassphraseCommand = "update-passphrase"
	MaintenanceCommand      = "maintenance"
)

const (
	providerCommandUpdatePhasephraseExamples = `# Update the Corso repository passphrase"
corso repo update-passphrase --new-passphrase 'newpass'`
)

var (
	ErrConnectingRepo   = clues.New("connecting repository")
	ErrInitializingRepo = clues.New("initializing repository")
)

var repoCommands = []func(cmd *cobra.Command) *cobra.Command{
	addS3Commands,
	addFilesystemCommands,
}

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	var (
		// Get new instances so that setting the context during tests works
		// properly.
		repoCmd             = repoCmd()
		initCmd             = initCmd()
		connectCmd          = connectCmd()
		maintenanceCmd      = maintenanceCmd()
		updatePassphraseCmd = updatePassphraseCmd()
	)

	cmd.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)
	repoCmd.AddCommand(maintenanceCmd)
	repoCmd.AddCommand(updatePassphraseCmd)

	flags.AddMaintenanceModeFlag(maintenanceCmd)
	flags.AddForceMaintenanceFlag(maintenanceCmd)
	flags.AddMaintenanceUserFlag(maintenanceCmd)
	flags.AddMaintenanceHostnameFlag(maintenanceCmd)

	flags.AddUpdatePassphraseFlags(updatePassphraseCmd, true)

	for _, addRepoTo := range repoCommands {
		addRepoTo(initCmd)
		addRepoTo(connectCmd)
	}
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
func repoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "repo",
		Short: "Manage your repositories",
		Long:  `Initialize, configure, connect and update to your account backup repositories`,
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
func initCmd() *cobra.Command {
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
func connectCmd() *cobra.Command {
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

func maintenanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   MaintenanceCommand,
		Short: "Run maintenance on an existing repository",
		Long:  `Run maintenance on an existing repository to optimize performance and storage use`,
		RunE:  handleMaintenanceCmd,
		Args:  cobra.NoArgs,
	}
}

func handleMaintenanceCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	t, err := getMaintenanceType(flags.MaintenanceModeFV)
	if err != nil {
		return err
	}

	r, _, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		// Need to give it a valid service so it won't error out on us even though
		// we don't need the graph client.
		path.OneDriveService)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	m, err := r.NewMaintenance(
		ctx,
		repository.Maintenance{
			Type:   t,
			Safety: repository.FullMaintenanceSafety,
			Force:  flags.ForceMaintenanceFV,
		})
	if err != nil {
		return Only(ctx, err)
	}

	err = m.Run(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	return nil
}

func getMaintenanceType(t string) (repository.MaintenanceType, error) {
	res, ok := repository.StringToMaintenanceType[t]
	if !ok {
		modes := maps.Keys(repository.StringToMaintenanceType)
		allButLast := []string{}

		for i := 0; i < len(modes)-1; i++ {
			allButLast = append(allButLast, string(modes[i]))
		}

		valuesStr := strings.Join(allButLast, ", ") + " or " + string(modes[len(modes)-1])

		return res, clues.New(t + " is an unrecognized maintenance mode; must be one of " + valuesStr)
	}

	return res, nil
}

// The repo update subcommand.
// `corso repo update-passphrase [<flag>...]`
func updatePassphraseCmd() *cobra.Command {
	return &cobra.Command{
		Use:     updatePassphraseCommand,
		Short:   "Update the repository passphrase",
		Long:    `Update the repository passphrase`,
		RunE:    handleUpdateCmd,
		Args:    cobra.NoArgs,
		Example: providerCommandUpdatePhasephraseExamples,
	}
}

// Handler for calls to `corso repo update-password`.
func handleUpdateCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Need to give it a valid service so it won't error out on us even though
	// we don't need the graph client.
	repos, rdao, err := utils.GetAccountAndConnect(ctx, cmd, path.OneDriveService)
	if err != nil {
		return Only(ctx, err)
	}

	opts := rdao.Opts

	defer utils.CloseRepo(ctx, repos)

	repoID := repos.GetID()
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	r, err := repo.New(
		ctx,
		rdao.Repo.Account,
		rdao.Repo.Storage,
		opts,
		repoID)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to create a repository controller"))
	}

	if err := r.UpdatePassword(ctx, flags.NewPhasephraseFV); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to update s3"))
	}

	Infof(ctx, "Updated repo password.")

	return nil
}
