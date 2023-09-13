package storage

import (
	"strconv"

	"github.com/alcionai/clues"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/spf13/cast"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/credentials"
)

const (
	// S3 config
	StorageProviderTypeKey    = "provider"
	BucketNameKey             = "bucket"
	EndpointKey               = "endpoint"
	PrefixKey                 = "prefix"
	DisableTLSKey             = "disable_tls"
	DisableTLSVerificationKey = "disable_tls_verification"
	RepoID                    = "repo_id"

	AccessKey       = "aws_access_key_id"
	SecretAccessKey = "aws_secret_access_key"
	SessionToken    = "aws_session_token"
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

func (c S3Config) Normalize() S3Config {
	return S3Config{
		Bucket:         common.NormalizeBucket(c.Bucket),
		Endpoint:       c.Endpoint,
		Prefix:         common.NormalizePrefix(c.Prefix),
		DoNotUseTLS:    c.DoNotUseTLS,
		DoNotVerifyTLS: c.DoNotVerifyTLS,
	}
}

// No need to return error here. Viper returns empty values.
func s3ConfigsFromStore(kvg KVStoreGetter) S3Config {
	var s3Config S3Config

	s3Config.Bucket = cast.ToString(kvg.Get(BucketNameKey))
	s3Config.Endpoint = cast.ToString(kvg.Get(EndpointKey))
	s3Config.Prefix = cast.ToString(kvg.Get(PrefixKey))
	s3Config.DoNotUseTLS = cast.ToBool(kvg.Get(DisableTLSKey))
	s3Config.DoNotVerifyTLS = cast.ToBool(kvg.Get(DisableTLSVerificationKey))

	return s3Config
}

func s3CredsFromStore(
	kvg KVStoreGetter,
	s3Config S3Config,
) S3Config {
	s3Config.AccessKey = cast.ToString(kvg.Get(AccessKey))
	s3Config.SecretKey = cast.ToString(kvg.Get(SecretAccessKey))
	s3Config.SessionToken = cast.ToString(kvg.Get(SessionToken))

	return s3Config
}

var _ Configurer = S3Config{}

func (c S3Config) FetchConfigFromStore(
	kvg KVStoreGetter,
	readConfigFromStore bool,
	matchFromConfig bool,
	overrides map[string]string,
) (Configurer, error) {
	var (
		s3Cfg S3Config
		err   error
	)

	if readConfigFromStore {
		s3Cfg = s3ConfigsFromStore(kvg)

		if b, ok := overrides[Bucket]; ok {
			overrides[Bucket] = common.NormalizeBucket(b)
		}

		if p, ok := overrides[Prefix]; ok {
			overrides[Prefix] = common.NormalizePrefix(p)
		}

		if matchFromConfig {
			providerType := cast.ToString(kvg.Get(StorageProviderTypeKey))
			if providerType != ProviderS3.String() {
				return S3Config{}, clues.New("unsupported storage provider: " + providerType)
			}

			// This is matching override values from config file.
			if err := mustMatchConfig(kvg, s3Overrides(overrides)); err != nil {
				return S3Config{}, clues.Wrap(err, "verifying s3 configs in corso config file")
			}
		}
	}

	s3Cfg = s3CredsFromStore(kvg, s3Cfg)
	aws := credentials.GetAWS(overrides)

	if len(aws.AccessKey) <= 0 || len(aws.SecretKey) <= 0 {
		_, err = defaults.CredChain(
			defaults.Config().WithCredentialsChainVerboseErrors(true),
			defaults.Handlers()).Get()
		if err != nil && (len(s3Cfg.AccessKey) > 0 || len(s3Cfg.SecretKey) > 0) {
			aws = credentials.AWS{
				AccessKey:    s3Cfg.AccessKey,
				SecretKey:    s3Cfg.SecretKey,
				SessionToken: s3Cfg.SessionToken,
			}
			err = nil
		}

		if err != nil {
			return S3Config{}, clues.Wrap(err, "validating aws credentials")
		}
	}

	s3Cfg = S3Config{
		AWS:      aws,
		Bucket:   str.First(overrides[Bucket], s3Cfg.Bucket),
		Endpoint: str.First(overrides[Endpoint], s3Cfg.Endpoint, "s3.amazonaws.com"),
		Prefix:   str.First(overrides[Prefix], s3Cfg.Prefix),
		DoNotUseTLS: str.ParseBool(str.First(
			overrides[DoNotUseTLS],
			strconv.FormatBool(s3Cfg.DoNotUseTLS),
			"false")),
		DoNotVerifyTLS: str.ParseBool(str.First(
			overrides[DoNotVerifyTLS],
			strconv.FormatBool(s3Cfg.DoNotVerifyTLS),
			"false")),
	}

	return s3Cfg, s3Cfg.validate()
}

var _ WriteConfigToStorer = S3Config{}

func (c S3Config) WriteConfigToStore(
	kvs KVStoreSetter,
) {
	s3Config := c.Normalize()

	kvs.Set(StorageProviderTypeKey, ProviderS3.String())
	kvs.Set(BucketNameKey, s3Config.Bucket)
	kvs.Set(EndpointKey, s3Config.Endpoint)
	kvs.Set(PrefixKey, s3Config.Prefix)
	kvs.Set(DisableTLSKey, s3Config.DoNotUseTLS)
	kvs.Set(DisableTLSVerificationKey, s3Config.DoNotVerifyTLS)
}

// StringConfig transforms a s3Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c S3Config) StringConfig() (map[string]string, error) {
	cn := c.Normalize()
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

	return cfg, c.validate()
}

// S3Config retrieves the S3Config details from the Storage config.
func MakeS3ConfigFromMap(config map[string]string) (S3Config, error) {
	c := S3Config{}

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

var constToTomlKeyMap = map[string]string{
	Bucket:                 BucketNameKey,
	Endpoint:               EndpointKey,
	Prefix:                 PrefixKey,
	StorageProviderTypeKey: StorageProviderTypeKey,
}

// mustMatchConfig compares the values of each key to their config file value in store.
// If any value differs from the store value, an error is returned.
// values in m that aren't stored in the config are ignored.
func mustMatchConfig(kvg KVStoreGetter, m map[string]string) error {
	for k, v := range m {
		if len(v) == 0 {
			continue // empty variables will get caught by configuration validators, if necessary
		}

		tomlK, ok := constToTomlKeyMap[k]
		if !ok {
			continue // m may declare values which aren't stored in the config file
		}

		vv := cast.ToString(kvg.Get(tomlK))
		if v != vv {
			return clues.New("value of " + k + " (" + v + ") does not match corso configuration value (" + vv + ")")
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
