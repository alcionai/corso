package test_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/store"
)

type restoreOpDependencies struct {
	acct account.Account
	ctrl *m365.Controller
	kms  *kopia.ModelStore
	kw   *kopia.Wrapper
	sel  selectors.Selector
	sss  streamstore.Streamer
	st   storage.Storage
	sw   store.BackupStorer

	closer func()
}

func (rod *restoreOpDependencies) close(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	if rod.closer != nil {
		rod.closer()
	}

	if rod.kw != nil {
		err := rod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kw close: %+v", clues.ToCore(err))
	}

	if rod.kms != nil {
		err := rod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kms close: %+v", clues.ToCore(err))
	}
}

// prepNewTestRestoreOp generates all clients required to run a restore operation,
// returning both a restore operation created with those clients, as well as
// the clients themselves.
func prepNewTestRestoreOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupStore storage.Storage,
	backupID model.StableID,
	bus events.Eventer,
	ctr *count.Bus,
	sel selectors.Selector,
	opts control.Options,
	restoreCfg control.RestoreConfig,
) (
	operations.RestoreOperation,
	*restoreOpDependencies,
) {
	var (
		rod = &restoreOpDependencies{
			acct: tconfig.NewM365Account(t),
			st:   backupStore,
		}
		k = kopia.NewConn(rod.st)
	)

	err := k.Connect(ctx, repository.Options{})
	require.NoError(t, err, clues.ToCore(err))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again
	// we're so safe to close here.
	defer func() {
		err := k.Close(ctx)
		assert.NoErrorf(t, err, "k close: %+v", clues.ToCore(err))
	}()

	rod.kw, err = kopia.NewWrapper(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.RestoreOperation{}, rod
	}

	rod.kms, err = kopia.NewModelStore(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.RestoreOperation{}, rod
	}

	rod.sw = store.NewWrapper(rod.kms)

	connectorResource := resource.Users

	switch sel.Service {
	case selectors.ServiceSharePoint:
		connectorResource = resource.Sites
	case selectors.ServiceGroups:
		connectorResource = resource.Groups
	}

	rod.ctrl, rod.sel = ControllerWithSelector(
		t,
		ctx,
		rod.acct,
		connectorResource,
		sel,
		nil,
		rod.close)

	ro := newTestRestoreOp(
		t,
		ctx,
		rod,
		backupID,
		bus,
		ctr,
		opts,
		restoreCfg)

	rod.sss = streamstore.NewStreamer(
		rod.kw,
		rod.acct.ID(),
		rod.sel.PathService())

	return ro, rod
}

// newTestRestoreOp accepts the clients required to compose a restore operation, plus
// any other metadata, and uses them to generate a new restore operation.  This
// allows restore chains to utilize the same temp directory and configuration
// details.
func newTestRestoreOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	rod *restoreOpDependencies,
	backupID model.StableID,
	bus events.Eventer,
	ctr *count.Bus,
	opts control.Options,
	restoreCfg control.RestoreConfig,
) operations.RestoreOperation {
	rod.ctrl.IDNameLookup = idname.NewCache(map[string]string{rod.sel.ID(): rod.sel.Name()})

	ro, err := operations.NewRestoreOperation(
		ctx,
		opts,
		rod.kw,
		rod.sw,
		rod.ctrl,
		rod.acct,
		backupID,
		rod.sel,
		restoreCfg,
		bus,
		ctr)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		rod.close(t, ctx)
		t.FailNow()
	}

	return ro
}

func runAndCheckRestore(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	ro *operations.RestoreOperation,
	mb *evmock.Bus,
	acceptNoData bool,
) *details.Details {
	deets, err := ro.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, ro.Results, "the restore had non-zero results")
	require.NotNil(t, deets, "restore details")

	expectStatus := []operations.OpStatus{operations.Completed}
	if acceptNoData {
		expectStatus = append(expectStatus, operations.NoData)
	}

	require.Contains(
		t,
		expectStatus,
		ro.Status,
		"restore doesn't match expectation, wanted any of %v, got %s",
		expectStatus,
		ro.Status)

	assert.NoError(t, ro.Errors.Failure(), "non-recoverable error", clues.ToCore(ro.Errors.Failure()))

	if assert.Empty(t, ro.Errors.Recovered(), "recoverable/iteration errors") {
		allErrs := ro.Errors.Errors()
		for i, err := range allErrs.Recovered {
			t.Log("recovered from test err", i, err)
		}
	}

	assert.NotZero(t, ro.Results.ItemsRead, "count of items read")
	assert.NotZero(t, ro.Results.BytesRead, "bytes read")
	assert.Equal(t, 1, ro.Results.ResourceOwners, "count of resource owners")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreStart], "restore-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.RestoreEnd], "restore-end events")

	return deets
}

type GetItemsInContainerByCollisionKeyer[T any] interface {
	GetItemsInContainerByCollisionKey(
		ctx context.Context,
		userID, containerID string,
	) (map[string]T, error)
}

func filterCollisionKeyResults[T any](
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	protectedResourceID, containerID string,
	giicbck GetItemsInContainerByCollisionKeyer[T],
	filterOut map[string]T,
) map[string]T {
	m, err := giicbck.GetItemsInContainerByCollisionKey(
		ctx,
		protectedResourceID,
		containerID)
	require.NoError(t, err, clues.ToCore(err))

	for k := range filterOut {
		delete(m, k)
	}

	return m
}

func checkRestoreCounts(
	t *testing.T,
	ctr *count.Bus,
	expectSkips, expectReplaces, expectNew int,
) {
	t.Log("counted values", ctr.Values())
	t.Log("counted totals", ctr.TotalValues())

	if expectSkips >= 0 {
		assert.Equal(
			t,
			int64(expectSkips),
			ctr.Total(count.CollisionSkip),
			"count of collisions resolved by skip")
	}

	if expectReplaces >= 0 {
		assert.Equal(
			t,
			int64(expectReplaces),
			ctr.Total(count.CollisionReplace),
			"count of collisions resolved by replace")
	}

	if expectNew >= 0 {
		assert.Equal(
			t,
			int64(expectNew),
			ctr.Total(count.NewItemCreated),
			"count of new items or collisions resolved by copy")
	}
}
