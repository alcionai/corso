package drive

import (
	"context"
	"sort"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// folderyMcFolderFace owns our delta processing tree.
type folderyMcFolderFace struct {
	// tenant/service/resource/category/driveID
	// (or whatever variant the service defines)
	// allows the tree to focus only on folder structure,
	// and minimizes the possibility of multi-prefix path bugs.
	prefix path.Path

	// the root of the tree;
	// new, moved, and notMoved root
	root *nodeyMcNodeFace

	// the ID of the actual root folder.
	// required to ensure correct population of the root node.
	rootID string

	// the majority of operations we perform can be handled with
	// a folder ID lookup instead of re-walking the entire tree.
	// Ex: adding a new file to its parent folder.
	folderIDToNode map[string]*nodeyMcNodeFace

	// tombstones don't need to form a tree.
	// We maintain the node data in case we swap back
	// and forth between live and tombstoned states.
	tombstones map[string]*nodeyMcNodeFace

	// will also be used to construct the excluded file id map
	// during the post-processing step
	fileIDToParentID map[string]string
	// required for populating the excluded file id map
	deletedFileIDs map[string]struct{}

	// true if Reset() was called
	hadReset bool
}

func newFolderyMcFolderFace(
	prefix path.Path,
	rootID string,
) *folderyMcFolderFace {
	return &folderyMcFolderFace{
		prefix:           prefix,
		rootID:           rootID,
		folderIDToNode:   map[string]*nodeyMcNodeFace{},
		tombstones:       map[string]*nodeyMcNodeFace{},
		fileIDToParentID: map[string]string{},
		deletedFileIDs:   map[string]struct{}{},
	}
}

// reset erases all data contained in the tree.  This is intended for
// tracking a delta enumeration reset, not for tree re-use, and will
// cause the tree to flag itself as dirty in order to appropriately
// post-process the data.
func (face *folderyMcFolderFace) reset() {
	face.hadReset = true
	face.root = nil
	face.folderIDToNode = map[string]*nodeyMcNodeFace{}
	face.tombstones = map[string]*nodeyMcNodeFace{}
	face.fileIDToParentID = map[string]string{}
}

type nodeyMcNodeFace struct {
	// required for mid-enumeration folder moves, else we have to walk
	// the tree completely to remove the node from its old parent.
	parent *nodeyMcNodeFace
	// the microsoft item ID.  Mostly because we might as well
	// attach that to the node if we're also attaching the dir.
	id string
	// single directory name, not a path
	name string
	// contains the complete previous path
	prev path.Path
	// folderID -> node
	children map[string]*nodeyMcNodeFace
	// file item ID -> file metadata
	files map[string]*custom.DriveItem
	// for special handling protocols around packages
	isPackage bool
}

func newNodeyMcNodeFace(
	parent *nodeyMcNodeFace,
	id, name string,
	isPackage bool,
) *nodeyMcNodeFace {
	return &nodeyMcNodeFace{
		parent:    parent,
		id:        id,
		name:      name,
		children:  map[string]*nodeyMcNodeFace{},
		files:     map[string]*custom.DriveItem{},
		isPackage: isPackage,
	}
}

// ---------------------------------------------------------------------------
// folder handling
// ---------------------------------------------------------------------------

// containsFolder returns true if the given folder id is present as either
// a live node or a tombstone.
func (face *folderyMcFolderFace) containsFolder(id string) bool {
	_, stillKicking := face.folderIDToNode[id]
	_, alreadyBuried := face.tombstones[id]

	return stillKicking || alreadyBuried
}

func (face *folderyMcFolderFace) getNode(id string) *nodeyMcNodeFace {
	if zombey, alreadyBuried := face.tombstones[id]; alreadyBuried {
		return zombey
	}

	return face.folderIDToNode[id]
}

// setFolder adds a node with the following details to the tree.
// If the node already exists with the given ID, the name and parent
// values are updated to match (isPackage is assumed not to change).
func (face *folderyMcFolderFace) setFolder(
	ctx context.Context,
	parentID, id, name string,
	isPackage bool,
) error {
	// need to ensure we have the minimum requirements met for adding a node.
	if len(id) == 0 {
		return clues.NewWC(ctx, "missing folder ID")
	}

	if len(name) == 0 {
		return clues.NewWC(ctx, "missing folder name")
	}

	if len(parentID) == 0 && id != face.rootID {
		return clues.NewWC(ctx, "non-root folder missing parent id")
	}

	// only set the root node once.
	if id == face.rootID {
		if face.root == nil {
			root := newNodeyMcNodeFace(nil, id, name, isPackage)
			face.root = root
			face.folderIDToNode[id] = root
		}

		return nil
	}

	// There are four possible changes that can happen at this point.
	// 1. new folder addition.
	// 2. duplicate folder addition.
	// 3. existing folder migrated to new location.
	// 4. tombstoned folder restored.

	parent, ok := face.folderIDToNode[parentID]
	if !ok {
		return clues.NewWC(ctx, "folder added before parent")
	}

	// Handling case 4 is exclusive to 1-3.  IE: we ensure tree state such
	// that a node's previous appearance can be either a tombstone or
	// a live node, but not both.  So if we find a tombstone, we assume
	// there is not also a node in the live tree for this id.

	// if a folder is deleted and restored, we'll get both the deletion marker
	// (which should be first in enumeration, since all deletion markers are first,
	// or it would have happened in one of the prior enumerations), followed by
	// the restoration of the folder.
	if zombey, tombstoned := face.tombstones[id]; tombstoned {
		delete(face.tombstones, id)

		zombey.parent = parent
		zombey.name = name
		parent.children[id] = zombey
		face.folderIDToNode[id] = zombey

		return nil
	}

	// if not previously a tombstone, handle change cases 1-3
	var (
		nodey   *nodeyMcNodeFace
		visited bool
	)

	// change type 2 & 3.  Update the existing node details to match current data.
	if nodey, visited = face.folderIDToNode[id]; visited {
		if nodey.parent == nil {
			// technically shouldn't be possible but better to keep the problem tracked
			// just in case.
			logger.Ctx(ctx).Info("non-root folder already exists with no parent ref")
		} else if nodey.parent != parent {
			// change type 3.  we need to ensure the old parent stops pointing to this node.
			delete(nodey.parent.children, id)
		}

		nodey.name = name
		nodey.parent = parent
	} else {
		// change type 1: new addition
		nodey = newNodeyMcNodeFace(parent, id, name, isPackage)
	}

	// ensure the parent points to this node, and that the node is registered
	// in the map of all nodes in the tree.
	parent.children[id] = nodey
	face.folderIDToNode[id] = nodey

	return nil
}

func (face *folderyMcFolderFace) setTombstone(
	ctx context.Context,
	id string,
) error {
	if len(id) == 0 {
		return clues.NewWC(ctx, "missing tombstone folder ID")
	}

	// since we run mutiple enumerations, it's possible to see a folder added on the
	// first enumeration that then gets deleted on the next.  This means that the folder
	// was deleted while the first enumeration was in flight, and will show up even if
	// the folder gets restored after deletion.
	// When this happens, we have to adjust the original tree accordingly.
	if zombey, stillKicking := face.folderIDToNode[id]; stillKicking {
		if zombey.parent != nil {
			delete(zombey.parent.children, id)
		}

		delete(face.folderIDToNode, id)

		zombey.parent = nil
		face.tombstones[id] = zombey

		// this handling is exclusive to updating an already-existing tombstone.
		// ie: if we find a living node in the tree, we assume there is no tombstone
		// entry with the same ID.
		return nil
	}

	if _, alreadyBuried := face.tombstones[id]; !alreadyBuried {
		face.tombstones[id] = newNodeyMcNodeFace(nil, id, "", false)
	}

	return nil
}

// setPreviousPath updates the previousPath for the folder with folderID.  If the folder
// already exists either as a tombstone or in the tree, the previous path on those nodes
// gets updated.  Otherwise the previous path update usually gets dropped, because we
// assume no changes have occurred.
// If the tree was Reset() at any point, any previous path that does not still exist in
// the tree- either as a tombstone or a live node- is assumed to have been deleted between
// deltas, and gets turned into a tombstone.
func (face *folderyMcFolderFace) setPreviousPath(
	folderID string,
	prev path.Path,
) error {
	if len(folderID) == 0 {
		return clues.New("missing folder id")
	}

	if prev == nil {
		return clues.New("missing previous path")
	}

	if zombey, isDie := face.tombstones[folderID]; isDie {
		zombey.prev = prev
		return nil
	}

	if nodey, exists := face.folderIDToNode[folderID]; exists {
		nodey.prev = prev
		return nil
	}

	// if no reset occurred, then we assume all previous folder entries are still
	// valid and continue to exist, even without a reference in the tree.  However,
	// if the delta was reset, then it's possible for a folder to be have been deleted
	// and the only way we'd know is if the previous paths map says the folder exists
	// but we haven't seen it again in this enumeration.
	if !face.hadReset {
		return nil
	}

	zombey := newNodeyMcNodeFace(nil, folderID, "", false)
	zombey.prev = prev
	face.tombstones[folderID] = zombey

	return nil
}

// ---------------------------------------------------------------------------
// file handling
// ---------------------------------------------------------------------------

func (face *folderyMcFolderFace) hasFile(id string) bool {
	_, exists := face.fileIDToParentID[id]
	return exists
}

// addFile places the file in the correct parent node.  If the
// file was already added to the tree and is getting relocated,
// this func will update and/or clean up all the old references.
func (face *folderyMcFolderFace) addFile(
	parentID, id string,
	file *custom.DriveItem,
) error {
	if len(parentID) == 0 {
		return clues.New("item added without parent folder ID")
	}

	if len(id) == 0 {
		return clues.New("item added without ID")
	}

	// in case of file movement, clean up any references
	// to the file in the old parent
	oldParentID, ok := face.fileIDToParentID[id]
	if ok && oldParentID != parentID {
		if nodey, ok := face.folderIDToNode[oldParentID]; ok {
			delete(nodey.files, id)
		}

		if zombey, ok := face.tombstones[oldParentID]; ok {
			delete(zombey.files, id)
		}
	}

	parent, ok := face.folderIDToNode[parentID]
	if !ok {
		return clues.New("item added before parent")
	}

	face.fileIDToParentID[id] = parentID
	parent.files[id] = file

	delete(face.deletedFileIDs, id)

	return nil
}

func (face *folderyMcFolderFace) deleteFile(id string) {
	parentID, ok := face.fileIDToParentID[id]
	if ok {
		if nodey, ok := face.folderIDToNode[parentID]; ok {
			delete(nodey.files, id)
		}

		if zombey, ok := face.tombstones[parentID]; ok {
			delete(zombey.files, id)
		}
	}

	delete(face.fileIDToParentID, id)

	face.deletedFileIDs[id] = struct{}{}
}

// ---------------------------------------------------------------------------
// post-processing
// ---------------------------------------------------------------------------

type collectable struct {
	currPath                  path.Path
	files                     map[string]*custom.DriveItem
	folderID                  string
	isPackageOrChildOfPackage bool
	prevPath                  path.Path
}

// produces a map of folderID -> collectable
func (face *folderyMcFolderFace) generateCollectables() (map[string]collectable, error) {
	result := map[string]collectable{}

	err := walkTreeAndBuildCollections(
		face.root,
		face.prefix,
		&path.Builder{},
		false, true,
		result)

	for id, tombstone := range face.tombstones {
		// in case we got a folder deletion marker for a folder
		// that has no previous path, drop the entry entirely.
		// it doesn't exist in storage, so there's nothing to delete.
		if tombstone.prev != nil {
			result[id] = collectable{
				folderID: id,
				prevPath: tombstone.prev,
			}
		}
	}

	return result, clues.Stack(err).OrNil()
}

func walkTreeAndBuildCollections(
	node *nodeyMcNodeFace,
	pathPfx path.Path,
	parentPath *path.Builder,
	isChildOfPackage, isRoot bool,
	result map[string]collectable,
) error {
	if node == nil {
		return nil
	}

	if !isRoot {
		parentPath = parentPath.Append(node.name)
	}

	for _, child := range node.children {
		err := walkTreeAndBuildCollections(
			child,
			pathPfx,
			parentPath,
			node.isPackage || isChildOfPackage,
			false,
			result)
		if err != nil {
			return err
		}
	}

	collectionPath, err := pathPfx.Append(false, parentPath.Elements()...)
	if err != nil {
		return clues.Wrap(err, "building collection path").
			With(
				"path_prefix", pathPfx,
				"path_suffix", parentPath.Elements())
	}

	cbl := collectable{
		currPath:                  collectionPath,
		files:                     node.files,
		folderID:                  node.id,
		isPackageOrChildOfPackage: node.isPackage || isChildOfPackage,
		prevPath:                  node.prev,
	}

	result[node.id] = cbl

	return nil
}

type idPrevPathTup struct {
	id       string
	prevPath string
}

// fuses the collectables and old prevPaths into a
// new prevPaths map.
func (face *folderyMcFolderFace) generateNewPreviousPaths(
	collectables map[string]collectable,
	prevPaths map[string]string,
) (map[string]string, error) {
	var (
		// id -> currentPath
		results = map[string]string{}
		// prevPath -> currentPath
		movedPaths = map[string]string{}
		// prevPath -> {}
		tombstoned = map[string]struct{}{}
	)

	// first, move all collectables into the new maps

	for id, cbl := range collectables {
		if cbl.currPath == nil {
			tombstoned[cbl.prevPath.String()] = struct{}{}
			continue
		}

		cp := cbl.currPath.String()
		results[id] = cp

		if cbl.prevPath != nil && cbl.prevPath.String() != cp {
			movedPaths[cbl.prevPath.String()] = cp
		}
	}

	// next, create a slice of tuples representing any
	// old prevPath entry whose ID isn't already bound to
	// a collectable.

	unseenPrevPaths := []idPrevPathTup{}

	for id, p := range prevPaths {
		// if the current folder was tombstoned, skip it
		if _, ok := tombstoned[p]; ok {
			continue
		}

		if _, ok := results[id]; !ok {
			unseenPrevPaths = append(unseenPrevPaths, idPrevPathTup{id, p})
		}
	}

	// sort the slice by path, ascending.
	// This ensures we work from root to leaf when replacing prefixes,
	// and thus we won't need to walk every unseen path from leaf to
	// root looking for a matching prefix.

	sortByLeastPath := func(i, j int) bool {
		return unseenPrevPaths[i].prevPath < unseenPrevPaths[j].prevPath
	}

	sort.Slice(unseenPrevPaths, sortByLeastPath)

	for _, un := range unseenPrevPaths {
		elems := path.NewElements(un.prevPath)

		pb, err := path.Builder{}.UnescapeAndAppend(elems...)
		if err != nil {
			return nil, err
		}

		parent := pb.Dir().String()

		// if the parent was tombstoned, add this prevPath entry to the
		// tombstoned map; that'll allow the tombstone identification to
		// cascade to children, and it won't get added to the results.
		if _, ok := tombstoned[parent]; ok {
			tombstoned[un.prevPath] = struct{}{}
			continue
		}

		// if the parent wasn't moved, add the same path to the result set
		parentCurrentPath, ok := movedPaths[parent]
		if !ok {
			results[un.id] = un.prevPath
			continue
		}

		// if the parent was moved, replace the prefix and
		// add it to the result set
		// TODO: should probably use path.UpdateParent for this.
		// but I want the quality-of-life of feeding it strings
		// instead of parsing strings to paths here first.
		newPath := strings.Replace(un.prevPath, parent, parentCurrentPath, 1)

		results[un.id] = newPath

		// add the current string to the moved list, that'll allow it to cascade to all children.
		movedPaths[un.prevPath] = newPath
	}

	return results, nil
}

func (face *folderyMcFolderFace) generateExcludeItemIDs() map[string]struct{} {
	result := map[string]struct{}{}

	for iID, pID := range face.fileIDToParentID {
		if _, itsAlive := face.folderIDToNode[pID]; !itsAlive {
			// don't worry about items whose parents are tombstoned.
			// those will get handled in the delete cascade.
			continue
		}

		result[iID+metadata.DataFileSuffix] = struct{}{}
		result[iID+metadata.MetaFileSuffix] = struct{}{}
	}

	for iID := range face.deletedFileIDs {
		result[iID+metadata.DataFileSuffix] = struct{}{}
		result[iID+metadata.MetaFileSuffix] = struct{}{}
	}

	return result
}

// ---------------------------------------------------------------------------
// quantification
// ---------------------------------------------------------------------------

// countLiveFolders returns a count of the number of folders held in the tree.
// Tombstones are not included in the count.  Only live folders.
func (face *folderyMcFolderFace) countLiveFolders() int {
	return len(face.folderIDToNode)
}

type countAndSize struct {
	numFiles   int
	totalBytes int64
}

// countLiveFilesAndSizes returns a count of the number of files in the tree
// and the sum of all of their sizes.  Only includes files that are not
// children of tombstoned containers.  If running an incremental backup, a
// live file may be either a creation or an update.
func (face *folderyMcFolderFace) countLiveFilesAndSizes() countAndSize {
	return countFilesAndSizes(face.root)
}

func countFilesAndSizes(nodey *nodeyMcNodeFace) countAndSize {
	if nodey == nil {
		return countAndSize{}
	}

	var (
		fileCount      int
		sumContentSize int64
	)

	for _, child := range nodey.children {
		countSize := countFilesAndSizes(child)
		fileCount += countSize.numFiles
		sumContentSize += countSize.totalBytes
	}

	for _, file := range nodey.files {
		sumContentSize += ptr.Val(file.GetSize())
	}

	return countAndSize{
		numFiles:   fileCount + len(nodey.files),
		totalBytes: sumContentSize,
	}
}
