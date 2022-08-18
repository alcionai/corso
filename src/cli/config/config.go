package config

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/storage"
)

const (
	// S3 config
	StorageProviderTypeKey = "provider"
	BucketNameKey          = "bucket"
	EndpointKey            = "endpoint"
	PrefixKey              = "prefix"

	// M365 config
	AccountProviderTypeKey = "account_provider"
	TenantIDKey            = "tenantid"
)

var configFilePath string

// adds the persistent flag --config-file to the provided command.
func AddConfigFileFlag(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		Err(cmd.Context(), "finding $HOME directory (default) for config file")
	}
	fs.StringVar(
		&configFilePath,
		"config-file",
		filepath.Join(homeDir, ".corso.toml"),
		"config file (default is $HOME/.corso)")
}

// ---------------------------------------------------------------------------------------------------------
// Initialization & Storage
// ---------------------------------------------------------------------------------------------------------

// InitFunc provides a func that lazily initializes viper and
// verifies that the configuration was able to read a file.
func InitFunc() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := initWithViper(GetViper(cmd.Context()), configFilePath)
		if err != nil {
			return err
		}
		return Read(cmd.Context())
	}
}

// initWithViper implements InitConfig, but takes in a viper
// struct for testing.
func initWithViper(vpr *viper.Viper, configFP string) error {
	// Configure default config file location
	if configFP == "" {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		// Search config in home directory with name ".corso" (without extension).
		vpr.AddConfigPath(home)
		vpr.SetConfigType("toml")
		vpr.SetConfigName(".corso")
		return nil
	}

	vpr.SetConfigFile(configFP)
	// We also configure the path, type and filename
	// because `vpr.SafeWriteConfig` needs these set to
	// work correctly (it does not use the configured file)
	vpr.AddConfigPath(path.Dir(configFP))

	fileName := path.Base(configFP)
	ext := path.Ext(configFP)
	if len(ext) == 0 {
		return errors.New("config file requires an extension e.g. `toml`")
	}
	fileName = strings.TrimSuffix(fileName, ext)
	vpr.SetConfigType(strings.TrimPrefix(ext, "."))
	vpr.SetConfigName(fileName)

	return nil
}

type viperCtx struct{}

// Seed embeds a viper instance in the context.
func Seed(ctx context.Context) context.Context {
	return SetViper(ctx, nil)
}

// Adds a viper instance to the context.
// If vpr is nil, sets the default (global) viper.
func SetViper(ctx context.Context, vpr *viper.Viper) context.Context {
	if vpr == nil {
		vpr = viper.GetViper()
	}
	return context.WithValue(ctx, viperCtx{}, vpr)
}

// Gets a viper instance from the context.
// If no viper instance is found, returns the default
// (global) viper instance.
func GetViper(ctx context.Context) *viper.Viper {
	vprIface := ctx.Value(viperCtx{})
	vpr, ok := vprIface.(*viper.Viper)
	if vpr == nil || !ok {
		return viper.GetViper()
	}
	return vpr
}

// ---------------------------------------------------------------------------------------------------------
// Reading & Writing the config
// ---------------------------------------------------------------------------------------------------------

// Read reads the config from the viper instance in the context.
// Primarily used as a test-check to ensure the instance was
// set up properly.
func Read(ctx context.Context) error {
	if err := viper.ReadInConfig(); err == nil {
		logger.Ctx(ctx).Debugw("found config file", "configFile", viper.ConfigFileUsed())
		return err
	}
	return nil
}

// WriteRepoConfig currently just persists corso config to the config file
// It does not check for conflicts or existing data.
func WriteRepoConfig(ctx context.Context, s3Config storage.S3Config, m365Config account.M365Config) error {
	return writeRepoConfigWithViper(GetViper(ctx), s3Config, m365Config)
}

// writeRepoConfigWithViper implements WriteRepoConfig, but takes in a viper
// struct for testing.
func writeRepoConfigWithViper(vpr *viper.Viper, s3Config storage.S3Config, m365Config account.M365Config) error {
	// Rudimentary support for persisting repo config
	// TODO: Handle conflicts, support other config types
	vpr.Set(StorageProviderTypeKey, storage.ProviderS3.String())
	vpr.Set(BucketNameKey, s3Config.Bucket)
	vpr.Set(EndpointKey, s3Config.Endpoint)
	vpr.Set(PrefixKey, s3Config.Prefix)

	vpr.Set(AccountProviderTypeKey, account.ProviderM365.String())
	vpr.Set(TenantIDKey, m365Config.TenantID)

	if err := vpr.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return vpr.WriteConfig()
		}
		return err
	}
	return nil
}

// GetStorageAndAccount creates a storage and account instance by mediating all the possible
// data sources (config file, env vars, flag overrides) and the config file.
func GetStorageAndAccount(
	ctx context.Context,
	readFromFile bool,
	overrides map[string]string,
) (storage.Storage, account.Account, error) {
	return getStorageAndAccountWithViper(GetViper(ctx), readFromFile, overrides)
}

// getSorageAndAccountWithViper implements GetSorageAndAccount, but takes in a viper
// struct for testing.
func getStorageAndAccountWithViper(
	vpr *viper.Viper,
	readFromFile bool,
	overrides map[string]string,
) (storage.Storage, account.Account, error) {
	var (
		store storage.Storage
		acct  account.Account
		err   error
	)

	readConfigFromViper := readFromFile

	// possibly read the prior config from a .corso file
	if readFromFile {
		err = vpr.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return store, acct, errors.Wrap(err, "reading corso config file: "+vpr.ConfigFileUsed())
			}
			readConfigFromViper = false
		}
	}

	acct, err = configureAccount(vpr, readConfigFromViper, overrides)
	if err != nil {
		return store, acct, errors.Wrap(err, "retrieving account configuration details")
	}

	store, err = configureStorage(vpr, readConfigFromViper, overrides)
	if err != nil {
		return store, acct, errors.Wrap(err, "retrieving storage provider details")
	}

	return store, acct, nil
}

// ---------------------------------------------------------------------------
// Helper funcs
// ---------------------------------------------------------------------------

var constToTomlKeyMap = map[string]string{
	account.TenantID:       TenantIDKey,
	AccountProviderTypeKey: AccountProviderTypeKey,
	storage.Bucket:         BucketNameKey,
	storage.Endpoint:       EndpointKey,
	storage.Prefix:         PrefixKey,
	StorageProviderTypeKey: StorageProviderTypeKey,
}

// mustMatchConfig compares the values of each key to their config file value in viper.
// If any value differs from the viper value, an error is returned.
// values in m that aren't stored in the config are ignored.
func mustMatchConfig(vpr *viper.Viper, m map[string]string) error {
	for k, v := range m {
		if len(v) == 0 {
			continue // empty variables will get caught by configuration validators, if necessary
		}
		tomlK, ok := constToTomlKeyMap[k]
		if !ok {
			continue // m may declare values which aren't stored in the config file
		}
		vv := vpr.GetString(tomlK)
		if v != vv {
			return errors.New("value of " + k + " (" + v + ") does not match corso configuration value (" + vv + ")")
		}
	}
	return nil
}
