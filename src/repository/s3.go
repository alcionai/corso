package repository

// S3Repository
type S3Repository struct {
	repo

	Provider        repoProvider // must be repository.S3Provider
	Bucket          string       // the S3 storage bucket name
	AccessKey       string       // access key to the S3 bucket
	SecretAccessKey string       // s3 access key secret
}

// NewS3 generates a S3 repository struct to use for interfacing with a s3 storage bucket.
func NewS3(
	bucket, accessKey, secretKey, tenantID, clientID, clientSecret string,
) S3Repository {
	return S3Repository{
		repo:            newRepo(tenantID, clientID, clientSecret),
		Provider:        S3Provider,
		Bucket:          bucket,
		AccessKey:       accessKey,
		SecretAccessKey: secretKey,
	}
}

// Initialize will:
//  * validate the m365 account & secrets
//  * connect to the m365 account to ensure communication capability
//  * validate the s3 bucket & secrets
//  * create the s3 bucket with its provided configuration
//  * store the configuration details
//  * connect to the s3 bucket
func (s3 S3Repository) Initialize() error {
	return nil
}

// Connect will:
//  * validate the m365 account details
//  * connect to the m365 account to ensure communication capability
//  * connect to the s3 bucket
func (s3 S3Repository) Connect() error {
	return nil
}
