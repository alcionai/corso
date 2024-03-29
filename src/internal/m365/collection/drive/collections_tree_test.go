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
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	countTD "github.com/alcionai/corso/src/pkg/count/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
	badPfxMBH := defaultOneDriveBH(user)
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
			c:          collWithMBH(defaultOneDriveBH(user)),
			expectErr:  assert.NoError,
			expect:     assert.Empty,
		},
		{
			name:       "none",
			tombstones: map[string]struct{}{},
			c:          collWithMBH(defaultOneDriveBH(user)),
			expectErr:  assert.NoError,
			expect:     assert.Empty,
		},
		{
			name:       "some tombstones",
			tombstones: twostones,
			c:          collWithMBH(defaultOneDriveBH(user)),
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
	badMetaPfxMBH := defaultOneDriveBH(user)
	badMetaPfxMBH.MetadataPathPrefixErr = assert.AnError

	table := []struct {
		name   string
		c      *Collections
		expect assert.ValueAssertionFunc
	}{
		{
			name:   "no errors",
			c:      collWithMBH(defaultOneDriveBH(user)),
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
// Note: these tests intentionally duplicate file and folder IDs and names
// across drives.  This mimics real drive api behavior, where ids and names are
// only unique within each drive.
func (suite *CollectionsTreeUnitSuite) TestCollections_GetTree() {
	t := suite.T()

	type expected struct {
		canUsePrevBackup      assert.BoolAssertionFunc
		collections           expectedCollections
		globalExcludedFileIDs *prefixmatcher.StringSetMatchBuilder
	}

	d1 := drive()
	d2 := drive(2)

	table := []struct {
		name       string
		enumerator enumerateDriveItemsDelta
		metadata   []data.RestoreCollection
		expect     expected
	}{
		{
			// in the real world we'd be guaranteed the root folder in the first
			// delta query, at minimum (ie, the backup before this one).
			// but this makes for a good sanity check on no-op behavior.
			name: "no data either previously or in delta",
			enumerator: driveEnumerator(
				d1.newEnumer().with(delta(nil)),
				d2.newEnumer().with(delta(nil))),
			metadata: multiDriveMetadata(t),
			expect: expected{
				canUsePrevBackup: assert.True,
				collections: expectCollections(
					false,
					true,
					aMetadata(nil, multiDrivePrevPaths(
						d1.newPrevPaths(t),
						d2.newPrevPaths(t)))),
				globalExcludedFileIDs: multiDriveExcludeMap(
					d1.newExcludes(t, makeExcludeMap()),
					d2.newExcludes(t, makeExcludeMap())),
			},
		},
		{
			name: "a new backup",
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "r"),
							d1.folderAt(root),
							d1.fileAt(folder, "f")))),
				d2.newEnumer().with(
					delta(nil).with(
						aPage(
							d2.fileAt(root, "r"),
							d2.folderAt(root),
							d2.fileAt(folder, "f"))))),
			metadata: multiDriveMetadata(t),
			expect: expected{
				canUsePrevBackup: assert.True,
				collections: expectCollections(
					false,
					true,
					aColl(
						d1.fullPath(t),
						nil,
						fileID("r")),
					aColl(
						d1.fullPath(t, folderName()),
						nil,
						fileID("f")),
					aColl(
						d2.fullPath(t),
						nil,
						fileID("r")),
					aColl(
						d2.fullPath(t, folderName()),
						nil,
						fileID("f")),
					aMetadata(nil, multiDrivePrevPaths(
						d1.newPrevPaths(
							t,
							rootID, d1.strPath(t),
							folderID(), d1.strPath(t, folderName())),
						d2.newPrevPaths(
							t,
							rootID, d2.strPath(t),
							folderID(), d2.strPath(t, folderName()))))),
				globalExcludedFileIDs: multiDriveExcludeMap(
					d1.newExcludes(t, makeExcludeMap(fileID("r"), fileID("f"))),
					d2.newExcludes(t, makeExcludeMap(fileID("r"), fileID("f")))),
			},
		},
		{
			name: "an incremental backup with no changes",
			enumerator: driveEnumerator(
				d1.newEnumer().with(delta(nil).with(aPage())),
				d2.newEnumer().with(delta(nil).with(aPage()))),
			metadata: multiDriveMetadata(
				t,
				d1.newPrevPaths(
					t,
					rootID, d1.strPath(t),
					folderID(), d1.strPath(t, folderName())),
				d2.newPrevPaths(
					t,
					rootID, d2.strPath(t),
					folderID(), d2.strPath(t, folderName()))),
			expect: expected{
				canUsePrevBackup: assert.True,
				collections: expectCollections(
					false,
					true,
					aColl(d1.fullPath(t), d1.fullPath(t)),
					aColl(d2.fullPath(t), d2.fullPath(t)),
					aMetadata(nil, multiDrivePrevPaths(
						d1.newPrevPaths(
							t,
							rootID, d1.strPath(t),
							folderID(), d1.strPath(t, folderName())),
						d2.newPrevPaths(
							t,
							rootID, d2.strPath(t),
							folderID(), d2.strPath(t, folderName()))))),
				globalExcludedFileIDs: multiDriveExcludeMap(
					d1.newExcludes(t, makeExcludeMap()),
					d2.newExcludes(t, makeExcludeMap())),
			},
		},
		{
			name: "an incremental backup with deletions in one drive and movements in the other",
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							delItem(fileID("r"), rootID, isFile),
							delItem(folderID(), rootID, isFolder),
							delItem(fileID("f"), folderID(), isFile)))),
				d2.newEnumer().with(
					delta(nil).with(
						aPage(
							driveItem(fileID("r"), fileName("r2"), d2.dir(), rootID, isFile),
							driveItem(folderID(), folderName(2), d2.dir(), rootID, isFolder),
							driveItem(fileID("f"), fileName("f2"), d2.dir(folderName()), folderID(), isFile))))),
			metadata: multiDriveMetadata(
				t,
				d1.newPrevPaths(
					t,
					rootID, d1.strPath(t),
					folderID(), d1.strPath(t, folderName())),
				d2.newPrevPaths(
					t,
					rootID, d2.strPath(t),
					folderID(), d2.strPath(t, folderName()))),
			expect: expected{
				canUsePrevBackup: assert.True,
				collections: expectCollections(
					false,
					true,
					aColl(d1.fullPath(t), d1.fullPath(t)),
					aTomb(d1.fullPath(t, folderName())),
					aColl(
						d2.fullPath(t),
						d2.fullPath(t),
						fileID("r")),
					aColl(
						d2.fullPath(t, folderName(2)),
						d2.fullPath(t, folderName()),
						fileID("f")),
					aMetadata(nil, multiDrivePrevPaths(
						d1.newPrevPaths(
							t,
							rootID, d1.strPath(t)),
						d2.newPrevPaths(
							t,
							rootID, d2.strPath(t),
							folderID(), d2.strPath(t, folderName(2)))))),
				globalExcludedFileIDs: multiDriveExcludeMap(
					d1.newExcludes(t, makeExcludeMap(fileID("r"), fileID("f"))),
					d2.newExcludes(t, makeExcludeMap(fileID("r"), fileID("f")))),
			},
		},
		{
			name: "tombstone an entire drive",
			enumerator: driveEnumerator(
				d1.newEnumer().with(delta(nil).with(aPage()))),
			metadata: multiDriveMetadata(
				t,
				d1.newPrevPaths(
					t,
					rootID, d1.strPath(t),
					folderID(), d1.strPath(t, folderName())),
				d2.newPrevPaths(
					t,
					rootID, d2.strPath(t),
					folderID(), d2.strPath(t, folderName()))),
			expect: expected{
				canUsePrevBackup: assert.True,
				collections: expectCollections(
					false,
					true,
					aColl(d1.fullPath(t), d1.fullPath(t)),
					aTomb(d2.fullPath(t)),
					aMetadata(nil, multiDrivePrevPaths(
						d1.newPrevPaths(
							t,
							rootID, d1.strPath(t),
							folderID(), d1.strPath(t, folderName()))))),
				globalExcludedFileIDs: multiDriveExcludeMap(
					d1.newExcludes(t, makeExcludeMap())),
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				mbh            = defaultDriveBHWith(user, test.enumerator)
				c              = collWithMBH(mbh)
				globalExcludes = prefixmatcher.NewStringSetBuilder()
				errs           = fault.New(true)
			)

			results, canUsePrevBackup, err := c.getTree(
				ctx,
				test.metadata,
				globalExcludes,
				errs)
			require.NoError(t, err, clues.ToCore(err))
			test.expect.canUsePrevBackup(t, canUsePrevBackup, "can use previous backup")

			expectColls := test.expect.collections
			expectColls.compare(t, results)
			expectColls.requireNoUnseenCollections(t)

			for _, d := range test.expect.globalExcludedFileIDs.Keys() {
				suite.Run(d, func() {
					expect, _ := test.expect.globalExcludedFileIDs.Get(d)
					result, rok := globalExcludes.Get(d)

					if len(test.metadata) > 0 {
						require.True(t, rok, "drive results have a global excludes entry")
						assert.Equal(t, expect, result, "global excluded file IDs")
					} else {
						require.False(t, rok, "drive results have no global excludes entry")
						assert.Empty(t, result, "global excluded file IDs")
					}
				})
			}
		})
	}
}

// this test is expressly aimed exercising coarse combinations of delta enumeration,
// previous path management, and post processing. Coarse here means the intent is not
// to evaluate every possible combination of inputs and outputs.  More granular tests
// at lower levels are better for verifing fine-grained concerns.  This test only needs
// to ensure we stitch the parts together correctly.
func (suite *CollectionsTreeUnitSuite) TestCollections_MakeDriveCollections() {
	d := drive()
	t := suite.T()

	type expected struct {
		collections           expectedCollections
		counts                countTD.Expected
		globalExcludedFileIDs map[string]struct{}
		prevPaths             map[string]string
	}

	table := []struct {
		name       string
		drive      *deltaDrive
		enumerator enumerateDriveItemsDelta
		prevPaths  map[string]string
		expect     expected
	}{
		{
			name:  "only root in delta, no prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage()))),
			prevPaths: map[string]string{},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    0,
					count.NewPrevPaths: 1,
				},
				collections: expectCollections(
					false,
					true,
					aColl(d.fullPath(t), nil)),
				globalExcludedFileIDs: makeExcludeMap(),
				prevPaths: map[string]string{
					rootID: d.strPath(t),
				},
			},
		},
		{
			name:  "only root in delta, with prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage()))),
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    2,
					count.NewPrevPaths: 2,
				},
				collections: expectCollections(
					false,
					true,
					aColl(d.fullPath(t), d.fullPath(t))),
				globalExcludedFileIDs: makeExcludeMap(),
				prevPaths: map[string]string{
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
		},
		{
			name:  "some items, no prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
			prevPaths: map[string]string{},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    0,
					count.NewPrevPaths: 2,
				},
				collections: expectCollections(
					false,
					true,
					aColl(d.fullPath(t), nil),
					aColl(
						d.fullPath(t, folderName()),
						nil,
						fileID())),
				globalExcludedFileIDs: makeExcludeMap(fileID()),
				prevPaths: map[string]string{
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
		},
		{
			name:  "some items, with prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    2,
					count.NewPrevPaths: 2,
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t)),
					aColl(
						d.fullPath(t, folderName()),
						d.fullPath(t, folderName()),
						fileID())),
				globalExcludedFileIDs: makeExcludeMap(fileID()),
				prevPaths: map[string]string{
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
		},
		{
			name:  "only previous path data",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage()))),
			prevPaths: map[string]string{
				rootID:           d.strPath(t),
				folderID():       d.strPath(t, folderName()),
				folderID("sib"):  d.strPath(t, folderName("sib")),
				folderID("chld"): d.strPath(t, folderName(), folderName("chld")),
			},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    4,
					count.NewPrevPaths: 4,
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t))),
				globalExcludedFileIDs: makeExcludeMap(),
				prevPaths: map[string]string{
					rootID:           d.strPath(t),
					folderID():       d.strPath(t, folderName()),
					folderID("sib"):  d.strPath(t, folderName("sib")),
					folderID("chld"): d.strPath(t, folderName(), folderName("chld")),
				},
			},
		},
		{
			name:  "tree had delta reset, only root after, no prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage()))),
			prevPaths: map[string]string{},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    0,
					count.NewPrevPaths: 1,
				},
				collections: expectCollections(
					true,
					true,
					aColl(d.fullPath(t), nil)),
				globalExcludedFileIDs: nil,
				prevPaths: map[string]string{
					rootID: d.strPath(t),
				},
			},
		},
		{
			name:  "tree had delta reset, only root after, with prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage()))),
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    2,
					count.NewPrevPaths: 1,
				},
				collections: expectCollections(
					true,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t)),
					aTomb(d.fullPath(t, folderName()))),
				globalExcludedFileIDs: nil,
				prevPaths: map[string]string{
					rootID: d.strPath(t),
				},
			},
		},
		{
			name:  "tree had delta reset, enumerate items after, no prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
			prevPaths: map[string]string{},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    0,
					count.NewPrevPaths: 2,
				},
				collections: expectCollections(
					true,
					true,
					aColl(
						d.fullPath(t),
						nil),
					aColl(
						d.fullPath(t, folderName()),
						nil,
						fileID())),
				globalExcludedFileIDs: nil,
				prevPaths: map[string]string{
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
		},
		{
			name:  "tree had delta reset, enumerate items after, with prev paths",
			drive: d,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
			prevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expect: expected{
				counts: countTD.Expected{
					count.PrevPaths:    2,
					count.NewPrevPaths: 2,
				},
				collections: expectCollections(
					true,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t)),
					aColl(
						d.fullPath(t, folderName()),
						d.fullPath(t, folderName()),
						fileID())),
				globalExcludedFileIDs: nil,
				prevPaths: map[string]string{
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbh := defaultOneDriveBH(user)
			mbh.DriveItemEnumeration = test.enumerator

			c := collWithMBH(mbh)
			globalExcludes := prefixmatcher.NewStringSetBuilder()

			results, prevPaths, _, err := c.makeDriveCollections(
				ctx,
				test.drive.able,
				test.prevPaths,
				deltaURL("prev"),
				newPagerLimiter(control.DefaultOptions()),
				globalExcludes,
				c.counter,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			test.expect.counts.Compare(t, c.counter)

			expectColls := test.expect.collections
			expectColls.compare(t, results)
			expectColls.requireNoUnseenCollections(t)

			assert.Equal(t, test.expect.prevPaths, prevPaths, "new previous paths")

			if test.expect.globalExcludedFileIDs != nil {
				excludes, ok := globalExcludes.Get(d.strPath(t))
				require.True(t, ok, "drive path produces a global file exclusions entry")
				assert.Equal(t, test.expect.globalExcludedFileIDs, excludes)
			} else {
				_, ok := globalExcludes.Get(d.strPath(t))
				require.False(t, ok, "drive path does not produce a global file exclusions entry")
			}
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddPrevPathsToTree_errors() {
	d := drive()
	t := suite.T()

	table := []struct {
		name      string
		tree      func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		prevPaths map[string]string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "no error - normal usage",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				folderID("parent"): d.strPath(t, folderName("parent")),
				folderID():         d.strPath(t, folderName("parent"), folderName()),
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
				id("santa"): d.strPath(t, name("santa")),
			},
			expectErr: require.NoError,
		},
		{
			name: "empty key in previous paths",
			tree: treeWithFolders,
			prevPaths: map[string]string{
				"": d.strPath(t, folderName("parent")),
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

			tree := test.tree(t, d)

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
	d := drive()
	t := suite.T()

	type expected struct {
		prevPaths             map[string]string
		collections           expectedCollections
		globalExcludedFileIDs map[string]struct{}
	}

	table := []struct {
		name           string
		tree           func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
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
					rootID:             d.strPath(t),
					folderID("parent"): d.strPath(t, folderName("parent")),
					folderID():         d.strPath(t, folderName("parent"), folderName()),
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						nil,
						fileID("r")),
					aColl(
						d.fullPath(t, folderName("parent")),
						nil,
						fileID("p")),
					aColl(
						d.fullPath(t, folderName("parent"), folderName()),
						nil,
						fileID("f"))),
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID("f")),
			},
		},
		{
			name:           "all folders moved",
			tree:           fullTree,
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                d.strPath(t),
				folderID("parent"):    d.strPath(t, folderName("parent-prev")),
				folderID():            d.strPath(t, folderName("parent-prev"), folderName()),
				folderID("prev"):      d.strPath(t, folderName("parent-prev"), folderName("prev")),
				folderID("prev-chld"): d.strPath(t, folderName("parent-prev"), folderName("prev"), folderName("prev-chld")),
				folderID("tombstone"): d.strPath(t, folderName("tombstone-prev")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:                d.strPath(t),
					folderID("parent"):    d.strPath(t, folderName("parent")),
					folderID():            d.strPath(t, folderName("parent"), folderName()),
					folderID("prev"):      d.strPath(t, folderName("parent"), folderName("prev")),
					folderID("prev-chld"): d.strPath(t, folderName("parent"), folderName("prev"), folderName("prev-chld")),
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t),
						fileID("r")),
					aColl(
						d.fullPath(t, folderName("parent")),
						d.fullPath(t, folderName("parent-prev")),
						fileID("p")),
					aColl(
						d.fullPath(t, folderName("parent"), folderName()),
						d.fullPath(t, folderName("parent-prev"), folderName()),
						fileID("f")),
					aColl(nil, d.fullPath(t, folderName("tombstone-prev")))),
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID("f")),
			},
		},
		{
			name:           "all folders moved - path separator string check",
			tree:           fullTreeWithNames("pa/rent", "to/mbstone"),
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                 d.strPath(t),
				folderID("pa/rent"):    d.strPath(t, folderName("parent/prev")),
				folderID():             d.strPath(t, folderName("parent/prev"), folderName()),
				folderID("pr/ev"):      d.strPath(t, folderName("parent/prev"), folderName("pr/ev")),
				folderID("prev/chld"):  d.strPath(t, folderName("parent/prev"), folderName("pr/ev"), folderName("prev/chld")),
				folderID("to/mbstone"): d.strPath(t, folderName("tombstone/prev")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:                d.strPath(t),
					folderID("pa/rent"):   d.strPath(t, folderName("pa/rent")),
					folderID():            d.strPath(t, folderName("pa/rent"), folderName()),
					folderID("pr/ev"):     d.strPath(t, folderName("pa/rent"), folderName("pr/ev")),
					folderID("prev/chld"): d.strPath(t, folderName("pa/rent"), folderName("pr/ev"), folderName("prev/chld")),
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t),
						fileID("r")),
					aColl(
						d.fullPath(t, folderName("pa/rent")),
						d.fullPath(t, folderName("parent/prev")),
						fileID("p")),
					aColl(
						d.fullPath(t, folderName("pa/rent"), folderName()),
						d.fullPath(t, folderName("parent/prev"), folderName()),
						fileID("f")),
					aColl(nil, d.fullPath(t, folderName("tombstone/prev")))),
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID("f")),
			},
		},
		{
			name: "nothing in the tree was moved " +
				"but there were some folders in the previous paths that " +
				"didn't appear in the delta so those have to appear in the " +
				"new previous paths but those weren't moved either so " +
				"everything should have the same path at the end",
			tree:           fullTree,
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                d.strPath(t),
				folderID("parent"):    d.strPath(t, folderName("parent")),
				folderID():            d.strPath(t, folderName("parent"), folderName()),
				folderID("tombstone"): d.strPath(t, folderName("tombstone")),
				folderID("prev"):      d.strPath(t, folderName("prev")),
				folderID("prev-chld"): d.strPath(t, folderName("prev"), folderName("prev-chld")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:                d.strPath(t),
					folderID("parent"):    d.strPath(t, folderName("parent")),
					folderID():            d.strPath(t, folderName("parent"), folderName()),
					folderID("prev"):      d.strPath(t, folderName("prev")),
					folderID("prev-chld"): d.strPath(t, folderName("prev"), folderName("prev-chld")),
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t),
						fileID("r")),
					aColl(
						d.fullPath(t, folderName("parent")),
						d.fullPath(t, folderName("parent")),
						fileID("p")),
					aColl(
						d.fullPath(t, folderName("parent"), folderName()),
						d.fullPath(t, folderName("parent"), folderName()),
						fileID("f")),
					aColl(nil, d.fullPath(t, folderName("tombstone")))),
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID("f")),
			},
		},
		{
			name: "nothing in the tree was moved " +
				"but there were some folders in the previous paths that " +
				"didn't appear in the delta so those have to appear in the " +
				"new previous paths but those weren't moved either so " +
				"everything should have the same path at the end " +
				"- the version with path separators chars in the directory names",
			tree:           fullTreeWithNames("pa/rent", "to/mbstone"),
			enableURLCache: true,
			prevPaths: map[string]string{
				rootID:                 d.strPath(t),
				folderID("pa/rent"):    d.strPath(t, folderName("pa/rent")),
				folderID():             d.strPath(t, folderName("pa/rent"), folderName()),
				folderID("pr/ev"):      d.strPath(t, folderName("pa/rent"), folderName("pr/ev")),
				folderID("prev/chld"):  d.strPath(t, folderName("pa/rent"), folderName("pr/ev"), folderName("prev/chld")),
				folderID("to/mbstone"): d.strPath(t, folderName("to/mbstone")),
			},
			expect: expected{
				prevPaths: map[string]string{
					rootID:                d.strPath(t),
					folderID("pa/rent"):   d.strPath(t, folderName("pa/rent")),
					folderID():            d.strPath(t, folderName("pa/rent"), folderName()),
					folderID("pr/ev"):     d.strPath(t, folderName("pa/rent"), folderName("pr/ev")),
					folderID("prev/chld"): d.strPath(t, folderName("pa/rent"), folderName("pr/ev"), folderName("prev/chld")),
				},
				collections: expectCollections(
					false,
					true,
					aColl(
						d.fullPath(t),
						d.fullPath(t),
						fileID("r")),
					aColl(
						d.fullPath(t, folderName("pa/rent")),
						d.fullPath(t, folderName("pa/rent")),
						fileID("p")),
					aColl(
						d.fullPath(t, folderName("pa/rent"), folderName()),
						d.fullPath(t, folderName("pa/rent"), folderName()),
						fileID("f")),
					aColl(nil, d.fullPath(t, folderName("to/mbstone")))),
				globalExcludedFileIDs: makeExcludeMap(
					fileID("r"),
					fileID("p"),
					fileID("d"),
					fileID("f")),
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			tree := test.tree(t, d)

			err := addPrevPathsToTree(ctx, tree, test.prevPaths, fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			c := collWithMBH(defaultOneDriveBH(user))

			countPages := 9001
			if test.enableURLCache {
				countPages = 1
			}

			colls, newPrevPaths, excluded, err := c.turnTreeIntoCollections(
				ctx,
				tree,
				d.able,
				test.prevPaths,
				deltaURL(),
				countPages,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.prevPaths, newPrevPaths, "new previous paths")

			expectColls := test.expect.collections
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
	enumerator enumerateDriveItemsDelta
	tree       func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
	limiter    *pagerLimiter
	expect     populateTreeExpected
}

// this test focuses on the population of a tree using a single delta's enumeration data.
// It is not concerned with unifying previous paths or post-processing collections.
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree_singleDelta() {
	d := drive()

	table := []populateTreeTest{
		{
			name: "nil page",
			tree: newTree,
			// special case enumerator to generate a null page.
			// otherwise all enumerators should be DriveEnumerator()s.
			enumerator: enumerateDriveItemsDelta{
				DrivePagers: map[string]*DeltaDriveEnumerator{
					d.id: {
						Drive: d,
						DeltaQueries: []*deltaQuery{{
							Pages:       nil,
							DeltaUpdate: pagers.DeltaUpdate{URL: deltaURL()},
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(d.folderAt(root)),
						aPage(d.folderAt(root, "sib")),
						aPage(
							d.folderAt(root),
							d.folderAt(folder, "chld"))))),
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder)),
						aPage(
							d.folderAt(root, "sib"),
							d.fileAt("sib", "fsib")),
						aPage(
							d.folderAt(root),
							d.folderAt(folder, "chld"),
							d.fileAt("chld", "fchld"))))),
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
				sizeBytes:      3 * defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(delItem(folderID(), folderID("parent"), isFolder))))),
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(aPage(
						delItem(folderID(), folderID("parent"), isFolder))))),
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder)),
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root, "parent"),
							driveItem(folderID(), folderName("moved"), d.dir(), folderID("parent"), isFolder),
							driveFile(d.dir(folderName("parent"), folderName()), folderID())),
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(delItem(folderID(), rootID, isFolder)),
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
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
				sizeBytes:      defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(delItem(folderID(), rootID, isFolder)),
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
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
				sizeBytes:      defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder)),
						aPage(
							d.folderAt(root, "sib"),
							d.fileAt("sib", "fsib")),
						aPage(
							d.folderAt(root),
							d.folderAt(folder, "chld"),
							d.fileAt("chld", "fchld"))))),
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
				sizeBytes:      defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder)),
						aPage(
							d.folderAt(root, "sib"),
							d.fileAt("sib", "fsib")),
						aPage(
							d.folderAt(root),
							d.folderAt(folder, "chld"),
							d.fileAt("chld", "fchld"))))),
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
			runPopulateTreeTest(suite.T(), d.able, test)
		})
	}
}

// this test focuses on quirks that can only arise from cases that occur across
// multiple delta enumerations.
// It is not concerned with unifying previous paths or post-processing collections.
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree_multiDelta() {
	d := drive()

	table := []populateTreeTest{
		{
			name: "sanity case: normal enumeration split across multiple deltas",
			tree: newTree,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).
						with(aPage(
							d.folderAt(root),
							d.fileAt(folder))),
					delta(nil).
						with(aPage(
							d.folderAt(root, "sib"),
							d.fileAt("sib", "fsib"))),
					delta(nil).
						with(aPage(
							d.folderAt(root),
							d.folderAt(folder, "chld"),
							d.fileAt("chld", "fchld"))))),
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
				sizeBytes:      3 * defaultFileSize,
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
			name: "file moved multiple times",
			tree: newTree,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).
						with(aPage(
							d.folderAt(root),
							d.fileAt(folder))),
					delta(nil).
						with(aPage(
							d.fileAt(root))),
					delta(nil).
						with(aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: populateTreeExpected{
				counts: countTD.Expected{
					count.TotalDeltasProcessed:        4,
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalDeleteFilesProcessed:   0,
					count.TotalFilesProcessed:         3,
					count.TotalFoldersProcessed:       5,
					count.TotalPagesEnumerated:        4,
				},
				err:            require.NoError,
				numLiveFiles:   1,
				numLiveFolders: 2,
				sizeBytes:      1 * defaultFileSize,
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
			name: "create->delete,create",
			tree: newTree,
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder))),
					// a (delete,create) pair in the same delta can occur when
					// a user deletes and restores an item in-between deltas.
					delta(nil).with(
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), folderID(), isFile)),
						aPage(
							d.folderAt(root),
							d.fileAt(folder))))),
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
				sizeBytes:      defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							d.folderAt(root),
							d.fileAt(folder))),
					delta(nil).with(
						aPage(
							driveItem(folderID(), folderName("rename"), d.dir(), rootID, isFolder),
							driveItem(fileID(), fileName("rename"), d.dir(folderName("rename")), folderID(), isFile))))),
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
				sizeBytes:      defaultFileSize,
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
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						// first page: create /root/folder and /root/folder/file
						aPage(
							d.folderAt(root),
							d.fileAt(folder)),
						// assume the user makes changes at this point:
						// * create a new /root/folder
						// * move /root/folder/file from old to new folder (same file ID)
						// * delete /root/folder
						// in drive deltas, this will show up as another folder creation sharing
						// the same dirname, but we won't see the delete until...
						aPage(
							driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
							driveItem(fileID(), fileName(), d.dir(folderName()), folderID(2), isFile))),
					// the next delta, containing the delete marker for the original /root/folder
					delta(nil).with(
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
				sizeBytes:      defaultFileSize,
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
			runPopulateTreeTest(suite.T(), d.able, test)
		})
	}
}

func runPopulateTreeTest(
	t *testing.T,
	d models.Driveable,
	test populateTreeTest,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mbh     = defaultDriveBHWith(user, test.enumerator)
		c       = collWithMBH(mbh)
		counter = count.New()
		tree    = test.tree(t, drive())
	)

	_, _, err := c.populateTree(
		ctx,
		tree,
		d,
		deltaURL(),
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
	d := drive()

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
		tree    func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		page    nextPage
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "nil page",
			tree:    treeWithRoot,
			page:    nextPage{},
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
			page:    nextPage{Items: []models.DriveItemable{}},
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
				d.folderAt(root),
				d.folderAt(folder, "chld"),
				d.folderAt(root, "sib")),
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
				d.folderAt(root),
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
				d.folderAt(root, "parent"),
				driveItem(folderID(), folderName("moved"), d.dir(folderName("parent")), folderID("parent"), isFolder),
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
				d.folderAt(root)),
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
				d.folderAt(root)),
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
				c       = collWithMBH(defaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t, d)
			)

			err := c.enumeratePageOfItems(
				ctx,
				tree,
				d.able,
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
		d      = drive()
		fld    = custom.ToCustomDriveItem(d.folderAt(root))
		subFld = custom.ToCustomDriveItem(driveFolder(d.dir(folderName("parent")), folderID("parent")))
		pack   = custom.ToCustomDriveItem(driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage))
		del    = custom.ToCustomDriveItem(delItem(folderID(), rootID, isFolder))
		mal    = custom.ToCustomDriveItem(malwareItem(folderID("mal"), folderName("mal"), d.dir(), rootID, isFolder))
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
		tree    func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
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
				c       = collWithMBH(defaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t, d)
			)

			skipped, err := c.addFolderToTree(
				ctx,
				tree,
				d.able,
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
	d := drive()

	basePath, err := path.Build(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DriveFolderPrefixBuilder(d.id).Elements()...)
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
			folder:    rootFolder(),
			expect:    basePath.String(),
			expectErr: require.NoError,
		},
		{
			name:      "folder",
			folder:    d.folderAt(root),
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
			c := collWithMBH(defaultOneDriveBH(user))

			ctx, flush := tester.NewContext(t)
			defer flush()

			p, err := c.makeFolderCollectionPath(
				ctx,
				d.id,
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
	d := drive()

	type expected struct {
		counts                        countTD.Expected
		err                           require.ErrorAssertionFunc
		treeContainsFileIDsWithParent map[string]string
		countLiveFiles                int
		countTotalBytes               int64
	}

	table := []struct {
		name   string
		tree   func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		page   nextPage
		expect expected
	}{
		{
			name: "one file at root",
			tree: treeWithRoot,
			page: aPage(d.fileAt(root)),
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
				countTotalBytes: defaultFileSize,
			},
		},
		{
			name: "many files in a hierarchy",
			tree: treeWithRoot,
			page: aPage(
				d.fileAt(root),
				d.folderAt(root),
				d.fileAt(folder, "fchld")),
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
				countTotalBytes: defaultFileSize * 2,
			},
		},
		{
			name: "many updates to the same file",
			tree: treeWithRoot,
			page: aPage(
				d.fileAt(root),
				driveItem(fileID(), fileName(1), d.dir(), rootID, isFile),
				driveItem(fileID(), fileName(2), d.dir(), rootID, isFile)),
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
				countTotalBytes: defaultFileSize,
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
				d.fileAt(root),
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
				d.folderAt(root),
				d.fileAt(folder),
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
				d.fileAt(root)),
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
				countTotalBytes: defaultFileSize,
			},
		},
		{
			name: "file not in tree: delete->restore",
			tree: treeWithRoot,
			page: aPage(
				delItem(fileID(), rootID, isFile),
				d.fileAt(root)),
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
				countTotalBytes: defaultFileSize,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(defaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t, d)
			)

			err := c.enumeratePageOfItems(
				ctx,
				tree,
				d.able,
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
	d := drive()

	unlimitedItemsPerContainer := newPagerLimiter(minimumLimitOpts())
	unlimitedItemsPerContainer.limits.MaxItemsPerContainer = 9001

	unlimitedTotalBytesAndFiles := newPagerLimiter(minimumLimitOpts())
	unlimitedTotalBytesAndFiles.limits.MaxBytes = 9001
	unlimitedTotalBytesAndFiles.limits.MaxItems = 9001

	type expected struct {
		counts                        countTD.Expected
		err                           require.ErrorAssertionFunc
		shouldHitLimit                bool
		shouldHitCollLimit            bool
		skipped                       assert.ValueAssertionFunc
		treeContainsFileIDsWithParent map[string]string
		countLiveFiles                int
		countTotalBytes               int64
	}

	table := []struct {
		name    string
		tree    func(t *testing.T, d *deltaDrive) *folderyMcFolderFace
		file    models.DriveItemable
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "add new file",
			tree:    treeWithRoot,
			file:    d.fileAt(root),
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
				countTotalBytes: defaultFileSize,
			},
		},
		{
			name:    "duplicate file",
			tree:    treeWithFileAtRoot,
			file:    d.fileAt(root),
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
				countTotalBytes: defaultFileSize,
			},
		},
		{
			name:    "error file seen before parent",
			tree:    treeWithRoot,
			file:    d.fileAt(folder),
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
			file:    malwareItem(fileID(), fileName(), d.dir(folderName()), rootID, isFile),
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
			file:    d.fileAt(root, 2),
			limiter: unlimitedTotalBytesAndFiles,
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:                require.Error,
				shouldHitCollLimit: true,
				skipped:            assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					fileID(): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultFileSize,
			},
		},
		{
			name:    "goes over total byte limit",
			tree:    treeWithRoot,
			file:    d.fileAt(root),
			limiter: unlimitedItemsPerContainer,
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				// no error here, since byte limit shouldn't
				// make the func return an error.
				err:                           require.NoError,
				skipped:                       assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{},
				countLiveFiles:                0,
				countTotalBytes:               0,
			},
		},
		{
			name:    "already over total byte limit",
			tree:    treeWithFileAtRoot,
			file:    d.fileAt(root, 2),
			limiter: unlimitedItemsPerContainer,
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
				countTotalBytes: defaultFileSize,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				c       = collWithMBH(defaultOneDriveBH(user))
				counter = count.New()
				tree    = test.tree(t, d)
			)

			skipped, err := c.addFileToTree(
				ctx,
				tree,
				d.able,
				custom.ToCustomDriveItem(test.file),
				test.limiter,
				counter)

			test.expect.err(t, err, clues.ToCore(err))
			test.expect.skipped(t, skipped)

			if test.expect.shouldHitCollLimit {
				require.ErrorIs(t, err, errHitCollectionLimit, clues.ToCore(err))
			}

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
