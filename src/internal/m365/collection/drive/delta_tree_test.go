package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

var loc = path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")

func treeWithRoot() *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(nil)
	rootey := newNodeyMcNodeFace(nil, rootID, rootName, false)
	tree.root = rootey
	tree.folderIDToNode[rootID] = rootey

	return tree
}

func treeWithTombstone() *folderyMcFolderFace {
	tree := treeWithRoot()
	tree.tombstones[id(folder)] = newNodeyMcNodeFace(nil, id(folder), "", false)

	return tree
}

func treeWithFolders() *folderyMcFolderFace {
	tree := treeWithRoot()

	o := newNodeyMcNodeFace(tree.root, idx(folder, "parent"), namex(folder, "parent"), true)
	tree.folderIDToNode[o.id] = o

	f := newNodeyMcNodeFace(o, id(folder), name(folder), false)
	tree.folderIDToNode[f.id] = f
	o.children[f.id] = f

	return tree
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type DeltaTreeUnitSuite struct {
	tester.Suite
}

func TestDeltaTreeUnitSuite(t *testing.T) {
	suite.Run(t, &DeltaTreeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DeltaTreeUnitSuite) TestNewFolderyMcFolderFace() {
	var (
		t      = suite.T()
		p, err = path.BuildPrefix(tenant, user, path.OneDriveService, path.FilesCategory)
	)

	require.NoError(t, err, clues.ToCore(err))

	folderFace := newFolderyMcFolderFace(p)
	assert.Equal(t, p, folderFace.prefix)
	assert.Nil(t, folderFace.root)
	assert.NotNil(t, folderFace.folderIDToNode)
	assert.NotNil(t, folderFace.tombstones)
	assert.NotNil(t, folderFace.excludeFileIDs)
}

func (suite *DeltaTreeUnitSuite) TestNewNodeyMcNodeFace() {
	var (
		t      = suite.T()
		parent = &nodeyMcNodeFace{}
	)

	nodeFace := newNodeyMcNodeFace(parent, "id", "name", true)
	assert.Equal(t, parent, nodeFace.parent)
	assert.Equal(t, "id", nodeFace.id)
	assert.Equal(t, "name", nodeFace.name)
	assert.NotEqual(t, loc, nodeFace.prev)
	assert.True(t, nodeFace.isPackage)
	assert.NotNil(t, nodeFace.children)
	assert.NotNil(t, nodeFace.items)
}

// note that this test is focused on the SetFolder function,
// and intentionally does not verify the resulting node tree
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder() {
	table := []struct {
		tname     string
		tree      *folderyMcFolderFace
		parentID  string
		id        string
		name      string
		isPackage bool
		expectErr assert.ErrorAssertionFunc
	}{
		{
			tname: "add root",
			tree: &folderyMcFolderFace{
				folderIDToNode: map[string]*nodeyMcNodeFace{},
			},
			id:        rootID,
			name:      rootName,
			isPackage: true,
			expectErr: assert.NoError,
		},
		{
			tname:     "root already exists",
			tree:      treeWithRoot(),
			id:        rootID,
			name:      rootName,
			expectErr: assert.NoError,
		},
		{
			tname:     "add folder",
			tree:      treeWithRoot(),
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			expectErr: assert.NoError,
		},
		{
			tname:     "add package",
			tree:      treeWithRoot(),
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.NoError,
		},
		{
			tname:     "missing ID",
			tree:      treeWithRoot(),
			parentID:  rootID,
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "missing name",
			tree:      treeWithRoot(),
			parentID:  rootID,
			id:        id(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "missing parentID",
			tree:      treeWithRoot(),
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "already tombstoned",
			tree:      treeWithTombstone(),
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			expectErr: assert.NoError,
		},
		{
			tname: "add folder before parent",
			tree: &folderyMcFolderFace{
				folderIDToNode: map[string]*nodeyMcNodeFace{},
			},
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "folder already exists",
			tree:      treeWithFolders(),
			parentID:  idx(folder, "parent"),
			id:        id(folder),
			name:      name(folder),
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			err := test.tree.SetFolder(
				ctx,
				test.parentID,
				test.id,
				test.name,
				test.isPackage)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := test.tree.folderIDToNode[test.id]
			require.NotNil(t, result)
			assert.Equal(t, test.id, result.id)
			assert.Equal(t, test.name, result.name)
			assert.Equal(t, test.isPackage, result.isPackage)

			_, ded := test.tree.tombstones[test.id]
			assert.False(t, ded)

			if len(test.parentID) > 0 {
				parent := test.tree.folderIDToNode[test.parentID]
				assert.Equal(t, parent, result.parent)
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_AddTombstone() {
	table := []struct {
		name      string
		id        string
		tree      *folderyMcFolderFace
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "add tombstone",
			id:        id(folder),
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.NoError,
		},
		{
			name:      "duplicate tombstone",
			id:        id(folder),
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
		{
			name:      "missing ID",
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.Error,
		},
		{
			name:      "conflict: folder alive",
			id:        id(folder),
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
		{
			name:      "already tombstoned",
			id:        id(folder),
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			err := test.tree.SetTombstone(ctx, test.id)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := test.tree.tombstones[test.id]
			require.NotNil(t, result)
		})
	}
}

// ---------------------------------------------------------------------------
// tree structure assertions tests
// ---------------------------------------------------------------------------

type assertNode struct {
	self     *nodeyMcNodeFace
	children []assertNode
}

func an(
	self *nodeyMcNodeFace,
	children ...assertNode,
) assertNode {
	return assertNode{
		self:     self,
		children: children,
	}
}

func (an assertNode) compare(
	t *testing.T,
	tree *folderyMcFolderFace,
	checkLiveNodeCount bool,
) {
	var nodeCount int

	t.Run("assert_tree_shape/root", func(_t *testing.T) {
		nodeCount = compareNodes(_t, tree.root, an)
	})

	if checkLiveNodeCount {
		require.Len(t, tree.folderIDToNode, nodeCount, "total count of live nodes")
	}
}

func compareNodes(
	t *testing.T,
	node *nodeyMcNodeFace,
	expect assertNode,
) int {
	// ensure the nodes match
	require.NotNil(t, node, "node does not exist in tree")
	require.Equal(
		t,
		expect.self,
		node,
		"non-matching node")

	// ensure the node has the expected number of children
	assert.Len(
		t,
		node.children,
		len(expect.children),
		"node has expected number of children")

	var nodeCount int

	for _, expectChild := range expect.children {
		t.Run(expectChild.self.id, func(_t *testing.T) {
			nodeChild := node.children[expectChild.self.id]
			require.NotNilf(
				_t,
				nodeChild,
				"child must exist with id %q",
				expectChild.self.id)

			// ensure each child points to the current node as its parent
			assert.Equal(
				_t,
				node,
				nodeChild.parent,
				"should point to correct parent")

			// recurse
			nodeCount += compareNodes(_t, nodeChild, expectChild)
		})
	}

	return nodeCount + 1
}

type tombs []assertNode

func entomb(nodes ...assertNode) tombs {
	if len(nodes) == 0 {
		return tombs{}
	}

	return append(tombs{}, nodes...)
}

func (ts tombs) compare(
	t *testing.T,
	tombstones map[string]*nodeyMcNodeFace,
) {
	require.Len(t, tombstones, len(ts), "count of tombstoned nodes")

	for _, entombed := range ts {
		zombey := tombstones[entombed.self.id]
		require.NotNil(t, zombey, "tombstone must exist")
		assert.Nil(t, zombey.parent, "tombstoned node should not have a parent reference")

		t.Run("assert_tombstones/"+zombey.id, func(_t *testing.T) {
			compareNodes(_t, zombey, entombed)
		})
	}
}

// unlike the prior test, this focuses entirely on whether or not the
// tree produced by folder additions is correct.
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder_correctTree() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	tree := treeWithRoot()

	set := func(
		parentID, fid, fname string,
		isPackage bool,
	) {
		err := tree.SetFolder(ctx, parentID, fid, fname, isPackage)
		require.NoError(t, err, clues.ToCore(err))
	}

	// assert the root exists

	assert.NotNil(t, tree.root)
	assert.Equal(t, rootID, tree.root.id)
	assert.Equal(t, rootID, tree.folderIDToNode[rootID].id)

	an(tree.root).compare(t, tree, true)

	// add a child at the root
	set(rootID, id("lefty"), name("l"), false)

	lefty := tree.folderIDToNode[id("lefty")]
	an(
		tree.root,
		an(lefty)).
		compare(t, tree, true)

	// add another child at the root
	set(rootID, id("righty"), name("r"), false)

	righty := tree.folderIDToNode[id("righty")]
	an(
		tree.root,
		an(lefty),
		an(righty)).
		compare(t, tree, true)

	// add a child to lefty
	set(lefty.id, id("bloaty"), name("bl"), false)

	bloaty := tree.folderIDToNode[id("bloaty")]
	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty)).
		compare(t, tree, true)

	// add another child to lefty
	set(lefty.id, id("brightly"), name("br"), false)

	brightly := tree.folderIDToNode[id("brightly")]
	an(
		tree.root,
		an(lefty, an(bloaty), an(brightly)),
		an(righty)).
		compare(t, tree, true)

	// relocate brightly underneath righty
	set(righty.id, brightly.id, brightly.name, false)

	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty, an(brightly))).
		compare(t, tree, true)

	// relocate righty and subtree beneath lefty
	set(lefty.id, righty.id, righty.name, false)

	an(
		tree.root,
		an(
			lefty,
			an(bloaty),
			an(righty, an(brightly)))).
		compare(t, tree, true)
}

// this test focuses on whether the tree is correct when bouncing back and forth
// between live and tombstoned states on the same folder
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder_correctTombstones() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	tree := treeWithRoot()

	set := func(
		parentID, fid, fname string,
		isPackage bool,
	) {
		err := tree.SetFolder(ctx, parentID, fid, fname, isPackage)
		require.NoError(t, err, clues.ToCore(err))
	}

	tomb := func(
		tid string,
		loc path.Elements,
	) {
		err := tree.SetTombstone(ctx, tid)
		require.NoError(t, err, clues.ToCore(err))
	}

	// create a simple tree
	// root > branchy > [leafy, bob]
	set(tree.root.id, id("branchy"), name("br"), false)
	branchy := tree.folderIDToNode[id("branchy")]

	set(branchy.id, id("leafy"), name("l"), false)
	set(branchy.id, id("bob"), name("bobbers"), false)

	leafy := tree.folderIDToNode[id("leafy")]
	bob := tree.folderIDToNode[id("bob")]

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	var (
		branchyLoc = path.NewElements("root/branchy")
		leafyLoc   = path.NewElements("root/branchy/leafy")
		bobLoc     = path.NewElements("root/branchy/bob")
	)

	// tombstone bob
	tomb(bob.id, bobLoc)
	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// tombstone leafy
	tomb(leafy.id, leafyLoc)
	an(
		tree.root,
		an(branchy)).
		compare(t, tree, true)

	entomb(an(bob), an(leafy)).compare(t, tree.tombstones)

	// resurrect leafy
	set(branchy.id, leafy.id, leafy.name, false)

	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(branchy.id, bob.id, bob.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(branchy.id, branchyLoc)

	an(tree.root).compare(t, tree, false)
	// note: the folder count here *will be wrong*.
	// since we've only tombstoned branchy, both leafy
	// and bob will remain in the folderIDToNode map.
	// If this were real graph behavior, the delete would
	// necessarily cascade and those children would get
	// tombstoned next.
	// So we skip the check here, just to minimize code.
	// It's safe to do so, for the scope of this test.
	// This should be part of the consideration for prev-
	// path iteration that could create improper state in
	// the post-processing stage if we're nott careful.

	entomb(
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree.tombstones)

	// resurrect branchy
	set(tree.root.id, branchy.id, branchy.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(branchy.id, branchyLoc)

	an(tree.root).compare(t, tree, false)

	entomb(
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree.tombstones)

	// tombstone bob
	tomb(bob.id, bobLoc)

	an(tree.root).compare(t, tree, false)

	entomb(
		an(branchy, an(leafy)),
		an(bob)).
		compare(t, tree.tombstones)

	// resurrect branchy
	set(tree.root.id, branchy.id, branchy.name, false)

	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, false)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(branchy.id, bob.id, bob.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)
}
