package backup_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/storage/testdata"
)

func prepM365Test(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) (
	account.Account,
	storage.Storage,
	repository.Repository,
	*viper.Viper,
	strings.Builder,
	string,
) {
	var (
		acct     = tconfig.NewM365Account(t)
		st       = testdata.NewPrefixedS3Storage(t)
		recorder = strings.Builder{}
	)

	cfg, err := st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tconfig.TestCfgAccountProvider: "M365",
		tconfig.TestCfgStorageProvider: "S3",
		tconfig.TestCfgPrefix:          cfg.Prefix,
	}

	vpr, cfgFP := tconfig.MakeTempTestConfigClone(t, force)
	ctx = config.SetViper(ctx, vpr)

	repo, err := repository.Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	return acct, st, repo, vpr, recorder, cfgFP
}
