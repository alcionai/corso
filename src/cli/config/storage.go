package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/alcionai/clues"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func s3ConfigsFromViper(vpr *viper.Viper) (storage.S3Config, error) {
	var s3Config storage.S3Config

	providerType := vpr.GetString(StorageProviderTypeKey)
	if providerType != storage.ProviderS3.String() {
		return s3Config, clues.New("unsupported storage provider: " + providerType)
	}

	s3Config.Bucket = vpr.GetString(BucketNameKey)
	s3Config.Endpoint = vpr.GetString(EndpointKey)
	s3Config.Prefix = vpr.GetString(PrefixKey)
	s3Config.DoNotUseTLS = vpr.GetBool(DisableTLSKey)
	s3Config.DoNotVerifyTLS = vpr.GetBool(DisableTLSVerificationKey)

	return s3Config, nil
}

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func s3CredsFromViper(vpr *viper.Viper, s3Config storage.S3Config) (storage.S3Config, error) {
	s3Config.AccessKey = vpr.GetString(AccessKey)
	s3Config.SecretKey = vpr.GetString(SecretAccessKey)
	s3Config.SessionToken = vpr.GetString(SessionToken)

	return s3Config, nil
}

func s3Overrides(in map[string]string) map[string]string {
	return map[string]string{
		storage.Bucket:         in[storage.Bucket],
		storage.Endpoint:       in[storage.Endpoint],
		storage.Prefix:         in[storage.Prefix],
		storage.DoNotUseTLS:    in[storage.DoNotUseTLS],
		storage.DoNotVerifyTLS: in[storage.DoNotVerifyTLS],
		StorageProviderTypeKey: in[StorageProviderTypeKey],
	}
}

// configureStorage builds a complete storage configuration from a mix of
// viper properties and manual overrides.
func configureStorage(
	vpr *viper.Viper,
	readConfigFromViper bool,
	matchFromConfig bool,
	overrides map[string]string,
) (storage.Storage, error) {
	var (
		s3Cfg storage.S3Config
		store storage.Storage
		err   error
	)

	if readConfigFromViper {
		if b, ok := overrides[storage.Bucket]; ok {
			overrides[storage.Bucket] = common.NormalizeBucket(b)
		}

		if p, ok := overrides[storage.Prefix]; ok {
			overrides[storage.Prefix] = common.NormalizePrefix(p)
		}
	}

	if matchFromConfig {
		if s3Cfg, err = s3ConfigsFromViper(vpr); err != nil {
			return store, clues.Wrap(err, "reading s3 configs from corso config file")
		}

		if err := mustMatchConfig(vpr, s3Overrides(overrides)); err != nil {
			return store, clues.Wrap(err, "verifying s3 configs in corso config file")
		}
	}

	if s3Cfg, err = s3CredsFromViper(vpr, s3Cfg); err != nil {
		return store, clues.Wrap(err, "reading s3 configs from corso config file")
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
			return store, clues.Wrap(err, "validating aws credentials")
		}
	}

	s3Cfg = storage.S3Config{
		AWS:      aws,
		Bucket:   str.First(overrides[storage.Bucket], s3Cfg.Bucket),
		Endpoint: str.First(overrides[storage.Endpoint], s3Cfg.Endpoint, "s3.amazonaws.com"),
		Prefix:   str.First(overrides[storage.Prefix], s3Cfg.Prefix),
		DoNotUseTLS: str.ParseBool(str.First(
			overrides[storage.DoNotUseTLS],
			strconv.FormatBool(s3Cfg.DoNotUseTLS),
			"false",
		)),
		DoNotVerifyTLS: str.ParseBool(str.First(
			overrides[storage.DoNotVerifyTLS],
			strconv.FormatBool(s3Cfg.DoNotVerifyTLS),
			"false",
		)),
	}

	// compose the common config and credentials
	corso := GetAndInsertCorso(vpr.GetString(CorsoPassphrase))
	if err := corso.Validate(); err != nil {
		return store, clues.Wrap(err, "validating corso credentials")
	}

	cCfg := storage.CommonConfig{
		Corso: corso,
	}
	// the following is a hack purely for integration testing.
	// the value is not required, and if empty, kopia will default
	// to its routine behavior
	if t, ok := vpr.Get("corso-testing").(bool); t && ok {
		dir, _ := filepath.Split(vpr.ConfigFileUsed())
		cCfg.KopiaCfgDir = dir
	}

	// ensure required properties are present
	if err := requireProps(map[string]string{
		storage.Bucket:              s3Cfg.Bucket,
		credentials.CorsoPassphrase: corso.CorsoPassphrase,
	}); err != nil {
		return storage.Storage{}, err
	}

	// build the storage
	store, err = storage.NewStorage(storage.ProviderS3, s3Cfg, cCfg)
	if err != nil {
		return store, clues.Wrap(err, "configuring repository storage")
	}

	return store, nil
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetAndInsertCorso(passphase string) credentials.Corso {
	// fetch data from flag, env var or func param giving priority to func param
	// Func param generally will be value fetched from config file using viper.
	corsoPassph := str.First(flags.CorsoPassphraseFV, os.Getenv(credentials.CorsoPassphrase), passphase)

	return credentials.Corso{
		CorsoPassphrase: corsoPassph,
	}
}
