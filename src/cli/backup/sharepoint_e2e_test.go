package backup_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
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
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

// ---------------------------------------------------------------------------
// tests that require no existing backups
// ---------------------------------------------------------------------------

type NoBackupSharePointE2ESuite struct {
	tester.Suite
	dpnd dependencies
}

func TestNoBackupSharePointE2ESuite(t *testing.T) {
	suite.Run(t, &NoBackupSharePointE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *NoBackupSharePointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx, path.SharePointService)
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
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupSharepointE2ESuite struct {
	tester.Suite
	dpnd dependencies
	its  intgTesterSetup
}

func TestBackupSharepointE2ESuite(t *testing.T) {
	suite.Run(t, &BackupSharepointE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *BackupSharepointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx, path.SharePointService)
}

func (suite *BackupSharepointE2ESuite) TestSharepointBackupCmd_lists() {
	runSharepointBackupCategoryTest(suite, flags.DataLists)
}

func runSharepointBackupCategoryTest(suite *BackupSharepointE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildSharepointBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		suite.its.site.ID,
		category,
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)
}

func (suite *BackupSharepointE2ESuite) TestSharepointBackupCmd_siteNotFound_lists() {
	runSharepointBackupSiteNotFoundTest(suite, flags.DataLists)
}

func runSharepointBackupSiteNotFoundTest(suite *BackupSharepointE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildSharepointBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		uuid.NewString(),
		category,
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
	assert.Contains(
		t,
		err.Error(),
		"Invalid hostname for this tenancy", "error missing group not found")
	assert.NotContains(t, err.Error(), "runtime error", "panic happened")

	t.Logf("backup error message: %s", err.Error())

	result := recorder.String()
	t.Log("backup results", result)
}

// ---------------------------------------------------------------------------
// tests prepared with a previous backup
// ---------------------------------------------------------------------------

type PreparedBackupSharepointE2ESuite struct {
	tester.Suite
	dpnd      dependencies
	backupOps map[path.CategoryType]string
	its       intgTesterSetup
}

func TestPreparedBackupSharepointE2ESuite(t *testing.T) {
	suite.Run(t, &PreparedBackupSharepointE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *PreparedBackupSharepointE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx, path.SharePointService)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		sites = []string{suite.its.site.ID}
		ins   = idname.NewCache(map[string]string{suite.its.site.ID: suite.its.site.ID})
		cats  = []path.CategoryType{
			path.ListsCategory,
		}
	)

	for _, set := range cats {
		var (
			sel    = selectors.NewSharePointBackup(sites)
			scopes []selectors.SharePointScope
		)

		switch set {
		case path.ListsCategory:
			scopes = selTD.SharePointBackupListsScope(sel)
		}

		sel.Include(scopes)

		bop, err := suite.dpnd.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
		require.NoError(t, err, clues.ToCore(err))

		err = bop.Run(ctx)
		require.NoError(t, err, clues.ToCore(err))

		bIDs := string(bop.Results.BackupID)

		// sanity check, ensure we can find the backup and its details immediately
		b, err := suite.dpnd.repo.Backup(ctx, string(bop.Results.BackupID))
		require.NoError(t, err, "retrieving recent backup by ID")
		require.Equal(t, bIDs, string(b.ID), "repo backup matches results id")

		_, b, errs := suite.dpnd.repo.GetBackupDetails(ctx, bIDs)
		require.NoError(t, errs.Failure(), "retrieving recent backup details by ID")
		require.Empty(t, errs.Recovered(), "retrieving recent backup details by ID")
		require.Equal(t, bIDs, string(b.ID), "repo details matches results id")

		suite.backupOps[set] = string(b.ID)
	}
}

func (suite *PreparedBackupSharepointE2ESuite) TestSharepointListCmd_lists() {
	runSharepointListCmdTest(suite, path.ListsCategory)
}

func runSharepointListCmdTest(suite *PreparedBackupSharepointE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "sharepoint",
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.dpnd.recorder.String()
	assert.Contains(t, result, suite.backupOps[category])

	t.Log("backup results", result)
}

func (suite *PreparedBackupSharepointE2ESuite) TestSharepointListCmd_badID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "sharepoint",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", uuid.NewString())
	cli.BuildCommandTree(cmd)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *PreparedBackupSharepointE2ESuite) TestSharepointDetailsCmd_lists() {
	runSharepointDetailsCmdTest(suite, path.ListsCategory)
}

func runSharepointDetailsCmdTest(suite *PreparedBackupSharepointE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	bID := suite.backupOps[category]

	// fetch the details from the repo first
	deets, _, errs := suite.dpnd.repo.GetBackupDetails(ctx, string(bID))
	require.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
	require.Empty(t, errs.Recovered())

	cmd := cliTD.StubRootCmd(
		"backup", "details", "sharepoint",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, string(bID))
	cli.BuildCommandTree(cmd)
	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.dpnd.recorder.String()

	i := 0
	foundList := 0

	for _, ent := range deets.Entries {
		if ent.SharePoint != nil && ent.SharePoint.ItemName != "" {
			suite.Run(fmt.Sprintf("detail %d", i), func() {
				assert.Contains(suite.T(), result, ent.ShortRef)
			})
			foundList++
			i++
		}
	}

	assert.GreaterOrEqual(t, foundList, 1)
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteSharePointE2ESuite struct {
	tester.Suite
	dpnd              dependencies
	backupOp          operations.BackupOperation
	secondaryBackupOp operations.BackupOperation
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

	suite.dpnd = prepM365Test(t, ctx, path.SharePointService)

	var (
		m365SiteID = tconfig.M365SiteID(t)
		sites      = []string{m365SiteID}
		ins        = idname.NewCache(map[string]string{m365SiteID: m365SiteID})
	)

	// some tests require an existing backup
	sel := selectors.NewSharePointBackup(sites)
	sel.Include(selTD.SharePointBackupFolderScope(sel))

	backupOp, err := suite.dpnd.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// secondary backup
	secondaryBackupOp, err := suite.dpnd.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
	require.NoError(t, err, clues.ToCore(err))

	suite.secondaryBackupOp = secondaryBackupOp

	err = suite.secondaryBackupOp.Run(ctx)
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
		"--"+flags.BackupIDsFN,
		fmt.Sprintf("%s,%s",
			string(suite.backupOp.Results.BackupID),
			string(suite.secondaryBackupOp.Results.BackupID)))
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
			fmt.Sprintf("Deleted SharePoint backup [%s %s]\n",
				string(suite.backupOp.Results.BackupID),
				string(suite.secondaryBackupOp.Results.BackupID))))
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
		"--"+flags.BackupIDsFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteSharePointE2ESuite) TestSharePointBackupDeleteCmd_NoBackupID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	// empty backupIDs should error since no data provided
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func buildSharepointBackupCmd(
	ctx context.Context,
	configFile, site, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := cliTD.StubRootCmd(
		"backup", "create", "sharepoint",
		"--config-file", configFile,
		"--"+flags.SiteIDFN, site,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
