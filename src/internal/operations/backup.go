package operations

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	ResourceOwner idname.Provider

	Results   BackupResults      `json:"results"`
	Selectors selectors.Selector `json:"selectors"`
	Version   string             `json:"version"`

	// BackupVersion ONLY controls the value that gets persisted to the
	// backup model after operation.  It does NOT modify the operation behavior
	// to match the version.  Its inclusion here is, unfortunately, purely to
	// facilitate integration testing that requires a certain backup version, and
	// should be removed when we have a more controlled workaround.
	BackupVersion int

	account account.Account
	bp      inject.BackupProducer

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
	bp inject.BackupProducer,
	acct account.Account,
	selector selectors.Selector,
	owner idname.Provider,
	bus events.Eventer,
) (BackupOperation, error) {
	op := BackupOperation{
		operation:     newOperation(opts, bus, count.New(), kw, sw),
		ResourceOwner: owner,
		Selectors:     selector,
		Version:       "v0",
		BackupVersion: version.Backup,
		account:       acct,
		incremental:   useIncrementalBackup(selector, opts),
		bp:            bp,
	}

	if err := op.validate(); err != nil {
		return BackupOperation{}, err
	}

	return op, nil
}

func (op BackupOperation) validate() error {
	if op.ResourceOwner == nil {
		return clues.New("backup requires a resource owner")
	}

	if len(op.ResourceOwner.ID()) == 0 {
		return clues.New("backup requires a resource owner with a populated ID")
	}

	if op.bp == nil {
		return clues.New("missing backup producer")
	}

	return op.operation.validate()
}

// aggregates stats from the backup.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type backupStats struct {
	k                   *kopia.BackupStats
	ctrl                *data.CollectionStats
	resourceCount       int
	hasNewDetailEntries bool
}

// To qualify as an assist backup, all of the following must be true:
// 1. new detail entries were produced during this operation
// 2. The updated details file was persisted without error
// 3. we have a valid snapshot ID
// 4. we don't have any non-recoverable errors
// 5. we recorded recoverable errors
// 6. We are not running in best effort mode. Reason being that there is
// no way to distinguish assist backups from merge backups in best effort mode
// Primary reason for persisting assist backup models is to ensure we don't
// lose corso extension data(deets) for items which were downloaded and
// processed by kopia during this backup operation.
// Note: kopia.DetailsMergeInfoer doesn't impact decision making for creating
// assist backups. It may be empty if it’s the very first backup so there is no
// merge base to source base details from, or non-empty, if there was a merge base.
// In summary, if there are no new deets, no new extension data was produced
// and hence no need to persist assist backup model.
func tagAsAssistBackup(
	newDeetsProduced bool,
	snapID, ssid string,
	failurePolicy control.FailurePolicy,
	err *fault.Bus,
) bool {
	return newDeetsProduced &&
		snapID != "" &&
		ssid != "" &&
		err.Failure() == nil &&
		len(err.Recovered()) > 0 &&
		failurePolicy != control.BestEffort
}

// ---------------------------------------------------------------------------
// Primary Controller
// ---------------------------------------------------------------------------

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "backup"); crErr != nil {
			err = crErr
		}
	}()

	ctx, end := diagnostics.Span(ctx, "operations:backup:run")
	defer func() {
		end()
	}()

	ctx, flushMetrics := events.NewMetrics(ctx, logger.Writer{Ctx: ctx})
	defer flushMetrics()

	var runnable bool

	// IsBackupRunnable checks if the user has services enabled to run a backup.
	// it also checks for conditions like mailbox full.
	runnable, err = op.bp.IsBackupRunnable(ctx, op.Selectors.PathService(), op.ResourceOwner.ID())
	if err != nil {
		logger.CtxErr(ctx, err).Error("verifying backup is runnable")
		op.Errors.Fail(clues.Wrap(err, "verifying backup is runnable"))

		return
	}

	if !runnable {
		logger.CtxErr(ctx, graph.ErrServiceNotEnabled).Error("checking if backup is enabled")
		op.Errors.Fail(clues.Wrap(err, "checking if backup is enabled"))

		return
	}

	// -----
	// Setup
	// -----

	var (
		opStats   backupStats
		startTime = time.Now()
		sstore    = streamstore.NewStreamer(op.kopia, op.account.ID(), op.Selectors.PathService())
	)

	op.Results.BackupID = model.StableID(uuid.NewString())

	ctx = clues.Add(
		ctx,
		"tenant_id", clues.Hide(op.account.ID()),
		"resource_owner_id", op.ResourceOwner.ID(),
		"resource_owner_name", clues.Hide(op.ResourceOwner.Name()),
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

	defer func() {
		op.bus.Event(
			ctx,
			events.BackupEnd,
			map[string]any{
				events.BackupID:   op.Results.BackupID,
				events.DataStored: op.Results.BytesUploaded,
				events.Duration:   op.Results.CompletedAt.Sub(op.Results.StartedAt),
				events.EndTime:    dttm.Format(op.Results.CompletedAt),
				events.Resources:  op.Results.ResourceOwners,
				events.Service:    op.Selectors.PathService().String(),
				events.StartTime:  dttm.Format(op.Results.StartedAt),
				events.Status:     op.Status.String(),
			})
	}()

	// -----
	// Execution
	// -----

	observe.Message(ctx, "Backing Up", observe.Bullet, clues.Hide(op.ResourceOwner.Name()))

	deets, err := op.do(
		ctx,
		&opStats,
		sstore,
		op.Results.BackupID)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.CtxErr(ctx, err).Error("running backup")
		op.Errors.Fail(clues.Wrap(err, "running backup"))
	}

	// finalizeErrorHandling(ctx, op.Options, op.Errors, "running backup")
	LogFaultErrors(ctx, op.Errors.Errors(), "running backup")

	// -----
	// Persistence
	// -----

	err = op.persistResults(startTime, &opStats)
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting backup results"))
		return op.Errors.Failure()
	}

	err = op.createBackupModels(
		ctx,
		sstore,
		opStats,
		op.Results.BackupID,
		op.BackupVersion,
		deets.Details())
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting backup models"))
		return op.Errors.Failure()
	}

	if op.Errors.Failure() == nil {
		logger.Ctx(ctx).Infow("completed backup", "results", op.Results)
	}

	return op.Errors.Failure()
}

// do is purely the action of running a backup.  All pre/post behavior
// is found in Run().
func (op *BackupOperation) do(
	ctx context.Context,
	opStats *backupStats,
	detailsStore streamstore.Streamer,
	backupID model.StableID,
) (*details.Builder, error) {
	var (
		reasons           = selectorToReasons(op.account.ID(), op.Selectors, false)
		fallbackReasons   = makeFallbackReasons(op.account.ID(), op.Selectors)
		lastBackupVersion = version.NoBackup
	)

	logger.Ctx(ctx).With(
		"control_options", op.Options,
		"selectors", op.Selectors).
		Info("backing up selection")

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1

	kbf, err := op.kopia.NewBaseFinder(op.store)
	if err != nil {
		return nil, clues.Stack(err)
	}

	mans, mdColls, canUseMetadata, err := produceManifestsAndMetadata(
		ctx,
		kbf,
		op.kopia,
		reasons, fallbackReasons,
		op.account.ID(),
		op.incremental)
	if err != nil {
		return nil, clues.Wrap(err, "producing manifests and metadata")
	}

	ctx = clues.Add(ctx, "can_use_metadata", canUseMetadata)

	if canUseMetadata {
		lastBackupVersion = mans.MinBackupVersion()
	}

	cs, ssmb, canUsePreviousBackup, err := produceBackupDataCollections(
		ctx,
		op.bp,
		op.ResourceOwner,
		op.Selectors,
		mdColls,
		lastBackupVersion,
		op.Options,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "producing backup data collections")
	}

	ctx = clues.Add(
		ctx,
		"can_use_previous_backup", canUsePreviousBackup,
		"collection_count", len(cs))

	writeStats, deets, toMerge, err := consumeBackupCollections(
		ctx,
		op.kopia,
		op.account.ID(),
		reasons,
		mans,
		cs,
		ssmb,
		backupID,
		op.incremental && canUseMetadata && canUsePreviousBackup,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "persisting collection backups")
	}

	opStats.hasNewDetailEntries = deets != nil && !deets.Empty()
	opStats.k = writeStats

	err = mergeDetails(
		ctx,
		detailsStore,
		mans.Backups(),
		toMerge,
		deets,
		writeStats,
		op.Selectors.PathService(),
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "merging details")
	}

	opStats.ctrl = op.bp.Wait()

	logger.Ctx(ctx).Debug(opStats.ctrl)

	return deets, nil
}

func makeFallbackReasons(tenant string, sel selectors.Selector) []kopia.Reasoner {
	if sel.PathService() != path.SharePointService &&
		sel.DiscreteOwner != sel.DiscreteOwnerName {
		return selectorToReasons(tenant, sel, true)
	}

	return nil
}

// checker to see if conditions are correct for incremental backup behavior such as
// retrieving metadata like delta tokens and previous paths.
func useIncrementalBackup(sel selectors.Selector, opts control.Options) bool {
	return !opts.ToggleFeatures.DisableIncrementals
}

// ---------------------------------------------------------------------------
// Producer funcs
// ---------------------------------------------------------------------------

// calls the producer to generate collections of data to backup
func produceBackupDataCollections(
	ctx context.Context,
	bp inject.BackupProducer,
	protectedResource idname.Provider,
	sel selectors.Selector,
	metadata []data.RestoreCollection,
	lastBackupVersion int,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	progressBar := observe.MessageWithCompletion(ctx, "Discovering items to backup")
	defer close(progressBar)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion:   lastBackupVersion,
		MetadataCollections: metadata,
		Options:             ctrlOpts,
		ProtectedResource:   protectedResource,
		Selector:            sel,
	}

	return bp.ProduceBackupCollections(ctx, bpc, errs)
}

// ---------------------------------------------------------------------------
// Consumer funcs
// ---------------------------------------------------------------------------

func selectorToReasons(
	tenant string,
	sel selectors.Selector,
	useOwnerNameForID bool,
) []kopia.Reasoner {
	service := sel.PathService()
	reasons := []kopia.Reasoner{}

	pcs, err := sel.PathCategories()
	if err != nil {
		// This is technically safe, it's just that the resulting backup won't be
		// usable as a base for future incremental backups.
		return nil
	}

	owner := sel.DiscreteOwner
	if useOwnerNameForID {
		owner = sel.DiscreteOwnerName
	}

	for _, sl := range [][]path.CategoryType{pcs.Includes, pcs.Filters} {
		for _, cat := range sl {
			reasons = append(reasons, kopia.NewReason(tenant, owner, service, cat))
		}
	}

	return reasons
}

// calls kopia to backup the collections of data
func consumeBackupCollections(
	ctx context.Context,
	bc kinject.BackupConsumer,
	tenantID string,
	reasons []kopia.Reasoner,
	bbs kopia.BackupBases,
	cs []data.BackupCollection,
	pmr prefixmatcher.StringSetReader,
	backupID model.StableID,
	isIncremental bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error) {
	ctx = clues.Add(ctx, "collection_source", "operations")

	progressBar := observe.MessageWithCompletion(ctx, "Backing up data")
	defer close(progressBar)

	tags := map[string]string{
		kopia.TagBackupID:       string(backupID),
		kopia.TagBackupCategory: "",
	}

	kopiaStats, deets, itemsSourcedFromBase, err := bc.ConsumeBackupCollections(
		ctx,
		reasons,
		bbs,
		cs,
		pmr,
		tags,
		isIncremental,
		errs)
	if err != nil {
		if kopiaStats == nil {
			return nil, nil, nil, clues.Stack(err)
		}

		return nil, nil, nil, clues.Stack(err).With(
			"kopia_errors", kopiaStats.ErrorCount,
			"kopia_ignored_errors", kopiaStats.IgnoredErrorCount)
	}

	ctx = clues.Add(
		ctx,
		"kopia_errors", kopiaStats.ErrorCount,
		"kopia_ignored_errors", kopiaStats.IgnoredErrorCount,
		"kopia_expected_ignored_errors", kopiaStats.ExpectedIgnoredErrorCount)

	if kopiaStats.ErrorCount > 0 {
		err = clues.New("building kopia snapshot").WithClues(ctx)
	} else if kopiaStats.IgnoredErrorCount > kopiaStats.ExpectedIgnoredErrorCount {
		logger.Ctx(ctx).Info("recoverable errors were seen during backup")
	}

	return kopiaStats, deets, itemsSourcedFromBase, err
}

func matchesReason(reasons []kopia.Reasoner, p path.Path) bool {
	for _, reason := range reasons {
		if p.ResourceOwner() == reason.ProtectedResource() &&
			p.Service() == reason.Service() &&
			p.Category() == reason.Category() {
			return true
		}
	}

	return false
}

// getNewPathRefs returns
//  1. the new RepoRef for the item if it needs merged
//  2. the new locationPath
//  3. if the location was likely updated
//  4. any errors encountered
func getNewPathRefs(
	dataFromBackup kopia.DetailsMergeInfoer,
	entry *details.Entry,
	repoRef path.Path,
	backupVersion int,
) (path.Path, *path.Builder, bool, error) {
	// Right now we can't guarantee that we have an old location in the
	// previous details entry so first try a lookup without a location to see
	// if it matches so we don't need to try parsing from the old entry.
	//
	// TODO(ashmrtn): In the future we can remove this first check as we'll be
	// able to assume we always have the location in the previous entry. We'll end
	// up doing some extra parsing, but it will simplify this code.
	if repoRef.Service() == path.ExchangeService {
		newPath, newLoc, err := dataFromBackup.GetNewPathRefs(
			repoRef.ToBuilder(),
			entry.Modified(),
			nil)
		if err != nil {
			return nil, nil, false, clues.Wrap(err, "getting new paths")
		} else if newPath == nil {
			// This entry doesn't need merging.
			return nil, nil, false, nil
		} else if newLoc == nil {
			return nil, nil, false, clues.New("unable to find new exchange location")
		}

		// This is kind of jank cause we're in a transitionary period, but even if
		// we're consesrvative here about marking something as updated the RepoRef
		// comparison in the caller should catch the change. Calendars is the only
		// exception, since it uses IDs for folders, but we should already be
		// populating the LocationRef for them.
		//
		// Without this, all OneDrive items will be marked as updated the first time
		// around because OneDrive hasn't been persisting LocationRef before now.
		updated := len(entry.LocationRef) > 0 && newLoc.String() != entry.LocationRef

		return newPath, newLoc, updated, nil
	}

	// We didn't have an exact entry, so retry with a location.
	locRef, err := entry.ToLocationIDer(backupVersion)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "getting previous item location")
	}

	if locRef == nil {
		return nil, nil, false, clues.New("entry with empty LocationRef")
	}

	newPath, newLoc, err := dataFromBackup.GetNewPathRefs(
		repoRef.ToBuilder(),
		entry.Modified(),
		locRef)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "getting new paths with old location")
	} else if newPath == nil {
		return nil, nil, false, nil
	} else if newLoc == nil {
		return nil, nil, false, clues.New("unable to get new paths")
	}

	updated := len(entry.LocationRef) > 0 && newLoc.String() != entry.LocationRef

	return newPath, newLoc, updated, nil
}

func mergeDetails(
	ctx context.Context,
	detailsStore streamstore.Streamer,
	backups []kopia.BackupEntry,
	dataFromBackup kopia.DetailsMergeInfoer,
	deets *details.Builder,
	writeStats *kopia.BackupStats,
	serviceType path.ServiceType,
	errs *fault.Bus,
) error {
	detailsModel := deets.Details().DetailsModel

	// getting the values in writeStats before anything else so that we don't get a return from
	// conditions like no backup data.
	writeStats.TotalNonMetaFileCount = len(detailsModel.FilterMetaFiles().Items())
	writeStats.TotalNonMetaUploadedBytes = detailsModel.SumNonMetaFileSizes()

	// Don't bother loading any of the base details if there's nothing we need to merge.
	if dataFromBackup == nil || dataFromBackup.ItemsToMerge() == 0 {
		return nil
	}

	var addedEntries int

	for _, baseBackup := range backups {
		var (
			mctx                 = clues.Add(ctx, "base_backup_id", baseBackup.ID)
			manifestAddedEntries int
		)

		baseDeets, err := getDetailsFromBackup(
			mctx,
			baseBackup.Backup,
			detailsStore,
			errs)
		if err != nil {
			return clues.New("fetching base details for backup")
		}

		for _, entry := range baseDeets.Items() {
			rr, err := path.FromDataLayerPath(entry.RepoRef, true)
			if err != nil {
				return clues.New("parsing base item info path").
					WithClues(mctx).
					With("repo_ref", path.NewElements(entry.RepoRef))
			}

			// Although this base has an entry it may not be the most recent. Check
			// the reasons a snapshot was returned to ensure we only choose the recent
			// entries.
			//
			// TODO(ashmrtn): This logic will need expanded to cover entries from
			// checkpoints if we start doing kopia-assisted incrementals for those.
			if !matchesReason(baseBackup.Reasons, rr) {
				continue
			}

			mctx = clues.Add(mctx, "repo_ref", rr)

			newPath, newLoc, locUpdated, err := getNewPathRefs(
				dataFromBackup,
				entry,
				rr,
				baseBackup.Version)
			if err != nil {
				return clues.Wrap(err, "getting updated info for entry").WithClues(mctx)
			}

			// This entry isn't merged.
			if newPath == nil {
				continue
			}

			// Fixup paths in the item.
			item := entry.ItemInfo
			details.UpdateItem(&item, newLoc)

			// TODO(ashmrtn): This may need updated if we start using this merge
			// strategry for items that were cached in kopia.
			itemUpdated := newPath.String() != rr.String() || locUpdated

			err = deets.Add(
				newPath,
				newLoc,
				itemUpdated,
				item)
			if err != nil {
				return clues.Wrap(err, "adding item to details")
			}

			// Track how many entries we added so that we know if we got them all when
			// we're done.
			addedEntries++
			manifestAddedEntries++
		}

		logger.Ctx(mctx).Infow(
			"merged details with base manifest",
			"base_item_count_unfiltered", len(baseDeets.Items()),
			"base_item_count_added", manifestAddedEntries)
	}

	checkCount := dataFromBackup.ItemsToMerge()

	if addedEntries != checkCount {
		return clues.New("incomplete migration of backup details").
			WithClues(ctx).
			With(
				"item_count", addedEntries,
				"expected_item_count", checkCount)
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

	// Non recoverable errors always result in a failed backup.
	// This holds true for all FailurePolicy.
	if op.Errors.Failure() != nil {
		op.Status = Failed
	}

	// If recoverable errors were seen & failure policy is set to fail after recovery,
	// mark the backup operation as failed.
	if op.Options.FailureHandling == control.FailAfterRecovery &&
		len(op.Errors.Recovered()) > 0 {
		op.Status = Failed
	}

	if opStats.k == nil {
		op.Status = Failed
		return clues.New("backup persistence never completed")
	}

	op.Results.BytesRead = opStats.k.TotalHashedBytes
	op.Results.BytesUploaded = opStats.k.TotalUploadedBytes
	op.Results.ItemsWritten = opStats.k.TotalFileCount
	op.Results.NonMetaBytesUploaded = opStats.k.TotalNonMetaUploadedBytes
	op.Results.NonMetaItemsWritten = opStats.k.TotalNonMetaFileCount
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.ctrl == nil {
		op.Status = Failed
		return clues.New("backup population never completed")
	}

	if op.Status != Failed && opStats.ctrl.IsZero() {
		op.Status = NoData
	}

	op.Results.ItemsRead = opStats.ctrl.Successes

	// Only return non-recoverable errors at this point.
	return op.Errors.Failure()
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	sscw streamstore.CollectorWriter,
	opStats backupStats,
	backupID model.StableID,
	backupVersion int,
	deets *details.Details,
) error {
	snapID := opStats.k.SnapshotID
	ctx = clues.Add(ctx,
		"snapshot_id",
		snapID,
		"backup_id",
		backupID)

	// generate a new fault bus so that we can maintain clean
	// separation between the errors we serialize and those that
	// are generated during the serialization process.
	errs := fault.New(true)

	if deets == nil {
		return clues.New("no backup details to record").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "details_entry_count", len(deets.Entries))

	err := sscw.Collect(ctx, streamstore.DetailsCollector(deets))
	if err != nil {
		return clues.Wrap(err, "collecting details for persistence").WithClues(ctx)
	}

	err = sscw.Collect(ctx, streamstore.FaultErrorsCollector(op.Errors.Errors()))
	if err != nil {
		return clues.Wrap(err, "collecting errors for persistence").WithClues(ctx)
	}

	ssid, err := sscw.Write(ctx, errs)
	if err != nil {
		return clues.Wrap(err, "persisting details and errors").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "streamstore_snapshot_id", ssid)

	isAssistBackup := tagAsAssistBackup(
		opStats.hasNewDetailEntries,
		snapID,
		ssid,
		op.Options.FailureHandling,
		op.Errors)

	ctx = clues.Add(ctx, "is_assist_backup", isAssistBackup)

	tags := map[string]string{
		model.ServiceTag: op.Selectors.PathService().String(),
	}

	// Add tags to mark this backup as either assist or merge. This is used to:
	// 1. Filter assist backups by tag during base selection process
	// 2. Differentiate assist backups from merge backups
	if isAssistBackup {
		tags[model.BackupTypeTag] = model.AssistBackup
	} else {
		tags[model.BackupTypeTag] = model.MergeBackup
	}

	b := backup.New(
		snapID, ssid,
		op.Status.String(),
		backupVersion,
		backupID,
		op.Selectors,
		op.ResourceOwner.ID(),
		op.ResourceOwner.Name(),
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
		op.Errors.Errors(),
		tags)

	logger.Ctx(ctx).Info("creating new backup")

	if err = op.store.Put(ctx, model.BackupSchema, b); err != nil {
		return clues.Wrap(err, "creating backup model").WithClues(ctx)
	}

	return nil
}
