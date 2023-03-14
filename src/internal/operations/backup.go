package operations

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/crash"
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
	"github.com/alcionai/corso/src/pkg/fault"
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
	k             *kopia.BackupStats
	gc            *support.ConnectorOperationStatus
	resourceCount int
}

// ---------------------------------------------------------------------------
// Primary Controller
// ---------------------------------------------------------------------------

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover()); crErr != nil {
			err = crErr
		}
	}()

	ctx, end := D.Span(ctx, "operations:backup:run")
	defer func() {
		end()
		// wait for the progress display to clean up
		observe.Complete()
	}()

	// -----
	// Setup
	// -----

	var (
		opStats      backupStats
		startTime    = time.Now()
		detailsStore = streamstore.NewDetails(op.kopia, op.account.ID(), op.Selectors.PathService())
		errorsStore  = streamstore.NewFaultErrors(op.kopia, op.account.ID(), op.Selectors.PathService())
	)

	op.Results.BackupID = model.StableID(uuid.NewString())

	ctx = clues.Add(
		ctx,
		"tenant_id", op.account.ID(), // TODO: pii
		"resource_owner", op.ResourceOwner, // TODO: pii
		"backup_id", op.Results.BackupID,
		"service", op.Selectors.Service,
		"incremental", op.incremental)

	op.bus.Event(
		ctx,
		events.BackupStart,
		map[string]any{
			events.StartTime: startTime,
			events.Service:   op.Selectors.Service.String(),
			events.BackupID:  op.Results.BackupID,
		})

	// -----
	// Execution
	// -----

	deets, err := op.do(
		ctx,
		&opStats,
		detailsStore,
		op.Results.BackupID)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.Ctx(ctx).
			With("err", err).
			Errorw("doing backup", clues.InErr(err).Slice()...)
		op.Errors.Fail(errors.Wrap(err, "doing backup"))
	}

	// TODO: the consumer (sdk or cli) should run this, not operations.
	recoverableCount := len(op.Errors.Recovered())
	for i, err := range op.Errors.Recovered() {
		logger.Ctx(ctx).
			With("error", err).
			With(clues.InErr(err).Slice()...).
			Errorf("doing backup: recoverable error %d of %d", i+1, recoverableCount)
	}

	skippedCount := len(op.Errors.Skipped())
	for i, skip := range op.Errors.Skipped() {
		logger.Ctx(ctx).With("skip", skip).Infof("doing backup: skipped item %d of %d", i+1, skippedCount)
	}

	// -----
	// Persistence
	// -----

	err = op.persistResults(startTime, &opStats)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "persisting backup results"))
		return op.Errors.Failure()
	}

	// force exit without backup in certain cases.
	// see: https://github.com/alcionai/corso/pull/2510#discussion_r1113532530
	for _, e := range op.Errors.Recovered() {
		if clues.HasLabel(e, fault.LabelForceNoBackupCreation) {
			logger.Ctx(ctx).
				With("error", e).
				With(clues.InErr(err).Slice()...).
				Infow("completed backup; conditional error forcing exit without model persistence",
					"results", op.Results)

			return op.Errors.Fail(errors.Wrap(e, "forced backup")).Failure()
		}
	}

	err = op.createBackupModels(
		ctx,
		detailsStore,
		errorsStore,
		opStats.k.SnapshotID,
		op.Results.BackupID,
		deets.Details())
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "persisting backup"))
		return op.Errors.Failure()
	}

	logger.Ctx(ctx).Infow("completed backup", "results", op.Results)

	return nil
}

// do is purely the action of running a backup.  All pre/post behavior
// is found in Run().
func (op *BackupOperation) do(
	ctx context.Context,
	opStats *backupStats,
	detailsStore streamstore.Streamer,
	backupID model.StableID,
) (*details.Builder, error) {
	reasons := selectorToReasons(op.Selectors)
	logger.Ctx(ctx).With("selectors", op.Selectors).Info("backing up selection")

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1

	mans, mdColls, canUseMetaData, err := produceManifestsAndMetadata(
		ctx,
		op.kopia,
		op.store,
		reasons,
		op.account.ID(),
		op.incremental,
		op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "producing manifests and metadata")
	}

	gc, err := connectToM365(ctx, op.Selectors, op.account, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "connectng to m365")
	}

	cs, excludes, err := produceBackupDataCollections(ctx, gc, op.Selectors, mdColls, op.Options, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "producing backup data collections")
	}

	ctx = clues.Add(ctx, "coll_count", len(cs))

	writeStats, deets, toMerge, err := consumeBackupDataCollections(
		ctx,
		op.kopia,
		op.account.ID(),
		reasons,
		mans,
		cs,
		excludes,
		backupID,
		op.incremental && canUseMetaData,
		op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "persisting collection backups")
	}

	opStats.k = writeStats

	err = mergeDetails(
		ctx,
		op.store,
		detailsStore,
		mans,
		toMerge,
		deets,
		op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "merging details")
	}

	opStats.gc = gc.AwaitStatus()

	logger.Ctx(ctx).Debug(gc.PrintableStatus())

	return deets, nil
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
	metadata []data.RestoreCollection,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	complete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Discovering items to backup"))
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	return gc.DataCollections(ctx, sel, metadata, ctrlOpts, errs)
}

// ---------------------------------------------------------------------------
// Consumer funcs
// ---------------------------------------------------------------------------

type backuper interface {
	BackupCollections(
		ctx context.Context,
		bases []kopia.IncrementalBase,
		cs []data.BackupCollection,
		excluded map[string]map[string]struct{},
		tags map[string]string,
		buildTreeWithBase bool,
		errs *fault.Bus,
	) (*kopia.BackupStats, *details.Builder, map[string]kopia.PrevRefs, error)
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
	p, err := path.Build(
		tenant,
		r.ResourceOwner,
		r.Service,
		r.Category,
		false,
		"tmp")
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
	cs []data.BackupCollection,
	excludes map[string]map[string]struct{},
	backupID model.StableID,
	isIncremental bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, map[string]kopia.PrevRefs, error) {
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
		// TODO(ashmrtn): When we're ready to enable incremental backups for
		// OneDrive replace this with `excludes`.
		nil,
		tags,
		isIncremental,
		errs)
	if err != nil {
		if kopiaStats == nil {
			return nil, nil, nil, err
		}

		return nil, nil, nil, clues.Stack(err).With(
			"kopia_errors", kopiaStats.ErrorCount,
			"kopia_ignored_errors", kopiaStats.IgnoredErrorCount)
	}

	if kopiaStats.ErrorCount > 0 ||
		(kopiaStats.IgnoredErrorCount > kopiaStats.ExpectedIgnoredErrorCount) {
		err = clues.New("building kopia snapshot").With(
			"kopia_errors", kopiaStats.ErrorCount,
			"kopia_ignored_errors", kopiaStats.IgnoredErrorCount)
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
	detailsStore streamstore.Streamer,
	mans []*kopia.ManifestEntry,
	shortRefsFromPrevBackup map[string]kopia.PrevRefs,
	deets *details.Builder,
	errs *fault.Bus,
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
			errs)
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

			prev, ok := shortRefsFromPrevBackup[rr.ShortRef()]
			if !ok {
				// This entry was not sourced from a base snapshot or cached from a
				// previous backup, skip it.
				continue
			}

			newPath := prev.Repo
			newLoc := prev.Location

			// Fixup paths in the item.
			item := entry.ItemInfo
			if err := details.UpdateItem(&item, newPath); err != nil {
				return clues.New("updating item details").WithClues(mctx)
			}

			// TODO(ashmrtn): This may need updated if we start using this merge
			// strategry for items that were cached in kopia.
			var (
				itemUpdated = newPath.String() != rr.String()
				newLocStr   string
				locBuilder  *path.Builder
			)

			if newLoc != nil {
				locBuilder = newLoc.ToBuilder()
				newLocStr = newLoc.Folder(true)
				itemUpdated = itemUpdated || newLocStr != entry.LocationRef
			}

			deets.Add(
				newPath.String(),
				newPath.ShortRef(),
				newPath.ToBuilder().Dir().ShortRef(),
				newLocStr,
				itemUpdated,
				item)

			folders := details.FolderEntriesForPath(newPath.ToBuilder().Dir(), locBuilder)
			deets.AddFoldersForItem(folders, item, itemUpdated)

			// Track how many entries we added so that we know if we got them all when
			// we're done.
			addedEntries++
		}
	}

	if addedEntries != len(shortRefsFromPrevBackup) {
		return clues.New("incomplete migration of backup details").
			WithClues(ctx).
			With("item_count", addedEntries, "expected_item_count", len(shortRefsFromPrevBackup))
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

	op.Status = Completed

	if op.Errors.Failure() != nil {
		op.Status = Failed
	}

	if opStats.k == nil {
		op.Status = Failed
		return errors.New("backup persistence never completed")
	}

	op.Results.BytesRead = opStats.k.TotalHashedBytes
	op.Results.BytesUploaded = opStats.k.TotalUploadedBytes
	op.Results.ItemsWritten = opStats.k.TotalFileCount
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.gc == nil {
		op.Status = Failed
		return errors.New("backup population never completed")
	}

	if op.Status != Failed && opStats.gc.Metrics.Successes == 0 {
		op.Status = NoData
	}

	op.Results.ItemsRead = opStats.gc.Metrics.Successes

	return op.Errors.Failure()
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	detailsStore, errorsStore streamstore.Writer,
	snapID string,
	backupID model.StableID,
	backupDetails *details.Details,
) error {
	ctx = clues.Add(ctx, "snapshot_id", snapID)
	errs := op.Errors

	if backupDetails == nil {
		return clues.New("no backup details to record").WithClues(ctx)
	}

	detailsID, err := detailsStore.Write(ctx, backupDetails, errs)
	if err != nil {
		return clues.Wrap(err, "creating backupDetails persistence").WithClues(ctx)
	}

	errorsID, err := errorsStore.Write(ctx, errs.Errors(), errs)
	if err != nil {
		return clues.Wrap(err, "creating errors persistence").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "details_id", detailsID)
	b := backup.New(
		snapID, detailsID, errorsID,
		op.Status.String(),
		backupID,
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
		errs)

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
		})

	return nil
}
