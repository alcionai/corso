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
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
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
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/store"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	ResourceOwner idname.Provider

	Results   BackupResults      `json:"results"`
	Selectors selectors.Selector `json:"selectors"`
	Version   string             `json:"version"`

	// backupVersion ONLY controls the value that gets persisted to the
	// backup model after operation.  It does NOT modify the operation behavior
	// to match the version.  Its inclusion here is, unfortunately, purely to
	// facilitate integration testing that requires a certain backup version, and
	// should be removed when we have a more controlled workaround.
	backupVersion int

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
		operation:     newOperation(opts, bus, kw, sw),
		ResourceOwner: owner,
		Selectors:     selector,
		Version:       "v0",
		backupVersion: version.Backup,
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
	k             *kopia.BackupStats
	gc            *data.CollectionStats
	resourceCount int
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
		// wait for the progress display to clean up
		observe.Complete()
	}()

	ctx, flushMetrics := events.NewMetrics(ctx, logger.Writer{Ctx: ctx})
	defer flushMetrics()

	//-----
	// Precheck
	//-----
	err = Precheck(ctx, op.account, op.Selectors.PathService(), op.Selectors.DiscreteOwner)
	if err != nil {
		logger.CtxErr(ctx, err).Error("running backup")
		op.Errors.Fail(clues.Wrap(err, "running backup"))
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
		"resource_owner", clues.Hide(op.ResourceOwner.Name()),
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

	finalizeErrorHandling(ctx, op.Options, op.Errors, "running backup")
	LogFaultErrors(ctx, op.Errors.Errors(), "running backup")

	// -----
	// Persistence
	// -----

	err = op.persistResults(startTime, &opStats)
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting backup results"))
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

			return op.Errors.Fail(clues.Wrap(e, "forced backup")).Failure()
		}
	}

	err = op.createBackupModels(
		ctx,
		sstore,
		opStats.k.SnapshotID,
		op.Results.BackupID,
		op.backupVersion,
		deets.Details())
	if err != nil {
		op.Errors.Fail(clues.Wrap(err, "persisting backup"))
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
	var (
		reasons           = selectorToReasons(op.Selectors, false)
		fallbackReasons   = makeFallbackReasons(op.Selectors)
		lastBackupVersion = version.NoBackup
	)

	logger.Ctx(ctx).With(
		"control_options", op.Options,
		"selectors", op.Selectors).
		Info("backing up selection")

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1

	mans, mdColls, canUseMetaData, err := produceManifestsAndMetadata(
		ctx,
		op.kopia,
		op.store,
		reasons, fallbackReasons,
		op.account.ID(),
		op.incremental)
	if err != nil {
		return nil, clues.Wrap(err, "producing manifests and metadata")
	}

	if canUseMetaData {
		_, lastBackupVersion, err = lastCompleteBackups(ctx, op.store, mans)
		if err != nil {
			return nil, clues.Wrap(err, "retrieving prior backups")
		}
	}

	cs, ssmb, err := produceBackupDataCollections(
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

	ctx = clues.Add(ctx, "coll_count", len(cs))

	writeStats, deets, toMerge, err := consumeBackupCollections(
		ctx,
		op.kopia,
		op.account.ID(),
		reasons,
		mans,
		cs,
		ssmb,
		backupID,
		op.incremental && canUseMetaData,
		op.Errors)
	if err != nil {
		return nil, clues.Wrap(err, "persisting collection backups")
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
		return nil, clues.Wrap(err, "merging details")
	}

	opStats.gc = op.bp.Wait()

	logger.Ctx(ctx).Debug(opStats.gc)

	return deets, nil
}

func Precheck(
	ctx context.Context,
	acc account.Account,
	service path.ServiceType,
	userID string,
) error {
	if service == path.SharePointService {
		// No "enabled" check required for sharepoint
		return nil
	}

	cred, err := acc.M365Config()
	if err != nil {
		return clues.Wrap(err, "getting creds")
	}

	client, err := api.NewClient(cred)
	if err != nil {
		return clues.Wrap(err, "constructing api client")
	}

	ui, err := client.Users().GetInfo(ctx, userID)
	if err != nil {
		return clues.Wrap(err, "unable to get user info")
	}

	if ui == nil || len(ui.ServicesEnabled) == 0 {
		return graph.ErrServiceNotEnabled
	}

	_, ok := ui.ServicesEnabled[service]
	if !ok {
		return graph.ErrServiceNotEnabled
	}

	return nil
}

func makeFallbackReasons(sel selectors.Selector) []kopia.Reason {
	if sel.PathService() != path.SharePointService &&
		sel.DiscreteOwner != sel.DiscreteOwnerName {
		return selectorToReasons(sel, true)
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
	resourceOwner idname.Provider,
	sel selectors.Selector,
	metadata []data.RestoreCollection,
	lastBackupVersion int,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, error) {
	complete, closer := observe.MessageWithCompletion(ctx, "Discovering items to backup")
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	return bp.ProduceBackupCollections(ctx, resourceOwner, sel, metadata, lastBackupVersion, ctrlOpts, errs)
}

// ---------------------------------------------------------------------------
// Consumer funcs
// ---------------------------------------------------------------------------

func selectorToReasons(sel selectors.Selector, useOwnerNameForID bool) []kopia.Reason {
	service := sel.PathService()
	reasons := []kopia.Reason{}

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
			reasons = append(reasons, kopia.Reason{
				ResourceOwner: owner,
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
func consumeBackupCollections(
	ctx context.Context,
	bc inject.BackupConsumer,
	tenantID string,
	reasons []kopia.Reason,
	mans []*kopia.ManifestEntry,
	cs []data.BackupCollection,
	pmr prefixmatcher.StringSetReader,
	backupID model.StableID,
	isIncremental bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error) {
	complete, closer := observe.MessageWithCompletion(ctx, "Backing up data")
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
				return nil, nil, nil, clues.Wrap(err, "getting subtree paths for bases")
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

		mbID, ok := m.GetTag(kopia.TagBackupID)
		if !ok {
			mbID = "no_backup_id_tag"
		}

		logger.Ctx(ctx).Infow(
			"using base for backup",
			"base_snapshot_id", m.ID,
			"services", svcs,
			"categories", cats,
			"base_backup_id", mbID)
	}

	kopiaStats, deets, itemsSourcedFromBase, err := bc.ConsumeBackupCollections(
		ctx,
		bases,
		cs,
		pmr,
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

	ctx = clues.Add(
		ctx,
		"kopia_errors", kopiaStats.ErrorCount,
		"kopia_ignored_errors", kopiaStats.IgnoredErrorCount,
		"kopia_expected_ignored_errors", kopiaStats.ExpectedIgnoredErrorCount)

	if kopiaStats.ErrorCount > 0 {
		err = clues.New("building kopia snapshot").WithClues(ctx)
	} else if kopiaStats.IgnoredErrorCount > kopiaStats.ExpectedIgnoredErrorCount {
		err = clues.New("downloading items for persistence").WithClues(ctx)
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
		newPath, newLoc, err := dataFromBackup.GetNewPathRefs(repoRef.ToBuilder(), nil)
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

	newPath, newLoc, err := dataFromBackup.GetNewPathRefs(repoRef.ToBuilder(), locRef)
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

func lastCompleteBackups(
	ctx context.Context,
	ms *store.Wrapper,
	mans []*kopia.ManifestEntry,
) (map[string]*backup.Backup, int, error) {
	var (
		oldestVersion = version.NoBackup
		result        = map[string]*backup.Backup{}
	)

	if len(mans) == 0 {
		return result, -1, nil
	}

	for _, man := range mans {
		// For now skip snapshots that aren't complete. We will need to revisit this
		// when we tackle restartability.
		if len(man.IncompleteReason) > 0 {
			continue
		}

		var (
			mctx    = clues.Add(ctx, "base_manifest_id", man.ID)
			reasons = man.Reasons
		)

		bID, ok := man.GetTag(kopia.TagBackupID)
		if !ok {
			return result, oldestVersion, clues.New("no backup ID in snapshot manifest").WithClues(mctx)
		}

		mctx = clues.Add(mctx, "base_manifest_backup_id", bID)

		bup, err := getBackupFromID(mctx, model.StableID(bID), ms)
		if err != nil {
			return result, oldestVersion, err
		}

		for _, r := range reasons {
			result[r.Key()] = bup
		}

		if oldestVersion == -1 || bup.Version < oldestVersion {
			oldestVersion = bup.Version
		}
	}

	return result, oldestVersion, nil
}

func mergeDetails(
	ctx context.Context,
	ms *store.Wrapper,
	detailsStore streamstore.Streamer,
	mans []*kopia.ManifestEntry,
	dataFromBackup kopia.DetailsMergeInfoer,
	deets *details.Builder,
	errs *fault.Bus,
) error {
	// Don't bother loading any of the base details if there's nothing we need to merge.
	if dataFromBackup == nil || dataFromBackup.ItemsToMerge() == 0 {
		return nil
	}

	var addedEntries int

	for _, man := range mans {
		var (
			mctx                 = clues.Add(ctx, "base_manifest_id", man.ID)
			manifestAddedEntries int
		)

		// For now skip snapshots that aren't complete. We will need to revisit this
		// when we tackle restartability.
		if len(man.IncompleteReason) > 0 {
			continue
		}

		bID, ok := man.GetTag(kopia.TagBackupID)
		if !ok {
			return clues.New("no backup ID in snapshot manifest").WithClues(mctx)
		}

		mctx = clues.Add(mctx, "base_manifest_backup_id", bID)

		baseBackup, baseDeets, err := getBackupAndDetailsFromID(
			mctx,
			model.StableID(bID),
			ms,
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
			if !matchesReason(man.Reasons, rr) {
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
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.gc == nil {
		op.Status = Failed
		return clues.New("backup population never completed")
	}

	if op.Status != Failed && opStats.gc.IsZero() {
		op.Status = NoData
	}

	op.Results.ItemsRead = opStats.gc.Successes

	return op.Errors.Failure()
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	sscw streamstore.CollectorWriter,
	snapID string,
	backupID model.StableID,
	backupVersion int,
	deets *details.Details,
) error {
	ctx = clues.Add(ctx, "snapshot_id", snapID, "backup_id", backupID)
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
		op.Errors.Errors())

	logger.Ctx(ctx).Info("creating new backup")

	if err = op.store.Put(ctx, model.BackupSchema, b); err != nil {
		return clues.Wrap(err, "creating backup model").WithClues(ctx)
	}

	return nil
}
