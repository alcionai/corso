package backup_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupOneDriveIntegrationSuite struct {
	suite.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	recorder   strings.Builder
}

func TestBackupOneDriveIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupOneDriveIntegrationSuite))
}

func (suite *BackupOneDriveIntegrationSuite) SetupSuite() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(t, err)

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}

	suite.vpr, suite.cfgFP, err = tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx = config.SetViper(ctx, suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)
}

func (suite *BackupOneDriveIntegrationSuite) TestOneDriveBackupListCmd_empty() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "list", "onedrive",
		"--config-file", suite.cfgFP)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))

	result := suite.recorder.String()

	// as an offhand check: the result should contain the m365 user id
	assert.Equal(t, result, "No backups available\n")
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteOneDriveIntegrationSuite struct {
	suite.Suite
	acct     account.Account
	st       storage.Storage
	vpr      *viper.Viper
	cfgFP    string
	repo     repository.Repository
	backupOp operations.BackupOperation
	recorder   strings.Builder
}

func TestBackupDeleteOneDriveIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupDeleteOneDriveIntegrationSuite))
}

func (suite *BackupDeleteOneDriveIntegrationSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(t, err)

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}
	suite.vpr, suite.cfgFP, err = tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)

	m365UserID := tester.M365UserID(t)

	// some tests require an existing backup
	sel := selectors.NewOneDriveBackup()
	sel.Include(sel.Folders([]string{m365UserID}, []string{"*"}))

	suite.backupOp, err = suite.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)
}

func (suite *BackupDeleteOneDriveIntegrationSuite) TestOneDriveBackupDeleteCmd() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "delete", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)
	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))

	result := suite.recorder.String()

	assert.Equal(t, result, fmt.Sprintf("Deleted OneDrive backup %s\n", string(suite.backupOp.Results.BackupID)))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = tester.StubRootCmd(
		"backup", "details", "onedrive",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	require.Error(t, cmd.ExecuteContext(ctx))
}

func (suite *BackupDeleteOneDriveIntegrationSuite) TestOneDriveBackupDeleteCmd_unknownID() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "delete", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	require.Error(t, cmd.ExecuteContext(ctx))
}
