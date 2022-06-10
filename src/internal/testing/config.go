package testing

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	// S3 config
	testCfgBucket   = "bucket"
	testCfgEndpoint = "endpoint"
	testCfgPrefix   = "prefix"

	// M365 config
	testCfgTenantID = "tenantid"
)

func newTestViper() (*viper.Viper, error) {
	configFilePath := os.Getenv("CORSO_TEST_CONFIG_FILE")

	vpr := viper.New()

	// Read from the default location
	if configFilePath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		// Search config in home directory with name ".corso" (without extension).
		vpr.AddConfigPath(home)
		vpr.SetConfigType("toml")
		vpr.SetConfigName(".corso_test")
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
	fallbackTo(testEnv, testCfgTenantID, vpr.GetString(testCfgTenantID))

	return testEnv, nil
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
