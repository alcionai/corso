package kopia

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	corsoHost = "corso-host"
	corsoUser = "corso"
)

var (
	errNotConnected  = errors.New("not connected to repo")
	errNoRestorePath = errors.New("no restore path given")
)

type BackupStats struct {
	SnapshotID          string
	TotalFileCount      int
	TotalHashedBytes    int64
	TotalDirectoryCount int
	IgnoredErrorCount   int
	ErrorCount          int
	Incomplete          bool
	IncompleteReason    string
}

func manifestToStats(man *snapshot.Manifest, progress *corsoProgress) BackupStats {
	return BackupStats{
		SnapshotID:          string(man.ID),
		TotalFileCount:      int(man.Stats.TotalFileCount),
		TotalHashedBytes:    progress.totalBytes,
		TotalDirectoryCount: int(man.Stats.TotalDirectoryCount),
		IgnoredErrorCount:   int(man.Stats.IgnoredErrorCount),
		ErrorCount:          int(man.Stats.ErrorCount),
		Incomplete:          man.IncompleteReason != "",
		IncompleteReason:    man.IncompleteReason,
	}
}

type itemDetails struct {
	info     details.ItemInfo
	repoPath path.Path
}

type corsoProgress struct {
	snapshotfs.UploadProgress
	pending    map[string]*itemDetails
	deets      *details.Details
	mu         sync.RWMutex
	totalBytes int64
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

	parent := d.repoPath.ToBuilder().Dir()

	cp.deets.Add(
		d.repoPath.String(),
		d.repoPath.ShortRef(),
		parent.ShortRef(),
		d.info,
	)

	folders := []details.FolderEntry{}

	for len(parent.Elements()) > 0 {
		nextParent := parent.Dir()

		folders = append(folders, details.FolderEntry{
			RepoRef:   parent.String(),
			ShortRef:  parent.ShortRef(),
			ParentRef: nextParent.ShortRef(),
			Info: details.ItemInfo{
				Folder: &details.FolderInfo{
					DisplayName: parent.Elements()[len(parent.Elements())-1],
				},
			},
		})

		parent = nextParent
	}

	cp.deets.AddFolders(folders)
}

// Kopia interface function used as a callback when kopia finishes hashing a file.
func (cp *corsoProgress) FinishedHashingFile(fname string, bytes int64) {
	// Pass the call through as well so we don't break expected functionality.
	defer cp.UploadProgress.FinishedHashingFile(fname, bytes)

	atomic.AddInt64(&cp.totalBytes, bytes)
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

				// For now assuming that item IDs don't need escaping.
				itemPath, err := streamedEnts.FullPath().Append(e.UUID(), true)
				if err != nil {
					err = errors.Wrap(err, "getting full item path")
					errs = multierror.Append(errs, err)

					logger.Ctx(ctx).Error(err)

					continue
				}

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
				d := &itemDetails{info: ei.Info(), repoPath: itemPath}
				progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)

				entry := virtualfs.StreamingFileFromReader(encodeAsPath(e.UUID()), e.ToReader())
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
		encodeAsPath(dirName),
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
		if s.FullPath() == nil {
			return nil, errors.New("no identifier for collection")
		}

		itemPath := s.FullPath().Elements()

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

	res := manifestToStats(man, progress)

	return &res, nil
}

func (w Wrapper) getSnapshotRoot(
	ctx context.Context,
	snapshotID string,
) (fs.Entry, error) {
	man, err := snapshot.LoadSnapshot(ctx, w.c, manifest.ID(snapshotID))
	if err != nil {
		return nil, errors.Wrap(err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, man)

	return rootDirEntry, errors.Wrap(err, "getting root directory")
}

// getItemStream looks up the item at the given path starting from snapshotRoot.
// If the item is a file in kopia then it returns a data.Stream of the item. If
// the item does not exist in kopia or is not a file an error is returned. The
// UUID of the returned data.Stream will be the name of the kopia file the data
// is sourced from.
func getItemStream(
	ctx context.Context,
	itemPath path.Path,
	snapshotRoot fs.Entry,
) (data.Stream, error) {
	if itemPath == nil {
		return nil, errors.WithStack(errNoRestorePath)
	}

	// GetNestedEntry handles nil properly.
	e, err := snapshotfs.GetNestedEntry(
		ctx,
		snapshotRoot,
		encodeElements(itemPath.PopFront().Elements()...),
	)
	if err != nil {
		return nil, errors.Wrap(err, "getting nested object handle")
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, errors.New("requested object is not a file")
	}

	r, err := f.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}

	decodedName, err := decodeElement(f.Name())
	if err != nil {
		return nil, errors.Wrap(err, "decoding file name")
	}

	return &kopiaDataStream{
		uuid:   decodedName,
		reader: r,
	}, nil
}

// RestoreMultipleItems looks up all paths- assuming each is an item declaration,
// not a directory- in the snapshot with id snapshotID. The path should be the
// full path of the item from the root.  Returns the results as a slice of single-
// item DataCollections, where the DataCollection.FullPath() matches the path.
// If the item does not exist in kopia or is not a file an error is returned.
// The UUID of the returned DataStreams will be the name of the kopia file the
// data is sourced from.
func (w Wrapper) RestoreMultipleItems(
	ctx context.Context,
	snapshotID string,
	paths []path.Path,
) ([]data.Collection, error) {
	if len(paths) == 0 {
		return nil, errors.WithStack(errNoRestorePath)
	}

	snapshotRoot, err := w.getSnapshotRoot(ctx, snapshotID)
	if err != nil {
		return nil, err
	}

	var (
		errs *multierror.Error
		// Maps short ID of parent path to data collection for that folder.
		cols = map[string]*kopiaDataCollection{}
	)

	for _, itemPath := range paths {
		ds, err := getItemStream(ctx, itemPath, snapshotRoot)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		parentPath, err := itemPath.Dir()
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "making directory collection"))
			continue
		}

		c, ok := cols[parentPath.ShortRef()]
		if !ok {
			cols[parentPath.ShortRef()] = &kopiaDataCollection{path: parentPath}
			c = cols[parentPath.ShortRef()]
		}

		c.streams = append(c.streams, ds)
	}

	res := make([]data.Collection, 0, len(cols))
	for _, c := range cols {
		res = append(res, c)
	}

	return res, errs.ErrorOrNil()
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
