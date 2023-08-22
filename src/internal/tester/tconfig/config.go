package tconfig

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alcionai/clues"
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
	TestCfgSecondarySiteID  = "secondarym365siteid"
	TestCfgSiteID           = "m365siteid"
	TestCfgSiteURL          = "m365siteurl"
	TestCfgTeamID           = "m365teamid"
	TestCfgGroupID          = "m365groupid"
	TestCfgUserID           = "m365userid"
	TestCfgSecondaryUserID  = "secondarym365userid"
	TestCfgSecondaryGroupID = "secondarym365groupid"
	TestCfgTertiaryUserID   = "tertiarym365userid"
	TestCfgLoadTestUserID   = "loadtestm365userid"
	TestCfgLoadTestOrgUsers = "loadtestm365orgusers"
	TestCfgAccountProvider  = "account_provider"
	TestCfgUnlicensedUserID = "unlicensedm365userid"
)

// test specific env vars
const (
	EnvCorsoM365LoadTestUserID       = "CORSO_M365_LOAD_TEST_USER_ID"
	EnvCorsoM365LoadTestOrgUsers     = "CORSO_M365_LOAD_TEST_ORG_USERS"
	EnvCorsoM365TestSiteID           = "CORSO_M365_TEST_SITE_ID"
	EnvCorsoM365TestSiteURL          = "CORSO_M365_TEST_SITE_URL"
	EnvCorsoM365TestTeamID           = "CORSO_M365_TEST_TEAM_ID"
	EnvCorsoM365TestGroupID          = "CORSO_M365_TEST_GROUP_ID"
	EnvCorsoM365TestUserID           = "CORSO_M365_TEST_USER_ID"
	EnvCorsoSecondaryM365TestSiteID  = "CORSO_SECONDARY_M365_TEST_SITE_ID"
	EnvCorsoSecondaryM365TestUserID  = "CORSO_SECONDARY_M365_TEST_USER_ID"
	EnvCorsoTertiaryM365TestUserID   = "CORSO_TERTIARY_M365_TEST_USER_ID"
	EnvCorsoTestConfigFilePath       = "CORSO_TEST_CONFIG_FILE"
	EnvCorsoUnlicensedM365TestUserID = "CORSO_M365_TEST_UNLICENSED_USER"
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
		return nil, clues.New("corso_test requires an extension")
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
func ReadTestConfig() (map[string]string, error) {
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
			return nil, clues.Wrap(err, "reading config file: "+viper.ConfigFileUsed())
		}
	}

	testEnv := map[string]string{}
	fallbackTo(testEnv, TestCfgStorageProvider, vpr.GetString(TestCfgStorageProvider))
	fallbackTo(testEnv, TestCfgAccountProvider, vpr.GetString(TestCfgAccountProvider))
	fallbackTo(testEnv, TestCfgBucket, os.Getenv("S3_BUCKET"), vpr.GetString(TestCfgBucket))
	fallbackTo(testEnv, TestCfgEndpoint, vpr.GetString(TestCfgEndpoint), "s3.amazonaws.com")
	fallbackTo(testEnv, TestCfgPrefix, vpr.GetString(TestCfgPrefix))
	fallbackTo(testEnv, TestCfgAzureTenantID, os.Getenv(account.AzureTenantID), vpr.GetString(TestCfgAzureTenantID))
	fallbackTo(
		testEnv,
		TestCfgUserID,
		os.Getenv(EnvCorsoM365TestUserID),
		vpr.GetString(TestCfgUserID),
		"LynneR@10rqc2.onmicrosoft.com")
	fallbackTo(
		testEnv,
		TestCfgSecondaryUserID,
		os.Getenv(EnvCorsoSecondaryM365TestUserID),
		vpr.GetString(TestCfgSecondaryUserID),
		"AdeleV@10rqc2.onmicrosoft.com")
	fallbackTo(
		testEnv,
		TestCfgTertiaryUserID,
		os.Getenv(EnvCorsoTertiaryM365TestUserID),
		vpr.GetString(TestCfgTertiaryUserID),
		"PradeepG@10rqc2.onmicrosoft.com")
	fallbackTo(
		testEnv,
		TestCfgLoadTestUserID,
		os.Getenv(EnvCorsoM365LoadTestUserID),
		vpr.GetString(TestCfgLoadTestUserID),
		"leeg@10rqc2.onmicrosoft.com")
	fallbackTo(
		testEnv,
		TestCfgLoadTestOrgUsers,
		os.Getenv(EnvCorsoM365LoadTestOrgUsers),
		vpr.GetString(TestCfgLoadTestOrgUsers),
		"AdeleV@10rqc2.onmicrosoft.com,LynneR@10rqc2.onmicrosoft.com")
	fallbackTo(
		testEnv,
		TestCfgSiteID,
		os.Getenv(EnvCorsoM365TestSiteID),
		vpr.GetString(TestCfgSiteID),
		"4892edf5-2ebf-46be-a6e5-a40b2cbf1c1a,38ab6d06-fc82-4417-af93-22d8733c22be")
	fallbackTo(
		testEnv,
		TestCfgTeamID,
		os.Getenv(EnvCorsoM365TestTeamID),
		vpr.GetString(TestCfgTeamID),
		"6f24b40d-b13d-4752-980f-f5fb9fba7aa0")
	fallbackTo(
		testEnv,
		TestCfgGroupID,
		os.Getenv(EnvCorsoM365TestGroupID),
		vpr.GetString(TestCfgGroupID),
		"6f24b40d-b13d-4752-980f-f5fb9fba7aa0")
	fallbackTo(
		testEnv,
		TestCfgSiteURL,
		os.Getenv(EnvCorsoM365TestSiteURL),
		vpr.GetString(TestCfgSiteURL),
		"https://10rqc2.sharepoint.com/sites/CorsoCI")
	fallbackTo(
		testEnv,
		TestCfgSecondarySiteID,
		os.Getenv(EnvCorsoSecondaryM365TestSiteID),
		vpr.GetString(TestCfgSecondarySiteID),
		"053684d8-ca6c-4376-a03e-2567816bb091,9b3e9abe-6a5e-4084-8b44-ea5a356fe02c")
	fallbackTo(
		testEnv,
		TestCfgUnlicensedUserID,
		os.Getenv(EnvCorsoUnlicensedM365TestUserID),
		vpr.GetString(TestCfgUnlicensedUserID),
		"testevents@10rqc2.onmicrosoft.com")

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
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "reading tester config", clues.ToCore(err))

	fName := filepath.Base(os.Getenv(EnvCorsoTestConfigFilePath))
	if len(fName) == 0 || fName == "." || fName == "/" {
		fName = ".corso_test.toml"
	}

	tDir := t.TempDir()
	tDirFp := filepath.Join(tDir, fName)

	_, err = os.Create(tDirFp)
	require.NoError(t, err, "creating temp test dir", clues.ToCore(err))

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

	err = vpr.WriteConfig()
	require.NoError(t, err, "writing temp dir viper config file", clues.ToCore(err))

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
