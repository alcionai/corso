package kopia

import (
	"context"
	"path"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/logger"
)

const (
	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	corsoHost = "corso-host"
	corsoUser = "corso"
)

var errNotConnected = errors.New("not connected to repo")

type BackupStats struct {
	SnapshotID          string
	TotalFileCount      int
	TotalDirectoryCount int
	IgnoredErrorCount   int
	ErrorCount          int
	Incomplete          bool
	IncompleteReason    string
}

func manifestToStats(man *snapshot.Manifest) BackupStats {
	return BackupStats{
		SnapshotID:          string(man.ID),
		TotalFileCount:      int(man.Stats.TotalFileCount),
		TotalDirectoryCount: int(man.Stats.TotalDirectoryCount),
		IgnoredErrorCount:   int(man.Stats.IgnoredErrorCount),
		ErrorCount:          int(man.Stats.ErrorCount),
		Incomplete:          man.IncompleteReason != "",
		IncompleteReason:    man.IncompleteReason,
	}
}

type itemDetails struct {
	info    details.ItemInfo
	repoRef string
}

type corsoProgress struct {
	snapshotfs.UploadProgress
	pending map[string]*itemDetails
	deets   *details.Details
	mu      sync.RWMutex
}

// Kopia interface function used as a callback when kopia finishes processing a
// file.
func (cp *corsoProgress) FinishedFile(relativePath string, err error) {
	// Pass the call through as well so we don't break expected functionality.
	defer cp.UploadProgress.FinishedFile(relativePath, err)
	// Whether it succeeded or failed, remove the entry from our pending set so we
	// don't leak references.
	defer func() {
		cp.mu.Lock()
		defer cp.mu.Unlock()

		delete(cp.pending, relativePath)
	}()

	if err != nil {
		return
	}

	d := cp.get(relativePath)
	if d == nil {
		return
	}

	cp.deets.Add(d.repoRef, d.info)
}

func (cp *corsoProgress) put(k string, v *itemDetails) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.pending[k] = v
}

func (cp *corsoProgress) get(k string) *itemDetails {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.pending[k]
}

func NewWrapper(c *conn) (*Wrapper, error) {
	if err := c.wrap(); err != nil {
		return nil, errors.Wrap(err, "creating Wrapper")
	}

	return &Wrapper{c}, nil
}

type Wrapper struct {
	c *conn
}

func (w *Wrapper) Close(ctx context.Context) error {
	if w.c == nil {
		return nil
	}

	err := w.c.Close(ctx)
	w.c = nil

	return errors.Wrap(err, "closing Wrapper")
}

// getStreamItemFunc returns a function that can be used by kopia's
// virtualfs.StreamingDirectory to iterate through directory entries and call
// kopia callbacks on directory entries. It binds the directory to the given
// DataCollection.
func getStreamItemFunc(
	staticEnts []fs.Entry,
	streamedEnts data.Collection,
	progress *corsoProgress,
) func(context.Context, func(context.Context, fs.Entry) error) error {
	return func(ctx context.Context, cb func(context.Context, fs.Entry) error) error {
		// Collect all errors and return them at the end so that iteration for this
		// directory doesn't end early.
		var errs *multierror.Error

		// Return static entries in this directory first.
		for _, d := range staticEnts {
			if err := cb(ctx, d); err != nil {
				return errors.Wrap(err, "executing callback on static directory")
			}
		}

		if streamedEnts == nil {
			return nil
		}

		items := streamedEnts.Items()

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case e, ok := <-items:
				if !ok {
					return errs.ErrorOrNil()
				}

				itemPath := path.Join(append(streamedEnts.FullPath(), e.UUID())...)

				ei, ok := e.(data.StreamInfo)
				if !ok {
					errs = multierror.Append(
						errs, errors.Errorf("item %q does not implement DataStreamInfo", itemPath))

					logger.Ctx(ctx).Errorw(
						"item does not implement DataStreamInfo; skipping", "path", itemPath)

					continue
				}

				// Relative path given to us in the callback is missing the root
				// element. Add to pending set before calling the callback to avoid race
				// conditions when the item is completed.
				p := path.Join(append(streamedEnts.FullPath()[1:], e.UUID())...)
				d := &itemDetails{info: ei.Info(), repoRef: itemPath}

				progress.put(p, d)

				entry := virtualfs.StreamingFileFromReader(e.UUID(), e.ToReader())
				if err := cb(ctx, entry); err != nil {
					// Kopia's uploader swallows errors in most cases, so if we see
					// something here it's probably a big issue and we should return.
					errs = multierror.Append(errs, errors.Wrapf(err, "executing callback on %q", itemPath))
					return errs.ErrorOrNil()
				}
			}
		}
	}
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are virtualfs.StreamingDirectory.
func buildKopiaDirs(dirName string, dir *treeMap, progress *corsoProgress) (fs.Directory, error) {
	// Need to build the directory tree from the leaves up because intermediate
	// directories need to have all their entries at creation time.
	var childDirs []fs.Entry

	for childName, childDir := range dir.childDirs {
		child, err := buildKopiaDirs(childName, childDir, progress)
		if err != nil {
			return nil, err
		}

		childDirs = append(childDirs, child)
	}

	return virtualfs.NewStreamingDirectory(
		dirName,
		getStreamItemFunc(childDirs, dir.collection, progress),
	), nil
}

type treeMap struct {
	childDirs  map[string]*treeMap
	collection data.Collection
}

func newTreeMap() *treeMap {
	return &treeMap{
		childDirs: map[string]*treeMap{},
	}
}

// inflateDirTree returns an fs.Directory tree rooted at the oldest common
// ancestor of the streams and uses virtualfs.StaticDirectory for internal nodes
// in the hierarchy. Leaf nodes are virtualfs.StreamingDirectory with the given
// DataCollections.
func inflateDirTree(
	ctx context.Context,
	collections []data.Collection,
	progress *corsoProgress,
) (fs.Directory, error) {
	roots := make(map[string]*treeMap)

	for _, s := range collections {
		itemPath := s.FullPath()

		if len(itemPath) == 0 {
			return nil, errors.New("no identifier for collection")
		}

		dir, ok := roots[itemPath[0]]
		if !ok {
			dir = newTreeMap()
			roots[itemPath[0]] = dir
		}

		// Single DataCollection with no ancestors.
		if len(itemPath) == 1 {
			dir.collection = s
			continue
		}

		for _, p := range itemPath[1 : len(itemPath)-1] {
			newDir := dir.childDirs[p]
			if newDir == nil {
				newDir = newTreeMap()

				if dir.childDirs == nil {
					dir.childDirs = map[string]*treeMap{}
				}

				dir.childDirs[p] = newDir
			}

			dir = newDir
		}

		// At this point we have all the ancestor directories of this DataCollection
		// as treeMap objects and `dir` is the parent directory of this
		// DataCollection.

		end := len(itemPath) - 1

		// Make sure this entry doesn't already exist.
		tmpDir := dir.childDirs[itemPath[end]]
		if tmpDir == nil {
			tmpDir = newTreeMap()
			dir.childDirs[itemPath[end]] = tmpDir
		}

		tmpDir.collection = s
	}

	if len(roots) > 1 {
		return nil, errors.New("multiple root directories")
	}

	var res fs.Directory

	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(dirName, dir, progress)
		if err != nil {
			return nil, err
		}

		res = tmp
	}

	return res, nil
}

func (w Wrapper) BackupCollections(
	ctx context.Context,
	collections []data.Collection,
) (*BackupStats, *details.Details, error) {
	if w.c == nil {
		return nil, nil, errNotConnected
	}

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		deets:   &details.Details{},
	}

	dirTree, err := inflateDirTree(ctx, collections, progress)
	if err != nil {
		return nil, nil, errors.Wrap(err, "building kopia directories")
	}

	stats, err := w.makeSnapshotWithRoot(ctx, dirTree, progress)
	if err != nil {
		return nil, nil, err
	}

	return stats, progress.deets, nil
}

func (w Wrapper) makeSnapshotWithRoot(
	ctx context.Context,
	root fs.Directory,
	progress *corsoProgress,
) (*BackupStats, error) {
	var man *snapshot.Manifest

	err := repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{
			Purpose: "KopiaWrapperBackup",
			// Always flush so we don't leak write sessions. Still uses reachability
			// for consistency.
			FlushOnFailure: true,
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
				err = errors.Wrap(err, "get policy tree")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			// By default Uploader is best-attempt.
			u := snapshotfs.NewUploader(rw)
			progress.UploadProgress = u.Progress
			u.Progress = progress

			man, err = u.Upload(innerCtx, root, policyTree, si)
			if err != nil {
				err = errors.Wrap(err, "uploading data")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			if _, err := snapshot.SaveSnapshot(innerCtx, rw, man); err != nil {
				err = errors.Wrap(err, "saving snapshot")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return nil, errors.Wrap(err, "kopia backup")
	}

	res := manifestToStats(man)

	return &res, nil
}

// getEntry returns the item that the restore operation is rooted at. For
// single-item restores, this is the kopia file the data is sourced from. For
// restores of directories or subtrees it is the directory at the root of the
// subtree.
func (w Wrapper) getEntry(
	ctx context.Context,
	snapshotID string,
	itemPath []string,
) (fs.Entry, error) {
	if len(itemPath) == 0 {
		return nil, errors.New("no restore path given")
	}

	man, err := snapshot.LoadSnapshot(ctx, w.c, manifest.ID(snapshotID))
	if err != nil {
		return nil, errors.Wrap(err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, man)
	if err != nil {
		return nil, errors.Wrap(err, "getting root directory")
	}

	// GetNestedEntry handles nil properly.
	e, err := snapshotfs.GetNestedEntry(ctx, rootDirEntry, itemPath[1:])
	if err != nil {
		return nil, errors.Wrap(err, "getting nested object handle")
	}

	return e, nil
}

// CollectItems pulls data from kopia for the given items in the snapshot with
// ID snapshotID. If isDirectory is true, it returns a slice of DataCollections
// with data from directories in the subtree rooted at itemPath. If isDirectory
// is false it returns a DataCollection (in a slice) with a single item for each
// requested item. If the item does not exist or a file is found when a directory
// is expected (or the opposite) it returns an error.
func (w Wrapper) collectItems(
	ctx context.Context,
	snapshotID string,
	itemPath []string,
	isDirectory bool,
) ([]data.Collection, error) {
	e, err := w.getEntry(ctx, snapshotID, itemPath)
	if err != nil {
		return nil, err
	}

	// The paths passed below is the path up to (but not including) the
	// file/directory passed.
	if isDirectory {
		dir, ok := e.(fs.Directory)
		if !ok {
			return nil, errors.New("requested object is not a directory")
		}

		c, err := restoreSubtree(ctx, dir, itemPath[:len(itemPath)-1])
		// For some reason tests error out if the multierror is nil but we don't
		// call ErrorOrNil.
		return c, err.ErrorOrNil()
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, errors.New("requested object is not a file")
	}

	c, err := restoreSingleItem(ctx, f, itemPath[:len(itemPath)-1])
	if err != nil {
		return nil, err
	}

	return []data.Collection{c}, nil
}

// RestoreSingleItem looks up the item at the given path in the snapshot with id
// snapshotID. The path should be the full path of the item from the root.
// If the item is a file in kopia then it returns a DataCollection with the item
// as its sole element and DataCollection.FullPath() set to
// split(dirname(itemPath), "/"). If the item does not exist in kopia or is not
// a file an error is returned. The UUID of the returned DataStreams will be the
// name of the kopia file the data is sourced from.
func (w Wrapper) RestoreSingleItem(
	ctx context.Context,
	snapshotID string,
	itemPath []string,
) (data.Collection, error) {
	c, err := w.collectItems(ctx, snapshotID, itemPath, false)
	if err != nil {
		return nil, err
	}

	return c[0], nil
}

// restoreSingleItem looks up the item at the given path starting from rootDir
// where rootDir is the root of a snapshot. If the item is a file in kopia then
// it returns a DataCollection with the item as its sole element and
// DataCollection.FullPath() set to split(dirname(itemPath), "/"). If the item
// does not exist in kopia or is not a file an error is returned. The UUID of
// the returned DataStreams will be the name of the kopia file the data is
// sourced from.
func restoreSingleItem(
	ctx context.Context,
	f fs.File,
	itemPath []string,
) (data.Collection, error) {
	r, err := f.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}

	return &kopiaDataCollection{
		streams: []data.Stream{
			&kopiaDataStream{
				uuid:   f.Name(),
				reader: r,
			},
		},
		path: itemPath,
	}, nil
}

func walkDirectory(
	ctx context.Context,
	dir fs.Directory,
) ([]fs.File, []fs.Directory, *multierror.Error) {
	var errs *multierror.Error

	files := []fs.File{}
	dirs := []fs.Directory{}

	err := dir.IterateEntries(ctx, func(innerCtx context.Context, e fs.Entry) error {
		// Early exit on context cancel.
		if err := innerCtx.Err(); err != nil {
			return err
		}

		switch e := e.(type) {
		case fs.Directory:
			dirs = append(dirs, e)
		case fs.File:
			files = append(files, e)
		default:
			errs = multierror.Append(errs, errors.Errorf("unexpected item type %T", e))
			logger.Ctx(ctx).Errorw(
				"unexpected item type; skipping", "type", e)
		}

		return nil
	})
	if err != nil {
		// If the iterator itself had an error add it to the list.
		errs = multierror.Append(errs, errors.Wrap(err, "getting directory data"))
	}

	return files, dirs, errs
}

// restoreSubtree returns DataCollections for each subdirectory (or the
// directory itself) that contains files. The FullPath of each returned
// DataCollection is the path from the root of the kopia directory structure to
// the directory. The UUID of each DataStream in each DataCollection is the name
// of the kopia file the data is sourced from.
func restoreSubtree(
	ctx context.Context,
	dir fs.Directory,
	relativePath []string,
) ([]data.Collection, *multierror.Error) {
	var errs *multierror.Error

	collections := []data.Collection{}
	// Want a local copy of relativePath with our new element.
	fullPath := append(append([]string{}, relativePath...), dir.Name())

	files, dirs, err := walkDirectory(ctx, dir)
	if err != nil {
		errs = multierror.Append(
			errs, errors.Wrapf(err, "walking directory %q", path.Join(fullPath...)))
	}

	if len(files) > 0 {
		if ctxErr := ctx.Err(); ctxErr != nil {
			errs = multierror.Append(errs, errors.WithStack(ctxErr))
			return nil, errs
		}

		streams := make([]data.Stream, 0, len(files))

		for _, f := range files {
			r, err := f.Open(ctx)
			if err != nil {
				fileFullPath := path.Join(append(append([]string{}, fullPath...), f.Name())...)
				errs = multierror.Append(
					errs, errors.Wrapf(err, "getting reader for file %q", fileFullPath))

				logger.Ctx(ctx).Errorw(
					"unable to get file reader; skipping", "path", fileFullPath)

				continue
			}

			streams = append(streams, &kopiaDataStream{
				reader: r,
				uuid:   f.Name(),
			})
		}

		collections = append(collections, &kopiaDataCollection{
			streams: streams,
			path:    fullPath,
		})
	}

	for _, d := range dirs {
		if ctxErr := ctx.Err(); ctxErr != nil {
			errs = multierror.Append(errs, errors.WithStack(ctxErr))
			return nil, errs
		}

		c, err := restoreSubtree(ctx, d, fullPath)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(
				err,
				"traversing subdirectory %q",
				path.Join(append(append([]string{}, fullPath...), d.Name())...),
			))
		}

		collections = append(collections, c...)
	}

	return collections, errs
}

func (w Wrapper) RestoreDirectory(
	ctx context.Context,
	snapshotID string,
	basePath []string,
) ([]data.Collection, error) {
	return w.collectItems(ctx, snapshotID, basePath, true)
}

// RestoreSingleItem looks up all paths- assuming each is an item declaration,
// not a directory- in the snapshot with id snapshotID. The path should be the
// full path of the item from the root.  Returns the results as a slice of single-
// item DataCollections, where the DataCollection.FullPath() matches the path.
// If the item does not exist in kopia or is not a file an error is returned.
// The UUID of the returned DataStreams will be the name of the kopia file the
// data is sourced from.
func (w Wrapper) RestoreMultipleItems(
	ctx context.Context,
	snapshotID string,
	paths [][]string,
) ([]data.Collection, error) {
	var (
		dcs  = []data.Collection{}
		errs *multierror.Error
	)

	for _, itemPath := range paths {
		dc, err := w.RestoreSingleItem(ctx, snapshotID, itemPath)
		if err != nil {
			errs = multierror.Append(errs, err)
		} else {
			dcs = append(dcs, dc)
		}
	}

	return dcs, errs.ErrorOrNil()
}

// DeleteSnapshot removes the provided manifest from kopia.
func (w Wrapper) DeleteSnapshot(
	ctx context.Context,
	snapshotID string,
) error {
	mid := manifest.ID(snapshotID)

	if len(mid) == 0 {
		return errors.New("attempt to delete unidentified snapshot")
	}

	err := repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{Purpose: "KopiaWrapperBackupDeletion"},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			if err := rw.DeleteManifest(ctx, mid); err != nil {
				return errors.Wrap(err, "deleting snapshot")
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return errors.Wrap(err, "kopia deleting backup manifest")
	}

	return nil
}
