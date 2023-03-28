package operations

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
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
	tester.Suite
}

func TestRestoreOpSuite(t *testing.T) {
	suite.Run(t, &RestoreOpSuite{Suite: tester.NewUnitSuite(t)})
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
		fail         error
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: restoreStats{
				resourceCount: 1,
				bytesRead: &stats.ByteCounter{
					NumBytes: 42,
				},
				cs: []data.RestoreCollection{
					data.NotFoundRestoreCollection{
						Collection: &mockconnector.MockExchangeDataCollection{},
					},
				},
				gc: &support.ConnectorOperationStatus{
					Metrics: support.CollectionMetrics{
						Objects:   1,
						Successes: 1,
					},
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: restoreStats{
				bytesRead: &stats.ByteCounter{},
				gc:        &support.ConnectorOperationStatus{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: restoreStats{
				bytesRead: &stats.ByteCounter{},
				cs:        []data.RestoreCollection{},
				gc:        &support.ConnectorOperationStatus{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.expectStatus.String(), func() {
			t := suite.T()

			op, err := NewRestoreOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				"foo",
				selectors.Selector{DiscreteOwner: "test"},
				dest,
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			err = op.persistResults(ctx, now, &test.stats)
			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, len(test.stats.cs), op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.gc.Metrics.Successes, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.bytesRead.NumBytes, op.Results.BytesRead, "resource owners")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type bupResults struct {
	selectorResourceOwners []string
	resourceOwner          string
	backupID               model.StableID
	items                  int
}

type RestoreOpIntegrationSuite struct {
	tester.Suite

	exchange   bupResults
	sharepoint bupResults

	kopiaCloser func(ctx context.Context)
	kw          *kopia.Wrapper
	sw          *store.Wrapper
	ms          *kopia.ModelStore
}

func TestRestoreOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &RestoreOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs}),
	})
}

func (suite *RestoreOpIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		m365UserID = tester.M365UserID(t)
		acct       = tester.NewM365Account(t)
		st         = tester.NewPrefixedS3Storage(t)
		k          = kopia.NewConn(st)
	)

	err := k.Initialize(ctx)
	require.NoError(t, err, clues.ToCore(err))

	suite.kopiaCloser = func(ctx context.Context) {
		k.Close(ctx)
	}

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.kw = kw

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.ms = ms

	sw := store.NewKopiaStore(ms)
	suite.sw = sw

	suite.Run("exchange_setup", func() {
		var (
			t     = suite.T()
			users = []string{m365UserID}
			bsel  = selectors.NewExchangeBackup(users)
		)

		bsel.DiscreteOwner = m365UserID
		bsel.Include(
			bsel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()),
			bsel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch()),
			bsel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()),
		)

		bo, err := NewBackupOperation(
			ctx,
			control.Options{},
			kw,
			sw,
			acct,
			bsel.Selector,
			bsel.Selector.DiscreteOwner,
			evmock.NewBus())
		require.NoError(t, err, clues.ToCore(err))

		err = bo.Run(ctx)
		require.NoError(t, err, clues.ToCore(err))
		require.NotEmpty(t, bo.Results.BackupID)

		suite.exchange = bupResults{
			selectorResourceOwners: users,
			resourceOwner:          m365UserID,
			backupID:               bo.Results.BackupID,
			// Discount metadata collection files (1 delta and one prev path for each category).
			// These meta files are used to aid restore, but are not themselves
			// restored (ie: counted as writes).
			items: bo.Results.ItemsWritten - 6,
		}
	})

	suite.Run("sharepoint_setup", func() {
		var (
			t      = suite.T()
			siteID = tester.M365SiteID(t)
			sites  = []string{siteID}
			spsel  = selectors.NewSharePointBackup(sites)
		)

		spsel.DiscreteOwner = siteID
		// assume a folder name "test" exists in the drive.
		// this is brittle, and requires us to backfill anytime
		// the site under test changes, but also prevents explosive
		// growth from re-backup/restore of restored files.
		spsel.Include(spsel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))

		bo, err := NewBackupOperation(
			ctx,
			control.Options{},
			kw,
			sw,
			acct,
			spsel.Selector,
			spsel.Selector.DiscreteOwner,
			evmock.NewBus())
		require.NoError(t, err, clues.ToCore(err))

		// get the count of drives
		m365, err := acct.M365Config()
		require.NoError(t, err, clues.ToCore(err))

		adpt, err := graph.CreateAdapter(
			m365.AzureTenantID,
			m365.AzureClientID,
			m365.AzureClientSecret)
		require.NoError(t, err, clues.ToCore(err))

		service := graph.NewService(adpt)
		spPgr := api.NewSiteDrivePager(service, siteID, []string{"id", "name"})

		drives, err := api.GetAllDrives(ctx, spPgr, true, 3)
		require.NoError(t, err, clues.ToCore(err))

		err = bo.Run(ctx)
		require.NoError(t, err, clues.ToCore(err))
		require.NotEmpty(t, bo.Results.BackupID)

		suite.sharepoint = bupResults{
			selectorResourceOwners: sites,
			resourceOwner:          siteID,
			backupID:               bo.Results.BackupID,
			// Discount metadata files (1 delta, 1 prev path)
			// assume only one folder, and therefore 1 dirmeta per drive
			// assume only one file in each folder, and therefore 1 meta per drive.
			// These meta files are used to aid restore, but are not themselves
			// restored (ie: counted as writes).
			items: bo.Results.ItemsWritten - 2 - len(drives) - len(drives),
		}
	})
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
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			_, err := NewRestoreOperation(
				ctx,
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				"backup-id",
				selectors.Selector{DiscreteOwner: "test"},
				dest,
				evmock.NewBus())
			test.errCheck(suite.T(), err, clues.ToCore(err))
		})
	}
}

//nolint:lll
func (suite *RestoreOpIntegrationSuite) TestRestore_Run() {
	ctx, flush := tester.NewContext()
	defer flush()

	tables := []struct {
		name          string
		bID           model.StableID
		expectedItems int
		dest          control.RestoreDestination
		getSelector   func(t *testing.T) selectors.Selector
		cleanup       func(t *testing.T, dest string)
	}{
		{
			name:          "Exchange_Restore",
			bID:           suite.exchange.backupID,
			expectedItems: suite.exchange.items,
			dest:          tester.DefaultTestRestoreDestination(),
			getSelector: func(t *testing.T) selectors.Selector {
				rsel := selectors.NewExchangeRestore(suite.exchange.selectorResourceOwners)
				rsel.Include(rsel.AllData())

				return rsel.Selector
			},
		},
		{
			name:          "SharePoint_Restore",
			bID:           suite.sharepoint.backupID,
			expectedItems: suite.sharepoint.items,
			dest:          control.DefaultRestoreDestination(common.SimpleDateTimeOneDrive),
			getSelector: func(t *testing.T) selectors.Selector {
				rsel := selectors.NewSharePointRestore(suite.sharepoint.selectorResourceOwners)
				rsel.Include(rsel.AllData())

				return rsel.Selector
			},
		},
	}

	for _, test := range tables {
		suite.Run(test.name, func() {
			t := suite.T()
			mb := evmock.NewBus()

			ro, err := NewRestoreOperation(
				ctx,
				control.Options{FailFast: true},
				suite.kw,
				suite.sw,
				tester.NewM365Account(t),
				test.bID,
				test.getSelector(t),
				test.dest,
				mb)
			require.NoError(t, err, clues.ToCore(err))

			ds, err := ro.Run(ctx)

			require.NoError(t, err, "restoreOp.Run() %+v", clues.ToCore(err))
			require.NotEmpty(t, ro.Results, "restoreOp results")
			require.NotNil(t, ds, "restored details")
			assert.Equal(t, ro.Status, Completed, "restoreOp status")
			assert.Equal(t, ro.Results.ItemsWritten, len(ds.Entries), "count of items written matches restored entries in details")
			assert.Less(t, 0, ro.Results.ItemsRead, "restore items read")
			assert.Less(t, int64(0), ro.Results.BytesRead, "bytes read")
			assert.Equal(t, 1, ro.Results.ResourceOwners, "resource Owners")
			assert.NoError(t, ro.Errors.Failure(), "non-recoverable error", clues.ToCore(ro.Errors.Failure()))
			assert.Empty(t, ro.Errors.Recovered(), "recoverable errors")
			assert.Equal(t, test.expectedItems, ro.Results.ItemsWritten, "backup and restore wrote the same num of items")
			assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
			assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run_ErrorNoResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	rsel := selectors.NewExchangeRestore(selectors.None())
	rsel.Include(rsel.AllData())

	dest := tester.DefaultTestRestoreDestination()
	mb := evmock.NewBus()

	ro, err := NewRestoreOperation(
		ctx,
		control.Options{},
		suite.kw,
		suite.sw,
		tester.NewM365Account(t),
		suite.exchange.backupID,
		rsel.Selector,
		dest,
		mb)
	require.NoError(t, err, clues.ToCore(err))

	ds, err := ro.Run(ctx)
	require.Error(t, err, "restoreOp.Run() should have errored")
	require.Nil(t, ds, "restoreOp.Run() should not produce details")
	assert.Zero(t, ro.Results.ResourceOwners, "resource owners")
	assert.Zero(t, ro.Results.BytesRead, "bytes read")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Zero(t, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}
