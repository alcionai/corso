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
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/storage"
)

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type NoBackupSharePointE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365SiteID string
	recorder   strings.Builder
}

func TestNoBackupSharePointE2ESuite(t *testing.T) {
	suite.Run(t, &NoBackupSharePointE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
	)})
}

func (suite *NoBackupSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.recorder = recorder
	suite.cfgFP = cfgFilePath
	suite.m365SiteID = tester.M365SiteID(t)
}

func (suite *NoBackupSharePointE2ESuite) TestSharePointBackupListCmd_empty() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "list", "sharepoint",
		"--config-file", suite.cfgFP)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.recorder.String()

	// as an offhand check: the result should contain the m365 sitet id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteSharePointE2ESuite struct {
	tester.Suite
	acct     account.Account
	st       storage.Storage
	vpr      *viper.Viper
	cfgFP    string
	repo     repository.Repository
	backupOp operations.BackupOperation
	recorder strings.Builder
}

func TestBackupDeleteSharePointE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteSharePointE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.recorder = recorder
	suite.cfgFP = cfgFilePath

	var (
		m365SiteID = tester.M365SiteID(t)
		sites      = []string{m365SiteID}
		ins        = idname.NewCache(map[string]string{m365SiteID: m365SiteID})
	)

	// some tests require an existing backup
	sel := selectors.NewSharePointBackup(sites)
	sel.Include(testdata.SharePointBackupFolderScope(sel))

	backupOp, err := suite.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "delete", "sharepoint",
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
			fmt.Sprintf("Deleted SharePoint backup %s\n", string(suite.backupOp.Results.BackupID)),
		),
	)
}

// moved out of the func above to make the linter happy
// // a follow-up details call should fail, due to the backup ID being deleted
// cmd = tester.StubRootCmd(
// 	"backup", "details", "sharepoint",
// 	"--config-file", suite.cfgFP,
// 	"--backup", string(suite.backupOp.Results.BackupID))
// cli.BuildCommandTree(cmd)

// err := cmd.ExecuteContext(ctx)
// require.Error(t, err, clues.ToCore(err))

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd_unknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "delete", "sharepoint",
		"--config-file", suite.cfgFP,
		"--"+flags.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
