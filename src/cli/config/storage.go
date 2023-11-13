package config

import (
	"context"
	"os"
	"path/filepath"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

// configureStorage builds a complete storage configuration from a mix of
// viper properties and manual overrides.
func configureStorage(
	vpr *viper.Viper,
	provider storage.ProviderType,
	readConfigFromViper bool,
	matchFromConfig bool,
	overrides map[string]string,
) (storage.Storage, error) {
	var store storage.Storage

	sc, err := storage.NewStorageConfig(provider)
	if err != nil {
		return store, clues.Stack(err)
	}

	err = sc.ApplyConfigOverrides(
		vpr,
		readConfigFromViper,
		matchFromConfig,
		overrides)
	if err != nil {
		return store, clues.Stack(err)
	}

	// compose the common config and credentials
	corso := GetAndInsertCorso(vpr.GetString(CorsoPassphrase))
	if err := corso.Validate(); err != nil {
		return store, clues.Wrap(err, "validating corso credentials")
	}

	configDir, _ := filepath.Split(vpr.ConfigFileUsed())

	cCfg := storage.CommonConfig{
		Corso:       corso,
		KopiaCfgDir: configDir,
	}
	// the following is a hack purely for integration testing.
	// the value is not required, and if empty, kopia will default
	// to its routine behavior
	if t, ok := vpr.Get("corso-testing").(bool); t && ok {
		cCfg.KopiaCfgDir = configDir
	}

	// ensure required properties are present
	if err := requireProps(map[string]string{
		credentials.CorsoPassphrase: corso.CorsoPassphrase,
	}); err != nil {
		return storage.Storage{}, err
	}

	// build the storage
	store, err = storage.NewStorage(provider, sc, cCfg)
	if err != nil {
		return store, clues.Wrap(err, "configuring repository storage")
	}

	return store, nil
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetAndInsertCorso(passphase string) credentials.Corso {
	// fetch data from flag, env var or func param giving priority to func param
	// Func param generally will be value fetched from config file using viper.
	corsoPassph := str.First(flags.PassphraseFV, os.Getenv(credentials.CorsoPassphrase), passphase)

	return credentials.Corso{
		CorsoPassphrase: corsoPassph,
	}
}

// GetStorageProviderFromConfigFile reads the storage provider from the config file.
// Storage provider can only be sourced from config file with the exception of
// commands that create or connect to a repo.
func GetStorageProviderFromConfigFile(ctx context.Context) (storage.ProviderType, error) {
	vpr := GetViper(ctx)

	err := vpr.ReadInConfig()
	if err != nil {
		return storage.ProviderUnknown, clues.Wrap(err, "reading config file")
	}

	provider := vpr.GetString(storage.StorageProviderTypeKey)

	return storage.StringToProviderType[provider], nil
}
