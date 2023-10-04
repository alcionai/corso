package repo_test

import (
	"os"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/config"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type S3E2ESuite struct {
	tester.Suite
}

func TestS3E2ESuite(t *testing.T) {
	suite.Run(t, &S3E2ESuite{Suite: tester.NewE2ESuite(
		t,
		[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs})})
}

func (suite *S3E2ESuite) TestInitS3Cmd() {
	table := []struct {
		name          string
		bucketPrefix  string
		hasConfigFile bool
	}{
		{
			name:          "NoPrefix",
			bucketPrefix:  "",
			hasConfigFile: true,
		},
		{
			name:          "S3Prefix",
			bucketPrefix:  "s3://",
			hasConfigFile: true,
		},
		{
			name:          "NoConfigFile",
			bucketPrefix:  "",
			hasConfigFile: false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			st := storeTD.NewPrefixedS3Storage(t)

			cfg, err := st.ToS3Config()
			require.NoError(t, err, clues.ToCore(err))

			vpr, configFP := tconfig.MakeTempTestConfigClone(t, nil)
			if !test.hasConfigFile {
				// Ideally we could use `/dev/null`, but you need a
				// toml file plus this works cross platform
				os.Remove(configFP)
			}

			ctx = config.SetViper(ctx, vpr)

			cmd := cliTD.StubRootCmd(
				"repo", "init", "s3",
				"--config-file", configFP,
				"--bucket", test.bucketPrefix+cfg.Bucket,
				"--prefix", cfg.Prefix)
			cli.BuildCommandTree(cmd)

			// run the command
			err = cmd.ExecuteContext(ctx)
			require.NoError(t, err, clues.ToCore(err))

			// a second initialization should result in an error
			err = cmd.ExecuteContext(ctx)
			assert.ErrorIs(t, err, repository.ErrorRepoAlreadyExists, clues.ToCore(err))
		})
	}
}

func (suite *S3E2ESuite) TestInitMultipleTimes() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	st := storeTD.NewPrefixedS3Storage(t)

	cfg, err := st.ToS3Config()
	require.NoError(t, err, clues.ToCore(err))

	vpr, configFP := tconfig.MakeTempTestConfigClone(t, nil)

	ctx = config.SetViper(ctx, vpr)

	for i := 0; i < 2; i++ {
		cmd := cliTD.StubRootCmd(
			"repo", "init", "s3",
			"--config-file", configFP,
			"--bucket", cfg.Bucket,
			"--prefix", cfg.Prefix,
			"--succeed-if-exists")
		cli.BuildCommandTree(cmd)

		// run the command
		err = cmd.ExecuteContext(ctx)
		require.NoError(t, err, clues.ToCore(err))
	}
}

func (suite *S3E2ESuite) TestInitS3Cmd_missingBucket() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()

	st := storeTD.NewPrefixedS3Storage(t)

	cfg, err := st.ToS3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tconfig.TestCfgBucket: "",
	}

	vpr, configFP := tconfig.MakeTempTestConfigClone(t, force)

	ctx = config.SetViper(ctx, vpr)

	cmd := cliTD.StubRootCmd(
		"repo", "init", "s3",
		"--config-file", configFP,
		"--prefix", cfg.Prefix)
	cli.BuildCommandTree(cmd)

	// run the command
	err = cmd.ExecuteContext(ctx)
	require.Error(t, err, clues.ToCore(err))
}

func (suite *S3E2ESuite) TestConnectS3Cmd() {
	table := []struct {
		name          string
		bucketPrefix  string
		hasConfigFile bool
	}{
		{
			name:          "NoPrefix",
			bucketPrefix:  "",
			hasConfigFile: true,
		},
		{
			name:          "S3Prefix",
			bucketPrefix:  "s3://",
			hasConfigFile: true,
		},
		{
			name:          "NoConfigFile",
			bucketPrefix:  "",
			hasConfigFile: false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			st := storeTD.NewPrefixedS3Storage(t)

			cfg, err := st.ToS3Config()
			require.NoError(t, err, clues.ToCore(err))

			force := map[string]string{
				tconfig.TestCfgAccountProvider: account.ProviderM365.String(),
				tconfig.TestCfgStorageProvider: storage.ProviderS3.String(),
				tconfig.TestCfgPrefix:          cfg.Prefix,
			}
			vpr, configFP := tconfig.MakeTempTestConfigClone(t, force)
			if !test.hasConfigFile {
				// Ideally we could use `/dev/null`, but you need a
				// toml file plus this works cross platform
				os.Remove(configFP)
			}

			ctx = config.SetViper(ctx, vpr)

			// init the repo first
			r, err := repository.New(
				ctx,
				tconfig.NewM365Account(t),
				st,
				control.DefaultOptions(),
				repository.NewRepoID)
			require.NoError(t, err, clues.ToCore(err))

			err = r.Initialize(ctx, repository.InitConfig{})
			require.NoError(t, err, clues.ToCore(err))

			// then test it
			cmd := cliTD.StubRootCmd(
				"repo", "connect", "s3",
				"--config-file", configFP,
				"--bucket", test.bucketPrefix+cfg.Bucket,
				"--prefix", cfg.Prefix)
			cli.BuildCommandTree(cmd)

			// run the command
			err = cmd.ExecuteContext(ctx)
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

func (suite *S3E2ESuite) TestConnectS3Cmd_badInputs() {
	table := []struct {
		name      string
		bucket    string
		prefix    string
		expectErr func(t *testing.T, err error)
	}{
		{
			name:   "bucket",
			bucket: "wrong",
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, storage.ErrVerifyingConfigStorage, clues.ToCore(err))
			},
		},
		{
			name:   "prefix",
			prefix: "wrong",
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, storage.ErrVerifyingConfigStorage, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			st := storeTD.NewPrefixedS3Storage(t)
			cfg, err := st.ToS3Config()
			require.NoError(t, err, clues.ToCore(err))

			bucket := str.First(test.bucket, cfg.Bucket)
			prefix := str.First(test.prefix, cfg.Prefix)

			over := map[string]string{}
			acct := tconfig.NewM365Account(t)

			maps.Copy(over, acct.Config)
			over[account.AccountProviderTypeKey] = account.ProviderM365.String()
			over[storage.StorageProviderTypeKey] = storage.ProviderS3.String()

			vpr, configFP := tconfig.MakeTempTestConfigClone(t, over)
			ctx = config.SetViper(ctx, vpr)

			cmd := cliTD.StubRootCmd(
				"repo", "connect", "s3",
				"--config-file", configFP,
				"--bucket", bucket,
				"--prefix", prefix)
			cli.BuildCommandTree(cmd)

			// run the command
			err = cmd.ExecuteContext(ctx)
			require.Error(t, err, clues.ToCore(err))
			test.expectErr(t, err)
		})
	}
}
