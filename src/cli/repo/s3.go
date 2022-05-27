package repo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
)

// s3 bucket info from flags
var (
	accessKey string
	bucket    string
	endpoint  string
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
	return c
}

// `corso repo init s3 [<flag>...]`
var s3InitCmd = &cobra.Command{
	Use:   "s3",
	Short: "Initialize a S3 repository",
	Long:  `Bootstraps a new S3 repository and connects it to your m356 account.`,
	Run:   initS3Cmd,
	Args:  cobra.NoArgs,
}

// initializes a s3 repo.
func initS3Cmd(cmd *cobra.Command, args []string) {
	mv := getM365Vars()
	s3Cfg := makeS3Config()
	fmt.Printf(
		"Called - %s\n\tbucket:\t%s\n\tkey:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n\tfound awsSecret:\t%v\n",
		cmd.CommandPath(),
		s3Cfg.Bucket,
		s3Cfg.AccessKey,
		mv.clientID,
		len(mv.clientSecret) > 0,
		len(s3Cfg.SecretKey) > 0)

	a := repository.Account{
		TenantID:     mv.tenantID,
		ClientID:     mv.clientID,
		ClientSecret: mv.clientSecret,
	}
	s := storage.NewStorage(storage.ProviderS3, s3Cfg)

	if _, err := repository.Initialize(cmd.Context(), a, s); err != nil {
		fmt.Printf("Failed to initialize a new S3 repository: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Initialized a S3 repository within bucket %s.\n", s3Cfg.Bucket)
}

// `corso repo connect s3 [<flag>...]`
var s3ConnectCmd = &cobra.Command{
	Use:   "s3",
	Short: "Connect to a S3 repository",
	Long:  `Ensures a connection to an existing S3 repository.`,
	Run:   connectS3Cmd,
	Args:  cobra.NoArgs,
}

// connects to an existing s3 repo.
func connectS3Cmd(cmd *cobra.Command, args []string) {
	mv := getM365Vars()
	s3Cfg := makeS3Config()
	fmt.Printf(
		"Called - %s\n\tbucket:\t%s\n\tkey:\t%s\n\t356Client:\t%s\n\tfound 356Secret:\t%v\n\tfound awsSecret:\t%v\n",
		cmd.CommandPath(),
		s3Cfg.Bucket,
		s3Cfg.AccessKey,
		mv.clientID,
		len(mv.clientSecret) > 0,
		len(s3Cfg.SecretKey) > 0)

	a := repository.Account{
		TenantID:     mv.tenantID,
		ClientID:     mv.clientID,
		ClientSecret: mv.clientSecret,
	}
	s := storage.NewStorage(storage.ProviderS3, s3Cfg)

	if _, err := repository.Connect(cmd.Context(), a, s); err != nil {
		fmt.Printf("Failed to connect to the S3 repository: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to S3 bucket %s.\n", s3Cfg.Bucket)
}

// helper for aggregating aws connection details.
func makeS3Config() storage.S3Config {
	ak := os.Getenv("AWS_ACCESS_KEY_ID")
	if len(accessKey) > 0 {
		ak = accessKey
	}
	return storage.S3Config{
		AccessKey:    ak,
		Bucket:       bucket,
		Endpoint:     endpoint,
		SecretKey:    os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionToken: os.Getenv("AWS_SESSION_TOKEN"),
	}
}
