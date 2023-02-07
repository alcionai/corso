package storage

import (
	"strconv"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
)

type S3Config struct {
	Bucket         string // required
	Endpoint       string
	Prefix         string
	DoNotUseTLS    bool
	DoNotVerifyTLS bool
}

// config key consts
const (
	keyS3Bucket         = "s3_bucket"
	keyS3Endpoint       = "s3_endpoint"
	keyS3Prefix         = "s3_prefix"
	keyS3DoNotUseTLS    = "s3_donotusetls"
	keyS3DoNotVerifyTLS = "s3_donotverifytls"
)

// config exported name consts
const (
	Bucket         = "bucket"
	Endpoint       = "endpoint"
	Prefix         = "prefix"
	DoNotUseTLS    = "donotusetls"
	DoNotVerifyTLS = "donotverifytls"
)

func (c S3Config) Normalize() S3Config {
	return S3Config{
		Bucket:         common.NormalizeBucket(c.Bucket),
		Endpoint:       c.Endpoint,
		Prefix:         common.NormalizePrefix(c.Prefix),
		DoNotUseTLS:    c.DoNotUseTLS,
		DoNotVerifyTLS: c.DoNotVerifyTLS,
	}
}

// StringConfig transforms a s3Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c S3Config) StringConfig() (map[string]string, error) {
	cn := c.Normalize()
	cfg := map[string]string{
		keyS3Bucket:         cn.Bucket,
		keyS3Endpoint:       cn.Endpoint,
		keyS3Prefix:         cn.Prefix,
		keyS3DoNotUseTLS:    strconv.FormatBool(cn.DoNotUseTLS),
		keyS3DoNotVerifyTLS: strconv.FormatBool(cn.DoNotVerifyTLS),
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
		c.DoNotUseTLS = common.ParseBool(s.Config[keyS3DoNotUseTLS])
		c.DoNotVerifyTLS = common.ParseBool(s.Config[keyS3DoNotVerifyTLS])
	}

	return c, c.validate()
}

func (c S3Config) validate() error {
	check := map[string]string{
		Bucket: c.Bucket,
	}
	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, errors.New(k))
		}
	}

	return nil
}
