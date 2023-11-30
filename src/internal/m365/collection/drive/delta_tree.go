package drive

import (
	"context"
	"time"

	"github.com/alcionai/clues"

	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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

	// the majority of operations we perform can be handled with
	// a folder ID lookup instead of re-walking the entire tree.
	// Ex: adding a new file to its parent folder.
	folderIDToNode map[string]*nodeyMcNodeFace

	// tombstones don't need to form a tree.
	// We maintain the node data in case we swap back
	// and forth between live and tombstoned states.
	tombstones map[string]*nodeyMcNodeFace

	// it's just a sensible place to store the data, since we're
	// already pushing file additions through the api.
	excludeFileIDs map[string]struct{}
}

func newFolderyMcFolderFace(
	prefix path.Path,
) *folderyMcFolderFace {
	return &folderyMcFolderFace{
		prefix:         prefix,
		folderIDToNode: map[string]*nodeyMcNodeFace{},
		tombstones:     map[string]*nodeyMcNodeFace{},
		excludeFileIDs: map[string]struct{}{},
	}
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
	// only contains the folders starting at and including '/root:'
	prev path.Elements
	// map folderID -> node
	children map[string]*nodeyMcNodeFace
	// items are keyed by item ID
	items map[string]time.Time
	// for special handling protocols around packages
	isPackage bool
}

func newNodeyMcNodeFace(
	parent *nodeyMcNodeFace,
	id, name string,
	prev path.Elements,
	isPackage bool,
) *nodeyMcNodeFace {
	return &nodeyMcNodeFace{
		parent:    parent,
		id:        id,
		name:      name,
		prev:      prev,
		children:  map[string]*nodeyMcNodeFace{},
		items:     map[string]time.Time{},
		isPackage: isPackage,
	}
}

// ---------------------------------------------------------------------------
// folder handling
// ---------------------------------------------------------------------------

// SetFolder adds a node with the following details to the tree.
// If the node already exists with the given ID, the name and parent
// values are updated to match (isPackage is assumed not to change).
func (face *folderyMcFolderFace) SetFolder(
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

	// drive doesn't normally allow the `:` character in folder names.
	// so `root:` is, by default, the only folder that can match this
	// name.  That makes this check a little bit brittle, but generally
	// reliable, since we should always see the root first and can rely
	// on the naming structure.
	if len(parentID) == 0 && name != odConsts.RootPathDir {
		return clues.NewWC(ctx, "non-root folder missing parent id")
	}

	// only set the root node once.
	if name == odConsts.RootPathDir {
		if face.root == nil {
			root := newNodeyMcNodeFace(nil, id, name, nil, isPackage)
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
		// the previous location is always nil, since previous path additions get their
		// own setter func.
		nodey = newNodeyMcNodeFace(parent, id, name, nil, isPackage)
	}

	// ensure the parent points to this node, and that the node is registered
	// in the map of all nodes in the tree.
	parent.children[id] = nodey
	face.folderIDToNode[id] = nodey

	return nil
}

func (face *folderyMcFolderFace) SetTombstone(
	ctx context.Context,
	id string,
	loc path.Elements,
) error {
	if len(id) == 0 {
		return clues.NewWC(ctx, "missing tombstone folder ID")
	}

	if len(loc) == 0 {
		return clues.NewWC(ctx, "missing tombstone location")
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

	zombey, alreadyBuried := face.tombstones[id]
	if alreadyBuried {
		if zombey.prev.String() != loc.String() {
			// logging for sanity
			logger.Ctx(ctx).Infow(
				"attempted to tombstone two paths with the same ID",
				"first_tombstone_path", zombey.prev,
				"second_tombstone_path", loc)
		}

		// since we're storing drive data by folder name in kopia, not id, we need
		// to make sure to preserve the original tombstone location.  If we get a
		// conflicting set of locations in the same delta enumeration, we can always
		// treat the original one as the canonical one.  IE: what we're deleting is
		// the original location as it exists in kopia.  So even if we get a newer
		// location in the drive enumeration, the original location is the one that
		// kopia uses, and the one we need to tombstone.
		//
		// this should also be asserted in the second step, where we compare the delta
		// changes to the backup previous paths metadata.
		face.tombstones[id] = zombey
	} else {
		face.tombstones[id] = newNodeyMcNodeFace(nil, id, "", loc, false)
	}

	return nil
}
