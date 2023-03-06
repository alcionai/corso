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

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
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

var backupDataSets = []path.CategoryType{email, contacts, events}

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
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests)})
}

func (suite *NoBackupExchangeE2ESuite) SetupSuite() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)
	suite.recorder = strings.Builder{}

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}

	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx = config.SetViper(ctx, suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)
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
	require.NoError(t, cmd.ExecuteContext(ctx))

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
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests)})
}

func (suite *BackupExchangeE2ESuite) SetupSuite() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

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

	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx = config.SetViper(ctx, suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd() {
	recorder := strings.Builder{}

	for _, set := range backupDataSets {
		recorder.Reset()

		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)
			defer flush()

			cmd := tester.StubRootCmd(
				"backup", "create", "exchange",
				"--config-file", suite.cfgFP,
				"--"+utils.UserFN, suite.m365UserID,
				"--"+utils.DataFN, set.String())
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

			result := recorder.String()
			t.Log("backup results", result)

			// as an offhand check: the result should contain the m365 user id
			assert.Contains(t, result, suite.m365UserID)
		})
	}
}

func (suite *BackupExchangeE2ESuite) TestExchangeBackupCmd_UserNotInTenant() {
	recorder := strings.Builder{}

	for _, set := range backupDataSets {
		recorder.Reset()

		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)
			defer flush()

			cmd := tester.StubRootCmd(
				"backup", "create", "exchange",
				"--config-file", suite.cfgFP,
				"--"+utils.UserFN, "foo@nothere.com",
				"--"+utils.DataFN, set.String())
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			err := cmd.ExecuteContext(ctx)
			require.Error(t, err)
			assert.Contains(
				t,
				err.Error(),
				"not found within tenant", "error missing user not found")
			assert.NotContains(t, err.Error(), "runtime error", "panic happened")

			t.Logf("backup error message: %s", err.Error())

			result := recorder.String()
			t.Log("backup results", result)
		})
	}
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
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests)})
}

func (suite *PreparedBackupExchangeE2ESuite) SetupSuite() {
	t := suite.T()

	// prepare common details
	suite.acct = tester.NewM365Account(t)
	suite.st = tester.NewPrefixedS3Storage(t)
	suite.recorder = strings.Builder{}

	cfg, err := suite.st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}
	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)

	suite.backupOps = make(map[path.CategoryType]string)

	users := []string{suite.m365UserID}

	for _, set := range backupDataSets {
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

		bop, err := suite.repo.NewBackup(ctx, sel.Selector)
		require.NoError(t, bop.Run(ctx))
		require.NoError(t, err)

		bIDs := string(bop.Results.BackupID)

		// sanity check, ensure we can find the backup and its details immediately
		b, err := suite.repo.Backup(ctx, bop.Results.BackupID)
		require.NoError(t, err, "retrieving recent backup by ID")
		require.Equal(t, bIDs, string(b.ID), "repo backup matches results id")
		_, b, errs := suite.repo.BackupDetails(ctx, bIDs)
		require.NoError(t, errs.Failure(), "retrieving recent backup details by ID")
		require.Empty(t, errs.Recovered(), "retrieving recent backup details by ID")
		require.Equal(t, bIDs, string(b.ID), "repo details matches results id")

		suite.backupOps[set] = string(b.ID)
	}
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd() {
	for _, set := range backupDataSets {
		suite.recorder.Reset()

		suite.Run(set.String(), func() {
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
			require.NoError(t, cmd.ExecuteContext(ctx))

			// compare the output
			result := suite.recorder.String()
			assert.Contains(t, result, suite.backupOps[set])
		})
	}
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_singleID() {
	for _, set := range backupDataSets {
		suite.recorder.Reset()

		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)
			defer flush()

			bID := suite.backupOps[set]

			cmd := tester.StubRootCmd(
				"backup", "list", "exchange",
				"--config-file", suite.cfgFP,
				"--backup", string(bID))
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&suite.recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

			// compare the output
			result := suite.recorder.String()
			assert.Contains(t, result, bID)
		})
	}
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeListCmd_badID() {
	for _, set := range backupDataSets {
		suite.Run(set.String(), func() {
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
			require.Error(t, cmd.ExecuteContext(ctx))
		})
	}
}

func (suite *PreparedBackupExchangeE2ESuite) TestExchangeDetailsCmd() {
	for _, set := range backupDataSets {
		suite.recorder.Reset()

		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)
			defer flush()

			bID := suite.backupOps[set]

			// fetch the details from the repo first
			deets, _, errs := suite.repo.BackupDetails(ctx, string(bID))
			require.NoError(t, errs.Failure())
			require.Empty(t, errs.Recovered())

			cmd := tester.StubRootCmd(
				"backup", "details", "exchange",
				"--config-file", suite.cfgFP,
				"--"+utils.BackupFN, string(bID))
			cli.BuildCommandTree(cmd)

			cmd.SetOut(&suite.recorder)

			ctx = print.SetRootCmd(ctx, cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))

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

			// At least the prefix of the path should be encoded as folders.
			assert.Greater(t, foundFolders, 4)
		})
	}
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
			tester.CorsoCITests,
			tester.CorsoCLITests,
			tester.CorsoCLIBackupTests),
	})
}

func (suite *BackupDeleteExchangeE2ESuite) SetupSuite() {
	t := suite.T()

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
	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	ctx, flush := tester.NewContext()
	ctx = config.SetViper(ctx, suite.vpr)

	defer flush()

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)

	m365UserID := tester.M365UserID(t)
	users := []string{m365UserID}

	// some tests require an existing backup
	sel := selectors.NewExchangeBackup(users)
	sel.Include(sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))

	suite.backupOp, err = suite.repo.NewBackup(ctx, sel.Selector)
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)
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
	require.NoError(t, cmd.ExecuteContext(ctx))

	// a follow-up details call should fail, due to the backup ID being deleted
	cmd = tester.StubRootCmd(
		"backup", "details", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	require.Error(t, cmd.ExecuteContext(ctx))
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
	require.Error(t, cmd.ExecuteContext(ctx))
}
