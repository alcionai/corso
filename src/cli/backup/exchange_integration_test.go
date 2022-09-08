package backup_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

var (
	email    = path.EmailCategory
	contacts = path.ContactsCategory
	events   = path.EventsCategory
)

var backupDataSets = []path.CategoryType{email, contacts, events}

// ---------------------------------------------------------------------------
// tests with no prior backup
// ---------------------------------------------------------------------------

type BackupExchangeIntegrationSuite struct {
	suite.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       *repository.Repository
	m365UserID string
}

func TestBackupExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupExchangeIntegrationSuite))
}

func (suite *BackupExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(t, err)

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}

	suite.vpr, suite.cfgFP, err = tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st)
	require.NoError(t, err)
}

func (suite *BackupExchangeIntegrationSuite) TestExchangeBackupCmd() {
	recorder := strings.Builder{}

	for _, set := range backupDataSets {
		recorder.Reset()

		suite.T().Run(set.String(), func(t *testing.T) {
			ctx := config.SetViper(tester.NewContext(), suite.vpr)

			cmd := tester.StubRootCmd(
				"backup", "create", "exchange",
				"--config-file", suite.cfgFP,
				"--user", suite.m365UserID,
				"--data", set.String())
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

			result := recorder.String()
			t.Log("backup results", result)

			// as an offhand check: the result should contain a string with the current hour
			assert.Contains(t, result, time.Now().UTC().Format("2006-01-02T15"))
			// and the m365 user id
			assert.Contains(t, result, suite.m365UserID)
		})
	}
}

// ---------------------------------------------------------------------------
// tests prepared with a previous backup
// ---------------------------------------------------------------------------

type PreparedBackupExchangeIntegrationSuite struct {
	suite.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       *repository.Repository
	m365UserID string
	backupOps  map[path.CategoryType]operations.BackupOperation
}

func TestPreparedBackupExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(PreparedBackupExchangeIntegrationSuite))
}

func (suite *PreparedBackupExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(t, err)

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}
	suite.vpr, suite.cfgFP, err = tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st)
	require.NoError(t, err)

	suite.backupOps = make(map[path.CategoryType]operations.BackupOperation)

	for _, set := range backupDataSets {
		var (
			sel    = selectors.NewExchangeBackup()
			scopes []selectors.ExchangeScope
		)

		switch set {
		case email:
			scopes = sel.MailFolders([]string{suite.m365UserID}, []string{"Inbox"})

		case contacts:
			scopes = sel.ContactFolders([]string{suite.m365UserID}, selectors.Any())

		case events:
			scopes = sel.EventCalendars([]string{suite.m365UserID}, selectors.Any())
		}

		sel.Include(scopes)

		bop, err := suite.repo.NewBackup(
			ctx,
			sel.Selector,
			control.NewOptions(false))
		require.NoError(t, bop.Run(ctx))
		require.NoError(t, err)

		suite.backupOps[set] = bop

		// sanity check, ensure we can find the backup and its details immediately
		_, err = suite.repo.Backup(ctx, bop.Results.BackupID)
		require.NoError(t, err, "retrieving recent backup by ID")
		_, _, err = suite.repo.BackupDetails(ctx, string(bop.Results.BackupID))
		require.NoError(t, err, "retrieving recent backup details by ID")
	}
}

func (suite *PreparedBackupExchangeIntegrationSuite) TestExchangeListCmd() {
	recorder := strings.Builder{}

	for _, set := range backupDataSets {
		recorder.Reset()

		suite.T().Run(set.String(), func(t *testing.T) {
			ctx := config.SetViper(tester.NewContext(), suite.vpr)

			cmd := tester.StubRootCmd(
				"backup", "list", "exchange",
				"--config-file", suite.cfgFP)
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

			// compare the output
			result := recorder.String()
			assert.Contains(t, result, suite.backupOps[set].Results.BackupID)
		})
	}
}

func (suite *PreparedBackupExchangeIntegrationSuite) TestExchangeDetailsCmd() {
	recorder := strings.Builder{}

	for _, set := range backupDataSets {
		recorder.Reset()

		suite.T().Run(set.String(), func(t *testing.T) {
			ctx := config.SetViper(tester.NewContext(), suite.vpr)
			bID := suite.backupOps[set].Results.BackupID

			// fetch the details from the repo first
			deets, _, err := suite.repo.BackupDetails(ctx, string(bID))
			require.NoError(t, err)

			cmd := tester.StubRootCmd(
				"backup", "details", "exchange",
				"--config-file", suite.cfgFP,
				"--backup", string(bID))
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

			// compare the output
			result := recorder.String()

			for i, ent := range deets.Entries {
				t.Run(fmt.Sprintf("detail %d", i), func(t *testing.T) {
					assert.Contains(t, result, ent.RepoRef)
				})
			}
		})
	}
}

// ---------------------------------------------------------------------------
// tests for deleting backups
// ---------------------------------------------------------------------------

type BackupDeleteExchangeIntegrationSuite struct {
	suite.Suite
	acct     account.Account
	st       storage.Storage
	vpr      *viper.Viper
	cfgFP    string
	repo     *repository.Repository
	backupOp operations.BackupOperation
}

func TestBackupDeleteExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupDeleteExchangeIntegrationSuite))
}

func (suite *BackupDeleteExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(t, err)

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}
	suite.vpr, suite.cfgFP, err = tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx := config.SetViper(tester.NewContext(), suite.vpr)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st)
	require.NoError(t, err)

	m365UserID := tester.M365UserID(t)

	// some tests require an existing backup
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders([]string{m365UserID}, []string{"Inbox"}))

	suite.backupOp, err = suite.repo.NewBackup(
		ctx,
		sel.Selector,
		control.NewOptions(false))
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)
}

func (suite *BackupDeleteExchangeIntegrationSuite) TestExchangeBackupDeleteCmd() {
	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	t := suite.T()

	cmd := tester.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = tester.StubRootCmd(
		"backup", "details", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	require.Error(t, cmd.ExecuteContext(ctx))
}

func (suite *BackupDeleteExchangeIntegrationSuite) TestExchangeBackupDeleteCmd_UnknownID() {
	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	t := suite.T()

	cmd := tester.StubRootCmd(
		"backup", "delete", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", uuid.NewString())
	cli.BuildCommandTree(cmd)

	// unknown backupIDs should error since the modelStore can't find the backup
	require.Error(t, cmd.ExecuteContext(ctx))
}
