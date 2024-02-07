package flags

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/pkg/credentials"
	"github.com/alcionai/canario/src/pkg/storage"
)

// S3 bucket flags
const (
	BucketFN         = "bucket"
	EndpointFN       = "endpoint"
	PrefixFN         = "prefix"
	DoNotUseTLSFN    = "disable-tls"
	DoNotVerifyTLSFN = "disable-tls-verification"
)

// S3 bucket flag values
var (
	BucketFV         string
	EndpointFV       string
	PrefixFV         string
	DoNotUseTLSFV    bool
	DoNotVerifyTLSFV bool
)

// S3 bucket flags
func AddS3BucketFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic and more frequently used flags take precedence.
	fs.StringVar(&BucketFV, BucketFN, "", "Name of S3 bucket for repo. (required)")
	fs.StringVar(&PrefixFV, PrefixFN, "", "Repo prefix within bucket.")
	fs.StringVar(&EndpointFV, EndpointFN, "", "S3 service endpoint.")
	fs.BoolVar(&DoNotUseTLSFV, DoNotUseTLSFN, false, "Disable TLS (HTTPS)")
	fs.BoolVar(&DoNotVerifyTLSFV, DoNotVerifyTLSFN, false, "Disable TLS (HTTPS) certificate verification.")
}

func S3FlagOverrides(cmd *cobra.Command) map[string]string {
	fs := GetPopulatedFlags(cmd)
	return PopulateS3Flags(fs)
}

func PopulateS3Flags(flagset PopulatedFlags) map[string]string {
	s3Overrides := map[string]string{
		storage.StorageProviderTypeKey: storage.ProviderS3.String(),
	}

	if _, ok := flagset[AWSAccessKeyFN]; ok {
		s3Overrides[credentials.AWSAccessKeyID] = AWSAccessKeyFV
	}

	if _, ok := flagset[AWSSecretAccessKeyFN]; ok {
		s3Overrides[credentials.AWSSecretAccessKey] = AWSSecretAccessKeyFV
	}

	if _, ok := flagset[AWSSessionTokenFN]; ok {
		s3Overrides[credentials.AWSSessionToken] = AWSSessionTokenFV
	}

	if _, ok := flagset[BucketFN]; ok {
		s3Overrides[storage.Bucket] = BucketFV
	}

	if _, ok := flagset[PrefixFN]; ok {
		s3Overrides[storage.Prefix] = PrefixFV
	}

	if _, ok := flagset[DoNotUseTLSFN]; ok {
		s3Overrides[storage.DoNotUseTLS] = strconv.FormatBool(DoNotUseTLSFV)
	}

	if _, ok := flagset[DoNotVerifyTLSFN]; ok {
		s3Overrides[storage.DoNotVerifyTLS] = strconv.FormatBool(DoNotVerifyTLSFV)
	}

	if _, ok := flagset[EndpointFN]; ok {
		s3Overrides[storage.Endpoint] = EndpointFV
	}

	return s3Overrides
}
