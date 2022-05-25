package kopia

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"

	"github.com/alcionai/corso/pkg/storage"
)

const (
	defaultS3Endpoint = "s3.amazonaws.com" // matches kopia's default value
)

func s3BlobStorage(ctx context.Context, cfg storage.S3Config) (blob.Storage, error) {
	opts := s3.Options{
		AccessKeyID:     cfg.AccessKey,
		BucketName:      cfg.Bucket,
		Endpoint:        defaultS3Endpoint,
		SecretAccessKey: cfg.SecretKey,
		SessionToken:    cfg.SessionToken,
	}
	return s3.New(ctx, &opts)
}
