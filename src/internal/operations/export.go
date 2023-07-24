package operations

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

const (
	// CopyBufferSize is the size of the copy buffer for disk
	// write operations
	// TODO(meain): tweak this value
	CopyBufferSize = 5 * 1024 * 1024
)

// ExportOperation wraps an operation with export-specific props.
type ExportOperation struct {
	operation

	BackupID  model.StableID
	Results   RestoreResults
	Selectors selectors.Selector
	ExportCfg control.ExportConfig
	Version   string

	acct account.Account
	ec   inject.ExportConsumer
}

// NewExportOperation constructs and validates a export operation.
func NewExportOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	ec inject.ExportConsumer,
	acct account.Account,
	backupID model.StableID,
	sel selectors.Selector,
	exportCfg control.ExportConfig,
	bus events.Eventer,
) (ExportOperation, error) {
	op := ExportOperation{
		operation: newOperation(opts, bus, count.New(), kw, sw),
		acct:      acct,
		BackupID:  backupID,
		ExportCfg: exportCfg,
		Selectors: sel,
		Version:   "v0",
		ec:        ec,
	}
	if err := op.validate(); err != nil {
		return ExportOperation{}, err
	}

	return op, nil
}

func (op ExportOperation) validate() error {
	if op.ec == nil {
		return clues.New("missing export consumer")
	}

	return op.operation.validate()
}

// aggregates stats from the export.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type exportStats struct {
	cs            []data.RestoreCollection
	ctrl          *data.CollectionStats
	bytesRead     *stats.ByteCounter
	resourceCount int

	// a transient value only used to pair up start-end events.
	exportID string
}

// Run begins a synchronous export operation.
func (op *ExportOperation) Run(ctx context.Context) (
	expColl []export.Collection,
	err error,
) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "export"); crErr != nil {
			err = crErr
		}
	}()

	var (
		opStats = exportStats{
			bytesRead: &stats.ByteCounter{},
			exportID:  uuid.NewString(),
		}
		start  = time.Now()
		sstore = streamstore.NewStreamer(op.kopia, op.acct.ID(), op.Selectors.PathService())
	)

	// -----
	// Setup
	// -----

	ctx, end := diagnostics.Span(ctx, "operations:export:run")
	defer func() {
		end()
		// wait for the progress display to clean up
		observe.Complete()
	}()

	ctx, flushMetrics := events.NewMetrics(ctx, logger.Writer{Ctx: ctx})
	defer flushMetrics()

	ctx = clues.Add(
		ctx,
		"tenant_id", clues.Hide(op.acct.ID()),
		"backup_id", op.BackupID,
		"service", op.Selectors.Service)

	defer func() {
		op.bus.Event(
			ctx,
			events.ExportEnd,
			map[string]any{
				events.BackupID:      op.BackupID,
				events.DataRetrieved: op.Results.BytesRead,
				events.Duration:      op.Results.CompletedAt.Sub(op.Results.StartedAt),
				events.EndTime:       dttm.Format(op.Results.CompletedAt),
				events.ItemsRead:     op.Results.ItemsRead,
				events.ItemsWritten:  op.Results.ItemsWritten,
				events.Resources:     op.Results.ResourceOwners,
				events.ExportID:      opStats.exportID,
				events.Service:       op.Selectors.Service.String(),
				events.StartTime:     dttm.Format(op.Results.StartedAt),
				events.Status:        op.Status.String(),
			})
	}()

	// -----
	// Execution
	// -----

	expCollections, err := op.do(ctx, &opStats, sstore, start)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.CtxErr(ctx, err).Error("running export")

		if errors.Is(err, kopia.ErrNoRestorePath) {
			op.Errors.Fail(clues.New("empty backup or unknown path provided"))
		}

		op.Errors.Fail(clues.Wrap(err, "running export"))
	}

	finalizeErrorHandling(ctx, op.Options, op.Errors, "running export")
	LogFaultErrors(ctx, op.Errors.Errors(), "running export")

	// -----
	// Persistence
	// -----

	err = op.persistResults(ctx, start, &opStats)
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting export results"))
		return nil, op.Errors.Failure()
	}

	logger.Ctx(ctx).Infow("completed export", "results", op.Results)

	return expCollections, nil
}

type zipExportCol struct {
	reader io.ReadCloser
}

func (z zipExportCol) BasePath() string {
	return ""
}

func (z zipExportCol) Items(ctx context.Context) <-chan export.Item {
	rc := make(chan export.Item, 1)
	defer close(rc)

	rc <- export.Item{
		Data: export.ItemData{
			Name: "Corso_Export_" + dttm.FormatNow(dttm.HumanReadable) + ".zip",
			Body: z.reader,
		},
	}

	return rc
}

// zipExportCollection takes a list of export collections and zips
// them into a single collection.
func zipExportCollection(
	ctx context.Context,
	expCollections []export.Collection,
) (export.Collection, error) {
	if len(expCollections) == 0 {
		return nil, clues.New("no export collections provided")
	}

	reader, writer := io.Pipe()
	wr := zip.NewWriter(writer)

	go func() {
		defer writer.Close()
		defer wr.Close()

		buf := make([]byte, CopyBufferSize)

		for _, ec := range expCollections {
			folder := ec.BasePath()
			items := ec.Items(ctx)

			for item := range items {
				err := item.Error
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "getting export item").With("id", item.ID))
					return
				}

				name := item.Data.Name

				// We assume folder and name to not contain any path separators.
				// Also, this should always use `/` as this is
				// created within a zip file and not written to disk.
				// TODO(meain): Exchange paths might contain a path
				// separator and will have to have special handling.

				//nolint:forbidigo
				f, err := wr.Create(path.Join(folder, name))
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "creating zip entry").With("name", name).With("id", item.ID))
					return
				}

				_, err = io.CopyBuffer(f, item.Data.Body, buf)
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "writing zip entry").With("name", name).With("id", item.ID))
					return
				}
			}
		}
	}()

	return zipExportCol{reader}, nil
}

func (op *ExportOperation) do(
	ctx context.Context,
	opStats *exportStats,
	detailsStore streamstore.Reader,
	start time.Time,
) ([]export.Collection, error) {
	logger.Ctx(ctx).
		With("control_options", op.Options, "selectors", op.Selectors).
		Info("exporting selection")

	bup, deets, err := getBackupAndDetailsFromID(
		ctx,
		op.BackupID,
		op.store,
		detailsStore,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "getting backup and details")
	}

	observe.Message(ctx, "Exporting", observe.Bullet, clues.Hide(bup.Selector.DiscreteOwner))

	paths, err := formatDetailsForRestoration(ctx, bup.Version, op.Selectors, deets, op.ec, op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "formatting paths from details")
	}

	ctx = clues.Add(
		ctx,
		"resource_owner_id", bup.Selector.ID(),
		"resource_owner_name", clues.Hide(bup.Selector.Name()),
		"details_entries", len(deets.Entries),
		"details_paths", len(paths),
		"backup_snapshot_id", bup.SnapshotID,
		"backup_version", bup.Version)

	op.bus.Event(
		ctx,
		events.ExportStart,
		map[string]any{
			events.StartTime:        start,
			events.BackupID:         op.BackupID,
			events.BackupCreateTime: bup.CreationTime,
			events.ExportID:         opStats.exportID,
		})

	observe.Message(ctx, fmt.Sprintf("Discovered %d items in backup %s to export", len(paths), op.BackupID))

	kopiaComplete := observe.MessageWithCompletion(ctx, "Enumerating items in repository")
	defer close(kopiaComplete)

	dcs, err := op.kopia.ProduceRestoreCollections(ctx, bup.SnapshotID, paths, opStats.bytesRead, op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "producing collections to export")
	}

	kopiaComplete <- struct{}{}

	ctx = clues.Add(ctx, "coll_count", len(dcs))

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1
	opStats.cs = dcs

	expCollections, err := exportRestoreCollections(
		ctx,
		op.ec,
		bup.Version,
		op.Selectors,
		op.ExportCfg,
		op.Options,
		dcs,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "exporting collections")
	}

	opStats.ctrl = op.ec.Wait()

	logger.Ctx(ctx).Debug(opStats.ctrl)

	if op.ExportCfg.Archive {
		zc, err := zipExportCollection(ctx, expCollections)
		if err != nil {
			return nil, clues.Wrap(err, "zipping export collections")
		}

		return []export.Collection{zc}, nil
	}

	return expCollections, nil
}

// persists details and statistics about the export operation.
func (op *ExportOperation) persistResults(
	ctx context.Context,
	started time.Time,
	opStats *exportStats,
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

	// We don't have data on what all items were written
	// op.Results.ItemsWritten = opStats.ctrl.Successes

	return op.Errors.Failure()
}

// ---------------------------------------------------------------------------
// Exportr funcs
// ---------------------------------------------------------------------------

func exportRestoreCollections(
	ctx context.Context,
	ec inject.ExportConsumer,
	backupVersion int,
	sel selectors.Selector,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
) ([]export.Collection, error) {
	complete := observe.MessageWithCompletion(ctx, "Preparing export")
	defer func() {
		complete <- struct{}{}
		close(complete)
	}()

	expCollections, err := ec.ExportRestoreCollections(
		ctx,
		backupVersion,
		sel,
		exportCfg,
		opts,
		dcs,
		errs)
	if err != nil {
		return nil, clues.Wrap(err, "exporting collections")
	}

	return expCollections, nil
}
