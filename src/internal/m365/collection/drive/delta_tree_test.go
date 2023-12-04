package drive

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

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

	folderFace := newFolderyMcFolderFace(p, rootID)
	assert.Equal(t, p, folderFace.prefix)
	assert.Nil(t, folderFace.root)
	assert.NotNil(t, folderFace.folderIDToNode)
	assert.NotNil(t, folderFace.tombstones)
	assert.NotNil(t, folderFace.fileIDToParentID)
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
	assert.NotEqual(t, defaultLoc, nodeFace.prev)
	assert.True(t, nodeFace.isPackage)
	assert.NotNil(t, nodeFace.children)
	assert.NotNil(t, nodeFace.files)
}

// ---------------------------------------------------------------------------
// folder tests
// ---------------------------------------------------------------------------

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
			tname:     "add root",
			tree:      newTree,
			id:        rootID,
			name:      rootName,
			isPackage: true,
			expectErr: assert.NoError,
		},
		{
			tname:     "root already exists",
			tree:      treeWithRoot,
			id:        rootID,
			name:      rootName,
			expectErr: assert.NoError,
		},
		{
			tname:     "add folder",
			tree:      treeWithRoot,
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			expectErr: assert.NoError,
		},
		{
			tname:     "add package",
			tree:      treeWithRoot,
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.NoError,
		},
		{
			tname:     "missing ID",
			tree:      treeWithRoot,
			parentID:  rootID,
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "missing name",
			tree:      treeWithRoot,
			parentID:  rootID,
			id:        id(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "missing parentID",
			tree:      treeWithRoot,
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "already tombstoned",
			tree:      treeWithTombstone,
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			expectErr: assert.NoError,
		},
		{
			tname: "add folder before parent",
			tree: func(t *testing.T) *folderyMcFolderFace {
				return &folderyMcFolderFace{
					folderIDToNode: map[string]*nodeyMcNodeFace{},
				}
			},
			parentID:  rootID,
			id:        id(folder),
			name:      name(folder),
			isPackage: true,
			expectErr: assert.Error,
		},
		{
			tname:     "folder already exists",
			tree:      treeWithFolders,
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

			tree := test.tree(t)

			err := tree.setFolder(
				ctx,
				test.parentID,
				test.id,
				test.name,
				test.isPackage)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := tree.folderIDToNode[test.id]
			require.NotNil(t, result)
			assert.Equal(t, test.id, result.id)
			assert.Equal(t, test.name, result.name)
			assert.Equal(t, test.isPackage, result.isPackage)

			_, ded := tree.tombstones[test.id]
			assert.False(t, ded)

			if len(test.parentID) > 0 {
				parent := tree.folderIDToNode[test.parentID]
				assert.Equal(t, parent, result.parent)
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_AddTombstone() {
	table := []struct {
		name      string
		id        string
		tree      func(t *testing.T) *folderyMcFolderFace
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "add tombstone",
			id:        id(folder),
			tree:      newTree,
			expectErr: assert.NoError,
		},
		{
			name:      "duplicate tombstone",
			id:        id(folder),
			tree:      treeWithTombstone,
			expectErr: assert.NoError,
		},
		{
			name:      "missing ID",
			tree:      newTree,
			expectErr: assert.Error,
		},
		{
			name:      "conflict: folder alive",
			id:        id(folder),
			tree:      treeWithTombstone,
			expectErr: assert.NoError,
		},
		{
			name:      "already tombstoned",
			id:        id(folder),
			tree:      treeWithTombstone,
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			tree := test.tree(t)

			err := tree.setTombstone(ctx, test.id)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := tree.tombstones[test.id]
			require.NotNil(t, result)
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetPreviousPath() {
	pathWith := func(loc path.Elements) path.Path {
		p, err := path.Build(tenant, user, path.OneDriveService, path.FilesCategory, false, loc...)
		require.NoError(suite.T(), err, clues.ToCore(err))

		return p
	}

	table := []struct {
		name            string
		id              string
		prev            path.Path
		tree            func(t *testing.T) *folderyMcFolderFace
		expectErr       assert.ErrorAssertionFunc
		expectLive      bool
		expectTombstone bool
	}{
		{
			name:            "no changes become a no-op",
			id:              id(folder),
			prev:            pathWith(defaultLoc()),
			tree:            newTree,
			expectErr:       assert.NoError,
			expectLive:      false,
			expectTombstone: false,
		},
		{
			name:            "added folders after reset",
			id:              id(folder),
			loc:             loc,
			tree:            treeWithFoldersAfterReset(),
			expectErr:       assert.NoError,
			expectLive:      true,
			expectTombstone: false,
		},
		{
			name:            "create tombstone after reset",
			id:              id(folder),
			prev:            pathWith(defaultLoc()),
			tree:            treeAfterReset,
			expectErr:       assert.NoError,
			expectLive:      false,
			expectTombstone: true,
		},
		{
			name:            "missing ID",
			prev:            pathWith(defaultLoc()),
			tree:            newTree,
			expectErr:       assert.Error,
			expectLive:      false,
			expectTombstone: false,
		},
		{
			name:            "missing prev",
			id:              id(folder),
			tree:            newTree,
			expectErr:       assert.Error,
			expectLive:      false,
			expectTombstone: false,
		},
		{
			name:            "update live folder",
			id:              id(folder),
			prev:            pathWith(defaultLoc()),
			tree:            treeWithFolders,
			expectErr:       assert.NoError,
			expectLive:      true,
			expectTombstone: false,
		},
		{
			name:            "update tombstone",
			id:              id(folder),
			prev:            pathWith(defaultLoc()),
			tree:            treeWithTombstone,
			expectErr:       assert.NoError,
			expectLive:      false,
			expectTombstone: true,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			tree := test.tree(t)

			err := tree.setPreviousPath(test.id, test.prev)
			test.expectErr(t, err, clues.ToCore(err))

			if test.expectLive {
				require.Contains(t, tree.folderIDToNode, test.id)
				assert.Equal(t, test.prev.String(), tree.folderIDToNode[test.id].prev.String())
			} else {
				require.NotContains(t, tree.folderIDToNode, test.id)
			}

			if test.expectTombstone {
				require.Contains(t, tree.tombstones, test.id)
				assert.Equal(t, test.prev, tree.tombstones[test.id].prev)
			} else {
				require.NotContains(t, tree.tombstones, test.id)
			}
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

	tree := treeWithRoot(t)

	set := func(
		parentID, fid, fname string,
		isPackage bool,
	) {
		err := tree.setFolder(ctx, parentID, fid, fname, isPackage)
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

	tree := treeWithRoot(t)

	set := func(
		parentID, fid, fname string,
		isPackage bool,
	) {
		err := tree.setFolder(ctx, parentID, fid, fname, isPackage)
		require.NoError(t, err, clues.ToCore(err))
	}

	tomb := func(
		tid string,
		loc path.Elements,
	) {
		err := tree.setTombstone(ctx, tid)
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

// ---------------------------------------------------------------------------
// file tests
// ---------------------------------------------------------------------------

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_AddFile() {
	table := []struct {
		tname       string
		tree        func(t *testing.T) *folderyMcFolderFace
		oldParentID string
		parentID    string
		contentSize int64
		expectErr   assert.ErrorAssertionFunc
		expectFiles map[string]string
	}{
		{
			tname:       "add file to root",
			tree:        treeWithRoot,
			oldParentID: "",
			parentID:    rootID,
			contentSize: 42,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{id(file): rootID},
		},
		{
			tname:       "add file to folder",
			tree:        treeWithFolders,
			oldParentID: "",
			parentID:    id(folder),
			contentSize: 24,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{id(file): id(folder)},
		},
		{
			tname:       "re-add file at the same location",
			tree:        treeWithFileAtRoot,
			oldParentID: rootID,
			parentID:    rootID,
			contentSize: 84,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{id(file): rootID},
		},
		{
			tname:       "move file from folder to root",
			tree:        treeWithFileInFolder,
			oldParentID: id(folder),
			parentID:    rootID,
			contentSize: 48,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{id(file): rootID},
		},
		{
			tname:       "move file from tombstone to root",
			tree:        treeWithFileInTombstone,
			oldParentID: id(folder),
			parentID:    rootID,
			contentSize: 2,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{id(file): rootID},
		},
		{
			tname:       "error adding file to tombstone",
			tree:        treeWithTombstone,
			oldParentID: "",
			parentID:    id(folder),
			contentSize: 4,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
		{
			tname:       "error adding file before parent",
			tree:        treeWithTombstone,
			oldParentID: "",
			parentID:    idx(folder, 1),
			contentSize: 8,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
		{
			tname:       "error adding file without parent id",
			tree:        treeWithTombstone,
			oldParentID: "",
			parentID:    "",
			contentSize: 16,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()
			tree := test.tree(t)

			err := tree.addFile(
				test.parentID,
				id(file),
				driveFile(file, parentDir(), test.parentID))
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectFiles, tree.fileIDToParentID)

			if err != nil {
				return
			}

			parent := tree.getNode(test.parentID)

			require.NotNil(t, parent)
			assert.Contains(t, parent.files, id(file))

			countSize := tree.countLiveFilesAndSizes()
			assert.Equal(t, 1, countSize.numFiles, "should have one file in the tree")
			assert.Equal(t, test.contentSize, countSize.totalBytes, "tree should be sized to test file contents")

			if len(test.oldParentID) > 0 && test.oldParentID != test.parentID {
				old := tree.GetNode(test.oldParentID)

				require.NotNil(t, old)
				assert.NotContains(t, old.files, id(file))
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_DeleteFile() {
	table := []struct {
		tname    string
		tree     func(t *testing.T) *folderyMcFolderFace
		parentID string
	}{
		{
			tname:    "delete unseen file",
			tree:     treeWithRoot,
			parentID: rootID,
		},
		{
			tname:    "delete file from root",
			tree:     treeWithFolders,
			parentID: rootID,
		},
		{
			tname:    "delete file from folder",
			tree:     treeWithFileInFolder,
			parentID: id(folder),
		},
		{
			tname:    "delete file from tombstone",
			tree:     treeWithFileInTombstone,
			parentID: id(folder),
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()
			tree := test.tree(t)

			tree.deleteFile(id(file))

			parent := tree.getNode(test.parentID)

			require.NotNil(t, parent)
			assert.NotContains(t, parent.files, id(file))
			assert.NotContains(t, tree.fileIDToParentID, id(file))
			assert.Contains(t, tree.deletedFileIDs, id(file))
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_addAndDeleteFile() {
	t := suite.T()
	tree := treeWithRoot()
	fID := id(file)

	require.Len(t, tree.fileIDToParentID, 0)
	require.Len(t, tree.deletedFileIDs, 0)

	tree.deleteFile(fID)

	assert.Len(t, tree.fileIDToParentID, 0)
	assert.NotContains(t, tree.fileIDToParentID, fID)
	assert.Len(t, tree.deletedFileIDs, 1)
	assert.Contains(t, tree.deletedFileIDs, fID)

	err := tree.addFile(rootID, fID, time.Now(), defaultItemSize)
	require.NoError(t, err, clues.ToCore(err))

	assert.Len(t, tree.fileIDToParentID, 1)
	assert.Contains(t, tree.fileIDToParentID, fID)
	assert.Len(t, tree.deletedFileIDs, 0)
	assert.NotContains(t, tree.deletedFileIDs, fID)

	tree.deleteFile(fID)

	assert.Len(t, tree.fileIDToParentID, 0)
	assert.NotContains(t, tree.fileIDToParentID, fID)
	assert.Len(t, tree.deletedFileIDs, 1)
	assert.Contains(t, tree.deletedFileIDs, fID)
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_GenerateExcludeItemIDs() {
	table := []struct {
		name   string
		tree   func(t *testing.T) *folderyMcFolderFace
		expect map[string]struct{}
	}{
		{
			name:   "no files",
			tree:   treeWithRoot,
			expect: map[string]struct{}{},
		},
		{
			name: "one file in a folder",
			tree: treeWithFileInFolder,
			expect: map[string]struct{}{
				id(file): {},
			},
		},
		{
			name:   "one file in a tombstone",
			tree:   treeWithFileInTombstone,
			expect: map[string]struct{}{},
		},
		{
			name: "one deleted file",
			tree: treeWithDeletedFile,
			expect: map[string]struct{}{
				id(file, "d"): {},
			},
		},
		{
			name: "files in folders and tombstones",
			tree: fullTree,
			expect: map[string]struct{}{
				id(file):       {},
				idx(file, "r"): {},
				idx(file, "p"): {},
				idx(file, "d"): {},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			tree := test.tree(t)

			result := tree.generateExcludeItemIDs()
			assert.Equal(t, test.expect, result)
		})
	}
}

// ---------------------------------------------------------------------------
// post-processing tests
// ---------------------------------------------------------------------------

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_GenerateCollectables() {
	t := suite.T()

	table := []struct {
		name      string
		tree      func(t *testing.T) *folderyMcFolderFace
		prevPaths map[string]string
		expectErr require.ErrorAssertionFunc
		expect    map[string]collectable
	}{
		{
			name:      "empty tree",
			tree:      newTree,
			expectErr: require.NoError,
			expect:    map[string]collectable{},
		},
		{
			name:      "root only",
			tree:      treeWithRoot,
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  fullPathPath(t),
					files:                     map[string]models.DriveItemable{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
				},
			},
		},
		{
			name:      "root with files",
			tree:      treeWithFileAtRoot,
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath: fullPathPath(t),
					files: map[string]models.DriveItemable{
						id(file): fileAtRoot(),
					},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
				},
			},
		},
		{
			name:      "folder hierarchy, no previous",
			tree:      treeWithFileInFolder,
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  fullPathPath(t),
					files:                     map[string]models.DriveItemable{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
				},
				idx(folder, "parent"): {
					currPath:                  fullPathPath(t, namex(folder, "parent")),
					files:                     map[string]models.DriveItemable{},
					folderID:                  idx(folder, "parent"),
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{rootName},
				},
				id(folder): {
					currPath: fullPathPath(t, namex(folder, "parent"), name(folder)),
					files: map[string]models.DriveItemable{
						id(file): fileAt("parent"),
					},
					folderID:                  id(folder),
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{rootName, namex(folder, "parent")},
				},
			},
		},
		{
			name: "package in hierarchy",
			tree: func(t *testing.T) *folderyMcFolderFace {
				ctx, flush := tester.NewContext(t)
				defer flush()

				tree := treeWithRoot(t)
				err := tree.setFolder(ctx, rootID, id(pkg), name(pkg), true)
				require.NoError(t, err, clues.ToCore(err))

				err = tree.setFolder(ctx, id(pkg), id(folder), name(folder), false)
				require.NoError(t, err, clues.ToCore(err))

				return tree
			},
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  fullPathPath(t),
					files:                     map[string]models.DriveItemable{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
				},
				id(pkg): {
					currPath:                  fullPathPath(t, name(pkg)),
					files:                     map[string]models.DriveItemable{},
					folderID:                  id(pkg),
					isPackageOrChildOfPackage: true,
					loc:                       path.Elements{rootName},
				},
				id(folder): {
					currPath:                  fullPathPath(t, name(pkg), name(folder)),
					files:                     map[string]models.DriveItemable{},
					folderID:                  id(folder),
					isPackageOrChildOfPackage: true,
					loc:                       path.Elements{rootName, name(pkg)},
				},
			},
		},
		{
			name:      "folder hierarchy with previous paths",
			tree:      treeWithFileInFolder,
			expectErr: require.NoError,
			prevPaths: map[string]string{
				rootID:                fullPath(),
				idx(folder, "parent"): fullPath(namex(folder, "parent-prev")),
				id(folder):            fullPath(namex(folder, "parent-prev"), name(folder)),
			},
			expect: map[string]collectable{
				rootID: {
					currPath:                  fullPathPath(t),
					files:                     map[string]models.DriveItemable{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
					prevPath:                  fullPathPath(t),
				},
				idx(folder, "parent"): {
					currPath:                  fullPathPath(t, namex(folder, "parent")),
					files:                     map[string]models.DriveItemable{},
					folderID:                  idx(folder, "parent"),
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{rootName},
					prevPath:                  fullPathPath(t, namex(folder, "parent-prev")),
				},
				id(folder): {
					currPath:                  fullPathPath(t, namex(folder, "parent"), name(folder)),
					folderID:                  id(folder),
					isPackageOrChildOfPackage: false,
					files: map[string]models.DriveItemable{
						id(file): fileAt("parent"),
					},
					loc:      path.Elements{rootName, namex(folder, "parent")},
					prevPath: fullPathPath(t, namex(folder, "parent-prev"), name(folder)),
				},
			},
		},
		{
			name: "root and tombstones",
			tree: treeWithFileInTombstone,
			prevPaths: map[string]string{
				rootID:     fullPath(),
				id(folder): fullPath(name(folder)),
			},
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  fullPathPath(t),
					files:                     map[string]models.DriveItemable{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					loc:                       path.Elements{},
					prevPath:                  fullPathPath(t),
				},
				id(folder): {
					files:                     map[string]models.DriveItemable{},
					folderID:                  id(folder),
					isPackageOrChildOfPackage: false,
					prevPath:                  fullPathPath(t, name(folder)),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			tree := test.tree(t)

			if len(test.prevPaths) > 0 {
				for id, ps := range test.prevPaths {
					pp, err := path.FromDataLayerPath(ps, false)
					require.NoError(t, err, clues.ToCore(err))

					err = tree.setPreviousPath(id, pp)
					require.NoError(t, err, clues.ToCore(err))
				}
			}

			results, err := tree.generateCollectables()
			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, results, len(test.expect))

			for id, expect := range test.expect {
				require.Contains(t, results, id)

				result := results[id]
				assert.Equal(t, id, result.folderID)

				if expect.currPath == nil {
					assert.Nil(t, result.currPath)
				} else {
					assert.Equal(t, expect.currPath.String(), result.currPath.String())
				}

				if expect.prevPath == nil {
					assert.Nil(t, result.prevPath)
				} else {
					assert.Equal(t, expect.prevPath.String(), result.prevPath.String())
				}

				if expect.loc == nil {
					assert.Nil(t, result.loc)
				} else {
					assert.Equal(t, expect.loc.PlainString(), result.loc.PlainString())
				}

				assert.ElementsMatch(t, maps.Keys(expect.files), maps.Keys(result.files))
			}
		})
	}
}
