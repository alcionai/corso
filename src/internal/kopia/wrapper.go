package kopia

import (
	"context"
	"path"

	"github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/logger"
)

const (
	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	corsoHost = "corso-host"
	corsoUser = "corso"
)

var (
	errNotConnected   = errors.New("not connected to repo")
	errUnsupportedDir = errors.New("unsupported static children in streaming directory")
)

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
	collection connector.DataCollection,
	details *backup.Details,
) func(context.Context, func(context.Context, fs.Entry) error) error {
	return func(ctx context.Context, cb func(context.Context, fs.Entry) error) error {
		items := collection.Items()
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case e, ok := <-items:
				if !ok {
					return nil
				}
				ei, ok := e.(connector.DataStreamInfo)
				if !ok {
					return errors.New("item does not implement DataStreamInfo")
				}

				entry := virtualfs.StreamingFileFromReader(e.UUID(), e.ToReader())
				if err := cb(ctx, entry); err != nil {
					return errors.Wrap(err, "executing callback")
				}

				// Populate BackupDetails
				ep := append(collection.FullPath(), e.UUID())
				details.Add(path.Join(ep...), ei.Info())
			}
		}
	}
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are either virtualfs.StreamingDirectory or
// virtualfs.staticDirectory.
func buildKopiaDirs(dirName string, dir *treeMap, details *backup.Details) (fs.Directory, error) {
	// Don't support directories that have both a DataCollection and a set of
	// static child directories.
	if dir.collection != nil && len(dir.childDirs) > 0 {
		return nil, errors.New(errUnsupportedDir.Error())
	}

	if dir.collection != nil {
		return virtualfs.NewStreamingDirectory(dirName, getStreamItemFunc(dir.collection, details)), nil
	}

	// Need to build the directory tree from the leaves up because intermediate
	// directories need to have all their entries at creation time.
	childDirs := []fs.Entry{}

	for childName, childDir := range dir.childDirs {
		child, err := buildKopiaDirs(childName, childDir, details)
		if err != nil {
			return nil, err
		}

		childDirs = append(childDirs, child)
	}

	return virtualfs.NewStaticDirectory(dirName, childDirs), nil
}

type treeMap struct {
	childDirs  map[string]*treeMap
	collection connector.DataCollection
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
func inflateDirTree(ctx context.Context, collections []connector.DataCollection, details *backup.Details) (fs.Directory, error) {
	roots := make(map[string]*treeMap)

	for _, s := range collections {
		path := s.FullPath()

		if len(path) == 0 {
			return nil, errors.New("no identifier for collection")
		}

		dir, ok := roots[path[0]]
		if !ok {
			dir = newTreeMap()
			roots[path[0]] = dir
		}

		// Single DataCollection with no ancestors.
		if len(path) == 1 {
			dir.collection = s
			continue
		}

		for _, p := range path[1 : len(path)-1] {
			newDir, ok := dir.childDirs[p]
			if !ok {
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

		end := len(path) - 1

		// Make sure this entry doesn't already exist.
		if _, ok := dir.childDirs[path[end]]; ok {
			return nil, errors.New(errUnsupportedDir.Error())
		}

		sd := newTreeMap()
		sd.collection = s
		dir.childDirs[path[end]] = sd
	}

	if len(roots) > 1 {
		return nil, errors.New("multiple root directories")
	}

	var res fs.Directory
	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(dirName, dir, details)
		if err != nil {
			return nil, err
		}

		res = tmp
	}

	return res, nil
}

func (w Wrapper) BackupCollections(
	ctx context.Context,
	collections []connector.DataCollection,
) (*BackupStats, *backup.Details, error) {
	if w.c == nil {
		return nil, nil, errNotConnected
	}

	details := &backup.Details{}

	dirTree, err := inflateDirTree(ctx, collections, details)
	if err != nil {
		return nil, nil, errors.Wrap(err, "building kopia directories")
	}

	stats, err := w.makeSnapshotWithRoot(ctx, dirTree, details)
	if err != nil {
		return nil, nil, err
	}

	return stats, details, nil
}

func (w Wrapper) makeSnapshotWithRoot(
	ctx context.Context,
	root fs.Directory,
	details *backup.Details,
) (*BackupStats, error) {
	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		// TODO(ashmrtnz): will this be something useful for snapshot lookups later?
		Path: root.Name(),
	}
	ctx, rw, err := w.c.NewWriter(ctx, repo.WriteSessionOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "get repo writer")
	}

	policyTree, err := policy.TreeForSource(ctx, w.c, si)
	if err != nil {
		return nil, errors.Wrap(err, "get policy tree")
	}

	u := snapshotfs.NewUploader(rw)

	man, err := u.Upload(ctx, root, policyTree, si)
	if err != nil {
		return nil, errors.Wrap(err, "uploading data")
	}

	if _, err := snapshot.SaveSnapshot(ctx, rw, man); err != nil {
		return nil, errors.Wrap(err, "saving snapshot")
	}

	if err := rw.Flush(ctx); err != nil {
		return nil, errors.Wrap(err, "flushing writer")
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

	manifest, err := snapshot.LoadSnapshot(ctx, w.c, manifest.ID(snapshotID))
	if err != nil {
		return nil, errors.Wrap(err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, manifest)
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
func (w Wrapper) CollectItems(
	ctx context.Context,
	snapshotID string,
	itemPath []string,
	isDirectory bool,
) ([]connector.DataCollection, error) {
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

	return []connector.DataCollection{c}, nil
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
) (connector.DataCollection, error) {
	c, err := w.CollectItems(ctx, snapshotID, itemPath, false)
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
) (connector.DataCollection, error) {
	r, err := f.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}

	return &kopiaDataCollection{
		streams: []connector.DataStream{
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
	files := []fs.File{}
	dirs := []fs.Directory{}
	var errs *multierror.Error

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
			logger.Ctx(ctx).Warnf("unexpected item of type %T; skipping", e)
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
) ([]connector.DataCollection, *multierror.Error) {
	collections := []connector.DataCollection{}
	// Want a local copy of relativePath with our new element.
	fullPath := append(append([]string{}, relativePath...), dir.Name())
	var errs *multierror.Error

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

		streams := make([]connector.DataStream, 0, len(files))

		for _, f := range files {
			r, err := f.Open(ctx)
			if err != nil {
				fileFullPath := path.Join(append(append([]string{}, fullPath...), f.Name())...)
				errs = multierror.Append(
					errs, errors.Wrapf(err, "getting reader for file %q", fileFullPath))
				logger.Ctx(ctx).Warnf("skipping file %q", fileFullPath)
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
) ([]connector.DataCollection, error) {
	return w.CollectItems(ctx, snapshotID, basePath, true)
}
