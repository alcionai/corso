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
	"github.com/alcionai/corso/src/pkg/backup/identity"
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
	// When true, disables kopia-assisted incremental backups. This forces
	// downloading and hashing all item data for items not in the merge base(s).
	disableAssistBackup bool
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
	sw store.BackupStorer,
	bp inject.BackupProducer,
	acct account.Account,
	selector selectors.Selector,
	owner idname.Provider,
	bus events.Eventer,
) (BackupOperation, error) {
	op := BackupOperation{
		operation:           newOperation(opts, bus, count.New(), kw, sw),
		ResourceOwner:       owner,
		Selectors:           selector,
		Version:             "v0",
		BackupVersion:       version.Backup,
		account:             acct,
		incremental:         useIncrementalBackup(selector, opts),
		disableAssistBackup: opts.ToggleFeatures.ForceItemDataDownload,
		bp:                  bp,
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

// An assist backup must meet the following criteria:
// 1. new detail entries were produced
// 2. valid details ssid & item snapshot ID
// 3. no non-recoverable errors
// 4. we observed recoverable errors
// 5. not running in best effort mode. Reason being that there is
// no way to distinguish assist backups from merge backups in best effort mode.
//
// Primary reason for persisting assist backup models is to ensure we don't
// lose corso extension data(deets) in the event of recoverable failures.
//
// Note: kopia.DetailsMergeInfoer doesn't impact decision making for creating
// assist backups. It may be empty if it’s the very first backup so there is no
// merge base to source base details from, or non-empty, if there was a merge
// base. In summary, if there are no new deets, no new extension data was produced
// and hence no need to persist assist backup model.
func isAssistBackup(
	newDeetsProduced bool,
	snapID, ssid string,
	failurePolicy control.FailurePolicy,
	err *fault.Bus,
) bool {
	return newDeetsProduced &&
		len(snapID) > 0 &&
		len(ssid) > 0 &&
		failurePolicy != control.BestEffort &&
		err.Failure() == nil &&
		len(err.Recovered()) > 0
}

// A merge backup must meet the following criteria:
// 1. valid details ssid & item snapshot ID
// 2. zero recoverable errors
// 3. no recoverable errors if not running in best effort mode
func isMergeBackup(
	snapID, ssid string,
	failurePolicy control.FailurePolicy,
	err *fault.Bus,
) bool {
	if len(snapID) == 0 || len(ssid) == 0 {
		return false
	}

	if err.Failure() != nil {
		return false
	}

	if failurePolicy == control.BestEffort {
		return true
	}

	return len(err.Recovered()) == 0
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

	// IsServiceEnabled checks if the user has the service enabled to run a backup.
	runnable, err := op.bp.IsServiceEnabled(
		ctx,
		op.Selectors.PathService(),
		op.ResourceOwner.ID())
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
		"incremental", op.incremental,
		"disable_assist_backup", op.disableAssistBackup)

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

	finalizeErrorHandling(ctx, op.Options, op.Errors, "running backup")

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
	lastBackupVersion := version.NoBackup

	reasons, err := op.Selectors.Reasons(op.account.ID(), false)
	if err != nil {
		return nil, clues.Wrap(err, "getting reasons")
	}

	fallbackReasons, err := makeFallbackReasons(op.account.ID(), op.Selectors)
	if err != nil {
		return nil, clues.Wrap(err, "getting fallback reasons")
	}

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
		op.incremental,
		op.disableAssistBackup)
	if err != nil {
		return nil, clues.Wrap(err, "producing manifests and metadata")
	}

	ctx = clues.Add(ctx, "can_use_metadata", canUseMetadata)

	if canUseMetadata {
		lastBackupVersion = mans.MinBackupVersion()
	}

	// TODO(ashmrtn): This should probably just return a collection that deletes
	// the entire subtree instead of returning an additional bool. That way base
	// selection is controlled completely by flags and merging is controlled
	// completely by collections.
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

	opStats.hasNewDetailEntries = (deets != nil && !deets.Empty()) ||
		(toMerge != nil && toMerge.ItemsToMerge() > 0)
	opStats.k = writeStats

	err = mergeDetails(
		ctx,
		detailsStore,
		mans,
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

func makeFallbackReasons(tenant string, sel selectors.Selector) ([]identity.Reasoner, error) {
	if sel.PathService() != path.SharePointService &&
		sel.DiscreteOwner != sel.DiscreteOwnerName {
		return sel.Reasons(tenant, true)
	}

	// return nil for fallback reasons since a nil value will no-op.
	return nil, nil
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

// calls kopia to backup the collections of data
func consumeBackupCollections(
	ctx context.Context,
	bc kinject.BackupConsumer,
	tenantID string,
	reasons []identity.Reasoner,
	bbs kopia.BackupBases,
	cs []data.BackupCollection,
	pmr prefixmatcher.StringSetReader,
	backupID model.StableID,
	isIncremental bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error) {
	ctx = clues.Add(
		ctx,
		"collection_source", "operations",
		"snapshot_type", "item data")

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

func matchesReason(reasons []identity.Reasoner, p path.Path) bool {
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
) (path.Path, *path.Builder, error) {
	locRef, err := entry.ToLocationIDer(backupVersion)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting previous item location")
	}

	if locRef == nil {
		return nil, nil, clues.New("entry with empty LocationRef")
	}

	newPath, newLoc, err := dataFromBackup.GetNewPathRefs(
		repoRef.ToBuilder(),
		entry.Modified(),
		locRef)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting new paths with old location")
	} else if newPath == nil {
		return nil, nil, nil
	} else if newLoc == nil {
		return nil, nil, clues.New("unable to get new paths")
	}

	return newPath, newLoc, nil
}

func mergeItemsFromBase(
	ctx context.Context,
	checkReason bool,
	baseBackup kopia.BackupEntry,
	detailsStore streamstore.Streamer,
	dataFromBackup kopia.DetailsMergeInfoer,
	deets *details.Builder,
	alreadySeenItems map[string]struct{},
	errs *fault.Bus,
) (int, error) {
	var (
		manifestAddedEntries int
		totalBaseItems       int
	)

	// Can't be in the above block else it's counted as a redeclaration.
	ctx = clues.Add(ctx, "base_backup_id", baseBackup.ID)

	baseDeets, err := getDetailsFromBackup(
		ctx,
		baseBackup.Backup,
		detailsStore,
		errs)
	if err != nil {
		return manifestAddedEntries,
			clues.New("fetching base details for backup").WithClues(ctx)
	}

	for _, entry := range baseDeets.Items() {
		// Track this here instead of calling Items() again to get the count since
		// it can be a bit expensive.
		totalBaseItems++

		rr, err := path.FromDataLayerPath(entry.RepoRef, true)
		if err != nil {
			return manifestAddedEntries, clues.New("parsing base item info path").
				WithClues(ctx).
				With("repo_ref", path.LoggableDir(entry.RepoRef))
		}

		// Although this base has an entry it may not be the most recent. Check
		// the reasons a snapshot was returned to ensure we only choose the recent
		// entries.
		//
		// We only really want to do this check for merge bases though because
		// kopia won't abide by reasons when determining if an item's cached. This
		// leaves us in a bit of a pickle if the user has run any concurrent backups
		// with overlapping reasons that then turn into assist bases, but the
		// modTime check in DetailsMergeInfoer should handle that.
		if checkReason && !matchesReason(baseBackup.Reasons, rr) {
			continue
		}

		// Skip items that were already found in a previous base backup.
		if _, ok := alreadySeenItems[rr.ShortRef()]; ok {
			continue
		}

		ictx := clues.Add(ctx, "repo_ref", rr)

		newPath, newLoc, err := getNewPathRefs(
			dataFromBackup,
			entry,
			rr,
			baseBackup.Version)
		if err != nil {
			return manifestAddedEntries,
				clues.Wrap(err, "getting updated info for entry").WithClues(ictx)
		}

		// This entry isn't merged.
		if newPath == nil {
			continue
		}

		// Fixup paths in the item.
		item := entry.ItemInfo
		details.UpdateItem(&item, newLoc)

		err = deets.Add(
			newPath,
			newLoc,
			item)
		if err != nil {
			return manifestAddedEntries,
				clues.Wrap(err, "adding item to details").WithClues(ictx)
		}

		// Make sure we won't add this again in another base.
		alreadySeenItems[rr.ShortRef()] = struct{}{}

		// Track how many entries we added so that we know if we got them all when
		// we're done.
		manifestAddedEntries++
	}

	logger.Ctx(ctx).Infow(
		"merged details with base manifest",
		"count_base_item_unfiltered", totalBaseItems,
		"count_base_item_added", manifestAddedEntries)

	return manifestAddedEntries, nil
}

func mergeDetails(
	ctx context.Context,
	detailsStore streamstore.Streamer,
	bases kopia.BackupBases,
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
	if bases == nil || dataFromBackup == nil || dataFromBackup.ItemsToMerge() == 0 {
		return nil
	}

	var (
		addedEntries int
		// alreadySeenEntries tracks items that we've already merged so we don't
		// accidentally merge them again. This could happen if, for example, there's
		// an assist backup and a merge backup that both have the same version of an
		// item at the same path.
		alreadySeenEntries = map[string]struct{}{}
	)

	// Merge details from assist bases first. It shouldn't technically matter
	// since the DetailsMergeInfoer should take into account the modTime of items,
	// but just to be on the safe side.
	//
	// We don't want to match entries based on Reason for assist bases because
	// kopia won't abide by Reasons when determining if an item's cached. This
	// leaves us in a bit of a pickle if the user has run any concurrent backups
	// with overlapping Reasons that turn into assist bases, but the modTime check
	// in DetailsMergeInfoer should handle that.
	for _, base := range bases.AssistBackups() {
		added, err := mergeItemsFromBase(
			ctx,
			false,
			base,
			detailsStore,
			dataFromBackup,
			deets,
			alreadySeenEntries,
			errs)
		if err != nil {
			return clues.Wrap(err, "merging assist backup base details")
		}

		addedEntries = addedEntries + added
	}

	// Now add entries from the merge base backups. These will be things that
	// weren't changed in the new backup. Items that were already added because
	// they were counted as cached in an assist base backup will be skipped due to
	// alreadySeenEntries.
	//
	// We do want to enable matching entries based on Reasons because we
	// explicitly control which subtrees from the merge base backup are grafted
	// onto the hierarchy for the currently running backup.
	for _, base := range bases.Backups() {
		added, err := mergeItemsFromBase(
			ctx,
			true,
			base,
			detailsStore,
			dataFromBackup,
			deets,
			alreadySeenEntries,
			errs)
		if err != nil {
			return clues.Wrap(err, "merging merge backup base details")
		}

		addedEntries = addedEntries + added
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
		"snapshot_id", snapID,
		"backup_id", backupID)

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

	tags := map[string]string{
		model.ServiceTag: op.Selectors.PathService().String(),
	}

	// Add tags to mark this backup as either assist or merge. This is used to:
	// 1. Filter assist backups by tag during base selection process
	// 2. Differentiate assist backups from merge backups
	if isMergeBackup(
		snapID,
		ssid,
		op.Options.FailureHandling,
		op.Errors) {
		tags[model.BackupTypeTag] = model.MergeBackup
	} else if isAssistBackup(
		opStats.hasNewDetailEntries,
		snapID,
		ssid,
		op.Options.FailureHandling,
		op.Errors) {
		tags[model.BackupTypeTag] = model.AssistBackup
	} else {
		return clues.New("backup is neither assist nor merge").WithClues(ctx)
	}

	ctx = clues.Add(ctx, model.BackupTypeTag, tags[model.BackupTypeTag])

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
