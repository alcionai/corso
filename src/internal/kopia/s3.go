package kopia

import (
	"context"

	awscreds "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"
	miniocreds "github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	defaultS3Endpoint = "s3.amazonaws.com" // matches kopia's default value
)

func s3BlobStorage(ctx context.Context, s storage.Storage) (blob.Storage, error) {
	cfg, err := s.S3Config()
	if err != nil {
		return nil, err
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
		Creds:          credentials(s.Creds),
	}

	return s3.New(ctx, &opts, false)
}

// credentials converts an AWS Credential to a Minio credential (which Kopia uses)
func credentials(creds *awscreds.Credentials) *miniocreds.Credentials {
	if creds == nil {
		return nil
	}

	return miniocreds.New(&minioProvider{creds: creds})
}

// minioProvider is a shim that implements the Minio `Provider` interface
// for an AWS credential
type minioProvider struct {
	creds *awscreds.Credentials
}

func (mp *minioProvider) Retrieve() (miniocreds.Value, error) {
	v, err := mp.creds.Get()
	if err != nil {
		return miniocreds.Value{}, err
	}

	return miniocreds.Value{
		AccessKeyID:     v.AccessKeyID,
		SecretAccessKey: v.SecretAccessKey,
		SessionToken:    v.SessionToken,
		SignerType:      miniocreds.SignatureV4,
	}, nil
}

func (mp *minioProvider) IsExpired() bool {
	return mp.creds.IsExpired()
}
