package s3

import (
	"context"

	"github.com/alcionai/corso/internal/kopia/s3"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/pkg/errors"
)

// Config defines communication with a s3 repository provider.
type Config struct {
	repository.InitConnector
	Bucket          string
	AccessKey       string
	SecretAccessKey string
}

// NewInitializer generates a S3 configuration that lets a Repository initailize and connect to s3.
func NewInitializer(ctx context.Context, bucket, accessKey, secretKey string) (Config, error) {
	init, err := s3.Initializer(ctx, bucket, accessKey, secretKey)
	if err != nil {
		return Config{}, errors.Wrap(err, "configuring s3 initialization")
	}
	return Config{
		InitConnector:   init,
		Bucket:          bucket,
		AccessKey:       accessKey,
		SecretAccessKey: secretKey,
	}, nil
}

// NewConnector generates a S3 configuration that lets a Repository connect to s3.
func NewConnector(ctx context.Context, bucket, accessKey, secretKey string) (Config, error) {
	conn, err := s3.Connector(ctx, bucket, accessKey, secretKey)
	if err != nil {
		return Config{}, errors.Wrap(err, "configuring s3 connection")
	}
	return Config{
		InitConnector:   conn,
		Bucket:          bucket,
		AccessKey:       accessKey,
		SecretAccessKey: secretKey,
	}, nil
}
