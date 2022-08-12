package backup_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli"
	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/repository"
)

type ExchangeIntegrationSuite struct {
	suite.Suite
}

func TestExchangeIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLITests,
		tester.CorsoCLIBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ExchangeIntegrationSuite))
}

func (suite *ExchangeIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(
		append(
			tester.AWSStorageCredEnvs,
			tester.M365AcctCredEnvs...,
		)...,
	)
	require.NoError(suite.T(), err)
}

func (suite *ExchangeIntegrationSuite) TestExchangeBackupCmd() {
	ctx := tester.NewContext()
	t := suite.T()

	acct := tester.NewM365Account(t)
	st := tester.NewPrefixedS3Storage(t)
	cfg, err := st.S3Config()
	require.NoError(t, err)

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}
	vpr, configFP, err := tester.MakeTempTestConfigClone(t, force)
	require.NoError(t, err)

	ctx = config.SetViper(ctx, vpr)

	// init the repo first
	_, err = repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	m365UserID := tester.M365UserID(t)

	// then test it
	cmd := tester.StubRootCmd(
		"backup", "create", "exchange",
		"--config-file", configFP,
		"--user", m365UserID,
		"--data", "email",
	)
	cli.BuildCommandTree(cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))
}
