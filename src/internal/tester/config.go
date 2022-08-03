package tester

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/pkg/account"
)

const (
	// S3 config
	testCfgBucket   = "bucket"
	testCfgEndpoint = "endpoint"
	testCfgPrefix   = "prefix"

	// M365 config
	testCfgTenantID = "tenantid"
	testCfgUserID   = "m365userid"
)

// test specific env vars
const (
	EnvCorsoM365TestUserID = "CORSO_M356_TEST_USER_ID"
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

func newTestViper() (*viper.Viper, error) {
	vpr := viper.New()

	configFilePath := os.Getenv("CORSO_TEST_CONFIG_FILE")
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

	vpr, err := newTestViper()
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
	fallbackTo(testEnv, testCfgBucket, vpr.GetString(testCfgBucket), "test-corso-repo-init")
	fallbackTo(testEnv, testCfgEndpoint, vpr.GetString(testCfgEndpoint), "s3.amazonaws.com")
	fallbackTo(testEnv, testCfgPrefix, vpr.GetString(testCfgPrefix))
	fallbackTo(testEnv, testCfgTenantID, os.Getenv(account.TenantID), vpr.GetString(testCfgTenantID))
	fallbackTo(testEnv, testCfgUserID, os.Getenv(EnvCorsoM365TestUserID), vpr.GetString(testCfgTenantID), "lidiah@8qzvrj.onmicrosoft.com")

	testConfig = testEnv
	return cloneTestConfig(), nil
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
