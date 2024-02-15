package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	strTD "github.com/alcionai/corso/src/internal/common/str/testdata"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

type BackupOpDependencies struct {
	Acct account.Account
	Ctrl *m365.Controller
	KMS  *kopia.ModelStore
	KW   *kopia.Wrapper
	Sel  selectors.Selector
	SSS  streamstore.Streamer
	St   storage.Storage
	SW   store.BackupStorer

	closer func()
}

func (bod *BackupOpDependencies) Close(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	bod.closer()

	if bod.KW != nil {
		err := bod.KW.Close(ctx)
		assert.NoErrorf(t, err, "kw close: %+v", clues.ToCore(err))
	}

	if bod.KMS != nil {
		err := bod.KW.Close(ctx)
		assert.NoErrorf(t, err, "kms close: %+v", clues.ToCore(err))
	}
}

// PrepNewTestBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
func PrepNewTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bus events.Eventer,
	sel selectors.Selector,
	opts control.Options,
	backupVersion int,
	counter *count.Bus,
) (
	operations.BackupOperation,
	*BackupOpDependencies,
) {
	bod := &BackupOpDependencies{
		Acct: tconfig.NewM365Account(t),
		St:   storeTD.NewPrefixedS3Storage(t),
	}
	repoNameHash := strTD.NewHashForRepoConfigName()

	k := kopia.NewConn(bod.St)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		if err != nil {
			bod.Close(t, ctx)
			t.FailNow()
		}
	}()

	// kopiaRef comes with a count of 1 and Wrapper bumps it again
	// we're so safe to close here.
	bod.closer = func() {
		err := k.Close(ctx)
		assert.NoErrorf(t, err, "k close: %+v", clues.ToCore(err))
	}

	bod.KW, err = kopia.NewWrapper(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.BackupOperation{}, nil
	}

	bod.KMS, err = kopia.NewModelStore(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.BackupOperation{}, nil
	}

	bod.SW = store.NewWrapper(bod.KMS)

	bod.Ctrl, bod.Sel = ControllerWithSelector(
		t,
		ctx,
		bod.Acct,
		sel,
		nil,
		bod.Close,
		counter)

	bo := NewTestBackupOp(
		t,
		ctx,
		bod,
		bus,
		opts,
		counter)
	bo.BackupVersion = backupVersion

	bod.SSS = streamstore.NewStreamer(
		bod.KW,
		bod.Acct.ID(),
		bod.Sel.PathService())

	return bo, bod
}

// NewTestBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
func NewTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bod *BackupOpDependencies,
	bus events.Eventer,
	opts control.Options,
	counter *count.Bus,
) operations.BackupOperation {
	bod.Ctrl.IDNameLookup = idname.NewCache(map[string]string{bod.Sel.ID(): bod.Sel.Name()})

	bo, err := operations.NewBackupOperation(
		ctx,
		opts,
		bod.KW,
		bod.SW,
		bod.Ctrl,
		bod.Acct,
		bod.Sel,
		bod.Sel,
		bus,
		counter)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		bod.Close(t, ctx)
		t.FailNow()
	}

	return bo
}

func RunAndCheckBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bo *operations.BackupOperation,
	mb *evmock.Bus,
	acceptNoData bool,
) {
	err := bo.Run(ctx)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		for i, err := range bo.Errors.Recovered() {
			t.Logf("recoverable err %d, %+v", i, err)
		}

		assert.Fail(t, "not allowed to error")
	}

	require.NotEmpty(t, bo.Results, "the backup had non-zero results")
	require.NotEmpty(t, bo.Results.BackupID, "the backup generated an ID")

	expectStatus := []operations.OpStatus{operations.Completed}
	if acceptNoData {
		expectStatus = append(expectStatus, operations.NoData)
	}

	require.Contains(
		t,
		expectStatus,
		bo.Status,
		"backup doesn't match expectation, wanted any of %v, got %s",
		expectStatus,
		bo.Status)

	require.NotZero(t, bo.Results.ItemsWritten)
	assert.NotZero(t, bo.Results.ItemsRead, "count of items read")
	assert.NotZero(t, bo.Results.BytesRead, "bytes read")
	assert.NotZero(t, bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners, "count of resource owners")
	assert.NoError(t, bo.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(bo.Errors.Failure()))
	assert.Empty(t, bo.Errors.Recovered(), "incremental recoverable/iteration errors")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
}

func CheckBackupIsInManifests(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	kw *kopia.Wrapper,
	sw store.BackupStorer,
	bo *operations.BackupOperation,
	sel selectors.Selector,
	resourceOwner string,
	categories ...path.CategoryType,
) {
	for _, category := range categories {
		t.Run(category.String(), func(t *testing.T) {
			var (
				r     = identity.NewReason("", resourceOwner, sel.PathService(), category)
				tags  = map[string]string{kopia.TagBackupCategory: ""}
				found bool
			)

			bf, err := kw.NewBaseFinder(sw)
			require.NoError(t, err, clues.ToCore(err))

			mans := bf.FindBases(ctx, []identity.Reasoner{r}, tags)
			for _, man := range mans.MergeBases() {
				bID, ok := man.GetSnapshotTag(kopia.TagBackupID)
				if !assert.Truef(t, ok, "snapshot manifest %s missing backup ID tag", man.ItemDataSnapshot.ID) {
					continue
				}

				if bID == string(bo.Results.BackupID) {
					found = true
					break
				}
			}

			assert.True(t, found, "backup retrieved by previous snapshot manifest")
		})
	}
}

func RunMergeBaseGroupsUpdate(
	suite tester.Suite,
	sel selectors.Selector,
	expectCached bool,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		opts    = control.DefaultOptions()
		whatSet = deeTD.CategoryFromRepoRef
	)

	opts.ToggleFeatures.UseDeltaTree = true

	// Need outside the inner test case so bod lasts for the entire test.
	bo, bod := PrepNewTestBackupOp(
		t,
		ctx,
		mb,
		sel,
		opts,
		version.All8MigrateUserPNToID,
		count.New())
	defer bod.Close(t, ctx)

	suite.Run("makeMergeBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		RunAndCheckBackup(t, ctx, &bo, mb, false)

		reasons, err := bod.Sel.Reasons(bod.Acct.ID(), false)
		require.NoError(t, err, clues.ToCore(err))

		for _, reason := range reasons {
			CheckBackupIsInManifests(
				t,
				ctx,
				bod.KW,
				bod.SW,
				&bo,
				bod.Sel,
				bod.Sel.ID(),
				reason.Category())
		}

		_, expectDeets := deeTD.GetDeetsInBackup(
			t,
			ctx,
			bo.Results.BackupID,
			bod.Acct.ID(),
			bod.Sel.ID(),
			bod.Sel.PathService(),
			whatSet,
			bod.KMS,
			bod.SSS)
		deeTD.CheckBackupDetails(
			t,
			ctx,
			bo.Results.BackupID,
			whatSet,
			bod.KMS,
			bod.SSS,
			expectDeets,
			false)
	})

	suite.Run("makeIncrementalBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		var (
			mb   = evmock.NewBus()
			opts = control.DefaultOptions()
		)

		forcedFull := NewTestBackupOp(
			t,
			ctx,
			bod,
			mb,
			opts,
			count.New())
		forcedFull.BackupVersion = version.Groups9Update

		RunAndCheckBackup(t, ctx, &forcedFull, mb, false)

		reasons, err := bod.Sel.Reasons(bod.Acct.ID(), false)
		require.NoError(t, err, clues.ToCore(err))

		for _, reason := range reasons {
			CheckBackupIsInManifests(
				t,
				ctx,
				bod.KW,
				bod.SW,
				&forcedFull,
				bod.Sel,
				bod.Sel.ID(),
				reason.Category())
		}

		_, expectDeets := deeTD.GetDeetsInBackup(
			t,
			ctx,
			forcedFull.Results.BackupID,
			bod.Acct.ID(),
			bod.Sel.ID(),
			bod.Sel.PathService(),
			whatSet,
			bod.KMS,
			bod.SSS)
		deeTD.CheckBackupDetails(
			t,
			ctx,
			forcedFull.Results.BackupID,
			whatSet,
			bod.KMS,
			bod.SSS,
			expectDeets,
			false)

		check := assert.Zero

		if expectCached {
			check = assert.NotZero
		}

		check(
			t,
			forcedFull.Results.Counts[string(count.PersistedCachedFiles)],
			"cached items")
	})
}

func RunBasicBackupTest(
	suite tester.Suite,
	sel selectors.Selector,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
		whatSet = deeTD.CategoryFromRepoRef
	)

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	reasons, err := bod.Sel.Reasons(bod.Acct.ID(), false)
	require.NoError(t, err, clues.ToCore(err))

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	for _, reason := range reasons {
		CheckBackupIsInManifests(
			t,
			ctx,
			bod.KW,
			bod.SW,
			&bo,
			bod.Sel,
			bod.Sel.ID(),
			reason.Category())
	}

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bo.Results.BackupID,
		bod.Acct.ID(),
		bod.Sel.ID(),
		sel.PathService(),
		whatSet,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		whatSet,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)

	// Basic, happy path incremental test.  No changes are dictated or expected.
	// This only tests that an incremental backup is runnable at all, and that it
	// produces fewer results than the last backup.
	//
	// Incremental testing for conversations is limited because of API restrictions.
	// Since graph doesn't provide us a way to programmatically delete conversations,
	// or create new conversations without a delegated token, we can't do incremental
	// testing with newly added items.
	incMB := evmock.NewBus()
	incBO := NewTestBackupOp(
		t,
		ctx,
		bod,
		incMB,
		opts,
		count.New())

	RunAndCheckBackup(t, ctx, &incBO, incMB, true)

	for _, reason := range reasons {
		CheckBackupIsInManifests(
			t,
			ctx,
			bod.KW,
			bod.SW,
			&incBO,
			bod.Sel,
			bod.Sel.ID(),
			reason.Category())
	}

	_, expectDeets = deeTD.GetDeetsInBackup(
		t,
		ctx,
		incBO.Results.BackupID,
		bod.Acct.ID(),
		bod.Sel.ID(),
		bod.Sel.PathService(),
		whatSet,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		incBO.Results.BackupID,
		whatSet,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)

	assert.NotZero(
		t,
		incBO.Results.Counts[string(count.PersistedCachedFiles)],
		"cached items")
	assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
	assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
	assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
}
