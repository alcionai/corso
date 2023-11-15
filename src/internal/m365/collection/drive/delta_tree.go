package drive

import (
	"time"

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
	// new, moved, and notMoved collections
	collections *nodeyMcNodeFace

	// the majority of operations we perform can be handled with
	// a folder id lookup instead of re-walking the entire tree.
	// Ex: adding a new file to its parent folder.
	folderIDToNode map[string]*nodeyMcNodeFace

	// tombstones don't need to form a tree.
	// we only need the folder ID and their previous path.
	tombstones map[string]path.Path

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
		tombstones:     map[string]path.Path{},
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
	prev path.Path
	// map folderID -> node
	childDirs map[string]*nodeyMcNodeFace
	// items are keyed by item ID
	items map[string]time.Time
	// for special handling protocols around packages
	isPackage bool
}

func newNodeyMcNodeFace(
	parent *nodeyMcNodeFace,
	id, name string,
	prev path.Path,
	isPackage bool,
) *nodeyMcNodeFace {
	return &nodeyMcNodeFace{
		parent:    parent,
		id:        id,
		name:      name,
		prev:      prev,
		childDirs: map[string]*nodeyMcNodeFace{},
		items:     map[string]time.Time{},
		isPackage: isPackage,
	}
}
