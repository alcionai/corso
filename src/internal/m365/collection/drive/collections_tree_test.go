package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	countTD "github.com/alcionai/corso/src/pkg/count/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

type CollectionsTreeUnitSuite struct {
	tester.Suite
}

func TestCollectionsTreeUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionsTreeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionsTreeUnitSuite) TestCollections_MakeDriveTombstones() {
	badPfxMBH := mock.DefaultOneDriveBH(user)
	badPfxMBH.PathPrefixErr = assert.AnError

	twostones := map[string]struct{}{
		"t1": {},
		"t2": {},
	}

	table := []struct {
		name       string
		tombstones map[string]struct{}
		c          *Collections
		expectErr  assert.ErrorAssertionFunc
		expect     assert.ValueAssertionFunc
	}{
		{
			name:       "nil",
			tombstones: nil,
			c:          collWithMBH(mock.DefaultOneDriveBH(user)),
			expectErr:  assert.NoError,
			expect:     assert.Empty,
		},
		{
			name:       "none",
			tombstones: map[string]struct{}{},
			c:          collWithMBH(mock.DefaultOneDriveBH(user)),
			expectErr:  assert.NoError,
			expect:     assert.Empty,
		},
		{
			name:       "some tombstones",
			tombstones: twostones,
			c:          collWithMBH(mock.DefaultOneDriveBH(user)),
			expectErr:  assert.NoError,
			expect:     assert.NotEmpty,
		},
		{
			name:       "bad prefix path",
			tombstones: twostones,
			c:          collWithMBH(badPfxMBH),
			expectErr:  assert.Error,
			expect:     assert.Empty,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			colls, err := test.c.makeDriveTombstones(ctx, test.tombstones, fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
			test.expect(t, colls)

			for _, coll := range colls {
				assert.Equal(t, data.DeletedState, coll.State(), "tombstones should always delete data")
			}
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_MakeMetadataCollections() {
	badMetaPfxMBH := mock.DefaultOneDriveBH(user)
	badMetaPfxMBH.MetadataPathPrefixErr = assert.AnError

	table := []struct {
		name   string
		c      *Collections
		expect assert.ValueAssertionFunc
	}{
		{
			name:   "no errors",
			c:      collWithMBH(mock.DefaultOneDriveBH(user)),
			expect: assert.NotEmpty,
		},
		{
			name:   "bad prefix path",
			c:      collWithMBH(badMetaPfxMBH),
			expect: assert.Empty,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t           = suite.T()
				deltaTokens = map[string]string{}
				prevPaths   = map[string]map[string]string{}
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			colls := test.c.makeMetadataCollections(ctx, deltaTokens, prevPaths)
			test.expect(t, colls)

			for _, coll := range colls {
				assert.NotEqual(t, data.DeletedState, coll.State(), "metadata is never deleted")
			}
		})
	}
}

// This test is primarily aimed at multi-drive handling.
// More complicated single-drive tests can be found in _MakeDriveCollections.
func (suite *CollectionsTreeUnitSuite) TestCollections_GetTree() {
	// metadataPath, err := path.BuildMetadata(
	// 	tenant,
	// 	user,
	// 	path.OneDriveService,
	// 	path.FilesCategory,
	// 	false)
	// require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))
	type expected struct {
		canUsePrevBackup assert.BoolAssertionFunc
		counts           countTD.Expected
		deltas           map[string]string
		prevPaths        map[string]map[string]string
		skips            int
	}

	table := []struct {
		name          string
		enumerator    mock.EnumerateDriveItemsDelta
		previousPaths map[string]map[string]string

		metadata []data.RestoreCollection
		expect   expected
	}{
		{
			name: "not yet implemented",
			enumerator: mock.DriveEnumerator(
				mock.Drive().NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage()))),
			expect: expected{
				canUsePrevBackup: assert.False,
				counts: countTD.Expected{
					count.PrevPaths: 0,
				},
				deltas:    map[string]string{},
				prevPaths: map[string]map[string]string{},
				skips:     0,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				mbh            = mock.DefaultDriveBHWith(user, test.enumerator)
				c              = collWithMBH(mbh)
				prevMetadata   = makePrevMetadataColls(t, mbh, test.previousPaths)
				globalExcludes = prefixmatcher.NewStringSetBuilder()
				errs           = fault.New(true)
			)

			_, _, err := c.getTree(
				ctx,
				prevMetadata,
				globalExcludes,
				errs)

			require.ErrorIs(t, err, errGetTreeNotImplemented, clues.ToCore(err))
			// TODO(keepers): awaiting implementation
			// assert.Empty(t, colls)
			// assert.Equal(t, test.expect.skips, len(errs.Skipped()))
			// test.expect.canUsePrevBackup(t, canUsePrevBackup)
			// test.expect.counts.Compare(t, c.counter)

			// if err != nil {
			// 	return
			// }

			// for _, coll := range colls {
			// 	collPath := fullOrPrevPath(t, coll)

			// 	if collPath.String() == metadataPath.String() {
			// 		compareMetadata(
			// 			t,
			// 			coll,
			// 			test.expect.deltas,
			// 			test.expect.prevPaths)

			// 		continue
			// 	}

			// 	test.expect.collAssertions.compare(t, coll, globalExcludes)
			// }

			// test.expect.collAssertions.requireNoUnseenCollections(t)
		})
	}
}

// this test is expressly aimed exercising coarse combinations of delta enumeration,
// previous path management, and post processing. Coarse here means the intent is not
// to evaluate every possible combination of inputs and outputs.  More granular tests
// at lower levels are better for verifing fine-grained concerns.  This test only needs
// to ensure we stitch the parts together correctly.
func (suite *CollectionsTreeUnitSuite) TestCollections_MakeDriveCollections() {
	drv := mock.Drive()

	table := []struct {
		name         string
		drive        *mock.DeltaDrive
		enumerator   mock.EnumerateDriveItemsDelta
		prevPaths    map[string]string
		expectCounts countTD.Expected
	}{
		{
			name:  "only root in delta, no prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage()))),
			prevPaths: map[string]string{},
			expectCounts: countTD.Expected{
				count.PrevPaths: 0,
			},
		},
		{
			name:  "only root in delta, with prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage()))),
			prevPaths: map[string]string{
				folderID(): fullPath(folderName()),
			},
			expectCounts: countTD.Expected{
				count.PrevPaths: 1,
			},
		},
		{
			name:  "some items, no prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(folderAtRoot(), fileAt(folder))))),
			prevPaths: map[string]string{},
			expectCounts: countTD.Expected{
				count.PrevPaths: 0,
			},
		},
		{
			name:  "some items, with prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(folderAtRoot(), fileAt(folder))))),
			prevPaths: map[string]string{
				folderID(): fullPath(folderName()),
			},
			expectCounts: countTD.Expected{
				count.PrevPaths: 1,
			},
		},
		{
			name:  "tree had delta reset, only root after, no prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage()))),
			prevPaths: map[string]string{},
			expectCounts: countTD.Expected{
				count.PrevPaths: 0,
			},
		},
		{
			name:  "tree had delta reset, only root after, with prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage()))),
			prevPaths: map[string]string{
				folderID(): fullPath(folderName()),
			},
			expectCounts: countTD.Expected{
				count.PrevPaths: 1,
			},
		},
		{
			name:  "tree had delta reset, enumerate items after, no prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(folderAtRoot(), fileAt(folder))))),
			prevPaths: map[string]string{},
			expectCounts: countTD.Expected{
				count.PrevPaths: 0,
			},
		},
		{
			name:  "tree had delta reset, enumerate items after, with prev paths",
			drive: drv,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(folderAtRoot(), fileAt(folder))))),
			prevPaths: map[string]string{
				folderID(): fullPath(folderName()),
			},
			expectCounts: countTD.Expected{
				count.PrevPaths: 1,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbh := mock.DefaultOneDriveBH(user)
			mbh.DriveItemEnumeration = test.enumerator

			c := collWithMBH(mbh)

			_, _, _, err := c.makeDriveCollections(
				ctx,
				test.drive.Able,
				test.prevPaths,
				id(delta, "prev"),
				newPagerLimiter(control.DefaultOptions()),
				prefixmatcher.NewStringSetBuilder(),
				c.counter,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			test.expectCounts.Compare(t, c.counter)
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddPrevPathsToTree_errors() {
	table := []struct {
		name      string
		tree      func(t *testing.T) *folderyMcFolderFace
		prevPaths map[string]string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "no error - normal usage",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				folderID("parent"): fullPath(folderName("parent")),
				folderID():         fullPath(folderName("parent"), folderName()),
			},
			expectErr: require.NoError,
		},
		{
			name:      "no error - prev paths are empty",
			tree:      treeWithFolders,
			prevPaths: map[string]string{},
			expectErr: require.NoError,
		},
		{
			name: "no error - folder not visited in this delta",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				id("santa"): fullPath(name("santa")),
			},
			expectErr: require.NoError,
		},
		{
			name: "empty key in previous paths",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				"": fullPath(folderName("parent")),
			},
			expectErr: require.Error,
		},
		{
			name: "empty value in previous paths",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				folderID(): "",
			},
			expectErr: require.Error,
		},
		{
			name: "malformed value in previous paths",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				folderID(): "not a path",
			},
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			tree := test.tree(t)

			err := addPrevPathsToTree(
				ctx,
				tree,
				test.prevPaths,
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_TurnTreeIntoCollections() {
	type expected struct {
		prevPaths             map[string]string
		collections           func(t *testing.T) expectedCollections
		globalExcludedFileIDs map[string]struct{}
	}

	table := []struct {
		name           string
		tree           func(t *testing.T) *folderyMcFolderFace
		prevPaths      map[string]string
		enableURLCache bool
		expect         expected
	}{
		{
			name:           "all new collections",
			tree:           fullTree,
			prevPaths:      map[string]string{},
			enableURLCache: true,
			expect: expected{
				prevPaths: map[string]string{
					rootID:             fullPath(),
					folderID("parent"): fullPath(folderName("parent")),
					folderID():         fullPath(folderName("parent"), folderName()),
				},
				collections: func(t *testing.T) expectedCollections {
					return expectCollections(
						false,
						true,
						aColl(
							fullPathPath(t),
							nil,
							fileID("r")),
						aColl(
							fullPathPath(t, folderName("parent")),
							nil,
							fileID("p")),
						aColl(
							fullPathPath(t, folderName("parent"), folderName()),
							nil,
							fileID()))
				},
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID()),
			},
		},
		{
			name:           "all folders moved",
			tree:           fullTree,
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                fullPath(),
				folderID("parent"):    fullPath(folderName("parent-prev")),
				folderID():            fullPath(folderName("parent-prev"), folderName()),
				folderID("tombstone"): fullPath(folderName("tombstone-prev")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:             fullPath(),
					folderID("parent"): fullPath(folderName("parent")),
					folderID():         fullPath(folderName("parent"), folderName()),
				},
				collections: func(t *testing.T) expectedCollections {
					return expectCollections(
						false,
						true,
						aColl(
							fullPathPath(t),
							fullPathPath(t),
							fileID("r")),
						aColl(
							fullPathPath(t, folderName("parent")),
							fullPathPath(t, folderName("parent-prev")),
							fileID("p")),
						aColl(
							fullPathPath(t, folderName("parent"), folderName()),
							fullPathPath(t, folderName("parent-prev"), folderName()),
							fileID()),
						aColl(nil, fullPathPath(t, folderName("tombstone-prev"))))
				},
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID()),
			},
		},
		{
			name:           "all folders moved - todo: path separator string check",
			tree:           fullTreeWithNames("parent", "tombstone"),
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                fullPath(),
				folderID("parent"):    fullPath(folderName("parent-prev")),
				folderID():            fullPath(folderName("parent-prev"), folderName()),
				folderID("tombstone"): fullPath(folderName("tombstone-prev")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:             fullPath(),
					folderID("parent"): fullPath(folderName("parent")),
					folderID():         fullPath(folderName("parent"), folderName()),
				},
				collections: func(t *testing.T) expectedCollections {
					return expectCollections(
						false,
						true,
						aColl(
							fullPathPath(t),
							fullPathPath(t),
							fileID("r")),
						aColl(
							fullPathPath(t, folderName("parent")),
							fullPathPath(t, folderName("parent-prev")),
							fileID("p")),
						aColl(
							fullPathPath(t, folderName("parent"), folderName()),
							fullPathPath(t, folderName("parent-prev"), folderName()),
							fileID()),
						aColl(nil, fullPathPath(t, folderName("tombstone-prev"))))
				},
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID()),
			},
		},
		{
			name:           "no folders moved",
			tree:           fullTree,
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                fullPath(),
				folderID("parent"):    fullPath(folderName("parent")),
				folderID():            fullPath(folderName("parent"), folderName()),
				folderID("tombstone"): fullPath(folderName("tombstone")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:             fullPath(),
					folderID("parent"): fullPath(folderName("parent")),
					folderID():         fullPath(folderName("parent"), folderName()),
				},
				collections: func(t *testing.T) expectedCollections {
					return expectCollections(
						false,
						true,
						aColl(
							fullPathPath(t),
							fullPathPath(t),
							fileID("r")),
						aColl(
							fullPathPath(t, folderName("parent")),
							fullPathPath(t, folderName("parent")),
							fileID("p")),
						aColl(
							fullPathPath(t, folderName("parent"), folderName()),
							fullPathPath(t, folderName("parent"), folderName()),
							fileID()),
						aColl(nil, fullPathPath(t, folderName("tombstone"))))
				},
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID()),
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			tree := test.tree(t)

			err := addPrevPathsToTree(ctx, tree, test.prevPaths, fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			c := collWithMBH(mock.DefaultOneDriveBH(user))

			countPages := 9001
			if test.enableURLCache {
				countPages = 1
			}

			colls, newPrevPaths, excluded, err := c.turnTreeIntoCollections(
				ctx,
				tree,
				id(drive),
				delta,
				countPages,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.prevPaths, newPrevPaths, "new previous paths")

			expectColls := test.expect.collections(t)
			expectColls.compare(t, colls)
			expectColls.requireNoUnseenCollections(t)

			assert.Equal(t, test.expect.globalExcludedFileIDs, excluded)
		})
	}
}

type populateTreeExpected struct {
	counts                        countTD.Expected
	err                           require.ErrorAssertionFunc
	numLiveFiles                  int
	numLiveFolders                int
	shouldHitLimit                bool
	sizeBytes                     int64
	treeContainsFolderIDs         []string
	treeContainsTombstoneIDs      []string
	treeContainsFileIDsWithParent map[string]string
}

type populateTreeTest struct {
	name       string
	enumerator mock.EnumerateDriveItemsDelta
	tree       func(t *testing.T) *folderyMcFolderFace
	limiter    *pagerLimiter
	expect     populateTreeExpected
}

// this test focuses on the population of a tree using a single delta's enumeration data.
// It is not concerned with unifying previous paths or post-processing collections.
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree_singleDelta() {
	drv := mock.Drive()

	table := []populateTreeTest{
		{
			name: "nil page",
			tree: newTree,
			// special case enumerator to generate a null page.
			// otherwise all enumerators should be DriveEnumerator()s.
			enumerator: mock.EnumerateDriveItemsDelta{
				DrivePagers: map[string]*mock.DeltaDriveEnumerator{
					drv.ID: {
						Drive: drv,
						DeltaQueries: []*mock.DeltaQuery{{
							Pages:       nil,
							DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
						}},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts:                        countTD.Expected{},
				err:                           require.NoError,
				numLiveFiles:                  0,
				numLiveFolders:                0,
				sizeBytes:                     0,
				treeContainsFolderIDs:         []string{},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "root only",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage()))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 1,
					count.TotalFilesProcessed:   0,
					count.TotalPagesEnumerated:  2,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 1,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "root only on two pages",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(),
						aPage()))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 2,
					count.TotalFilesProcessed:   0,
					count.TotalPagesEnumerated:  3,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 1,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "many folders in a hierarchy across multiple pages",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(folderAtRoot()),
						aPage(folderAtRoot("sib")),
						aPage(
							folderAtRoot(),
							folderAt(folder, "chld"))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 7,
					count.TotalPagesEnumerated:  4,
					count.TotalFilesProcessed:   0,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 4,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
					folderID("sib"),
					folderID("chld"),
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "many folders with files",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder)),
						aPage(
							folderAtRoot("sib"),
							fileAt("sib", "fsib")),
						aPage(
							folderAtRoot(),
							folderAt(folder, "chld"),
							fileAt("chld", "fchld")),
					))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 7,
					count.TotalFilesProcessed:   3,
					count.TotalPagesEnumerated:  4,
				},
				err:            require.NoError,
				numLiveFiles:   3,
				numLiveFolders: 4,
				sizeBytes:      3 * 42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
					folderID("sib"),
					folderID("chld"),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID():        folderID(),
					fileID("fsib"):  folderID("sib"),
					fileID("fchld"): folderID("chld"),
				},
			},
		},
		{
			name: "tombstone with unpopulated tree",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(delItem(folderID(), folderID("parent"), isFolder)),
					))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFilesProcessed:         0,
					count.TotalPagesEnumerated:        2,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 1,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{
					folderID(),
				},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "tombstone with populated tree",
			tree: treeWithFileInFolder,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						delItem(folderID(), folderID("parent"), isFolder),
					)))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFilesProcessed:         0,
					count.TotalPagesEnumerated:        2,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 2,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
					folderID("parent"),
				},
				treeContainsTombstoneIDs: []string{
					folderID(),
				},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			// technically you won't see this behavior from graph deltas, since deletes always
			// precede creates/updates.  But it's worth checking that we can handle it anyways.
			name: "create->delete folder",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder)),
						aPage(delItem(folderID(), rootID, isFolder))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       3,
					count.TotalFilesProcessed:         1,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalPagesEnumerated:        3,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 1,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			// technically you won't see this behavior from graph deltas, since deletes always
			// precede creates/updates.  But it's worth checking that we can handle it anyways.
			name: "move->delete folder with populated tree",
			tree: treeWithFolders,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot("parent"),
							driveItem(folderID(), folderName("moved"), parentDir(), folderID("parent"), isFolder),
							driveFile(parentDir(folderName("parent"), folderName()), folderID())),
						aPage(delItem(folderID(), folderID("parent"), isFolder))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       4,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFilesProcessed:         1,
					count.TotalPagesEnumerated:        3,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 2,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
					folderID("parent"),
				},
				treeContainsTombstoneIDs: []string{
					folderID(),
				},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			name: "delete->create folder with previous path",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(delItem(folderID(), rootID, isFolder)),
						aPage(
							folderAtRoot(),
							fileAt(folder)),
					))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFoldersProcessed:       3,
					count.TotalFilesProcessed:         1,
					count.TotalPagesEnumerated:        3,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			name: "delete->create folder without previous path",
			tree: treeWithRoot,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(delItem(folderID(), rootID, isFolder)),
						aPage(
							folderAtRoot(),
							fileAt(folder)),
					))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFoldersProcessed:       3,
					count.TotalFilesProcessed:         1,
					count.TotalPagesEnumerated:        3,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			name: "at folder limit before enumeration",
			tree: treeWithFileAtRoot,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder)),
						aPage(
							folderAtRoot("sib"),
							fileAt("sib", "fsib")),
						aPage(
							folderAtRoot(),
							folderAt(folder, "chld"),
							fileAt("chld", "fchld"))))),
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.TotalFilesProcessed:         0,
					count.TotalPagesEnumerated:        1,
				},
				err:            require.NoError,
				shouldHitLimit: true,
				numLiveFiles:   1,
				numLiveFolders: 1,
				sizeBytes:      defaultItemSize,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "hit folder limit during enumeration",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder)),
						aPage(
							folderAtRoot("sib"),
							fileAt("sib", "fsib")),
						aPage(
							folderAtRoot(),
							folderAt(folder, "chld"),
							fileAt("chld", "fchld"))))),
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.TotalFilesProcessed:         0,
					count.TotalPagesEnumerated:        1,
				},
				err:            require.NoError,
				shouldHitLimit: true,
				numLiveFiles:   0,
				numLiveFolders: 1,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			runPopulateTreeTest(suite.T(), drv.Able, test)
		})
	}
}

// this test focuses on quirks that can only arise from cases that occur across
// multiple delta enumerations.
// It is not concerned with unifying previous paths or post-processing collections.
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree_multiDelta() {
	drv := mock.Drive()

	table := []populateTreeTest{
		{
			name: "sanity case: normal enumeration split across multiple deltas",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder))),
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot("sib"),
							fileAt("sib", "fsib"))),
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							folderAt(folder, "chld"),
							fileAt("chld", "fchld"))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeltasProcessed:        4,
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalDeleteFilesProcessed:   0,
					count.TotalFilesProcessed:         3,
					count.TotalFoldersProcessed:       7,
					count.TotalPagesEnumerated:        4,
				},
				err:            require.NoError,
				numLiveFiles:   3,
				numLiveFolders: 4,
				sizeBytes:      3 * 42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
					folderID("sib"),
					folderID("chld"),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID():        folderID(),
					fileID("fsib"):  folderID("sib"),
					fileID("fchld"): folderID("chld"),
				},
			},
		},
		{
			name: "create->delete,create",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder))),
					// a (delete,create) pair in the same delta can occur when
					// a user deletes and restores an item in-between deltas.
					mock.Delta(id(delta), nil).With(
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), folderID(), isFile)),
						aPage(
							folderAtRoot(),
							fileAt(folder))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeltasProcessed:        3,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalDeleteFilesProcessed:   1,
					count.TotalFilesProcessed:         2,
					count.TotalFoldersProcessed:       5,
					count.TotalPagesEnumerated:        4,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "visit->rename",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						aPage(
							folderAtRoot(),
							fileAt(folder))),
					mock.Delta(id(delta), nil).With(
						aPage(
							driveItem(folderID(), folderName("rename"), parentDir(), rootID, isFolder),
							driveItem(fileID(), fileName("rename"), parentDir(folderName("rename")), folderID(), isFile))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeltasProcessed:        3,
					count.TotalDeleteFilesProcessed:   0,
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFilesProcessed:         2,
					count.TotalFoldersProcessed:       4,
					count.TotalPagesEnumerated:        3,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(),
				},
			},
		},
		{
			name: "duplicate folder name from deferred delete marker",
			tree: newTree,
			enumerator: mock.DriveEnumerator(
				drv.NewEnumer().With(
					mock.Delta(id(delta), nil).With(
						// first page: create /root/folder and /root/folder/file
						aPage(
							folderAtRoot(),
							fileAt(folder)),
						// assume the user makes changes at this point:
						// * create a new /root/folder
						// * move /root/folder/file from old to new folder (same file ID)
						// * delete /root/folder
						// in drive deltas, this will show up as another folder creation sharing
						// the same dirname, but we won't see the delete until...
						aPage(
							driveItem(folderID(2), folderName(), parentDir(), rootID, isFolder),
							driveItem(fileID(), fileName(), parentDir(folderName()), folderID(2), isFile))),
					// the next delta, containing the delete marker for the original /root/folder
					mock.Delta(id(delta), nil).With(
						aPage(
							delItem(folderID(), rootID, isFolder))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeltasProcessed:        3,
					count.TotalDeleteFilesProcessed:   0,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFilesProcessed:         2,
					count.TotalFoldersProcessed:       5,
					count.TotalPagesEnumerated:        4,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      42,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(2),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): folderID(2),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			runPopulateTreeTest(suite.T(), drv.Able, test)
		})
	}
}

func runPopulateTreeTest(
	t *testing.T,
	drv models.Driveable,
	test populateTreeTest,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mbh     = mock.DefaultDriveBHWith(user, test.enumerator)
		c       = collWithMBH(mbh)
		counter = count.New()
		tree    = test.tree(t)
	)

	_, _, err := c.populateTree(
		ctx,
		tree,
		drv,
		id(delta),
		test.limiter,
		counter,
		fault.New(true))

	test.expect.err(t, err, clues.ToCore(err))

	assert.Equal(
		t,
		test.expect.numLiveFolders,
		tree.countLiveFolders(),
		"count live folders in tree")

	cAndS := tree.countLiveFilesAndSizes()
	assert.Equal(
		t,
		test.expect.numLiveFiles,
		cAndS.numFiles,
		"count live files in tree")
	assert.Equal(
		t,
		test.expect.sizeBytes,
		cAndS.totalBytes,
		"count total bytes in tree")
	test.expect.counts.Compare(t, counter)

	for _, id := range test.expect.treeContainsFolderIDs {
		assert.NotNil(t, tree.folderIDToNode[id], "node exists")
	}

	for _, id := range test.expect.treeContainsTombstoneIDs {
		assert.NotNil(t, tree.tombstones[id], "tombstone exists")
	}

	for iID, pID := range test.expect.treeContainsFileIDsWithParent {
		assert.Contains(t, tree.fileIDToParentID, iID, "file should exist in tree")
		assert.Equal(t, pID, tree.fileIDToParentID[iID], "file should reference correct parent")
	}
}

// ---------------------------------------------------------------------------
// folder tests
// ---------------------------------------------------------------------------

// This test focuses on folder assertions when enumerating a page of items.
// File-specific assertions are focused in the _folders test variant.
func (suite *CollectionsTreeUnitSuite) TestCollections_EnumeratePageOfItems_folders() {
	drv := mock.Drive()

	type expected struct {
		counts                   countTD.Expected
		err                      require.ErrorAssertionFunc
		shouldHitLimit           bool
		treeSize                 int
		treeContainsFolderIDs    []string
		treeContainsTombstoneIDs []string
	}

	table := []struct {
		name    string
		tree    func(t *testing.T) *folderyMcFolderFace
		page    mock.NextPage
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "nil page",
			tree:    treeWithRoot,
			page:    mock.NextPage{},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts:   countTD.Expected{},
				err:      require.NoError,
				treeSize: 1,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name:    "empty page",
			tree:    treeWithRoot,
			page:    mock.NextPage{Items: []models.DriveItemable{}},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts:   countTD.Expected{},
				err:      require.NoError,
				treeSize: 1,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name:    "root only",
			tree:    treeWithRoot,
			page:    aPage(),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 1,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "many folders in a hierarchy",
			tree: treeWithRoot,
			page: aPage(
				folderAtRoot(),
				folderAtRoot("sib"),
				folderAt(folder, "chld")),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 4,
				},
				err:      require.NoError,
				treeSize: 4,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
					folderID("sib"),
					folderID("chld"),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "create->delete",
			tree: treeWithRoot,
			page: aPage(
				folderAtRoot(),
				delItem(folderID(), rootID, isFolder)),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       2,
					count.TotalDeleteFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "move->delete",
			tree: treeWithFolders,
			page: aPage(
				folderAtRoot("parent"),
				driveItem(folderID(), folderName("moved"), parentDir(folderName("parent")), folderID("parent"), isFolder),
				delItem(folderID(), folderID("parent"), isFolder)),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       3,
					count.TotalDeleteFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 3,
				treeContainsFolderIDs: []string{
					rootID,
					folderID("parent"),
				},
				treeContainsTombstoneIDs: []string{
					folderID(),
				},
			},
		},
		{
			name: "delete->create with previous path",
			tree: treeWithRoot,
			page: aPage(
				delItem(folderID(), rootID, isFolder),
				folderAtRoot()),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       2,
					count.TotalDeleteFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "delete->create without previous path",
			tree: treeWithRoot,
			page: aPage(
				delItem(folderID(), rootID, isFolder),
				folderAtRoot()),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       2,
					count.TotalDeleteFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					folderID(),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(mock.DefaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t)
			)

			err := c.enumeratePageOfItems(
				ctx,
				tree,
				drv.Able,
				test.page.Items,
				test.limiter,
				counter,
				fault.New(true))

			test.expect.err(t, err, clues.ToCore(err))
			if test.expect.shouldHitLimit {
				assert.ErrorIs(t, err, errHitLimit, clues.ToCore(err))
			}

			assert.Equal(
				t,
				test.expect.treeSize,
				len(tree.tombstones)+tree.countLiveFolders(),
				"count folders in tree")
			test.expect.counts.Compare(t, counter)

			for _, id := range test.expect.treeContainsFolderIDs {
				assert.NotNil(t, tree.folderIDToNode[id], "node exists")
			}

			for _, id := range test.expect.treeContainsTombstoneIDs {
				assert.NotNil(t, tree.tombstones[id], "tombstone exists")
			}
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddFolderToTree() {
	var (
		drv    = mock.Drive()
		fld    = custom.ToCustomDriveItem(folderAtRoot())
		subFld = custom.ToCustomDriveItem(driveFolder(driveParentDir(drv, folderName("parent")), folderID("parent")))
		pack   = custom.ToCustomDriveItem(driveItem(id(pkg), name(pkg), parentDir(), rootID, isPackage))
		del    = custom.ToCustomDriveItem(delItem(folderID(), rootID, isFolder))
		mal    = custom.ToCustomDriveItem(malwareItem(folderID("mal"), folderName("mal"), parentDir(), rootID, isFolder))
	)

	type expected struct {
		countLiveFolders   int
		counts             countTD.Expected
		err                require.ErrorAssertionFunc
		shouldHitLimit     bool
		treeSize           int
		treeContainsFolder assert.BoolAssertionFunc
		skipped            assert.ValueAssertionFunc
	}

	table := []struct {
		name    string
		tree    func(t *testing.T) *folderyMcFolderFace
		folder  *custom.DriveItem
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "add folder",
			tree:    treeWithRoot,
			folder:  fld,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 2,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 0,
				},
				treeSize:           2,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "re-add folder that already exists",
			tree:    treeWithFolders,
			folder:  subFld,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 3,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 0,
				},
				treeSize:           3,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "add package",
			tree:    treeWithRoot,
			folder:  pack,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 2,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      1,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 0,
				},
				treeSize:           2,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "tombstone a folder in a populated tree",
			tree:    treeWithFolders,
			folder:  del,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 2,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 1,
				},
				treeSize:           3,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "tombstone new folder in unpopulated tree",
			tree:    newTree,
			folder:  del,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				err: require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 1,
				},
				treeSize:           1,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "re-add tombstone that already exists",
			tree:    treeWithTombstone,
			folder:  del,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 1,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 1,
				},
				treeSize:           2,
				treeContainsFolder: assert.True,
				skipped:            assert.Nil,
			},
		},
		{
			name:    "add malware",
			tree:    treeWithRoot,
			folder:  mal,
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				countLiveFolders: 1,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       1,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 0,
				},
				treeSize:           1,
				treeContainsFolder: assert.False,
				skipped:            assert.NotNil,
			},
		},
		{
			name:    "already over container limit, folder seen twice",
			tree:    treeWithFolders,
			folder:  fld,
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				countLiveFolders: 3,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 0,
				},
				shouldHitLimit:     false,
				skipped:            assert.Nil,
				treeSize:           3,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:    "already at container limit",
			tree:    treeWithRoot,
			folder:  fld,
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				countLiveFolders: 1,
				err:              require.Error,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 0,
				},
				shouldHitLimit:     true,
				skipped:            assert.Nil,
				treeSize:           1,
				treeContainsFolder: assert.False,
			},
		},
		{
			name:    "process tombstone when over folder limits",
			tree:    treeWithFolders,
			folder:  del,
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				countLiveFolders: 2,
				err:              require.NoError,
				counts: countTD.Expected{
					count.TotalMalwareProcessed:       0,
					count.TotalPackagesProcessed:      0,
					count.TotalFoldersProcessed:       0,
					count.TotalDeleteFoldersProcessed: 1,
				},
				shouldHitLimit:     false,
				skipped:            assert.Nil,
				treeSize:           3,
				treeContainsFolder: assert.True,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(mock.DefaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t)
			)

			skipped, err := c.addFolderToTree(
				ctx,
				tree,
				drv.Able,
				test.folder,
				test.limiter,
				counter)

			test.expect.err(t, err, clues.ToCore(err))
			test.expect.skipped(t, skipped)

			if test.expect.shouldHitLimit {
				assert.ErrorIs(t, err, errHitLimit, clues.ToCore(err))
			}

			test.expect.counts.Compare(t, counter)
			assert.Equal(t, test.expect.countLiveFolders, tree.countLiveFolders(), "live folders")
			assert.Equal(
				t,
				test.expect.treeSize,
				len(tree.tombstones)+tree.countLiveFolders(),
				"folders in tree")
			test.expect.treeContainsFolder(t, tree.containsFolder(ptr.Val(test.folder.GetId())))
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_MakeFolderCollectionPath() {
	drv := mock.Drive()

	basePath, err := odConsts.DriveFolderPrefixBuilder(drv.ID).
		ToDataLayerOneDrivePath(tenant, user, false)
	require.NoError(suite.T(), err, clues.ToCore(err))

	folderPath, err := basePath.Append(false, folderName())
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name      string
		folder    models.DriveItemable
		expect    string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "root",
			folder:    driveRootFolder(),
			expect:    basePath.String(),
			expectErr: require.NoError,
		},
		{
			name:      "folder",
			folder:    folderAtRoot(),
			expect:    folderPath.String(),
			expectErr: require.NoError,
		},
		{
			name:      "folder without parent ref",
			folder:    models.NewDriveItem(),
			expect:    folderPath.String(),
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			c := collWithMBH(mock.DefaultOneDriveBH(user))

			p, err := c.makeFolderCollectionPath(
				ctx,
				drv.ID,
				custom.ToCustomDriveItem(test.folder))
			test.expectErr(t, err, clues.ToCore(err))

			if err == nil {
				assert.Equal(t, test.expect, p.String())
			}
		})
	}
}

// ---------------------------------------------------------------------------
// file tests
// ---------------------------------------------------------------------------

// this test focuses on folder assertions when enumerating a page of items
// file-specific assertions are in the next test
func (suite *CollectionsTreeUnitSuite) TestCollections_EnumeratePageOfItems_files() {
	drv := mock.Drive()

	type expected struct {
		counts                        countTD.Expected
		err                           require.ErrorAssertionFunc
		treeContainsFileIDsWithParent map[string]string
		countLiveFiles                int
		countTotalBytes               int64
	}

	table := []struct {
		name   string
		tree   func(t *testing.T) *folderyMcFolderFace
		page   mock.NextPage
		expect expected
	}{
		{
			name: "one file at root",
			tree: treeWithRoot,
			page: aPage(fileAtRoot()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "many files in a hierarchy",
			tree: treeWithRoot,
			page: aPage(
				fileAtRoot(),
				folderAtRoot(),
				fileAt(folder, "fchld")),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     2,
					count.TotalFilesProcessed:       2,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					fileID():        rootID,
					fileID("fchld"): folderID(),
				},
				countLiveFiles:  2,
				countTotalBytes: defaultItemSize * 2,
			},
		},
		{
			name: "many updates to the same file",
			tree: treeWithRoot,
			page: aPage(
				fileAtRoot(),
				driveItem(fileID(), fileName(1), parentDir(), rootID, isFile),
				driveItem(fileID(), fileName(2), parentDir(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       3,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "delete an existing file",
			tree: treeWithFileAtRoot,
			page: aPage(delItem(fileID(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       0,
				},
				err:                           require.NoError,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name: "delete the same file twice",
			tree: treeWithFileAtRoot,
			page: aPage(
				delItem(fileID(), rootID, isFile),
				delItem(fileID(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 2,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       0,
				},
				err:                           require.NoError,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name: "create->delete",
			tree: treeWithRoot,
			page: aPage(
				fileAtRoot(),
				delItem(fileID(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err:                           require.NoError,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name: "move->delete",
			tree: treeWithFileAtRoot,
			page: aPage(
				folderAtRoot(),
				fileAt(folder),
				delItem(fileID(), folderID(), isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     2,
					count.TotalFilesProcessed:       1,
				},
				err:                           require.NoError,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name: "file already in tree: delete->restore",
			tree: treeWithFileAtRoot,
			page: aPage(
				delItem(fileID(), rootID, isFile),
				fileAtRoot()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "file not in tree: delete->restore",
			tree: treeWithRoot,
			page: aPage(
				delItem(fileID(), rootID, isFile),
				fileAtRoot()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(mock.DefaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t)
			)

			err := c.enumeratePageOfItems(
				ctx,
				tree,
				drv.Able,
				test.page.Items,
				newPagerLimiter(control.DefaultOptions()),
				counter,
				fault.New(true))
			test.expect.err(t, err, clues.ToCore(err))

			countSize := tree.countLiveFilesAndSizes()
			assert.Equal(t, test.expect.countLiveFiles, countSize.numFiles, "count of files")
			assert.Equal(t, test.expect.countTotalBytes, countSize.totalBytes, "total size in bytes")
			assert.Equal(t, test.expect.treeContainsFileIDsWithParent, tree.fileIDToParentID)
			test.expect.counts.Compare(t, counter)
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddFileToTree() {
	drv := mock.Drive()

	type expected struct {
		counts                        countTD.Expected
		err                           require.ErrorAssertionFunc
		shouldHitLimit                bool
		skipped                       assert.ValueAssertionFunc
		treeContainsFileIDsWithParent map[string]string
		countLiveFiles                int
		countTotalBytes               int64
	}

	table := []struct {
		name    string
		tree    func(t *testing.T) *folderyMcFolderFace
		file    models.DriveItemable
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "add new file",
			tree:    treeWithRoot,
			file:    fileAtRoot(),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:     require.NoError,
				skipped: assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "duplicate file",
			tree:    treeWithFileAtRoot,
			file:    fileAtRoot(),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:     require.NoError,
				skipped: assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "error file seen before parent",
			tree:    treeWithRoot,
			file:    fileAt(folder),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:                           require.Error,
				skipped:                       assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name:    "malware file",
			tree:    treeWithRoot,
			file:    malwareItem(fileID(), fileName(), parentDir(folderName()), rootID, isFile),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalMalwareProcessed: 1,
				},
				err:                           require.NoError,
				skipped:                       assert.NotNil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name:    "delete non-existing file",
			tree:    treeWithRoot,
			file:    delItem(fileID(), folderID(), isFile),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
				},
				err:                           require.NoError,
				skipped:                       assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name:    "delete existing file",
			tree:    treeWithFileAtRoot,
			file:    delItem(fileID(), rootID, isFile),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
				},
				err:                           require.NoError,
				skipped:                       assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name:    "already at container file limit",
			tree:    treeWithFileAtRoot,
			file:    fileAtRoot(2),
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:            require.Error,
				shouldHitLimit: true,
				skipped:        assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "goes over total byte limit",
			tree:    treeWithRoot,
			file:    fileAtRoot(),
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:                           require.Error,
				shouldHitLimit:                true,
				skipped:                       assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(mock.DefaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t)
			)

			skipped, err := c.addFileToTree(
				ctx,
				tree,
				drv.Able,
				custom.ToCustomDriveItem(test.file),
				test.limiter,
				counter)

			test.expect.err(t, err, clues.ToCore(err))
			test.expect.skipped(t, skipped)

			if test.expect.shouldHitLimit {
				require.ErrorIs(t, err, errHitLimit, clues.ToCore(err))
			}

			assert.Equal(t, test.expect.treeContainsFileIDsWithParent, tree.fileIDToParentID)
			test.expect.counts.Compare(t, counter)

			countSize := tree.countLiveFilesAndSizes()
			assert.Equal(t, test.expect.countLiveFiles, countSize.numFiles, "count of files")
			assert.Equal(t, test.expect.countTotalBytes, countSize.totalBytes, "total size in bytes")
		})
	}
}
