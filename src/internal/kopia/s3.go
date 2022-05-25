package kopia

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"

	"github.com/alcionai/corso/pkg/storage"
)

func s3BlobStorage(ctx context.Context, cfg storage.S3Config) (blob.Storage, error) {
	opts := s3.Options{
		BucketName:      cfg.Bucket,
		AccessKeyID:     cfg.AccessKey,
		SecretAccessKey: cfg.SecretKey,
	}
	return s3.New(ctx, &opts)
}
