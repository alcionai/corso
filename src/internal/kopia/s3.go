package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"

	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	defaultS3Endpoint = "s3.amazonaws.com" // matches kopia's default value
)

func s3BlobStorage(ctx context.Context, s storage.Storage) (blob.Storage, error) {
	cfg, err := s.S3Config()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	endpoint := defaultS3Endpoint
	if len(cfg.Endpoint) > 0 {
		endpoint = cfg.Endpoint
	}

	opts := s3.Options{
		BucketName:     cfg.Bucket,
		Endpoint:       endpoint,
		Prefix:         cfg.Prefix,
		DoNotUseTLS:    cfg.DoNotUseTLS,
		DoNotVerifyTLS: cfg.DoNotVerifyTLS,
	}

	store, err := s3.New(ctx, &opts, false)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return store, nil
}
