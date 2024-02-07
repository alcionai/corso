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

type NoBackupTeamsChatsE2ESuite struct {
	tester.Suite
	dpnd dependencies
	m365 its.M365IntgTestSetup
}

func TestNoBackupTeamsChatsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupTeamsChatsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *NoBackupTeamsChatsE2ESuite) SetupSuite() {
	t := suite.T()
	t.Skip("not fully implemented")

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.TeamsChatsService)
}

func (suite *NoBackupTeamsChatsE2ESuite) TestTeamsChatsBackupListCmd_noBackups() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()

	// as an offhand check: the result should contain the m365 teamschat id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupTeamsChatsE2ESuite struct {
	tester.Suite
	dpnd dependencies
	m365 its.M365IntgTestSetup
}

func TestBackupTeamsChatsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupTeamsChatsE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *BackupTeamsChatsE2ESuite) SetupSuite() {
	t := suite.T()
	t.Skip("not fully implemented")

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.TeamsChatsService)
}

func (suite *BackupTeamsChatsE2ESuite) TestTeamsChatsBackupCmd_chats() {
	runTeamsChatsBackupCategoryTest(suite, flags.DataChats)
}

func runTeamsChatsBackupCategoryTest(suite *BackupTeamsChatsE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildTeamsChatsBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		suite.m365.User.ID,
		category,
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)
}

func (suite *BackupTeamsChatsE2ESuite) TestTeamsChatsBackupCmd_teamschatNotFound_chats() {
	runTeamsChatsBackupTeamsChatNotFoundTest(suite, flags.DataChats)
}

func runTeamsChatsBackupTeamsChatNotFoundTest(suite *BackupTeamsChatsE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildTeamsChatsBackupCmd(
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

func (suite *BackupTeamsChatsE2ESuite) TestBackupCreateTeamsChats_badAzureClientIDFlag() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "chats",
		"--teamschat", suite.m365.User.ID,
		"--azure-client-id", "invalid-value")
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupTeamsChatsE2ESuite) TestBackupCreateTeamsChats_fromConfigFile() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "chats",
		"--teamschat", suite.m365.User.ID,
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

// AWS flags
func (suite *BackupTeamsChatsE2ESuite) TestBackupCreateTeamsChats_badAWSFlags() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "chats",
		"--teamschat", suite.m365.User.ID,
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

type PreparedBackupTeamsChatsE2ESuite struct {
	tester.Suite
	dpnd      dependencies
	backupOps map[path.CategoryType]string
	m365      its.M365IntgTestSetup
}

func TestPreparedBackupTeamsChatsE2ESuite(t *testing.T) {
	suite.Run(t, &PreparedBackupTeamsChatsE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *PreparedBackupTeamsChatsE2ESuite) SetupSuite() {
	t := suite.T()
	t.Skip("not fully implemented")

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.m365 = its.GetM365(t)
	suite.dpnd = prepM365Test(t, ctx, path.TeamsChatsService)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		teamschats = []string{suite.m365.User.ID}
		ins        = idname.NewCache(map[string]string{suite.m365.User.ID: suite.m365.User.ID})
		cats       = []path.CategoryType{
			path.ChatsCategory,
		}
	)

	for _, set := range cats {
		var (
			sel    = selectors.NewTeamsChatsBackup(teamschats)
			scopes []selectors.TeamsChatsScope
		)

		switch set {
		case path.ChatsCategory:
			scopes = selTD.TeamsChatsBackupChatScope(sel)
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

func (suite *PreparedBackupTeamsChatsE2ESuite) TestTeamsChatsListCmd_chats() {
	runTeamsChatsListCmdTest(suite, path.ChatsCategory)
}

func runTeamsChatsListCmdTest(suite *PreparedBackupTeamsChatsE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "chats",
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

func (suite *PreparedBackupTeamsChatsE2ESuite) TestTeamsChatsListCmd_singleID_chats() {
	runTeamsChatsListSingleCmdTest(suite, path.ChatsCategory)
}

func runTeamsChatsListSingleCmdTest(suite *PreparedBackupTeamsChatsE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	bID := suite.backupOps[category]

	cmd := cliTD.StubRootCmd(
		"backup", "list", "chats",
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

func (suite *PreparedBackupTeamsChatsE2ESuite) TestTeamsChatsListCmd_badID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backup", "smarfs")
	cli.BuildCommandTree(cmd)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *PreparedBackupTeamsChatsE2ESuite) TestTeamsChatsDetailsCmd_chats() {
	runTeamsChatsDetailsCmdTest(suite, path.ChatsCategory)
}

func runTeamsChatsDetailsCmdTest(suite *PreparedBackupTeamsChatsE2ESuite, category path.CategoryType) {
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
		"backup", "details", "chats",
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
		// Skip folders as they don't mean anything to the end teamschat.
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

type BackupDeleteTeamsChatsE2ESuite struct {
	tester.Suite
	dpnd      dependencies
	backupOps [3]operations.BackupOperation
}

func TestBackupDeleteTeamsChatsE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteTeamsChatsE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteTeamsChatsE2ESuite) SetupSuite() {
	t := suite.T()
	t.Skip("not fully implemented")

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx, path.TeamsChatsService)

	m365TeamsChatID := tconfig.M365TeamID(t)
	teamschats := []string{m365TeamsChatID}

	// some tests require an existing backup
	sel := selectors.NewTeamsChatsBackup(teamschats)
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	for i := 0; i < cap(suite.backupOps); i++ {
		backupOp, err := suite.dpnd.repo.NewBackup(ctx, sel.Selector)
		require.NoError(t, err, clues.ToCore(err))

		suite.backupOps[i] = backupOp

		err = suite.backupOps[i].Run(ctx)
		require.NoError(t, err, clues.ToCore(err))
	}
}

func (suite *BackupDeleteTeamsChatsE2ESuite) TestTeamsChatsBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "chats",
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
		"backup", "details", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backups", string(suite.backupOps[0].Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteTeamsChatsE2ESuite) TestTeamsChatsBackupDeleteCmd_SingleID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--"+flags.BackupFN,
		string(suite.backupOps[2].Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--backup", string(suite.backupOps[2].Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteTeamsChatsE2ESuite) TestTeamsChatsBackupDeleteCmd_UnknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath,
		"--"+flags.BackupIDsFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteTeamsChatsE2ESuite) TestTeamsChatsBackupDeleteCmd_NoBackupID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "chats",
		"--"+flags.ConfigFileFN, suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	// empty backupIDs should error since no data provided
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func buildTeamsChatsBackupCmd(
	ctx context.Context,
	configFile, resource, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := cliTD.StubRootCmd(
		"backup", "create", "chats",
		"--"+flags.ConfigFileFN, configFile,
		"--"+flags.UserFN, resource,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
