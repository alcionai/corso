package repo

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/events"
	ctrlRepo "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	fsProviderCommand      = "filesystem"
	fsProviderCmdUseSuffix = "--path <path>"
)

const (
	fsProviderCmdInitExamples = `# Create a new Corso repository on local or network attached storage
corso repo init filesystem --path /tmp/corso-repo`

	fsProviderCmdConnectExamples = `# Connect to a Corso repository on local or network attached storage
corso repo connect filesystem --path /tmp/corso-repo`
)

func addFilesystemCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case initCommand:
		init := filesystemInitCmd()
		c, _ = utils.AddCommand(cmd, init)

	case connectCommand:
		c, _ = utils.AddCommand(cmd, filesystemConnectCmd())
	}

	c.Use = c.Use + " " + fsProviderCmdUseSuffix
	c.SetUsageTemplate(cmd.UsageTemplate())

	flags.AddFilesystemFlags(c)

	return c
}

// `corso repo init filesystem [<flag>...]`
func filesystemInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:     fsProviderCommand,
		Short:   "Initialize a repository on local or network storage.",
		Long:    `Bootstraps a new repository on local or network storage and connects it to your m365 account.`,
		RunE:    initFilesystemCmd,
		Args:    cobra.NoArgs,
		Example: fsProviderCmdInitExamples,
	}
}

// initializes a filesystem repo.
func initFilesystemCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	overrides := flags.FilesystemFlagOverrides(cmd)

	// TODO(pandeyabs): Move filepath conversion to FilesystemConfig scope.
	abs, err := utils.MakeAbsoluteFilePath(overrides[flags.FilesystemPathFN])
	if err != nil {
		return Only(ctx, clues.Wrap(err, "getting absolute repo path"))
	}

	overrides[flags.FilesystemPathFN] = abs

	cfg, err := config.GetConfigRepoDetails(
		ctx,
		storage.ProviderFilesystem,
		true,
		false,
		flags.FilesystemFlagOverrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	opt := utils.ControlWithConfig(cfg)
	// Retention is not supported for filesystem repos.
	retention := ctrlRepo.Retention{}

	// SendStartCorsoEvent uses distict ID as tenant ID because repoID is still not generated
	utils.SendStartCorsoEvent(
		ctx,
		cfg.Storage,
		cfg.Account.ID(),
		map[string]any{"command": "init repo"},
		cfg.Account.ID(),
		opt)

	sc, err := cfg.Storage.StorageConfig()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving filesystem configuration"))
	}

	storageCfg := sc.(*storage.FilesystemConfig)

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.New(
		ctx,
		cfg.Account,
		cfg.Storage,
		opt,
		"")
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to construct the repository controller"))
	}

	if err = r.Initialize(ctx, retention); err != nil {
		if flags.SucceedIfExistsFV && errors.Is(err, repository.ErrorRepoAlreadyExists) {
			return nil
		}

		return Only(ctx, clues.Wrap(err, "Failed to initialize a new filesystem repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Initialized a repository at path %s", storageCfg.Path)

	if err = config.WriteRepoConfig(ctx, sc, m365, opt.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------
// Connect
// ---------------------------------------------------------------------------------------------------------

// `corso repo connect filesystem [<flag>...]`
func filesystemConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:     fsProviderCommand,
		Short:   "Connect to a repository on local or network storage.",
		Long:    `Ensures a connection to an existing repository on local or network storage.`,
		RunE:    connectFilesystemCmd,
		Args:    cobra.NoArgs,
		Example: fsProviderCmdConnectExamples,
	}
}

// connects to an existing filesystem repo.
func connectFilesystemCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	overrides := flags.FilesystemFlagOverrides(cmd)

	// TODO(pandeyabs): Move filepath conversion to FilesystemConfig scope.
	abs, err := utils.MakeAbsoluteFilePath(overrides[flags.FilesystemPathFN])
	if err != nil {
		return Only(ctx, clues.Wrap(err, "getting absolute repo path"))
	}

	overrides[flags.FilesystemPathFN] = abs

	cfg, err := config.GetConfigRepoDetails(
		ctx,
		storage.ProviderFilesystem,
		true,
		true,
		overrides)
	if err != nil {
		return Only(ctx, err)
	}

	repoID := cfg.RepoID
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	sc, err := cfg.Storage.StorageConfig()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving filesystem configuration"))
	}

	storageCfg := sc.(*storage.FilesystemConfig)

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	opts := utils.ControlWithConfig(cfg)

	r, err := repository.New(
		ctx,
		cfg.Account,
		cfg.Storage,
		opts,
		repoID)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to create a repository controller"))
	}

	if err := r.Connect(ctx); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to connect to the filesystem repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Connected to repository at path %s", storageCfg.Path)

	if err = config.WriteRepoConfig(ctx, sc, m365, opts.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}
