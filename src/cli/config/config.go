package config

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/storage"
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

	// M365 config
	AccountProviderTypeKey = "account_provider"
	AzureTenantIDKey       = "azure_tenantid"
	AzureClientID          = "azure_client_id"
	AzureSecret            = "azure_secret"

	// Corso passphrase in config
	CorsoPassphrase = "passphrase"
	CorsoUser       = "corso_user"
	CorsoHost       = "corso_host"
)

var (
	configFilePath     string
	configFilePathFlag string
	configDir          string
	displayDefaultFP   = filepath.Join("$HOME", ".corso.toml")
)

// RepoDetails holds the repository configuration retrieved from
// the .corso.toml configuration file.
type RepoDetails struct {
	Storage  storage.Storage
	Account  account.Account
	RepoID   string
	RepoUser string
	RepoHost string
}

// Attempts to set the default dir and config file path.
// Default is always $HOME.
func init() {
	envDir := os.Getenv("CORSO_CONFIG_DIR")
	if len(envDir) > 0 {
		if _, err := os.Stat(envDir); err != nil {
			Infof(context.Background(), "cannot stat CORSO_CONFIG_DIR [%s]: %v", envDir, err)
		} else {
			configDir = envDir
			configFilePath = filepath.Join(configDir, ".corso.toml")
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		Infof(context.Background(), "cannot stat user's $HOME directory: %v", err)
	}

	if len(configDir) == 0 {
		configDir = homeDir
		configFilePath = filepath.Join(configDir, ".corso.toml")
	}
}

// adds the persistent flag --config-file to the provided command.
func AddConfigFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVar(
		&configFilePathFlag,
		"config-file", displayDefaultFP, "config file location")
}

// ---------------------------------------------------------------------------------------------------------
// Initialization & Storage
// ---------------------------------------------------------------------------------------------------------

// InitFunc provides a func that lazily initializes viper and
// verifies that the configuration was able to read a file.
func InitFunc(cmd *cobra.Command, args []string) error {
	fp := configFilePathFlag
	if len(fp) == 0 || fp == displayDefaultFP {
		fp = configFilePath
	}

	err := initWithViper(GetViper(cmd.Context()), fp)
	if err != nil {
		return err
	}

	return Read(cmd.Context())
}

// initWithViper implements InitConfig, but takes in a viper
// struct for testing.
func initWithViper(vpr *viper.Viper, configFP string) error {
	// Configure default config file location
	if configFP == "" || configFP == displayDefaultFP {
		// Find home directory.
		_, err := os.Stat(configDir)
		if err != nil {
			return err
		}

		// Search config in home directory with name ".corso" (without extension).
		vpr.AddConfigPath(configDir)
		vpr.SetConfigType("toml")
		vpr.SetConfigName(".corso")

		return nil
	}

	vpr.SetConfigFile(configFP)
	// We also configure the path, type and filename
	// because `vpr.SafeWriteConfig` needs these set to
	// work correctly (it does not use the configured file)
	vpr.AddConfigPath(filepath.Dir(configFP))

	ext := filepath.Ext(configFP)
	if len(ext) == 0 {
		return clues.New("config file requires an extension e.g. `toml`")
	}

	fileName := filepath.Base(configFP)
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
func WriteRepoConfig(
	ctx context.Context,
	storageConfig StorageConfigurer,
	m365Config account.M365Config,
	repoOpts repository.Options,
	repoID string,
) error {
	return writeRepoConfigWithViper(
		GetViper(ctx),
		storageConfig,
		m365Config,
		repoOpts,
		repoID)
}

// writeRepoConfigWithViper implements WriteRepoConfig, but takes in a viper
// struct for testing.
func writeRepoConfigWithViper(
	vpr *viper.Viper,
	storageConfig StorageConfigurer,
	m365Config account.M365Config,
	repoOpts repository.Options,
	repoID string,
) error {
	storageConfig.WriteConfigToViper(vpr)

	// Rudimentary support for persisting repo config
	// TODO: Handle conflicts, support other config types

	vpr.Set(RepoID, repoID)

	// Need if-checks as Viper will write empty values otherwise.
	if len(repoOpts.User) > 0 {
		vpr.Set(CorsoUser, repoOpts.User)
	}

	if len(repoOpts.Host) > 0 {
		vpr.Set(CorsoHost, repoOpts.Host)
	}

	vpr.Set(AccountProviderTypeKey, account.ProviderM365.String())
	vpr.Set(AzureTenantIDKey, m365Config.AzureTenantID)

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
func GetConfigRepoDetails(
	ctx context.Context,
	readFromFile bool,
	mustMatchFromConfig bool,
	overrides map[string]string,
) (
	RepoDetails,
	error,
) {
	config, err := getStorageAndAccountWithViper(GetViper(ctx), readFromFile, mustMatchFromConfig, overrides)
	return config, err
}

// getSorageAndAccountWithViper implements GetSorageAndAccount, but takes in a viper
// struct for testing.
func getStorageAndAccountWithViper(
	vpr *viper.Viper,
	readFromFile bool,
	mustMatchFromConfig bool,
	overrides map[string]string,
) (
	RepoDetails,
	error,
) {
	var (
		config RepoDetails
		err    error
	)

	readConfigFromViper := readFromFile

	// possibly read the prior config from a .corso file
	if readFromFile {
		err = vpr.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return config, clues.Wrap(err, "reading corso config file: "+vpr.ConfigFileUsed())
			}

			readConfigFromViper = false
		}

		// in case of existing config, fetch repoid from config file
		config.RepoID = vpr.GetString(RepoID)
	}

	config.Account, err = configureAccount(vpr, readConfigFromViper, mustMatchFromConfig, overrides)
	if err != nil {
		return config, clues.Wrap(err, "retrieving account configuration details")
	}

	config.Storage, err = configureStorage(vpr, readConfigFromViper, mustMatchFromConfig, overrides)
	if err != nil {
		return config, clues.Wrap(err, "retrieving storage provider details")
	}

	config.RepoUser, config.RepoHost = getUserHost(vpr, readConfigFromViper)

	return config, nil
}

func getUserHost(vpr *viper.Viper, readConfigFromViper bool) (string, string) {
	user := str.First(flags.UserMaintenanceFV, vpr.GetString(CorsoUser))
	host := str.First(flags.HostnameMaintenanceFV, vpr.GetString(CorsoHost))

	// Fine if these are empty; later code will assign a meaningful default if
	// needed.
	return user, host
}

// ---------------------------------------------------------------------------
// Helper funcs
// ---------------------------------------------------------------------------

// TODO: This is not really needed?
var constToTomlKeyMap = map[string]string{
	account.AzureTenantID:  AzureTenantIDKey,
	AccountProviderTypeKey: AccountProviderTypeKey,
	Bucket:                 BucketNameKey,
	Endpoint:               EndpointKey,
	Prefix:                 PrefixKey,
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
			return clues.New("value of " + k + " (" + v + ") does not match corso configuration value (" + vv + ")")
		}
	}

	return nil
}

// requireProps validates the existence of the properties
// in the map.  Expects the format map[propName]propVal.
func requireProps(props map[string]string) error {
	for name, val := range props {
		if len(val) == 0 {
			return clues.New(name + " is required to perform this command")
		}
	}

	return nil
}
