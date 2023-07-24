package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

const (
	configFileTemplate = `
` + BucketNameKey + ` = '%s'
` + EndpointKey + ` = 's3.amazonaws.com'
` + PrefixKey + ` = 'test-prefix/'
` + StorageProviderTypeKey + ` = 'S3'
` + AccountProviderTypeKey + ` = 'M365'
` + AzureTenantIDKey + ` = '%s'
` + AccessKey + ` = '%s'
` + SecretAccessKey + ` = '%s'
` + SessionToken + ` = '%s'
` + CorsoPassphrase + ` = '%s'
` + AzureClientID + ` = '%s'
` + AzureSecret + ` = '%s'
` + DisableTLSKey + ` = '%s'
` + DisableTLSVerificationKey + ` = '%s'
`
)

type ConfigSuite struct {
	tester.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConfigSuite) TestRequireProps() {
	table := []struct {
		name     string
		props    map[string]string
		errCheck assert.ErrorAssertionFunc
	}{
		{
			props:    map[string]string{"exists": "I have seen the fnords!"},
			errCheck: assert.NoError,
		},
		{
			props:    map[string]string{"not-exists": ""},
			errCheck: assert.Error,
		},
	}
	for _, test := range table {
		err := requireProps(test.props)
		test.errCheck(suite.T(), err, clues.ToCore(err))
	}
}

func (suite *ConfigSuite) TestReadRepoConfigBasic() {
	var (
		t   = suite.T()
		vpr = viper.New()
	)

	const (
		b                      = "read-repo-config-basic-bucket"
		tID                    = "6f34ac30-8196-469b-bf8f-d83deadbbbba"
		accKey                 = "aws-test-access-key"
		secret                 = "aws-test-secret-key"
		token                  = "aws-test-session-token"
		passphrase             = "passphrase-test"
		azureClientID          = "azure-client-id-test"
		azureSecret            = "azure-secret-test"
		endpoint               = "s3-test"
		disableTLS             = "true"
		disableTLSVerification = "true"
	)

	// Generate test config file
	testConfigData := fmt.Sprintf(configFileTemplate, b, tID, accKey, secret,
		token, passphrase, azureClientID, azureSecret,
		disableTLS, disableTLSVerification)
	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	err := os.WriteFile(testConfigFilePath, []byte(testConfigData), 0o700)
	require.NoError(t, err, clues.ToCore(err))

	// Configure viper to read test config file
	vpr.SetConfigFile(testConfigFilePath)

	// Read and validate config
	err = vpr.ReadInConfig()
	require.NoError(t, err, "reading repo config", clues.ToCore(err))

	s3Cfg, err := s3ConfigsFromViper(vpr)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, b, s3Cfg.Bucket)
	assert.Equal(t, "test-prefix/", s3Cfg.Prefix)
	assert.Equal(t, disableTLS, strconv.FormatBool(s3Cfg.DoNotUseTLS))
	assert.Equal(t, disableTLSVerification, strconv.FormatBool(s3Cfg.DoNotVerifyTLS))

	s3Cfg, err = s3CredsFromViper(vpr, s3Cfg)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, accKey, s3Cfg.AWS.AccessKey)
	assert.Equal(t, secret, s3Cfg.AWS.SecretKey)
	assert.Equal(t, token, s3Cfg.AWS.SessionToken)

	m365, err := m365ConfigsFromViper(vpr)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, azureClientID, m365.AzureClientID)
	assert.Equal(t, azureSecret, m365.AzureClientSecret)
	assert.Equal(t, tID, m365.AzureTenantID)
}

func (suite *ConfigSuite) TestWriteReadConfig() {
	var (
		t   = suite.T()
		vpr = viper.New()
		// Configure viper to read test config file
		testConfigFilePath = filepath.Join(t.TempDir(), "corso.toml")
	)

	const (
		bkt    = "write-read-config-bucket"
		tid    = "3c0748d2-470e-444c-9064-1268e52609d5"
		repoID = "repoid"
		user   = "a-user"
		host   = "some-host"
	)

	err := initWithViper(vpr, testConfigFilePath)
	require.NoError(t, err, "initializing repo config", clues.ToCore(err))

	s3Cfg := storage.S3Config{Bucket: bkt, DoNotUseTLS: true, DoNotVerifyTLS: true}
	m365 := account.M365Config{AzureTenantID: tid}

	rOpts := repository.Options{
		User: user,
		Host: host,
	}

	err = writeRepoConfigWithViper(vpr, s3Cfg, m365, rOpts, repoID)
	require.NoError(t, err, "writing repo config", clues.ToCore(err))

	err = vpr.ReadInConfig()
	require.NoError(t, err, "reading repo config", clues.ToCore(err))

	readS3Cfg, err := s3ConfigsFromViper(vpr)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)

	readM365, err := m365ConfigsFromViper(vpr)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)

	gotUser, gotHost := getUserHost(vpr, true)
	assert.Equal(t, user, gotUser)
	assert.Equal(t, host, gotHost)
}

func (suite *ConfigSuite) TestMustMatchConfig() {
	var (
		t   = suite.T()
		vpr = viper.New()
		// Configure viper to read test config file
		testConfigFilePath = filepath.Join(t.TempDir(), "corso.toml")
	)

	const (
		bkt = "must-match-config-bucket"
		tid = "dfb12063-7598-458b-85ab-42352c5c25e2"
	)

	err := initWithViper(vpr, testConfigFilePath)
	require.NoError(t, err, "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt}
	m365 := account.M365Config{AzureTenantID: tid}

	err = writeRepoConfigWithViper(vpr, s3Cfg, m365, repository.Options{}, "repoid")
	require.NoError(t, err, "writing repo config", clues.ToCore(err))

	err = vpr.ReadInConfig()
	require.NoError(t, err, "reading repo config", clues.ToCore(err))

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
			errCheck: assert.NoError,
		},
		{
			name: "empty values",
			input: map[string]string{
				storage.Bucket:        "",
				account.AzureTenantID: "",
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
				storage.Bucket:        tid,
				account.AzureTenantID: bkt,
			},
			errCheck: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.errCheck(suite.T(), mustMatchConfig(vpr, test.input), clues.ToCore(err))
		})
	}
}

func (suite *ConfigSuite) TestReadFromFlags() {
	var (
		t   = suite.T()
		vpr = viper.New()
	)

	const (
		b                      = "read-repo-config-basic-bucket"
		tID                    = "6f34ac30-8196-469b-bf8f-d83deadbbbba"
		accKey                 = "aws-test-access-key"
		secret                 = "aws-test-secret-key"
		token                  = "aws-test-session-token"
		passphrase             = "passphrase-test"
		azureClientID          = "azure-client-id-test"
		azureSecret            = "azure-secret-test"
		prefix                 = "prefix-test"
		disableTLS             = "true"
		disableTLSVerification = "true"
	)

	t.Cleanup(func() {
		// reset values
		flags.AzureClientTenantFV = ""
		flags.AzureClientIDFV = ""
		flags.AzureClientSecretFV = ""

		flags.AWSAccessKeyFV = ""
		flags.AWSSecretAccessKeyFV = ""
		flags.AWSSessionTokenFV = ""

		flags.CorsoPassphraseFV = ""
	})

	// Generate test config file
	testConfigData := fmt.Sprintf(configFileTemplate, b, tID, accKey, secret, token,
		passphrase, azureClientID, azureSecret,
		disableTLS, disableTLSVerification)

	testConfigFilePath := filepath.Join(t.TempDir(), "corso.toml")
	err := os.WriteFile(testConfigFilePath, []byte(testConfigData), 0o700)
	require.NoError(t, err, clues.ToCore(err))

	// Configure viper to read test config file
	vpr.SetConfigFile(testConfigFilePath)

	// Read and validate config
	err = vpr.ReadInConfig()
	require.NoError(t, err, "reading repo config", clues.ToCore(err))

	overrides := map[string]string{}
	flags.AzureClientTenantFV = "6f34ac30-8196-469b-bf8f-d83deadbbbba"
	flags.AzureClientIDFV = "azure-id-flag-value"
	flags.AzureClientSecretFV = "azure-secret-flag-value"

	flags.AWSAccessKeyFV = "aws-access-key"
	flags.AWSSecretAccessKeyFV = "aws-access-secret-flag-value"
	flags.AWSSessionTokenFV = "aws-access-session-flag-value"

	overrides[storage.Bucket] = "flag-bucket"
	overrides[storage.Endpoint] = "flag-endpoint"
	overrides[storage.Prefix] = "flag-prefix"
	overrides[storage.DoNotUseTLS] = "true"
	overrides[storage.DoNotVerifyTLS] = "true"
	overrides[credentials.AWSAccessKeyID] = flags.AWSAccessKeyFV
	overrides[credentials.AWSSecretAccessKey] = flags.AWSSecretAccessKeyFV
	overrides[credentials.AWSSessionToken] = flags.AWSSessionTokenFV

	flags.CorsoPassphraseFV = "passphrase-flags"

	repoDetails, err := getStorageAndAccountWithViper(
		vpr,
		true,
		false,
		overrides,
	)

	m365Config, _ := repoDetails.Account.M365Config()
	s3Cfg, _ := repoDetails.Storage.S3Config()
	commonConfig, _ := repoDetails.Storage.CommonConfig()
	pass := commonConfig.Corso.CorsoPassphrase

	require.NoError(t, err, "reading repo config", clues.ToCore(err))

	assert.Equal(t, flags.AWSAccessKeyFV, s3Cfg.AWS.AccessKey)
	assert.Equal(t, flags.AWSSecretAccessKeyFV, s3Cfg.AWS.SecretKey)
	assert.Equal(t, flags.AWSSessionTokenFV, s3Cfg.AWS.SessionToken)

	assert.Equal(t, overrides[storage.Bucket], s3Cfg.Bucket)
	assert.Equal(t, overrides[storage.Endpoint], s3Cfg.Endpoint)
	assert.Equal(t, overrides[storage.Prefix], s3Cfg.Prefix)
	assert.Equal(t, str.ParseBool(overrides[storage.DoNotUseTLS]), s3Cfg.DoNotUseTLS)
	assert.Equal(t, str.ParseBool(overrides[storage.DoNotVerifyTLS]), s3Cfg.DoNotVerifyTLS)

	assert.Equal(t, flags.AzureClientIDFV, m365Config.AzureClientID)
	assert.Equal(t, flags.AzureClientSecretFV, m365Config.AzureClientSecret)
	assert.Equal(t, flags.AzureClientTenantFV, m365Config.AzureTenantID)

	assert.Equal(t, flags.CorsoPassphraseFV, pass)
}

// ------------------------------------------------------------
// integration tests
// ------------------------------------------------------------

type ConfigIntegrationSuite struct {
	tester.Suite
}

func TestConfigIntegrationSuite(t *testing.T) {
	suite.Run(t, &ConfigIntegrationSuite{Suite: tester.NewIntegrationSuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
	)})
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

	err := initWithViper(vpr, testConfigFilePath)
	require.NoError(t, err, "initializing repo config", clues.ToCore(err))

	s3Cfg := storage.S3Config{
		Bucket:         bkt,
		Endpoint:       end,
		Prefix:         pfx,
		DoNotVerifyTLS: true,
		DoNotUseTLS:    true,
	}
	m365 := account.M365Config{AzureTenantID: tid}

	err = writeRepoConfigWithViper(vpr, s3Cfg, m365, repository.Options{}, "repoid")
	require.NoError(t, err, "writing repo config", clues.ToCore(err))

	err = vpr.ReadInConfig()
	require.NoError(t, err, "reading repo config", clues.ToCore(err))

	cfg, err := getStorageAndAccountWithViper(vpr, true, true, nil)
	require.NoError(t, err, "getting storage and account from config", clues.ToCore(err))

	readS3Cfg, err := cfg.Storage.S3Config()
	require.NoError(t, err, "reading s3 config from storage", clues.ToCore(err))
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.Endpoint, s3Cfg.Endpoint)
	assert.Equal(t, readS3Cfg.Prefix, s3Cfg.Prefix)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)
	assert.Equal(t, cfg.RepoID, "repoid")

	common, err := cfg.Storage.CommonConfig()
	require.NoError(t, err, "reading common config from storage", clues.ToCore(err))
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := cfg.Account.M365Config()
	require.NoError(t, err, "reading m365 config from account", clues.ToCore(err))
	// Env var gets preference here. Where to get env tenantID from
	// assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
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

	cfg, err := getStorageAndAccountWithViper(vpr, false, true, overrides)
	require.NoError(t, err, "getting storage and account from config", clues.ToCore(err))

	readS3Cfg, err := cfg.Storage.S3Config()
	require.NoError(t, err, "reading s3 config from storage", clues.ToCore(err))
	assert.Equal(t, readS3Cfg.Bucket, bkt)
	assert.Equal(t, cfg.RepoID, "")
	assert.Equal(t, readS3Cfg.Endpoint, end)
	assert.Equal(t, readS3Cfg.Prefix, pfx)
	assert.True(t, readS3Cfg.DoNotUseTLS)
	assert.True(t, readS3Cfg.DoNotVerifyTLS)

	common, err := cfg.Storage.CommonConfig()
	require.NoError(t, err, "reading common config from storage", clues.ToCore(err))
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := cfg.Account.M365Config()
	require.NoError(t, err, "reading m365 config from account", clues.ToCore(err))
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
	assert.Equal(t, readM365.AzureClientID, os.Getenv(credentials.AzureClientID))
	assert.Equal(t, readM365.AzureClientSecret, os.Getenv(credentials.AzureClientSecret))
}
