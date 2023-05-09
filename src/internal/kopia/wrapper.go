package kopia

import (
	"context"
	"errors"
	"strings"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/maintenance"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/kopia/kopia/snapshot/snapshotmaintenance"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	corsoHost = "corso-host"
	corsoUser = "corso"

	serializationVersion uint32 = 1
)

// common manifest tags
const (
	TagBackupID       = "backup-id"
	TagBackupCategory = "is-canon-backup"
)

var (
	errNotConnected  = clues.New("not connected to repo")
	errNoRestorePath = clues.New("no restore path given")
)

type BackupStats struct {
	SnapshotID string

	TotalHashedBytes   int64
	TotalUploadedBytes int64

	TotalFileCount      int
	CachedFileCount     int
	UncachedFileCount   int
	TotalDirectoryCount int
	ErrorCount          int

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
		return clues.Wrap(err, "closing Wrapper").WithClues(ctx)
	}

	return nil
}

type IncrementalBase struct {
	*snapshot.Manifest
	SubtreePaths []*path.Builder
}

// ConsumeBackupCollections takes a set of collections and creates a kopia snapshot
// with the data that they contain. previousSnapshots is used for incremental
// backups and should represent the base snapshot from which metadata is sourced
// from as well as any incomplete snapshot checkpoints that may contain more
// recent data than the base snapshot. The absence of previousSnapshots causes a
// complete backup of all data.
func (w Wrapper) ConsumeBackupCollections(
	ctx context.Context,
	previousSnapshots []IncrementalBase,
	collections []data.BackupCollection,
	globalExcludeSet prefixmatcher.StringSetReader,
	tags map[string]string,
	buildTreeWithBase bool,
	errs *fault.Bus,
) (*BackupStats, *details.Builder, DetailsMergeInfoer, error) {
	if w.c == nil {
		return nil, nil, nil, clues.Stack(errNotConnected).WithClues(ctx)
	}

	ctx, end := diagnostics.Span(ctx, "kopia:consumeBackupCollections")
	defer end()

	if len(collections) == 0 && (globalExcludeSet == nil || globalExcludeSet.Empty()) {
		return &BackupStats{}, &details.Builder{}, nil, nil
	}

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		deets:   &details.Builder{},
		toMerge: newMergeDetails(),
		errs:    errs,
	}

	// When running an incremental backup, we need to pass the prior
	// snapshot bases into inflateDirTree so that the new snapshot
	// includes historical data.
	var base []IncrementalBase
	if buildTreeWithBase {
		base = previousSnapshots
	}

	dirTree, err := inflateDirTree(
		ctx,
		w.c,
		base,
		collections,
		globalExcludeSet,
		progress)
	if err != nil {
		return nil, nil, nil, clues.Wrap(err, "building kopia directories")
	}

	s, err := w.makeSnapshotWithRoot(
		ctx,
		previousSnapshots,
		dirTree,
		tags,
		progress)
	if err != nil {
		return nil, nil, nil, err
	}

	return s, progress.deets, progress.toMerge, progress.errs.Failure()
}

func (w Wrapper) makeSnapshotWithRoot(
	ctx context.Context,
	prevSnapEntries []IncrementalBase,
	root fs.Directory,
	addlTags map[string]string,
	progress *corsoProgress,
) (*BackupStats, error) {
	var (
		man *snapshot.Manifest
		bc  = &stats.ByteCounter{}
	)

	snapIDs := make([]manifest.ID, 0, len(prevSnapEntries)) // just for logging
	prevSnaps := make([]*snapshot.Manifest, 0, len(prevSnapEntries))

	for _, ent := range prevSnapEntries {
		prevSnaps = append(prevSnaps, ent.Manifest)
		snapIDs = append(snapIDs, ent.ID)
	}

	ctx = clues.Add(
		ctx,
		"len_prev_base_snapshots", len(prevSnapEntries),
		"assist_snap_ids", snapIDs,
		"additional_tags", addlTags)

	if len(snapIDs) > 0 {
		logger.Ctx(ctx).Info("using snapshots for kopia-assisted incrementals")
	} else {
		logger.Ctx(ctx).Info("no base snapshots for kopia-assisted incrementals")
	}

	tags := map[string]string{}

	for k, v := range addlTags {
		mk, mv := makeTagKV(k)

		if len(v) == 0 {
			v = mv
		}

		tags[mk] = v
	}

	err := repo.WriteSession(
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
				Host:     corsoHost,
				UserName: corsoUser,
				// TODO(ashmrtnz): will this be something useful for snapshot lookups later?
				Path: root.Name(),
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
				err = clues.Wrap(err, "get policy tree").WithClues(ctx)
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
				err = clues.Wrap(err, "uploading data").WithClues(ctx)
				logger.CtxErr(innerCtx, err).Error("uploading kopia backup")

				return err
			}

			man.Tags = tags

			if _, err := snapshot.SaveSnapshot(innerCtx, rw, man); err != nil {
				err = clues.Wrap(err, "saving snapshot").WithClues(ctx)
				logger.CtxErr(innerCtx, err).Error("persisting kopia backup snapshot")

				return err
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return nil, clues.Wrap(err, "kopia backup")
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
		return nil, clues.Wrap(err, "getting snapshot handle").WithClues(ctx)
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, man)
	if err != nil {
		return nil, clues.Wrap(err, "getting root directory").WithClues(ctx)
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
		return nil, clues.Wrap(errNoRestorePath, "getting directory").WithClues(ctx)
	}

	// GetNestedEntry handles nil properly.
	e, err := snapshotfs.GetNestedEntry(
		ctx,
		snapshotRoot,
		encodeElements(dirPath.PopFront().Elements()...))
	if err != nil {
		if isErrEntryNotFound(err) {
			err = clues.Stack(data.ErrNotFound, err)
		}

		return nil, clues.Wrap(err, "getting nested object handle").WithClues(ctx)
	}

	f, ok := e.(fs.Directory)
	if !ok {
		return nil, clues.New("requested object is not a directory").WithClues(ctx)
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
		el        = bus.Local()
		res       = make([]data.RestoreCollection, 0, len(toLoad))
		loadCount = 0
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
				el.AddRecoverable(clues.Wrap(err, "loading storage directory").
					WithClues(ictx).
					Label(fault.LabelForceNoBackupCreation))

				continue
			}

			dc := &kopiaDataCollection{
				path:            col.restorePath,
				dir:             dir,
				counter:         bcounter,
				expectedVersion: serializationVersion,
			}

			if err := mergeCol.addCollection(dirItems.dir.String(), dc); err != nil {
				el.AddRecoverable(clues.Wrap(err, "adding collection to merge collection").
					WithClues(ctx).
					Label(fault.LabelForceNoBackupCreation))

				continue
			}

			for _, item := range dirItems.items {
				if el.Failure() != nil {
					return nil, el.Failure()
				}

				err := dc.addStream(ictx, item)
				if err != nil {
					el.AddRecoverable(clues.Wrap(err, "loading item").
						WithClues(ictx).
						Label(fault.LabelForceNoBackupCreation))

					continue
				}

				loadCount++
				if loadCount%1000 == 0 {
					logger.Ctx(ctx).Infow(
						"loading items from kopia",
						"loaded_items", loadCount)
				}
			}
		}
	}

	logger.Ctx(ctx).Infow(
		"done loading items from kopia",
		"loaded_items", loadCount)

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
		return nil, clues.Stack(errNoRestorePath).WithClues(ctx)
	}

	// Used later on, but less confusing to follow error propagation if we just
	// load it here.
	snapshotRoot, err := w.getSnapshotRoot(ctx, snapshotID)
	if err != nil {
		return nil, clues.Wrap(err, "loading snapshot root")
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
			"item_path", itemPaths.StoragePath.String(),
			"restore_path", itemPaths.RestorePath.String())

		parentStoragePath, err := itemPaths.StoragePath.Dir()
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "getting storage directory path").
				WithClues(ictx).
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
		return nil, clues.Wrap(err, "loading items")
	}

	return res, el.Failure()
}

// DeleteSnapshot removes the provided manifest from kopia.
func (w Wrapper) DeleteSnapshot(
	ctx context.Context,
	snapshotID string,
) error {
	mid := manifest.ID(snapshotID)
	if len(mid) == 0 {
		return clues.New("snapshot ID required for deletion").WithClues(ctx)
	}

	err := repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{Purpose: "KopiaWrapperBackupDeletion"},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			if err := rw.DeleteManifest(ctx, mid); err != nil {
				return clues.Wrap(err, "deleting snapshot").WithClues(ctx)
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return clues.Wrap(err, "deleting backup manifest").WithClues(ctx)
	}

	return nil
}

// FetchPrevSnapshotManifests returns a set of manifests for complete and maybe
// incomplete snapshots for the given (resource owner, service, category)
// tuples. Up to two manifests can be returned per tuple: one complete and one
// incomplete. An incomplete manifest may be returned if it is newer than the
// newest complete manifest for the tuple. Manifests are deduped such that if
// multiple tuples match the same manifest it will only be returned once.
// If tags are provided, manifests must include a superset of the k:v pairs
// specified by those tags.  Tags should pass their raw values, and will be
// normalized inside the func using MakeTagKV.
func (w Wrapper) FetchPrevSnapshotManifests(
	ctx context.Context,
	reasons []Reason,
	tags map[string]string,
) ([]*ManifestEntry, error) {
	if w.c == nil {
		return nil, clues.Stack(errNotConnected).WithClues(ctx)
	}

	return fetchPrevSnapshotManifests(ctx, w.c, reasons, tags), nil
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
	opts repository.Maintenance,
) error {
	kopiaSafety, err := translateSafety(opts.Safety)
	if err != nil {
		return clues.Wrap(err, "identifying safety level")
	}

	mode, err := translateMode(opts.Type)
	if err != nil {
		return clues.Wrap(err, "identifying maintenance mode")
	}

	currentOwner := w.c.ClientOptions().UsernameAtHost()

	ctx = clues.Add(
		ctx,
		"kopia_safety", kopiaSafety,
		"kopia_maintenance_mode", mode,
		"force", opts.Force,
		"current_local_owner", clues.Hide(currentOwner))

	dr, ok := w.c.Repository.(repo.DirectRepository)
	if !ok {
		return clues.New("unable to get valid handle to repo").WithClues(ctx)
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
				return clues.Wrap(err, "getting maintenance user@host").WithClues(ctx)
			}

			// Need to do some fixup here as the user/host may not have been set.
			if len(params.Owner) == 0 || (params.Owner != currentOwner && opts.Force) {
				observe.Message(
					ctx,
					"updating maintenance user@host to ",
					clues.Hide(currentOwner))

				if err := w.setMaintenanceParams(ctx, dw, params, currentOwner); err != nil {
					return clues.Wrap(err, "updating maintenance parameters").
						WithClues(ctx)
				}
			}

			ctx = clues.Add(ctx, "expected_owner", clues.Hide(params.Owner))

			logger.Ctx(ctx).Info("running kopia maintenance")

			err = snapshotmaintenance.Run(ctx, dw, mode, opts.Force, kopiaSafety)
			if err != nil {
				return clues.Wrap(err, "running kopia maintenance").WithClues(ctx)
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
