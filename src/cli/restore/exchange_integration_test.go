package restore_test

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli"
	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/internal/operations"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/storage"
)

type RestoreExchangeIntegrationSuite struct {
	suite.Suite
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
	suite.Run(t, new(RestoreExchangeIntegrationSuite))
}

func (suite *RestoreExchangeIntegrationSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
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
	sel.Include(sel.MailFolders([]string{suite.m365UserID}, []string{"Inbox"}))
	suite.backupOp, err = suite.repo.NewBackup(
		ctx,
		sel.Selector,
		control.NewOptions(false))
	require.NoError(t, suite.backupOp.Run(ctx))
	require.NoError(t, err)

	time.Sleep(3 * time.Second)
}

func (suite *RestoreExchangeIntegrationSuite) TestExchangeRestoreCmd() {
	ctx := config.SetViper(tester.NewContext(), suite.vpr)
	t := suite.T()

	cmd := tester.StubRootCmd(
		"restore", "exchange",
		"--config-file", suite.cfgFP,
		"--backup", string(suite.backupOp.Results.BackupID))
	cli.BuildCommandTree(cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))
}
