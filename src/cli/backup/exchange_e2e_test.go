package backup_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

var (
	email    = path.EmailCategory
	contacts = path.ContactsCategory
	events   = path.EventsCategory
)

// ---------------------------------------------------------------------------
// tests with no backups
// ---------------------------------------------------------------------------

type NoBackupExchangeE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	recorder   strings.Builder
}

func TestNoBackupExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &NoBackupExchangeE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
	)})
}

func (suite *NoBackupExchangeE2ESuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.recorder = recorder
	suite.cfgFP = cfgFilePath
	suite.m365UserID = tester.M365UserID(t)
}

func (suite *NoBackupExchangeE2ESuite) TestExchangeBackupListCmd_empty() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.recorder.Reset()

	cmd := tester.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.cfgFP)
	cli.BuildCommandTree(cmd)

	cmd.SetErr(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := suite.recorder.String()

	// as an offhand check: the result should contain the m365 user id
	assert.Equal(t, "No backups available\n", result)
}

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupExchangeE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
}

func TestBackupExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &BackupExchangeE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
	)})
}

func (suite *BackupExchangeE2ESuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	acct, st, repo, vpr, _, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.cfgFP = cfgFilePath
	suite.m365UserID = tester.M365UserID(t)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_email() {
	runExchangeBackupCategoryTest(suite, "email")
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_contacts() {
	runExchangeBackupCategoryTest(suite, "contacts")
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_events() {
	runExchangeBackupCategoryTest(suite, "events")
}

func runExchangeBackupCategoryTest(suite *BackupExchangeE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd, ctx := buildExchangeBackupCmd(ctx, suite.cfgFP, suite.m365UserID, category, &recorder)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 user id
	assert.Contains(t, result, suite.m365UserID)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_ServiceNotEnables_email() {
	runExchangeBackupServiceNotEnabledTest(suite, "email")
}

func runExchangeBackupServiceNotEnabledTest(suite *BackupExchangeE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	// run the command

	cmd, ctx := buildExchangeBackupCmd(
		ctx,
		suite.cfgFP,
		fmt.Sprintf("testevents@10rqc2.onmicrosoft.com,%s", suite.m365UserID),
		category,
		&recorder)
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	result := recorder.String()
	t.Log("backup results", result)

	// as an offhand check: the result should contain the m365 user id
	assert.Contains(t, result, suite.m365UserID)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_email() {
	runExchangeBackupUserNotFoundTest(suite, "email")
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_contacts() {
	runExchangeBackupUserNotFoundTest(suite, "contacts")
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_userNotFound_events() {
	runExchangeBackupUserNotFoundTest(suite, "events")
}

func runExchangeBackupUserNotFoundTest(suite *BackupExchangeE2ESuite, category string) {
	recorder := strings.Builder{}
	recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd, ctx := buildExchangeBackupCmd(ctx, suite.cfgFP, "foo@not-there.com", category, &recorder)

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
// tests prepared with a previous backup
// ---------------------------------------------------------------------------

type PreparedBackupExchangeE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	backupOps  map[path.CategoryType]string
	recorder   strings.Builder
}

func TestPreparedBackupExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &PreparedBackupExchangeE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
	)})
}

func (suite *PreparedBackupExchangeE2ESuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	acct, st, repo, vpr, recorder, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.recorder = recorder
	suite.cfgFP = cfgFilePath
	suite.m365UserID = tester.M365UserID(t)
	suite.backupOps = make(map[path.CategoryType]string)

	var (
		users = []string{suite.m365UserID}
		ins   = idname.NewCache(map[string]string{suite.m365UserID: suite.m365UserID})
	)

	for _, set := range []path.CategoryType{email, contacts, events} {
		var (
			sel    = selectors.NewExchangeBackup(users)
			scopes []selectors.ExchangeScope
		)

		switch set {
		case email:
			scopes = sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch())

		case contacts:
			scopes = sel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch())

		case events:
			scopes = sel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch())
		}

		sel.Include(scopes)

		bop, err := suite.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
		require.NoError(t, err, clues.ToCore(err))

		err = bop.Run(ctx)
		require.NoError(t, err, clues.ToCore(err))

		bIDs := string(bop.Results.BackupID)

		// sanity check, ensure we can find the backup and its details immediately
		b, err := suite.repo.Backup(ctx, string(bop.Results.BackupID))
		require.NoError(t, err, "retrieving recent backup by ID")
		require.Equal(t, bIDs, string(b.ID), "repo backup matches results id")

		_, b, errs := suite.repo.GetBackupDetails(ctx, bIDs)
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
	suite.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.cfgFP)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.recorder.String()
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
	suite.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	bID := suite.backupOps[category]

	cmd := tester.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(bID))
	cli.BuildCommandTree(cmd)

	cmd.SetOut(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.recorder.String()
	assert.Contains(t, result, bID)
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_badID() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "list", "exchange",
		"--config-file", suite.cfgFP,
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
	suite.recorder.Reset()

	t := suite.T()

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	bID := suite.backupOps[category]

	// fetch the details from the repo first
	deets, _, errs := suite.repo.GetBackupDetails(ctx, string(bID))
	require.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
	require.Empty(t, errs.Recovered())

	cmd := tester.StubRootCmd(
		"backup", "details", "exchange",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, string(bID))
	cli.BuildCommandTree(cmd)
	cmd.SetOut(&suite.recorder)

	ctx = print.SetRootCmd(ctx, cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// compare the output
	result := suite.recorder.String()

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
	acct     account.Account
	st       storage.Storage
	vpr      *viper.Viper
	cfgFP    string
	repo     repository.Repository
	backupOp operations.BackupOperation
}

func TestBackupDeleteExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &BackupDeleteExchangeE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
		),
	})
}

func (suite *BackupDeleteExchangeE2ESuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	acct, st, repo, vpr, _, cfgFilePath := prepM365Test(t, ctx)

	suite.acct = acct
	suite.st = st
	suite.repo = repo
	suite.vpr = vpr
	suite.cfgFP = cfgFilePath

	m365UserID := tester.M365UserID(t)
	users := []string{m365UserID}

	// some tests require an existing backup
	sel := selectors.NewExchangeBackup(users)
	sel.Include(sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))

	backupOp, err := suite.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOp = backupOp

	err = suite.backupOp.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteExchangeE2ESuite) TestExchangeBackupDeleteCmd() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	err := cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = tester.StubRootCmd(
		"backup", "details", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BackupDeleteExchangeE2ESuite) TestExchangeBackupDeleteCmd_UnknownID() {
	t := suite.T()
	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	cmd := tester.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.cfgFP,
		"--"+utils.BackupFN, uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func buildExchangeBackupCmd(
	ctx context.Context,
	configFile, user, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := tester.StubRootCmd(
		"backup", "create", "exchange",
		"--config-file", configFile,
		"--"+utils.UserFN, user,
		"--"+utils.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
