package repo

import (
	"fmt"
	"os"

	"github.com/alcionai/corso/repository"
	"github.com/spf13/cobra"
)

// s3 bucket info from flags
var (
	bucket    string
	accessKey string
)

// called by repo.go to map parent subcommands to provider-specific handling.
func addS3Commands(parent *cobra.Command) *cobra.Command {
	var c *cobra.Command
	switch parent.Use {
	case initUse:
		c = s3InitCmd
	case connectUse:
		c = s3ConnectCmd
	}

	parent.AddCommand(c)

	fs := c.Flags()
	fs.StringVar(&bucket, "bucket", "", "Name of the S3 bucket (required).")
	fs.StringVar(&accessKey, "access-key", "", "Access key ID (replaces the AWS_ACCESS_KEY_ID env variable).")

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
	oci := os.Getenv("O365_CLIENT_ID")
	ocs := os.Getenv("O356_SECRET")
	asak := os.Getenv("AWS_SECRET_ACCESS_KEY")
	fmt.Printf(
		"Called -\n`corso repo init s3`\nbucket:\t%s\nkey:\t%s\n356Client:\t%s\nfound 356Secret:\t%v\nfound awsSecret:\t%v\n",
		bucket,
		accessKey,
		oci,
		len(ocs) > 0,
		len(asak) > 0)

	r := repository.NewS3(bucket, accessKey, asak, m365Tenant, m365ClientID, m365ClientSecret)

	if err := r.Initialize(); err != nil {
		fmt.Printf("Failed to initialize a new S3 repository: %v", err)
		os.Exit(1)
	}
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
	oci := os.Getenv("O365_CLIENT_ID")
	ocs := os.Getenv("O356_SECRET")
	asak := os.Getenv("AWS_SECRET_ACCESS_KEY")
	fmt.Printf(
		"Called -\n`corso repo connect s3`\nbucket:\t%s\nkey:\t%s\n356Client:\t%s\nfound 356Secret:\t%v\nfound awsSecret:\t%v\n",
		bucket,
		accessKey,
		oci,
		len(ocs) > 0,
		len(asak) > 0)

	// TODO: this should retrieve an existing repo, not generate a new one
	r := repository.NewS3(bucket, accessKey, asak, m365Tenant, m365ClientID, m365ClientSecret)

	if err := r.Connect(); err != nil {
		fmt.Printf("Failed to connect to the S3 repository: %v", err)
		os.Exit(1)
	}
}
