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
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

var (
	channelMessages = path.ChannelMessagesCategory
	libraries       = path.LibrariesCategory
)

// ---------------------------------------------------------------------------
// tests that require no existing backups
// ---------------------------------------------------------------------------

type NoBackupGroupsE2ESuite struct {
	tester.Suite
	dpnd dependencies
	its  intgTesterSetup
}

func TestNoBackupGroupsE2ESuite(t *testing.T) {
	t.Skip("enable when e2e is complete for groups")

	suite.Run(t, &BackupGroupsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
	)})
}

func (suite *NoBackupGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx)
}

func (suite *NoBackupGroupsE2ESuite) TestGroupsBackupListCmd_noBackups() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()

	// as an offhand check: the result should contain the m365 group id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupGroupsE2ESuite struct {
	tester.Suite
	dpnd dependencies
	its  intgTesterSetup
}

func TestBackupGroupsE2ESuite(t *testing.T) {
	t.Skip("enable when e2e is complete for groups")

	suite.Run(t, &BackupGroupsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
	)})
}

func (suite *BackupGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_channelMessages() {
	runGroupsBackupCategoryTest(suite, channelMessages)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_libraries() {
	runGroupsBackupCategoryTest(suite, libraries)
}

func runGroupsBackupCategoryTest(suite *BackupGroupsE2ESuite, category path.CategoryType) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildGroupsBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		suite.its.group.ID,
		category.String(),
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 group id
	assert.Contains(t, result, suite.its.group.ID)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_groupNotFound_channelMessages() {
	runGroupsBackupGroupNotFoundTest(suite, channelMessages)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_groupNotFound_libraries() {
	runGroupsBackupGroupNotFoundTest(suite, libraries)
}

func runGroupsBackupGroupNotFoundTest(suite *BackupGroupsE2ESuite, category path.CategoryType) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildGroupsBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		"foo@not-there.com",
		category.String(),
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
	assert.Contains(
		t,
		err.Error(),
		"not found in tenant", "error missing group not found")
	assert.NotContains(t, err.Error(), "runtime error", "panic happened")

	t.Logf("backup error message: %s", err.Error())

	result := recorder.String()
	t.Log("backup results", result)
}

func (suite *BackupGroupsE2ESuite) TestBackupCreateGroups_badAzureClientIDFlag() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "groups",
		"--group", suite.its.group.ID,
		"--azure-client-id", "invalid-value")
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupGroupsE2ESuite) TestBackupCreateGroups_fromConfigFile() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "groups",
		"--group", suite.its.group.ID,
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 group id
	assert.Contains(t, result, suite.its.group.ID)
}

// AWS flags
func (suite *BackupGroupsE2ESuite) TestBackupCreateGroups_badAWSFlags() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "groups",
		"--group", suite.its.group.ID,
		"--aws-access-key", "invalid-value",
		"--aws-secret-access-key", "some-invalid-value",
	)
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	// since invalid aws creds are explicitly set, should see a failure
	require.Error(t, err, clues.ToCore(err))
}

// ---------------------------------------------------------------------------
// tests prepared with a previous backup
// ---------------------------------------------------------------------------

type PreparedBackupGroupsE2ESuite struct {
	tester.Suite
	dpnd      dependencies
	backupOps map[path.CategoryType]string
	its       intgTesterSetup
}

func TestPreparedBackupGroupsE2ESuite(t *testing.T) {
	t.Skip("enable when e2e is complete for groups")

	suite.Run(t, &PreparedBackupGroupsE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *PreparedBackupGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		groups = []string{suite.its.group.ID}
		ins    = idname.NewCache(map[string]string{suite.its.group.ID: suite.its.group.ID})
	)

	for _, set := range []path.CategoryType{channelMessages, libraries} {
		var (
			sel    = selectors.NewGroupsBackup(groups)
			scopes []selectors.GroupsScope
		)

		switch set {
		case channelMessages:
			scopes = sel.Channel("TODO-test-channel-const")

		case libraries:
			scopes = sel.LibraryFolders([]string{"TODO-test-library-folder-const"}, selectors.PrefixMatch())
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

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_channelMessages() {
	runGroupsListCmdTest(suite, channelMessages)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_libraries() {
	runGroupsListCmdTest(suite, libraries)
}

func runGroupsListCmdTest(suite *PreparedBackupGroupsE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
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
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_singleID_channelMessages() {
	runGroupsListSingleCmdTest(suite, channelMessages)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_singleID_libraries() {
	runGroupsListSingleCmdTest(suite, libraries)
}

func runGroupsListSingleCmdTest(suite *PreparedBackupGroupsE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	bID := suite.backupOps[category]

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", string(bID))
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.dpnd.recorder.String()
	assert.Contains(t, result, bID)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_badID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", "smarfs")
	cli.BuildCommandTree(cmd)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsDetailsCmd_channelMessages() {
	runGroupsDetailsCmdTest(suite, channelMessages)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsDetailsCmd_libraries() {
	runGroupsDetailsCmdTest(suite, libraries)
}

func runGroupsDetailsCmdTest(suite *PreparedBackupGroupsE2ESuite, category path.CategoryType) {
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
		"backup", "details", "groups",
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
	foundFolders := 0

	for _, ent := range deets.Entries {
		// Skip folders as they don't mean anything to the end group.
		if ent.Folder != nil {
			foundFolders++
			continue
		}

		suite.Run(fmt.Sprintf("detail %d", i), func() {
			assert.Contains(suite.T(), result, ent.ShortRef)
		})

		i++
	}

	// We only backup the default folder for each category so there should be at
	// least that folder (we don't make details entries for prefix folders).
	assert.GreaterOrEqual(t, foundFolders, 1)
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteGroupsE2ESuite struct {
	tester.Suite
	dpnd     dependencies
	backupOp operations.BackupOperation
}

func TestBackupDeleteGroupsE2ESuite(t *testing.T) {
	t.Skip("enable when e2e is complete for groups")

	suite.Run(t, &BackupDeleteGroupsE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *BackupDeleteGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx)

	m365GroupID := tconfig.M365GroupID(t)
	groups := []string{m365GroupID}

	// some tests require an existing backup
	sel := selectors.NewGroupsBackup(groups)
	sel.Include(sel.Channel("TODO-test-channel-const"))

	backupOp, err := suite.dpnd.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteGroupsE2ESuite) TestGroupsBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "groups",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteGroupsE2ESuite) TestGroupsBackupDeleteCmd_UnknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func buildGroupsBackupCmd(
	ctx context.Context,
	configFile, group, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := cliTD.StubRootCmd(
		"backup", "create", "groups",
		"--config-file", configFile,
		"--"+flags.GroupFN, group,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
