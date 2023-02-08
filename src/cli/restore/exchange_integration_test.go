package restore_test

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
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
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIRestoreTests)

	suite.Run(t, new(RestoreExchangeIntegrationSuite))
}

func (suite *RestoreExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	tester.MustGetEnvSets(t, tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs)

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
	suite.vpr, suite.cfgFP = tester.MakeTempTestConfigClone(t, force)

	suite.m365UserID = tester.M365UserID(t)
	users := []string{suite.m365UserID}

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st, control.Options{})
	require.NoError(t, err)

	suite.backupOps = make(map[path.CategoryType]operations.BackupOperation)

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

		suite.backupOps[set] = bop

		// sanity check, ensure we can find the backup and its details immediately
		_, err = suite.repo.Backup(ctx, bop.Results.BackupID)
		require.NoError(t, err, "retrieving recent backup by ID")

		_, _, errs := suite.repo.BackupDetails(ctx, string(bop.Results.BackupID))
		require.NoError(t, errs.Err(), "retrieving recent backup details by ID")
		require.Empty(t, errs.Errs(), "retrieving recent backup details by ID")
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
				"--"+utils.BackupFN, string(suite.backupOps[set].Results.BackupID))
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
				timeFilter = "--" + utils.EmailReceivedAfterFN
			case events:
				timeFilter = "--" + utils.EventStartsAfterFN
			}

			cmd := tester.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--"+utils.BackupFN, string(suite.backupOps[set].Results.BackupID),
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
			//nolint:forbidigo
			ctx := config.SetViper(context.Background(), suite.vpr)
			ctx, flush := tester.WithContext(ctx)
			defer flush()

			var timeFilter string
			switch set {
			case events:
				timeFilter = "--" + utils.EventRecursFN
			}

			cmd := tester.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--"+utils.BackupFN, string(suite.backupOps[set].Results.BackupID),
				timeFilter, "wingbat")
			cli.BuildCommandTree(cmd)

			// run the command
			require.Error(t, cmd.ExecuteContext(ctx))
		})
	}
}
