package storage

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/credentials"
)

type S3Config struct {
	credentials.AWS // requires: AccessKey, SecretKey, SessionToken

	Bucket   string // required
	Endpoint string
	Prefix   string
}

// config key consts
const (
	keyS3AccessKey    = "s3_accessKey"
	keyS3Bucket       = "s3_bucket"
	keyS3Endpoint     = "s3_endpoint"
	keyS3Prefix       = "s3_prefix"
	keyS3SecretKey    = "s3_secretKey"
	keyS3SessionToken = "s3_sessionToken"
)

// config exported name consts
const (
	Bucket   = "bucket"
	Endpoint = "endpoint"
	Prefix   = "prefix"
)

func (c S3Config) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		keyS3AccessKey:    c.AccessKey,
		keyS3Bucket:       c.Bucket,
		keyS3Endpoint:     c.Endpoint,
		keyS3Prefix:       c.Prefix,
		keyS3SecretKey:    c.SecretKey,
		keyS3SessionToken: c.SessionToken,
	}
	return cfg, c.validate()
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() (S3Config, error) {
	c := S3Config{}
	if len(s.Config) > 0 {
		c.AccessKey = orEmptyString(s.Config[keyS3AccessKey])
		c.Bucket = orEmptyString(s.Config[keyS3Bucket])
		c.Endpoint = orEmptyString(s.Config[keyS3Endpoint])
		c.Prefix = orEmptyString(s.Config[keyS3Prefix])
		c.SecretKey = orEmptyString(s.Config[keyS3SecretKey])
		c.SessionToken = orEmptyString(s.Config[keyS3SessionToken])
	}
	return c, c.validate()
}

func (c S3Config) validate() error {
	check := map[string]string{
		credentials.AWSAccessKeyID:     c.AccessKey,
		credentials.AWSSecretAccessKey: c.SecretKey,
		credentials.AWSSessionToken:    c.SessionToken,
		Bucket:                         c.Bucket,
	}
	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}
	return nil
}
