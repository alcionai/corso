package config

import (
	"strconv"

	"github.com/alcionai/clues"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func s3ConfigsFromViper(vpr *viper.Viper) (S3Config, error) {
	var s3Config S3Config

	s3Config.Bucket = vpr.GetString(BucketNameKey)
	s3Config.Endpoint = vpr.GetString(EndpointKey)
	s3Config.Prefix = vpr.GetString(PrefixKey)
	s3Config.DoNotUseTLS = vpr.GetBool(DisableTLSKey)
	s3Config.DoNotVerifyTLS = vpr.GetBool(DisableTLSVerificationKey)

	return s3Config, nil
}

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func s3CredsFromViper(vpr *viper.Viper, s3Config S3Config) (S3Config, error) {
	s3Config.AccessKey = vpr.GetString(AccessKey)
	s3Config.SecretKey = vpr.GetString(SecretAccessKey)
	s3Config.SessionToken = vpr.GetString(SessionToken)

	return s3Config, nil
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

type StorageConfigurer interface {
	common.StringConfigurer
	//Normalize() StorageConfigurer
	WriteConfigToViper(vpr *viper.Viper)
}

// func NewStorageConfig(
// 	provider storage.StorageProvider,
// ) (StorageConfigurer, error) {
// 	switch provider {
// 	case storage.ProviderS3:
// 		return S3Config{}, nil
// 	default:
// 		return nil, clues.New("unsupported storage type")
// 	}
// }

// Hydrate from a config map
func NewStorageConfigFrom(s storage.Storage) (StorageConfigurer, error) {
	switch s.Provider {
	case storage.ProviderS3:
		return makeS3Config(s.Config)
	default:
		return nil, clues.New("unsupported storage type")
	}
}

type S3Config struct {
	credentials.AWS
	Bucket         string // required
	Endpoint       string
	Prefix         string
	DoNotUseTLS    bool
	DoNotVerifyTLS bool
}

// MakeS3Config retrieves the S3Config details from the Storage config.
func makeS3Config(config map[string]string) (StorageConfigurer, error) {
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

// TODO: Remove storageconfigurer return value.
// Have it contained in current s3config.
func fetchS3ConfigFromViper(
	vpr *viper.Viper,
	readConfigFromViper bool,
	matchFromConfig bool,
	overrides map[string]string,
) (StorageConfigurer, error) {
	var (
		s3Cfg S3Config
		err   error
	)

	if readConfigFromViper {
		if s3Cfg, err = s3ConfigsFromViper(vpr); err != nil {
			clues.Wrap(err, "reading s3 configs from corso config file")
		}

		if b, ok := overrides[Bucket]; ok {
			overrides[Bucket] = common.NormalizeBucket(b)
		}

		if p, ok := overrides[Prefix]; ok {
			overrides[Prefix] = common.NormalizePrefix(p)
		}

		if matchFromConfig {
			providerType := vpr.GetString(StorageProviderTypeKey)
			if providerType != storage.ProviderS3.String() {
				return nil, clues.New("unsupported storage provider" + providerType)
			}

			if err := mustMatchConfig(vpr, s3Overrides(overrides)); err != nil {
				return nil, clues.Wrap(err, "verifying s3 configs in corso config file")
			}
		}
	}

	if s3Cfg, err = s3CredsFromViper(vpr, s3Cfg); err != nil {
		return nil, clues.Wrap(err, "reading s3 configs from corso config file")
	}

	s3Overrides(overrides)
	aws := credentials.GetAWS(overrides)

	if len(aws.AccessKey) <= 0 || len(aws.SecretKey) <= 0 {
		_, err = defaults.CredChain(defaults.Config().WithCredentialsChainVerboseErrors(true), defaults.Handlers()).Get()
		if err != nil && (len(s3Cfg.AccessKey) > 0 || len(s3Cfg.SecretKey) > 0) {
			aws = credentials.AWS{
				AccessKey:    s3Cfg.AccessKey,
				SecretKey:    s3Cfg.SecretKey,
				SessionToken: s3Cfg.SessionToken,
			}
			err = nil
		}

		if err != nil {
			return nil, clues.Wrap(err, "validating aws credentials")
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
			"false",
		)),
		DoNotVerifyTLS: str.ParseBool(str.First(
			overrides[DoNotVerifyTLS],
			strconv.FormatBool(s3Cfg.DoNotVerifyTLS),
			"false",
		)),
	}

	return s3Cfg, s3Cfg.validate()
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
// TODO: Move these to storage?
const (
	Bucket         = "bucket"
	Endpoint       = "endpoint"
	Prefix         = "prefix"
	DoNotUseTLS    = "donotusetls"
	DoNotVerifyTLS = "donotverifytls"
)

func normalize(c S3Config) S3Config {
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
	cn := normalize(c)
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

func (c S3Config) WriteConfigToViper(vpr *viper.Viper) {
	cn := normalize(c)

	vpr.Set(StorageProviderTypeKey, storage.ProviderS3.String())
	vpr.Set(BucketNameKey, cn.Bucket)
	vpr.Set(EndpointKey, cn.Endpoint)
	vpr.Set(PrefixKey, cn.Prefix)
	vpr.Set(DisableTLSKey, cn.DoNotUseTLS)
	vpr.Set(DisableTLSVerificationKey, cn.DoNotVerifyTLS)
}

// storage parsing errors
var (
	errMissingRequired = clues.New("missing required storage configuration")
)

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
