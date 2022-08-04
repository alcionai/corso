package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli"
	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/internal/tester"
)

// ---------------------------------------------------------------------------------------------------------
// Integration
// ---------------------------------------------------------------------------------------------------------

type S3IntegrationSuite struct {
	suite.Suite
}

func TestS3IntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoCLIRepoTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(S3IntegrationSuite))
}

func (suite *S3IntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(
		append(
			tester.AWSStorageCredEnvs,
			tester.M365AcctCredEnvs...,
		)...,
	)
	require.NoError(suite.T(), err)
}

func (suite *S3IntegrationSuite) TestInitS3Cmd() {
	ctx := tester.NewContext()
	t := suite.T()

	st, err := tester.NewPrefixedS3Storage(t)
	require.NoError(t, err)
	cfg, err := st.S3Config()
	require.NoError(t, err)

	vpr, configFP, err := tester.MakeTempTestConfigClone(t)
	require.NoError(t, err)
	ctx = config.SetViper(ctx, vpr)

	require.NoError(t, err)

	cmd := tester.StubRootCmd(
		"repo", "init", "s3",
		"--config-file", configFP,
		"--bucket", cfg.Bucket,
		"--prefix", cfg.Prefix)
	cli.BuildCommandTree(cmd)

	// run the command
	require.NoError(t, cmd.ExecuteContext(ctx))
}
