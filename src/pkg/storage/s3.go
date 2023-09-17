package storage

import (
	"os"
	"strconv"

	"github.com/alcionai/clues"
	"github.com/spf13/cast"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/credentials"
)

type S3Config struct {
	credentials.AWS
	Bucket         string // required
	Endpoint       string
	Prefix         string
	DoNotUseTLS    bool
	DoNotVerifyTLS bool
}

// config key consts
const (
	keyS3AccessKey      = "s3_access_key"
	keyS3Bucket         = "s3_bucket"
	keyS3Endpoint       = "s3_endpoint"
	keyS3Prefix         = "s3_prefix"
	keyS3SecretKey      = "s3_secret_key"
	keyS3SessionToken   = "s3_session_token"
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

// config file keys
const (
	BucketNameKey             = "bucket"
	EndpointKey               = "endpoint"
	PrefixKey                 = "prefix"
	DisableTLSKey             = "disable_tls"
	DisableTLSVerificationKey = "disable_tls_verification"

	AccessKey       = "aws_access_key_id"
	SecretAccessKey = "aws_secret_access_key"
	SessionToken    = "aws_session_token"
)

var s3constToTomlKeyMap = map[string]string{
	Bucket:                 BucketNameKey,
	Endpoint:               EndpointKey,
	Prefix:                 PrefixKey,
	StorageProviderTypeKey: StorageProviderTypeKey,
}

func (c *S3Config) normalize() S3Config {
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
func (c *S3Config) StringConfig() (map[string]string, error) {
	cn := c.normalize()
	cfg := map[string]string{
		keyS3AccessKey:      c.AccessKey,
		keyS3Bucket:         cn.Bucket,
		keyS3Endpoint:       cn.Endpoint,
		keyS3Prefix:         cn.Prefix,
		keyS3SecretKey:      c.SecretKey,
		keyS3SessionToken:   c.SessionToken,
		keyS3DoNotUseTLS:    strconv.FormatBool(cn.DoNotUseTLS),
		keyS3DoNotVerifyTLS: strconv.FormatBool(cn.DoNotVerifyTLS),
	}

	return cfg, cn.validate()
}

func buildS3ConfigFromMap(config map[string]string) (*S3Config, error) {
	c := &S3Config{}

	if len(config) > 0 {
		c.AccessKey = orEmptyString(config[keyS3AccessKey])
		c.SecretKey = orEmptyString(config[keyS3SecretKey])
		c.SessionToken = orEmptyString(config[keyS3SessionToken])

		c.Bucket = orEmptyString(config[keyS3Bucket])
		c.Endpoint = orEmptyString(config[keyS3Endpoint])
		c.Prefix = orEmptyString(config[keyS3Prefix])
		c.DoNotUseTLS = str.ParseBool(config[keyS3DoNotUseTLS])
		c.DoNotVerifyTLS = str.ParseBool(config[keyS3DoNotVerifyTLS])
	}

	return c, c.validate()
}

func (c S3Config) validate() error {
	check := map[string]string{
		Bucket: c.Bucket,
	}
	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}

func s3Overrides(in map[string]string) map[string]string {
	return map[string]string{
		Bucket:                 in[Bucket],
		Endpoint:               in[Endpoint],
		Prefix:                 in[Prefix],
		DoNotUseTLS:            in[DoNotUseTLS],
		DoNotVerifyTLS:         in[DoNotVerifyTLS],
		StorageProviderTypeKey: in[StorageProviderTypeKey],
	}
}

func (c *S3Config) s3ConfigsFromStore(kvg Getter) {
	c.Bucket = cast.ToString(kvg.Get(BucketNameKey))
	c.Endpoint = cast.ToString(kvg.Get(EndpointKey))
	c.Prefix = cast.ToString(kvg.Get(PrefixKey))
	c.DoNotUseTLS = cast.ToBool(kvg.Get(DisableTLSKey))
	c.DoNotVerifyTLS = cast.ToBool(kvg.Get(DisableTLSVerificationKey))
}

func (c *S3Config) s3CredsFromStore(kvg Getter) {
	c.AccessKey = cast.ToString(kvg.Get(AccessKey))
	c.SecretKey = cast.ToString(kvg.Get(SecretAccessKey))
	c.SessionToken = cast.ToString(kvg.Get(SessionToken))
}

var _ Configurer = &S3Config{}

func (c *S3Config) ApplyConfigOverrides(
	kvg Getter,
	readConfigFromStore bool,
	matchFromConfig bool,
	overrides map[string]string,
) error {
	if readConfigFromStore {
		c.s3ConfigsFromStore(kvg)

		overrides[Bucket] = common.NormalizeBucket(overrides[Bucket])
		overrides[Prefix] = common.NormalizePrefix(overrides[Prefix])

		if matchFromConfig {
			providerType := cast.ToString(kvg.Get(StorageProviderTypeKey))
			if providerType != ProviderS3.String() {
				return clues.New("unsupported storage provider: " + providerType)
			}

			if err := mustMatchConfig(kvg, s3constToTomlKeyMap, s3Overrides(overrides)); err != nil {
				return clues.Wrap(err, "verifying s3 configs in corso config file")
			}
		}
	}

	c.s3CredsFromStore(kvg)

	aws := credentials.AWS{
		AccessKey: str.First(
			overrides[credentials.AWSAccessKeyID],
			os.Getenv(credentials.AWSAccessKeyID),
			c.AccessKey),
		SecretKey: str.First(
			overrides[credentials.AWSSecretAccessKey],
			os.Getenv(credentials.AWSSecretAccessKey),
			c.SecretKey),
		SessionToken: str.First(
			overrides[credentials.AWSSessionToken],
			os.Getenv(credentials.AWSSessionToken),
			c.SessionToken),
	}

	c.AWS = aws
	c.Bucket = str.First(overrides[Bucket], c.Bucket)
	c.Endpoint = str.First(overrides[Endpoint], c.Endpoint, "s3.amazonaws.com")
	c.Prefix = str.First(overrides[Prefix], c.Prefix)
	c.DoNotUseTLS = str.ParseBool(str.First(
		overrides[DoNotUseTLS],
		strconv.FormatBool(c.DoNotUseTLS),
		"false"))
	c.DoNotVerifyTLS = str.ParseBool(str.First(
		overrides[DoNotVerifyTLS],
		strconv.FormatBool(c.DoNotVerifyTLS),
		"false"))

	return c.validate()
}

var _ WriteConfigToStorer = &S3Config{}

func (c *S3Config) WriteConfigToStore(
	kvs Setter,
) {
	s3Config := c.normalize()

	kvs.Set(StorageProviderTypeKey, ProviderS3.String())
	kvs.Set(BucketNameKey, s3Config.Bucket)
	kvs.Set(EndpointKey, s3Config.Endpoint)
	kvs.Set(PrefixKey, s3Config.Prefix)
	kvs.Set(DisableTLSKey, s3Config.DoNotUseTLS)
	kvs.Set(DisableTLSVerificationKey, s3Config.DoNotVerifyTLS)
}
