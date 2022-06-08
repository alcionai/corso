package config_test

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	configFileTemplate = `
bucket = '%s'
endpoint = 's3.amazonaws.com'
prefix = 'test-prefix'
provider = 'S3'
tenantid = '%s'
`
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestReadRepoConfigBasic() {
	// Generate test config file
	b := "test-bucket"
	tID := "6f34ac30-8196-469b-bf8f-d83deadbbbba"
	testConfigData := fmt.Sprintf(configFileTemplate, b, tID)
	testConfigFilePath := path.Join(suite.T().TempDir(), "corso.toml")
	err := ioutil.WriteFile(testConfigFilePath, []byte(testConfigData), 0700)
	assert.NoError(suite.T(), err)

	// Configure viper to read test config file
	viper.SetConfigFile(testConfigFilePath)

	// Read and validate config
	s3Cfg, account, err := config.ReadRepoConfig()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), b, s3Cfg.Bucket)
	assert.Equal(suite.T(), tID, account.TenantID)
}

func (suite *ConfigSuite) TestWriteReadConfig() {
	// Configure viper to read test config file
	tempDir := suite.T().TempDir()
	testConfigFilePath := path.Join(tempDir, "corso.toml")
	err := config.InitConfig(testConfigFilePath)
	assert.NoError(suite.T(), err)

	table := []struct {
		name          string
		s3Cfg         storage.S3Config
		account       repository.Account
		writeErrCheck assert.ErrorAssertionFunc
		readErrCheck  assert.ErrorAssertionFunc
	}{
		{
			name:          "good",
			s3Cfg:         storage.S3Config{Bucket: "bucket"},
			account:       repository.Account{TenantID: "6f34ac30-8196-469b-bf8f-d83deadbbbbd"},
			writeErrCheck: assert.NoError,
			readErrCheck:  assert.NoError,
		},
	}
	for _, test := range table {
		err := config.WriteRepoConfig(test.s3Cfg, test.account)
		test.writeErrCheck(suite.T(), err)

		if err != nil {
			break
		}
		readS3Cfg, readAccount, err := config.ReadRepoConfig()
		test.readErrCheck(suite.T(), err)
		assert.Equal(suite.T(), test.s3Cfg, readS3Cfg)
		assert.Equal(suite.T(), test.account, readAccount)
	}
}
