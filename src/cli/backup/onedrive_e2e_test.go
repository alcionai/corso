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

type NoBackupOneDriveE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	recorder   strings.Builder
}

func TestNoBackupOneDriveE2ESuite(t *testing.T) {
	suite.Run(t, &NoBackupOneDriveE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoCITests,
		),
	})
}

func (suite *NoBackupOneDriveE2ESuite) SetupSuite() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

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

	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx = config.SetViper(ctx, suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(
		ctx,
		suite.acct,
		suite.st,
		control.Options{
			// TODO: turn back on when this stops throttling-out the tests.
			// ToggleFeatures: control.Toggles{EnablePermissionsBackup: true},
		})
	require.NoError(t, err)
}

func (suite *NoBackupOneDriveE2ESuite) TestOneDriveBackupListCmd_empty() {
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
	assert.Equal(t, "No backups available\n", result)
}

func (suite *NoBackupOneDriveE2ESuite) TestOneDriveBackupCmd_UserNotInTenant() {
	recorder := strings.Builder{}

	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	ctx = config.SetViper(ctx, suite.vpr)

	cmd := tester.StubRootCmd(
		"backup", "create", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+utils.UserFN, "foo@nothere.com")
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"not found within tenant", "error missing user not found")
	assert.NotContains(t, err.Error(), "runtime error", "panic happened")

	t.Logf("backup error message: %s", err.Error())

	result := recorder.String()
	t.Log("backup results", result)
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteOneDriveE2ESuite struct {
	tester.Suite
	acct     account.Account
	st       storage.Storage
	vpr      *viper.Viper
	cfgFP    string
	repo     repository.Repository
	backupOp operations.BackupOperation
	recorder strings.Builder
}

func TestBackupDeleteOneDriveE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteOneDriveE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoCITests,
		),
	})
}

func (suite *BackupDeleteOneDriveE2ESuite) SetupSuite() {
	t := suite.T()

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
	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	// init the repo first
	suite.repo, err = repository.Initialize(
		ctx,
		suite.acct,
		suite.st,
		control.Options{
			// TODO: turn back on when this stops throttling-out the tests.
			// ToggleFeatures: control.Toggles{EnablePermissionsBackup: true},
		})
	require.NoError(t, err)

	m365UserID := tester.M365UserID(t)
	users := []string{m365UserID}

	// some tests require an existing backup
	sel := selectors.NewOneDriveBackup(users)
	sel.Include(sel.Folders(selectors.Any()))

	suite.backupOp, err = suite.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)
}

func (suite *BackupDeleteOneDriveE2ESuite) TestOneDriveBackupDeleteCmd() {
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

	assert.Equal(t, fmt.Sprintf("Deleted OneDrive backup %s\n", string(suite.backupOp.Results.BackupID)), result)

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = tester.StubRootCmd(
		"backup", "details", "onedrive",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	require.Error(t, cmd.ExecuteContext(ctx))
}

func (suite *BackupDeleteOneDriveE2ESuite) TestOneDriveBackupDeleteCmd_unknownID() {
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
