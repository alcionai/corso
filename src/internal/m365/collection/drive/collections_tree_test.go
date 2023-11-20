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

	var p path.Path

	switch coll.State() {
	case data.DeletedState:
		p = coll.PreviousPath()
	default:
		p = coll.FullPath()
	}

	for {
		itm, ok := <-itemCh
		if !ok {
			break
		}

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

// TODO(keepers): implement tree version of populateDriveCollections tests

// TODO(keepers): implement tree version of TestGet single-drive tests

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

// TODO(keepers): implement tree version of TestGet multi-drive tests

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
