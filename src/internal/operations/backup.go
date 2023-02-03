package operations

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
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

	ResourceOwner string             `json:"resourceOwner"`
	Results       BackupResults      `json:"results"`
	Selectors     selectors.Selector `json:"selectors"`
	Version       string             `json:"version"`

	account account.Account

	// when true, this allows for incremental backups instead of full data pulls
	incremental bool
}

// BackupResults aggregate the details of the result of the operation.
type BackupResults struct {
	stats.Errs // deprecated in place of fault.Errors in the base operation.
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
		operation:     newOperation(opts, bus, kw, sw),
		ResourceOwner: selector.DiscreteOwner,
		Selectors:     selector,
		Version:       "v0",
		account:       acct,
		incremental:   useIncrementalBackup(selector, opts),
	}
	if err := op.validate(); err != nil {
		return BackupOperation{}, err
	}

	return op, nil
}

func (op BackupOperation) validate() error {
	if len(op.ResourceOwner) == 0 {
		return errors.New("backup requires a resource owner")
	}

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
	readErr, writeErr error
}

type detailsWriter interface {
	WriteBackupDetails(context.Context, *details.Details) (string, error)
}

// ---------------------------------------------------------------------------
// Primary Controller
// ---------------------------------------------------------------------------

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = clues.Wrap(r.(error), "panic recovery").WithClues(ctx).With("stacktrace", debug.Stack())
		}
	}()

	ctx, end := D.Span(ctx, "operations:backup:run")
	defer func() {
		end()
		// wait for the progress display to clean up
		observe.Complete()
	}()

	ctx = clues.AddAll(
		ctx,
		"tenant_id", op.account.ID(), // TODO: pii
		"resource_owner", op.ResourceOwner, // TODO: pii
		"backup_id", op.Results.BackupID,
		"service", op.Selectors.Service,
		"incremental", op.incremental)

	if err := op.do(ctx); err != nil {
		logger.Ctx(ctx).
			With("err", err).
			Errorw("backup operation", clues.InErr(err).Slice()...)

		return err
	}

	logger.Ctx(ctx).Infow("completed backup", "results", op.Results)

	return nil
}

func (op *BackupOperation) do(ctx context.Context) (err error) {
	var (
		opStats       backupStats
		backupDetails *details.Builder
		toMerge       map[string]path.Path
		tenantID      = op.account.ID()
		startTime     = time.Now()
		detailsStore  = streamstore.New(op.kopia, tenantID, op.Selectors.PathService())
		reasons       = selectorToReasons(op.Selectors)
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
		err = op.persistResults(startTime, &opStats)
		if err != nil {
			op.Errors.Fail(errors.Wrap(err, "persisting backup results"))
			return
		}

		err = op.createBackupModels(
			ctx,
			detailsStore,
			opStats.k.SnapshotID,
			backupDetails.Details())
		if err != nil {
			op.Errors.Fail(errors.Wrap(err, "persisting backup"))
			opStats.writeErr = op.Errors.Err()
		}
	}()

	mans, mdColls, canUseMetaData, err := produceManifestsAndMetadata(
		ctx,
		op.kopia,
		op.store,
		reasons,
		tenantID,
		op.incremental,
		op.Errors)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "collecting manifest heuristics"))
		opStats.readErr = op.Errors.Err()
		logger.Ctx(ctx).With("err", err).Errorw("producing manifests and metadata", clues.InErr(err).Slice()...)

		return opStats.readErr
	}

	gc, err := connectToM365(ctx, op.Selectors, op.account)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "connecting to m365"))
		opStats.readErr = op.Errors.Err()
		logger.Ctx(ctx).With("err", err).Errorw("connectng to m365", clues.InErr(err).Slice()...)

		return opStats.readErr
	}

	cs, err := produceBackupDataCollections(ctx, gc, op.Selectors, mdColls, op.Options)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "retrieving data to backup"))
		opStats.readErr = op.Errors.Err()
		logger.Ctx(ctx).With("err", err).Errorw("producing backup data collections", clues.InErr(err).Slice()...)

		return opStats.readErr
	}

	ctx = clues.Add(ctx, "coll_count", len(cs))

	opStats.k, backupDetails, toMerge, err = consumeBackupDataCollections(
		ctx,
		op.kopia,
		tenantID,
		reasons,
		mans,
		cs,
		op.Results.BackupID,
		op.incremental && canUseMetaData)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "backing up service data"))
		opStats.writeErr = op.Errors.Err()
		logger.Ctx(ctx).With("err", err).Errorw("persisting collection backups", clues.InErr(err).Slice()...)

		return opStats.writeErr
	}

	if err = mergeDetails(
		ctx,
		op.store,
		detailsStore,
		mans,
		toMerge,
		backupDetails,
	); err != nil {
		op.Errors.Fail(errors.Wrap(err, "merging backup details"))
		opStats.writeErr = op.Errors.Err()
		logger.Ctx(ctx).With("err", err).Errorw("merging details", clues.InErr(err).Slice()...)

		return opStats.writeErr
	}

	opStats.gc = gc.AwaitStatus()

	// TODO(keepers): remove when fault.Errors handles all iterable error aggregation.
	if opStats.gc.ErrorCount > 0 {
		merr := multierror.Append(opStats.readErr, errors.Wrap(opStats.gc.Err, "retrieving data"))
		opStats.readErr = merr.ErrorOrNil()

		// Need to exit before we set started to true else we'll report no errors.
		return opStats.readErr
	}

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1

	return err
}

// checker to see if conditions are correct for incremental backup behavior such as
// retrieving metadata like delta tokens and previous paths.
func useIncrementalBackup(sel selectors.Selector, opts control.Options) bool {
	// Delta-based incrementals currently only supported for Exchange
	if sel.Service != selectors.ServiceExchange {
		return false
	}

	return !opts.ToggleFeatures.DisableIncrementals
}

// ---------------------------------------------------------------------------
// Producer funcs
// ---------------------------------------------------------------------------

// calls the producer to generate collections of data to backup
func produceBackupDataCollections(
	ctx context.Context,
	gc *connector.GraphConnector,
	sel selectors.Selector,
	metadata []data.Collection,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	complete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Discovering items to backup"))
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	// TODO(ashmrtn): When we're ready to wire up the global exclude list return
	// all values.
	cols, _, errs := gc.DataCollections(ctx, sel, metadata, ctrlOpts)

	return cols, errs
}

// ---------------------------------------------------------------------------
// Consumer funcs
// ---------------------------------------------------------------------------

type backuper interface {
	BackupCollections(
		ctx context.Context,
		bases []kopia.IncrementalBase,
		cs []data.Collection,
		excluded map[string]struct{},
		tags map[string]string,
		buildTreeWithBase bool,
	) (*kopia.BackupStats, *details.Builder, map[string]path.Path, error)
}

func selectorToReasons(sel selectors.Selector) []kopia.Reason {
	service := sel.PathService()
	reasons := []kopia.Reason{}

	pcs, err := sel.PathCategories()
	if err != nil {
		// This is technically safe, it's just that the resulting backup won't be
		// usable as a base for future incremental backups.
		return nil
	}

	for _, sl := range [][]path.CategoryType{pcs.Includes, pcs.Filters} {
		for _, cat := range sl {
			reasons = append(reasons, kopia.Reason{
				ResourceOwner: sel.DiscreteOwner,
				Service:       service,
				Category:      cat,
			})
		}
	}

	return reasons
}

func builderFromReason(ctx context.Context, tenant string, r kopia.Reason) (*path.Builder, error) {
	ctx = clues.Add(ctx, "category", r.Category.String())

	// This is hacky, but we want the path package to format the path the right
	// way (e.x. proper order for service, category, etc), but we don't care about
	// the folders after the prefix.
	p, err := path.Builder{}.Append("tmp").ToDataLayerPath(
		tenant,
		r.ResourceOwner,
		r.Service,
		r.Category,
		false,
	)
	if err != nil {
		return nil, clues.Wrap(err, "building path").WithClues(ctx)
	}

	return p.ToBuilder().Dir(), nil
}

// calls kopia to backup the collections of data
func consumeBackupDataCollections(
	ctx context.Context,
	bu backuper,
	tenantID string,
	reasons []kopia.Reason,
	mans []*kopia.ManifestEntry,
	cs []data.Collection,
	backupID model.StableID,
	isIncremental bool,
) (*kopia.BackupStats, *details.Builder, map[string]path.Path, error) {
	complete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Backing up data"))
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	tags := map[string]string{
		kopia.TagBackupID:       string(backupID),
		kopia.TagBackupCategory: "",
	}

	for _, reason := range reasons {
		for _, k := range reason.TagKeys() {
			tags[k] = ""
		}
	}

	bases := make([]kopia.IncrementalBase, 0, len(mans))

	for _, m := range mans {
		paths := make([]*path.Builder, 0, len(m.Reasons))
		services := map[string]struct{}{}
		categories := map[string]struct{}{}

		for _, reason := range m.Reasons {
			pb, err := builderFromReason(ctx, tenantID, reason)
			if err != nil {
				return nil, nil, nil, errors.Wrap(err, "getting subtree paths for bases")
			}

			paths = append(paths, pb)
			services[reason.Service.String()] = struct{}{}
			categories[reason.Category.String()] = struct{}{}
		}

		bases = append(bases, kopia.IncrementalBase{
			Manifest:     m.Manifest,
			SubtreePaths: paths,
		})

		svcs := make([]string, 0, len(services))
		for k := range services {
			svcs = append(svcs, k)
		}

		cats := make([]string, 0, len(categories))
		for k := range categories {
			cats = append(cats, k)
		}

		logger.Ctx(ctx).Infow(
			"using base for backup",
			"snapshot_id", m.ID,
			"services", svcs,
			"categories", cats)
	}

	kopiaStats, deets, itemsSourcedFromBase, err := bu.BackupCollections(
		ctx,
		bases,
		cs,
		nil,
		tags,
		isIncremental,
	)

	if kopiaStats.ErrorCount > 0 || kopiaStats.IgnoredErrorCount > 0 {
		if err != nil {
			err = errors.Wrapf(
				err,
				"kopia snapshot failed with %v catastrophic errors and %v ignored errors",
				kopiaStats.ErrorCount,
				kopiaStats.IgnoredErrorCount,
			)
		} else {
			err = errors.Errorf(
				"kopia snapshot failed with %v catastrophic errors and %v ignored errors",
				kopiaStats.ErrorCount,
				kopiaStats.IgnoredErrorCount,
			)
		}
	}

	return kopiaStats, deets, itemsSourcedFromBase, err
}

func matchesReason(reasons []kopia.Reason, p path.Path) bool {
	for _, reason := range reasons {
		if p.ResourceOwner() == reason.ResourceOwner &&
			p.Service() == reason.Service &&
			p.Category() == reason.Category {
			return true
		}
	}

	return false
}

func mergeDetails(
	ctx context.Context,
	ms *store.Wrapper,
	detailsStore detailsReader,
	mans []*kopia.ManifestEntry,
	shortRefsFromPrevBackup map[string]path.Path,
	deets *details.Builder,
) error {
	// Don't bother loading any of the base details if there's nothing we need to
	// merge.
	if len(shortRefsFromPrevBackup) == 0 {
		return nil
	}

	var addedEntries int

	for _, man := range mans {
		mctx := clues.Add(ctx, "manifest_id", man.ID)

		// For now skip snapshots that aren't complete. We will need to revisit this
		// when we tackle restartability.
		if len(man.IncompleteReason) > 0 {
			continue
		}

		bID, ok := man.GetTag(kopia.TagBackupID)
		if !ok {
			return clues.New("no backup ID in snapshot manifest").WithClues(mctx)
		}

		mctx = clues.Add(mctx, "manifest_backup_id", bID)

		_, baseDeets, err := getBackupAndDetailsFromID(
			ctx,
			model.StableID(bID),
			ms,
			detailsStore,
		)
		if err != nil {
			return clues.New("fetching base details for backup").WithClues(mctx)
		}

		for _, entry := range baseDeets.Items() {
			rr, err := path.FromDataLayerPath(entry.RepoRef, true)
			if err != nil {
				return clues.New("parsing base item info path").
					WithClues(mctx).
					With("repo_ref", entry.RepoRef) // todo: pii
			}

			// Although this base has an entry it may not be the most recent. Check
			// the reasons a snapshot was returned to ensure we only choose the recent
			// entries.
			//
			// TODO(ashmrtn): This logic will need expanded to cover entries from
			// checkpoints if we start doing kopia-assisted incrementals for those.
			if !matchesReason(man.Reasons, rr) {
				continue
			}

			newPath := shortRefsFromPrevBackup[rr.ShortRef()]
			if newPath == nil {
				// This entry was not sourced from a base snapshot or cached from a
				// previous backup, skip it.
				continue
			}

			// Fixup paths in the item.
			item := entry.ItemInfo
			if err := details.UpdateItem(&item, newPath); err != nil {
				return clues.New("updating item details").WithClues(mctx)
			}

			// TODO(ashmrtn): This may need updated if we start using this merge
			// strategry for items that were cached in kopia.
			itemUpdated := newPath.String() != rr.String()

			deets.Add(
				newPath.String(),
				newPath.ShortRef(),
				newPath.ToBuilder().Dir().ShortRef(),
				itemUpdated,
				item,
			)

			folders := details.FolderEntriesForPath(newPath.ToBuilder().Dir())
			deets.AddFoldersForItem(folders, item, itemUpdated)

			// Track how many entries we added so that we know if we got them all when
			// we're done.
			addedEntries++
		}
	}

	if addedEntries != len(shortRefsFromPrevBackup) {
		return clues.New("incomplete migration of backup details").
			WithClues(ctx).
			WithAll("item_count", addedEntries, "expected_item_count", len(shortRefsFromPrevBackup))
	}

	return nil
}

// writes the results metrics to the operation results.
// later stored in the manifest using createBackupModels.
func (op *BackupOperation) persistResults(
	started time.Time,
	opStats *backupStats,
) error {
	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()
	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr

	op.Status = Completed

	if opStats.readErr != nil || opStats.writeErr != nil {
		op.Status = Failed

		// TODO(keepers): replace with fault.Errors handling.
		return multierror.Append(
			errors.New("errors prevented the operation from processing"),
			opStats.readErr,
			opStats.writeErr)
	}

	op.Results.BytesRead = opStats.k.TotalHashedBytes
	op.Results.BytesUploaded = opStats.k.TotalUploadedBytes
	op.Results.ItemsWritten = opStats.k.TotalFileCount
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.gc == nil {
		op.Status = Failed
		return errors.New("data population never completed")
	}

	if opStats.readErr == nil && opStats.writeErr == nil && opStats.gc.Successful == 0 {
		op.Status = NoData
	}

	op.Results.ItemsRead = opStats.gc.Successful

	return nil
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	detailsStore detailsWriter,
	snapID string,
	backupDetails *details.Details,
) error {
	ctx = clues.Add(ctx, "snapshot_id", snapID)

	if backupDetails == nil {
		return clues.New("no backup details to record").WithClues(ctx)
	}

	detailsID, err := detailsStore.WriteBackupDetails(ctx, backupDetails)
	if err != nil {
		return clues.Wrap(err, "creating backupDetails model").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "details_id", detailsID)
	b := backup.New(
		snapID, detailsID, op.Status.String(),
		op.Results.BackupID,
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
		op.Errors,
	)

	if err = op.store.Put(ctx, model.BackupSchema, b); err != nil {
		return clues.Wrap(err, "creating backup model").WithClues(ctx)
	}

	dur := op.Results.CompletedAt.Sub(op.Results.StartedAt)

	op.bus.Event(
		ctx,
		events.BackupEnd,
		map[string]any{
			events.BackupID:   b.ID,
			events.DataStored: op.Results.BytesUploaded,
			events.Duration:   dur,
			events.EndTime:    common.FormatTime(op.Results.CompletedAt),
			events.Resources:  op.Results.ResourceOwners,
			events.Service:    op.Selectors.PathService().String(),
			events.StartTime:  common.FormatTime(op.Results.StartedAt),
			events.Status:     op.Status.String(),
		},
	)

	return nil
}
