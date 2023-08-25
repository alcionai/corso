package backup_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
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
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/selectors/testdata"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type NoBackupSharePointE2ESuite struct {
	tester.Suite
	dpnd dependencies
}

func TestNoBackupSharePointE2ESuite(t *testing.T) {
	suite.Run(t, &NoBackupSharePointE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
	)})
}

func (suite *NoBackupSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx)
}

func (suite *NoBackupSharePointE2ESuite) TestSharePointBackupListCmd_empty() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "sharepoint",
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()

	// as an offhand check: the result should contain the m365 sitet id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteSharePointE2ESuite struct {
	tester.Suite
	dpnd     dependencies
	backupOp operations.BackupOperation
}

func TestBackupDeleteSharePointE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteSharePointE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx)

	var (
		m365SiteID = tconfig.M365SiteID(t)
		sites      = []string{m365SiteID}
		ins        = idname.NewCache(map[string]string{m365SiteID: m365SiteID})
	)

	// some tests require an existing backup
	sel := selectors.NewSharePointBackup(sites)
	sel.Include(testdata.SharePointBackupFolderScope(sel))

	backupOp, err := suite.dpnd.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "sharepoint",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)
	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()
	assert.True(t,
		strings.HasSuffix(
			result,
			fmt.Sprintf("Deleted SharePoint backup %s\n", string(suite.backupOp.Results.BackupID)),
		),
	)
}

// moved out of the func above to make the linter happy
// // a follow-up details call should fail, due to the backup ID being deleted
// cmd = cliTD.StubRootCmd(
// 	"backup", "details", "sharepoint",
// 	"--config-file", suite.cfgFP,
// 	"--backup", string(suite.backupOp.Results.BackupID))
// cli.BuildCommandTree(cmd)

// err := cmd.ExecuteContext(ctx)
// require.Error(t, err, clues.ToCore(err))

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd_unknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "sharepoint",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
