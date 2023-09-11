package flags

import "github.com/spf13/cobra"

// S3 bucket flags
const (
	BucketFN          = "bucket"
	EndpointFN        = "endpoint"
	PrefixFN          = "prefix"
	DoNotUseTLSFN     = "disable-tls"
	DoNotVerifyTLSFN  = "disable-tls-verification"
	SucceedIfExistsFN = "succeed-if-exists"
)

// S3 bucket flag values
var (
	BucketFV          string
	EndpointFV        string
	PrefixFV          string
	DoNotUseTLSFV     bool
	DoNotVerifyTLSFV  bool
	SucceedIfExistsFV bool
)

// S3 bucket flags
func AddS3BucketFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&CorsoPassphraseFV,
		CorsoPassphraseFN,
		"",
		"Passphrase to protect encrypted repository contents")

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic and more frequently used flags take precedence.
	fs.StringVar(&BucketFV, BucketFN, "", "Name of S3 bucket for repo. (required)")
	fs.StringVar(&PrefixFV, PrefixFN, "", "Repo prefix within bucket.")
	fs.StringVar(&EndpointFV, EndpointFN, "", "S3 service endpoint.")
	fs.BoolVar(&DoNotUseTLSFV, DoNotUseTLSFN, false, "Disable TLS (HTTPS)")
	fs.BoolVar(&DoNotVerifyTLSFV, DoNotVerifyTLSFN, false, "Disable TLS (HTTPS) certificate verification.")

	// In general, we don't want to expose this flag to users and have them mistake it
	// for a broad-scale idempotency solution.  We can un-hide it later the need arises.
	fs.BoolVar(&SucceedIfExistsFV, SucceedIfExistsFN, false, "Exit with success if the repo has already been initialized.")
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))
}
