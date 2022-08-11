package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
)

const (
	configFileTemplate = `
` + BucketNameKey + ` = '%s'
` + EndpointKey + ` = 's3.amazonaws.com'
` + PrefixKey + ` = 'test-prefix'
` + StorageProviderTypeKey + ` = 'S3'
` + AccountProviderTypeKey + ` = 'M365'
` + TenantIDKey + ` = '%s'
`
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestReadRepoConfigBasic() {
	var (
		t   = suite.T()
		vpr = viper.New()
	)

	const (
		b   = "read-repo-config-basic-bucket"
		tID = "6f34ac30-8196-469b-bf8f-d83deadbbbba"
	)

	// Generate test config file
	testConfigData := fmt.Sprintf(configFileTemplate, b, tID)
	testConfigFilePath := path.Join(t.TempDir(), "corso.toml")
	err := ioutil.WriteFile(testConfigFilePath, []byte(testConfigData), 0o700)
	require.NoError(t, err)

	// Configure viper to read test config file
	vpr.SetConfigFile(testConfigFilePath)

	// Read and validate config
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	s3Cfg, err := s3ConfigsFromViper(vpr)
	require.NoError(t, err)
	assert.Equal(t, b, s3Cfg.Bucket)

	m365, err := m365ConfigsFromViper(vpr)
	require.NoError(t, err)
	assert.Equal(t, tID, m365.TenantID)
}

func (suite *ConfigSuite) TestWriteReadConfig() {
	var (
		t   = suite.T()
		vpr = viper.New()
	)

	const (
		bkt = "write-read-config-bucket"
		tid = "3c0748d2-470e-444c-9064-1268e52609d5"
	)

	// Configure viper to read test config file
	testConfigFilePath := path.Join(t.TempDir(), "corso.toml")
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt}
	m365 := account.M365Config{TenantID: tid}

	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	readS3Cfg, err := s3ConfigsFromViper(vpr)
	require.NoError(t, err)
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)

	readM365, err := m365ConfigsFromViper(vpr)
	require.NoError(t, err)
	assert.Equal(t, readM365.TenantID, m365.TenantID)
}

func (suite *ConfigSuite) TestMustMatchConfig() {
	var (
		t   = suite.T()
		vpr = viper.New()
	)

	const (
		bkt = "must-match-config-bucket"
		tid = "dfb12063-7598-458b-85ab-42352c5c25e2"
	)

	// Configure viper to read test config file
	testConfigFilePath := path.Join(t.TempDir(), "corso.toml")
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt}
	m365 := account.M365Config{TenantID: tid}

	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	table := []struct {
		name     string
		input    map[string]string
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name: "full match",
			input: map[string]string{
				storage.Bucket:   bkt,
				account.TenantID: tid,
			},
			errCheck: assert.NoError,
		},
		{
			name: "empty values",
			input: map[string]string{
				storage.Bucket:   "",
				account.TenantID: "",
			},
			errCheck: assert.NoError,
		},
		{
			name:     "no overrides",
			input:    map[string]string{},
			errCheck: assert.NoError,
		},
		{
			name:     "nil map",
			input:    nil,
			errCheck: assert.NoError,
		},
		{
			name: "no recognized keys",
			input: map[string]string{
				"fnords":   "smurfs",
				"nonsense": "",
			},
			errCheck: assert.NoError,
		},
		{
			name: "mismatch",
			input: map[string]string{
				storage.Bucket:   tid,
				account.TenantID: bkt,
			},
			errCheck: assert.Error,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			test.errCheck(t, mustMatchConfig(vpr, test.input))
		})
	}
}

// ------------------------------------------------------------
// integration tests
// ------------------------------------------------------------

type ConfigIntegrationSuite struct {
	suite.Suite
}

func TestConfigIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLIConfigTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(ConfigIntegrationSuite))
}

func (suite *ConfigIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *ConfigIntegrationSuite) TestGetStorageAndAccount() {
	t := suite.T()
	vpr := viper.New()

	const (
		bkt = "get-storage-and-account-bucket"
		end = "https://get-storage-and-account.com"
		pfx = "get-storage-and-account-prefix"
		tid = "3a2faa4e-a882-445c-9d27-f552ef189381"
	)

	// Configure viper to read test config file
	testConfigFilePath := path.Join(t.TempDir(), "corso.toml")
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{
		Bucket:   bkt,
		Endpoint: end,
		Prefix:   pfx,
	}
	m365 := account.M365Config{TenantID: tid}

	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	st, ac, err := getStorageAndAccountWithViper(vpr, true, nil)
	require.NoError(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	require.NoError(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.Endpoint, s3Cfg.Endpoint)
	assert.Equal(t, readS3Cfg.Prefix, s3Cfg.Prefix)

	assert.Equal(t, readS3Cfg.AccessKey, os.Getenv(credentials.AWSAccessKeyID))
	assert.Equal(t, readS3Cfg.SecretKey, os.Getenv(credentials.AWSSecretAccessKey))
	assert.Equal(t, readS3Cfg.SessionToken, os.Getenv(credentials.AWSSessionToken))

	common, err := st.CommonConfig()
	require.NoError(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassword, os.Getenv(credentials.CorsoPassword))

	readM365, err := ac.M365Config()
	require.NoError(t, err, "reading m365 config from account")
	assert.Equal(t, readM365.TenantID, m365.TenantID)
	assert.Equal(t, readM365.ClientID, os.Getenv(credentials.ClientID))
	assert.Equal(t, readM365.ClientSecret, os.Getenv(credentials.ClientSecret))
}

func (suite *ConfigIntegrationSuite) TestGetStorageAndAccount_noFileOnlyOverrides() {
	t := suite.T()
	vpr := viper.New()

	const (
		bkt = "get-storage-and-account-no-file-bucket"
		end = "https://get-storage-and-account.com/no-file"
		pfx = "get-storage-and-account-no-file-prefix"
		tid = "88f8522b-18e4-4d0f-b514-2d7b34d4c5a1"
	)

	// Configure viper to read test config file
	s3Cfg := storage.S3Config{
		Bucket:   bkt,
		Endpoint: end,
		Prefix:   pfx,
	}
	m365 := account.M365Config{TenantID: tid}

	overrides := map[string]string{
		account.TenantID:       tid,
		AccountProviderTypeKey: account.ProviderM365.String(),
		storage.Bucket:         bkt,
		storage.Endpoint:       end,
		storage.Prefix:         pfx,
		StorageProviderTypeKey: storage.ProviderS3.String(),
	}

	st, ac, err := getStorageAndAccountWithViper(vpr, false, overrides)
	require.NoError(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	require.NoError(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.Endpoint, s3Cfg.Endpoint)
	assert.Equal(t, readS3Cfg.Prefix, s3Cfg.Prefix)

	assert.Equal(t, readS3Cfg.AccessKey, os.Getenv(credentials.AWSAccessKeyID))
	assert.Equal(t, readS3Cfg.SecretKey, os.Getenv(credentials.AWSSecretAccessKey))
	assert.Equal(t, readS3Cfg.SessionToken, os.Getenv(credentials.AWSSessionToken))

	common, err := st.CommonConfig()
	require.NoError(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassword, os.Getenv(credentials.CorsoPassword))

	readM365, err := ac.M365Config()
	require.NoError(t, err, "reading m365 config from account")
	assert.Equal(t, readM365.TenantID, m365.TenantID)
	assert.Equal(t, readM365.ClientID, os.Getenv(credentials.ClientID))
	assert.Equal(t, readM365.ClientSecret, os.Getenv(credentials.ClientSecret))
}
