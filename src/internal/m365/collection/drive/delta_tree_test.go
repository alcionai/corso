package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type DeltaTreeUnitSuite struct {
	tester.Suite
}

func TestDeltaTreeUnitSuite(t *testing.T) {
	suite.Run(t, &DeltaTreeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DeltaTreeUnitSuite) TestNewFolderyMcFolderFace() {
	var (
		t      = suite.T()
		p, err = path.BuildPrefix("t", "r", path.OneDriveService, path.FilesCategory)
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
		loc    = path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")
	)

	nodeFace := newNodeyMcNodeFace(parent, "id", "name", loc, true)
	assert.Equal(t, parent, nodeFace.parent)
	assert.Equal(t, "id", nodeFace.id)
	assert.Equal(t, "name", nodeFace.name)
	assert.Equal(t, loc, nodeFace.prev)
	assert.True(t, nodeFace.isPackage)
	assert.NotNil(t, nodeFace.children)
	assert.NotNil(t, nodeFace.items)
}

// note that this test is focused on the SetFolder function,
// and intentionally does not verify the resulting node tree
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder() {
	treeWithRoot := func() *folderyMcFolderFace {
		rootey := newNodeyMcNodeFace(nil, odConsts.RootID, odConsts.RootPathDir, nil, false)
		tree := newFolderyMcFolderFace(nil)
		tree.root = rootey
		tree.folderIDToNode[odConsts.RootID] = rootey

		return tree
	}

	treeWithTombstone := func() *folderyMcFolderFace {
		tree := treeWithRoot()
		tree.tombstones["folder"] = newNodeyMcNodeFace(nil, "folder", "", path.NewElements(""), false)

		return tree
	}

	treeWithFolders := func() *folderyMcFolderFace {
		tree := treeWithRoot()

		o := newNodeyMcNodeFace(tree.root, "other", "o", nil, true)
		tree.folderIDToNode[o.id] = o

		f := newNodeyMcNodeFace(o, "folder", "f", nil, false)
		tree.folderIDToNode[f.id] = f
		o.children[f.id] = f

		return tree
	}

	table := []struct {
		tname      string
		tree       *folderyMcFolderFace
		parentID   string
		id         string
		name       string
		isPackage  bool
		expectErr  assert.ErrorAssertionFunc
		expectPrev assert.ValueAssertionFunc
	}{
		{
			tname: "add root",
			tree: &folderyMcFolderFace{
				folderIDToNode: map[string]*nodeyMcNodeFace{},
			},
			id:         odConsts.RootID,
			name:       odConsts.RootPathDir,
			isPackage:  true,
			expectErr:  assert.NoError,
			expectPrev: assert.Nil,
		},
		{
			tname:      "root already exists",
			tree:       treeWithRoot(),
			id:         odConsts.RootID,
			name:       odConsts.RootPathDir,
			expectErr:  assert.NoError,
			expectPrev: assert.Nil,
		},
		{
			tname:      "add folder",
			tree:       treeWithRoot(),
			parentID:   odConsts.RootID,
			id:         "folder",
			name:       "nameyMcNameFace",
			expectErr:  assert.NoError,
			expectPrev: assert.Nil,
		},
		{
			tname:      "add package",
			tree:       treeWithRoot(),
			parentID:   odConsts.RootID,
			id:         "folder",
			name:       "nameyMcNameFace",
			isPackage:  true,
			expectErr:  assert.NoError,
			expectPrev: assert.Nil,
		},
		{
			tname:      "missing ID",
			tree:       treeWithRoot(),
			parentID:   odConsts.RootID,
			name:       "nameyMcNameFace",
			isPackage:  true,
			expectErr:  assert.Error,
			expectPrev: assert.Nil,
		},
		{
			tname:      "missing name",
			tree:       treeWithRoot(),
			parentID:   odConsts.RootID,
			id:         "folder",
			isPackage:  true,
			expectErr:  assert.Error,
			expectPrev: assert.Nil,
		},
		{
			tname:      "missing parentID",
			tree:       treeWithRoot(),
			id:         "folder",
			name:       "nameyMcNameFace",
			isPackage:  true,
			expectErr:  assert.Error,
			expectPrev: assert.Nil,
		},
		{
			tname:      "already tombstoned",
			tree:       treeWithTombstone(),
			parentID:   odConsts.RootID,
			id:         "folder",
			name:       "nameyMcNameFace",
			expectErr:  assert.NoError,
			expectPrev: assert.NotNil,
		},
		{
			tname: "add folder before parent",
			tree: &folderyMcFolderFace{
				folderIDToNode: map[string]*nodeyMcNodeFace{},
			},
			parentID:   odConsts.RootID,
			id:         "folder",
			name:       "nameyMcNameFace",
			isPackage:  true,
			expectErr:  assert.Error,
			expectPrev: assert.Nil,
		},
		{
			tname:      "folder already exists",
			tree:       treeWithFolders(),
			parentID:   "other",
			id:         "folder",
			name:       "nameyMcNameFace",
			expectErr:  assert.NoError,
			expectPrev: assert.Nil,
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
			test.expectPrev(t, result.prev)
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
	loc := path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")
	treeWithTombstone := func() *folderyMcFolderFace {
		tree := newFolderyMcFolderFace(nil)
		tree.tombstones["id"] = newNodeyMcNodeFace(nil, "id", "", loc, false)

		return tree
	}

	table := []struct {
		name      string
		id        string
		loc       path.Elements
		tree      *folderyMcFolderFace
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "add tombstone",
			id:        "id",
			loc:       loc,
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.NoError,
		},
		{
			name:      "duplicate tombstone",
			id:        "id",
			loc:       loc,
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
		{
			name:      "missing ID",
			loc:       loc,
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.Error,
		},
		{
			name:      "missing loc",
			id:        "id",
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.Error,
		},
		{
			name:      "empty loc",
			id:        "id",
			loc:       path.Elements{},
			tree:      newFolderyMcFolderFace(nil),
			expectErr: assert.Error,
		},
		{
			name:      "conflict: folder alive",
			id:        "id",
			loc:       loc,
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
		{
			name:      "already tombstoned with different path",
			id:        "id",
			loc:       append(path.Elements{"foo"}, loc...),
			tree:      treeWithTombstone(),
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			err := test.tree.SetTombstone(ctx, test.id, test.loc)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := test.tree.tombstones[test.id]
			require.NotNil(t, result)
			require.NotEmpty(t, result.prev)
			assert.Equal(t, loc, result.prev)
		})
	}
}

// ---------------------------------------------------------------------------
// tree sructure assertions tests
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

	tree := newFolderyMcFolderFace(nil)
	rootID := odConsts.RootID

	set := func(
		parentID, id, name string,
		isPackage bool,
	) {
		err := tree.SetFolder(ctx, parentID, id, name, isPackage)
		require.NoError(t, err, clues.ToCore(err))
	}

	assert.Nil(t, tree.root)
	assert.Empty(t, tree.folderIDToNode)

	// add the root
	set("", rootID, odConsts.RootPathDir, false)

	assert.NotNil(t, tree.root)
	assert.Equal(t, rootID, tree.root.id)
	assert.Equal(t, rootID, tree.folderIDToNode[rootID].id)

	an(tree.root).compare(t, tree, true)

	// add a child at the root
	set(tree.root.id, "lefty", "l", false)

	lefty := tree.folderIDToNode["lefty"]
	an(
		tree.root,
		an(lefty),
	).compare(t, tree, true)

	// add another child at the root
	set(tree.root.id, "righty", "r", false)

	righty := tree.folderIDToNode["righty"]
	an(
		tree.root,
		an(lefty),
		an(righty),
	).compare(t, tree, true)

	// add a child to lefty
	set(lefty.id, "bloaty", "bl", false)

	bloaty := tree.folderIDToNode["bloaty"]
	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty),
	).compare(t, tree, true)

	// add a child to righty
	set(lefty.id, "brightly", "br", false)

	brightly := tree.folderIDToNode["brightly"]
	an(
		tree.root,
		an(lefty, an(bloaty), an(brightly)),
		an(righty),
	).compare(t, tree, true)

	// relocate brightly underneath righty
	set(righty.id, brightly.id, brightly.name, false)

	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty, an(brightly)),
	).compare(t, tree, true)

	// relocate righty and subtree beneath lefty
	set(lefty.id, righty.id, righty.name, false)

	an(
		tree.root,
		an(
			lefty,
			an(bloaty),
			an(righty, an(brightly))),
	).compare(t, tree, true)
}

// this test focuses on whether the tree is correct when bouncing back and forth
// between live and tombstoned states on the same folder
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder_correctTombstones() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	tree := newFolderyMcFolderFace(nil)
	rootID := odConsts.RootID

	set := func(
		parentID, id, name string,
		isPackage bool,
	) {
		err := tree.SetFolder(ctx, parentID, id, name, isPackage)
		require.NoError(t, err, clues.ToCore(err))
	}

	tomb := func(
		id string,
		loc path.Elements,
	) {
		err := tree.SetTombstone(ctx, id, loc)
		require.NoError(t, err, clues.ToCore(err))
	}

	assert.Nil(t, tree.root)
	assert.Empty(t, tree.folderIDToNode)

	// create a simple tree
	// root > branchy > [leafy, bob]
	set("", rootID, odConsts.RootPathDir, false)

	set(tree.root.id, "branchy", "br", false)
	branchy := tree.folderIDToNode["branchy"]

	set(branchy.id, "leafy", "l", false)
	set(branchy.id, "bob", "bobbers", false)

	leafy := tree.folderIDToNode["leafy"]
	bob := tree.folderIDToNode["bob"]

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob)),
	).compare(t, tree, true)

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
		an(branchy, an(leafy)),
	).compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// tombstone leafy
	tomb(leafy.id, leafyLoc)
	an(
		tree.root,
		an(branchy),
	).compare(t, tree, true)

	entomb(an(bob), an(leafy)).compare(t, tree.tombstones)

	// resurrect leafy
	set(branchy.id, leafy.id, leafy.name, false)

	an(
		tree.root,
		an(branchy, an(leafy)),
	).compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(branchy.id, bob.id, bob.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob)),
	).compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(branchy.id, branchyLoc)

	an(
		tree.root,
	).compare(t, tree, false)
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
			an(bob)),
	).compare(t, tree.tombstones)

	// resurrect branchy
	set(tree.root.id, branchy.id, branchy.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob)),
	).compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(branchy.id, branchyLoc)

	an(
		tree.root,
	).compare(t, tree, false)

	entomb(
		an(
			branchy,
			an(leafy),
			an(bob)),
	).compare(t, tree.tombstones)

	// tombstone bob
	tomb(bob.id, bobLoc)

	an(
		tree.root,
	).compare(t, tree, false)

	entomb(
		an(branchy, an(leafy)),
		an(bob),
	).compare(t, tree.tombstones)

	// resurrect branchy
	set(tree.root.id, branchy.id, branchy.name, false)

	an(
		tree.root,
		an(branchy, an(leafy)),
	).compare(t, tree, false)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(branchy.id, bob.id, bob.name, false)

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob)),
	).compare(t, tree, true)

	entomb().compare(t, tree.tombstones)
}
