package restore_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	rep "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

var (
	email    = path.EmailCategory
	contacts = path.ContactsCategory
	events   = path.EventsCategory
)

var backupDataSets = []path.CategoryType{email, contacts, events}

type RestoreExchangeE2ESuite struct {
	tester.Suite
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       repository.Repository
	m365UserID string
	backupOps  map[path.CategoryType]operations.BackupOperation
}

func TestRestoreExchangeE2ESuite(t *testing.T) {
	suite.Run(t, &RestoreExchangeE2ESuite{
		Suite: tester.NewE2ESuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *RestoreExchangeE2ESuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// aggregate required details
	suite.acct = tconfig.NewM365Account(t)
	suite.st = storeTD.NewPrefixedS3Storage(t)

	cfg, err := suite.st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tconfig.TestCfgAccountProvider: "M365",
		tconfig.TestCfgStorageProvider: "S3",
		tconfig.TestCfgPrefix:          cfg.Prefix,
	}
	suite.vpr, suite.cfgFP = tconfig.MakeTempTestConfigClone(t, force)

	suite.m365UserID = strings.ToLower(tconfig.M365UserID(t))

	var (
		users = []string{suite.m365UserID}
		ins   = idname.NewCache(map[string]string{suite.m365UserID: suite.m365UserID})
	)

	// init the repo first
	suite.repo, err = repository.Initialize(
		ctx,
		suite.acct,
		suite.st,
		control.Options{},
		rep.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	suite.backupOps = make(map[path.CategoryType]operations.BackupOperation)

	for _, set := range backupDataSets {
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

		bop, err := suite.repo.NewBackupWithLookup(ctx, sel.Selector, ins)
		require.NoError(t, err, clues.ToCore(err))

		err = bop.Run(ctx)
		require.NoError(t, err, clues.ToCore(err))

		suite.backupOps[set] = bop

		// sanity check, ensure we can find the backup and its details immediately
		_, err = suite.repo.Backup(ctx, string(bop.Results.BackupID))
		require.NoError(t, err, "retrieving recent backup by ID", clues.ToCore(err))

		_, _, errs := suite.repo.GetBackupDetails(ctx, string(bop.Results.BackupID))
		require.NoError(t, errs.Failure(), "retrieving recent backup details by ID", clues.ToCore(err))
		require.Empty(t, errs.Recovered(), "retrieving recent backup details by ID")
	}
}

func (suite *RestoreExchangeE2ESuite) TestExchangeRestoreCmd() {
	for _, set := range backupDataSets {
		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			ctx = config.SetViper(ctx, suite.vpr)

			defer flush()

			cmd := cliTD.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--"+flags.BackupFN, string(suite.backupOps[set].Results.BackupID))
			cli.BuildCommandTree(cmd)

			// run the command
			err := cmd.ExecuteContext(ctx)
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

func (suite *RestoreExchangeE2ESuite) TestExchangeRestoreCmd_badTimeFlags() {
	for _, set := range backupDataSets {
		if set == contacts {
			suite.T().Skip()
		}

		suite.Run(set.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			ctx = config.SetViper(ctx, suite.vpr)

			defer flush()

			var timeFilter string
			switch set {
			case email:
				timeFilter = "--" + flags.EmailReceivedAfterFN
			case events:
				timeFilter = "--" + flags.EventStartsAfterFN
			}

			cmd := cliTD.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--"+flags.BackupFN, string(suite.backupOps[set].Results.BackupID),
				timeFilter, "smarf")
			cli.BuildCommandTree(cmd)

			// run the command
			err := cmd.ExecuteContext(ctx)
			require.Error(t, err, clues.ToCore(err))
		})
	}
}

func (suite *RestoreExchangeE2ESuite) TestExchangeRestoreCmd_badBoolFlags() {
	for _, set := range backupDataSets {
		if set != events {
			suite.T().Skip()
		}

		suite.Run(set.String(), func() {
			t := suite.T()

			//nolint:forbidigo
			ctx := config.SetViper(context.Background(), suite.vpr)
			ctx, flush := tester.WithContext(t, ctx)
			defer flush()

			var timeFilter string
			switch set {
			case events:
				timeFilter = "--" + flags.EventRecursFN
			}

			cmd := cliTD.StubRootCmd(
				"restore", "exchange",
				"--config-file", suite.cfgFP,
				"--"+flags.BackupFN, string(suite.backupOps[set].Results.BackupID),
				timeFilter, "wingbat")
			cli.BuildCommandTree(cmd)

			// run the command
			err := cmd.ExecuteContext(ctx)
			require.Error(t, err, clues.ToCore(err))
		})
	}
}
