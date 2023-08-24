package repo

import (
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

// // filesystem flags
// const (
// 	filesystemPathFN = "path"
// )

const (
	fsProviderCommand          = "filesystem"
	fsProviderCommandUseSuffix = "--path <path>"
)

const (
	fsProviderCommandInitExamples = `# Create a new Corso repo in local or 
	network attached storage
corso repo init filesystem --path /mnt/corso-repo

# Create a new Corso repo in local or network attached storage using a prefix
corso repo init filesystem --path /mnt/corso-repo --prefix my-prefix`

	fsProviderCommandConnectExamples = `# Connect to a Corso repo in local or
	network attached storage"
corso repo connect filesystem --path /mnt/corso-repo

# Connect to a Corso repo in local or network attached storage using a prefix
corso repo connect filesystem --path /mnt/corso-repo --prefix my-prefix`
)

// called by repo.go to map subcommands to provider-specific handling.
func addFsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case initCommand:
		init := fsInitCmd()
		flags.AddRetentionConfigFlags(init)
		c, fs = utils.AddCommand(cmd, init)

	case connectCommand:
		c, fs = utils.AddCommand(cmd, fsConnectCmd())
	}

	c.Use = c.Use + " " + fsProviderCommandUseSuffix
	c.SetUsageTemplate(cmd.UsageTemplate())

	flags.AddAzureCredsFlags(c)
	flags.AddCorsoPassphaseFlags(c)
	flags.AddFilesystemPathFlag(c, true)

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic and more frequently used flags take precedence.
	fs.StringVar(&prefix, prefixFN, "", "Repo prefix.")

	// In general, we don't want to expose this flag to users and have them mistake it
	// for a broad-scale idempotency solution.  We can un-hide it later the need arises.
	fs.BoolVar(&succeedIfExists, "succeed-if-exists", false, "Exit with success if the repo has already been initialized.")
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))

	return c
}

// `corso repo init filesystem [<flag>...]`
func fsInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   fsProviderCommand,
		Short: "Initialize a repository in local or network attached storage",
		Long: `Bootstraps a new local or network attached storage repository
		and connects it to your m365 account.`,
		RunE:    initFsCmd,
		Args:    cobra.NoArgs,
		Example: fsProviderCommandInitExamples,
	}
}

// initializes a filesystem repo.
func initFsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// TODO(pandeyabs): Modify this func to accept non s3.
	cfg, err := config.GetConfigRepoDetails(ctx, true, false, nil)
	if err != nil {
		return Only(ctx, err)
	}

	opt := utils.ControlWithConfig(cfg)

	retentionOpts, err := utils.MakeRetentionOpts(cmd)
	if err != nil {
		return Only(ctx, err)
	}

	// SendStartCorsoEvent uses distict ID as tenant ID because repoID is still not generated
	utils.SendStartCorsoEvent(
		ctx,
		cfg.Storage,
		cfg.Account.ID(),
		map[string]any{"command": "init repo"},
		cfg.Account.ID(),
		opt)

	// TODO(pandeyabs): Need prefix for filesystem repo.
	// s3Cfg, err := cfg.Storage.S3Config()
	// if err != nil {
	// 	return Only(ctx, clues.Wrap(err, "Retrieving s3 configuration"))
	// }

	_, err = cfg.Storage.FsConfig()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving filesystem configuration"))
	}

	cfg.Storage.Provider = storage.ProviderFS

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.Initialize(
		ctx,
		cfg.Account,
		cfg.Storage,
		opt,
		retentionOpts)
	if err != nil {
		if succeedIfExists && errors.Is(err, repository.ErrorRepoAlreadyExists) {
			return nil
		}

		return Only(ctx, clues.Wrap(err, "Failed to initialize a new filesystem repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Initialized a repository at path %s", flags.FilesystemPathFV)

	if err = config.WriteRepoConfig(ctx, storage.S3Config{}, m365, opt.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------
// Connect
// ---------------------------------------------------------------------------------------------------------

// `corso repo connect s3 [<flag>...]`
func fsConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:     fsProviderCommand,
		Short:   "Connect to a local repository",
		Long:    `Ensures a connection to an existing local repository.`,
		RunE:    connectFsCmd,
		Args:    cobra.NoArgs,
		Example: fsProviderCommandConnectExamples,
	}
}

// connects to an existing s3 repo.
func connectFsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// s3 values from flags
	s3Override := S3Overrides(cmd)

	cfg, err := config.GetConfigRepoDetails(ctx, true, true, s3Override)
	if err != nil {
		return Only(ctx, err)
	}

	repoID := cfg.RepoID
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	s3Cfg, err := cfg.Storage.S3Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving s3 configuration"))
	}

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	if strings.HasPrefix(s3Cfg.Endpoint, "http://") || strings.HasPrefix(s3Cfg.Endpoint, "https://") {
		invalidEndpointErr := "endpoint doesn't support specifying protocol. " +
			"pass --disable-tls flag to use http:// instead of default https://"

		return Only(ctx, clues.New(invalidEndpointErr))
	}

	opts := utils.ControlWithConfig(cfg)

	r, err := repository.ConnectAndSendConnectEvent(
		ctx,
		cfg.Account,
		cfg.Storage,
		repoID,
		opts)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to connect to the S3 repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Connected to repository at path %s", flags.FilesystemPathFV)

	if err = config.WriteRepoConfig(ctx, s3Cfg, m365, opts.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}
