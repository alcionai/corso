package operations

import (
	"context"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/internal/kopia"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/store"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type RestoreOpSuite struct {
	suite.Suite
}

func TestRestoreOpSuite(t *testing.T) {
	suite.Run(t, new(RestoreOpSuite))
}

// TODO: after modelStore integration is added, mock the store and/or
// move this to an integration test.
func (suite *RestoreOpSuite) TestRestoreOperation_PersistResults() {
	t := suite.T()
	ctx := context.Background()

	var (
		kw    = &kopia.Wrapper{}
		sw    = &store.Wrapper{}
		acct  = account.Account{}
		now   = time.Now()
		stats = restoreStats{
			readErr:  multierror.Append(nil, assert.AnError),
			writeErr: assert.AnError,
			cs:       []data.Collection{&connector.ExchangeDataCollection{}},
			gc: &support.ConnectorOperationStatus{
				ObjectCount: 1,
			},
		}
	)

	op, err := NewRestoreOperation(ctx, Options{}, kw, sw, acct, "foo", selectors.Selector{})
	require.NoError(t, err)

	op.persistResults(now, &stats)

	assert.Equal(t, op.Status, Failed)
	assert.Equal(t, op.Results.ItemsRead, len(stats.cs))
	assert.Equal(t, op.Results.ReadErrors, stats.readErr)
	assert.Equal(t, op.Results.ItemsWritten, stats.gc.ObjectCount)
	assert.Equal(t, op.Results.WriteErrors, stats.writeErr)
	assert.Equal(t, op.Results.StartedAt, now)
	assert.Less(t, now, op.Results.CompletedAt)
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type RestoreOpIntegrationSuite struct {
	suite.Suite
}

func TestRestoreOpIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoOperationTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(RestoreOpIntegrationSuite))
}

func (suite *RestoreOpIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *RestoreOpIntegrationSuite) TestNewRestoreOperation() {
	kw := &kopia.Wrapper{}
	sw := &store.Wrapper{}
	acct, err := ctesting.NewM365Account()
	require.NoError(suite.T(), err)

	table := []struct {
		name     string
		opts     Options
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", Options{}, kw, sw, acct, nil, assert.NoError},
		{"missing kopia", Options{}, nil, sw, acct, nil, assert.Error},
		{"missing modelstore", Options{}, kw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := NewRestoreOperation(
				context.Background(),
				Options{},
				test.kw,
				test.sw,
				test.acct,
				"backup-id",
				selectors.Selector{})
			test.errCheck(t, err)
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run() {
	t := suite.T()
	ctx := context.Background()

	m365User := "lidiah@8qzvrj.onmicrosoft.com"
	acct, err := ctesting.NewM365Account()
	require.NoError(t, err)

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))
	defer k.Close(ctx)

	w, err := kopia.NewWrapper(k)
	require.NoError(t, err)
	defer w.Close(ctx)

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)
	defer ms.Close(ctx)

	sw := store.NewKopiaStore(ms)

	bsel := selectors.NewExchangeBackup()
	bsel.Include(bsel.Users([]string{m365User}))

	bo, err := NewBackupOperation(
		ctx,
		Options{},
		w,
		sw,
		acct,
		bsel.Selector)
	require.NoError(t, err)
	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results.BackupID)

	rsel := selectors.NewExchangeRestore()
	rsel.Include(rsel.Users([]string{m365User}))

	ro, err := NewRestoreOperation(
		ctx,
		Options{},
		w,
		sw,
		acct,
		bo.Results.BackupID,
		rsel.Selector)
	require.NoError(t, err)

	require.NoError(t, ro.Run(ctx), "restoreOp.Run()")
	require.NotEmpty(t, ro.Results, "restoreOp results")
	assert.Equal(t, ro.Status, Successful, "restoreOp status")
	assert.Greater(t, ro.Results.ItemsRead, 0, "restore items read")
	assert.Greater(t, ro.Results.ItemsWritten, 0, "restored items written")
	assert.Zero(t, ro.Results.ReadErrors, "errors while reading restore data")
	assert.Zero(t, ro.Results.WriteErrors, "errors while writing restore data")
	assert.Equal(t, bo.Results.ItemsWritten, ro.Results.ItemsWritten, "backup and restore wrote the same num of items")
}
