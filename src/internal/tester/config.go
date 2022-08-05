package tester

import (
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/pkg/account"
)

const (
	// S3 config
	TestCfgBucket          = "bucket"
	TestCfgEndpoint        = "endpoint"
	TestCfgPrefix          = "prefix"
	TestCfgStorageProvider = "provider"

	// M365 config
	TestCfgTenantID        = "tenantid"
	TestCfgUserID          = "m365userid"
	TestCfgAccountProvider = "account_provider"
)

// test specific env vars
const (
	EnvCorsoM365TestUserID     = "CORSO_M356_TEST_USER_ID"
	EnvCorsoTestConfigFilePath = "CORSO_TEST_CONFIG_FILE"
)

// global to hold the test config results.
var testConfig map[string]string

// call this instead of returning testConfig directly.
func cloneTestConfig() map[string]string {
	if testConfig == nil {
		return map[string]string{}
	}
	clone := map[string]string{}
	for k, v := range testConfig {
		clone[k] = v
	}
	return clone
}

func NewTestViper() (*viper.Viper, error) {
	vpr := viper.New()

	configFilePath := os.Getenv(EnvCorsoTestConfigFilePath)
	if len(configFilePath) == 0 {
		return vpr, nil
	}

	// Or use a custom file location
	fileName := path.Base(configFilePath)
	ext := path.Ext(configFilePath)
	if len(ext) == 0 {
		return nil, errors.New("corso_test requires an extension")
	}

	vpr.SetConfigFile(configFilePath)
	vpr.AddConfigPath(path.Dir(configFilePath))
	vpr.SetConfigType(ext[1:])
	fileName = strings.TrimSuffix(fileName, ext)
	vpr.SetConfigName(fileName)

	return vpr, nil
}

// reads a corso configuration file with values specific to
// local integration test controls.  Populates values with
// defaults where standard.
func readTestConfig() (map[string]string, error) {
	if testConfig != nil {
		return cloneTestConfig(), nil
	}

	vpr, err := NewTestViper()
	if err != nil {
		return nil, err
	}

	// only error if reading an existing file failed.  No problem if we're missing files.
	if err = vpr.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return nil, errors.Wrap(err, "reading config file: "+viper.ConfigFileUsed())
		}
	}

	testEnv := map[string]string{}
	fallbackTo(testEnv, TestCfgStorageProvider, vpr.GetString(TestCfgStorageProvider))
	fallbackTo(testEnv, TestCfgAccountProvider, vpr.GetString(TestCfgAccountProvider))
	fallbackTo(testEnv, TestCfgBucket, vpr.GetString(TestCfgBucket), "test-corso-repo-init")
	fallbackTo(testEnv, TestCfgEndpoint, vpr.GetString(TestCfgEndpoint), "s3.amazonaws.com")
	fallbackTo(testEnv, TestCfgPrefix, vpr.GetString(TestCfgPrefix))
	fallbackTo(testEnv, TestCfgTenantID, os.Getenv(account.TenantID), vpr.GetString(TestCfgTenantID))
	fallbackTo(testEnv, TestCfgUserID, os.Getenv(EnvCorsoM365TestUserID), vpr.GetString(TestCfgTenantID), "lidiah@8qzvrj.onmicrosoft.com")
	testEnv[EnvCorsoTestConfigFilePath] = os.Getenv(EnvCorsoTestConfigFilePath)

	testConfig = testEnv
	return cloneTestConfig(), nil
}

// MakeTempTestConfigClone makes a copy of the test config file in a temp directory.
// This allows tests which interface with reading and writing to a config file
// (such as the CLI) to safely manipulate file contents without amending the user's
// original file.
//
// Attempts to copy values sourced from the caller's test config file.
// The overrides prop replaces config values with the provided value.
//
// Returns a context containing the new viper instance,
// and a filepath string pointing to the location of the temp file.
func MakeTempTestConfigClone(
	ctx context.Context,
	t *testing.T,
	overrides map[string]string,
) (context.Context, string, error) {
	cfg, err := readTestConfig()
	if err != nil {
		return ctx, "", err
	}

	fName := path.Base(os.Getenv(EnvCorsoTestConfigFilePath))
	if len(fName) == 0 || fName == "." || fName == "/" {
		fName = ".corso_test.toml"
	}
	tDir := t.TempDir()
	tDirFp := path.Join(tDir, fName)

	if _, err := os.Create(tDirFp); err != nil {
		return ctx, "", err
	}

	vpr := viper.New()
	ext := path.Ext(fName)
	vpr.SetConfigFile(tDirFp)
	vpr.AddConfigPath(tDir)
	vpr.SetConfigType(strings.TrimPrefix(ext, "."))
	vpr.SetConfigName(strings.TrimSuffix(fName, ext))

	for k, v := range cfg {
		vpr.Set(k, v)
	}

	for k, v := range overrides {
		vpr.Set(k, v)
	}

	if err := vpr.WriteConfig(); err != nil {
		return ctx, "", err
	}

	return config.SetViper(ctx, vpr), tDirFp, nil
}

// writes the first non-zero valued string to the map at the key.
// fallback priority should match viper ordering (manually handled
// here since viper fails to provide fallbacks on fileNotFoundErr):
// manual overrides > flags > env vars > config file > default value
func fallbackTo(m map[string]string, key string, fallbacks ...string) {
	for _, fb := range fallbacks {
		if len(fb) > 0 {
			m[key] = fb
			return
		}
	}
}
