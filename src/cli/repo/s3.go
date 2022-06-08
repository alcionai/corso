package repo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

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
	log := logger.Ctx(cmd.Context())

	m365 := credentials.GetM365()
	s3Cfg, commonCfg, err := makeS3Config()
	if err != nil {
		return err
	}

	log.Debug().
		Str("bucket", s3Cfg.Bucket).
		Str("clientID", m365.ClientID).
		Bool("hasClientSecret", len(m365.ClientSecret) > 0).
		Str("accessKey", s3Cfg.AccessKey).
		Bool("hasSecretKey", len(s3Cfg.SecretKey) > 0).
		Msg("Called - " + cmd.CommandPath())

	a := repository.Account{
		TenantID:     m365.TenantID,
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
	}
	s, err := storage.NewStorage(storage.ProviderS3, s3Cfg, commonCfg)
	if err != nil {
		return errors.Wrap(err, "Failed to configure storage provider")
	}

	r, err := repository.Initialize(cmd.Context(), a, s)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize a new S3 repository")
	}
	defer utils.CloseRepo(cmd.Context(), r)

	fmt.Printf("Initialized a S3 repository within bucket %s.\n", s3Cfg.Bucket)
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
	log := logger.Ctx(cmd.Context())

	m365 := credentials.GetM365()
	s3Cfg, commonCfg, err := makeS3Config()
	if err != nil {
		return err
	}

	log.Debug().
		Str("bucket", s3Cfg.Bucket).
		Str("clientID", m365.ClientID).
		Bool("hasClientSecret", len(m365.ClientSecret) > 0).
		Str("accessKey", s3Cfg.AccessKey).
		Bool("hasSecretKey", len(s3Cfg.SecretKey) > 0).
		Msg("Called - " + cmd.CommandPath())

	a := repository.Account{
		TenantID:     m365.TenantID,
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
	}
	s, err := storage.NewStorage(storage.ProviderS3, s3Cfg, commonCfg)
	if err != nil {
		return errors.Wrap(err, "Failed to configure storage provider")
	}

	r, err := repository.Connect(cmd.Context(), a, s)
	if err != nil {
		return errors.Wrap(err, "Failed to connect to the S3 repository")
	}
	defer utils.CloseRepo(cmd.Context(), r)

	fmt.Printf("Connected to S3 bucket %s.\n", s3Cfg.Bucket)
	return nil
}

// helper for aggregating aws connection details.
func makeS3Config() (storage.S3Config, storage.CommonConfig, error) {
	aws := credentials.GetAWS(map[string]string{credentials.AWSAccessKeyID: accessKey})
	corso := credentials.GetCorso()
	return storage.S3Config{
			AWS:      aws,
			Bucket:   bucket,
			Endpoint: endpoint,
			Prefix:   prefix,
		},
		storage.CommonConfig{
			Corso: corso,
		},
		utils.RequireProps(map[string]string{
			credentials.AWSAccessKeyID:     aws.AccessKey,
			"bucket":                       bucket,
			credentials.AWSSecretAccessKey: aws.SecretKey,
			credentials.AWSSessionToken:    aws.SessionToken,
			credentials.CorsoPassword:      corso.CorsoPassword,
		})
}
