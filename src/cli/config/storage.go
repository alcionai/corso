package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
)

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func s3ConfigsFromViper(vpr *viper.Viper) (storage.S3Config, error) {
	var s3Config storage.S3Config

	providerType := vpr.GetString(StorageProviderTypeKey)
	if providerType != storage.ProviderS3.String() {
		return s3Config, errors.New("unsupported storage provider: " + providerType)
	}

	s3Config.Bucket = vpr.GetString(BucketNameKey)
	s3Config.Endpoint = vpr.GetString(EndpointKey)
	s3Config.Prefix = vpr.GetString(PrefixKey)
	return s3Config, nil
}

func s3Overrides(in map[string]string) map[string]string {
	return map[string]string{
		storage.Bucket:         in[storage.Bucket],
		storage.Endpoint:       in[storage.Endpoint],
		storage.Prefix:         in[storage.Prefix],
		StorageProviderTypeKey: in[StorageProviderTypeKey],
	}
}

// configureStorage builds a complete storage configuration from a mix of
// viper properties and manual overrides.
func configureStorage(
	vpr *viper.Viper,
	readConfigFromViper bool,
	overrides map[string]string,
) (storage.Storage, error) {
	var (
		s3Cfg storage.S3Config
		store storage.Storage
		err   error
	)

	if readConfigFromViper {
		if s3Cfg, err = s3ConfigsFromViper(vpr); err != nil {
			return store, errors.Wrap(err, "reading s3 configs from corso config file")
		}

		if err := mustMatchConfig(vpr, s3Overrides(overrides)); err != nil {
			return store, errors.Wrap(err, "verifying s3 configs in corso config file")
		}
	}

	// compose the s3 storage config and credentials
	aws := credentials.GetAWS(overrides)
	if err := aws.Validate(); err != nil {
		return store, errors.Wrap(err, "validating aws credentials")
	}

	s3Cfg = storage.S3Config{
		AWS:      aws,
		Bucket:   common.First(overrides[storage.Bucket], s3Cfg.Bucket, os.Getenv(storage.BucketKey)),
		Endpoint: common.First(overrides[storage.Endpoint], s3Cfg.Endpoint, os.Getenv(storage.EndpointKey)),
		Prefix:   common.First(overrides[storage.Prefix], s3Cfg.Prefix, os.Getenv(storage.PrefixKey)),
	}

	// compose the common config and credentials
	corso := credentials.GetCorso()
	if err := corso.Validate(); err != nil {
		return store, errors.Wrap(err, "validating corso credentials")
	}
	cCfg := storage.CommonConfig{
		Corso: corso,
	}
	// the following is a hack purely for integration testing.
	// the value is not required, and if empty, kopia will default
	// to its routine behavior
	if vpr.Get("corso-testing") == true {
		dir, _ := filepath.Split(vpr.ConfigFileUsed())
		cCfg.KopiaCfgDir = dir
	}

	// ensure required properties are present
	if err := utils.RequireProps(map[string]string{
		credentials.AWSAccessKeyID:     aws.AccessKey,
		storage.Bucket:                 s3Cfg.Bucket,
		credentials.AWSSecretAccessKey: aws.SecretKey,
		credentials.AWSSessionToken:    aws.SessionToken,
		credentials.CorsoPassword:      corso.CorsoPassword,
	}); err != nil {
		return storage.Storage{}, err
	}

	// build the storage
	store, err = storage.NewStorage(storage.ProviderS3, s3Cfg, cCfg)
	if err != nil {
		return store, errors.Wrap(err, "configuring repository storage")
	}

	return store, nil
}
