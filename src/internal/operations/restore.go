package operations

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/operations/pathtransformer"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	BackupID   model.StableID
	Results    RestoreResults
	Selectors  selectors.Selector
	RestoreCfg control.RestoreConfig
	Version    string

	acct account.Account
	rc   inject.RestoreConsumer
}

// RestoreResults aggregate the details of the results of the operation.
type RestoreResults struct {
	stats.ReadWrites
	stats.StartAndEndTime
}

// NewRestoreOperation constructs and validates a restore operation.
func NewRestoreOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	sw store.BackupStorer,
	rc inject.RestoreConsumer,
	acct account.Account,
	backupID model.StableID,
	sel selectors.Selector,
	restoreCfg control.RestoreConfig,
	bus events.Eventer,
	ctr *count.Bus,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation:  newOperation(opts, bus, ctr, kw, sw),
		acct:       acct,
		BackupID:   backupID,
		RestoreCfg: control.EnsureRestoreConfigDefaults(ctx, restoreCfg),
		Selectors:  sel,
		Version:    "v0",
		rc:         rc,
	}
	if err := op.validate(); err != nil {
		return RestoreOperation{}, err
	}

	return op, nil
}

func (op RestoreOperation) validate() error {
	if op.rc == nil {
		return clues.New("missing restore consumer")
	}

	return op.operation.validate()
}

// aggregates stats from the restore.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type restoreStats struct {
	cs            []data.RestoreCollection
	ctrl          *data.CollectionStats
	bytesRead     *stats.ByteCounter
	resourceCount int

	// a transient value only used to pair up start-end events.
	restoreID string
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (restoreDetails *details.Details, err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "restore"); crErr != nil {
			err = crErr
		}
	}()

	var (
		opStats = restoreStats{
			bytesRead: &stats.ByteCounter{},
			restoreID: uuid.NewString(),
		}
		start  = time.Now()
		sstore = streamstore.NewStreamer(op.kopia, op.acct.ID(), op.Selectors.PathService())
	)

	// -----
	// Setup
	// -----

	ctx, end := diagnostics.Span(ctx, "operations:restore:run")
	defer func() {
		end()
	}()

	ctx, flushMetrics := events.NewMetrics(ctx, logger.Writer{Ctx: ctx})
	defer flushMetrics()

	ctx = clues.Add(
		ctx,
		"tenant_id", clues.Hide(op.acct.ID()),
		"backup_id", op.BackupID,
		"service", op.Selectors.Service,
		"destination_container", clues.Hide(op.RestoreCfg.Location))

	defer func() {
		op.bus.Event(
			ctx,
			events.RestoreEnd,
			map[string]any{
				events.BackupID:      op.BackupID,
				events.DataRetrieved: op.Results.BytesRead,
				events.Duration:      op.Results.CompletedAt.Sub(op.Results.StartedAt),
				events.EndTime:       dttm.Format(op.Results.CompletedAt),
				events.ItemsRead:     op.Results.ItemsRead,
				events.ItemsWritten:  op.Results.ItemsWritten,
				events.Resources:     op.Results.ResourceOwners,
				events.RestoreID:     opStats.restoreID,
				events.Service:       op.Selectors.Service.String(),
				events.StartTime:     dttm.Format(op.Results.StartedAt),
				events.Status:        op.Status.String(),
			})
	}()

	// -----
	// Execution
	// -----

	deets, err := op.do(ctx, &opStats, sstore, start)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.CtxErr(ctx, err).Error("running restore")

		if errors.Is(err, kopia.ErrNoRestorePath) {
			op.Errors.Fail(clues.Wrap(err, "empty backup or unknown path provided"))
		}

		op.Errors.Fail(clues.Wrap(err, "running restore"))
	}

	finalizeErrorHandling(ctx, op.Options, op.Errors, "running restore")
	LogFaultErrors(ctx, op.Errors.Errors(), "running restore")
	logger.Ctx(ctx).With("total_counts", op.Counter.Values()).Info("restore stats")

	// -----
	// Persistence
	// -----

	err = op.persistResults(ctx, start, &opStats)
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting restore results"))
		return nil, op.Errors.Failure()
	}

	logger.Ctx(ctx).Infow("completed restore", "results", op.Results)

	return deets, nil
}

func (op *RestoreOperation) do(
	ctx context.Context,
	opStats *restoreStats,
	detailsStore streamstore.Reader,
	start time.Time,
) (*details.Details, error) {
	logger.Ctx(ctx).
		With("control_options", op.Options, "selectors", op.Selectors).
		Info("restoring selection")

	bup, deets, err := getBackupAndDetailsFromID(
		ctx,
		op.BackupID,
		op.store,
		detailsStore,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "getting backup and details")
	}

	restoreToProtectedResource, err := chooseRestoreResource(ctx, op.rc, op.RestoreCfg, bup.Selector)
	if err != nil {
		return nil, clues.Wrap(err, "getting destination protected resource")
	}

	ctx = clues.Add(
		ctx,
		"backup_protected_resource_id", bup.Selector.ID(),
		"backup_protected_resource_name", clues.Hide(bup.Selector.Name()),
		"restore_protected_resource_id", restoreToProtectedResource.ID(),
		"restore_protected_resource_name", clues.Hide(restoreToProtectedResource.Name()))

	// IsServiceEnabled checks if the resource has the service enabled to be able to restore.
	runnable, err := op.rc.IsServiceEnabled(
		ctx,
		op.Selectors.PathService(),
		restoreToProtectedResource.ID())
	if err != nil {
		logger.CtxErr(ctx, err).Error("verifying restore is runnable")
		op.Errors.Fail(clues.Wrap(err, "verifying restore is runnable"))

		return nil, clues.Stack(err).WithClues(ctx)
	}

	if !runnable {
		logger.CtxErr(ctx, graph.ErrServiceNotEnabled).Error("checking if restore is enabled")
		op.Errors.Fail(clues.Wrap(err, "checking if restore is enabled"))

		return nil, clues.Stack(graph.ErrServiceNotEnabled).WithClues(ctx)
	}

	observe.Message(ctx, "Restoring", observe.Bullet, clues.Hide(restoreToProtectedResource.Name()))

	paths, err := formatDetailsForRestoration(
		ctx,
		bup.Version,
		op.Selectors,
		deets,
		op.rc,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "formatting paths from details")
	}

	ctx = clues.Add(
		ctx,
		"details_entries", len(deets.Entries),
		"details_paths", len(paths),
		"backup_snapshot_id", bup.SnapshotID,
		"backup_version", bup.Version)

	op.bus.Event(
		ctx,
		events.RestoreStart,
		map[string]any{
			events.StartTime:        start,
			events.BackupID:         op.BackupID,
			events.BackupCreateTime: bup.CreationTime,
			events.RestoreID:        opStats.restoreID,
		})

	observe.Message(ctx, fmt.Sprintf("Discovered %d items in backup %s to restore", len(paths), op.BackupID))

	progressBar := observe.MessageWithCompletion(ctx, "Enumerating items in repository")
	defer close(progressBar)

	dcs, err := op.kopia.ProduceRestoreCollections(
		ctx,
		bup.SnapshotID,
		paths,
		opStats.bytesRead,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "producing collections to restore")
	}

	ctx = clues.Add(ctx, "coll_count", len(dcs))

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1
	opStats.cs = dcs

	deets, err = consumeRestoreCollections(
		ctx,
		op.rc,
		bup.Version,
		restoreToProtectedResource,
		op.Selectors,
		op.RestoreCfg,
		op.Options,
		dcs,
		op.Errors,
		op.Counter)
	if err != nil {
		return nil, clues.Stack(err)
	}

	opStats.ctrl = op.rc.Wait()

	logger.Ctx(ctx).Debug(opStats.ctrl)

	return deets, nil
}

// persists details and statistics about the restore operation.
func (op *RestoreOperation) persistResults(
	ctx context.Context,
	started time.Time,
	opStats *restoreStats,
) error {
	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	op.Status = Completed

	if op.Errors.Failure() != nil {
		op.Status = Failed
	}

	op.Results.BytesRead = opStats.bytesRead.NumBytes
	op.Results.ItemsRead = len(opStats.cs) // TODO: file count, not collection count
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.ctrl == nil {
		op.Status = Failed
		return clues.New("restoration never completed")
	}

	if op.Status != Failed && opStats.ctrl.IsZero() {
		op.Status = NoData
	}

	op.Results.ItemsWritten = opStats.ctrl.Successes

	return op.Errors.Failure()
}

func chooseRestoreResource(
	ctx context.Context,
	pprian inject.PopulateProtectedResourceIDAndNamer,
	restoreCfg control.RestoreConfig,
	orig idname.Provider,
) (idname.Provider, error) {
	if len(restoreCfg.ProtectedResource) == 0 {
		return orig, nil
	}

	id, name, err := pprian.PopulateProtectedResourceIDAndName(
		ctx,
		restoreCfg.ProtectedResource,
		nil)

	return idname.NewProvider(id, name), clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// Restorer funcs
// ---------------------------------------------------------------------------

func consumeRestoreCollections(
	ctx context.Context,
	rc inject.RestoreConsumer,
	backupVersion int,
	toProtectedResource idname.Provider,
	sel selectors.Selector,
	restoreCfg control.RestoreConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, error) {
	progressBar := observe.MessageWithCompletion(ctx, "Restoring data")
	defer close(progressBar)

	rcc := inject.RestoreConsumerConfig{
		BackupVersion:     backupVersion,
		Options:           opts,
		ProtectedResource: toProtectedResource,
		RestoreConfig:     restoreCfg,
		Selector:          sel,
	}

	deets, err := rc.ConsumeRestoreCollections(ctx, rcc, dcs, errs, ctr)
	if err != nil {
		return nil, clues.Wrap(err, "restoring collections")
	}

	return deets, nil
}

// formatDetailsForRestoration reduces the provided detail entries according to the
// selector specifications.
func formatDetailsForRestoration(
	ctx context.Context,
	backupVersion int,
	sel selectors.Selector,
	deets *details.Details,
	cii inject.CacheItemInfoer,
	errs *fault.Bus,
) ([]path.RestorePaths, error) {
	fds, err := sel.Reduce(ctx, deets, errs)
	if err != nil {
		return nil, err
	}

	// allow restore controllers to iterate over item metadata
	for _, ent := range fds.Entries {
		cii.CacheItemInfo(ent.ItemInfo)
	}

	paths, err := pathtransformer.GetPaths(ctx, backupVersion, fds.Items(), errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting restore paths")
	}

	if sel.Service == selectors.ServiceOneDrive || sel.Service == selectors.ServiceSharePoint {
		paths, err = onedrive.AugmentRestorePaths(backupVersion, paths)
		if err != nil {
			return nil, clues.Wrap(err, "augmenting paths")
		}
	}

	return paths, nil
}
