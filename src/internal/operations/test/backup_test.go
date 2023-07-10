package test_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
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
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// singleton
// ---------------------------------------------------------------------------

type backupInstance struct {
	obo *operations.BackupOperation
	bod *backupOpDependencies
	// forms a linked list of incremental backups
	incremental *backupInstance
}

func (bi *backupInstance) close(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	if bi.incremental != nil {
		bi.incremental.close(t, ctx)
	}

	bi.bod.close(t, ctx)
}

func (bi *backupInstance) runAndCheckBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	mb *evmock.Bus,
	acceptNoData bool,
) {
	err := bi.obo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bi.obo.Results, "the backup had non-zero results")
	require.NotEmpty(t, bi.obo.Results.BackupID, "the backup generated an ID")

	checkBackup(t, *bi.obo, mb, acceptNoData)
}

func (bi *backupInstance) runAndCheckIncrementalBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	mb *evmock.Bus,
) *backupInstance {
	// bi.incremental = prepNewTestBackupOp(t, ctx, mb, bi.bod.sel, bi.obo.Options, bi.obo.BackupVersion)
	incremental := &backupInstance{
		bod: &backupOpDependencies{},
	}

	// copy the old bod connection references
	*incremental.bod = *bi.bod

	// generate a new controller to avoid statefulness
	incremental.bod.renewController(t, ctx)

	incremental.obo = newTestBackupOp(
		t,
		ctx,
		incremental.bod,
		mb,
		bi.obo.Options)

	err := bi.incremental.obo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bi.incremental.obo.Results, "the incremental backup had non-zero results")
	require.NotEmpty(t, bi.incremental.obo.Results.BackupID, "the incremental backup generated an ID")

	return bi.incremental
}

func checkBackup(
	t *testing.T,
	obo operations.BackupOperation,
	mb *evmock.Bus,
	acceptNoData bool,
) {
	expectStatus := []operations.OpStatus{operations.Completed}
	if acceptNoData {
		expectStatus = append(expectStatus, operations.NoData)
	}

	require.Contains(
		t,
		expectStatus,
		obo.Status,
		"backup doesn't match expectation, wanted any of %v, got %s",
		expectStatus,
		obo.Status)

	require.Less(t, 0, obo.Results.ItemsWritten)
	assert.Less(t, 0, obo.Results.ItemsRead, "count of items read")
	assert.Less(t, int64(0), obo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), obo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, obo.Results.ResourceOwners, "count of resource owners")
	assert.NoErrorf(
		t,
		obo.Errors.Failure(),
		"incremental non-recoverable error %+v",
		clues.ToCore(obo.Errors.Failure()))
	assert.Empty(t, obo.Errors.Recovered(), "incremental recoverable/iteration errors")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		obo.Results.BackupID, "incremental pre-run backupID event")
}

func checkIncrementalBackup(
	t *testing.T,
	obo operations.BackupOperation,
	mb *evmock.Bus,
) {
	assert.NoError(t, obo.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(obo.Errors.Failure()))
	assert.Empty(t, obo.Errors.Recovered(), "incremental recoverable/iteration errors")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "incremental backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "incremental backup-end events")
	// FIXME: commented tests are flaky due to delta calls retaining data that is
	// out of scope of the test data.
	// we need to find a better way to make isolated assertions here.
	// The addition of the deeTD package gives us enough coverage to comment
	// out the tests for now and look to their improvemeng later.

	// do some additional checks to ensure the incremental dealt with fewer items.
	// +4 on read/writes to account for metadata: 1 delta and 1 path for each type.
	// if !toggles.DisableDelta {
	// assert.Equal(t, test.deltaItemsRead+4, incBO.Results.ItemsRead, "incremental items read")
	// assert.Equal(t, test.deltaItemsWritten+4, incBO.Results.ItemsWritten, "incremental items written")
	// } else {
	// assert.Equal(t, test.nonDeltaItemsRead+4, incBO.Results.ItemsRead, "non delta items read")
	// assert.Equal(t, test.nonDeltaItemsWritten+4, incBO.Results.ItemsWritten, "non delta items written")
	// }
	// assert.Equal(t, test.nonMetaItemsWritten, incBO.Results.ItemsWritten, "non meta incremental items write")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		obo.Results.BackupID, "incremental pre-run backupID event")
}

// ---------------------------------------------------------------------------
// initialization and dependencies
// ---------------------------------------------------------------------------

type backupOpDependencies struct {
	acct account.Account
	ctrl *m365.Controller
	kms  *kopia.ModelStore
	kw   *kopia.Wrapper
	sel  selectors.Selector
	sss  streamstore.Streamer
	st   storage.Storage
	sw   *store.Wrapper

	closer func()
}

// prepNewTestBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
func prepNewTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bus events.Eventer,
	sel selectors.Selector,
	opts control.Options,
	backupVersion int,
) *backupInstance {
	bod := &backupOpDependencies{
		acct: tconfig.NewM365Account(t),
		st:   storeTD.NewPrefixedS3Storage(t),
	}

	k := kopia.NewConn(bod.st)

	err := k.Initialize(ctx, repository.Options{})
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		if err != nil {
			bod.close(t, ctx)
			t.FailNow()
		}
	}()

	// kopiaRef comes with a count of 1 and Wrapper bumps it again
	// we're so safe to close here.
	bod.closer = func() {
		err := k.Close(ctx)
		assert.NoErrorf(t, err, "k close: %+v", clues.ToCore(err))
	}

	bod.kw, err = kopia.NewWrapper(k)
	require.NoError(t, err, clues.ToCore(err))

	bod.kms, err = kopia.NewModelStore(k)
	require.NoError(t, err, clues.ToCore(err))

	bod.sw = store.NewKopiaStore(bod.kms)

	connectorResource := resource.Users
	if sel.Service == selectors.ServiceSharePoint {
		connectorResource = resource.Sites
	}

	bod.ctrl, bod.sel = ControllerWithSelector(
		t,
		ctx,
		bod.acct,
		connectorResource,
		sel,
		nil,
		bod.close)

	obo := newTestBackupOp(
		t,
		ctx,
		bod,
		bus,
		opts)

	bod.sss = streamstore.NewStreamer(
		bod.kw,
		bod.acct.ID(),
		bod.sel.PathService())

	return &backupInstance{
		obo: obo,
		bod: bod,
	}
}

func (bod *backupOpDependencies) close(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	bod.closer()

	if bod.kw != nil {
		err := bod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kw close: %+v", clues.ToCore(err))
	}

	if bod.kms != nil {
		err := bod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kms close: %+v", clues.ToCore(err))
	}
}

// generates a new controller, and replaces bod.ctrl with that instance.
// useful for clearing controller state between runs.
func (bod *backupOpDependencies) renewController(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	rc := resource.Users
	if bod.sel.PathService() == path.SharePointService {
		rc = resource.Sites
	}

	newCtrl, err := m365.NewController(
		ctx,
		bod.acct,
		rc,
		bod.sel.PathService(),
		control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	bod.ctrl = newCtrl
}

// newTestBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
func newTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bod *backupOpDependencies,
	bus events.Eventer,
	opts control.Options,
) *operations.BackupOperation {
	bod.ctrl.IDNameLookup = idname.NewCache(map[string]string{bod.sel.ID(): bod.sel.Name()})

	bo, err := operations.NewBackupOperation(
		ctx,
		opts,
		bod.kw,
		bod.sw,
		bod.ctrl,
		bod.acct,
		bod.sel,
		bod.sel,
		bus)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		bod.close(t, ctx)
		t.FailNow()
	}

	return &bo
}

func checkBackupIsInManifests(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bod *backupOpDependencies,
	bo *operations.BackupOperation,
	sel selectors.Selector,
	resourceOwner string,
	categories ...path.CategoryType,
) {
	for _, category := range categories {
		t.Run("backup_in_manifests_"+category.String(), func(t *testing.T) {
			var (
				reasons = []kopia.Reason{
					{
						ResourceOwner: resourceOwner,
						Service:       sel.PathService(),
						Category:      category,
					},
				}
				tags  = map[string]string{kopia.TagBackupCategory: ""}
				found bool
			)

			bf, err := bod.kw.NewBaseFinder(bod.sw)
			require.NoError(t, err, clues.ToCore(err))

			fmt.Printf("\n-----\nR %+v\nT %+v\n-----\n", reasons, tags)

			mans := bf.FindBases(ctx, reasons, tags)

			mmb := mans.MergeBases()
			require.NotEmpty(t, mmb, "should find at least one merge base")

			t.Log("Backup IDs from merge bases:")

			for _, man := range mmb {
				bID, ok := man.GetTag(kopia.TagBackupID)
				if !assert.Truef(t, ok, "snapshot manifest %s missing backup ID tag", man.ID) {
					continue
				}

				t.Log("-", bID)

				if bID == string(bo.Results.BackupID) {
					found = true
					break
				}
			}

			assert.True(t, found, "backup %q retrieved by previous snapshot manifest", bo.Results.BackupID)
		})
	}
}

func checkMetadataFilesExist(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupID model.StableID,
	bod *backupOpDependencies,
	tenant, resourceOwner string,
	service path.ServiceType,
	filesByCat map[path.CategoryType][]string,
) {
	for category, files := range filesByCat {
		t.Run("metadata_files_exist_"+category.String(), func(t *testing.T) {
			bup := &backup.Backup{}

			err := bod.kms.Get(ctx, model.BackupSchema, backupID, bup)
			if !assert.NoError(t, err, clues.ToCore(err)) {
				return
			}

			paths := []path.RestorePaths{}
			pathsByRef := map[string][]string{}

			for _, fName := range files {
				p, err := path.Builder{}.
					Append(fName).
					ToServiceCategoryMetadataPath(tenant, resourceOwner, service, category, true)
				if !assert.NoError(t, err, "bad metadata path", clues.ToCore(err)) {
					continue
				}

				dir, err := p.Dir()
				if !assert.NoError(t, err, "parent path", clues.ToCore(err)) {
					continue
				}

				paths = append(
					paths,
					path.RestorePaths{StoragePath: p, RestorePath: dir})
				pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fName)
			}

			cols, err := bod.kw.ProduceRestoreCollections(
				ctx,
				bup.SnapshotID,
				paths,
				nil,
				fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))

			for _, col := range cols {
				itemNames := []string{}

				for item := range col.Items(ctx, fault.New(true)) {
					assert.Implements(t, (*data.StreamSize)(nil), item)

					s := item.(data.StreamSize)
					assert.Greaterf(
						t,
						s.Size(),
						int64(0),
						"empty metadata file: %s/%s",
						col.FullPath(),
						item.UUID(),
					)

					itemNames = append(itemNames, item.UUID())
				}

				assert.ElementsMatchf(
					t,
					pathsByRef[col.FullPath().ShortRef()],
					itemNames,
					"collection %s missing expected files",
					col.FullPath(),
				)
			}
		})
	}
}
