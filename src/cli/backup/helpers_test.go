package backup_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/print"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/config"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/storage/testdata"
)

type dependencies struct {
	st             storage.Storage
	repo           repository.Repositoryer
	vpr            *viper.Viper
	recorder       strings.Builder
	configFilePath string
}

func prepM365Test(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	pst path.ServiceType,
) dependencies {
	var (
		acct     = tconfig.NewM365Account(t)
		st       = testdata.NewPrefixedS3Storage(t)
		recorder = strings.Builder{}
	)

	cfg, err := st.ToS3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tconfig.TestCfgAccountProvider: account.ProviderM365.String(),
		tconfig.TestCfgStorageProvider: storage.ProviderS3.String(),
		tconfig.TestCfgPrefix:          cfg.Prefix,
	}

	vpr, cfgFP := tconfig.MakeTempTestConfigClone(t, force)
	ctx = config.SetViper(ctx, vpr)

	repo, err := repository.New(
		ctx,
		acct,
		st,
		control.DefaultOptions(),
		repository.NewRepoID)
	require.NoError(t, err, clues.ToCore(err))

	err = repo.Initialize(ctx, repository.InitConfig{
		Service: pst,
	})
	require.NoError(t, err, clues.ToCore(err))

	return dependencies{
		st:             st,
		repo:           repo,
		vpr:            vpr,
		recorder:       recorder,
		configFilePath: cfgFP,
	}
}

// ---------------------------------------------------------------------------
// funcs
// ---------------------------------------------------------------------------

func buildExchangeBackupCmd(
	ctx context.Context,
	configFile, user, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := cliTD.StubRootCmd(
		"backup", "create", "exchange",
		"--"+flags.ConfigFileFN, configFile,
		"--"+flags.UserFN, user,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
