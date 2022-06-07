package config

import (
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	// S3 config
	ProviderTypeKey = "provider"
	BucketNameKey   = "bucket"
	EndpointKey     = "endpoint"
	PrefixKey       = "prefix"

	// M365 config
	TenantIDKey = "tenantid"
)

// WriteRepoConfig currently just persists corso config to the config file
// It does not check for conflicts or existing data.
func WriteRepoConfig(s3Config storage.S3Config, account repository.Account) error {
	// Rudimentary support for persisting repo config
	// TODO: Handle conflicts, support other config types
	viper.Set(ProviderTypeKey, storage.ProviderS3.String())
	viper.Set(BucketNameKey, s3Config.Bucket)
	viper.Set(EndpointKey, s3Config.Endpoint)
	viper.Set(PrefixKey, s3Config.Prefix)
	viper.Set(TenantIDKey, account.TenantID)

	if err := viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return viper.GetViper().WriteConfig()
		}
		return err
	}
	return nil
}

func ReadRepoConfig() (s3Config storage.S3Config, account repository.Account, err error) {
	if err = viper.ReadInConfig(); err != nil {
		return s3Config, account, errors.Wrap(err, "reading config file: "+viper.ConfigFileUsed())
	}

	if providerType := viper.GetString(ProviderTypeKey); providerType != storage.ProviderS3.String() {
		return s3Config, account, errors.New("Unsupported storage provider: " + providerType)
	}

	s3Config.Bucket = viper.GetString(BucketNameKey)
	s3Config.Endpoint = viper.GetString(EndpointKey)
	s3Config.Prefix = viper.GetString(PrefixKey)
	account.TenantID = viper.GetString(TenantIDKey)

	return s3Config, account, nil
}
