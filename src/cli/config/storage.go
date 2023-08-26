package config

import (
	"fmt"
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
	readConfigFromViper bool,
	matchFromConfig bool,
	overrides map[string]string,
) (storage.Storage, error) {
	var (
		store    storage.Storage
		err      error
		provider = storage.ProviderS3 // temporary
	)

	storageCfg, err := fetchStorageConfigFromViper(
		vpr,
		provider, // Make it generic
		readConfigFromViper,
		matchFromConfig,
		overrides)
	if err != nil {
		return store, err
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
		credentials.CorsoPassphrase: corso.CorsoPassphrase,
	}); err != nil {
		return storage.Storage{}, err
	}

	// build the storage
	store, err = storage.NewStorage(provider, storageCfg, cCfg)
	if err != nil {
		return store, clues.Wrap(err, "configuring repository storage")
	}

	return store, nil
}

func fetchStorageConfigFromViper(
	vpr *viper.Viper,
	p storage.StorageProvider,
	readConfigFromViper bool,
	matchFromConfig bool,
	overrides map[string]string,
) (StorageConfigurer, error) {
	switch p {
	case storage.ProviderS3:
		return fetchS3ConfigFromViper(
			vpr,
			readConfigFromViper,
			matchFromConfig,
			overrides)
	}

	return nil, clues.New("unsupported storage provider")
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

// Helper for parsing the values in a config object.
// If the value is nil or not a string, returns an empty string.
func orEmptyString(v any) string {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("panic recovery casting %v to string\n", v)
		}
	}()

	if v == nil {
		return ""
	}

	return v.(string)
}
