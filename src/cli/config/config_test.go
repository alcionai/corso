package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/store"
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
	tester.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigSuite{Suite: tester.NewUnitSuite(t)})
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
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt, DoNotUseTLS: true, DoNotVerifyTLS: true}
	m365 := account.M365Config{AzureTenantID: tid}
	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365, mockRepo{}), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	readS3Cfg, err := s3ConfigsFromViper(vpr)
	require.NoError(t, err)
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)

	readM365, err := m365ConfigsFromViper(vpr)
	require.NoError(t, err)
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
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{Bucket: bkt}
	m365 := account.M365Config{AzureTenantID: tid}

	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365, mockRepo{}), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

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
			test.errCheck(suite.T(), mustMatchConfig(vpr, test.input))
		})
	}
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
		[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
		tester.CorsoCLIConfigTests)})
}

// mockRepo implementing Repository for test
type mockRepo struct{}

func (repo mockRepo) GetID() string { return "repoid" }

func (repo mockRepo) Close(context.Context) error { return nil }

func (repo mockRepo) NewBackup(
	ctx context.Context,
	self selectors.Selector,
) (operations.BackupOperation, error) {
	return operations.BackupOperation{}, nil
}

func (repo mockRepo) NewRestore(
	ctx context.Context,
	backupID string,
	sel selectors.Selector,
	dest control.RestoreDestination,
) (operations.RestoreOperation, error) {
	return operations.RestoreOperation{}, nil
}

func (repo mockRepo) DeleteBackup(ctx context.Context, id model.StableID) error { return nil }

// backups lists a backup by id
func (repo mockRepo) Backup(ctx context.Context, id model.StableID) (*backup.Backup, error) {
	return &backup.Backup{}, nil
}

func (repo mockRepo) BackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, *fault.Errors) {
	return nil, nil, nil
}

func (repo mockRepo) Backups(
	context.Context,
	[]model.StableID,
) ([]*backup.Backup, *fault.Errors) {
	return nil, nil
}

// backups lists backups in a repository
func (repo mockRepo) BackupsByTag(ctx context.Context, fs ...store.FilterOption) ([]*backup.Backup, error) {
	// sw := store.NewKopiaStore(r.modelStore)
	return []*backup.Backup{}, nil
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
	require.NoError(t, initWithViper(vpr, testConfigFilePath), "initializing repo config")

	s3Cfg := storage.S3Config{
		Bucket:         bkt,
		Endpoint:       end,
		Prefix:         pfx,
		DoNotVerifyTLS: true,
		DoNotUseTLS:    true,
	}
	m365 := account.M365Config{AzureTenantID: tid}

	require.NoError(t, writeRepoConfigWithViper(vpr, s3Cfg, m365, mockRepo{}), "writing repo config")
	require.NoError(t, vpr.ReadInConfig(), "reading repo config")

	st, ac, repoID, err := getStorageAndAccountWithViper(vpr, true, nil)
	require.NoError(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	require.NoError(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, s3Cfg.Bucket)
	assert.Equal(t, readS3Cfg.Endpoint, s3Cfg.Endpoint)
	assert.Equal(t, readS3Cfg.Prefix, s3Cfg.Prefix)
	assert.Equal(t, readS3Cfg.DoNotUseTLS, s3Cfg.DoNotUseTLS)
	assert.Equal(t, readS3Cfg.DoNotVerifyTLS, s3Cfg.DoNotVerifyTLS)
	assert.Equal(t, repoID, "repoid")

	common, err := st.CommonConfig()
	require.NoError(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := ac.M365Config()
	require.NoError(t, err, "reading m365 config from account")
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

	st, ac, repoid, err := getStorageAndAccountWithViper(vpr, false, overrides)
	require.NoError(t, err, "getting storage and account from config")

	readS3Cfg, err := st.S3Config()
	require.NoError(t, err, "reading s3 config from storage")
	assert.Equal(t, readS3Cfg.Bucket, bkt)
	assert.Equal(t, repoid, "")
	assert.Equal(t, readS3Cfg.Endpoint, end)
	assert.Equal(t, readS3Cfg.Prefix, pfx)
	assert.True(t, readS3Cfg.DoNotUseTLS)
	assert.True(t, readS3Cfg.DoNotVerifyTLS)

	common, err := st.CommonConfig()
	require.NoError(t, err, "reading common config from storage")
	assert.Equal(t, common.CorsoPassphrase, os.Getenv(credentials.CorsoPassphrase))

	readM365, err := ac.M365Config()
	require.NoError(t, err, "reading m365 config from account")
	assert.Equal(t, readM365.AzureTenantID, m365.AzureTenantID)
	assert.Equal(t, readM365.AzureClientID, os.Getenv(credentials.AzureClientID))
	assert.Equal(t, readM365.AzureClientSecret, os.Getenv(credentials.AzureClientSecret))
}
