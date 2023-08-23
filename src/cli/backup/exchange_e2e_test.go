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
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

var (
	email    = path.EmailCategory
	contacts = path.ContactsCategory
	events   = path.EventsCategory
)

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupExchangeE2ESuite struct {
	tester.Suite
	dpnd dependencies
	its  intgTesterSetup
}

func TestBackupExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &BackupExchangeE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
	)})
}

func (suite *BackupExchangeE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_email() {
	runExchangeBackupCategoryTest(suite, email)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_contacts() {
	runExchangeBackupCategoryTest(suite, contacts)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_events() {
	runExchangeBackupCategoryTest(suite, events)
}

func runExchangeBackupCategoryTest(suite *BackupExchangeE2ESuite, category path.CategoryType) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildExchangeBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		suite.its.user.ID,
		category.String(),
		&recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 user id
	assert.Contains(t, result, suite.its.user.ID)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_ServiceNotEnabled_email() {
	runExchangeBackupServiceNotEnabledTest(suite, email)
}

func runExchangeBackupServiceNotEnabledTest(suite *BackupExchangeE2ESuite, category path.CategoryType) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	// run the command

	cmd, ctx := buildExchangeBackupCmd(
		ctx,
		suite.dpnd.configFilePath,
		fmt.Sprintf("%s,%s", tconfig.UnlicensedM365UserID(suite.T()), suite.its.user.ID),
		category.String(),
		&recorder)
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 user id
	assert.Contains(t, result, suite.its.user.ID)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_email() {
	runExchangeBackupUserNotFoundTest(suite, email)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_contacts() {
	runExchangeBackupUserNotFoundTest(suite, contacts)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_events() {
	runExchangeBackupUserNotFoundTest(suite, events)
}

func runExchangeBackupUserNotFoundTest(suite *BackupExchangeE2ESuite, category path.CategoryType) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd, ctx := buildExchangeBackupCmd(
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
		"not found in tenant", "error missing user not found")
	assert.NotContains(t, err.Error(), "runtime error", "panic happened")

	t.Logf("backup error message: %s", err.Error())

	result := recorder.String()
	t.Log("backup results", result)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupListCmd_empty() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()

	// as an offhand check: the result should contain the m365 user id
	assert.True(t, strings.HasSuffix(result, "No backups available\n"))
}

func (suite *BackupExchangeE2ESuite) TestBackupCreateExchange_badAzureClientIDFlag() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "exchange",
		"--user", suite.its.user.ID,
		"--azure-client-id", "invalid-value",
	)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupExchangeE2ESuite) TestBackupCreateExchange_fromConfigFile() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "exchange",
		"--user", suite.its.user.ID,
		"--config-file", suite.dpnd.configFilePath)
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.dpnd.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.dpnd.recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 user id
	assert.Contains(t, result, suite.its.user.ID)
}

// AWS flags
func (suite *BackupExchangeE2ESuite) TestBackupCreateExchange_badAWSFlags() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	suite.dpnd.recorder.Reset()

	cmd := cliTD.StubRootCmd(
		"backup", "create", "exchange",
		"--user", suite.its.user.ID,
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

type PreparedBackupExchangeE2ESuite struct {
	tester.Suite
	dpnd      dependencies
	backupOps map[path.CategoryType]string
	its       intgTesterSetup
}

func TestPreparedBackupExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &PreparedBackupExchangeE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *PreparedBackupExchangeE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(t)
	suite.dpnd = prepM365Test(t, ctx)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		users = []string{suite.its.user.ID}
		ins   = idname.NewCache(map[string]string{suite.its.user.ID: suite.its.user.ID})
	)

	for _, set := range []path.CategoryType{email, contacts, events} {
		var (
			sel    = selectors.NewExchangeBackup(users)
			scopes []selectors.ExchangeScope
		)

		switch set {
		case email:
			scopes = sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch())

		case contacts:
			scopes = sel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch())

		case events:
			scopes = sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch())
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

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_email() {
	runExchangeListCmdTest(suite, email)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_contacts() {
	runExchangeListCmdTest(suite, contacts)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_events() {
	runExchangeListCmdTest(suite, events)
}

func runExchangeListCmdTest(suite *PreparedBackupExchangeE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "exchange",
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

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_singleID_email() {
	runExchangeListSingleCmdTest(suite, email)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_singleID_contacts() {
	runExchangeListSingleCmdTest(suite, contacts)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_singleID_events() {
	runExchangeListSingleCmdTest(suite, events)
}

func runExchangeListSingleCmdTest(suite *PreparedBackupExchangeE2ESuite, category path.CategoryType) {
	suite.dpnd.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	bID := suite.backupOps[category]

	cmd := cliTD.StubRootCmd(
		"backup", "list", "exchange",
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

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_badID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", "smarfs")
	cli.BuildCommandTree(cmd)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeDetailsCmd_email() {
	runExchangeDetailsCmdTest(suite, email)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeDetailsCmd_contacts() {
	runExchangeDetailsCmdTest(suite, contacts)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeDetailsCmd_events() {
	runExchangeDetailsCmdTest(suite, events)
}

func runExchangeDetailsCmdTest(suite *PreparedBackupExchangeE2ESuite, category path.CategoryType) {
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
		"backup", "details", "exchange",
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
		// Skip folders as they don't mean anything to the end user.
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

type BackupDeleteExchangeE2ESuite struct {
	tester.Suite
	dpnd     dependencies
	backupOp operations.BackupOperation
}

func TestBackupDeleteExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteExchangeE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupDeleteExchangeE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.dpnd = prepM365Test(t, ctx)

	m365UserID := tconfig.M365UserID(t)
	users := []string{m365UserID}

	// some tests require an existing backup
	sel := selectors.NewExchangeBackup(users)
	sel.Include(sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))

	backupOp, err := suite.dpnd.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteExchangeE2ESuite) TestExchangeBackupDeleteCmd() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = cliTD.StubRootCmd(
		"backup", "details", "exchange",
		"--config-file", suite.dpnd.configFilePath,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteExchangeE2ESuite) TestExchangeBackupDeleteCmd_UnknownID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	ctx = config.SetViper(ctx, suite.dpnd.vpr)

	defer flush()

	cmd := cliTD.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.dpnd.configFilePath,
		"--"+flags.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
