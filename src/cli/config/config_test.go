package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	configFileTemplate = `
` + BucketNameKey + ` = '%s'
` + EndpointKey + ` = 's3.amazonaws.com'
` + PrefixKey + ` = 'test-prefix/'
` + StorageProviderTypeKey + ` = 'S3'
` + AccountProviderTypeKey + ` = 'M365'
` + AzureTenantIDKey + ` = '%s'
` + DisableTLSKey + ` = 'false'
` + DisableTLSVerificationKey + ` = 'false'
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
	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	err := os.WriteFile(testConfigFilePath, []byte(testConfigData), 0o700)
	aw.MustNoErr(t, err)

	// Configure viper to read test config file
	vpr.SetConfigFile(testConfigFilePath)

	// Read and validate config
	aw.MustNoErr(t, vpr.ReadInConfig(), "reading repo config")

	s3Cfg, err := s3ConfigsFromViper(vpr)
	aw.MustNoErr(t, err)
	assert.Equal(t, b, s3Cfg.Bucket)

	m365, err := m365ConfigsFromViper(vpr)
	aw.MustNoErr(t, err)
	assert.Equal(t, tID, m365.AzureTenantID)
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
	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	aw.MustNoErr(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt, DoNotUseTLS: true, DoNotVerifyTLS: true}
	m365 := account.M365Config{AzureTenantID: tid}

	aw.MustNoErr(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	aw.MustNoErr(t, vpr.ReadInConfig(), "reading repo config")

	readS3Cfg, err := s3ConfigsFromViper(vpr)
	aw.MustNoErr(t, err)
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)

	readM365, err := m365ConfigsFromViper(vpr)
	aw.MustNoErr(t, err)
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
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
	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	aw.MustNoErr(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt}
	m365 := account.M365Config{AzureTenantID: tid}

	aw.MustNoErr(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	aw.MustNoErr(t, vpr.ReadInConfig(), "reading repo config")

	table := []struct {
		name     string
		input    map[string]string
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name: "full match",
			input: map[string]string{
				storage.Bucket:        bkt,
				account.AzureTenantID: tid,
			},
			errCheck: aw.NoErr,
		},
		{
			name: "empty values",
			input: map[string]string{
				storage.Bucket:        "",
				account.AzureTenantID: "",
			},
			errCheck: aw.NoErr,
		},
		{
			name:     "no overrides",
			input:    map[string]string{},
			errCheck: aw.NoErr,
		},
		{
			name:     "nil map",
			input:    nil,
			errCheck: aw.NoErr,
		},
		{
			name: "no recognized keys",
			input: map[string]string{
				"fnords":   "smurfs",
				"nonsense": "",
			},
			errCheck: aw.NoErr,
		},
		{
			name: "mismatch",
			input: map[string]string{
				storage.Bucket:        tid,
				account.AzureTenantID: bkt,
			},
			errCheck: aw.Err,
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
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoCLIConfigTests)

	suite.Run(t, new(ConfigIntegrationSuite))
}

func (suite *ConfigIntegrationSuite) SetupSuite() {
	tester.MustGetEnvSets(suite.T(), tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs)
}

func (suite *ConfigIntegrationSuite) TestGetStorageAndAccount() {
	t := suite.T()
	vpr := viper.New()

	const (
		bkt = "get-storage-and-account-bucket"
		end = "https://get-storage-and-account.com"
		pfx = "get-storage-and-account-prefix/"
		tid = "3a2faa4e-a882-445c-9d27-f552ef189381"
	)

	// Configure viper to read test config file
	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	aw.MustNoErr(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{
		Bucket:         bkt,
		Endpoint:       end,
		Prefix:         pfx,
		DoNotVerifyTLS: true,
		DoNotUseTLS:    true,
	}
	m365 := account.M365Config{AzureTenantID: tid}

	aw.MustNoErr(t, writeRepoConfigWithViper(vpr, s3Cfg, m365), "writing repo config")
	aw.MustNoErr(t, vpr.ReadInConfig(), "reading repo config")

	st, ac, err := getStorageAndAccountWithViper(vpr, true, nil)
	aw.MustNoErr(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	aw.MustNoErr(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.Endpoint, s3Cfg.Endpoint)
	assert.Equal(t, readS3Cfg.Prefix, s3Cfg.Prefix)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)

	common, err := st.CommonConfig()
	aw.MustNoErr(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := ac.M365Config()
	aw.MustNoErr(t, err, "reading m365 config from account")
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
	assert.Equal(t, readM365.AzureClientID, os.Getenv(credentials.AzureClientID))
	assert.Equal(t, readM365.AzureClientSecret, os.Getenv(credentials.AzureClientSecret))
}

func (suite *ConfigIntegrationSuite) TestGetStorageAndAccount_noFileOnlyOverrides() {
	t := suite.T()
	vpr := viper.New()

	const (
		bkt = "get-storage-and-account-no-file-bucket"
		end = "https://get-storage-and-account.com/no-file"
		pfx = "get-storage-and-account-no-file-prefix/"
		tid = "88f8522b-18e4-4d0f-b514-2d7b34d4c5a1"
	)

	m365 := account.M365Config{AzureTenantID: tid}

	overrides := map[string]string{
		account.AzureTenantID:  tid,
		AccountProviderTypeKey: account.ProviderM365.String(),
		storage.Bucket:         bkt,
		storage.Endpoint:       end,
		storage.Prefix:         pfx,
		storage.DoNotUseTLS:    "true",
		storage.DoNotVerifyTLS: "true",
		StorageProviderTypeKey: storage.ProviderS3.String(),
	}

	st, ac, err := getStorageAndAccountWithViper(vpr, false, overrides)
	aw.MustNoErr(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	aw.MustNoErr(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, bkt)
	assert.Equal(t, readS3Cfg.Endpoint, end)
	assert.Equal(t, readS3Cfg.Prefix, pfx)
	assert.True(t, readS3Cfg.DoNotUseTLS)
	assert.True(t, readS3Cfg.DoNotVerifyTLS)

	common, err := st.CommonConfig()
	aw.MustNoErr(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := ac.M365Config()
	aw.MustNoErr(t, err, "reading m365 config from account")
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
	assert.Equal(t, readM365.AzureClientID, os.Getenv(credentials.AzureClientID))
	assert.Equal(t, readM365.AzureClientSecret, os.Getenv(credentials.AzureClientSecret))
}
