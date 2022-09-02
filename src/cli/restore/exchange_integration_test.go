package restore_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	email    = "email"
	contacts = "contacts"
	events   = "events"
)

var backupDataSets = []string{email, contacts, events}

type RestoreExchangeIntegrationSuite struct {
	suite.Suite
	dataSet    string
	acct       account.Account
	st         storage.Storage
	vpr        *viper.Viper
	cfgFP      string
	repo       *repository.Repository
	m365UserID string
	backupOp   operations.BackupOperation
}

func TestRestoreExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIRestoreTests,
	); err != nil {
		t.Skip(err)
	}

	for _, set := range backupDataSets {
		s := new(RestoreExchangeIntegrationSuite)
		s.dataSet = set
		suite.Run(t, s)
	}
}

func (suite *RestoreExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()
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

	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	suite.m365UserID = tester.M365UserID(t)

	// init the repo first
	suite.repo, err = repository.Initialize(ctx, suite.acct, suite.st)
	require.NoError(t, err)

	// restoration requires an existing backup
	sel := selectors.NewExchangeBackup()

	var scopes []selectors.ExchangeScope

	switch suite.dataSet {
	case email:
		scopes = sel.MailFolders([]string{suite.m365UserID}, []string{"Inbox"})

	case contacts:
		scopes = sel.ContactFolders([]string{suite.m365UserID}, selectors.Any())

	case events:
		scopes = sel.EventCalendars([]string{suite.m365UserID}, selectors.Any())
	}

	sel.Include(scopes)

	suite.backupOp, err = suite.repo.NewBackup(
		ctx,
		sel.Selector,
		control.NewOptions(false))
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)
}

func (suite *RestoreExchangeIntegrationSuite) TestExchangeRestoreCmd() {
	suite.T().Run(suite.dataSet, func(t *testing.T) {
		ctx := config.SetViper(tester.NewContext(), suite.vpr)

		cmd := tester.StubRootCmd(
			"restore", "exchange",
			"--config-file", suite.cfgFP,
			"--backup", string(suite.backupOp.Results.BackupID))
		cli.BuildCommandTree(cmd)

		// run the command
		require.NoError(t, cmd.ExecuteContext(ctx))
	})
}
