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

	case updateCommand:
		update := filesystemUpdateCmd()
		flags.AddCorsoUpdatePassphraseFlags(update)
		c, _ = utils.AddCommand(cmd, update)
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
	retentionOpts := ctrlRepo.Retention{}

	// SendStartCorsoEvent uses distict ID as tenant ID because repoID is still not generated
	utils.SendStartCorsoEvent(
		ctx,
		cfg.Storage,
		cfg.Account.ID(),
		map[string]any{"command": "init repo"},
		cfg.Account.ID(),
		opt)

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
		if flags.SucceedIfExistsFV && errors.Is(err, repository.ErrorRepoAlreadyExists) {
			return nil
		}

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

// ---------------------------------------------------------------------------------------------------------
// Update
// ---------------------------------------------------------------------------------------------------------

// `corso repo update filesystem [<flag>...]`
func filesystemUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     fsProviderCommand,
		Short:   "Update to a filesystem repository",
		Long:    `Update to an existing repository on local or network storage.`,
		RunE:    updateFilesystemCmd,
		Args:    cobra.NoArgs,
		Example: fsProviderCmdConnectExamples,
	}
}

// updates to an existing filesystem repo.
func updateFilesystemCmd(cmd *cobra.Command, args []string) error {
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

	opts := utils.ControlWithConfig(cfg)

	r, err := repository.New(
		ctx,
		cfg.Account,
		cfg.Storage,
		opts,
		cfg.RepoID)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to create a repository controller"))
	}

	if err := r.UpdatePassword(ctx, flags.UpdateCorsoPhasephraseFV); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to update s3"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Updated repo password.")

	return nil
}
