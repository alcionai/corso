package repo

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/flags"
)

// `corso repo <subcmd> s3 [<flag>...]`
var providerS3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Initialize a s3 repository.",
	Long:  `TODO: exhaustive details on initializing a s3 repo.`,
	Args:  cobra.NoArgs,
}

var (
	bucket    string
	accessKey string

	providerS3Flags = []flags.CliFlag{
		{
			Name:        "bucket",
			Description: "Name of the S3 bucket (required).",
			Required:    true,
			VarType:     flags.StringType,
			Var:         &bucket,
		},
		{
			Name:        "access-key",
			Description: "Access key ID (replaces the AWS_ACCESS_KEY_ID env variable).",
			Required:    false,
			VarType:     flags.StringType,
			Var:         &accessKey,
		},
	}
)

// initializes a s3 repo.
func initS3Cmd(cmd *cobra.Command, args []string) {
	fmt.Printf(
		"Called -\n`corso repo init s3`\nbucket:\t%s\nkey:\t%s\n",
		bucket,
		accessKey)
}

// connects to an existing s3 repo.
func connectS3Cmd(cmd *cobra.Command, args []string) {
	fmt.Printf(
		"Called -\n`corso repo connect s3`\nbucket:\t%s\nkey:\t%s\n",
		bucket,
		accessKey)
}
