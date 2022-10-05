package repo

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

// s3 bucket info from flags
var (
	bucket          string
	endpoint        string
	prefix          string
	succeedIfExists bool
)

// called by repo.go to map parent subcommands to provider-specific handling.
func addS3Commands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case initCommand:
		c, fs = utils.AddCommand(parent, s3InitCmd())
	case connectCommand:
		c, fs = utils.AddCommand(parent, s3ConnectCmd())
	}

	c.Use = c.Use + s3ProviderCommandUseSuffix

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic (ex: --all) and more frequently used flags take precedence.
	fs.StringVar(&bucket, "bucket", "", "Name of S3 bucket for repo. (required)")
	cobra.CheckErr(c.MarkFlagRequired("bucket"))
	fs.StringVar(&prefix, "prefix", "", "Repo prefix within bucket.")
	fs.StringVar(&endpoint, "endpoint", "s3.amazonaws.com", "S3 service endpoint.")
	fs.BoolVar(&succeedIfExists, "succeed-if-exists", false, "Exit with success if the repo has already been initialized.")
	// In general, we don't want to expose this flag to users and have them mistake it
	// for a broad-scale idempotency solution.  We can un-hide it later the need arises.
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))

	return c
}

const (
	s3ProviderCommand          = "s3"
	s3ProviderCommandUseSuffix = " --bucket <bucket>"
)

const (
	s3ProviderCommandInitExamples = `# Create a new Corso repo in AWS S3 bucket named "my-bucket"
corso repo init s3 --bucket my-bucket

# Create a new Corso repo in AWS S3 bucket named "my-bucket" using a prefix
corso repo init s3 --bucket my-bucket --prefix my-prefix

# Create a new Corso repo in an S3 compliant storage provider
corso repo init s3 --bucket my-bucket --endpoint https://my-s3-server-endpoint`

	s3ProviderCommandConnectExamples = `# Connect to a Corso repo in AWS S3 bucket named "my-bucket"
corso repo connect s3 --bucket my-bucket

# Connect to a Corso repo in AWS S3 bucket named "my-bucket" using a prefix
corso repo connect s3 --bucket my-bucket --prefix my-prefix

# Connect to a Corso repo in an S3 compliant storage provider
corso repo connect s3 --bucket my-bucket --endpoint https://my-s3-server-endpoint`
)

// ---------------------------------------------------------------------------------------------------------
// Init
// ---------------------------------------------------------------------------------------------------------

// `corso repo init s3 [<flag>...]`
func s3InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:     s3ProviderCommand,
		Short:   "Initialize a S3 repository",
		Long:    `Bootstraps a new S3 repository and connects it to your m356 account.`,
		RunE:    initS3Cmd,
		Args:    cobra.NoArgs,
		Example: s3ProviderCommandInitExamples,
	}
}

// initializes a s3 repo.
func initS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, a, err := config.GetStorageAndAccount(ctx, false, s3Overrides())
	if err != nil {
		return Only(ctx, err)
	}

	s3Cfg, err := s.S3Config()
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Retrieving s3 configuration"))
	}

	m365, err := a.M365Config()
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.Initialize(ctx, a, s, options.Control())
	if err != nil {
		if succeedIfExists && kopia.IsRepoAlreadyExistsError(err) {
			return nil
		}

		return Only(ctx, errors.Wrap(err, "Failed to initialize a new S3 repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Initialized a S3 repository within bucket %s.", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(ctx, s3Cfg, m365); err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------
// Connect
// ---------------------------------------------------------------------------------------------------------

// `corso repo connect s3 [<flag>...]`
func s3ConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:     s3ProviderCommand,
		Short:   "Connect to a S3 repository",
		Long:    `Ensures a connection to an existing S3 repository.`,
		RunE:    connectS3Cmd,
		Args:    cobra.NoArgs,
		Example: s3ProviderCommandConnectExamples,
	}
}

// connects to an existing s3 repo.
func connectS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, a, err := config.GetStorageAndAccount(ctx, true, s3Overrides())
	if err != nil {
		return Only(ctx, err)
	}

	s3Cfg, err := s.S3Config()
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Retrieving s3 configuration"))
	}

	m365, err := a.M365Config()
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.Connect(ctx, a, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to connect to the S3 repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Connected to S3 bucket %s.", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(ctx, s3Cfg, m365); err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

func s3Overrides() map[string]string {
	return map[string]string{
		config.AccountProviderTypeKey: account.ProviderM365.String(),
		config.StorageProviderTypeKey: storage.ProviderS3.String(),
		storage.Bucket:                bucket,
		storage.Endpoint:              endpoint,
		storage.Prefix:                prefix,
	}
}
