package config

import (
	"os"
	"path"
	"strings"

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

func InitConfig(configFilePath string) error {
	// Configure default config file location
	if configFilePath == "" {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		// Search config in home directory with name ".corso" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".corso")
		return nil
	}
	// Use a custom file location

	viper.SetConfigFile(configFilePath)
	// We also configure the path, type and filename
	// because `viper.SafeWriteConfig` needs these set to
	// work correctly (it does not use the configured file)
	viper.AddConfigPath(path.Dir(configFilePath))

	fileName := path.Base(configFilePath)
	ext := path.Ext(configFilePath)
	if len(ext) == 0 {
		return errors.New("config file requires an extension e.g. `toml`")
	}
	fileName = strings.TrimSuffix(fileName, ext)
	viper.SetConfigType(ext[1:])
	viper.SetConfigName(fileName)
	return nil
}

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
