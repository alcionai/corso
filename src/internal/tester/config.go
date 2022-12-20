package tester

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/pkg/account"
)

const (
	// S3 config
	TestCfgBucket          = "bucket"
	TestCfgEndpoint        = "endpoint"
	TestCfgPrefix          = "prefix"
	TestCfgStorageProvider = "provider"

	// M365 config
	TestCfgAzureTenantID    = "azure_tenantid"
	TestCfgSiteID           = "m365siteid"
	TestCfgUserID           = "m365userid"
	TestCfgSecondaryUserID  = "secondarym365userid"
	TestCfgLoadTestUserID   = "loadtestm365userid"
	TestCfgLoadTestOrgUsers = "loadtestm365orgusers"
	TestCfgAccountProvider  = "account_provider"
)

// test specific env vars
const (
	EnvCorsoM365TestSiteID          = "CORSO_M365_TEST_SITE_ID"
	EnvCorsoM365TestUserID          = "CORSO_M365_TEST_USER_ID"
	EnvCorsoSecondaryM365TestUserID = "CORSO_SECONDARY_M365_TEST_USER_ID"
	EnvCorsoM365LoadTestUserID      = "CORSO_M365_LOAD_TEST_USER_ID"
	EnvCorsoM365LoadTestOrgUsers    = "CORSO_M365_LOAD_TEST_ORG_USERS"
	EnvCorsoTestConfigFilePath      = "CORSO_TEST_CONFIG_FILE"
)

// global to hold the test config results.
var testConfig map[string]string

// call this instead of returning testConfig directly.
func cloneTestConfig() map[string]string {
	if testConfig == nil {
		return map[string]string{}
	}

	return maps.Clone(testConfig)
}

func NewTestViper() (*viper.Viper, error) {
	vpr := viper.New()

	configFilePath := os.Getenv(EnvCorsoTestConfigFilePath)
	if len(configFilePath) == 0 {
		return vpr, nil
	}

	// Or use a custom file location
	ext := filepath.Ext(configFilePath)
	if len(ext) == 0 {
		return nil, errors.New("corso_test requires an extension")
	}

	vpr.SetConfigFile(configFilePath)
	vpr.AddConfigPath(filepath.Dir(configFilePath))
	vpr.SetConfigType(ext[1:])

	fileName := filepath.Base(configFilePath)
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
	fallbackTo(testEnv, TestCfgAzureTenantID, os.Getenv(account.AzureTenantID), vpr.GetString(TestCfgAzureTenantID))
	fallbackTo(
		testEnv,
		TestCfgUserID,
		os.Getenv(EnvCorsoM365TestUserID),
		vpr.GetString(TestCfgUserID),
		"lynner@8qzvrj.onmicrosoft.com",
		//"lidiah@8qzvrj.onmicrosoft.com",
	)
	fallbackTo(
		testEnv,
		TestCfgSecondaryUserID,
		os.Getenv(EnvCorsoSecondaryM365TestUserID),
		vpr.GetString(TestCfgSecondaryUserID),
		"lidiah@8qzvrj.onmicrosoft.com",
		//"lynner@8qzvrj.onmicrosoft.com",
	)
	fallbackTo(
		testEnv,
		TestCfgLoadTestUserID,
		os.Getenv(EnvCorsoM365LoadTestUserID),
		vpr.GetString(TestCfgLoadTestUserID),
		"leeg@8qzvrj.onmicrosoft.com",
	)
	fallbackTo(
		testEnv,
		TestCfgLoadTestOrgUsers,
		os.Getenv(EnvCorsoM365LoadTestOrgUsers),
		vpr.GetString(TestCfgLoadTestOrgUsers),
		"lidiah@8qzvrj.onmicrosoft.com,lynner@8qzvrj.onmicrosoft.com",
	)
	fallbackTo(
		testEnv,
		TestCfgSiteID,
		os.Getenv(EnvCorsoM365TestSiteID),
		vpr.GetString(TestCfgSiteID),
		"8qzvrj.sharepoint.com,1c9ef309-f47c-4e69-832b-a83edd69fa7f,c57f6e0e-3e4b-472c-b528-b56a2ccd0507",
	)

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
// Returns a filepath string pointing to the location of the temp file.
func MakeTempTestConfigClone(t *testing.T, overrides map[string]string) (*viper.Viper, string) {
	cfg, err := readTestConfig()
	require.NoError(t, err, "reading tester config")

	fName := filepath.Base(os.Getenv(EnvCorsoTestConfigFilePath))
	if len(fName) == 0 || fName == "." || fName == "/" {
		fName = ".corso_test.toml"
	}

	tDir := t.TempDir()
	tDirFp := filepath.Join(tDir, fName)

	_, err = os.Create(tDirFp)
	require.NoError(t, err, "creating temp test dir")

	ext := filepath.Ext(fName)
	vpr := viper.New()
	vpr.SetConfigFile(tDirFp)
	vpr.AddConfigPath(tDir)
	vpr.SetConfigType(strings.TrimPrefix(ext, "."))
	vpr.SetConfigName(strings.TrimSuffix(fName, ext))
	vpr.Set("corso-testing", true)

	for k, v := range cfg {
		vpr.Set(k, v)
	}

	for k, v := range overrides {
		vpr.Set(k, v)
	}

	require.NoError(t, vpr.WriteConfig(), "writing temp dir viper config file")

	return vpr, tDirFp
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
