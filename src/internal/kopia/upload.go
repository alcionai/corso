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
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/graph/metadata"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
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

	var errs *clues.Err

	for _, r := range rw.readers {
		err := r.Close()
		if err != nil {
			errs = clues.Stack(clues.Wrap(err, "closing reader"), errs)
		}
	}

	return errs.OrNil()
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
			return clues.Wrap(err, "reading data format version")
		}

		newlyRead += n
	}

	version := binary.BigEndian.Uint32(versionBuf)

	if version != rw.expectedVersion {
		return clues.New("unexpected data format").With("read_version", version)
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
	locationPath *path.Builder
	cached       bool
	modTime      *time.Time
}

type corsoProgress struct {
	// this is an unwanted hack.  We can't extend the kopia interface
	// funcs to pass through a context.  This is the second best way to
	// get an at least partially formed context into funcs that need it
	// for logging and other purposes.
	ctx context.Context

	snapshotfs.UploadProgress
	pending map[string]*itemDetails
	// deets contains entries that are complete and don't need merged with base
	// backup data at all.
	deets *details.Builder
	// toMerge represents items that we either don't have in-memory item info or
	// that need sourced from a base backup due to caching etc.
	toMerge    *mergeDetails
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

	ctx := clues.Add(
		cp.ctx,
		"service", d.repoPath.Service().String(),
		"category", d.repoPath.Category().String())

	// These items were sourced from a base snapshot or were cached in kopia so we
	// never had to materialize their details in-memory.
	if d.info == nil || d.cached {
		if d.prevPath == nil {
			cp.errs.AddRecoverable(ctx, clues.New("finished file sourced from previous backup with no previous path").
				Label(fault.LabelForceNoBackupCreation))

			return
		}

		cp.mu.Lock()
		defer cp.mu.Unlock()

		err := cp.toMerge.addRepoRef(
			d.prevPath.ToBuilder(),
			d.modTime,
			d.repoPath,
			d.locationPath)
		if err != nil {
			cp.errs.AddRecoverable(ctx, clues.Wrap(err, "adding finished file to merge list").
				Label(fault.LabelForceNoBackupCreation))
		}

		return
	}

	err = cp.deets.Add(
		d.repoPath,
		d.locationPath,
		*d.info)
	if err != nil {
		cp.errs.AddRecoverable(ctx, clues.Wrap(err, "adding finished file to details").
			Label(fault.LabelForceNoBackupCreation))

		return
	}
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

	logger.Ctx(cp.ctx).Debugw(
		"finished hashing file",
		"path", clues.Hide(path.Elements(sl[2:])))

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
	// LabelsSkippable is set of malware items or not found items.
	// The malware case is an artifact of being unable to skip the
	// item if we catch detection at a late enough stage in collection
	// enumeration. The not found could be items deleted in between a
	// delta query and a fetch.  This is our next point of error
	// handling, where we can identify and skip over the case.
	if clues.HasLabel(err, graph.LabelsSkippable) {
		cp.incExpectedErrs()
		return
	}

	defer cp.UploadProgress.Error(relpath, err, isIgnored)

	cp.errs.AddRecoverable(cp.ctx, clues.Wrap(err, "kopia reported error").
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
	ctr func(context.Context, fs.Entry) error,
	streamedEnts data.BackupCollection,
	progress *corsoProgress,
) (map[string]struct{}, error) {
	if streamedEnts == nil {
		return nil, nil
	}

	var (
		locationPath *path.Builder
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

			encodedName := encodeAsPath(e.ID())

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
			itemPath, err := streamedEnts.FullPath().AppendItem(e.ID())
			if err != nil {
				err = clues.Wrap(err, "getting full item path")
				progress.errs.AddRecoverable(ctx, err)

				logger.CtxErr(ctx, err).Error("getting full item path")

				continue
			}

			trace.Log(ctx, "kopia:streamEntries:item", itemPath.String())

			if e.Deleted() {
				continue
			}

			modTime := time.Now()
			if smt, ok := e.(data.ItemModTime); ok {
				modTime = smt.ModTime()
			}

			// Not all items implement StreamInfo. For example, the metadata files
			// do not because they don't contain information directly backed up or
			// used for restore. If progress does not contain information about a
			// finished file it just returns without an error so it's safe to skip
			// adding something to it.
			ei, ok := e.(data.ItemInfo)
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
					info:     &itemInfo,
					repoPath: itemPath,
					// Also use the current path as the previous path for this item. This
					// is so that if the item is marked as cached and we need to merge
					// details with an assist backup base which sourced the cached item we
					// can find it with the lookup in DetailsMergeInfoer.
					//
					// This all works out because cached item checks in kopia are direct
					// path + metadata comparisons.
					prevPath:     itemPath,
					locationPath: locationPath,
					modTime:      &modTime,
				}
				progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)
			}

			entry := virtualfs.StreamingFileWithModTimeFromReader(
				encodedName,
				modTime,
				newBackupStreamReader(serializationVersion, e.ToReader()))

			err = ctr(ctx, entry)
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
	ctr func(context.Context, fs.Entry) error,
	curPath path.Path,
	prevPath path.Path,
	locationPath *path.Builder,
	dir fs.Directory,
	encodedSeen map[string]struct{},
	globalExcludeSet prefixmatcher.StringSetReader,
	progress *corsoProgress,
) error {
	if dir == nil {
		return nil
	}

	var (
		longest    string
		excludeSet map[string]struct{}
	)

	if globalExcludeSet != nil {
		longest, excludeSet, _ = globalExcludeSet.LongestPrefix(curPath.String())
	}

	ctx = clues.Add(
		ctx,
		"current_directory_path", curPath,
		"longest_prefix", path.LoggableDir(longest))

	err := dir.IterateEntries(ctx, func(innerCtx context.Context, entry fs.Entry) error {
		if err := innerCtx.Err(); err != nil {
			return clues.Stack(err).WithClues(ctx)
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
			return clues.Wrap(err, "decoding entry name").
				WithClues(ctx).
				With("entry_name", entry.Name())
		}

		// This entry was marked as deleted by a service that can't tell us the
		// previous path of deleted items, only the item ID.
		if _, ok := excludeSet[entName]; ok {
			return nil
		}

		// For now assuming that item IDs don't need escaping.
		itemPath, err := curPath.AppendItem(entName)
		if err != nil {
			return clues.Wrap(err, "getting full item path for base entry").WithClues(ctx)
		}

		// We need the previous path so we can find this item in the base snapshot's
		// backup details. If the item moved and we had only the new path, we'd be
		// unable to find it in the old backup details because we wouldn't know what
		// to look for.
		prevItemPath, err := prevPath.AppendItem(entName)
		if err != nil {
			return clues.Wrap(err, "getting previous full item path for base entry").WithClues(ctx)
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
				modTime:      ptr.To(entry.ModTime()),
			}
			progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)
		}

		if err := ctr(ctx, entry); err != nil {
			return clues.Wrap(err, "executing callback on item").
				WithClues(ctx).
				With("item_path", itemPath)
		}

		return nil
	})
	if err != nil {
		return clues.Wrap(err, "traversing items in base snapshot directory").WithClues(ctx)
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
	globalExcludeSet prefixmatcher.StringSetReader,
	progress *corsoProgress,
) func(context.Context, func(context.Context, fs.Entry) error) error {
	return func(ctx context.Context, ctr func(context.Context, fs.Entry) error) error {
		ctx, end := diagnostics.Span(ctx, "kopia:getStreamItemFunc")
		defer end()

		// Return static entries in this directory first.
		for _, d := range staticEnts {
			if err := ctr(ctx, d); err != nil {
				return clues.Wrap(err, "executing callback on static directory").WithClues(ctx)
			}
		}

		var locationPath *path.Builder

		if lp, ok := streamedEnts.(data.LocationPather); ok {
			locationPath = lp.LocationPath()
		}

		seen, err := collectionEntries(ctx, ctr, streamedEnts, progress)
		if err != nil {
			return clues.Wrap(err, "streaming collection entries")
		}

		if err := streamBaseEntries(
			ctx,
			ctr,
			curPath,
			prevPath,
			locationPath,
			baseDir,
			seen,
			globalExcludeSet,
			progress,
		); err != nil {
			return clues.Wrap(err, "streaming base snapshot entries")
		}

		return nil
	}
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are virtualfs.StreamingDirectory.
func buildKopiaDirs(
	dirName string,
	dir *treeMap,
	globalExcludeSet prefixmatcher.StringSetReader,
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

func addMergeLocation(col data.BackupCollection, toMerge *mergeDetails) error {
	lp, ok := col.(data.PreviousLocationPather)
	if !ok {
		return nil
	}

	prevLoc := lp.PreviousLocationPath()
	newLoc := lp.LocationPath()

	if prevLoc == nil {
		return clues.New("moved collection with nil previous location")
	} else if newLoc == nil {
		return clues.New("moved collection with nil location")
	}

	if err := toMerge.addLocation(prevLoc, newLoc); err != nil {
		return clues.Wrap(err, "building updated location set").
			With(
				"collection_previous_location", prevLoc,
				"collection_location", newLoc)
	}

	return nil
}

func inflateCollectionTree(
	ctx context.Context,
	collections []data.BackupCollection,
	toMerge *mergeDetails,
) (map[string]*treeMap, map[string]path.Path, error) {
	roots := make(map[string]*treeMap)
	// Contains the old path for collections that are not new.
	// Allows resolving what the new path should be when walking the base
	// snapshot(s)'s hierarchy. Nil represents a collection that was deleted.
	updatedPaths := make(map[string]path.Path)
	// Temporary variable just to track the things that have been marked as
	// changed while keeping a reference to their path.
	changedPaths := []path.Path{}

	for _, s := range collections {
		ictx := clues.Add(
			ctx,
			"collection_full_path", s.FullPath(),
			"collection_previous_path", s.PreviousPath())

		switch s.State() {
		case data.DeletedState:
			if s.PreviousPath() == nil {
				return nil, nil, clues.New("nil previous path on deleted collection").WithClues(ictx)
			}

			changedPaths = append(changedPaths, s.PreviousPath())

			if _, ok := updatedPaths[s.PreviousPath().String()]; ok {
				return nil, nil, clues.New("multiple previous state changes to collection").
					WithClues(ictx)
			}

			updatedPaths[s.PreviousPath().String()] = nil

			continue

		case data.MovedState:
			changedPaths = append(changedPaths, s.PreviousPath())

			if _, ok := updatedPaths[s.PreviousPath().String()]; ok {
				return nil, nil, clues.New("multiple previous state changes to collection").
					WithClues(ictx)
			}

			updatedPaths[s.PreviousPath().String()] = s.FullPath()

			// Only safe when collections are moved since we only need prefix matching
			// if a nested folder's path changed in some way that didn't generate a
			// collection. For that to the be case, the nested folder's path must have
			// changed via one of the ancestor folders being moved. This catches the
			// ancestor folder move.
			if err := addMergeLocation(s, toMerge); err != nil {
				return nil, nil, clues.Wrap(err, "adding merge location").WithClues(ictx)
			}
		case data.NotMovedState:
			p := s.PreviousPath().String()
			if _, ok := updatedPaths[p]; ok {
				return nil, nil, clues.New("multiple previous state changes to collection").
					WithClues(ictx)
			}

			updatedPaths[p] = s.FullPath()
		}

		if s.FullPath() == nil || len(s.FullPath().Elements()) == 0 {
			return nil, nil, clues.New("no identifier for collection").WithClues(ictx)
		}

		node := getTreeNode(roots, s.FullPath().Elements())
		if node == nil {
			return nil, nil, clues.New("getting tree node").WithClues(ictx)
		}

		// Make sure there's only a single collection adding items for any given
		// path in the new hierarchy.
		if node.collection != nil {
			return nil, nil, clues.New("multiple instances of collection").WithClues(ictx)
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
			return nil, nil, clues.New("conflicting states for collection").
				WithClues(ctx).
				With("changed_path", p)
		}
	}

	return roots, updatedPaths, nil
}

// traverseBaseDir is an unoptimized function that reads items in a directory
// and traverses subdirectories in the given directory. oldDirPath is the path
// the directory would be at if the hierarchy was unchanged. expectedDirPath is the
// path the directory would be at if all changes from the root to this directory
// were taken into account. Both are needed to detect some changes like moving
// a parent directory and moving one of the child directories out of the parent.
// If a directory on the path was deleted, expectedDirPath is set to nil.
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
	expectedDirPath *path.Builder,
	dir fs.Directory,
	roots map[string]*treeMap,
	stats *count.Bus,
) error {
	ctx = clues.Add(ctx,
		"old_parent_dir_path", oldDirPath,
		"expected_parent_dir_path", expectedDirPath)

	if depth >= maxInflateTraversalDepth {
		return clues.New("base snapshot tree too tall").WithClues(ctx)
	}

	// Wrapper base64 encodes all file and folder names to avoid issues with
	// special characters. Since we're working directly with files and folders
	// from kopia we need to do the decoding here.
	dirName, err := decodeElement(dir.Name())
	if err != nil {
		return clues.Wrap(err, "decoding base directory name").
			WithClues(ctx).
			With("dir_name", clues.Hide(dir.Name()))
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
	currentPath := expectedDirPath

	if currentPath != nil {
		currentPath = currentPath.Append(dirName)
	}

	var explicitMention bool

	if upb, ok := updatedPaths[oldDirPath.String()]; ok {
		// This directory was deleted.
		if upb == nil {
			currentPath = nil

			stats.Inc(statDel)
		} else {
			// This directory was explicitly mentioned and the new (possibly
			// unchanged) location is in upb.
			currentPath = upb.ToBuilder()

			// Below we check if the collection was marked as new or DoNotMerge which
			// disables merging behavior. That means we can't directly update stats
			// here else we'll miss delta token refreshes and whatnot. Instead note
			// that we did see the path explicitly so it's not counted as a recursive
			// operation.
			explicitMention = true
		}
	} else if currentPath == nil {
		// Just stats tracking stuff.
		stats.Inc(statRecursiveDel)
	}

	ctx = clues.Add(ctx, "new_path", currentPath)

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
			stats)
	})
	if err != nil {
		return clues.Wrap(err, "traversing base directory").WithClues(ctx)
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
			return clues.New("getting tree node").WithClues(ctx)
		}

		// Now that we have the node we need to check if there is a collection
		// marked DoNotMerge. If there is, skip adding a reference to this base dir
		// in the node. That allows us to propagate subtree operations (e.x. move)
		// while selectively skipping merging old and new versions for some
		// directories. The expected usecase for this is delta token expiry in M365.
		if node.collection != nil &&
			(node.collection.DoNotMergeItems() || node.collection.State() == data.NewState) {
			stats.Inc(statSkipMerge)

			return nil
		}

		// Just stats tracking stuff.
		if oldDirPath.String() == currentPath.String() {
			stats.Inc(statNoMove)
		} else if explicitMention {
			stats.Inc(statMove)
		} else {
			stats.Inc(statRecursiveMove)
		}

		curP, err := path.FromDataLayerPath(currentPath.String(), false)
		if err != nil {
			return clues.New("converting current path to path.Path").WithClues(ctx)
		}

		oldP, err := path.FromDataLayerPath(oldDirPath.String(), false)
		if err != nil {
			return clues.New("converting old path to path.Path").WithClues(ctx)
		}

		node.baseDir = dir
		node.currentPath = curP
		node.prevPath = oldP
	}

	return nil
}

func logBaseInfo(ctx context.Context, m ManifestEntry) {
	svcs := map[string]struct{}{}
	cats := map[string]struct{}{}

	for _, r := range m.Reasons {
		svcs[r.Service().String()] = struct{}{}
		cats[r.Category().String()] = struct{}{}
	}

	mbID, _ := m.GetTag(TagBackupID)
	if len(mbID) == 0 {
		mbID = "no_backup_id_tag"
	}

	logger.Ctx(ctx).Infow(
		"using base for backup",
		"base_snapshot_id", m.ID,
		"services", maps.Keys(svcs),
		"categories", maps.Keys(cats),
		"base_backup_id", mbID)
}

const (
	// statNoMove denotes an directory that wasn't moved at all.
	statNoMove = "directories_not_moved"
	// statMove denotes an directory that was explicitly moved.
	statMove = "directories_explicitly_moved"
	// statRecursiveMove denotes an directory that moved because one or more or
	// its ancestors moved and it wasn't explicitly mentioned.
	statRecursiveMove = "directories_recursively_moved"
	// statDel denotes a directory that was explicitly deleted.
	statDel = "directories_explicitly_deleted"
	// statRecursiveDel denotes a directory that was deleted because one or more
	// of its ancestors was deleted and it wasn't explicitly mentioned.
	statRecursiveDel = "directories_recursively_deleted"
	// statSkipMerge denotes the number of directories that weren't merged because
	// they were marked either DoNotMerge or New.
	statSkipMerge = "directories_skipped_merging"
)

func inflateBaseTree(
	ctx context.Context,
	loader snapshotLoader,
	snap ManifestEntry,
	updatedPaths map[string]path.Path,
	roots map[string]*treeMap,
) error {
	// Only complete snapshots should be used to source base information.
	// Snapshots for checkpoints will rely on kopia-assisted dedupe to efficiently
	// handle items that were completely uploaded before Corso crashed.
	if len(snap.IncompleteReason) > 0 {
		return nil
	}

	ctx = clues.Add(ctx, "snapshot_base_id", snap.ID)

	root, err := loader.SnapshotRoot(snap.Manifest)
	if err != nil {
		return clues.Wrap(err, "getting snapshot root directory").WithClues(ctx)
	}

	dir, ok := root.(fs.Directory)
	if !ok {
		return clues.New("snapshot root is not a directory").WithClues(ctx)
	}

	// Some logging to help track things.
	logBaseInfo(ctx, snap)

	// For each subtree corresponding to the tuple
	// (resource owner, service, category) merge the directories in the base with
	// what has been reported in the collections we got.
	for _, r := range snap.Reasons {
		ictx := clues.Add(
			ctx,
			"subtree_service", r.Service().String(),
			"subtree_category", r.Category().String())

		subtreePath, err := r.SubtreePath()
		if err != nil {
			return clues.Wrap(err, "building subtree path").WithClues(ictx)
		}

		// We're starting from the root directory so don't need it in the path.
		pathElems := encodeElements(subtreePath.PopFront().Elements()...)

		ent, err := snapshotfs.GetNestedEntry(ictx, dir, pathElems)
		if err != nil {
			if isErrEntryNotFound(err) {
				logger.CtxErr(ictx, err).Infow("base snapshot missing subtree")
				continue
			}

			return clues.Wrap(err, "getting subtree root").WithClues(ictx)
		}

		subtreeDir, ok := ent.(fs.Directory)
		if !ok {
			return clues.Wrap(err, "subtree root is not directory").WithClues(ictx)
		}

		// This ensures that a migration on the directory prefix can complete.
		// The prefix is the tenant/service/owner/category set, which remains
		// otherwise unchecked in tree inflation below this point.
		newSubtreePath := subtreePath.ToBuilder()

		if p, ok := updatedPaths[subtreePath.String()]; ok {
			newSubtreePath = p.ToBuilder()
		}

		stats := count.New()

		if err = traverseBaseDir(
			ictx,
			0,
			updatedPaths,
			subtreePath.ToBuilder().Dir(),
			newSubtreePath.Dir(),
			subtreeDir,
			roots,
			stats,
		); err != nil {
			return clues.Wrap(err, "traversing base snapshot").WithClues(ictx)
		}

		logger.Ctx(ctx).Infow(
			"merge subtree stats",
			statNoMove, stats.Get(statNoMove),
			statMove, stats.Get(statMove),
			statRecursiveMove, stats.Get(statRecursiveMove),
			statDel, stats.Get(statDel),
			statRecursiveDel, stats.Get(statRecursiveDel),
			statSkipMerge, stats.Get(statSkipMerge))
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
	baseSnaps []ManifestEntry,
	collections []data.BackupCollection,
	globalExcludeSet prefixmatcher.StringSetReader,
	progress *corsoProgress,
) (fs.Directory, error) {
	roots, updatedPaths, err := inflateCollectionTree(ctx, collections, progress.toMerge)
	if err != nil {
		return nil, clues.Wrap(err, "inflating collection tree")
	}

	baseIDs := make([]manifest.ID, 0, len(baseSnaps))
	for _, snap := range baseSnaps {
		baseIDs = append(baseIDs, snap.ID)
	}

	ctx = clues.Add(ctx, "len_base_snapshots", len(baseSnaps), "base_snapshot_ids", baseIDs)

	if len(baseIDs) > 0 {
		logger.Ctx(ctx).Info("merging hierarchies from base snapshots")
	} else {
		logger.Ctx(ctx).Info("no base snapshots to merge")
	}

	for _, snap := range baseSnaps {
		if err = inflateBaseTree(ctx, loader, snap, updatedPaths, roots); err != nil {
			return nil, clues.Wrap(err, "inflating base snapshot tree(s)")
		}
	}

	if len(roots) > 1 {
		return nil, clues.New("multiple root directories").WithClues(ctx)
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
