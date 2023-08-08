package backup_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/print"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
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
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *NoBackupOneDriveE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.recorder = recorder
	suite.vpr = vpr
	suite.cfgFP = cfgFilePath
	suite.m365UserID = tconfig.M365UserID(t)
}

func (suite *NoBackupOneDriveE2ESuite) TestOneDriveBackupListCmd_empty() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "onedrive",
		"--config-file", suite.cfgFP)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.recorder.String()

	// as an offhand check: the result should contain the m365 user id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

func (suite *NoBackupOneDriveE2ESuite) TestOneDriveBackupCmd_UserNotInTenant() {
	recorder := strings.Builder{}

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	ctx = config.SetViper(ctx, suite.vpr)

	cmd := cliTD.StubRootCmd(
		"backup", "create", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+flags.UserFN, "foo@nothere.com")
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
	assert.Contains(
		t,
		err.Error(),
		"not found in tenant", "error missing user not found")
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
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteOneDriveE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.recorder = recorder
	suite.vpr = vpr
	suite.cfgFP = cfgFilePath

	var (
		m365UserID = tconfig.M365UserID(t)
		users      = []string{m365UserID}
		ins        = idname.NewCache(map[string]string{m365UserID: m365UserID})
	)

	// some tests require an existing backup
	sel := selectors.NewOneDriveBackup(users)
	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	backupOp, err := suite.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteOneDriveE2ESuite) TestOneDriveBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+flags.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)
	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.recorder.String()
	assert.True(t,
		strings.HasSuffix(
			result,
			fmt.Sprintf("Deleted OneDrive backup %s\n", string(suite.backupOp.Results.BackupID)),
		),
	)

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "onedrive",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteOneDriveE2ESuite) TestOneDriveBackupDeleteCmd_unknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "onedrive",
		"--config-file", suite.cfgFP,
		"--"+flags.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
