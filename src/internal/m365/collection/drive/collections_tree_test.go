package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	countTD "github.com/alcionai/corso/src/pkg/count/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func collWithMBH(mbh BackupHandler) *Collections {
	return NewCollections(
		mbh,
		tenant,
		idname.NewProvider(user, user),
		func(*support.ControllerOperationStatus) {},
		control.Options{ToggleFeatures: control.Toggles{
			UseDeltaTree: true,
		}},
		count.New())
}

func fullOrPrevPath(
	t *testing.T,
	coll data.BackupCollection,
) path.Path {
	var collPath path.Path

	if coll.State() != data.DeletedState {
		collPath = coll.FullPath()
	} else {
		collPath = coll.PreviousPath()
	}

	require.False(
		t,
		len(collPath.Elements()) < 4,
		"malformed or missing collection path")

	return collPath
}

func pagerForDrives(drives ...models.Driveable) *apiMock.Pager[models.Driveable] {
	return &apiMock.Pager[models.Driveable]{
		ToReturn: []apiMock.PagerResult[models.Driveable]{
			{Values: drives},
		},
	}
}

func makePrevMetadataColls(
	t *testing.T,
	mbh BackupHandler,
	previousPaths map[string]map[string]string,
) []data.RestoreCollection {
	pathPrefix, err := mbh.MetadataPathPrefix(tenant)
	require.NoError(t, err, clues.ToCore(err))

	prevDeltas := map[string]string{}

	for driveID := range previousPaths {
		prevDeltas[driveID] = idx(delta, "prev")
	}

	mdColl, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(bupMD.DeltaURLsFileName, prevDeltas),
			graph.NewMetadataEntry(bupMD.PreviousPathFileName, previousPaths),
		},
		func(*support.ControllerOperationStatus) {},
		count.New())
	require.NoError(t, err, "creating metadata collection", clues.ToCore(err))

	return []data.RestoreCollection{
		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
	}
}

func compareMetadata(
	t *testing.T,
	mdColl data.Collection,
	expectDeltas map[string]string,
	expectPrevPaths map[string]map[string]string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	colls := []data.RestoreCollection{
		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
	}

	deltas, prevs, _, err := deserializeAndValidateMetadata(
		ctx,
		colls,
		count.New(),
		fault.New(true))
	require.NoError(t, err, "deserializing metadata", clues.ToCore(err))
	assert.Equal(t, expectDeltas, deltas, "delta urls")
	assert.Equal(t, expectPrevPaths, prevs, "previous paths")
}

// for comparisons done by collection state
type stateAssertion struct {
	itemIDs []string
	// should never get set by the user.
	// this flag gets flipped when calling assertions.compare.
	// any unseen collection will error on requireNoUnseenCollections
	sawCollection bool
}

// for comparisons done by a given collection path
type collectionAssertion struct {
	doNotMerge    assert.BoolAssertionFunc
	states        map[data.CollectionState]*stateAssertion
	excludedItems map[string]struct{}
}

type statesToItemIDs map[data.CollectionState][]string

// TODO(keepers): move excludeItems to a more global position.
func newCollAssertion(
	doNotMerge bool,
	itemsByState statesToItemIDs,
	excludeItems ...string,
) collectionAssertion {
	states := map[data.CollectionState]*stateAssertion{}

	for state, itemIDs := range itemsByState {
		states[state] = &stateAssertion{
			itemIDs: itemIDs,
		}
	}

	dnm := assert.False
	if doNotMerge {
		dnm = assert.True
	}

	return collectionAssertion{
		doNotMerge:    dnm,
		states:        states,
		excludedItems: makeExcludeMap(excludeItems...),
	}
}

// to aggregate all collection-related expectations in the backup
// map collection path -> collection state -> assertion
type collectionAssertions map[string]collectionAssertion

// ensure the provided collection matches expectations as set by the test.
func (cas collectionAssertions) compare(
	t *testing.T,
	coll data.BackupCollection,
	excludes *prefixmatcher.StringSetMatchBuilder,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		itemCh  = coll.Items(ctx, fault.New(true))
		itemIDs = []string{}
	)

	p := fullOrPrevPath(t, coll)

	for itm := range itemCh {
		itemIDs = append(itemIDs, itm.ID())
	}

	expect := cas[p.String()]
	expectState := expect.states[coll.State()]
	expectState.sawCollection = true

	assert.ElementsMatchf(
		t,
		expectState.itemIDs,
		itemIDs,
		"expected all items to match in collection with:\nstate %q\npath %q",
		coll.State(),
		p)

	expect.doNotMerge(
		t,
		coll.DoNotMergeItems(),
		"expected collection to have the appropariate doNotMerge flag")

	if result, ok := excludes.Get(p.String()); ok {
		assert.Equal(
			t,
			expect.excludedItems,
			result,
			"excluded items")
	}
}

// ensure that no collections in the expected set are still flagged
// as sawCollection == false.
func (cas collectionAssertions) requireNoUnseenCollections(
	t *testing.T,
) {
	for p, withPath := range cas {
		for _, state := range withPath.states {
			require.True(
				t,
				state.sawCollection,
				"results should have contained collection:\n\t%q\t\n%q",
				state, p)
		}
	}
}

func rootAnd(items ...models.DriveItemable) []models.DriveItemable {
	return append([]models.DriveItemable{driveItem(rootID, rootName, parent(0), "", isFolder)}, items...)
}

func pagesOf(pages ...[]models.DriveItemable) []mock.NextPage {
	mnp := []mock.NextPage{}

	for _, page := range pages {
		mnp = append(mnp, mock.NextPage{Items: page})
	}

	return mnp
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

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

// this test is primarily aimed at multi-drive handling
// more complicated single-drive tests are in _MakeDriveCollections
func (suite *CollectionsTreeUnitSuite) TestCollections_GetTree() {
	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	type expected struct {
		canUsePrevBackup assert.BoolAssertionFunc
		collAssertions   collectionAssertions
		counts           countTD.Expected
		deltas           map[string]string
		err              require.ErrorAssertionFunc
		prevPaths        map[string]map[string]string
		skips            int
	}

	table := []struct {
		name          string
		drivePager    *apiMock.Pager[models.Driveable]
		enumerator    mock.EnumerateItemsDeltaByDrive
		previousPaths map[string]map[string]string

		metadata []data.RestoreCollection
		expect   expected
	}{
		{
			name:       "not yet implemented",
			drivePager: pagerForDrives(drive1),
			expect: expected{
				canUsePrevBackup: assert.False,
				collAssertions: collectionAssertions{
					fullPath(1): newCollAssertion(
						doNotMergeItems,
						statesToItemIDs{data.NotMovedState: {}},
						id(file)),
				},
				counts: countTD.Expected{
					count.PrevPaths: 0,
				},
				deltas:    map[string]string{},
				err:       require.Error,
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
				mbh            = mock.DefaultDriveBHWith(user, test.drivePager, test.enumerator)
				c              = collWithMBH(mbh)
				prevMetadata   = makePrevMetadataColls(t, mbh, test.previousPaths)
				globalExcludes = prefixmatcher.NewStringSetBuilder()
				errs           = fault.New(true)
			)

			colls, canUsePrevBackup, err := c.getTree(
				ctx,
				prevMetadata,
				globalExcludes,
				errs)

			test.expect.err(t, err, clues.ToCore(err))
			// TODO(keepers): awaiting implementation
			assert.Empty(t, colls)
			assert.Equal(t, test.expect.skips, len(errs.Skipped()))
			test.expect.canUsePrevBackup(t, canUsePrevBackup)
			test.expect.counts.Compare(t, c.counter)

			if err != nil {
				return
			}

			for _, coll := range colls {
				collPath := fullOrPrevPath(t, coll)

				if collPath.String() == metadataPath.String() {
					compareMetadata(
						t,
						coll,
						test.expect.deltas,
						test.expect.prevPaths)

					continue
				}

				test.expect.collAssertions.compare(t, coll, globalExcludes)
			}

			test.expect.collAssertions.requireNoUnseenCollections(t)
		})
	}
}

// this test is primarily aimed exercising the full breadth of single-drive delta enumeration
// and broad contracts
// more granular testing can be found in the functions below.
func (suite *CollectionsTreeUnitSuite) TestCollections_MakeDriveCollections() {
	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	table := []struct {
		name         string
		c            *Collections
		drive        models.Driveable
		prevPaths    map[string]string
		expectErr    require.ErrorAssertionFunc
		expectCounts countTD.Expected
	}{
		{
			name:      "not yet implemented",
			c:         collWithMBH(mock.DefaultOneDriveBH(user)),
			drive:     drive1,
			expectErr: require.Error,
			expectCounts: countTD.Expected{
				count.PrevPaths: 0,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			colls, paths, delta, err := test.c.makeDriveCollections(
				ctx,
				test.drive,
				test.prevPaths,
				idx(delta, "prev"),
				test.c.counter,
				fault.New(true))

			// TODO(keepers): awaiting implementation
			test.expectErr(t, err, clues.ToCore(err))
			assert.Empty(t, colls)
			assert.Empty(t, paths)
			assert.Empty(t, delta.URL)

			test.expectCounts.Compare(t, test.c.counter)
		})
	}
}

// this test focuses on the population of a tree using delta enumeration data,
// and is not concerned with unifying previous paths or post-processing collections
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	type expected struct {
		counts                   countTD.Expected
		err                      require.ErrorAssertionFunc
		treeSize                 int
		treeContainsFolderIDs    []string
		treeContainsTombstoneIDs []string
	}

	table := []struct {
		name       string
		enumerator mock.EnumerateItemsDeltaByDrive
		tree       *folderyMcFolderFace
		limiter    *pagerLimiter
		expect     expected
	}{
		{
			name: "nil page",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       nil,
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts:                   countTD.Expected{},
				err:                      require.NoError,
				treeSize:                 0,
				treeContainsFolderIDs:    []string{},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "root only",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(rootAnd()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 1,
					count.PagesEnumerated:       1,
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
			name: "root only on two pages",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(rootAnd(), rootAnd()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 2,
					count.PagesEnumerated:       2,
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
			name: "one",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 2,
					count.PagesEnumerated:       1,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "many",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
							rootAnd(driveItem(idx(folder, "sib"), namex(folder, "sib"), parent(0), rootID, isFolder)),
							rootAnd(
								driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parent(0), id(folder), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 7,
					count.PagesEnumerated:       3,
				},
				err:      require.NoError,
				treeSize: 4,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
					idx(folder, "sib"),
					idx(folder, "chld"),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "tombstone with unpopulated tree",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(rootAnd(delItem(id(folder), parent(0), idx(folder, "parent"), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 1,
					count.PagesEnumerated:             1,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
			},
		},
		{
			name: "tombstone with populated tree",
			tree: treeWithFolders(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(rootAnd(delItem(id(folder), parent(0), idx(folder, "parent"), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 1,
					count.PagesEnumerated:             1,
				},
				err:      require.NoError,
				treeSize: 3,
				treeContainsFolderIDs: []string{
					rootID,
					idx(folder, "parent"),
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
			},
		},
		{
			name: "create->delete",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
							rootAnd(delItem(id(folder), parent(0), rootID, isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       3,
					count.TotalDeleteFoldersProcessed: 1,
					count.PagesEnumerated:             2,
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
			name: "move->delete with populated tree",
			tree: treeWithFolders(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(folder, "parent"), namex(folder, "parent"), parent(0), rootID, isFolder),
								driveItem(id(folder), namex(folder, "moved"), parent(0), idx(folder, "parent"), isFolder)),
							rootAnd(delItem(id(folder), parent(0), idx(folder, "parent"), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       4,
					count.TotalDeleteFoldersProcessed: 1,
					count.PagesEnumerated:             2,
				},
				err:      require.NoError,
				treeSize: 3,
				treeContainsFolderIDs: []string{
					rootID,
					idx(folder, "parent"),
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
			},
		},
		{
			name: "delete->create with previous path",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(delItem(id(folder), parent(0), rootID, isFolder)),
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFoldersProcessed:       3,
					count.PagesEnumerated:             2,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "delete->create without previous path",
			tree: treeWithRoot(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(delItem(id(folder), parent(0), rootID, isFolder)),
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFoldersProcessed:       3,
					count.PagesEnumerated:             2,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "at limit before enumeration",
			tree: treeWithRoot(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
							rootAnd(driveItem(idx(folder, "sib"), namex(folder, "sib"), parent(0), rootID, isFolder)),
							rootAnd(
								driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parent(0), id(folder), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.PagesEnumerated:             1,
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
			name: "hit limit during enumeration",
			tree: newFolderyMcFolderFace(nil),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
							rootAnd(driveItem(idx(folder, "sib"), namex(folder, "sib"), parent(0), rootID, isFolder)),
							rootAnd(
								driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parent(0), id(folder), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.PagesEnumerated:             1,
				},
				err:      require.NoError,
				treeSize: 1,
				treeContainsFolderIDs: []string{
					rootID,
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

			mockDrivePager := &apiMock.Pager[models.Driveable]{
				ToReturn: []apiMock.PagerResult[models.Driveable]{
					{Values: []models.Driveable{drv}},
				},
			}

			mbh := mock.DefaultDriveBHWith(user, mockDrivePager, test.enumerator)
			c := collWithMBH(mbh)
			counter := count.New()

			_, err := c.populateTree(
				ctx,
				test.tree,
				test.limiter,
				&driveEnumerationStats{},
				drv,
				id(delta),
				counter,
				fault.New(true))
			test.expect.err(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.treeSize, test.tree.CountFolders(), "count folders in tree")
			test.expect.counts.Compare(t, counter)

			for _, id := range test.expect.treeContainsFolderIDs {
				require.NotNil(t, test.tree.folderIDToNode[id], "node exists")
			}

			for _, id := range test.expect.treeContainsTombstoneIDs {
				require.NotNil(t, test.tree.tombstones[id], "tombstone exists")
			}
		})
	}
}

// this test focuses on folder assertions when enumerating a page of items
// item-specific assertions are in the next test
func (suite *CollectionsTreeUnitSuite) TestCollections_EnumeratePageOfItems_folders() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	type expected struct {
		counts                   countTD.Expected
		err                      require.ErrorAssertionFunc
		treeSize                 int
		treeContainsFolderIDs    []string
		treeContainsTombstoneIDs []string
	}

	table := []struct {
		name    string
		tree    *folderyMcFolderFace
		page    []models.DriveItemable
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "nil page",
			tree:    treeWithRoot(),
			page:    nil,
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
			tree:    treeWithRoot(),
			page:    []models.DriveItemable{},
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
			tree:    treeWithRoot(),
			page:    rootAnd(),
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
			name:    "one",
			tree:    treeWithRoot(),
			page:    rootAnd(driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 2,
				},
				err:      require.NoError,
				treeSize: 2,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "many",
			tree: treeWithRoot(),
			page: rootAnd(
				driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
				driveItem(idx(folder, "sib"), namex(folder, "sib"), parent(0), rootID, isFolder),
				driveItem(idx(folder, "chld"), namex(folder, "chld"), parent(0), id(folder), isFolder)),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 4,
				},
				err:      require.NoError,
				treeSize: 4,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
					idx(folder, "sib"),
					idx(folder, "chld"),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "already hit limit",
			tree: treeWithRoot(),
			page: rootAnd(
				driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
				driveItem(idx(folder, "sib"), namex(folder, "sib"), parent(0), rootID, isFolder),
				driveItem(idx(folder, "chld"), namex(folder, "chld"), parent(0), id(folder), isFolder)),
			limiter: newPagerLimiter(minimumLimitOpts()),
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
			name:    "tombstone",
			tree:    treeWithFolders(),
			page:    rootAnd(delItem(id(folder), parent(0), idx(folder, "parent"), isFolder)),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       1,
					count.TotalDeleteFoldersProcessed: 1,
				},
				err:      require.NoError,
				treeSize: 3,
				treeContainsFolderIDs: []string{
					rootID,
					idx(folder, "parent"),
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
			},
		},
		{
			name: "create->delete",
			tree: treeWithRoot(),
			page: rootAnd(
				driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
				delItem(id(folder), parent(0), rootID, isFolder)),
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
			tree: treeWithFolders(),
			page: rootAnd(
				driveItem(idx(folder, "parent"), namex(folder, "parent"), parent(0), rootID, isFolder),
				driveItem(id(folder), namex(folder, "moved"), parent(0), idx(folder, "parent"), isFolder),
				delItem(id(folder), parent(0), idx(folder, "parent"), isFolder)),
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
					idx(folder, "parent"),
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
			},
		},
		{
			name: "delete->create with previous path",
			tree: treeWithRoot(),
			page: rootAnd(
				delItem(id(folder), parent(0), rootID, isFolder),
				driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
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
					id(folder),
				},
				treeContainsTombstoneIDs: []string{},
			},
		},
		{
			name: "delete->create without previous path",
			tree: treeWithRoot(),
			page: rootAnd(
				delItem(id(folder), parent(0), rootID, isFolder),
				driveItem(id(folder), name(folder), parent(0), rootID, isFolder)),
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
					id(folder),
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

			c := collWithMBH(mock.DefaultOneDriveBH(user))
			counter := count.New()

			err := c.enumeratePageOfItems(
				ctx,
				test.tree,
				test.limiter,
				&driveEnumerationStats{},
				drv,
				test.page,
				counter,
				fault.New(true))
			test.expect.err(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.treeSize, test.tree.CountFolders(), "count folders in tree")
			test.expect.counts.Compare(t, counter)

			for _, id := range test.expect.treeContainsFolderIDs {
				assert.NotNil(t, test.tree.folderIDToNode[id], "node exists")
			}

			for _, id := range test.expect.treeContainsTombstoneIDs {
				assert.NotNil(t, test.tree.tombstones[id], "tombstone exists")
			}
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddFolderToTree() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	fld := driveItem(id(folder), name(folder), parent(0), rootID, isFolder)
	subFld := driveItem(id(folder), name(folder), parent(drv, namex(folder, "parent")), idx(folder, "parent"), isFolder)
	pack := driveItem(id(pkg), name(pkg), parent(0), rootID, isPackage)
	del := delItem(id(folder), parent(0), rootID, isFolder)
	mal := malwareItem(idx(folder, "mal"), namex(folder, "mal"), parent(0), rootID, isFolder)

	type expected struct {
		counts             countTD.Expected
		err                require.ErrorAssertionFunc
		treeSize           int
		treeContainsFolder assert.BoolAssertionFunc
	}

	table := []struct {
		name   string
		tree   *folderyMcFolderFace
		folder models.DriveItemable
		expect expected
	}{
		{
			name:   "add folder",
			tree:   treeWithRoot(),
			folder: fld,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalFoldersProcessed: 1},
				treeSize:           2,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:   "re-add folder that already exists",
			tree:   treeWithFolders(),
			folder: subFld,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalFoldersProcessed: 1},
				treeSize:           3,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:   "add package",
			tree:   treeWithRoot(),
			folder: pack,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalPackagesProcessed: 1},
				treeSize:           2,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:   "add tombstone",
			tree:   treeWithFolders(),
			folder: del,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalDeleteFoldersProcessed: 1},
				treeSize:           3,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:   "re-add tombstone that already exists",
			tree:   treeWithTombstone(),
			folder: del,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalDeleteFoldersProcessed: 1},
				treeSize:           2,
				treeContainsFolder: assert.True,
			},
		},
		{
			name:   "add malware",
			tree:   treeWithRoot(),
			folder: mal,
			expect: expected{
				err:                require.NoError,
				counts:             countTD.Expected{count.TotalMalwareProcessed: 1},
				treeSize:           1,
				treeContainsFolder: assert.False,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			c := collWithMBH(mock.DefaultOneDriveBH(user))
			counter := count.New()
			des := &driveEnumerationStats{}

			err := c.addFolderToTree(
				ctx,
				test.tree,
				drv,
				test.folder,
				des,
				counter,
				fault.New(true))
			test.expect.err(t, err, clues.ToCore(err))
			test.expect.counts.Compare(t, counter)
			assert.Equal(t, test.expect.treeSize, test.tree.CountFolders(), "folders in tree")
			test.expect.treeContainsFolder(t, test.tree.ContainsFolder(ptr.Val(test.folder.GetId())))
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_MakeFolderCollectionPath() {
	basePath, err := odConsts.DriveFolderPrefixBuilder(id(drive)).ToDataLayerOneDrivePath(tenant, user, false)
	require.NoError(suite.T(), err, clues.ToCore(err))

	folderPath, err := basePath.Append(false, name(folder))
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name      string
		folder    models.DriveItemable
		expect    string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "root",
			folder:    driveRootItem(rootID),
			expect:    basePath.String(),
			expectErr: require.NoError,
		},
		{
			name:      "folder",
			folder:    driveItem(id(folder), name(folder), parent(0), rootID, isFolder),
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
			c := collWithMBH(mock.DefaultOneDriveBH(user))

			p, err := c.makeFolderCollectionPath(id(drive), test.folder)
			test.expectErr(t, err, clues.ToCore(err))

			if err == nil {
				assert.Equal(t, test.expect, p.String())
			}
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddFileToTree() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	c := collWithMBH(mock.DefaultOneDriveBH(user))

	err := c.addFileToTree(
		ctx,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	require.ErrorContains(t, err, "not yet implemented", clues.ToCore(err))
}
