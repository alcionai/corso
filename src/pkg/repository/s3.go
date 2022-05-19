package repository

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
	kopiaS3 "github.com/kopia/kopia/repo/blob/s3"
)

// S3Config defines communication with a s3 repository provider.
type S3Config struct {
	Bucket          string // the S3 storage bucket name
	AccessKey       string // access key to the S3 bucket
	SecretAccessKey string // s3 access key secret
}

// NewS3 generates a S3 repository struct to use for interfacing with a s3 storage bucket.
func NewS3(
	bucket, accessKey, secretKey, tenantID, clientID, clientSecret string,
) Repository {
	return newRepo(
		S3Provider,
		Account{
			TenantID:     tenantID,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		S3Config{
			Bucket:          bucket,
			AccessKey:       accessKey,
			SecretAccessKey: secretKey,
		},
	)
}

// KopiaStorage produces a kopia/blob Storage handle for connecting to s3.
func (c S3Config) KopiaStorage(ctx context.Context, create bool) (blob.Storage, error) {
	opts := kopiaS3.Options{
		BucketName:      c.Bucket,
		AccessKeyID:     c.AccessKey,
		SecretAccessKey: c.SecretAccessKey,
	}
	return kopiaS3.New(ctx, &opts)
}
