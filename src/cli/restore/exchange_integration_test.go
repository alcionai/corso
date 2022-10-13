package restore_test

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
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

// TODO: bring back event restore testing when they no longer produce
// notification emails.  Currently, the duplication causes our tests
// dataset to grow until timeouts occur.
// var backupDataSets = []path.CategoryType{email, contacts, events}

var backupDataSets = []path.CategoryType{contacts, email}

type RestoreExchangeIntegrationSuite struct {
	suite.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	backupOps  map[path.CategoryType]operations.BackupOperation
}

func TestRestoreExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIRestoreTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RestoreExchangeIntegrationSuite))
}

func (suite *RestoreExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs,
	)
	require.NoError(t, err)

	// aggregate required details
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

	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)

	suite.backupOps = make(map[path.CategoryType]operations.BackupOperation)

	for _, set := range backupDataSets {
		var (
			sel    = selectors.NewExchangeBackup()
			scopes []selectors.ExchangeScope
		)

		switch set {
		case email:
			scopes = sel.MailFolders([]string{suite.m365UserID}, []string{exchange.DefaultMailFolder})

		case contacts:
			scopes = sel.ContactFolders([]string{suite.m365UserID}, []string{exchange.DefaultContactFolder})

		case events:
			scopes = sel.EventCalendars([]string{suite.m365UserID}, []string{exchange.DefaultCalendar})
		}

		sel.Include(scopes)

		bop, err := suite.repo.NewBackup(ctx, sel.Selector)
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

func (suite *RestoreExchangeIntegrationSuite) TestExchangeRestoreCmd() {
	for _, set := range backupDataSets {
		suite.T().Run(set.String(), func(t *testing.T) {
			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)

			defer flush()

			cmd := tester.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--backup", string(suite.backupOps[set].Results.BackupID))
			cli.BuildCommandTree(cmd)

			// run the command
			require.NoError(t, cmd.ExecuteContext(ctx))
		})
	}
}

func (suite *RestoreExchangeIntegrationSuite) TestExchangeRestoreCmd_badTimeFlags() {
	for _, set := range backupDataSets {
		if set == contacts {
			suite.T().Skip()
		}

		suite.T().Run(set.String(), func(t *testing.T) {
			ctx, flush := tester.NewContext()
			ctx = config.SetViper(ctx, suite.vpr)

			defer flush()

			var timeFilter string
			switch set {
			case email:
				timeFilter = "--email-received-after"
			case events:
				timeFilter = "--event-starts-after"
			}

			cmd := tester.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--backup", string(suite.backupOps[set].Results.BackupID),
				timeFilter, "smarf")
			cli.BuildCommandTree(cmd)

			// run the command
			require.Error(t, cmd.ExecuteContext(ctx))
		})
	}
}

func (suite *RestoreExchangeIntegrationSuite) TestExchangeRestoreCmd_badBoolFlags() {
	for _, set := range backupDataSets {
		if set != events {
			suite.T().Skip()
		}

		suite.T().Run(set.String(), func(t *testing.T) {
			ctx := config.SetViper(context.Background(), suite.vpr)
			ctx, flush := tester.WithContext(ctx)
			defer flush()

			var timeFilter string
			switch set {
			case events:
				timeFilter = "--event-recurs"
			}

			cmd := tester.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--backup", string(suite.backupOps[set].Results.BackupID),
				timeFilter, "wingbat")
			cli.BuildCommandTree(cmd)

			// run the command
			require.Error(t, cmd.ExecuteContext(ctx))
		})
	}
}
