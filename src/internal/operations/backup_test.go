package operations

import (
	"context"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/kopia"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/store"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type BackupOpSuite struct {
	suite.Suite
}

func TestBackupOpSuite(t *testing.T) {
	suite.Run(t, new(BackupOpSuite))
}

// TODO: after modelStore integration is added, mock the store and/or
// move this to an integration test.
func (suite *BackupOpSuite) TestBackupOperation_PersistResults() {
	t := suite.T()
	ctx := context.Background()

	var (
		kw    = &kopia.Wrapper{}
		sw    = &store.Wrapper{}
		acct  = account.Account{}
		now   = time.Now()
		stats = backupStats{
			readErr:  multierror.Append(nil, assert.AnError),
			writeErr: assert.AnError,
			k: &kopia.BackupStats{
				TotalFileCount: 1,
			},
			gc: &support.ConnectorOperationStatus{
				Successful: 1,
			},
		}
	)

	op, err := NewBackupOperation(ctx, Options{}, kw, sw, acct, selectors.Selector{})
	require.NoError(t, err)

	op.persistResults(now, &stats)

	assert.Equal(t, op.Status, Failed)
	assert.Equal(t, op.Results.ItemsRead, stats.gc.Successful)
	assert.Equal(t, op.Results.ReadErrors, stats.readErr)
	assert.Equal(t, op.Results.ItemsWritten, stats.k.TotalFileCount)
	assert.Equal(t, op.Results.WriteErrors, stats.writeErr)
	assert.Equal(t, op.Results.StartedAt, now)
	assert.Less(t, now, op.Results.CompletedAt)
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	suite.Suite
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoOperationTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(BackupOpIntegrationSuite))
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(
		append(
			ctesting.AWSStorageCredEnvs,
			ctesting.M365AcctCredEnvs...,
		)...,
	)
	require.NoError(suite.T(), err)
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
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
			_, err := NewBackupOperation(
				context.Background(),
				Options{},
				test.kw,
				test.sw,
				test.acct,
				selectors.Selector{})
			test.errCheck(t, err)
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run() {
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

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	defer k.Close(ctx)

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err)
	defer kw.Close(ctx)

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)
	defer ms.Close(ctx)

	sw := store.NewKopiaStore(ms)

	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{m365User}))

	bo, err := NewBackupOperation(
		ctx,
		Options{},
		kw,
		sw,
		acct,
		sel.Selector)
	require.NoError(t, err)

	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results)
	require.NotEmpty(t, bo.Results.BackupID)
	assert.Equal(t, bo.Status, Successful)
	assert.Greater(t, bo.Results.ItemsRead, 0)
	assert.Greater(t, bo.Results.ItemsWritten, 0)
	assert.Zero(t, bo.Results.ReadErrors)
	assert.Zero(t, bo.Results.WriteErrors)
}
