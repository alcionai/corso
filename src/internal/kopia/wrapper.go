package kopia

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/maintenance"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/kopia/kopia/snapshot/snapshotmaintenance"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/store"
)

const defaultCorsoPin = "corso"

// common manifest tags
const (
	TagBackupID       = "backup-id"
	TagBackupCategory = "is-canon-backup"
)

var (
	errNotConnected  = clues.New("not connected to repo")
	ErrNoRestorePath = clues.New("no restore path given")
)

type BackupStats struct {
	SnapshotID string

	TotalHashedBytes          int64
	TotalUploadedBytes        int64
	TotalNonMetaUploadedBytes int64

	TotalFileCount        int
	TotalNonMetaFileCount int
	CachedFileCount       int
	UncachedFileCount     int
	TotalDirectoryCount   int
	ErrorCount            int

	IgnoredErrorCount         int
	ExpectedIgnoredErrorCount int

	Incomplete       bool
	IncompleteReason string
}

func manifestToStats(
	man *snapshot.Manifest,
	progress *corsoProgress,
	uploadCount *stats.ByteCounter,
) BackupStats {
	progress.counter.Add(count.PersistedFiles, int64(man.Stats.TotalFileCount))
	progress.counter.Add(count.PersistedCachedFiles, int64(man.Stats.CachedFiles))
	progress.counter.Add(count.PersistedNonCachedFiles, int64(man.Stats.NonCachedFiles))
	progress.counter.Add(count.PersistedDirectories, int64(man.Stats.TotalDirectoryCount))
	progress.counter.Add(count.PersistenceErrs, int64(man.Stats.ErrorCount))
	progress.counter.Add(count.PersistenceIgnoredErrs, int64(man.Stats.IgnoredErrorCount))

	return BackupStats{
		SnapshotID: string(man.ID),

		TotalHashedBytes:   progress.totalBytes,
		TotalUploadedBytes: uploadCount.NumBytes,

		TotalFileCount:      int(man.Stats.TotalFileCount),
		CachedFileCount:     int(man.Stats.CachedFiles),
		UncachedFileCount:   int(man.Stats.NonCachedFiles),
		TotalDirectoryCount: int(man.Stats.TotalDirectoryCount),
		ErrorCount:          int(man.Stats.ErrorCount),

		IgnoredErrorCount:         int(man.Stats.IgnoredErrorCount),
		ExpectedIgnoredErrorCount: progress.expectedIgnoredErrors,

		Incomplete:       man.IncompleteReason != "",
		IncompleteReason: man.IncompleteReason,
	}
}

func NewWrapper(c *conn) (*Wrapper, error) {
	if err := c.wrap(); err != nil {
		return nil, clues.Wrap(err, "creating Wrapper")
	}

	return &Wrapper{c}, nil
}

// FIXME: Circular references.
// must comply with restore producer and backup consumer
// var (
// _ inject.BackupConsumer  = &Wrapper{}
// _ inject.RestoreProducer = &Wrapper{}
// )

type Wrapper struct {
	c *conn
}

func (w *Wrapper) Close(ctx context.Context) error {
	if w.c == nil {
		return nil
	}

	err := w.c.Close(ctx)
	w.c = nil

	if err != nil {
		return clues.WrapWC(ctx, err, "closing Wrapper")
	}

	return nil
}

// ConsumeBackupCollections takes a set of collections and creates a kopia snapshot
// with the data that they contain. previousSnapshots is used for incremental
// backups and should represent the base snapshot from which metadata is sourced
// from as well as any incomplete snapshot checkpoints that may contain more
// recent data than the base snapshot. The absence of previousSnapshots causes a
// complete backup of all data.
func (w Wrapper) ConsumeBackupCollections(
	ctx context.Context,
	backupReasons []identity.Reasoner,
	bases BackupBases,
	collections []data.BackupCollection,
	globalExcludeSet prefixmatcher.StringSetReader,
	additionalTags map[string]string,
	buildTreeWithBase bool,
	counter *count.Bus,
	errs *fault.Bus,
) (*BackupStats, *details.Builder, DetailsMergeInfoer, error) {
	if w.c == nil {
		return nil, nil, nil, clues.StackWC(ctx, errNotConnected)
	}

	ctx, end := diagnostics.Span(ctx, "kopia:consumeBackupCollections")
	defer end()

	if len(collections) == 0 && (globalExcludeSet == nil || globalExcludeSet.Empty()) {
		return &BackupStats{}, &details.Builder{}, nil, nil
	}

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		deets:   &details.Builder{},
		toMerge: newMergeDetails(),
		errs:    errs,
		counter: counter,
	}

	// When running an incremental backup, we need to pass the prior
	// snapshot bases into inflateDirTree so that the new snapshot
	// includes historical data.
	var (
		mergeBase  []BackupBase
		assistBase []BackupBase
	)

	if bases != nil {
		if buildTreeWithBase {
			mergeBase = bases.MergeBases()
		}

		assistBase = bases.SnapshotAssistBases()
	}

	dirTree, err := inflateDirTree(
		ctx,
		w.c,
		mergeBase,
		collections,
		globalExcludeSet,
		progress)
	if err != nil {
		return nil, nil, nil, clues.Wrap(err, "building kopia directories")
	}

	s, err := w.makeSnapshotWithRoot(
		ctx,
		backupReasons,
		assistBase,
		dirTree,
		additionalTags,
		progress)
	if err != nil {
		return nil, nil, nil, err
	}

	return s, progress.deets, progress.toMerge, progress.errs.Failure()
}

// userAndHost is used as a passing mechanism for values that will be fed into
// kopia's UserName and Host fields for SourceInfo. It exists to avoid returning
// two strings from the hostAndUserFromReasons function.
type userAndHost struct {
	user string
	host string
}

func hostAndUserFromReasons(reasons []identity.Reasoner) (userAndHost, error) {
	var (
		tenant   string
		resource string
		// reasonMap is a hash set of the concatenation of the service and category.
		reasonMap = map[string]struct{}{}
	)

	for i, reason := range reasons {
		// Use a check on the iteration index instead of empty string so we can
		// differentiate between the first iteration and a reason with an empty
		// value (should result in an error if there's another reason with a
		// non-empty value).
		if i == 0 {
			tenant = reason.Tenant()
		} else if tenant != reason.Tenant() {
			return userAndHost{}, clues.New("multiple tenant IDs in backup reasons").
				With(
					"old_tenant_id", tenant,
					"new_tenant_id", reason.Tenant())
		}

		if i == 0 {
			resource = reason.ProtectedResource()
		} else if resource != reason.ProtectedResource() {
			return userAndHost{}, clues.New("multiple protected resource IDs in backup reasons").
				With(
					"old_resource_id", resource,
					"new_resource_id", reason.ProtectedResource())
		}

		dataType := reason.Service().String() + reason.Category().String()
		reasonMap[dataType] = struct{}{}
	}

	allReasons := maps.Keys(reasonMap)
	slices.Sort(allReasons)

	host := strings.Join(allReasons, "-")
	user := strings.Join([]string{tenant, resource}, "-")

	if len(user) == 0 || user == "-" {
		return userAndHost{}, clues.New("empty user value")
	}

	if len(host) == 0 {
		return userAndHost{}, clues.New("empty host value")
	}

	return userAndHost{
		host: host,
		user: user,
	}, nil
}

func (w Wrapper) makeSnapshotWithRoot(
	ctx context.Context,
	backupReasons []identity.Reasoner,
	prevBases []BackupBase,
	root fs.Directory,
	addlTags map[string]string,
	progress *corsoProgress,
) (*BackupStats, error) {
	var (
		man *snapshot.Manifest
		bc  = &stats.ByteCounter{
			// duplicate the count in the progress count.Bus.  Later we can
			// replace the ByteCounter with the progress counter entirely.
			Counter: progress.counter.AdderFor(count.PersistedUploadedBytes),
		}
	)

	snapIDs := make([]manifest.ID, 0, len(prevBases)) // just for logging
	prevSnaps := make([]*snapshot.Manifest, 0, len(prevBases))

	for _, ent := range prevBases {
		prevSnaps = append(prevSnaps, ent.ItemDataSnapshot)
		snapIDs = append(snapIDs, ent.ItemDataSnapshot.ID)
	}

	// Add some extra tags so we can look things up by reason.
	allTags := maps.Clone(addlTags)
	if allTags == nil {
		// Some platforms seem to return nil if the input is nil.
		allTags = map[string]string{}
	}

	for _, r := range backupReasons {
		for _, k := range tagKeys(r) {
			allTags[k] = ""
		}
	}

	ctx = clues.Add(
		ctx,
		"num_assist_snapshots", len(prevBases),
		"assist_snapshot_ids", snapIDs,
		"additional_tags", allTags)

	if len(snapIDs) > 0 {
		logger.Ctx(ctx).Info("using snapshots for kopia-assisted incrementals")
	} else {
		logger.Ctx(ctx).Info("no base snapshots for kopia-assisted incrementals")
	}

	tags := map[string]string{}

	for k, v := range allTags {
		mk, mv := makeTagKV(k)

		if len(v) == 0 {
			v = mv
		}

		tags[mk] = v
	}

	// Set the SourceInfo to the tenant ID, resource ID, and the concatenation
	// of the service/data types being backed up. This will give us unique
	// values for each set of backups with the assumption that no concurrent
	// backups for the same set of things is being run on this repo.
	userHost, err := hostAndUserFromReasons(backupReasons)
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	err = repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{
			Purpose: "KopiaWrapperBackup",
			// Always flush so we don't leak write sessions. Still uses reachability
			// for consistency.
			FlushOnFailure: true,
			OnUpload:       bc.Count,
		},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			si := snapshot.SourceInfo{
				Host:     userHost.host,
				UserName: userHost.user,
				Path:     root.Name(),
			}

			trueVal := policy.OptionalBool(true)
			errPolicy := &policy.Policy{
				ErrorHandlingPolicy: policy.ErrorHandlingPolicy{
					IgnoreFileErrors:      &trueVal,
					IgnoreDirectoryErrors: &trueVal,
				},
			}

			policyTree, err := policy.TreeForSourceWithOverride(innerCtx, w.c, si, errPolicy)
			if err != nil {
				err = clues.WrapWC(ctx, err, "get policy tree")
				logger.CtxErr(innerCtx, err).Error("building kopia backup")

				return err
			}

			// By default Uploader is best-attempt.
			u := snapshotfs.NewUploader(rw)
			progress.UploadProgress = u.Progress
			u.Progress = progress
			u.CheckpointLabels = tags

			man, err = u.Upload(innerCtx, root, policyTree, si, prevSnaps...)
			if err != nil {
				err = clues.WrapWC(ctx, err, "uploading data")
				logger.CtxErr(innerCtx, err).Error("uploading kopia backup")

				return err
			}

			man.Tags = tags
			// Add one pin to keep kopia's retention policy from collecting it if it
			// ends up enabled for some reason. The value in the pin doesn't matter.
			// We don't need to remove any pins.
			man.UpdatePins(append(man.Pins, defaultCorsoPin), nil)

			if _, err := snapshot.SaveSnapshot(innerCtx, rw, man); err != nil {
				err = clues.WrapWC(ctx, err, "saving snapshot")
				logger.CtxErr(innerCtx, err).Error("persisting kopia backup snapshot")

				return err
			}

			return nil
		})
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "kopia backup")
	}

	res := manifestToStats(man, progress, bc)

	return &res, nil
}

func (w Wrapper) getSnapshotRoot(
	ctx context.Context,
	snapshotID string,
) (fs.Entry, error) {
	man, err := snapshot.LoadSnapshot(ctx, w.c, manifest.ID(snapshotID))
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, man)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "getting root directory")
	}

	return rootDirEntry, nil
}

// getDir looks up the directory at the given path starting from snapshotRoot.
// If the item is a directory in kopia then it returns the kopia fs.Directory
// handle. If the item does not exist in kopia or is not a directory an error is
// returned.
func getDir(
	ctx context.Context,
	dirPath path.Path,
	snapshotRoot fs.Entry,
) (fs.Directory, error) {
	if dirPath == nil {
		return nil, clues.WrapWC(ctx, ErrNoRestorePath, "getting directory")
	}

	toGet := dirPath.PopFront()

	ctx = clues.Add(ctx, "entry_path", toGet)

	// GetNestedEntry handles nil properly.
	e, err := snapshotfs.GetNestedEntry(
		ctx,
		snapshotRoot,
		encodeElements(toGet.Elements()...))
	if err != nil {
		if isErrEntryNotFound(err) {
			err = clues.StackWC(ctx, data.ErrNotFound, err)
		}

		return nil, clues.WrapWC(ctx, err, "getting nested object handle")
	}

	f, ok := e.(fs.Directory)
	if !ok {
		return nil, clues.NewWC(ctx, "requested object is not a directory")
	}

	return f, nil
}

type ByteCounter interface {
	Count(numBytes int64)
}

type restoreCollection struct {
	restorePath path.Path
	storageDirs map[string]*dirAndItems
}

type dirAndItems struct {
	dir   path.Path
	items []string
}

// loadDirsAndItems takes a set of ShortRef -> (directory path, []item names)
// and creates a collection for each tuple in the set. Non-fatal errors are
// accumulated into bus. Any fatal errors will stop processing and return the
// error directly.
//
// All data is loaded from the given snapshot.
func loadDirsAndItems(
	ctx context.Context,
	snapshotRoot fs.Entry,
	bcounter ByteCounter,
	toLoad map[string]*restoreCollection,
	bus *fault.Bus,
) ([]data.RestoreCollection, error) {
	var (
		el  = bus.Local()
		res = make([]data.RestoreCollection, 0, len(toLoad))
	)

	for _, col := range toLoad {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		ictx := clues.Add(ctx, "restore_path", col.restorePath)

		mergeCol := &mergeCollection{fullPath: col.restorePath}
		res = append(res, mergeCol)

		for _, dirItems := range col.storageDirs {
			if el.Failure() != nil {
				return nil, el.Failure()
			}

			ictx = clues.Add(ictx, "storage_directory_path", dirItems.dir)

			dir, err := getDir(ictx, dirItems.dir, snapshotRoot)
			if err != nil {
				el.AddRecoverable(ctx, clues.WrapWC(ictx, err, "loading storage directory").
					Label(fault.LabelForceNoBackupCreation))

				continue
			}

			dc := &kopiaDataCollection{
				path:            col.restorePath,
				dir:             dir,
				items:           dirItems.items,
				counter:         bcounter,
				expectedVersion: readers.DefaultSerializationVersion,
			}

			if err := mergeCol.addCollection(dirItems.dir.String(), dc); err != nil {
				el.AddRecoverable(ctx, clues.WrapWC(ictx, err, "adding collection to merge collection").
					Label(fault.LabelForceNoBackupCreation))

				continue
			}
		}
	}

	return res, el.Failure()
}

// ProduceRestoreCollections looks up all paths- assuming each is an item declaration,
// not a directory- in the snapshot with id snapshotID. The path should be the
// full path of the item from the root.  Returns the results as a slice of single-
// item DataCollections, where the DataCollection.FullPath() matches the path.
// If the item does not exist in kopia or is not a file an error is returned.
// The UUID of the returned DataStreams will be the name of the kopia file the
// data is sourced from.
func (w Wrapper) ProduceRestoreCollections(
	ctx context.Context,
	snapshotID string,
	paths []path.RestorePaths,
	bcounter ByteCounter,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	ctx, end := diagnostics.Span(ctx, "kopia:produceRestoreCollections")
	defer end()

	if len(paths) == 0 {
		return nil, clues.StackWC(ctx, ErrNoRestorePath)
	}

	// Used later on, but less confusing to follow error propagation if we just
	// load it here.
	snapshotRoot, err := w.getSnapshotRoot(ctx, snapshotID)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "loading snapshot root")
	}

	var (
		loadCount int
		// RestorePath -> []StoragePath directory -> set of items to load from the
		// directory.
		dirsToItems = map[string]*restoreCollection{}
		el          = errs.Local()
	)

	for _, itemPaths := range paths {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		// Group things by RestorePath and then StoragePath so we can load multiple
		// items from a single directory instance lower down.
		ictx := clues.Add(
			ctx,
			"item_path", itemPaths.StoragePath,
			"restore_path", itemPaths.RestorePath)

		parentStoragePath, err := itemPaths.StoragePath.Dir()
		if err != nil {
			el.AddRecoverable(ictx, clues.WrapWC(ictx, err, "getting storage directory path").
				Label(fault.LabelForceNoBackupCreation))

			continue
		}

		// Find the location this item is restored to.
		rc := dirsToItems[itemPaths.RestorePath.ShortRef()]
		if rc == nil {
			dirsToItems[itemPaths.RestorePath.ShortRef()] = &restoreCollection{
				restorePath: itemPaths.RestorePath,
				storageDirs: map[string]*dirAndItems{},
			}
			rc = dirsToItems[itemPaths.RestorePath.ShortRef()]
		}

		// Find the collection this item is sourced from.
		di := rc.storageDirs[parentStoragePath.ShortRef()]
		if di == nil {
			rc.storageDirs[parentStoragePath.ShortRef()] = &dirAndItems{
				dir: parentStoragePath,
			}
			di = rc.storageDirs[parentStoragePath.ShortRef()]
		}

		di.items = append(di.items, itemPaths.StoragePath.Item())

		loadCount++
		if loadCount%1000 == 0 {
			logger.Ctx(ctx).Infow(
				"grouping items to load from kopia",
				"group_items", loadCount)
		}
	}

	// Now that we've grouped everything, go through and load each directory and
	// then load the items from the directory.
	res, err := loadDirsAndItems(ctx, snapshotRoot, bcounter, dirsToItems, errs)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "loading items")
	}

	return res, el.Failure()
}

func (w Wrapper) NewBaseFinder(bg store.BackupGetter) (*baseFinder, error) {
	return newBaseFinder(w.c, bg)
}

func isErrEntryNotFound(err error) bool {
	// Calling Child on a directory may return this.
	if errors.Is(err, fs.ErrEntryNotFound) {
		return true
	}

	// This is returned when walking the hierarchy of a backup.
	return strings.Contains(err.Error(), "entry not found") &&
		!strings.Contains(err.Error(), "parent is not a directory")
}

func (w Wrapper) RepoMaintenance(
	ctx context.Context,
	storer store.Storer,
	opts repository.Maintenance,
	errs *fault.Bus,
) error {
	// Check the existing config parameters first so that even if we fail for some
	// reason below we know we checked the config.
	w.c.verifyDefaultConfigOptions(ctx, errs)

	kopiaSafety, err := translateSafety(opts.Safety)
	if err != nil {
		return clues.WrapWC(ctx, err, "identifying safety level")
	}

	mode, err := translateMode(opts.Type)
	if err != nil {
		return clues.WrapWC(ctx, err, "identifying maintenance mode")
	}

	currentOwner := w.c.ClientOptions().UsernameAtHost()

	ctx = clues.Add(
		ctx,
		"kopia_safety", kopiaSafety,
		"kopia_maintenance_mode", mode,
		"force", opts.Force,
		"current_local_owner", clues.Hide(currentOwner))

	// Check if we should do additional cleanup prior to running kopia's
	// maintenance.
	if opts.Type == repository.CompleteMaintenance {
		buffer := time.Hour * 24 * 7
		if opts.CleanupBuffer != nil {
			buffer = *opts.CleanupBuffer
		}

		// Even if we fail this we don't want to fail the overall maintenance
		// operation since there's other useful work we can still do.
		if err := cleanupOrphanedData(ctx, storer, w.c, buffer, time.Now); err != nil {
			errs.AddRecoverable(ctx, clues.Wrap(
				err,
				"cleaning up failed backups, some space may not be freed"))
		}
	}

	dr, ok := w.c.Repository.(repo.DirectRepository)
	if !ok {
		return clues.NewWC(ctx, "unable to get valid handle to repo")
	}

	// Below write session options pulled from kopia's CLI code that runs
	// maintenance.
	err = repo.DirectWriteSession(
		ctx,
		dr,
		repo.WriteSessionOptions{
			Purpose: "Corso maintenance",
		},
		func(ctx context.Context, dw repo.DirectRepositoryWriter) error {
			params, err := maintenance.GetParams(ctx, w.c)
			if err != nil {
				return clues.WrapWC(ctx, err, "getting maintenance user@host")
			}

			// Need to do some fixup here as the user/host may not have been set.
			if len(params.Owner) == 0 || (params.Owner != currentOwner && opts.Force) {
				observe.Message(
					ctx,
					observe.ProgressCfg{},
					"updating maintenance user@host to ",
					clues.Hide(currentOwner))

				if err := w.setMaintenanceParams(ctx, dw, params, currentOwner); err != nil {
					return clues.WrapWC(ctx, err, "updating maintenance parameters")
				}
			}

			ctx = clues.Add(ctx, "expected_owner", clues.Hide(params.Owner))

			logger.Ctx(ctx).Info("running kopia maintenance")

			err = snapshotmaintenance.Run(ctx, dw, mode, opts.Force, kopiaSafety)
			if err != nil {
				return clues.WrapWC(ctx, err, "running kopia maintenance")
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

func translateSafety(
	s repository.MaintenanceSafety,
) (maintenance.SafetyParameters, error) {
	switch s {
	case repository.FullMaintenanceSafety:
		return maintenance.SafetyFull, nil
	case repository.NoMaintenanceSafety:
		return maintenance.SafetyNone, nil
	default:
		return maintenance.SafetyParameters{}, clues.New("bad safety value").
			With("input_safety", s.String())
	}
}

func translateMode(t repository.MaintenanceType) (maintenance.Mode, error) {
	switch t {
	case repository.CompleteMaintenance:
		return maintenance.ModeFull, nil

	case repository.MetadataMaintenance:
		return maintenance.ModeQuick, nil

	default:
		return maintenance.ModeNone, clues.New("bad maintenance type").
			With("input_maintenance_type", t.String())
	}
}

// setMaintenanceUserHost sets the user and host for maintenance to the the
// user and host in the kopia config.
func (w Wrapper) setMaintenanceParams(
	ctx context.Context,
	drw repo.DirectRepositoryWriter,
	p *maintenance.Params,
	userAtHost string,
) error {
	// This will source user/host from the kopia config file or fallback to
	// fetching the values from the OS.
	p.Owner = userAtHost
	// Disable automatic maintenance for now since it can start matching on the
	// user/host of at least one machine now.
	p.QuickCycle.Enabled = false
	p.FullCycle.Enabled = false

	err := maintenance.SetParams(ctx, drw, p)
	if err != nil {
		return clues.Wrap(err, "setting maintenance user/host")
	}

	return nil
}

// SetRetentionParameters sets configuration values related to immutable backups
// and retention policies on the storage bucket. The minimum retention period
// must be >= 24hrs due to kopia default expectations about full maintenance.
func (w *Wrapper) SetRetentionParameters(
	ctx context.Context,
	retention repository.Retention,
) error {
	return clues.Stack(w.c.setRetentionParameters(ctx, retention)).OrNil()
}

func (w *Wrapper) UpdatePersistentConfig(
	ctx context.Context,
	config repository.PersistentConfig,
) error {
	return clues.Stack(w.c.updatePersistentConfig(ctx, config)).OrNil()
}
