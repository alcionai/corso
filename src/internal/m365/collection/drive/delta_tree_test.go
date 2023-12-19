package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
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
		d      = drive()
		fld    = custom.ToCustomDriveItem(d.folderAt(root))
	)

	nodeFace := newNodeyMcNodeFace(parent, fld)
	assert.Equal(t, parent, nodeFace.parent)
	assert.Equal(t, folderID(), ptr.Val(nodeFace.folder.GetId()))
	assert.Equal(t, folderName(), ptr.Val(nodeFace.folder.GetName()))
	assert.Nil(t, nodeFace.prev)
	assert.NotNil(t, nodeFace.folder.GetParentReference())
	assert.NotNil(t, nodeFace.children)
	assert.NotNil(t, nodeFace.files)
}

// ---------------------------------------------------------------------------
// folder tests
// ---------------------------------------------------------------------------

// note that this test is focused on the SetFolder function,
// and intentionally does not verify the resulting node tree
func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_SetFolder() {
	d := drive()

	table := []struct {
		tname     string
		tree      func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		folder    func() *custom.DriveItem
		expectErr assert.ErrorAssertionFunc
	}{
		{
			tname: "add root",
			tree:  newTree,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(rootFolder())
			},
			expectErr: assert.NoError,
		},
		{
			tname: "root already exists",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(rootFolder())
			},
			expectErr: assert.NoError,
		},
		{
			tname: "add folder",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(d.folderAt(root))
			},
			expectErr: assert.NoError,
		},
		{
			tname: "add package",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(d.packageAtRoot())
			},
			expectErr: assert.NoError,
		},
		{
			tname: "missing ID",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				far := d.folderAt(root)
				far.SetId(nil)

				return custom.ToCustomDriveItem(far)
			},
			expectErr: assert.Error,
		},
		{
			tname: "missing name",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				far := d.folderAt(root)
				far.SetName(nil)

				return custom.ToCustomDriveItem(far)
			},
			expectErr: assert.Error,
		},
		{
			tname: "missing parent",
			tree:  treeWithRoot,
			folder: func() *custom.DriveItem {
				far := d.folderAt(root)
				far.SetParentReference(nil)

				return custom.ToCustomDriveItem(far)
			},
			expectErr: assert.Error,
		},
		{
			tname: "already tombstoned",
			tree:  treeWithTombstone,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(d.folderAt(root))
			},
			expectErr: assert.NoError,
		},
		{
			tname: "add folder before parent",
			tree:  newTree,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(d.folderAt(root))
			},
			expectErr: assert.Error,
		},
		{
			tname: "folder already exists",
			tree:  treeWithFolders,
			folder: func() *custom.DriveItem {
				return custom.ToCustomDriveItem(d.folderAt(root))
			},
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			tree := test.tree(t, drive())
			folder := test.folder()

			err := tree.setFolder(ctx, folder)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			result := tree.folderIDToNode[ptr.Val(folder.GetId())]
			require.NotNil(t, result)

			var (
				expectID        = ptr.Val(folder.GetId())
				expectName      = ptr.Val(folder.GetName())
				expectIsPackage = folder.GetPackageEscaped() == nil
				resultID        = ptr.Val(result.folder.GetId())
				resultName      = ptr.Val(result.folder.GetName())
				resultIsPackage = result.folder.GetPackageEscaped() == nil
			)

			assert.Equal(t, expectID, resultID)
			assert.Equal(t, expectName, resultName)
			assert.Equal(t, expectIsPackage, resultIsPackage)

			_, ded := tree.tombstones[expectID]
			assert.False(t, ded)

			if folder.GetParentReference() != nil {
				expectParentID := ptr.Val(folder.GetParentReference().GetId())
				parent := tree.folderIDToNode[expectParentID]
				assert.Equal(t, parent, result.parent)
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_AddTombstone() {
	table := []struct {
		name      string
		id        string
		tree      func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "add tombstone",
			id:        folderID(),
			tree:      newTree,
			expectErr: assert.NoError,
		},
		{
			name:      "duplicate tombstone",
			id:        folderID(),
			tree:      treeWithTombstone,
			expectErr: assert.NoError,
		},
		{
			name:      "missing ID",
			tree:      newTree,
			expectErr: assert.Error,
		},
		{
			name:      "folder exists and is alive",
			id:        folderID(),
			tree:      treeWithTombstone,
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			d := drive()
			tree := test.tree(t, d)
			tomb := delItem(test.id, rootID, isFolder)

			err := tree.setTombstone(ctx, custom.ToCustomDriveItem(tomb))
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
		tree            func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		expectErr       assert.ErrorAssertionFunc
		expectLive      bool
		expectTombstone bool
	}{
		{
			name:            "no changes become a no-op",
			id:              folderID(),
			prev:            pathWith(defaultLoc()),
			tree:            newTree,
			expectErr:       assert.NoError,
			expectLive:      false,
			expectTombstone: false,
		},
		{
			name:            "added folders after reset",
			id:              id(folder),
			prev:            pathWith(defaultLoc()),
			tree:            treeWithFoldersAfterReset,
			expectErr:       assert.NoError,
			expectLive:      true,
			expectTombstone: false,
		},
		{
			name:            "create tombstone after reset",
			id:              folderID(),
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
			id:              folderID(),
			tree:            newTree,
			expectErr:       assert.Error,
			expectLive:      false,
			expectTombstone: false,
		},
		{
			name:            "update live folder",
			id:              folderID(),
			prev:            pathWith(defaultLoc()),
			tree:            treeWithFolders,
			expectErr:       assert.NoError,
			expectLive:      true,
			expectTombstone: false,
		},
		{
			name:            "update tombstone",
			id:              folderID(),
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
			tree := test.tree(t, drive())

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

	t.Run("assert_tree_shape-root", func(_t *testing.T) {
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
		expectID := ptr.Val(expectChild.self.folder.GetId())

		t.Run(expectID, func(_t *testing.T) {
			nodeChild := node.children[expectID]
			require.NotNilf(
				_t,
				nodeChild,
				"child must exist with id %q",
				expectID)

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
		expectID := ptr.Val(entombed.self.folder.GetId())

		zombey := tombstones[expectID]
		require.NotNil(t, zombey, "tombstone must exist")
		assert.Nil(t, zombey.parent, "tombstoned node should not have a parent reference")

		resultID := ptr.Val(zombey.folder.GetId())

		t.Run("assert_tombstones-"+resultID, func(_t *testing.T) {
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

	d := drive()
	tree := treeWithRoot(t, d)

	set := func(folder *custom.DriveItem) {
		err := tree.setFolder(ctx, folder)
		require.NoError(t, err, clues.ToCore(err))
	}

	idOf := func(node *nodeyMcNodeFace) string {
		return ptr.Val(node.folder.GetId())
	}

	customFolder := func(parent, self string) *custom.DriveItem {
		var di models.DriveItemable

		if parent == rootID {
			di = d.folderAt(root, self)
		} else {
			di = driveFolder(
				d.dir("doesn't matter for this test"),
				folderID(parent),
				self)
		}

		return custom.ToCustomDriveItem(di)
	}

	// assert the root exists

	assert.NotNil(t, tree.root)
	assert.Equal(t, rootID, idOf(tree.root))
	assert.Equal(t, rootID, idOf(tree.folderIDToNode[rootID]))

	an(tree.root).compare(t, tree, true)

	// add a child at the root
	set(customFolder(rootID, "lefty"))

	lefty := tree.folderIDToNode[folderID("lefty")]
	an(
		tree.root,
		an(lefty)).
		compare(t, tree, true)

	// add another child at the root
	set(customFolder(rootID, "righty"))

	righty := tree.folderIDToNode[folderID("righty")]
	an(
		tree.root,
		an(lefty),
		an(righty)).
		compare(t, tree, true)

	// add a child to lefty
	set(customFolder("lefty", "bloaty"))

	bloaty := tree.folderIDToNode[folderID("bloaty")]
	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty)).
		compare(t, tree, true)

	// add another child to lefty
	set(customFolder("lefty", "brightly"))

	brightly := tree.folderIDToNode[folderID("brightly")]
	an(
		tree.root,
		an(lefty, an(bloaty), an(brightly)),
		an(righty)).
		compare(t, tree, true)

	// relocate brightly underneath righty
	set(customFolder("righty", "brightly"))

	an(
		tree.root,
		an(lefty, an(bloaty)),
		an(righty, an(brightly))).
		compare(t, tree, true)

	// relocate righty and subtree beneath lefty
	set(customFolder("lefty", "righty"))

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

	d := drive()
	tree := treeWithRoot(t, d)

	set := func(folder *custom.DriveItem) {
		err := tree.setFolder(ctx, folder)
		require.NoError(t, err, clues.ToCore(err))
	}

	customFolder := func(parent, self string) *custom.DriveItem {
		var di models.DriveItemable

		if parent == rootID {
			di = d.folderAt(root, self)
		} else {
			di = driveFolder(
				d.dir("doesn't matter for this test"),
				folderID(parent),
				self)
		}

		return custom.ToCustomDriveItem(di)
	}

	tomb := func(folder *custom.DriveItem) {
		err := tree.setTombstone(ctx, folder)
		require.NoError(t, err, clues.ToCore(err))
	}

	// create a simple tree
	// root > branchy > [leafy, bob]
	set(customFolder(rootID, "branchy"))

	branchy := tree.folderIDToNode[folderID("branchy")]

	set(customFolder("branchy", "leafy"))
	set(customFolder("branchy", "bob"))

	leafy := tree.folderIDToNode[folderID("leafy")]
	bob := tree.folderIDToNode[folderID("bob")]

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone bob
	tomb(customFolder("branchy", "bob"))
	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// tombstone leafy
	tomb(customFolder("branchy", "leafy"))
	an(
		tree.root,
		an(branchy)).
		compare(t, tree, true)

	entomb(an(bob), an(leafy)).compare(t, tree.tombstones)

	// resurrect leafy
	set(customFolder("branchy", "leafy"))

	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, true)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(customFolder("branchy", "bob"))

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(customFolder(rootID, "branchy"))

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
	set(customFolder(rootID, "branchy"))

	an(
		tree.root,
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree, true)

	entomb().compare(t, tree.tombstones)

	// tombstone branchy
	tomb(customFolder(rootID, "branchy"))

	an(tree.root).compare(t, tree, false)

	entomb(
		an(
			branchy,
			an(leafy),
			an(bob))).
		compare(t, tree.tombstones)

	// tombstone bob
	tomb(customFolder("branchy", "bob"))

	an(tree.root).compare(t, tree, false)

	entomb(
		an(branchy, an(leafy)),
		an(bob)).
		compare(t, tree.tombstones)

	// resurrect branchy
	set(customFolder(rootID, "branchy"))

	an(
		tree.root,
		an(branchy, an(leafy))).
		compare(t, tree, false)

	entomb(an(bob)).compare(t, tree.tombstones)

	// resurrect bob
	set(customFolder("branchy", "bob"))

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
		tree        func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		id          string
		oldParentID string
		parent      any
		contentSize int64
		expectErr   assert.ErrorAssertionFunc
		expectFiles map[string]string
	}{
		{
			tname:       "add file to root",
			tree:        treeWithRoot,
			id:          fileID(),
			oldParentID: "",
			parent:      root,
			contentSize: defaultFileSize,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{fileID(): rootID},
		},
		{
			tname:       "add file to folder",
			tree:        treeWithFolders,
			id:          fileID(),
			oldParentID: "",
			parent:      folder,
			contentSize: 24,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{fileID(): folderID()},
		},
		{
			tname:       "re-add file at the same location",
			tree:        treeWithFileAtRoot,
			id:          fileID(),
			oldParentID: rootID,
			parent:      root,
			contentSize: 84,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{fileID(): rootID},
		},
		{
			tname:       "move file from folder to root",
			tree:        treeWithFileInFolder,
			id:          fileID(),
			oldParentID: folderID(),
			parent:      root,
			contentSize: 48,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{fileID(): rootID},
		},
		{
			tname:       "move file from tombstone to root",
			tree:        treeWithFileInTombstone,
			id:          fileID(),
			oldParentID: folderID(),
			parent:      root,
			contentSize: 2,
			expectErr:   assert.NoError,
			expectFiles: map[string]string{fileID(): rootID},
		},
		{
			tname:       "adding file with no ID",
			tree:        treeWithTombstone,
			id:          "",
			oldParentID: "",
			parent:      folder,
			contentSize: 4,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
		{
			tname:       "error adding file to tombstone",
			tree:        treeWithTombstone,
			id:          fileID(),
			oldParentID: "",
			parent:      folder,
			contentSize: 8,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
		{
			tname:       "error adding file before parent",
			tree:        treeWithTombstone,
			id:          fileID(),
			oldParentID: "",
			parent:      "not-in-tree",
			contentSize: 16,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
		{
			tname:       "error adding file without parent id",
			tree:        treeWithTombstone,
			id:          fileID(),
			oldParentID: "",
			parent:      nil,
			contentSize: 16,
			expectErr:   assert.Error,
			expectFiles: map[string]string{},
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			var (
				t    = suite.T()
				d    = drive()
				tree = test.tree(t, d)
				df   = custom.ToCustomDriveItem(d.fileWSizeAt(test.contentSize, test.parent))
			)

			err := tree.addFile(df)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectFiles, tree.fileIDToParentID)

			if err != nil {
				return
			}

			parentID := folderID(test.parent)
			if test.parent == root {
				parentID = rootID
			}

			parent := tree.getNode(parentID)

			require.NotNil(t, parent)
			assert.Contains(t, parent.files, fileID())

			countSize := tree.countLiveFilesAndSizes()
			assert.Equal(t, 1, countSize.numFiles, "should have one file in the tree")
			assert.Equal(t, test.contentSize, countSize.totalBytes, "tree should be sized to test file contents")

			if len(test.oldParentID) > 0 && test.oldParentID != parentID {
				old := tree.getNode(test.oldParentID)

				require.NotNil(t, old)
				assert.NotContains(t, old.files, fileID())
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_DeleteFile() {
	table := []struct {
		tname    string
		tree     func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
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
			parentID: folderID(),
		},
		{
			tname:    "delete file from tombstone",
			tree:     treeWithFileInTombstone,
			parentID: folderID(),
		},
	}
	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()
			tree := test.tree(t, drive())

			tree.deleteFile(fileID())

			parent := tree.getNode(test.parentID)

			require.NotNil(t, parent)
			assert.NotContains(t, parent.files, fileID())
			assert.NotContains(t, tree.fileIDToParentID, fileID())
			assert.Contains(t, tree.deletedFileIDs, fileID())
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_addAndDeleteFile() {
	var (
		t    = suite.T()
		d    = drive()
		tree = treeWithRoot(t, d)
		fID  = fileID()
	)

	require.Len(t, tree.fileIDToParentID, 0)
	require.Len(t, tree.deletedFileIDs, 0)

	tree.deleteFile(fID)

	assert.Len(t, tree.fileIDToParentID, 0)
	assert.NotContains(t, tree.fileIDToParentID, fID)
	assert.Len(t, tree.deletedFileIDs, 1)
	assert.Contains(t, tree.deletedFileIDs, fID)

	err := tree.addFile(custom.ToCustomDriveItem(d.fileAt(root)))
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
		tree   func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		expect map[string]struct{}
	}{
		{
			name:   "no files",
			tree:   treeWithRoot,
			expect: map[string]struct{}{},
		},
		{
			name:   "one file in a folder",
			tree:   treeWithFileInFolder,
			expect: makeExcludeMap(fileID()),
		},
		{
			name:   "one file in a tombstone",
			tree:   treeWithFileInTombstone,
			expect: map[string]struct{}{},
		},
		{
			name:   "one deleted file",
			tree:   treeWithDeletedFile,
			expect: makeExcludeMap(fileID("d")),
		},
		{
			name: "files in folders and tombstones",
			tree: fullTree,
			expect: makeExcludeMap(
				fileID("r"),
				fileID("p"),
				fileID("f"),
				fileID("d")),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			tree := test.tree(t, drive())

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
	d := drive()

	table := []struct {
		name      string
		tree      func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
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
					currPath:                  d.fullPath(t),
					files:                     map[string]*custom.DriveItem{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
				},
			},
		},
		{
			name:      "root with files",
			tree:      treeWithFileAtRoot,
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath: d.fullPath(t),
					files: map[string]*custom.DriveItem{
						fileID(): custom.ToCustomDriveItem(d.fileAt(root)),
					},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
				},
			},
		},
		{
			name:      "folder hierarchy, no previous",
			tree:      treeWithFileInFolder,
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  d.fullPath(t),
					files:                     map[string]*custom.DriveItem{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
				},
				folderID("parent"): {
					currPath: d.fullPath(t, folderName("parent")),
					files: map[string]*custom.DriveItem{
						folderID("parent"): custom.ToCustomDriveItem(d.folderAt(root)),
					},
					folderID:                  folderID("parent"),
					isPackageOrChildOfPackage: false,
				},
				folderID(): {
					currPath: d.fullPath(t, folderName("parent"), folderName()),
					files: map[string]*custom.DriveItem{
						folderID(): custom.ToCustomDriveItem(d.folderAt("parent")),
						fileID():   custom.ToCustomDriveItem(d.fileAt(folder)),
					},
					folderID:                  folderID(),
					isPackageOrChildOfPackage: false,
				},
			},
		},
		{
			name: "package in hierarchy",
			tree: func(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
				ctx, flush := tester.NewContext(t)
				defer flush()

				tree := treeWithRoot(t, d)
				err := tree.setFolder(ctx, custom.ToCustomDriveItem(d.packageAtRoot()))
				require.NoError(t, err, clues.ToCore(err))

				err = tree.setFolder(ctx, custom.ToCustomDriveItem(d.folderAt(pkg)))
				require.NoError(t, err, clues.ToCore(err))

				return tree
			},
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  d.fullPath(t),
					files:                     map[string]*custom.DriveItem{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
				},
				folderID(pkg): {
					currPath: d.fullPath(t, folderName(pkg)),
					files: map[string]*custom.DriveItem{
						folderID(pkg): custom.ToCustomDriveItem(d.packageAtRoot()),
					},
					folderID:                  folderID(pkg),
					isPackageOrChildOfPackage: true,
				},
				folderID(): {
					currPath: d.fullPath(t, folderName(pkg), folderName()),
					files: map[string]*custom.DriveItem{
						folderID(): custom.ToCustomDriveItem(d.folderAt("parent")),
					},
					folderID:                  folderID(),
					isPackageOrChildOfPackage: true,
				},
			},
		},
		{
			name:      "folder hierarchy with previous paths",
			tree:      treeWithFileInFolder,
			expectErr: require.NoError,
			prevPaths: map[string]string{
				rootID:             d.strPath(t),
				folderID("parent"): d.strPath(t, folderName("parent-prev")),
				folderID():         d.strPath(t, folderName("parent-prev"), folderName()),
			},
			expect: map[string]collectable{
				rootID: {
					currPath:                  d.fullPath(t),
					files:                     map[string]*custom.DriveItem{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					prevPath:                  d.fullPath(t),
				},
				folderID("parent"): {
					currPath: d.fullPath(t, folderName("parent")),
					files: map[string]*custom.DriveItem{
						folderID("parent"): custom.ToCustomDriveItem(d.folderAt(root)),
					},
					folderID:                  folderID("parent"),
					isPackageOrChildOfPackage: false,
					prevPath:                  d.fullPath(t, folderName("parent-prev")),
				},
				folderID(): {
					currPath:                  d.fullPath(t, folderName("parent"), folderName()),
					folderID:                  folderID(),
					isPackageOrChildOfPackage: false,
					files: map[string]*custom.DriveItem{
						folderID(): custom.ToCustomDriveItem(d.folderAt("parent")),
						fileID():   custom.ToCustomDriveItem(d.fileAt(folder)),
					},
					prevPath: d.fullPath(t, folderName("parent-prev"), folderName()),
				},
			},
		},
		{
			name: "root and tombstones",
			tree: treeWithFileInTombstone,
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expectErr: require.NoError,
			expect: map[string]collectable{
				rootID: {
					currPath:                  d.fullPath(t),
					files:                     map[string]*custom.DriveItem{},
					folderID:                  rootID,
					isPackageOrChildOfPackage: false,
					prevPath:                  d.fullPath(t),
				},
				folderID(): {
					files:                     map[string]*custom.DriveItem{},
					folderID:                  folderID(),
					isPackageOrChildOfPackage: false,
					prevPath:                  d.fullPath(t, folderName()),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			tree := test.tree(t, d)

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

				assert.ElementsMatch(t, maps.Keys(expect.files), maps.Keys(result.files))
			}
		})
	}
}

func (suite *DeltaTreeUnitSuite) TestFolderyMcFolderFace_GenerateNewPreviousPaths() {
	t := suite.T()
	d := drive()

	table := []struct {
		name         string
		collectables map[string]collectable
		prevPaths    map[string]string
		expect       map[string]string
	}{
		{
			name:         "empty collectables, empty prev paths",
			collectables: map[string]collectable{},
			prevPaths:    map[string]string{},
			expect:       map[string]string{},
		},
		{
			name:         "empty collectables",
			collectables: map[string]collectable{},
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expect: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
		},
		{
			name: "empty prev paths",
			collectables: map[string]collectable{
				rootID:     {currPath: d.fullPath(t)},
				folderID(): {currPath: d.fullPath(t, folderName())},
			},
			prevPaths: map[string]string{},
			expect: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
		},
		{
			name: "collectables replace old prev as new location",
			collectables: map[string]collectable{
				rootID: {currPath: d.fullPath(t)},
				folderID(): {
					prevPath: d.fullPath(t, folderName("old")),
					currPath: d.fullPath(t, folderName()),
				},
			},
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName("old")),
			},
			expect: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
		},
		{
			name: "children of parents not moved maintain location",
			collectables: map[string]collectable{
				rootID: {currPath: d.fullPath(t)},
				folderID(): {
					prevPath: d.fullPath(t, folderName()),
					currPath: d.fullPath(t, folderName()),
				},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName()),
				folderID("c1"): d.strPath(t, folderName(), folderName("c1")),
			},
			expect: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName()),
				folderID("c1"): d.strPath(t, folderName(), folderName("c1")),
			},
		},
		{
			name: "updates cascade to unseen children",
			collectables: map[string]collectable{
				rootID: {currPath: d.fullPath(t)},
				folderID(): {
					prevPath: d.fullPath(t, folderName("old")),
					currPath: d.fullPath(t, folderName()),
				},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("old")),
				folderID("c1"): d.strPath(t, folderName("old"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("old"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("old"), folderName("c2"), folderName("c3")),
			},
			expect: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName()),
				folderID("c1"): d.strPath(t, folderName(), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName(), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName(), folderName("c2"), folderName("c3")),
			},
		},
		{
			name: "updates cascade to unseen children - escaped path separator",
			collectables: map[string]collectable{
				rootID: {currPath: d.fullPath(t)},
				folderID(): {
					prevPath: d.fullPath(t, folderName("o/ld")),
					currPath: d.fullPath(t, folderName("n/ew")),
				},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("o/ld")),
				folderID("c1"): d.strPath(t, folderName("o/ld"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("o/ld"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("o/ld"), folderName("c2"), folderName("c3")),
			},
			expect: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("n/ew")),
				folderID("c1"): d.strPath(t, folderName("n/ew"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("n/ew"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("n/ew"), folderName("c2"), folderName("c3")),
			},
		},
		{
			name: "tombstoned directories get removed",
			collectables: map[string]collectable{
				rootID:     {currPath: d.fullPath(t)},
				folderID(): {prevPath: d.fullPath(t, folderName("old"))},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("old")),
				folderID("c1"): d.strPath(t, folderName("old"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("old"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("old"), folderName("c2"), folderName("c3")),
			},
			expect: map[string]string{
				rootID: d.strPath(t),
			},
		},
		{
			name: "mix of moved and tombstoned",
			collectables: map[string]collectable{
				rootID: {currPath: d.fullPath(t)},
				folderID(): {
					prevPath: d.fullPath(t, folderName("old")),
					currPath: d.fullPath(t, folderName()),
				},
				folderID("c3"): {prevPath: d.fullPath(t, folderName("old"), folderName("c2"), folderName("c3"))},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("old")),
				folderID("c1"): d.strPath(t, folderName("old"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("old"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("old"), folderName("c2"), folderName("c3")),
				folderID("c4"): d.strPath(t, folderName("old"), folderName("c2"), folderName("c3"), folderName("c4")),
			},
			expect: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName()),
				folderID("c1"): d.strPath(t, folderName(), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName(), folderName("c2")),
			},
		},
		{
			// tests the equivalent of:
			// mv    root:/foo         -> root:/bar
			// mkdir root:/foo
			// mkdir root:/foo/c1
			// mv    root:/bar/c1/c2   -> root:/foo/c1/c2
			name: "moved and replaced with same name",
			collectables: map[string]collectable{
				rootID: {
					prevPath: d.fullPath(t),
					currPath: d.fullPath(t),
				},
				folderID(): {
					prevPath: d.fullPath(t, folderName("foo")),
					currPath: d.fullPath(t, folderName("bar")),
				},
				folderID(2): {
					currPath: d.fullPath(t, folderName("foo")),
				},
				folderID("2c1"): {
					currPath: d.fullPath(t, folderName("foo"), folderName("c1")),
				},
				folderID("c2"): {
					prevPath: d.fullPath(t, folderName("foo"), folderName("c1"), folderName("c2")),
					currPath: d.fullPath(t, folderName("foo"), folderName("c1"), folderName("c2")),
				},
			},
			prevPaths: map[string]string{
				rootID:         d.strPath(t),
				folderID():     d.strPath(t, folderName("foo")),
				folderID("c1"): d.strPath(t, folderName("foo"), folderName("c1")),
				folderID("c2"): d.strPath(t, folderName("foo"), folderName("c1"), folderName("c2")),
				folderID("c3"): d.strPath(t, folderName("foo"), folderName("c1"), folderName("c2"), folderName("c3")),
			},
			expect: map[string]string{
				rootID:          d.strPath(t),
				folderID():      d.strPath(t, folderName("bar")),
				folderID("c1"):  d.strPath(t, folderName("bar"), folderName("c1")),
				folderID(2):     d.strPath(t, folderName("foo")),
				folderID("2c1"): d.strPath(t, folderName("foo"), folderName("c1")),
				folderID("c2"):  d.strPath(t, folderName("foo"), folderName("c1"), folderName("c2")),
				folderID("c3"):  d.strPath(t, folderName("foo"), folderName("c1"), folderName("c2"), folderName("c3")),
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tree := newTree(t, d)

			results, err := tree.generateNewPreviousPaths(
				test.collectables,
				test.prevPaths)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, results)
		})
	}
}
