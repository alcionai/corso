package repo

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	initCommand        = "init"
	connectCommand     = "connect"
	maintenanceCommand = "maintenance"
)

var repoCommands = []func(cmd *cobra.Command) *cobra.Command{
	addS3Commands,
}

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	var (
		// Get new instances so that setting the context during tests works
		// properly.
		repoCmd        = repoCmd()
		initCmd        = initCmd()
		connectCmd     = connectCmd()
		maintenanceCmd = maintenanceCmd()
	)

	cmd.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)
	repoCmd.AddCommand(maintenanceCmd)

	flags.AddMaintenanceModeFlag(maintenanceCmd)
	flags.AddForceMaintenanceFlag(maintenanceCmd)
	flags.AddMaintenanceUserFlag(maintenanceCmd)
	flags.AddMaintenanceHostnameFlag(maintenanceCmd)

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
		Use:   maintenanceCommand,
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

	// Change this to override too?
	r, _, err := utils.AccountConnectAndWriteRepoConfig(
		ctx, path.UnknownService, storage.ProviderS3, S3Overrides(cmd))
	if err != nil {
		return print.Only(ctx, err)
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
		return print.Only(ctx, err)
	}

	err = m.Run(ctx)
	if err != nil {
		return print.Only(ctx, err)
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

func GetStorageOverrides(
	ctx context.Context,
	cmd *cobra.Command,
	storageProvider string,
) (map[string]string, error) {
	overrides := map[string]string{}

	switch storageProvider {
	case storage.ProviderS3.String():
		overrides = S3Overrides(cmd)
	}

	return overrides, nil
}

func GetStorageProviderAndOverrides(
	ctx context.Context,
	cmd *cobra.Command,
) (storage.ProviderType, map[string]string, error) {
	provider, err := config.GetStorageProviderFromConfigFile(ctx)
	if err != nil {
		return provider, nil, clues.Stack(err)
	}

	overrides := map[string]string{}

	switch provider {
	case storage.ProviderS3:
		overrides = S3Overrides(cmd)
	}

	return provider, overrides, nil
}
