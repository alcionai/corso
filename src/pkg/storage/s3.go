package storage

import (
	"strings"

	"github.com/pkg/errors"
)

type S3Config struct {
	Bucket   string // required
	Endpoint string
	Prefix   string
}

// config key consts
const (
	keyS3Bucket   = "s3_bucket"
	keyS3Endpoint = "s3_endpoint"
	keyS3Prefix   = "s3_prefix"
)

// config exported name consts
const (
	Bucket   = "bucket"
	Endpoint = "endpoint"
	Prefix   = "prefix"
)

func (c S3Config) Normalize() S3Config {
	// Minio doesn't accept scheme prefixes.
	c.Bucket = strings.TrimPrefix(c.Bucket, "s3://")
	return c
}

// StringConfig transforms a s3Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c S3Config) StringConfig() (map[string]string, error) {
	c = c.Normalize()
	cfg := map[string]string{
		keyS3Bucket:   c.Bucket,
		keyS3Endpoint: c.Endpoint,
		keyS3Prefix:   c.Prefix,
	}

	return cfg, c.validate()
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() (S3Config, error) {
	c := S3Config{}

	if len(s.Config) > 0 {
		c.Bucket = orEmptyString(s.Config[keyS3Bucket])
		c.Endpoint = orEmptyString(s.Config[keyS3Endpoint])
		c.Prefix = orEmptyString(s.Config[keyS3Prefix])
	}

	return c, c.validate()
}

func (c S3Config) validate() error {
	check := map[string]string{
		Bucket: c.Bucket,
	}
	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}

	return nil
}
