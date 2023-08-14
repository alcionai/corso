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
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/mock"
	"github.com/alcionai/corso/src/internal/m365/resource"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/selectors"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type RestoreOpUnitSuite struct {
	tester.Suite
}

func TestRestoreOpUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreOpUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreOpUnitSuite) TestRestoreOperation_PersistResults() {
	var (
		kw         = &kopia.Wrapper{}
		sw         = store.NewWrapper(&kopia.ModelStore{})
		ctrl       = &mock.Controller{}
		now        = time.Now()
		restoreCfg = testdata.DefaultRestoreConfig("")
	)

	table := []struct {
		expectStatus OpStatus
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
				ctrl: &data.CollectionStats{
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
				ctrl:      &data.CollectionStats{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: restoreStats{
				bytesRead: &stats.ByteCounter{},
				cs:        []data.RestoreCollection{},
				ctrl:      &data.CollectionStats{},
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
				control.DefaultOptions(),
				kw,
				sw,
				ctrl,
				account.Account{},
				"foo",
				selectors.Selector{DiscreteOwner: "test"},
				restoreCfg,
				evmock.NewBus(),
				count.New())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			err = op.persistResults(ctx, now, &test.stats)
			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, len(test.stats.cs), op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.ctrl.Successes, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.bytesRead.NumBytes, op.Results.BytesRead, "resource owners")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

func (suite *RestoreOpUnitSuite) TestChooseRestoreResource() {
	var (
		id        = "id"
		name      = "name"
		cfgWithPR = control.DefaultRestoreConfig(dttm.HumanReadable)
	)

	cfgWithPR.ProtectedResource = "cfgid"

	table := []struct {
		name           string
		cfg            control.RestoreConfig
		ctrl           *mock.Controller
		orig           idname.Provider
		expectErr      assert.ErrorAssertionFunc
		expectProvider assert.ValueAssertionFunc
		expectID       string
		expectName     string
	}{
		{
			name: "use original",
			cfg:  control.DefaultRestoreConfig(dttm.HumanReadable),
			ctrl: &mock.Controller{
				ProtectedResourceID:   id,
				ProtectedResourceName: name,
			},
			orig:       idname.NewProvider("oid", "oname"),
			expectErr:  assert.NoError,
			expectID:   "oid",
			expectName: "oname",
		},
		{
			name: "look up resource with iface",
			cfg:  cfgWithPR,
			ctrl: &mock.Controller{
				ProtectedResourceID:   id,
				ProtectedResourceName: name,
			},
			orig:       idname.NewProvider("oid", "oname"),
			expectErr:  assert.NoError,
			expectID:   id,
			expectName: name,
		},
		{
			name: "error looking up protected resource",
			cfg:  cfgWithPR,
			ctrl: &mock.Controller{
				ProtectedResourceErr: assert.AnError,
			},
			orig:      idname.NewProvider("oid", "oname"),
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := chooseRestoreResource(ctx, test.ctrl, test.cfg, test.orig)
			test.expectErr(t, err, clues.ToCore(err))
			require.NotNil(t, result)
			assert.Equal(t, test.expectID, result.ID())
			assert.Equal(t, test.expectName, result.Name())
		})
	}
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type RestoreOpIntegrationSuite struct {
	tester.Suite

	kopiaCloser func(ctx context.Context)
	acct        account.Account
	kw          *kopia.Wrapper
	sw          store.BackupStorer
	ms          *kopia.ModelStore
}

func TestRestoreOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &RestoreOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *RestoreOpIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	var (
		st = storeTD.NewPrefixedS3Storage(t)
		k  = kopia.NewConn(st)
	)

	suite.acct = tconfig.NewM365Account(t)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{})
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

	sw := store.NewWrapper(ms)
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
		sw         = store.NewWrapper(&kopia.ModelStore{})
		ctrl       = &mock.Controller{}
		restoreCfg = testdata.DefaultRestoreConfig("")
		opts       = control.DefaultOptions()
	)

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       store.BackupStorer
		rc       inject.RestoreConsumer
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kw, sw, ctrl, nil, assert.NoError},
		{"missing kopia", nil, sw, ctrl, nil, assert.Error},
		{"missing modelstore", kw, nil, ctrl, nil, assert.Error},
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
				tconfig.NewM365Account(t),
				"backup-id",
				selectors.Selector{DiscreteOwner: "test"},
				restoreCfg,
				evmock.NewBus(),
				count.New())
			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}

func (suite *RestoreOpIntegrationSuite) TestRestore_Run_errorNoBackup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		restoreCfg = testdata.DefaultRestoreConfig("")
		mb         = evmock.NewBus()
	)

	rsel := selectors.NewExchangeRestore(selectors.None())
	rsel.Include(rsel.AllData())

	ctrl, err := m365.NewController(
		ctx,
		suite.acct,
		resource.Users,
		rsel.PathService(),
		control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	ro, err := NewRestoreOperation(
		ctx,
		control.DefaultOptions(),
		suite.kw,
		suite.sw,
		ctrl,
		tconfig.NewM365Account(t),
		"backupID",
		rsel.Selector,
		restoreCfg,
		mb,
		count.New())
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
