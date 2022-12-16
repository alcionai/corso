package operations

import (
	"context"
	"testing"
	"time"

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
	"github.com/alcionai/corso/src/internal/stats"
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

func (suite *RestoreOpSuite) TestRestoreOperation_PersistResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		acct = account.Account{}
		now  = time.Now()
		dest = tester.DefaultTestRestoreDestination()
	)

	table := []struct {
		expectStatus opStatus
		expectErr    assert.ErrorAssertionFunc
		stats        restoreStats
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: restoreStats{
				started:       true,
				resourceCount: 1,
				bytesRead: &stats.ByteCounter{
					NumBytes: 42,
				},
				cs: []data.Collection{&exchange.Collection{}},
				gc: &support.ConnectorOperationStatus{
					ObjectCount: 1,
					Successful:  1,
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			stats: restoreStats{
				started:   false,
				bytesRead: &stats.ByteCounter{},
				gc:        &support.ConnectorOperationStatus{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: restoreStats{
				started:   true,
				bytesRead: &stats.ByteCounter{},
				cs:        []data.Collection{},
				gc:        &support.ConnectorOperationStatus{},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.expectStatus.String(), func(t *testing.T) {
			op, err := NewRestoreOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				"foo",
				selectors.Selector{},
				dest,
				evmock.NewBus())
			require.NoError(t, err)
			test.expectErr(t, op.persistResults(ctx, now, &test.stats))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, len(test.stats.cs), op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.readErr, op.Results.ReadErrors, "read errors")
			assert.Equal(t, test.stats.gc.Successful, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.bytesRead.NumBytes, op.Results.BytesRead, "resource owners")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, test.stats.writeErr, op.Results.WriteErrors, "write errors")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
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
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	t := suite.T()

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
	bsel.Include(
		bsel.MailFolders([]string{m365UserID}, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()),
		bsel.ContactFolders([]string{m365UserID}, []string{exchange.DefaultContactFolder}, selectors.PrefixMatch()),
		bsel.EventCalendars([]string{m365UserID}, []string{exchange.DefaultCalendar}, selectors.PrefixMatch()),
	)

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
	// Discount metadata files (3 paths, 2 deltas) as
	// they are not part of the data restored.
	suite.numItems = bo.Results.ItemsWritten - 5
}

func (suite *RestoreOpIntegrationSuite) TearDownSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

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
	dest := tester.DefaultTestRestoreDestination()

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
			ctx, flush := tester.NewContext()
			defer flush()

			_, err := NewRestoreOperation(
				ctx,
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				"backup-id",
				selectors.Selector{},
				dest,
				evmock.NewBus())
			test.errCheck(t, err)
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	rsel := selectors.NewExchangeRestore()
	rsel.Include(rsel.Users([]string{tester.M365UserID(t)}))

	dest := tester.DefaultTestRestoreDestination()
	mb := evmock.NewBus()

	ro, err := NewRestoreOperation(
		ctx,
		control.Options{},
		suite.kw,
		suite.sw,
		tester.NewM365Account(t),
		suite.backupID,
		rsel.Selector,
		dest,
		mb)
	require.NoError(t, err)

	ds, err := ro.Run(ctx)

	require.NoError(t, err, "restoreOp.Run()")
	require.NotEmpty(t, ro.Results, "restoreOp results")
	require.NotNil(t, ds, "restored details")
	assert.Equal(t, ro.Status, Completed, "restoreOp status")
	assert.Equal(t, ro.Results.ItemsWritten, len(ds.Entries), "count of items written matches restored entries in details")
	assert.Less(t, 0, ro.Results.ItemsRead, "restore items read")
	assert.Less(t, 0, ro.Results.ItemsWritten, "restored items written")
	assert.Less(t, int64(0), ro.Results.BytesRead, "bytes read")
	assert.Equal(t, 1, ro.Results.ResourceOwners, "resource Owners")
	assert.Zero(t, ro.Results.ReadErrors, "errors while reading restore data")
	assert.Zero(t, ro.Results.WriteErrors, "errors while writing restore data")
	assert.Equal(t, suite.numItems, ro.Results.ItemsWritten, "backup and restore wrote the same num of items")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run_ErrorNoResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	rsel := selectors.NewExchangeRestore()
	rsel.Include(rsel.Users(selectors.None()))

	dest := tester.DefaultTestRestoreDestination()
	mb := evmock.NewBus()

	ro, err := NewRestoreOperation(
		ctx,
		control.Options{},
		suite.kw,
		suite.sw,
		tester.NewM365Account(t),
		suite.backupID,
		rsel.Selector,
		dest,
		mb)
	require.NoError(t, err)

	ds, err := ro.Run(ctx)
	require.Error(t, err, "restoreOp.Run() should have errored")
	require.Nil(t, ds, "restoreOp.Run() should not produce details")
	assert.Zero(t, ro.Results.ResourceOwners, "resource owners")
	assert.Zero(t, ro.Results.BytesRead, "bytes read")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Zero(t, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}
