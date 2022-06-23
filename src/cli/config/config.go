package config

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
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
func WriteRepoConfig(s3Config storage.S3Config, account account.M365Config) error {
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

func ReadRepoConfig() (storage.S3Config, account.Account, error) {
	var (
		s3Config storage.S3Config
		acct     account.Account
		err      error
	)

	if err = viper.ReadInConfig(); err != nil {
		return s3Config, acct, errors.Wrap(err, "reading config file: "+viper.ConfigFileUsed())
	}

	if providerType := viper.GetString(ProviderTypeKey); providerType != storage.ProviderS3.String() {
		return s3Config, acct, errors.New("Unsupported storage provider: " + providerType)
	}

	s3Config.Bucket = viper.GetString(BucketNameKey)
	s3Config.Endpoint = viper.GetString(EndpointKey)
	s3Config.Prefix = viper.GetString(PrefixKey)

	m365Creds := credentials.GetM365()
	tenantID := os.Getenv(account.TenantID)
	cfgTenantID := viper.GetString(TenantIDKey)
	if len(tenantID) == 0 || len(cfgTenantID) > 0 {
		tenantID = cfgTenantID
	}
	acct, err = account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365:     m365Creds,
			TenantID: tenantID,
		},
	)

	return s3Config, acct, err
}

// GetStorageAndAccount creates a storage and account instance by mediating all the possible
// data sources (config file, env vars, flag overrides) and the config file.
func GetStorageAndAccount(readFromFile bool, overrides map[string]string) (storage.Storage, account.Account, error) {
	var (
		s3Cfg storage.S3Config
		acct  account.Account
		err   error
	)

	// possibly read the prior config from a .corso file
	if readFromFile {
		s3Cfg, acct, err = ReadRepoConfig()
		if err != nil {
			return storage.Storage{}, acct, errors.Wrap(err, "reading corso config file")
		}
	}

	// compose the s3 storage config and credentials
	aws := credentials.GetAWS(overrides)
	if err := aws.Validate(); err != nil {
		return storage.Storage{}, acct, errors.Wrap(err, "validating aws credentials")
	}
	s3Cfg = storage.S3Config{
		AWS:      aws,
		Bucket:   first(overrides[storage.Bucket], s3Cfg.Bucket),
		Endpoint: first(overrides[storage.Endpoint], s3Cfg.Endpoint),
		Prefix:   first(overrides[storage.Prefix], s3Cfg.Prefix),
	}

	// compose the common config and credentials
	corso := credentials.GetCorso()
	if err := corso.Validate(); err != nil {
		return storage.Storage{}, acct, errors.Wrap(err, "validating corso credentials")
	}
	cCfg := storage.CommonConfig{
		Corso: corso,
	}

	// ensure requried properties are present
	if err := utils.RequireProps(map[string]string{
		credentials.AWSAccessKeyID:     aws.AccessKey,
		storage.Bucket:                 s3Cfg.Bucket,
		credentials.AWSSecretAccessKey: aws.SecretKey,
		credentials.AWSSessionToken:    aws.SessionToken,
		credentials.CorsoPassword:      corso.CorsoPassword,
	}); err != nil {
		return storage.Storage{}, acct, err
	}

	// return a complete storage
	s, err := storage.NewStorage(storage.ProviderS3, s3Cfg, cCfg)
	if err != nil {
		return storage.Storage{}, acct, errors.Wrap(err, "configuring repository storage")
	}
	return s, acct, nil
}

// returns the first non-zero valued string
func first(vs ...string) string {
	for _, v := range vs {
		if len(v) > 0 {
			return v
		}
	}
	return ""
}
