package repo

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	. "github.com/alcionai/canario/src/cli/print"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/internal/events"
	"github.com/alcionai/canario/src/pkg/config"
	ctrlRepo "github.com/alcionai/canario/src/pkg/control/repository"
	"github.com/alcionai/canario/src/pkg/repository"
	"github.com/alcionai/canario/src/pkg/storage"
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

	cfg, err := config.ReadCorsoConfig(
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
	retentionOpts := ctrlRepo.Retention{}

	storageCfg, err := cfg.Storage.ToFilesystemConfig()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving filesystem configuration"))
	}

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.New(
		ctx,
		cfg.Account,
		cfg.Storage,
		opt,
		repository.NewRepoID)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to construct the repository controller"))
	}

	ric := repository.InitConfig{RetentionOpts: retentionOpts}

	if err = r.Initialize(ctx, ric); err != nil {
		return Only(ctx, clues.Stack(ErrInitializingRepo, err))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Initialized a repository at path %s", storageCfg.Path)

	err = config.WriteRepoConfig(
		ctx,
		storageCfg,
		m365,
		opt.Repo,
		r.GetID())
	if err != nil {
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

	cfg, err := config.ReadCorsoConfig(
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

	storageCfg, err := cfg.Storage.ToFilesystemConfig()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving filesystem configuration"))
	}

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

	if err := r.Connect(ctx, repository.ConnConfig{}); err != nil {
		return Only(ctx, clues.Stack(ErrConnectingRepo, err))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Connected to repository at path %s", storageCfg.Path)

	err = config.WriteRepoConfig(
		ctx,
		storageCfg,
		m365,
		opts.Repo,
		r.GetID())
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}
