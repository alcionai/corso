package flags

import (
	"github.com/spf13/cobra"
)

const (
	BackupFN             = "backup"
	AWSAccessKeyFN       = "aws-access-key"
	AWSSecretAccessKeyFN = "aws-secret-access-key"
	AWSSessionTokenFN    = "aws-session-token"
	FilesystemPathFN     = "path"
	// Corso Flags
	CorsoPassphraseFN = "passphrase"
)

var (
	BackupIDFV           string
	AWSAccessKeyFV       string
	AWSSecretAccessKeyFV string
	AWSSessionTokenFV    string
	FilesystemPathFV     string
	CorsoPassphraseFV    string
)

// AddBackupIDFlag adds the --backup flag.
func AddBackupIDFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringVar(&BackupIDFV, BackupFN, "", "ID of the backup to retrieve.")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(BackupFN))
	}
}

func AddAWSCredsFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&AWSAccessKeyFV, AWSAccessKeyFN, "", "S3 access key")
	fs.StringVar(&AWSSecretAccessKeyFV, AWSSecretAccessKeyFN, "", "S3 access secret")
	fs.StringVar(&AWSSessionTokenFV, AWSSessionTokenFN, "", "S3 session token")
}

// M365 flags
func AddCorsoPassphaseFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&CorsoPassphraseFV,
		CorsoPassphraseFN,
		"",
		"Passphrase to protect encrypted repository contents")
}

// AddFilesystemPathFlag adds the --path flag.
func AddFilesystemPathFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringVar(&FilesystemPathFV, FilesystemPathFN, "", "filesystem path to repo")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(FilesystemPathFN))
	}
}
