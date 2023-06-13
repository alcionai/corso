package operations

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
	var (
		kw         = &kopia.Wrapper{}
		sw         = &store.Wrapper{}
		gc         = &mock.GraphConnector{}
		now        = time.Now()
		restoreCfg = tester.DefaultTestRestoreConfig("")
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
					data.NoFetchRestoreCollection{
						Collection: &exchMock.DataCollection{},
					},
				},
				gc: &data.CollectionStats{
					Objects:   1,
					Successes: 1,
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: restoreStats{
				bytesRead: &stats.ByteCounter{},
				gc:        &data.CollectionStats{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: restoreStats{
				bytesRead: &stats.ByteCounter{},
				cs:        []data.RestoreCollection{},
				gc:        &data.CollectionStats{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.expectStatus.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			op, err := NewRestoreOperation(
				ctx,
				control.Defaults(),
				kw,
				sw,
				gc,
				account.Account{},
				"foo",
				selectors.Selector{DiscreteOwner: "test"},
				restoreCfg,
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			err = op.persistResults(ctx, now, &test.stats)
			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, len(test.stats.cs), op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.gc.Successes, op.Results.ItemsWritten, "items written")
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
	backupID               model.StableID
	items                  int
	gc                     *connector.GraphConnector
}

type RestoreOpIntegrationSuite struct {
	tester.Suite

	kopiaCloser func(ctx context.Context)
	acct        account.Account
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
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		st = tester.NewPrefixedS3Storage(t)
		k  = kopia.NewConn(st)
	)

	suite.acct = tester.NewM365Account(t)

	err := k.Initialize(ctx, repository.Options{})
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
}

func (suite *RestoreOpIntegrationSuite) TearDownSuite() {
	ctx, flush := tester.NewContext(suite.T())
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
	var (
		kw         = &kopia.Wrapper{}
		sw         = &store.Wrapper{}
		gc         = &mock.GraphConnector{}
		restoreCfg = tester.DefaultTestRestoreConfig("")
		opts       = control.Defaults()
	)

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		rc       inject.RestoreConsumer
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kw, sw, gc, nil, assert.NoError},
		{"missing kopia", nil, sw, gc, nil, assert.Error},
		{"missing modelstore", kw, nil, gc, nil, assert.Error},
		{"missing restore consumer", kw, sw, nil, nil, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := NewRestoreOperation(
				ctx,
				opts,
				test.kw,
				test.sw,
				test.rc,
				tester.NewM365Account(t),
				"backup-id",
				selectors.Selector{DiscreteOwner: "test"},
				restoreCfg,
				evmock.NewBus())
			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}

func setupExchangeBackup(
	t *testing.T,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	owner string,
) bupResults {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		users = []string{owner}
		esel  = selectors.NewExchangeBackup(users)
	)

	esel.DiscreteOwner = owner
	esel.Include(
		esel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()),
		esel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch()),
		esel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()))

	gc, sel := GCWithSelector(t, ctx, acct, connector.Users, esel.Selector, nil, nil)

	bo, err := NewBackupOperation(
		ctx,
		control.Defaults(),
		kw,
		sw,
		gc,
		acct,
		sel,
		inMock.NewProvider(owner, owner),
		evmock.NewBus())
	require.NoError(t, err, clues.ToCore(err))

	err = bo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bo.Results.BackupID)

	return bupResults{
		selectorResourceOwners: users,
		backupID:               bo.Results.BackupID,
		// Discount metadata collection files (1 delta and one prev path for each category).
		// These meta files are used to aid restore, but are not themselves
		// restored (ie: counted as writes).
		items: bo.Results.ItemsWritten - 6,
		gc:    gc,
	}
}

func setupSharePointBackup(
	t *testing.T,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	owner string,
) bupResults {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		sites = []string{owner}
		ssel  = selectors.NewSharePointBackup(sites)
	)

	// assume a folder name "test" exists in the drive.
	// this is brittle, and requires us to backfill anytime
	// the site under test changes, but also prevents explosive
	// growth from re-backup/restore of restored files.
	ssel.Include(ssel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))
	ssel.DiscreteOwner = owner

	gc, sel := GCWithSelector(t, ctx, acct, connector.Sites, ssel.Selector, nil, nil)

	bo, err := NewBackupOperation(
		ctx,
		control.Defaults(),
		kw,
		sw,
		gc,
		acct,
		sel,
		inMock.NewProvider(owner, owner),
		evmock.NewBus())
	require.NoError(t, err, clues.ToCore(err))

	spPgr := gc.AC.Drives().NewSiteDrivePager(owner, []string{"id", "name"})

	drives, err := api.GetAllDrives(ctx, spPgr, true, 3)
	require.NoError(t, err, clues.ToCore(err))

	err = bo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bo.Results.BackupID)

	return bupResults{
		selectorResourceOwners: sites,
		backupID:               bo.Results.BackupID,
		// Discount metadata files (1 delta, 1 prev path)
		// assume only one folder, and therefore 1 dirmeta per drive
		// assume only one file in each folder, and therefore 1 meta per drive.
		// These meta files are used to aid restore, but are not themselves
		// restored (ie: counted as writes).
		items: bo.Results.ItemsWritten - 2 - len(drives) - len(drives),
		gc:    gc,
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run() {
	tables := []struct {
		name        string
		owner       string
		restoreCfg  control.RestoreConfig
		getSelector func(t *testing.T, owners []string) selectors.Selector
		setup       func(t *testing.T, kw *kopia.Wrapper, sw *store.Wrapper, acct account.Account, owner string) bupResults
	}{
		{
			name:       "Exchange_Restore",
			owner:      tester.M365UserID(suite.T()),
			restoreCfg: tester.DefaultTestRestoreConfig(""),
			getSelector: func(t *testing.T, owners []string) selectors.Selector {
				rsel := selectors.NewExchangeRestore(owners)
				rsel.Include(rsel.AllData())

				return rsel.Selector
			},
			setup: setupExchangeBackup,
		},
		{
			name:       "SharePoint_Restore",
			owner:      tester.M365SiteID(suite.T()),
			restoreCfg: control.DefaultRestoreConfig(dttm.SafeForTesting),
			getSelector: func(t *testing.T, owners []string) selectors.Selector {
				rsel := selectors.NewSharePointRestore(owners)
				rsel.Include(rsel.AllData())

				return rsel.Selector
			},
			setup: setupSharePointBackup,
		},
	}

	for _, test := range tables {
		suite.Run(test.name, func() {
			var (
				t   = suite.T()
				mb  = evmock.NewBus()
				bup = test.setup(t, suite.kw, suite.sw, suite.acct, test.owner)
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			require.NotZero(t, bup.items)
			require.NotEmpty(t, bup.backupID)

			ro, err := NewRestoreOperation(
				ctx,
				control.Options{FailureHandling: control.FailFast},
				suite.kw,
				suite.sw,
				bup.gc,
				tester.NewM365Account(t),
				bup.backupID,
				test.getSelector(t, bup.selectorResourceOwners),
				test.restoreCfg,
				mb)
			require.NoError(t, err, clues.ToCore(err))

			ds, err := ro.Run(ctx)

			require.NoError(t, err, "restoreOp.Run() %+v", clues.ToCore(err))
			require.NotEmpty(t, ro.Results, "restoreOp results")
			require.NotNil(t, ds, "restored details")
			assert.Equal(t, ro.Status, Completed, "restoreOp status")
			assert.Equal(t, ro.Results.ItemsWritten, len(ds.Items()), "item write count matches len details")
			assert.Less(t, 0, ro.Results.ItemsRead, "restore items read")
			assert.Less(t, int64(0), ro.Results.BytesRead, "bytes read")
			assert.Equal(t, 1, ro.Results.ResourceOwners, "resource Owners")
			assert.NoError(t, ro.Errors.Failure(), "non-recoverable error", clues.ToCore(ro.Errors.Failure()))
			assert.Empty(t, ro.Errors.Recovered(), "recoverable errors")
			assert.Equal(t, bup.items, ro.Results.ItemsWritten, "backup and restore wrote the same num of items")
			assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
			assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run_errorNoBackup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		restoreCfg = tester.DefaultTestRestoreConfig("")
		mb         = evmock.NewBus()
	)

	rsel := selectors.NewExchangeRestore(selectors.None())
	rsel.Include(rsel.AllData())

	gc, err := connector.NewGraphConnector(
		ctx,
		suite.acct,
		connector.Users)
	require.NoError(t, err, clues.ToCore(err))

	ro, err := NewRestoreOperation(
		ctx,
		control.Defaults(),
		suite.kw,
		suite.sw,
		gc,
		tester.NewM365Account(t),
		"backupID",
		rsel.Selector,
		restoreCfg,
		mb)
	require.NoError(t, err, clues.ToCore(err))

	ds, err := ro.Run(ctx)
	require.Error(t, err, "restoreOp.Run() should have errored")
	require.Nil(t, ds, "restoreOp.Run() should not produce details")
	assert.Zero(t, ro.Results.ResourceOwners, "resource owners")
	assert.Zero(t, ro.Results.BytesRead, "bytes read")
	// no restore start, because we'd need to find the backup first.
	assert.Equal(t, 0, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")
}
