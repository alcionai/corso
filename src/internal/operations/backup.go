package operations

import (
	"context"
	"time"

	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/snapshot"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	Results   BackupResults      `json:"results"`
	Selectors selectors.Selector `json:"selectors"`
	Version   string             `json:"version"`

	account account.Account
}

// BackupResults aggregate the details of the result of the operation.
type BackupResults struct {
	stats.Errs
	stats.ReadWrites
	stats.StartAndEndTime
	BackupID model.StableID `json:"backupID"`
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	selector selectors.Selector,
	bus events.Eventer,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, bus, kw, sw),
		Selectors: selector,
		Version:   "v0",
		account:   acct,
	}
	if err := op.validate(); err != nil {
		return BackupOperation{}, err
	}

	return op, nil
}

func (op BackupOperation) validate() error {
	return op.operation.validate()
}

// aggregates stats from the backup.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type backupStats struct {
	k                 *kopia.BackupStats
	gc                *support.ConnectorOperationStatus
	resourceCount     int
	started           bool
	readErr, writeErr error
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	ctx, end := D.Span(ctx, "operations:backup:run")
	defer end()

	var (
		opStats       backupStats
		backupDetails *details.Details
		startTime     = time.Now()
	)

	op.Results.BackupID = model.StableID(uuid.NewString())

	op.bus.Event(
		ctx,
		events.BackupStart,
		map[string]any{
			events.StartTime: startTime,
			events.Service:   op.Selectors.Service.String(),
			events.BackupID:  op.Results.BackupID,
		},
	)

	// persist operation results to the model store on exit
	defer func() {
		// wait for the progress display to clean up
		observe.Complete()

		err = op.persistResults(startTime, &opStats)
		if err != nil {
			return
		}

		err = op.createBackupModels(ctx, opStats.k.SnapshotID, backupDetails)
		if err != nil {
			opStats.writeErr = err
		}
	}()

	_, mdColls, err := produceManifestsAndMetadata(ctx, op.kopia, op.store, op.Selectors, op.account)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "connecting to M365")
		return opStats.readErr
	}

	gc, err := connectToM365(ctx, op.Selectors, op.account)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "connecting to M365")
		return opStats.readErr
	}

	cs, err := produceBackupDataCollections(ctx, gc, op.Selectors, mdColls)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "retrieving data to backup")
		return opStats.readErr
	}

	opStats.k, backupDetails, err = consumeBackupDataCollections(ctx, op.kopia, op.Selectors, cs, op.Results.BackupID)
	if err != nil {
		opStats.writeErr = errors.Wrap(err, "backing up service data")
		return opStats.writeErr
	}

	logger.Ctx(ctx).Debugf(
		"Backed up %d directories and %d files",
		opStats.k.TotalDirectoryCount, opStats.k.TotalFileCount,
	)

	// TODO: should always be 1, since backups are 1:1 with resourceOwners now.
	opStats.resourceCount = len(data.ResourceOwnerSet(cs))
	opStats.started = true
	opStats.gc = gc.AwaitStatus()

	return err
}

// calls kopia to retrieve prior backup manifests, metadata collections to supply backup heuristics.
func produceManifestsAndMetadata(
	ctx context.Context,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	sel selectors.Selector,
	acct account.Account,
) ([]*snapshot.Manifest, []data.Collection, error) {
	complete, closer := observe.MessageWithCompletion("Fetching backup heuristics:")
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	m365, err := acct.M365Config()
	if err != nil {
		return nil, nil, err
	}

	var (
		tid         = m365.AzureTenantID
		oc          = selectorToOwnersCats(sel)
		ms          = kopia.FetchPrevSnapshotManifests(ctx, kw.Conn(), &oc)
		collections []data.Collection
	)

	for _, man := range ms {
		bup := backup.Backup{}
		if err := sw.Get(
			ctx,
			model.BackupSchema,
			model.StableID(man.Tags[kopia.BackupIDTag]),
			&bup,
		); err != nil {
			return nil, nil, err
		}

		for _, fn := range []string{graph.DeltaTokenFileName} {
			for ro := range oc.ResourceOwners {
				for _, sc := range oc.ServiceCats {
					colls, err := collectMetadata(ctx, kw, fn, tid, ro, bup.SnapshotID, sc)
					if err != nil {
						return nil, nil, err
					}

					collections = append(collections, colls...)
				}
			}
		}
	}

	return ms, collections, err
}

func collectMetadata(
	ctx context.Context,
	kw *kopia.Wrapper,
	fileName, tenantID, resourceOwner, snapshotID string,
	sc kopia.ServiceCat,
) ([]data.Collection, error) {
	paths := []path.Path{}
	pathsByRef := map[string][]string{}

	p, err := path.Builder{}.
		Append(fileName).
		ToServiceCategoryMetadataPath(
			tenantID,
			resourceOwner,
			sc.Service,
			sc.Category,
			true)
	if err != nil {
		return nil, errors.Wrap(err, "building metadata path")
	}

	dir, err := p.Dir()
	if err != nil {
		return nil, errors.Wrap(err, "malformed metadata path: "+p.String())
	}

	paths = append(paths, p)
	pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fileName)

	dcs, err := kw.RestoreMultipleItems(ctx, snapshotID, paths, nil)
	if err != nil {
		return nil, errors.Wrap(err, "collecting prior metadata")
	}

	return dcs, nil
}

func selectorToOwnersCats(sel selectors.Selector) kopia.OwnersCats {
	service := sel.PathService()
	oc := kopia.OwnersCats{
		ResourceOwners: map[string]struct{}{},
		ServiceCats:    map[string]kopia.ServiceCat{},
	}

	ros, err := sel.ResourceOwners()
	if err != nil {
		return oc
	}

	for _, sl := range [][]string{ros.Includes, ros.Filters} {
		for _, ro := range sl {
			oc.ResourceOwners[ro] = struct{}{}
		}
	}

	pcs, err := sel.PathCategories()
	if err != nil {
		return oc
	}

	for _, sl := range [][]path.CategoryType{pcs.Includes, pcs.Filters} {
		for _, cat := range sl {
			oc.ServiceCats[kopia.MakeServiceCat(service, cat)] = kopia.ServiceCat{
				Service:  service,
				Category: cat,
			}
		}
	}

	return oc
}

// calls the producer to generate collections of data to backup
func produceBackupDataCollections(
	ctx context.Context,
	gc *connector.GraphConnector,
	sel selectors.Selector,
	metadata []data.Collection,
) ([]data.Collection, error) {
	complete, closer := observe.MessageWithCompletion("Discovering items to backup:")
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	return gc.DataCollections(ctx, sel, metadata)
}

// calls kopia to backup the collections of data
func consumeBackupDataCollections(
	ctx context.Context,
	kw *kopia.Wrapper,
	sel selectors.Selector,
	cs []data.Collection,
	backupID model.StableID,
) (*kopia.BackupStats, *details.Details, error) {
	complete, closer := observe.MessageWithCompletion("Backing up data:")
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	return kw.BackupCollections(ctx, nil, cs, sel.PathService(), string(backupID))
}

// writes the results metrics to the operation results.
// later stored in the manifest using createBackupModels.
func (op *BackupOperation) persistResults(
	started time.Time,
	opStats *backupStats,
) error {
	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	op.Status = Completed
	if !opStats.started {
		op.Status = Failed

		return multierror.Append(
			errors.New("errors prevented the operation from processing"),
			opStats.readErr,
			opStats.writeErr)
	}

	if opStats.readErr == nil && opStats.writeErr == nil && opStats.gc.Successful == 0 {
		op.Status = NoData
	}

	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr

	op.Results.BytesRead = opStats.k.TotalHashedBytes
	op.Results.BytesUploaded = opStats.k.TotalUploadedBytes
	op.Results.ItemsRead = opStats.gc.Successful
	op.Results.ItemsWritten = opStats.k.TotalFileCount
	op.Results.ResourceOwners = opStats.resourceCount

	return nil
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	snapID string,
	backupDetails *details.Details,
) error {
	if backupDetails == nil {
		return errors.New("no backup details to record")
	}

	detailsID, err := streamstore.New(
		op.kopia,
		op.account.ID(),
		op.Selectors.PathService(),
	).WriteBackupDetails(ctx, backupDetails)
	if err != nil {
		return errors.Wrap(err, "creating backupdetails model")
	}

	b := backup.New(
		snapID, detailsID, op.Status.String(),
		op.Results.BackupID,
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
	)

	err = op.store.Put(ctx, model.BackupSchema, b)
	if err != nil {
		return errors.Wrap(err, "creating backup model")
	}

	op.bus.Event(
		ctx,
		events.BackupEnd,
		map[string]any{
			events.BackupID:   b.ID,
			events.DataStored: op.Results.BytesUploaded,
			events.Duration:   op.Results.CompletedAt.Sub(op.Results.StartedAt),
			events.EndTime:    op.Results.CompletedAt,
			events.Resources:  op.Results.ResourceOwners,
			events.Service:    op.Selectors.PathService().String(),
			events.StartTime:  op.Results.StartedAt,
			events.Status:     op.Status,
		},
	)

	return nil
}
