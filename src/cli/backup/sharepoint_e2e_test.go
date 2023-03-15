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

	"github.com/alcionai/clues"
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
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests)})
}

func (suite *NoBackupSharePointE2ESuite) SetupSuite() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}

	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx = config.SetViper(ctx, suite.vpr)
	suite.m365SiteID = tester.M365SiteID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *NoBackupSharePointE2ESuite) TestSharePointBackupListCmd_empty() {
	t := suite.T()
	ctx, flush := tester.NewContext()
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
	assert.Equal(t, "No backups available\n", result)
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
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoCITests,
			tester.CorsoCLITests,
			tester.CorsoCLIBackupTests),
	})
}

func (suite *BackupDeleteSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

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
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err, clues.ToCore(err))

	m365SiteID := tester.M365SiteID(t)
	sites := []string{m365SiteID}

	// some tests require an existing backup
	sel := selectors.NewSharePointBackup(sites)
	sel.Include(sel.LibraryFolders(selectors.Any()))

	suite.backupOp, err = suite.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, err, clues.ToCore(err))

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "delete", "sharepoint",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)
	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.recorder.String()
	expect := fmt.Sprintf("Deleted SharePoint backup %s\n", string(suite.backupOp.Results.BackupID))
	assert.Equal(t, expect, result)
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
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "delete", "sharepoint",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
