package repo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"

	kopiarepo "github.com/kopia/kopia/repo" // question for reviewer: do you always wrap kopia packages or is it ok to import this directly here?
)

// s3 bucket info from flags
var (
	accessKey       string
	bucket          string
	endpoint        string
	prefix          string
	successOnExists bool
)

// called by repo.go to map parent subcommands to provider-specific handling.
func addS3Commands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)
	switch parent.Use {
	case initCommand:
		c, fs = utils.AddCommand(parent, s3InitCmd)
	case connectCommand:
		c, fs = utils.AddCommand(parent, s3ConnectCmd)
	}
	fs.StringVar(&accessKey, "access-key", "", "Access key ID (replaces the AWS_ACCESS_KEY_ID env variable).")
	fs.StringVar(&bucket, "bucket", "", "Name of the S3 bucket (required).")
	cobra.CheckErr(c.MarkFlagRequired("bucket"))
	fs.StringVar(&endpoint, "endpoint", "s3.amazonaws.com", "Server endpoint for S3 communication.")
	fs.StringVar(&prefix, "prefix", "", "Prefix applied to objects in the bucket.")
	fs.BoolVar(&successOnExists, "success-on-exists", false, "Do not throw error if repo is already initialized.")
	return c
}

const s3ProviderCommand = "s3"

// `corso repo init s3 [<flag>...]`
var s3InitCmd = &cobra.Command{
	Use:   s3ProviderCommand,
	Short: "Initialize a S3 repository",
	Long:  `Bootstraps a new S3 repository and connects it to your m356 account.`,
	RunE:  initS3Cmd,
	Args:  cobra.NoArgs,
}

// initializes a s3 repo.
func initS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	log := logger.Ctx(ctx)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, a, err := config.GetStorageAndAccount(false, s3Overrides())
	if err != nil {
		return err
	}
	s3Cfg, err := s.S3Config()
	if err != nil {
		return errors.Wrap(err, "Retrieving s3 configuration")
	}

	m365, err := a.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	log.Debugw(
		"Called - "+cmd.CommandPath(),
		"bucket", s3Cfg.Bucket,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0,
		"accessKey", s3Cfg.AccessKey,
		"hasSecretKey", len(s3Cfg.SecretKey) > 0)

	r, err := repository.Initialize(ctx, a, s)
	if err != nil {
		if successOnExists && errors.Is(err, kopiarepo.ErrAlreadyInitialized) {
			return nil
		}
		return errors.Wrap(err, "Failed to initialize a new S3 repository")
	}
	defer utils.CloseRepo(ctx, r)

	fmt.Printf("Initialized a S3 repository within bucket %s.\n", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(s3Cfg, m365); err != nil {
		return errors.Wrap(err, "Failed to write repository configuration")
	}
	return nil
}

// `corso repo connect s3 [<flag>...]`
var s3ConnectCmd = &cobra.Command{
	Use:   s3ProviderCommand,
	Short: "Connect to a S3 repository",
	Long:  `Ensures a connection to an existing S3 repository.`,
	RunE:  connectS3Cmd,
	Args:  cobra.NoArgs,
}

// connects to an existing s3 repo.
func connectS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	log := logger.Ctx(ctx)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, a, err := config.GetStorageAndAccount(true, s3Overrides())
	if err != nil {
		return err
	}
	s3Cfg, err := s.S3Config()
	if err != nil {
		return errors.Wrap(err, "Retrieving s3 configuration")
	}
	m365, err := a.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	log.Debugw(
		"Called - "+cmd.CommandPath(),
		"bucket", s3Cfg.Bucket,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0,
		"accessKey", s3Cfg.AccessKey,
		"hasSecretKey", len(s3Cfg.SecretKey) > 0)

	r, err := repository.Connect(ctx, a, s)
	if err != nil {
		return errors.Wrap(err, "Failed to connect to the S3 repository")
	}
	defer utils.CloseRepo(ctx, r)

	fmt.Printf("Connected to S3 bucket %s.\n", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(s3Cfg, m365); err != nil {
		return errors.Wrap(err, "Failed to write repository configuration")
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
