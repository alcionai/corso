package repo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
)

// s3 bucket info from flags
var (
	accessKey string
	bucket    string
	endpoint  string
	prefix    string
)

// called by repo.go to map parent subcommands to provider-specific handling.
func addS3Commands(parent *cobra.Command) *cobra.Command {
	var c *cobra.Command
	switch parent.Use {
	case initCommand:
		c = s3InitCmd
	case connectCommand:
		c = s3ConnectCmd
	}
	parent.AddCommand(c)
	fs := c.Flags()
	fs.StringVar(&accessKey, "access-key", "", "Access key ID (replaces the AWS_ACCESS_KEY_ID env variable).")
	fs.StringVar(&bucket, "bucket", "", "Name of the S3 bucket (required).")
	c.MarkFlagRequired("bucket")
	fs.StringVar(&endpoint, "endpoint", "s3.amazonaws.com", "Server endpoint for S3 communication.")
	fs.StringVar(&prefix, "prefix", "", "Prefix applied to objects in the bucket.")
	return c
}

// `corso repo init s3 [<flag>...]`
var s3InitCmd = &cobra.Command{
	Use:   "s3",
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

	overrides := map[string]string{
		credentials.AWSAccessKeyID: accessKey,
		storage.Bucket:             bucket,
		storage.Endpoint:           endpoint,
		storage.Prefix:             prefix,
	}
	s, cfgTenantID, err := config.MakeS3Config(false, overrides)
	if err != nil {
		return err
	}
	s3Cfg, err := s.S3Config()
	if err != nil {
		return errors.Wrap(err, "Retrieving s3 configuration")
	}

	m365 := credentials.GetM365()
	a := repository.Account{
		TenantID:     m365.TenantID,
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
	}
	if len(cfgTenantID) > 0 {
		a.TenantID = cfgTenantID
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
		return errors.Wrap(err, "Failed to initialize a new S3 repository")
	}
	defer utils.CloseRepo(ctx, r)

	fmt.Printf("Initialized a S3 repository within bucket %s.\n", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(s3Cfg, a); err != nil {
		return errors.Wrap(err, "Failed to write repository configuration")
	}
	return nil
}

// `corso repo connect s3 [<flag>...]`
var s3ConnectCmd = &cobra.Command{
	Use:   "s3",
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

	overrides := map[string]string{
		credentials.AWSAccessKeyID: accessKey,
		storage.Bucket:             bucket,
		storage.Endpoint:           endpoint,
		storage.Prefix:             prefix,
	}
	s, cfgTenantID, err := config.MakeS3Config(true, overrides)
	if err != nil {
		return err
	}
	s3Cfg, err := s.S3Config()
	if err != nil {
		return errors.Wrap(err, "Retrieving s3 configuration")
	}

	m365 := credentials.GetM365()
	a := repository.Account{
		TenantID:     m365.TenantID,
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
	}
	if len(cfgTenantID) > 0 {
		a.TenantID = cfgTenantID
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

	if err = config.WriteRepoConfig(s3Cfg, a); err != nil {
		return errors.Wrap(err, "Failed to write repository configuration")
	}
	return nil
}
