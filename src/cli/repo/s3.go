package repo

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/utils"
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
	mv := utils.GetM365Vars()
	s3Cfg, commonCfg, err := makeS3Config()
	if err != nil {
		return err
	}

	fmt.Printf(
		"Called - %s\n\tbucket:\t%s\n\tkey:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n\tfound awsSecret:\t%v\n",
		cmd.CommandPath(),
		s3Cfg.Bucket,
		s3Cfg.AccessKey,
		mv.ClientID,
		len(mv.ClientSecret) > 0,
		len(s3Cfg.SecretKey) > 0)

	a := repository.Account{
		TenantID:     mv.TenantID,
		ClientID:     mv.ClientID,
		ClientSecret: mv.ClientSecret,
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
	mv := utils.GetM365Vars()
	s3Cfg, commonCfg, err := makeS3Config()
	if err != nil {
		return err
	}

	fmt.Printf(
		"Called - %s\n\tbucket:\t%s\n\tkey:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n\tfound awsSecret:\t%v\n",
		cmd.CommandPath(),
		s3Cfg.Bucket,
		s3Cfg.AccessKey,
		mv.ClientID,
		len(mv.ClientSecret) > 0,
		len(s3Cfg.SecretKey) > 0)

	a := repository.Account{
		TenantID:     mv.TenantID,
		ClientID:     mv.ClientID,
		ClientSecret: mv.ClientSecret,
	}
	s, err := storage.NewStorage(storage.ProviderS3, s3Cfg, commonCfg)
	if err != nil {
		errors.Wrap(err, "Failed to configure storage provider")
	}

	r, err := repository.Connect(cmd.Context(), a, s)
	if err != nil {
		errors.Wrap(err, "Failed to connect to the S3 repository")
	}
	defer utils.CloseRepo(cmd.Context(), r)

	fmt.Printf("Connected to S3 bucket %s.\n", s3Cfg.Bucket)
	return nil
}

// helper for aggregating aws connection details.
func makeS3Config() (storage.S3Config, storage.CommonConfig, error) {
	ak := os.Getenv(storage.AWS_ACCESS_KEY_ID)
	if len(accessKey) > 0 {
		ak = accessKey
	}
	secretKey := os.Getenv(storage.AWS_SECRET_ACCESS_KEY)
	sessToken := os.Getenv(storage.AWS_SESSION_TOKEN)
	corsoPasswd := os.Getenv(storage.CORSO_PASSWORD)

	return storage.S3Config{
			AccessKey:    ak,
			Bucket:       bucket,
			Endpoint:     endpoint,
			Prefix:       prefix,
			SecretKey:    secretKey,
			SessionToken: sessToken,
		},
		storage.CommonConfig{
			CorsoPassword: corsoPasswd,
		},
		utils.RequireProps(map[string]string{
			storage.AWS_ACCESS_KEY_ID:     ak,
			"bucket":                      bucket,
			storage.AWS_SECRET_ACCESS_KEY: secretKey,
			storage.AWS_SESSION_TOKEN:     sessToken,
			storage.CORSO_PASSWORD:        corsoPasswd,
		})
}
