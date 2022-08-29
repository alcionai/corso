package repo

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
)

// s3 bucket info from flags
var (
	accessKey       string
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

	fs.StringVar(&accessKey, "access-key", "", "Access key ID (replaces the AWS_ACCESS_KEY_ID env variable).")
	fs.StringVar(&bucket, "bucket", "", "Name of the S3 bucket (required).")
	cobra.CheckErr(c.MarkFlagRequired("bucket"))
	fs.StringVar(&endpoint, "endpoint", "s3.amazonaws.com", "Server endpoint for S3 communication.")
	fs.StringVar(&prefix, "prefix", "", "Prefix applied to objects in the bucket.")
	fs.BoolVar(&succeedIfExists, "succeed-if-exists", false, "Exit with success if the repo has already been initialized.")
	// In general, we don't want to expose this flag to users and have them mistake it
	// for a broad-scale idempotency solution.  We can un-hide it later the need arises.
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))

	return c
}

const s3ProviderCommand = "s3"

// ---------------------------------------------------------------------------------------------------------
// Init
// ---------------------------------------------------------------------------------------------------------

// `corso repo init s3 [<flag>...]`
func s3InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   s3ProviderCommand,
		Short: "Initialize a S3 repository",
		Long:  `Bootstraps a new S3 repository and connects it to your m356 account.`,
		RunE:  initS3Cmd,
		Args:  cobra.NoArgs,
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

	r, err := repository.Initialize(ctx, a, s)
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
		Use:   s3ProviderCommand,
		Short: "Connect to a S3 repository",
		Long:  `Ensures a connection to an existing S3 repository.`,
		RunE:  connectS3Cmd,
		Args:  cobra.NoArgs,
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

	r, err := repository.Connect(ctx, a, s)
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
		credentials.AWSAccessKeyID:    accessKey,
		storage.Bucket:                bucket,
		storage.Endpoint:              endpoint,
		storage.Prefix:                prefix,
	}
}
