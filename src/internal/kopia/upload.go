package kopia

import (
	"context"
	"encoding/base64"
	"errors"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph/metadata"
)

const maxInflateTraversalDepth = 500

type itemDetails struct {
	infoer       data.ItemInfo
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
	totalFiles int64
	errs       *fault.Bus
	counter    *count.Bus
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

	atomic.AddInt64(&cp.totalFiles, 1)

	// Log every 1000 items uploaded
	if cp.totalFiles%1000 == 0 {
		logger.Ctx(cp.ctx).Infow("finishedfile", "totalFiles", cp.totalFiles, "totalBytes", cp.totalBytes)
	}

	d := cp.get(relativePath)
	if d == nil {
		return
	}

	ctx := clues.Add(
		cp.ctx,
		"service", d.repoPath.Service().String(),
		"category", d.repoPath.Category().String(),
		"item_path", d.repoPath,
		"item_loc", d.locationPath)

	// These items were sourced from a base snapshot or were cached in kopia so we
	// never had to materialize their details in-memory.
	if d.infoer == nil || d.cached {
		if d.prevPath == nil {
			cp.errs.AddRecoverable(ctx, clues.NewWC(ctx, "finished file sourced from previous backup with no previous path").
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
			cp.errs.AddRecoverable(ctx, clues.WrapWC(ctx, err, "adding finished file to merge list").
				Label(fault.LabelForceNoBackupCreation))
		}

		return
	}

	info, err := d.infoer.Info()
	if errors.Is(err, data.ErrNotFound) {
		// The item was deleted between enumeration and trying to get data. Skip
		// adding it to details since there's no data for it.
		return
	} else if err != nil {
		cp.errs.AddRecoverable(ctx, clues.WrapWC(ctx, err, "getting ItemInfo").
			Label(fault.LabelForceNoBackupCreation))

		return
	} else if !ptr.Val(d.modTime).Equal(info.Modified()) {
		cp.errs.AddRecoverable(ctx, clues.NewWC(ctx, "item modTime mismatch").
			Label(fault.LabelForceNoBackupCreation))

		return
	} else if info.Modified().IsZero() {
		cp.errs.AddRecoverable(ctx, clues.NewWC(ctx, "zero-valued mod time").
			Label(fault.LabelForceNoBackupCreation))
	}

	err = cp.deets.Add(d.repoPath, d.locationPath, info)
	if err != nil {
		cp.errs.AddRecoverable(ctx, clues.WrapWC(ctx, err, "adding finished file to details").
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
			logger.Ctx(cp.ctx).Infow(
				"unable to decode base64 path segment",
				"segment", sl[i])
		} else {
			sl[i] = string(rdt)
		}
	}

	logger.Ctx(cp.ctx).Debugw(
		"finished hashing file",
		"path", clues.Hide(path.Elements(sl[2:])))

	cp.counter.Add(count.PersistedHashedBytes, bs)
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
		cp.counter.Inc(count.PersistenceExpectedErrors)
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

// These define a small state machine as to which source to return an entry from
// next. Since these are in-memory only values we can use iota. Phases are
// traversed in the order defined unless the underlying data source isn't
// present. If an underlying data source is missing, the non-pre/post phase
// associated with that data source is skipped.
//
// Since some phases require initialization of the underlying data source we
// insert additional phases to allow that. Once initialization is completed the
// phase should be updated to the next phase.
//
// A similar tactic can be used to handle tearing down resources for underlying
// data sources if needed.
const (
	initPhase = iota
	staticEntsPhase
	preStreamEntsPhase
	streamEntsPhase
	postStreamEntsPhase
	preBaseDirEntsPhase
	baseDirEntsPhase
	postBaseDirEntsPhase
	terminationPhase
)

type corsoDirectoryIterator struct {
	ctx    context.Context
	params snapshotParams
	// staticEnts is the set of fs.StreamingDirectory child directories that we've
	// generated based on the collections passed in. These entries may or may not
	// contain an underlying data.BackupCollection (depending on what was passed
	// in) and may or may not contain an fs.Directory (depending on hierarchy
	// merging).
	staticEnts       []fs.Entry
	globalExcludeSet prefixmatcher.StringSetReader
	progress         *corsoProgress

	// endSpan is the callback to stop diagnostic span collection for iteration.
	endSpan func()

	// seenEnts contains the encoded names of entries that we've already streamed
	// so we can skip returning them again when looking at base entries.
	seenEnts map[string]struct{}
	// locationPath contains the human-readable location of the underlying
	// collection.
	locationPath *path.Builder

	// excludeSet is the individual exclude set to use for the longest prefix for
	// this iterator.
	excludeSet map[string]struct{}

	// traversalPhase is the current state in the state machine.
	traversalPhase int

	// streamItemsChan contains the channel for the backing collection if there is
	// one. Once the backing collection has been traversed this is set to nil.
	streamItemsChan <-chan data.Item
	// staticEntsIdx contains the index in staticEnts of the next item to be
	// returned. Once all static entries have been traversed this is set to
	// len(staticEnts).
	staticEntsIdx int
	// baseDirIter contains the handle to the iterator for the base directory
	// found during hierarchy merging. Once all base directory entries have been
	// traversed this is set to nil.
	baseDirIter fs.DirectoryIterator
}

// Close releases any remaining resources the iterator may have at the end of
// iteration.
func (d *corsoDirectoryIterator) Close() {
	if d.endSpan != nil {
		d.endSpan()
	}
}

func (d *corsoDirectoryIterator) Next(ctx context.Context) (fs.Entry, error) {
	// Execute the state machine until either:
	//   * we get an entry to return
	//   * we exhaust all underlying data sources (end of iteration)
	//
	// Multiple executions of the state machine may be required for things like
	// setting up underlying data sources or finding that there's no more entries
	// in the current data source and needing to switch to the next one.
	//
	// Returned entries are handled with inline return statements.
	//
	// When an error is encountered it's added to the fault.Bus. We can't return
	// these errors since doing so will result in kopia stopping iteration of the
	// directory. Since these errors are recorded we won't lose track of them at
	// the end of the backup.
	for d.traversalPhase != terminationPhase {
		switch d.traversalPhase {
		case initPhase:
			d.ctx, d.endSpan = diagnostics.Span(d.ctx, "kopia:DirectoryIterator")
			d.traversalPhase = staticEntsPhase

		case staticEntsPhase:
			if d.staticEntsIdx < len(d.staticEnts) {
				ent := d.staticEnts[d.staticEntsIdx]
				d.staticEntsIdx++

				return ent, nil
			}

			d.traversalPhase = preStreamEntsPhase

		case preStreamEntsPhase:
			if d.params.collection == nil {
				d.traversalPhase = preBaseDirEntsPhase
				break
			}

			if lp, ok := d.params.collection.(data.LocationPather); ok {
				d.locationPath = lp.LocationPath()
			}

			d.streamItemsChan = d.params.collection.Items(d.ctx, d.progress.errs)
			d.seenEnts = map[string]struct{}{}
			d.traversalPhase = streamEntsPhase

		case streamEntsPhase:
			ent, err := d.nextStreamEnt(d.ctx)
			if ent != nil {
				return ent, nil
			}

			// This assumes that once we hit an error we won't generate any more valid
			// entries. Record the error in progress but don't return it to kopia
			// since doing so will terminate iteration.
			if err != nil {
				d.progress.errs.AddRecoverable(d.ctx, clues.Stack(err))
			}

			// Done iterating through stream entries, advance the state machine state.
			d.traversalPhase = postStreamEntsPhase

		case postStreamEntsPhase:
			d.streamItemsChan = nil
			d.traversalPhase = preBaseDirEntsPhase

		case preBaseDirEntsPhase:
			// We have no iterator from which to pull entries, switch to the next
			// state machine state.
			if d.params.baseDir == nil {
				d.traversalPhase = postBaseDirEntsPhase
				break
			}

			var err error

			d.baseDirIter, err = d.params.baseDir.Iterate(d.ctx)
			if err != nil {
				// We have no iterator from which to pull entries, switch to the next
				// state machine state.
				d.traversalPhase = postBaseDirEntsPhase
				d.progress.errs.AddRecoverable(
					d.ctx,
					clues.Wrap(err, "getting base directory iterator"))

				break
			}

			if d.globalExcludeSet != nil {
				longest, excludeSet, _ := d.globalExcludeSet.LongestPrefix(
					d.params.currentPath.String())
				d.excludeSet = excludeSet

				logger.Ctx(d.ctx).Debugw("found exclude set", "set_prefix", longest)
			}

			d.traversalPhase = baseDirEntsPhase

		case baseDirEntsPhase:
			ent, err := d.nextBaseEnt(d.ctx)
			if ent != nil {
				return ent, nil
			}

			// This assumes that once we hit an error we won't generate any more valid
			// entries. Record the error in progress but don't return it to kopia
			// since doing so will terminate iteration.
			if err != nil {
				d.progress.errs.AddRecoverable(d.ctx, clues.Stack(err))
			}

			// Done iterating through base entries, advance the state machine state.
			d.traversalPhase = postBaseDirEntsPhase

		case postBaseDirEntsPhase:
			// Making a separate phase so adding additional phases after this one is
			// less error prone if we ever need to do that.
			if d.baseDirIter != nil {
				d.baseDirIter.Close()
				d.baseDirIter = nil
			}

			d.seenEnts = nil
			d.excludeSet = nil

			d.traversalPhase = terminationPhase
		}
	}

	return nil, nil
}

func (d *corsoDirectoryIterator) nextStreamEnt(
	ctx context.Context,
) (fs.Entry, error) {
	// Loop over results until we get something we can return. Required because we
	// could see deleted items.
	for {
		select {
		case <-ctx.Done():
			return nil, clues.StackWC(ctx, ctx.Err())

		case e, ok := <-d.streamItemsChan:
			// Channel was closed, no more entries to return.
			if !ok {
				return nil, nil
			}

			// Got an entry to process, see if it's a deletion marker or something to
			// return to kopia.
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
			d.seenEnts[encodedName] = struct{}{}

			// For now assuming that item IDs don't need escaping.
			itemPath, err := d.params.currentPath.AppendItem(e.ID())
			if err != nil {
				err = clues.Wrap(err, "getting full item path")
				d.progress.errs.AddRecoverable(ctx, err)

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
				// element.
				deetsEnt := &itemDetails{
					infoer:   ei,
					repoPath: itemPath,
					// Also use the current path as the previous path for this item. This
					// is so that if the item is marked as cached and we need to merge
					// details with an assist backup base which sourced the cached item we
					// can find it with the lookup in DetailsMergeInfoer.
					//
					// This all works out because cached item checks in kopia are direct
					// path + metadata comparisons.
					prevPath:     itemPath,
					locationPath: d.locationPath,
					modTime:      &modTime,
				}
				d.progress.put(
					encodeAsPath(itemPath.PopFront().Elements()...),
					deetsEnt)
			}

			return virtualfs.StreamingFileWithModTimeFromReader(
				encodedName,
				modTime,
				e.ToReader()), nil
		}
	}
}

func (d *corsoDirectoryIterator) nextBaseEnt(
	ctx context.Context,
) (fs.Entry, error) {
	var (
		entry fs.Entry
		err   error
	)

	for entry, err = d.baseDirIter.Next(ctx); entry != nil && err == nil; entry, err = d.baseDirIter.Next(ctx) {
		entName, err := decodeElement(entry.Name())
		if err != nil {
			err = clues.Wrap(err, "decoding entry name").
				WithClues(ctx).
				With("entry_name", clues.Hide(entry.Name()))
			d.progress.errs.AddRecoverable(ctx, err)

			continue
		}

		ctx = clues.Add(ctx, "entry_name", clues.Hide(entName))

		if dir, ok := entry.(fs.Directory); ok {
			// Don't walk subdirectories in this function.
			if !d.params.streamBaseEnts {
				continue
			}

			// Do walk subdirectories. The previous and current path of the
			// directory can be generated by appending the directory name onto the
			// previous and current path of this directory. Since the directory has
			// no BackupCollection associated with it (part of the criteria for
			// allowing walking directories in this function) there shouldn't be any
			// LocationPath information associated with the directory.
			newP, err := d.params.currentPath.Append(false, entName)
			if err != nil {
				err = clues.Wrap(err, "getting current directory path").
					WithClues(ctx)
				d.progress.errs.AddRecoverable(ctx, err)

				continue
			}

			ctx = clues.Add(ctx, "child_directory_path", newP)

			oldP, err := d.params.prevPath.Append(false, entName)
			if err != nil {
				err = clues.Wrap(err, "getting previous directory path").
					WithClues(ctx)
				d.progress.errs.AddRecoverable(ctx, err)

				continue
			}

			return virtualfs.NewStreamingDirectory(
				entry.Name(),
				&corsoDirectoryIterator{
					ctx: ctx,
					params: snapshotParams{
						currentPath:    newP,
						prevPath:       oldP,
						collection:     nil,
						baseDir:        dir,
						streamBaseEnts: d.params.streamBaseEnts,
					},
					globalExcludeSet: d.globalExcludeSet,
					progress:         d.progress,
				}), nil
		}

		// This entry was either updated or deleted. In either case, the external
		// service notified us about it and it's already been handled so we should
		// skip it here.
		if _, ok := d.seenEnts[entry.Name()]; ok {
			continue
		}

		// This entry was marked as deleted by a service that can't tell us the
		// previous path of deleted items, only the item ID.
		if _, ok := d.excludeSet[entName]; ok {
			continue
		}

		// This is a path used in corso not kopia so it doesn't need to encode the
		// item name.
		itemPath, err := d.params.currentPath.AppendItem(entName)
		if err != nil {
			err = clues.Wrap(err, "getting full item path for base entry").
				WithClues(ctx)
			d.progress.errs.AddRecoverable(ctx, err)

			continue
		}

		ctx = clues.Add(ctx, "item_path", itemPath)

		// We need the previous path so we can find this item in the base snapshot's
		// backup details. If the item moved and we had only the new path, we'd be
		// unable to find it in the old backup details because we wouldn't know what
		// to look for.
		prevItemPath, err := d.params.prevPath.AppendItem(entName)
		if err != nil {
			err = clues.Wrap(err, "getting previous full item path for base entry").
				WithClues(ctx)
			d.progress.errs.AddRecoverable(ctx, err)

			continue
		}

		// Meta files aren't in backup details since it's the set of items the
		// user sees.
		//
		// TODO(ashmrtn): We may eventually want to make this a function that is
		// passed in so that we can more easily switch it between different
		// external service provider implementations.
		if !metadata.IsMetadataFile(itemPath) {
			// All items have item info in the base backup. However, we need to make
			// sure we have enough metadata to find those entries. To do that we add
			// the item to progress and having progress aggregate everything for
			// later.
			detailsEnt := &itemDetails{
				repoPath:     itemPath,
				prevPath:     prevItemPath,
				locationPath: d.locationPath,
				modTime:      ptr.To(entry.ModTime()),
			}
			d.progress.put(
				encodeAsPath(itemPath.PopFront().Elements()...),
				detailsEnt)
		}

		return entry, nil
	}

	return nil, clues.Stack(err).OrNil()
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are virtualfs.StreamingDirectory.
func buildKopiaDirs(
	ctx context.Context,
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
		child, err := buildKopiaDirs(
			ctx,
			childName,
			childDir,
			globalExcludeSet,
			progress)
		if err != nil {
			return nil, err
		}

		childDirs = append(childDirs, child)
	}

	return virtualfs.NewStreamingDirectory(
		encodeAsPath(dirName),
		&corsoDirectoryIterator{
			ctx:              ctx,
			params:           dir.snapshotParams,
			staticEnts:       childDirs,
			globalExcludeSet: globalExcludeSet,
			progress:         progress,
		}), nil
}

type snapshotParams struct {
	// path.Path representing the node's path. This is passed as a parameter to
	// the stream item function so that even baseDir directories can properly
	// generate the full path of items.
	currentPath path.Path
	// Previous path this directory may have resided at if it is sourced from a
	// base snapshot.
	prevPath path.Path

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

	// subtreeChanged denotes whether any directories under this node have been
	// moved, renamed, deleted, or added. If not then we should return both the
	// kopia files and the kopia directories in the base entry because we're also
	// doing selective subtree pruning during hierarchy merging.
	streamBaseEnts bool
}

type treeMap struct {
	snapshotParams
	// Child directories of this directory.
	childDirs map[string]*treeMap
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

type pathUpdate struct {
	p     path.Path
	state data.CollectionState
}

func inflateCollectionTree(
	ctx context.Context,
	collections []data.BackupCollection,
	toMerge *mergeDetails,
) (map[string]*treeMap, map[string]pathUpdate, error) {
	// failed is temporary and just allows us to log all conflicts before
	// returning an error.
	var firstErr error

	roots := make(map[string]*treeMap)
	// Contains the old path for collections that are not new.
	// Allows resolving what the new path should be when walking the base
	// snapshot(s)'s hierarchy. Nil represents a collection that was deleted.
	updatedPaths := make(map[string]pathUpdate)
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
				return nil, nil, clues.NewWC(ictx, "nil previous path on deleted collection")
			}

			changedPaths = append(changedPaths, s.PreviousPath())

			if p, ok := updatedPaths[s.PreviousPath().String()]; ok {
				err := clues.NewWC(ictx, "multiple previous state changes").
					With("updated_path", p, "current_state", data.DeletedState)
				logger.CtxErr(ictx, err).Error("previous path state collision")

				if firstErr == nil {
					firstErr = err
				}
			}

			updatedPaths[s.PreviousPath().String()] = pathUpdate{state: data.DeletedState}

			continue

		case data.MovedState:
			changedPaths = append(changedPaths, s.PreviousPath())

			if p, ok := updatedPaths[s.PreviousPath().String()]; ok {
				err := clues.NewWC(ictx, "multiple previous state changes").
					With("updated_path", p, "current_state", data.MovedState)
				logger.CtxErr(ictx, err).Error("previous path state collision")

				if firstErr == nil {
					firstErr = err
				}
			}

			updatedPaths[s.PreviousPath().String()] = pathUpdate{
				p:     s.FullPath(),
				state: data.MovedState,
			}

			// Only safe when collections are moved since we only need prefix matching
			// if a nested folder's path changed in some way that didn't generate a
			// collection. For that to the be case, the nested folder's path must have
			// changed via one of the ancestor folders being moved. This catches the
			// ancestor folder move.
			if err := addMergeLocation(s, toMerge); err != nil {
				return nil, nil, clues.WrapWC(ictx, err, "adding merge location")
			}

		case data.NotMovedState:
			p := s.PreviousPath().String()
			if p, ok := updatedPaths[p]; ok {
				err := clues.NewWC(ictx, "multiple previous state changes").
					With("updated_path", p, "current_state", data.NotMovedState)
				logger.CtxErr(ictx, err).Error("previous path state collision")

				if firstErr == nil {
					firstErr = err
				}
			}

			updatedPaths[p] = pathUpdate{
				p:     s.FullPath(),
				state: data.NotMovedState,
			}
		}

		if s.FullPath() == nil || len(s.FullPath().Elements()) == 0 {
			return nil, nil, clues.NewWC(ictx, "no identifier for collection")
		}

		node := getTreeNode(roots, s.FullPath().Elements())
		if node == nil {
			return nil, nil, clues.NewWC(ictx, "getting tree node")
		}

		// Make sure there's only a single collection adding items for any given
		// path in the new hierarchy.
		if node.collection != nil {
			return nil, nil, clues.NewWC(ictx, "multiple instances of collection")
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
			err := clues.NewWC(ctx, "conflicting states for collection")
			logger.CtxErr(ctx, err).Error("adding node to tree")

			if firstErr == nil {
				firstErr = err
			}
		}
	}

	return roots, updatedPaths, clues.Stack(firstErr).OrNil()
}

func subtreeChanged(
	roots map[string]*treeMap,
	updatedPaths map[string]pathUpdate,
	oldDirPath *path.Builder,
	currentPath *path.Builder,
) bool {
	// Can't combine with the inner if-block because go evaluates the assignment
	// prior to all conditional checks.
	if currentPath != nil {
		// Either there's an input collection for this path or there's some
		// descendant that has a collection with this path as a prefix.
		//
		// The base traversal code is single-threaded and this check is before
		// processing child base directories so we don't have to worry about
		// something else creating the in-memory subtree for unchanged things.
		// Having only one base per Reason also means we haven't added this branch
		// of the in-memory tree in a base processed before this one.
		if node := maybeGetTreeNode(roots, currentPath.Elements()); node != nil &&
			(node.collection != nil || len(node.childDirs) > 0) {
			return true
		}
	}

	oldPath := oldDirPath.String()

	// We only need to check the old paths here because the check on the in-memory
	// tree above will catch cases where the new path is in the subtree rooted at
	// currentPath. We're mostly concerned with things being moved out of the
	// subtree or deleted within the subtree in this block. Renames, moves within,
	// and moves into the subtree will be caught by the above in-memory tree
	// check.
	//
	// We can ignore exact matches on the old path because they would correspond
	// to deleting the root of the subtree which we can handle as long as
	// everything under the subtree root is also deleted.
	for oldP := range updatedPaths {
		if strings.HasPrefix(oldP, oldPath) && len(oldP) != len(oldPath) {
			return true
		}
	}

	return false
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
	updatedPaths map[string]pathUpdate,
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
		return clues.NewWC(ctx, "base snapshot tree too tall")
	}

	// Wrapper base64 encodes all file and folder names to avoid issues with
	// special characters. Since we're working directly with files and folders
	// from kopia we need to do the decoding here.
	dirName, err := decodeElement(dir.Name())
	if err != nil {
		return clues.WrapWC(ctx, err, "decoding base directory name").
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
		if upb.p == nil {
			currentPath = nil

			stats.Inc(statDel)
		} else {
			// This directory was explicitly mentioned and the new (possibly
			// unchanged) location is in upb.
			currentPath = upb.p.ToBuilder()

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

	// Figure out if the subtree rooted at this directory is either unchanged or
	// completely deleted. We can accomplish this by checking the in-memory tree
	// and the updatedPaths map.
	changed := subtreeChanged(roots, updatedPaths, oldDirPath, currentPath)

	var hasItems bool

	if changed {
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
			return clues.WrapWC(ctx, err, "traversing base directory")
		}
	} else {
		stats.Inc(statPruned)
	}

	// We only need to add this base directory to the tree we're building if it
	// has items in it. The traversal of the directory here just finds
	// subdirectories. This optimization will not be valid if we dynamically
	// determine the subdirectories this directory has when handing items to
	// kopia.
	if currentPath != nil && (hasItems || !changed) {
		// Having this in the if-block has the effect of removing empty directories
		// from backups that have a base snapshot. If we'd like to preserve empty
		// directories across incremental backups, move getting the node outside of
		// the if-block. That will be sufficient to create a StreamingDirectory that
		// kopia will pick up on. Assigning the baseDir of the node should remain
		// in the if-block though as that is an optimization.
		node := getTreeNode(roots, currentPath.Elements())
		if node == nil {
			return clues.NewWC(ctx, "getting tree node")
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

		curP, err := path.PrefixOrPathFromDataLayerPath(currentPath.String(), false)
		if err != nil {
			return clues.NewWC(ctx, "converting current path to path.Path")
		}

		oldP, err := path.PrefixOrPathFromDataLayerPath(oldDirPath.String(), false)
		if err != nil {
			return clues.NewWC(ctx, "converting old path to path.Path")
		}

		node.baseDir = dir
		node.currentPath = curP
		node.prevPath = oldP
		node.streamBaseEnts = !changed
	}

	return nil
}

func logBaseInfo(ctx context.Context, b BackupBase) {
	svcs := map[string]struct{}{}
	cats := map[string]struct{}{}

	for _, r := range b.Reasons {
		svcs[r.Service().String()] = struct{}{}
		cats[r.Category().String()] = struct{}{}
	}

	// Base backup ID and base snapshot ID are already in context clues.
	logger.Ctx(ctx).Infow(
		"using base for backup",
		"services", maps.Keys(svcs),
		"categories", maps.Keys(cats))
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
	// statPruned denotes the number of subtree roots that selective subtree
	// pruning applied to.
	statPruned = "subtrees_pruned"
)

func inflateBaseTree(
	ctx context.Context,
	loader snapshotLoader,
	base BackupBase,
	updatedPaths map[string]pathUpdate,
	roots map[string]*treeMap,
) error {
	bupID := "no_backup_id"
	if base.Backup != nil && len(base.Backup.ID) > 0 {
		bupID = string(base.Backup.ID)
	}

	ctx = clues.Add(
		ctx,
		"base_backup_id", bupID,
		"base_snapshot_id", base.ItemDataSnapshot.ID)

	// Only complete snapshots should be used to source base information.
	// Snapshots for checkpoints will rely on kopia-assisted dedupe to efficiently
	// handle items that were completely uploaded before Corso crashed.
	if len(base.ItemDataSnapshot.IncompleteReason) > 0 {
		logger.Ctx(ctx).Info("skipping incomplete snapshot")
		return nil
	}

	// Some logging to help track things.
	logBaseInfo(ctx, base)

	root, err := loader.SnapshotRoot(base.ItemDataSnapshot)
	if err != nil {
		return clues.WrapWC(ctx, err, "getting snapshot root directory")
	}

	dir, ok := root.(fs.Directory)
	if !ok {
		return clues.NewWC(ctx, "snapshot root is not a directory")
	}

	// For each subtree corresponding to the tuple
	// (resource owner, service, category) merge the directories in the base with
	// what has been reported in the collections we got.
	for _, r := range base.Reasons {
		ictx := clues.Add(
			ctx,
			"subtree_service", r.Service().String(),
			"subtree_category", r.Category().String())

		subtreePath, err := r.SubtreePath()
		if err != nil {
			return clues.WrapWC(ictx, err, "building subtree path")
		}

		// We're starting from the root directory so don't need it in the path.
		pathElems := encodeElements(subtreePath.PopFront().Elements()...)

		ent, err := snapshotfs.GetNestedEntry(ictx, dir, pathElems)
		if err != nil {
			if isErrEntryNotFound(err) {
				logger.CtxErr(ictx, err).Infow("base snapshot missing subtree")
				continue
			}

			return clues.WrapWC(ictx, err, "getting subtree root")
		}

		subtreeDir, ok := ent.(fs.Directory)
		if !ok {
			return clues.WrapWC(ictx, err, "subtree root is not directory")
		}

		// This ensures that a migration on the directory prefix can complete.
		// The prefix is the tenant/service/owner/category set, which remains
		// otherwise unchecked in tree inflation below this point.
		newSubtreePath := subtreePath.ToBuilder()

		if up, ok := updatedPaths[subtreePath.String()]; ok {
			newSubtreePath = up.p.ToBuilder()
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
			stats); err != nil {
			return clues.WrapWC(ictx, err, "traversing base snapshot")
		}

		logger.Ctx(ctx).Infow(
			"merge subtree stats",
			statNoMove, stats.Get(statNoMove),
			statMove, stats.Get(statMove),
			statRecursiveMove, stats.Get(statRecursiveMove),
			statDel, stats.Get(statDel),
			statRecursiveDel, stats.Get(statRecursiveDel),
			statSkipMerge, stats.Get(statSkipMerge),
			statPruned, stats.Get(statPruned))
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
	bases []BackupBase,
	collections []data.BackupCollection,
	globalExcludeSet prefixmatcher.StringSetReader,
	progress *corsoProgress,
) (fs.Directory, error) {
	roots, updatedPaths, err := inflateCollectionTree(ctx, collections, progress.toMerge)
	if err != nil {
		return nil, clues.Wrap(err, "inflating collection tree")
	}

	// Individual backup/snapshot IDs will be logged when merging their hierarchy.
	ctx = clues.Add(ctx, "len_bases", len(bases))

	if len(bases) > 0 {
		logger.Ctx(ctx).Info("merging hierarchies from base backups")
	} else {
		logger.Ctx(ctx).Info("no base backups to merge")
	}

	for _, base := range bases {
		if err = inflateBaseTree(ctx, loader, base, updatedPaths, roots); err != nil {
			return nil, clues.Wrap(err, "inflating base backup tree(s)")
		}
	}

	if len(roots) > 1 {
		return nil, clues.NewWC(ctx, "multiple root directories")
	}

	var res fs.Directory

	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(ctx, dirName, dir, globalExcludeSet, progress)
		if err != nil {
			return nil, err
		}

		res = tmp
	}

	return res, nil
}
