package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"

	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	defaultS3Endpoint = "s3.amazonaws.com" // matches kopia's default value
)

func s3BlobStorage(
	ctx context.Context,
	repoOpts repository.Options,
	s storage.Storage,
) (blob.Storage, error) {
	cfg, err := s.GetStorageConfig()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	// Cast to S3Config
	s3Cfg := cfg.(storage.S3Config)

	endpoint := defaultS3Endpoint
	if len(s3Cfg.Endpoint) > 0 {
		endpoint = s3Cfg.Endpoint
	}

	opts := s3.Options{
		BucketName:          s3Cfg.Bucket,
		Endpoint:            endpoint,
		Prefix:              s3Cfg.Prefix,
		DoNotUseTLS:         s3Cfg.DoNotUseTLS,
		DoNotVerifyTLS:      s3Cfg.DoNotVerifyTLS,
		Tags:                s.SessionTags,
		SessionName:         s.SessionName,
		RoleARN:             s.Role,
		RoleDuration:        s.SessionDuration,
		AccessKeyID:         s3Cfg.AccessKey,
		SecretAccessKey:     s3Cfg.SecretKey,
		SessionToken:        s3Cfg.SessionToken,
		TLSHandshakeTimeout: 60,
		PointInTime:         repoOpts.ViewTimestamp,
	}

	store, err := s3.New(ctx, &opts, false)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return store, nil
}
