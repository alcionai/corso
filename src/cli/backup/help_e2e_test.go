package backup_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
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
		acct     = tester.NewM365Account(t)
		st       = tester.NewPrefixedS3Storage(t)
		recorder = strings.Builder{}
	)

	cfg, err := st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tester.TestCfgAccountProvider: "M365",
		tester.TestCfgStorageProvider: "S3",
		tester.TestCfgPrefix:          cfg.Prefix,
	}

	vpr, cfgFP := tester.MakeTempTestConfigClone(t, force)
	ctx = config.SetViper(ctx, vpr)

	repo, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err, clues.ToCore(err))

	return acct, st, repo, vpr, recorder, cfgFP
}
