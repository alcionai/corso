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

	"github.com/alcionai/canario/src/cli"
	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/cli/print"
	cliTD "github.com/alcionai/canario/src/cli/testdata"
	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/internal/operations"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/its"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/config"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
	selTD "github.com/alcionai/canario/src/pkg/selectors/testdata"
	storeTD "github.com/alcionai/canario/src/pkg/storage/testdata"
)

// ---------------------------------------------------------------------------
// tests that require no existing backups
// ---------------------------------------------------------------------------

type NoBackupGroupsE2ESuite struct {
	tester.Suite
	dpnd dependencies
	m365 its.M365IntgTestSetup
}

func TestNoBackupGroupsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupGroupsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *NoBackupGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.GroupsService)
}

func (suite *NoBackupGroupsE2ESuite) TestGroupsBackupListCmd_noBackups() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
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
	m365 its.M365IntgTestSetup
}

func TestBackupGroupsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupGroupsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *BackupGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.GroupsService)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_channelMessages() {
	runGroupsBackupCategoryTest(suite, flags.DataMessages)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_conversations() {
	runGroupsBackupCategoryTest(suite, flags.DataConversations)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_libraries() {
	runGroupsBackupCategoryTest(suite, flags.DataLibraries)
}

func runGroupsBackupCategoryTest(suite *BackupGroupsE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildGroupsBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		suite.m365.Group.ID,
		category,
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_groupNotFound_channelMessages() {
	runGroupsBackupGroupNotFoundTest(suite, flags.DataMessages)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_groupNotFound_conversations() {
	runGroupsBackupGroupNotFoundTest(suite, flags.DataConversations)
}

func (suite *BackupGroupsE2ESuite) TestGroupsBackupCmd_groupNotFound_libraries() {
	runGroupsBackupGroupNotFoundTest(suite, flags.DataLibraries)
}

func runGroupsBackupGroupNotFoundTest(suite *BackupGroupsE2ESuite, category string) {
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
		category,
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
	assert.Contains(
		t,
		err.Error(),
		"not found",
		"error missing user not found")
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
		"--group", suite.m365.Group.ID,
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
		"--group", suite.m365.Group.ID,
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

// AWS flags
func (suite *BackupGroupsE2ESuite) TestBackupCreateGroups_badAWSFlags() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "groups",
		"--group", suite.m365.Group.ID,
		"--aws-access-key", "invalid-value",
		"--aws-secret-access-key", "some-invalid-value")
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
	m365      its.M365IntgTestSetup
}

func TestPreparedBackupGroupsE2ESuite(t *testing.T) {
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

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.GroupsService)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		groups = []string{suite.m365.Group.ID}
		ins    = idname.NewCache(map[string]string{suite.m365.Group.ID: suite.m365.Group.ID})
		cats   = []path.CategoryType{
			path.ChannelMessagesCategory,
			path.ConversationPostsCategory,
			path.LibrariesCategory,
		}
	)

	for _, set := range cats {
		var (
			sel    = selectors.NewGroupsBackup(groups)
			scopes []selectors.GroupsScope
		)

		switch set {
		case path.ChannelMessagesCategory:
			scopes = selTD.GroupsBackupChannelScope(sel)

		case path.ConversationPostsCategory:
			scopes = selTD.GroupsBackupConversationScope(sel)

		case path.LibrariesCategory:
			scopes = selTD.GroupsBackupLibraryFolderScope(sel)
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
	runGroupsListCmdTest(suite, path.ChannelMessagesCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_conversations() {
	runGroupsListCmdTest(suite, path.ConversationPostsCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_libraries() {
	runGroupsListCmdTest(suite, path.LibrariesCategory)
}

func runGroupsListCmdTest(suite *PreparedBackupGroupsE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
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
	runGroupsListSingleCmdTest(suite, path.ChannelMessagesCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_singleID_conversations() {
	runGroupsListSingleCmdTest(suite, path.ConversationPostsCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsListCmd_singleID_libraries() {
	runGroupsListSingleCmdTest(suite, path.LibrariesCategory)
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
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
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
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backup", "smarfs")
	cli.BuildCommandTree(cmd)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsDetailsCmd_channelMessages() {
	runGroupsDetailsCmdTest(suite, path.ChannelMessagesCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsDetailsCmd_conversations() {
	runGroupsDetailsCmdTest(suite, path.ConversationPostsCategory)
}

func (suite *PreparedBackupGroupsE2ESuite) TestGroupsDetailsCmd_libraries() {
	runGroupsDetailsCmdTest(suite, path.LibrariesCategory)
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
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
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
	dpnd      dependencies
	backupOps [3]operations.BackupOperation
}

func TestBackupDeleteGroupsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteGroupsE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteGroupsE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx, path.GroupsService)

	m365GroupID := tconfig.M365TeamID(t)
	groups := []string{m365GroupID}

	// some tests require an existing backup
	sel := selectors.NewGroupsBackup(groups)
	sel.Include(selTD.GroupsBackupChannelScope(sel))

	for i := 0; i < cap(suite.backupOps); i++ {
		backupOp, err := suite.dpnd.repo.NewBackup(ctx, sel.Selector)
		require.NoError(t, err, clues.ToCore(err))

		suite.backupOps[i] = backupOp

		err = suite.backupOps[i].Run(ctx)
		require.NoError(t, err, clues.ToCore(err))
	}
}

func (suite *BackupDeleteGroupsE2ESuite) TestGroupsBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--"+flags.BackupIDsFN,
		fmt.Sprintf("%s,%s",
			string(suite.backupOps[0].Results.BackupID),
			string(suite.backupOps[1].Results.BackupID)))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backups", string(suite.backupOps[0].Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteGroupsE2ESuite) TestGroupsBackupDeleteCmd_SingleID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--"+flags.BackupFN,
		string(suite.backupOps[2].Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backup", string(suite.backupOps[2].Results.BackupID))
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
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--"+flags.BackupIDsFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteGroupsE2ESuite) TestGroupsBackupDeleteCmd_NoBackupID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "groups",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	// empty backupIDs should error since no data provided
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
		"--"+flags.ConfigFileFN, configFile,
		"--"+flags.GroupFN, group,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
