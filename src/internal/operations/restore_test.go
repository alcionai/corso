package operations

import (
	"context"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
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
			started:  true,
			readErr:  multierror.Append(nil, assert.AnError),
			writeErr: assert.AnError,
			cs:       []data.Collection{&exchange.Collection{}},
			gc: &support.ConnectorOperationStatus{
				ObjectCount: 1,
			},
		}
	)

	op, err := NewRestoreOperation(
		ctx,
		control.Options{},
		kw,
		sw,
		acct,
		"foo",
		selectors.Selector{},
		evmock.NewBus())
	require.NoError(t, err)

	require.NoError(t, op.persistResults(ctx, now, &stats))

	assert.Equal(t, op.Status.String(), Completed.String(), "status")
	assert.Equal(t, op.Results.ItemsRead, len(stats.cs), "items read")
	assert.Equal(t, op.Results.ReadErrors, stats.readErr, "read errors")
	assert.Equal(t, op.Results.ItemsWritten, stats.gc.Successful, "items written")
	assert.Equal(t, op.Results.WriteErrors, stats.writeErr, "write errors")
	assert.Equal(t, op.Results.StartedAt, now, "started at")
	assert.Less(t, now, op.Results.CompletedAt, "completed at")
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type RestoreOpIntegrationSuite struct {
	suite.Suite

	backupID    model.StableID
	numItems    int
	kopiaCloser func(ctx context.Context)
	kw          *kopia.Wrapper
	sw          *store.Wrapper
	ms          *kopia.ModelStore
}

func TestRestoreOpIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoOperationTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RestoreOpIntegrationSuite))
}

func (suite *RestoreOpIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	t := suite.T()
	ctx := context.Background()

	m365UserID := tester.M365UserID(t)
	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	suite.kopiaCloser = func(ctx context.Context) {
		k.Close(ctx)
	}

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err)

	suite.kw = kw

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)

	suite.ms = ms

	sw := store.NewKopiaStore(ms)
	suite.sw = sw

	bsel := selectors.NewExchangeBackup()
	bsel.Include(bsel.MailFolders([]string{m365UserID}, []string{"Inbox"}))

	bo, err := NewBackupOperation(
		ctx,
		control.Options{},
		kw,
		sw,
		acct,
		bsel.Selector,
		evmock.NewBus())
	require.NoError(t, err)
	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results.BackupID)

	suite.backupID = bo.Results.BackupID
	suite.numItems = bo.Results.ItemsWritten
}

func (suite *RestoreOpIntegrationSuite) TearDownSuite() {
	ctx := context.Background()
	if suite.ms != nil {
		suite.ms.Close(ctx)
	}

	if suite.kw != nil {
		suite.kw.Close(ctx)
	}

	if suite.kopiaCloser != nil {
		suite.kopiaCloser(ctx)
	}
}

func (suite *RestoreOpIntegrationSuite) TestNewRestoreOperation() {
	kw := &kopia.Wrapper{}
	sw := &store.Wrapper{}
	acct := tester.NewM365Account(suite.T())

	table := []struct {
		name     string
		opts     control.Options
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", control.Options{}, kw, sw, acct, nil, assert.NoError},
		{"missing kopia", control.Options{}, nil, sw, acct, nil, assert.Error},
		{"missing modelstore", control.Options{}, kw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := NewRestoreOperation(
				context.Background(),
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				"backup-id",
				selectors.Selector{},
				evmock.NewBus())
			test.errCheck(t, err)
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run() {
	t := suite.T()
	ctx := context.Background()

	rsel := selectors.NewExchangeRestore()
	rsel.Include(rsel.Users([]string{tester.M365UserID(t)}))

	mb := evmock.NewBus()

	ro, err := NewRestoreOperation(
		ctx,
		control.Options{},
		suite.kw,
		suite.sw,
		tester.NewM365Account(t),
		suite.backupID,
		rsel.Selector,
		mb)
	require.NoError(t, err)

	require.NoError(t, ro.Run(ctx), "restoreOp.Run()")
	require.NotEmpty(t, ro.Results, "restoreOp results")
	assert.Equal(t, ro.Status, Completed, "restoreOp status")
	assert.Greater(t, ro.Results.ItemsRead, 0, "restore items read")
	assert.Greater(t, ro.Results.ItemsWritten, 0, "restored items written")
	assert.Zero(t, ro.Results.ReadErrors, "errors while reading restore data")
	assert.Zero(t, ro.Results.WriteErrors, "errors while writing restore data")
	assert.Equal(t, suite.numItems, ro.Results.ItemsWritten, "backup and restore wrote the same num of items")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run_ErrorNoResults() {
	t := suite.T()
	ctx := context.Background()

	rsel := selectors.NewExchangeRestore()
	rsel.Include(rsel.Users(selectors.None()))

	mb := evmock.NewBus()

	ro, err := NewRestoreOperation(
		ctx,
		control.Options{},
		suite.kw,
		suite.sw,
		tester.NewM365Account(t),
		suite.backupID,
		rsel.Selector,
		mb)
	require.NoError(t, err)
	require.Error(t, ro.Run(ctx), "restoreOp.Run() should have 0 results")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Equal(t, 0, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}
