package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/storage"
)

const (
	configFileTemplate = `
bucket = '%s'
endpoint = 's3.amazonaws.com'
prefix = 'test-prefix'
provider = 'S3'
account_provider = 'M365'
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
	b := "read-repo-config-basic-bucket"
	tID := "6f34ac30-8196-469b-bf8f-d83deadbbbba"
	testConfigData := fmt.Sprintf(configFileTemplate, b, tID)
	testConfigFilePath := path.Join(suite.T().TempDir(), "corso.toml")
	err := ioutil.WriteFile(testConfigFilePath, []byte(testConfigData), 0700)
	assert.NoError(suite.T(), err)

	// Configure viper to read test config file
	viper.SetConfigFile(testConfigFilePath)

	// Read and validate config
	err = readRepoConfig()
	require.NoError(suite.T(), err)

	s3Cfg, err := s3ConfigsFromViper()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), b, s3Cfg.Bucket)

	m365, err := m365ConfigsFromViper()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), tID, m365.TenantID)
}

func (suite *ConfigSuite) TestWriteReadConfig() {
	// Configure viper to read test config file
	tempDir := suite.T().TempDir()
	testConfigFilePath := path.Join(tempDir, "corso.toml")
	err := InitConfig(testConfigFilePath)
	assert.NoError(suite.T(), err)

	s3Cfg := storage.S3Config{Bucket: "write-read-config-bucket"}
	m365 := account.M365Config{TenantID: "3c0748d2-470e-444c-9064-1268e52609d5"}

	err = WriteRepoConfig(s3Cfg, m365)
	require.NoError(suite.T(), err)

	err = readRepoConfig()
	require.NoError(suite.T(), err)

	readS3Cfg, err := s3ConfigsFromViper()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), readS3Cfg.Bucket, s3Cfg.Bucket)

	readM365, err := m365ConfigsFromViper()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), readM365.TenantID, m365.TenantID)
}
