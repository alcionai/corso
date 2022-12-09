package kopia

import (
	"context"
	"encoding/binary"
	"io"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var versionSize = int(unsafe.Sizeof(serializationVersion))

// backupStreamReader is a wrapper around the io.Reader that other Corso
// components return when backing up information. It injects a version number at
// the start of the data stream. Future versions of Corso may not need this if
// they use more complex serialization logic as serialization/version injection
// will be handled by other components.
type backupStreamReader struct {
	io.ReadCloser
	version   uint32
	readBytes int
}

func (rw *backupStreamReader) Read(p []byte) (n int, err error) {
	if rw.readBytes < versionSize {
		marshalled := make([]byte, versionSize)

		toCopy := len(marshalled) - rw.readBytes
		if len(p) < toCopy {
			toCopy = len(p)
		}

		binary.BigEndian.PutUint32(marshalled, rw.version)

		copy(p, marshalled[rw.readBytes:rw.readBytes+toCopy])
		rw.readBytes += toCopy

		return toCopy, nil
	}

	return rw.ReadCloser.Read(p)
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

func collectionEntries(
	ctx context.Context,
	cb func(context.Context, fs.Entry) error,
	streamedEnts data.Collection,
	progress *corsoProgress,
) (map[string]struct{}, *multierror.Error) {
	if streamedEnts == nil {
		return nil, nil
	}

	var (
		errs *multierror.Error
		// Track which items have already been seen so we can skip them if we see
		// them again in the data from the base snapshot.
		seen  = map[string]struct{}{}
		items = streamedEnts.Items()
		log   = logger.Ctx(ctx)
	)

	for {
		select {
		case <-ctx.Done():
			errs = multierror.Append(errs, ctx.Err())
			return seen, errs

		case e, ok := <-items:
			if !ok {
				return seen, errs
			}

			// Even if this item has been deleted and should not appear at all in
			// the new snapshot we need to record that we've seen it here so we know
			// to skip it if it's also present in the base snapshot.
			//
			// TODO(ashmrtn): Determine if we want to try to use the old version of
			// the data (if it exists in the base) if we fail uploading the new
			// version. If so, we should split this call into where we check for the
			// item being deleted and then again after we do the kopia callback.
			seen[e.UUID()] = struct{}{}

			// For now assuming that item IDs don't need escaping.
			itemPath, err := streamedEnts.FullPath().Append(e.UUID(), true)
			if err != nil {
				err = errors.Wrap(err, "getting full item path")
				errs = multierror.Append(errs, err)

				log.Error(err)

				continue
			}

			log.Debugw("reading item", "path", itemPath.String())
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
				d := &itemDetails{info: ei.Info(), repoPath: itemPath}
				progress.put(encodeAsPath(itemPath.PopFront().Elements()...), d)
			}

			modTime := time.Now()
			if smt, ok := e.(data.StreamModTime); ok {
				modTime = smt.ModTime()
			}

			entry := virtualfs.StreamingFileWithModTimeFromReader(
				encodeAsPath(e.UUID()),
				modTime,
				&backupStreamReader{
					version:    serializationVersion,
					ReadCloser: e.ToReader(),
				},
			)
			if err := cb(ctx, entry); err != nil {
				// Kopia's uploader swallows errors in most cases, so if we see
				// something here it's probably a big issue and we should return.
				errs = multierror.Append(errs, errors.Wrapf(err, "executing callback on %q", itemPath))
				return seen, errs
			}
		}
	}
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
		ctx, end := D.Span(ctx, "kopia:getStreamItemFunc")
		defer end()

		// Return static entries in this directory first.
		for _, d := range staticEnts {
			if err := cb(ctx, d); err != nil {
				return errors.Wrap(err, "executing callback on static directory")
			}
		}

		_, errs := collectionEntries(ctx, cb, streamedEnts, progress)

		// TODO(ashmrtn): Stream entries from base snapshot if they exist.

		return errs.ErrorOrNil()
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
	collections []data.Collection,
) (map[string]*treeMap, map[string]path.Path, *OwnersCats, error) {
	roots := make(map[string]*treeMap)
	// Contains the old path for collections that have been moved or renamed.
	// Allows resolving what the new path should be when walking the base
	// snapshot(s)'s hierarchy. Nil represents a collection that was deleted.
	updatedPaths := make(map[string]path.Path)
	ownerCats := &OwnersCats{
		ResourceOwners: make(map[string]struct{}),
		ServiceCats:    make(map[string]struct{}),
	}

	for _, s := range collections {
		switch s.State() {
		case data.DeletedState:
			updatedPaths[s.PreviousPath().String()] = nil
			continue

		case data.MovedState:
			updatedPaths[s.PreviousPath().String()] = s.FullPath()
		}

		if s.FullPath() == nil || len(s.FullPath().Elements()) == 0 {
			return nil, nil, nil, errors.New("no identifier for collection")
		}

		node := getTreeNode(roots, s.FullPath().Elements())
		if node == nil {
			return nil, nil, nil, errors.Errorf(
				"unable to get tree node for path %s",
				s.FullPath(),
			)
		}

		serviceCat := serviceCatTag(s.FullPath())
		ownerCats.ServiceCats[serviceCat] = struct{}{}
		ownerCats.ResourceOwners[s.FullPath().ResourceOwner()] = struct{}{}

		node.collection = s
	}

	return roots, updatedPaths, ownerCats, nil
}

// inflateDirTree returns a set of tags representing all the resource owners and
// service/categories in the snapshot and a fs.Directory tree rooted at the
// oldest common ancestor of the streams. All nodes are
// virtualfs.StreamingDirectory with the given DataCollections if there is one
// for that node. Tags can be used in future backups to fetch old snapshots for
// caching reasons.
func inflateDirTree(
	ctx context.Context,
	collections []data.Collection,
	progress *corsoProgress,
) (fs.Directory, *OwnersCats, error) {
	roots, _, ownerCats, err := inflateCollectionTree(ctx, collections)
	if err != nil {
		return nil, nil, errors.Wrap(err, "inflating collection tree")
	}

	if len(roots) > 1 {
		return nil, nil, errors.New("multiple root directories")
	}

	var res fs.Directory

	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(dirName, dir, progress)
		if err != nil {
			return nil, nil, err
		}

		res = tmp
	}

	return res, ownerCats, nil
}
