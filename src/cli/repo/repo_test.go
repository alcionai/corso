package repo_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/repo"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/config"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type RepoUnitSuite struct {
	tester.Suite
}

func TestRepoUnitSuite(t *testing.T) {
	suite.Run(t, &RepoUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RepoUnitSuite) TestAddRepoCommands() {
	t := suite.T()
	cmd := &cobra.Command{}

	repo.AddCommands(cmd)

	var found bool

	// This is the repo command.
	repoCmds := cmd.Commands()
	require.Len(t, repoCmds, 1)

	for _, c := range repoCmds[0].Commands() {
		if c.Use == repo.MaintenanceCommand {
			found = true
		}
	}

	assert.True(t, found, "looking for maintenance command")
}

type RepoE2ESuite struct {
	tester.Suite
}

func TestRepoE2ESuite(t *testing.T) {
	suite.Run(t, &RepoE2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *RepoE2ESuite) TestUpdatePassphraseCmd() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	st := storeTD.NewPrefixedS3Storage(t)
	sc, err := st.StorageConfig()
	require.NoError(t, err, clues.ToCore(err))

	cfg := sc.(*storage.S3Config)

	vpr, configFP := tconfig.MakeTempTestConfigClone(t, nil)

	ctx = config.SetViper(ctx, vpr)

	cmd := cliTD.StubRootCmd(
		"repo", "init", "s3",
		"--config-file", configFP,
		"--prefix", cfg.Prefix)

	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// connect with old passphrase
	cmd = cliTD.StubRootCmd(
		"repo", "connect", "s3",
		"--config-file", configFP,
		"--bucket", cfg.Bucket,
		"--prefix", cfg.Prefix)
	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	cmd = cliTD.StubRootCmd(
		"repo", "update-passphrase",
		"--config-file", configFP,
		"--new-passphrase", "newpass")
	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// connect again with new passphrase
	cmd = cliTD.StubRootCmd(
		"repo", "connect", "s3",
		"--config-file", configFP,
		"--bucket", cfg.Bucket,
		"--prefix", cfg.Prefix,
		"--passphrase", "newpass")
	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// connect with old passphrase - it will fail
	cmd = cliTD.StubRootCmd(
		"repo", "connect", "s3",
		"--config-file", configFP,
		"--bucket", cfg.Bucket,
		"--prefix", cfg.Prefix)
	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}
