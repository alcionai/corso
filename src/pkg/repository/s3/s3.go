package s3

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
	kopiaS3 "github.com/kopia/kopia/repo/blob/s3"
)

// Config defines communication with a s3 repository provider.
type Config struct {
	Bucket          string // the S3 storage bucket name
	AccessKey       string // access key to the S3 bucket
	SecretAccessKey string // s3 access key secret
}

// NewConfig generates a S3 configuration struct to use for interfacing with a s3 storage
// bucket using a repository.Repository.
func NewConfig(bucket, accessKey, secretKey string) Config {
	return Config{
		Bucket:          bucket,
		AccessKey:       accessKey,
		SecretAccessKey: secretKey,
	}
}

// KopiaStorage produces a kopia/blob Storage handle for connecting to s3.
func (c Config) KopiaStorage(ctx context.Context, create bool) (blob.Storage, error) {
	opts := kopiaS3.Options{
		BucketName:      c.Bucket,
		AccessKeyID:     c.AccessKey,
		SecretAccessKey: c.SecretAccessKey,
	}
	return kopiaS3.New(ctx, &opts)
}
