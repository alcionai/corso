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

func collWithMBHAndOpts(
	mbh BackupHandler,
	opts control.Options,
) *Collections {
	return NewCollections(
		mbh,
		tenant,
		idname.NewProvider(user, user),
		func(*support.ControllerOperationStatus) {},
		opts,
		count.New())
}

// func fullOrPrevPath(
// 	t *testing.T,
// 	coll data.BackupCollection,
// ) path.Path {
// 	var collPath path.Path

// 	if coll.State() != data.DeletedState {
// 		collPath = coll.FullPath()
// 	} else {
// 		collPath = coll.PreviousPath()
// 	}

// 	require.False(
// 		t,
// 		len(collPath.Elements()) < 4,
// 		"malformed or missing collection path")

// 	return collPath
// }

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

// func compareMetadata(
// 	t *testing.T,
// 	mdColl data.Collection,
// 	expectDeltas map[string]string,
// 	expectPrevPaths map[string]map[string]string,
// ) {
// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	colls := []data.RestoreCollection{
// 		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
// 	}

// 	deltas, prevs, _, err := deserializeAndValidateMetadata(
// 		ctx,
// 		colls,
// 		count.New(),
// 		fault.New(true))
// 	require.NoError(t, err, "deserializing metadata", clues.ToCore(err))
// 	assert.Equal(t, expectDeltas, deltas, "delta urls")
// 	assert.Equal(t, expectPrevPaths, prevs, "previous paths")
// }

// for comparisons done by collection state
type stateAssertion struct {
	itemIDs []string
	// should never get set by the user.
	// this flag gets flipped when calling assertions.compare.
	// any unseen collection will error on requireNoUnseenCollections
	// sawCollection bool
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
// func (cas collectionAssertions) compare(
// 	t *testing.T,
// 	coll data.BackupCollection,
// 	excludes *prefixmatcher.StringSetMatchBuilder,
// ) {
// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	var (
// 		itemCh  = coll.Items(ctx, fault.New(true))
// 		itemIDs = []string{}
// 	)

// 	p := fullOrPrevPath(t, coll)

// 	for itm := range itemCh {
// 		itemIDs = append(itemIDs, itm.ID())
// 	}

// 	expect := cas[p.String()]
// 	expectState := expect.states[coll.State()]
// 	expectState.sawCollection = true

// 	assert.ElementsMatchf(
// 		t,
// 		expectState.itemIDs,
// 		itemIDs,
// 		"expected all items to match in collection with:\nstate %q\npath %q",
// 		coll.State(),
// 		p)

// 	expect.doNotMerge(
// 		t,
// 		coll.DoNotMergeItems(),
// 		"expected collection to have the appropariate doNotMerge flag")

// 	if result, ok := excludes.Get(p.String()); ok {
// 		assert.Equal(
// 			t,
// 			expect.excludedItems,
// 			result,
// 			"excluded items")
// 	}
// }

// ensure that no collections in the expected set are still flagged
// as sawCollection == false.
// func (cas collectionAssertions) requireNoUnseenCollections(
// 	t *testing.T,
// ) {
// 	for p, withPath := range cas {
// 		for _, state := range withPath.states {
// 			require.True(
// 				t,
// 				state.sawCollection,
// 				"results should have contained collection:\n\t%q\t\n%q",
// 				state, p)
// 		}
// 	}
// }

func pageItems(items ...models.DriveItemable) []models.DriveItemable {
	return append([]models.DriveItemable{driveRootItem()}, items...)
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
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

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
			drivePager: pagerForDrives(drv),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(pageItems()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expect: expected{
				canUsePrevBackup: assert.False,
				collAssertions: collectionAssertions{
					driveFullPath(1): newCollAssertion(
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

			_, _, err := c.getTree(
				ctx,
				prevMetadata,
				globalExcludes,
				errs)

			test.expect.err(t, err, clues.ToCore(err))
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

// This test is primarily aimed exercising the full breadth of single-drive delta enumeration
// and broad contracts.
// More granular testing can be found in the lower level test functions below.
func (suite *CollectionsTreeUnitSuite) TestCollections_MakeDriveCollections() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	table := []struct {
		name         string
		drive        models.Driveable
		drivePager   *apiMock.Pager[models.Driveable]
		enumerator   mock.EnumerateItemsDeltaByDrive
		prevPaths    map[string]string
		expectErr    require.ErrorAssertionFunc
		expectCounts countTD.Expected
	}{
		{
			name:       "not yet implemented",
			drive:      drv,
			drivePager: pagerForDrives(drv),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(pageItems()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
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

			mbh := mock.DefaultDriveBHWith(user, test.drivePager, test.enumerator)
			c := collWithMBH(mbh)

			colls, paths, du, err := c.makeDriveCollections(
				ctx,
				test.drive,
				test.prevPaths,
				idx(delta, "prev"),
				newPagerLimiter(control.DefaultOptions()),
				c.counter,
				fault.New(true))

			// TODO(keepers): awaiting implementation
			test.expectErr(t, err, clues.ToCore(err))
			assert.Empty(t, colls)
			assert.Empty(t, paths)
			assert.Equal(t, id(delta), du.URL)

			test.expectCounts.Compare(t, c.counter)
		})
	}
}

// This test focuses on the population of a tree using delta enumeration data,
// and is not concerned with unifying previous paths or post-processing collections.
func (suite *CollectionsTreeUnitSuite) TestCollections_PopulateTree() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	type expected struct {
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

	table := []struct {
		name       string
		enumerator mock.EnumerateItemsDeltaByDrive
		tree       *folderyMcFolderFace
		limiter    *pagerLimiter
		expect     expected
	}{
		{
			name: "nil page",
			tree: newFolderyMcFolderFace(nil, rootID),
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
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(pageItems()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 1,
					count.TotalFilesProcessed:   0,
					count.PagesEnumerated:       1,
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
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages:       pagesOf(pageItems(), pageItems()),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 2,
					count.TotalFilesProcessed:   0,
					count.PagesEnumerated:       2,
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
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(driveItem(id(folder), name(folder), parentDir(), rootID, isFolder)),
							pageItems(driveItem(idx(folder, "sib"), namex(folder, "sib"), parentDir(), rootID, isFolder)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parentDir(), id(folder), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 7,
					count.PagesEnumerated:       3,
					count.TotalFilesProcessed:   0,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 4,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
					idx(folder, "sib"),
					idx(folder, "chld"),
				},
				treeContainsTombstoneIDs:      []string{},
				treeContainsFileIDsWithParent: map[string]string{},
			},
		},
		{
			name: "many folders with files",
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile)),
							pageItems(
								driveItem(idx(folder, "sib"), namex(folder, "sib"), parentDir(), rootID, isFolder),
								driveItem(idx(file, "sib"), namex(file, "sib"), parentDir(namex(folder, "sib")), idx(folder, "sib"), isFile)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parentDir(), id(folder), isFolder),
								driveItem(idx(file, "chld"), namex(file, "chld"), parentDir(namex(folder, "chld")), idx(folder, "chld"), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed: 7,
					count.TotalFilesProcessed:   3,
					count.PagesEnumerated:       3,
				},
				err:            require.NoError,
				numLiveFiles:   3,
				numLiveFolders: 4,
				sizeBytes:      3 * defaultItemSize,
				treeContainsFolderIDs: []string{
					rootID,
					id(folder),
					idx(folder, "sib"),
					idx(folder, "chld"),
				},
				treeContainsTombstoneIDs: []string{},
				treeContainsFileIDsWithParent: map[string]string{
					id(file):          id(folder),
					idx(file, "sib"):  idx(folder, "sib"),
					idx(file, "chld"): idx(folder, "chld"),
				},
			},
		},
		{
			// technically you won't see this behavior from graph deltas, since deletes always
			// precede creates/updates.  But it's worth checking that we can handle it anyways.
			name: "create, delete on next page",
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile)),
							pageItems(delItem(id(folder), parentDir(), rootID, isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       3,
					count.TotalFilesProcessed:         1,
					count.TotalDeleteFoldersProcessed: 1,
					count.PagesEnumerated:             2,
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
					id(file): id(folder),
				},
			},
		},
		{
			// technically you won't see this behavior from graph deltas, since deletes always
			// precede creates/updates.  But it's worth checking that we can handle it anyways.
			name: "move->delete folder with populated tree",
			tree: treeWithFolders(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(folder, "parent"), namex(folder, "parent"), parentDir(), rootID, isFolder),
								driveItem(id(folder), namex(folder, "moved"), parentDir(), idx(folder, "parent"), isFolder),
								driveItem(id(file), name(file), parentDir(namex(folder, "parent"), name(folder)), id(folder), isFile)),
							pageItems(delItem(id(folder), parentDir(), idx(folder, "parent"), isFolder))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFoldersProcessed:       4,
					count.TotalDeleteFoldersProcessed: 1,
					count.TotalFilesProcessed:         1,
					count.PagesEnumerated:             2,
				},
				err:            require.NoError,
				numLiveFiles:   0,
				numLiveFolders: 2,
				sizeBytes:      0,
				treeContainsFolderIDs: []string{
					rootID,
					idx(folder, "parent"),
				},
				treeContainsTombstoneIDs: []string{
					id(folder),
				},
				treeContainsFileIDsWithParent: map[string]string{
					id(file): id(folder),
				},
			},
		},
		{
			name: "at folder limit before enumeration",
			tree: treeWithFileAtRoot(),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile)),
							pageItems(
								driveItem(idx(folder, "sib"), namex(folder, "sib"), parentDir(), rootID, isFolder),
								driveItem(idx(file, "sib"), namex(file, "sib"), parentDir(namex(folder, "sib")), idx(folder, "sib"), isFile)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parentDir(), id(folder), isFolder),
								driveItem(idx(file, "chld"), namex(file, "chld"), parentDir(namex(folder, "chld")), idx(folder, "chld"), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.TotalFilesProcessed:         0,
					count.PagesEnumerated:             0,
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
			tree: newFolderyMcFolderFace(nil, rootID),
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile)),
							pageItems(
								driveItem(idx(folder, "sib"), namex(folder, "sib"), parentDir(), rootID, isFolder),
								driveItem(idx(file, "sib"), namex(file, "sib"), parentDir(namex(folder, "sib")), idx(folder, "sib"), isFile)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(folder, "chld"), namex(folder, "chld"), parentDir(), id(folder), isFolder),
								driveItem(idx(file, "chld"), namex(file, "chld"), parentDir(namex(folder, "chld")), idx(folder, "chld"), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFoldersProcessed: 0,
					count.TotalFoldersProcessed:       1,
					count.TotalFilesProcessed:         0,
					count.PagesEnumerated:             0,
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
				drv,
				id(delta),
				test.limiter,
				counter,
				fault.New(true))

			test.expect.err(t, err, clues.ToCore(err))

			assert.Equal(
				t,
				test.expect.numLiveFolders,
				test.tree.countLiveFolders(),
				"count folders in tree")

			countSize := test.tree.countLiveFilesAndSizes()
			assert.Equal(
				t,
				test.expect.numLiveFiles,
				countSize.numFiles,
				"count files in tree")
			assert.Equal(
				t,
				test.expect.sizeBytes,
				countSize.totalBytes,
				"count total bytes in tree")
			test.expect.counts.Compare(t, counter)

			for _, id := range test.expect.treeContainsFolderIDs {
				assert.NotNil(t, test.tree.folderIDToNode[id], "node exists")
			}

			for _, id := range test.expect.treeContainsTombstoneIDs {
				assert.NotNil(t, test.tree.tombstones[id], "tombstone exists")
			}

			for iID, pID := range test.expect.treeContainsFileIDsWithParent {
				assert.Contains(t, test.tree.fileIDToParentID, iID, "file should exist in tree")
				assert.Equal(t, pID, test.tree.fileIDToParentID[iID], "file should reference correct parent")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// folder tests
// ---------------------------------------------------------------------------

// This test focuses on folder assertions when enumerating a page of items.
// File-specific assertions are focused in the _folders test variant.
func (suite *CollectionsTreeUnitSuite) TestCollections_EnumeratePageOfItems_folders() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

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
			page:    pageItems(),
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
			tree: treeWithRoot(),
			page: pageItems(
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
				driveItem(idx(folder, "sib"), namex(folder, "sib"), parentDir(), rootID, isFolder),
				driveItem(idx(folder, "chld"), namex(folder, "chld"), parentDir(name(folder)), id(folder), isFolder)),
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
			name: "create->delete",
			tree: treeWithRoot(),
			page: pageItems(
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
				delItem(id(folder), parentDir(), rootID, isFolder)),
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
			page: pageItems(
				driveItem(idx(folder, "parent"), namex(folder, "parent"), parentDir(), rootID, isFolder),
				driveItem(id(folder), namex(folder, "moved"), parentDir(namex(folder, "parent")), idx(folder, "parent"), isFolder),
				delItem(id(folder), parentDir(), idx(folder, "parent"), isFolder)),
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
			page: pageItems(
				delItem(id(folder), parentDir(), rootID, isFolder),
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder)),
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
			page: pageItems(
				delItem(id(folder), parentDir(), rootID, isFolder),
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder)),
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
				drv,
				test.page,
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
				len(test.tree.tombstones)+test.tree.countLiveFolders(),
				"count folders in tree")
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

	fld := driveItem(id(folder), name(folder), parentDir(), rootID, isFolder)
	subFld := driveItem(id(folder), name(folder), driveParentDir(drv, namex(folder, "parent")), idx(folder, "parent"), isFolder)
	pack := driveItem(id(pkg), name(pkg), parentDir(), rootID, isPackage)
	del := delItem(id(folder), parentDir(), rootID, isFolder)
	mal := malwareItem(idx(folder, "mal"), namex(folder, "mal"), parentDir(), rootID, isFolder)

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
		tree    *folderyMcFolderFace
		folder  models.DriveItemable
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "add folder",
			tree:    treeWithRoot(),
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
			tree:    treeWithFolders(),
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
			tree:    treeWithRoot(),
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
			tree:    treeWithFolders(),
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
			tree:    newFolderyMcFolderFace(nil, rootID),
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
			tree:    treeWithTombstone(),
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
			tree:    treeWithRoot(),
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
			tree:    treeWithFolders(),
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
			tree:    treeWithRoot(),
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
			tree:    treeWithFolders(),
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

			c := collWithMBH(mock.DefaultOneDriveBH(user))
			counter := count.New()

			skipped, err := c.addFolderToTree(
				ctx,
				test.tree,
				drv,
				test.folder,
				test.limiter,
				counter)

			test.expect.err(t, err, clues.ToCore(err))
			test.expect.skipped(t, skipped)

			if test.expect.shouldHitLimit {
				assert.ErrorIs(t, err, errHitLimit, clues.ToCore(err))
			}

			test.expect.counts.Compare(t, counter)
			assert.Equal(t, test.expect.countLiveFolders, test.tree.countLiveFolders(), "live folders")
			assert.Equal(
				t,
				test.expect.treeSize,
				len(test.tree.tombstones)+test.tree.countLiveFolders(),
				"folders in tree")
			test.expect.treeContainsFolder(t, test.tree.containsFolder(ptr.Val(test.folder.GetId())))
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
			folder:    driveRootItem(),
			expect:    basePath.String(),
			expectErr: require.NoError,
		},
		{
			name:      "folder",
			folder:    driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
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

			p, err := c.makeFolderCollectionPath(ctx, id(drive), test.folder)
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
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

	type expected struct {
		counts                        countTD.Expected
		err                           require.ErrorAssertionFunc
		treeContainsFileIDsWithParent map[string]string
		countLiveFiles                int
		countTotalBytes               int64
	}

	table := []struct {
		name   string
		tree   *folderyMcFolderFace
		page   []models.DriveItemable
		expect expected
	}{
		{
			name: "one file at root",
			tree: treeWithRoot(),
			page: pageItems(driveItem(id(file), name(file), parentDir(name(folder)), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "one file in a folder",
			tree: newFolderyMcFolderFace(nil, rootID),
			page: pageItems(
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
				driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     2,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): id(folder),
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "many files in a hierarchy",
			tree: treeWithRoot(),
			page: pageItems(
				driveItem(id(file), name(file), parentDir(), rootID, isFile),
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
				driveItem(idx(file, "chld"), namex(file, "chld"), parentDir(name(folder)), id(folder), isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     2,
					count.TotalFilesProcessed:       2,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file):          rootID,
					idx(file, "chld"): id(folder),
				},
				countLiveFiles:  2,
				countTotalBytes: defaultItemSize * 2,
			},
		},
		{
			name: "many updates to the same file",
			tree: treeWithRoot(),
			page: pageItems(
				driveItem(id(file), name(file), parentDir(), rootID, isFile),
				driveItem(id(file), namex(file, 1), parentDir(), rootID, isFile),
				driveItem(id(file), namex(file, 2), parentDir(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 0,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       3,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "delete an existing file",
			tree: treeWithFileAtRoot(),
			page: pageItems(delItem(id(file), parentDir(), rootID, isFile)),
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
			tree: treeWithFileAtRoot(),
			page: pageItems(
				delItem(id(file), parentDir(), rootID, isFile),
				delItem(id(file), parentDir(), rootID, isFile)),
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
			tree: treeWithRoot(),
			page: pageItems(
				driveItem(id(file), name(file), parentDir(), rootID, isFile),
				delItem(id(file), parentDir(), rootID, isFile)),
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
			tree: treeWithFileAtRoot(),
			page: pageItems(
				driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
				driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile),
				delItem(id(file), parentDir(name(folder)), id(folder), isFile)),
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
			name: "delete->create an existing file",
			tree: treeWithFileAtRoot(),
			page: pageItems(
				delItem(id(file), parentDir(), rootID, isFile),
				driveItem(id(file), name(file), parentDir(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name: "delete->create a non-existing file",
			tree: treeWithRoot(),
			page: pageItems(
				delItem(id(file), parentDir(), rootID, isFile),
				driveItem(id(file), name(file), parentDir(), rootID, isFile)),
			expect: expected{
				counts: countTD.Expected{
					count.TotalDeleteFilesProcessed: 1,
					count.TotalFoldersProcessed:     1,
					count.TotalFilesProcessed:       1,
				},
				err: require.NoError,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
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

			c := collWithMBH(mock.DefaultOneDriveBH(user))
			counter := count.New()

			err := c.enumeratePageOfItems(
				ctx,
				test.tree,
				drv,
				test.page,
				newPagerLimiter(control.DefaultOptions()),
				counter,
				fault.New(true))
			test.expect.err(t, err, clues.ToCore(err))

			countSize := test.tree.countLiveFilesAndSizes()
			assert.Equal(t, test.expect.countLiveFiles, countSize.numFiles, "count of files")
			assert.Equal(t, test.expect.countTotalBytes, countSize.totalBytes, "total size in bytes")
			assert.Equal(t, test.expect.treeContainsFileIDsWithParent, test.tree.fileIDToParentID)
			test.expect.counts.Compare(t, counter)
		})
	}
}

func (suite *CollectionsTreeUnitSuite) TestCollections_AddFileToTree() {
	drv := models.NewDrive()
	drv.SetId(ptr.To(id(drive)))
	drv.SetName(ptr.To(name(drive)))

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
		tree    *folderyMcFolderFace
		file    models.DriveItemable
		limiter *pagerLimiter
		expect  expected
	}{
		{
			name:    "add new file",
			tree:    treeWithRoot(),
			file:    driveItem(id(file), name(file), parentDir(), rootID, isFile),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:     require.NoError,
				skipped: assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "duplicate file",
			tree:    treeWithFileAtRoot(),
			file:    driveItem(id(file), name(file), parentDir(), rootID, isFile),
			limiter: newPagerLimiter(control.DefaultOptions()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:     require.NoError,
				skipped: assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "error file seen before parent",
			tree:    treeWithRoot(),
			file:    driveItem(id(file), name(file), parentDir(name(folder)), id(folder), isFile),
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
			tree:    treeWithRoot(),
			file:    malwareItem(id(file), name(file), parentDir(name(folder)), rootID, isFile),
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
			tree:    treeWithRoot(),
			file:    delItem(id(file), parentDir(name(folder)), id(folder), isFile),
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
			tree:    treeWithFileAtRoot(),
			file:    delItem(id(file), parentDir(), rootID, isFile),
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
			tree:    treeWithFileAtRoot(),
			file:    driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
			limiter: newPagerLimiter(minimumLimitOpts()),
			expect: expected{
				counts: countTD.Expected{
					count.TotalFilesProcessed: 1,
				},
				err:            require.Error,
				shouldHitLimit: true,
				skipped:        assert.Nil,
				treeContainsFileIDsWithParent: map[string]string{
					id(file): rootID,
				},
				countLiveFiles:  1,
				countTotalBytes: defaultItemSize,
			},
		},
		{
			name:    "goes over total byte limit",
			tree:    treeWithRoot(),
			file:    driveItem(id(file), name(file), parentDir(), rootID, isFile),
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

			c := collWithMBH(mock.DefaultOneDriveBH(user))
			counter := count.New()

			skipped, err := c.addFileToTree(
				ctx,
				test.tree,
				drv,
				test.file,
				test.limiter,
				counter)

			test.expect.err(t, err, clues.ToCore(err))
			test.expect.skipped(t, skipped)

			if test.expect.shouldHitLimit {
				require.ErrorIs(t, err, errHitLimit, clues.ToCore(err))
			}

			assert.Equal(t, test.expect.treeContainsFileIDsWithParent, test.tree.fileIDToParentID)
			test.expect.counts.Compare(t, counter)

			countSize := test.tree.countLiveFilesAndSizes()
			assert.Equal(t, test.expect.countLiveFiles, countSize.numFiles, "count of files")
			assert.Equal(t, test.expect.countTotalBytes, countSize.totalBytes, "total size in bytes")
		})
	}
}
