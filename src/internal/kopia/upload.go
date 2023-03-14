package kopia

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/metadata"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const maxInflateTraversalDepth = 500

var versionSize = int(unsafe.Sizeof(serializationVersion))

func newBackupStreamReader(version uint32, reader io.ReadCloser) *backupStreamReader {
	buf := make([]byte, versionSize)
	binary.BigEndian.PutUint32(buf, version)
	bufReader := io.NopCloser(bytes.NewReader(buf))

	return &backupStreamReader{
		readers:  []io.ReadCloser{bufReader, reader},
		combined: io.NopCloser(io.MultiReader(bufReader, reader)),
	}
}

// backupStreamReader is a wrapper around the io.Reader that other Corso
// components return when backing up information. It injects a version number at
// the start of the data stream. Future versions of Corso may not need this if
// they use more complex serialization logic as serialization/version injection
// will be handled by other components.
type backupStreamReader struct {
	readers  []io.ReadCloser
	combined io.ReadCloser
}

func (rw *backupStreamReader) Read(p []byte) (n int, err error) {
	if rw.combined == nil {
		return 0, os.ErrClosed
	}

	return rw.combined.Read(p)
}

func (rw *backupStreamReader) Close() error {
	if rw.combined == nil {
		return nil
	}

	rw.combined = nil

	for _, r := range rw.readers {
		r.Close()
	}

	return nil
}

// restoreStreamReader is a wrapper around the io.Reader that kopia returns when
// reading data from an item. It examines and strips off the version number of
// the restored data. Future versions of Corso may not need this if they use
// more complex serialization logic as version checking/deserialization will be
// handled by other components. A reader that returns a version error is no
// longer valid and should not be used once the version error is returned.
type restoreStreamReader struct {
	io.ReadCloser
	expectedVersion uint32
	readVersion     bool
}

func (rw *restoreStreamReader) checkVersion() error {
	versionBuf := make([]byte, versionSize)

	for newlyRead := 0; newlyRead < versionSize; {
		n, err := rw.ReadCloser.Read(versionBuf[newlyRead:])
		if err != nil {
			return errors.Wrap(err, "reading data format version")
		}

		newlyRead += n
	}

	version := binary.BigEndian.Uint32(versionBuf)

	if version != rw.expectedVersion {
		return errors.Errorf("unexpected data format %v", version)
	}

	return nil
}

func (rw *restoreStreamReader) Read(p []byte) (n int, err error) {
	if !rw.readVersion {
		rw.readVersion = true

		if err := rw.checkVersion(); err != nil {
			return 0, err
		}
	}

	return rw.ReadCloser.Read(p)
}

type itemDetails struct {
	info         *details.ItemInfo
	repoPath     path.Path
	prevPath     path.Path
	locationPath path.Path
	cached       bool
}

type corsoProgress struct {
	snapshotfs.UploadProgress
	pending map[string]*itemDetails
	deets   *details.Builder
	// toMerge represents items that we don't have in-memory item info for. The
	// item info for these items should be sourced from a base snapshot later on.
	toMerge    map[string]PrevRefs
	mu         sync.RWMutex
	totalBytes int64
	errs       *fault.Bus
	// expectedIgnoredErrors is a count of error cases caught in the Error wrapper
	// which are well known and actually ignorable.  At the end of a run, if the
	// manifest ignored error count is equal to this count, then everything is good.
	expectedIgnoredErrors int
}

// mutexted wrapper around expectedIgnoredErrors++
func (cp *corsoProgress) incExpectedErrs() {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.expectedIgnoredErrors++
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

	// These items were sourced from a base snapshot or were cached in kopia so we
	// never had to materialize their details in-memory.
	if d.info == nil {
		if d.prevPath == nil {
			cp.errs.AddRecoverable(clues.New("item sourced from previous backup with no previous path").
				With(
					"service", d.repoPath.Service().String(),
					"category", d.repoPath.Category().String(),
				).
				Label(fault.LabelForceNoBackupCreation))

			return
		}

		cp.mu.Lock()
		defer cp.mu.Unlock()

		cp.toMerge[d.prevPath.ShortRef()] = PrevRefs{
			Repo:     d.repoPath,
			Location: d.locationPath,
		}

		return
	}

	var (
		locationFolders string
		locPB           *path.Builder
		parent          = d.repoPath.ToBuilder().Dir()
	)

	if d.locationPath != nil {
		locationFolders = d.locationPath.Folder(true)

		locPB = d.locationPath.ToBuilder()

		// folderEntriesForPath assumes the location will
		// not have an item element appended
		if len(d.locationPath.Item()) > 0 {
			locPB = locPB.Dir()
		}
	}

	err = cp.deets.Add(
		d.repoPath.String(),
		d.repoPath.ShortRef(),
		parent.ShortRef(),
		locationFolders,
		!d.cached,
		*d.info)
	if err != nil {
		cp.errs.AddRecoverable(clues.New("adding item to details").
			With(
				"service", d.repoPath.Service().String(),
				"category", d.repoPath.Category().String(),
			).
			Label(fault.LabelForceNoBackupCreation))

		return
	}

	folders := details.FolderEntriesForPath(parent, locPB)
	cp.deets.AddFoldersForItem(
		folders,
		*d.info,
		!d.cached)
}

// Kopia interface function used as a callback when kopia finishes hashing a file.
func (cp *corsoProgress) FinishedHashingFile(fname string, bs int64) {
	// Pass the call through as well so we don't break expected functionality.
	defer cp.UploadProgress.FinishedHashingFile(fname, bs)

	sl := strings.Split(fname, "/")

	for i := range sl {
		rdt, err := base64.StdEncoding.DecodeString(sl[i])
		if err != nil {
			fmt.Println("f did not decode")
		}

		sl[i] = string(rdt)
	}

	logger.Ctx(context.Background()).Debugw("finished hashing file", "path", sl[2:])

	atomic.AddInt64(&cp.totalBytes, bs)
}

// Kopia interface function used as a callback when kopia detects a previously
// uploaded file that matches the current file and skips uploading the new
// (duplicate) version.
func (cp *corsoProgress) CachedFile(fname string, size int64) {
	defer cp.UploadProgress.CachedFile(fname, size)

	d := cp.get(fname)
	if d == nil {
		return
	}

	d.cached = true
}

// Kopia interface function used as a callback when kopia encounters an error
// during the upload process. This could be from reading a file or something
// else.
func (cp *corsoProgress) Error(relpath string, err error, isIgnored bool) {
	// The malware case is an artifact of being unable to skip the item
	// if we catch detection at a late enough stage in collection enumeration.
	// This is our next point of error handling, where we can identify and
	// skip over the case.
	if clues.HasLabel(err, graph.LabelsMalware) {
		cp.incExpectedErrs()
		return
	}

	defer cp.UploadProgress.Error(relpath, err, isIgnored)

	cp.errs.AddRecoverable(clues.Wrap(err, "kopia reported error").
		With("is_ignored", isIgnored, "relative_path", relpath).
		Label(fault.LabelForceNoBackupCreation))
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

func collectionEntries(
	ctx context.Context,
	cb func(context.Context, fs.Entry) error,
	streamedEnts data.BackupCollection,
	progress *corsoProgress,
) (map[string]struct{}, error) {
	if streamedEnts == nil {
		return nil, nil
	}

	var (
		locationPath path.Path
		// Track which items have already been seen so we can skip them if we see
		// them again in the data from the base snapshot.
		seen  = map[string]struct{}{}
		items = streamedEnts.Items(ctx, progress.errs)
	)

	if lp, ok := streamedEnts.(data.LocationPather); ok {
		locationPath = lp.LocationPath()
	}

	for {
		select {
		case <-ctx.Done():
			return seen, clues.Stack(ctx.Err()).WithClues(ctx)

		case e, ok := <-items:
			if !ok {
				return seen, nil
			}

			encodedName := encodeAsPath(e.UUID())

			// Even if this item has been deleted and should not appear at all in
			// the new snapshot we need to record that we've seen it here so we know
			// to skip it if it's also present in the base snapshot.
			//
			// TODO(ashmrtn): Determine if we want to try to use the old version of
			// the data (if it exists in the base) if we fail uploading the new
			// version. If so, we should split this call into where we check for the
			// item being deleted and then again after we do the kopia callback.
			//
			// TODO(ashmrtn): With a little more info, we could reduce the number of
			// items we need to track. Namely, we can track the created time of the
			// item and if it's after the base snapshot was finalized we can skip it
			// because it's not possible for the base snapshot to contain that item.
			seen[encodedName] = struct{}{}

			// For now assuming that item IDs don't need escaping.
			itemPath, err := streamedEnts.FullPath().Append(e.UUID(), true)
			if err != nil {
				err = errors.Wrap(err, "getting full item path")
				progress.errs.AddRecoverable(err)

				logger.CtxErr(ctx, err).Error("getting full item path")

				continue
			}

			trace.Log(ctx, "kopia:streamEntries:item", itemPath.String())

			if e.Deleted() {
				continue
			}

			// Not all items implement StreamInfo. For example, the metadata files
			// do not because they don't contain information directly backed up or
			// used for restore. If progress does not contain information about a
			// finished file it just returns without an error so it's safe to skip
			// adding something to it.
			ei, ok := e.(data.StreamInfo)
			if ok {
				// Relative path given to us in the callback is missing the root
				// element. Add to pending set before calling the callback to avoid race
				// conditions when the item is completed.
				//
				// TODO(ashmrtn): If we want to pull item info for cached item from a
				// previous snapshot then we should populate prevPath here and leave
				// info nil.
				itemInfo := ei.Info()
				d := &itemDetails{
					info:         &itemInfo,
					repoPath:     itemPath,
					locationPath: locationPath,
				}
				progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)
			}

			modTime := time.Now()
			if smt, ok := e.(data.StreamModTime); ok {
				modTime = smt.ModTime()
			}

			entry := virtualfs.StreamingFileWithModTimeFromReader(
				encodedName,
				modTime,
				newBackupStreamReader(serializationVersion, e.ToReader()))

			err = cb(ctx, entry)
			if err != nil {
				// Kopia's uploader swallows errors in most cases, so if we see
				// something here it's probably a big issue and we should return.
				return seen, clues.Wrap(err, "executing callback").WithClues(ctx).With("item_path", itemPath)
			}
		}
	}
}

func streamBaseEntries(
	ctx context.Context,
	cb func(context.Context, fs.Entry) error,
	curPath path.Path,
	prevPath path.Path,
	locationPath path.Path,
	dir fs.Directory,
	encodedSeen map[string]struct{},
	globalExcludeSet map[string]map[string]struct{},
	progress *corsoProgress,
) error {
	if dir == nil {
		return nil
	}

	var (
		excludeSet    map[string]struct{}
		curPrefix     string
		curPathString = curPath.String()
	)

	for prefix, excludes := range globalExcludeSet {
		// Select the set with the longest prefix to be most precise.
		if strings.HasPrefix(curPathString, prefix) && len(prefix) >= len(curPrefix) {
			excludeSet = excludes
			curPrefix = prefix
		}
	}

	err := dir.IterateEntries(ctx, func(innerCtx context.Context, entry fs.Entry) error {
		if err := innerCtx.Err(); err != nil {
			return err
		}

		// Don't walk subdirectories in this function.
		if _, ok := entry.(fs.Directory); ok {
			return nil
		}

		// This entry was either updated or deleted. In either case, the external
		// service notified us about it and it's already been handled so we should
		// skip it here.
		if _, ok := encodedSeen[entry.Name()]; ok {
			return nil
		}

		entName, err := decodeElement(entry.Name())
		if err != nil {
			return errors.Wrapf(err, "unable to decode entry name %s", entry.Name())
		}

		// This entry was marked as deleted by a service that can't tell us the
		// previous path of deleted items, only the item ID.
		if _, ok := excludeSet[entName]; ok {
			return nil
		}

		// For now assuming that item IDs don't need escaping.
		itemPath, err := curPath.Append(entName, true)
		if err != nil {
			return errors.Wrap(err, "getting full item path for base entry")
		}

		// We need the previous path so we can find this item in the base snapshot's
		// backup details. If the item moved and we had only the new path, we'd be
		// unable to find it in the old backup details because we wouldn't know what
		// to look for.
		prevItemPath, err := prevPath.Append(entName, true)
		if err != nil {
			return errors.Wrap(err, "getting previous full item path for base entry")
		}

		// Meta files aren't in backup details since it's the set of items the user
		// sees.
		//
		// TODO(ashmrtn): We may eventually want to make this a function that is
		// passed in so that we can more easily switch it between different external
		// service provider implementations.
		if !metadata.IsMetadataFile(itemPath) {
			// All items have item info in the base backup. However, we need to make
			// sure we have enough metadata to find those entries. To do that we add
			// the item to progress and having progress aggregate everything for
			// later.
			d := &itemDetails{
				info:         nil,
				repoPath:     itemPath,
				prevPath:     prevItemPath,
				locationPath: locationPath,
			}
			progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)
		}

		if err := cb(ctx, entry); err != nil {
			return errors.Wrapf(err, "executing callback on item %q", itemPath)
		}

		return nil
	})
	if err != nil {
		return errors.Wrapf(
			err,
			"traversing items in base snapshot directory %q",
			curPath,
		)
	}

	return nil
}

// getStreamItemFunc returns a function that can be used by kopia's
// virtualfs.StreamingDirectory to iterate through directory entries and call
// kopia callbacks on directory entries. It binds the directory to the given
// DataCollection.
func getStreamItemFunc(
	curPath path.Path,
	prevPath path.Path,
	staticEnts []fs.Entry,
	streamedEnts data.BackupCollection,
	baseDir fs.Directory,
	globalExcludeSet map[string]map[string]struct{},
	progress *corsoProgress,
) func(context.Context, func(context.Context, fs.Entry) error) error {
	return func(ctx context.Context, cb func(context.Context, fs.Entry) error) error {
		ctx, end := D.Span(ctx, "kopia:getStreamItemFunc")
		defer end()

		// Return static entries in this directory first.
		for _, d := range staticEnts {
			if err := cb(ctx, d); err != nil {
				return clues.Wrap(err, "executing callback on static directory").WithClues(ctx)
			}
		}

		var locationPath path.Path

		if lp, ok := streamedEnts.(data.LocationPather); ok {
			locationPath = lp.LocationPath()
		}

		seen, err := collectionEntries(ctx, cb, streamedEnts, progress)
		if err != nil {
			return errors.Wrap(err, "streaming collection entries")
		}

		if err := streamBaseEntries(
			ctx,
			cb,
			curPath,
			prevPath,
			locationPath,
			baseDir,
			seen,
			globalExcludeSet,
			progress,
		); err != nil {
			return errors.Wrap(err, "streaming base snapshot entries")
		}

		return nil
	}
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are virtualfs.StreamingDirectory.
func buildKopiaDirs(
	dirName string,
	dir *treeMap,
	globalExcludeSet map[string]map[string]struct{},
	progress *corsoProgress,
) (fs.Directory, error) {
	// Need to build the directory tree from the leaves up because intermediate
	// directories need to have all their entries at creation time.
	var childDirs []fs.Entry

	// TODO(ashmrtn): Reuse kopia directories directly if the subtree rooted at
	// them is unchanged.
	//
	// This has a few restrictions though:
	//   * if we allow for moved folders, we need to make sure we update folder
	//     names properly
	//   * we need some way to know what items need to be pulled from the base
	//     backup's backup details

	for childName, childDir := range dir.childDirs {
		child, err := buildKopiaDirs(childName, childDir, globalExcludeSet, progress)
		if err != nil {
			return nil, err
		}

		childDirs = append(childDirs, child)
	}

	return virtualfs.NewStreamingDirectory(
		encodeAsPath(dirName),
		getStreamItemFunc(
			dir.currentPath,
			dir.prevPath,
			childDirs,
			dir.collection,
			dir.baseDir,
			globalExcludeSet,
			progress,
		),
	), nil
}

type treeMap struct {
	// path.Path representing the node's path. This is passed as a parameter to
	// the stream item function so that even baseDir directories can properly
	// generate the full path of items.
	currentPath path.Path
	// Previous path this directory may have resided at if it is sourced from a
	// base snapshot.
	prevPath path.Path

	// Child directories of this directory.
	childDirs map[string]*treeMap
	// Reference to data pulled from the external service. Contains only items in
	// this directory. Does not contain references to subdirectories.
	collection data.BackupCollection
	// Reference to directory in base snapshot. The referenced directory itself
	// may contain files and subdirectories, but the subdirectories should
	// eventually be added when walking the base snapshot to build the hierarchy,
	// not when handing items to kopia for the new snapshot. Subdirectories should
	// be added to childDirs while building the hierarchy. They will be ignored
	// when iterating through the directory to hand items to kopia.
	baseDir fs.Directory
}

func newTreeMap() *treeMap {
	return &treeMap{
		childDirs: map[string]*treeMap{},
	}
}

// maybeGetTreeNode walks the tree(s) with roots roots and returns the node
// specified by pathElements if all nodes on the path exist. If pathElements is
// nil or empty then returns nil.
func maybeGetTreeNode(roots map[string]*treeMap, pathElements []string) *treeMap {
	if len(pathElements) == 0 {
		return nil
	}

	dir := roots[pathElements[0]]

	for i := 1; i < len(pathElements); i++ {
		if dir == nil {
			return nil
		}

		p := pathElements[i]

		dir = dir.childDirs[p]
	}

	return dir
}

// getTreeNode walks the tree(s) with roots roots and returns the node specified
// by pathElements. If pathElements is nil or empty then returns nil. Tree nodes
// are created for any path elements where a node is not already present.
func getTreeNode(roots map[string]*treeMap, pathElements []string) *treeMap {
	if len(pathElements) == 0 {
		return nil
	}

	dir, ok := roots[pathElements[0]]
	if !ok {
		dir = newTreeMap()
		roots[pathElements[0]] = dir
	}

	// Use actual indices so this is automatically skipped if
	// len(pathElements) == 1.
	for i := 1; i < len(pathElements); i++ {
		p := pathElements[i]

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

	return dir
}

func inflateCollectionTree(
	ctx context.Context,
	collections []data.BackupCollection,
) (map[string]*treeMap, map[string]path.Path, error) {
	roots := make(map[string]*treeMap)
	// Contains the old path for collections that have been moved or renamed.
	// Allows resolving what the new path should be when walking the base
	// snapshot(s)'s hierarchy. Nil represents a collection that was deleted.
	updatedPaths := make(map[string]path.Path)
	// Temporary variable just to track the things that have been marked as
	// changed while keeping a reference to their path.
	changedPaths := []path.Path{}

	for _, s := range collections {
		switch s.State() {
		case data.DeletedState:
			if s.PreviousPath() == nil {
				return nil, nil, errors.Errorf("nil previous path on deleted collection")
			}

			changedPaths = append(changedPaths, s.PreviousPath())

			if _, ok := updatedPaths[s.PreviousPath().String()]; ok {
				return nil, nil, errors.Errorf(
					"multiple previous state changes to collection %s",
					s.PreviousPath())
			}

			updatedPaths[s.PreviousPath().String()] = nil

			continue

		case data.MovedState:
			changedPaths = append(changedPaths, s.PreviousPath())

			if _, ok := updatedPaths[s.PreviousPath().String()]; ok {
				return nil, nil, errors.Errorf(
					"multiple previous state changes to collection %s",
					s.PreviousPath(),
				)
			}

			updatedPaths[s.PreviousPath().String()] = s.FullPath()
		}

		if s.FullPath() == nil || len(s.FullPath().Elements()) == 0 {
			return nil, nil, errors.New("no identifier for collection")
		}

		node := getTreeNode(roots, s.FullPath().Elements())
		if node == nil {
			return nil, nil, errors.Errorf(
				"unable to get tree node for path %s",
				s.FullPath(),
			)
		}

		// Make sure there's only a single collection adding items for any given
		// path in the new hierarchy.
		if node.collection != nil {
			return nil, nil, errors.Errorf("multiple instances of collection at %s", s.FullPath())
		}

		node.collection = s
		node.currentPath = s.FullPath()
		node.prevPath = s.PreviousPath()
	}

	// Check that each previous path has only one of the states of deleted, moved,
	// or notmoved. Check at the end to avoid issues like seeing a notmoved state
	// collection and then a deleted state collection.
	for _, p := range changedPaths {
		node := maybeGetTreeNode(roots, p.Elements())
		if node == nil {
			continue
		}

		if node.collection != nil && node.collection.State() == data.NotMovedState {
			return nil, nil, errors.Errorf("conflicting states for collection %s", p)
		}
	}

	return roots, updatedPaths, nil
}

// traverseBaseDir is an unoptimized function that reads items in a directory
// and traverses subdirectories in the given directory. oldDirPath is the path
// the directory would be at if the hierarchy was unchanged. newDirPath is the
// path the directory would be at if all changes from the root to this directory
// were taken into account. Both are needed to detect some changes like moving
// a parent directory and moving one of the child directories out of the parent.
// If a directory on the path was deleted, newDirPath is set to nil.
//
// TODO(ashmrtn): A potentially more memory efficient version of this would
// traverse only the directories that we know are present in the collections
// passed in. The other directories could be dynamically discovered when kopia
// was requesting items.
func traverseBaseDir(
	ctx context.Context,
	depth int,
	updatedPaths map[string]path.Path,
	oldDirPath *path.Builder,
	newDirPath *path.Builder,
	dir fs.Directory,
	roots map[string]*treeMap,
) error {
	if depth >= maxInflateTraversalDepth {
		return errors.Errorf("base snapshot tree too tall %s", oldDirPath)
	}

	// Wrapper base64 encodes all file and folder names to avoid issues with
	// special characters. Since we're working directly with files and folders
	// from kopia we need to do the decoding here.
	dirName, err := decodeElement(dir.Name())
	if err != nil {
		return errors.Wrapf(err, "decoding base directory name %s", dir.Name())
	}

	// Form the path this directory would be at if the hierarchy remained the same
	// as well as where it would be at if we take into account ancestor
	// directories that may have had changes. The former is used to check if this
	// directory specifically has been moved. The latter is used to handle
	// deletions and moving subtrees in the hierarchy.
	//
	// Explicit movement of directories should have the final say though so we
	// override any subtree movement with what's in updatedPaths if an entry
	// exists.
	oldDirPath = oldDirPath.Append(dirName)
	currentPath := newDirPath

	if currentPath != nil {
		currentPath = currentPath.Append(dirName)
	}

	if upb, ok := updatedPaths[oldDirPath.String()]; ok {
		// This directory was deleted.
		if upb == nil {
			currentPath = nil
		} else {
			// This directory was moved/renamed and the new location is in upb.
			currentPath = upb.ToBuilder()
		}
	}

	// TODO(ashmrtn): If we can do prefix matching on elements in updatedPaths and
	// we know that the tree node for this directory has no collection reference
	// and no child nodes then we can skip traversing this directory. This will
	// only work if we know what directory deleted items used to belong in (e.x.
	// it won't work for OneDrive because we only know the ID of the deleted
	// item).

	var hasItems bool

	err = dir.IterateEntries(ctx, func(innerCtx context.Context, entry fs.Entry) error {
		dEntry, ok := entry.(fs.Directory)
		if !ok {
			hasItems = true
			return nil
		}

		return traverseBaseDir(
			innerCtx,
			depth+1,
			updatedPaths,
			oldDirPath,
			currentPath,
			dEntry,
			roots,
		)
	})
	if err != nil {
		return errors.Wrapf(err, "traversing base directory %s", oldDirPath)
	}

	// We only need to add this base directory to the tree we're building if it
	// has items in it. The traversal of the directory here just finds
	// subdirectories. This optimization will not be valid if we dynamically
	// determine the subdirectories this directory has when handing items to
	// kopia.
	if currentPath != nil && hasItems {
		// Having this in the if-block has the effect of removing empty directories
		// from backups that have a base snapshot. If we'd like to preserve empty
		// directories across incremental backups, move getting the node outside of
		// the if-block. That will be sufficient to create a StreamingDirectory that
		// kopia will pick up on. Assigning the baseDir of the node should remain
		// in the if-block though as that is an optimization.
		node := getTreeNode(roots, currentPath.Elements())
		if node == nil {
			return errors.Errorf("unable to get tree node for path %s", currentPath)
		}

		// Now that we have the node we need to check if there is a collection
		// marked DoNotMerge. If there is, skip adding a reference to this base dir
		// in the node. That allows us to propagate subtree operations (e.x. move)
		// while selectively skipping merging old and new versions for some
		// directories. The expected usecase for this is delta token expiry in M365.
		if node.collection != nil &&
			(node.collection.DoNotMergeItems() || node.collection.State() == data.NewState) {
			return nil
		}

		curP, err := path.FromDataLayerPath(currentPath.String(), false)
		if err != nil {
			return errors.Errorf(
				"unable to convert current path %s to path.Path",
				currentPath,
			)
		}

		oldP, err := path.FromDataLayerPath(oldDirPath.String(), false)
		if err != nil {
			return errors.Errorf(
				"unable to convert old path %s to path.Path",
				oldDirPath,
			)
		}

		node.baseDir = dir
		node.currentPath = curP
		node.prevPath = oldP
	}

	return nil
}

func inflateBaseTree(
	ctx context.Context,
	loader snapshotLoader,
	snap IncrementalBase,
	updatedPaths map[string]path.Path,
	roots map[string]*treeMap,
) error {
	// Only complete snapshots should be used to source base information.
	// Snapshots for checkpoints will rely on kopia-assisted dedupe to efficiently
	// handle items that were completely uploaded before Corso crashed.
	if len(snap.IncompleteReason) > 0 {
		return nil
	}

	root, err := loader.SnapshotRoot(snap.Manifest)
	if err != nil {
		return errors.Wrapf(err, "getting snapshot %s root directory", snap.ID)
	}

	dir, ok := root.(fs.Directory)
	if !ok {
		return errors.Errorf("snapshot %s root is not a directory", snap.ID)
	}

	// For each subtree corresponding to the tuple
	// (resource owner, service, category) merge the directories in the base with
	// what has been reported in the collections we got.
	for _, subtreePath := range snap.SubtreePaths {
		// We're starting from the root directory so don't need it in the path.
		pathElems := encodeElements(subtreePath.PopFront().Elements()...)

		ent, err := snapshotfs.GetNestedEntry(ctx, dir, pathElems)
		if err != nil {
			if isErrEntryNotFound(err) {
				logger.Ctx(ctx).Infow("base snapshot missing subtree", "error", err)
				continue
			}

			return errors.Wrapf(err, "snapshot %s getting subtree root", snap.ID)
		}

		subtreeDir, ok := ent.(fs.Directory)
		if !ok {
			return errors.Wrapf(err, "snapshot %s subtree root is not directory", snap.ID)
		}

		// We're assuming here that the prefix for the path has not changed (i.e.
		// all of tenant, service, resource owner, and category are the same in the
		// old snapshot (snap) and the snapshot we're currently trying to make.
		if err = traverseBaseDir(
			ctx,
			0,
			updatedPaths,
			subtreePath.Dir(),
			subtreePath.Dir(),
			subtreeDir,
			roots,
		); err != nil {
			return errors.Wrapf(err, "traversing base snapshot %s", snap.ID)
		}
	}

	return nil
}

// inflateDirTree returns a set of tags representing all the resource owners and
// service/categories in the snapshot and a fs.Directory tree rooted at the
// oldest common ancestor of the streams. All nodes are
// virtualfs.StreamingDirectory with the given DataCollections if there is one
// for that node. Tags can be used in future backups to fetch old snapshots for
// caching reasons.
//
// globalExcludeSet represents a set of items, represented with file names, to
// exclude from base directories when uploading the snapshot. As items in *all*
// base directories will be checked for in every base directory, this assumes
// that items in the bases are unique. Deletions of directories or subtrees
// should be represented as changes in the status of a BackupCollection, not an
// entry in the globalExcludeSet.
func inflateDirTree(
	ctx context.Context,
	loader snapshotLoader,
	baseSnaps []IncrementalBase,
	collections []data.BackupCollection,
	globalExcludeSet map[string]map[string]struct{},
	progress *corsoProgress,
) (fs.Directory, error) {
	roots, updatedPaths, err := inflateCollectionTree(ctx, collections)
	if err != nil {
		return nil, errors.Wrap(err, "inflating collection tree")
	}

	baseIDs := make([]manifest.ID, 0, len(baseSnaps))
	for _, snap := range baseSnaps {
		baseIDs = append(baseIDs, snap.ID)
	}

	logger.Ctx(ctx).Infow(
		"merging hierarchies from base snapshots",
		"snapshot_ids", baseIDs)

	for _, snap := range baseSnaps {
		if err = inflateBaseTree(ctx, loader, snap, updatedPaths, roots); err != nil {
			return nil, errors.Wrap(err, "inflating base snapshot tree(s)")
		}
	}

	if len(roots) > 1 {
		return nil, errors.New("multiple root directories")
	}

	var res fs.Directory

	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(dirName, dir, globalExcludeSet, progress)
		if err != nil {
			return nil, err
		}

		res = tmp
	}

	return res, nil
}
