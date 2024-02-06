package exchange

import (
	"context"
	"fmt"
	stdpath "path"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// mocks and helpers
// ---------------------------------------------------------------------------

var _ graph.CachedContainer = &mockContainer{}

type mockContainer struct {
	id          *string
	displayName *string
	parentID    *string
	p           *path.Builder
	l           *path.Builder
}

//revive:disable-next-line:var-naming
func (m mockContainer) GetId() *string { return m.id }

//revive:disable-next-line:var-naming
func (m mockContainer) GetParentFolderId() *string  { return m.parentID }
func (m mockContainer) GetDisplayName() *string     { return m.displayName }
func (m mockContainer) Location() *path.Builder     { return m.l }
func (m mockContainer) SetLocation(p *path.Builder) {}
func (m mockContainer) Path() *path.Builder         { return m.p }
func (m mockContainer) SetPath(p *path.Builder)     {}

func strPtr(s string) *string {
	return &s
}

var _ graph.ContainerResolver = mockContainerResolver{}

type mockContainerResolver struct {
	containersByID map[string]graph.CachedContainer
}

func (mc mockContainerResolver) IDToPath(
	ctx context.Context,
	id string,
) (*path.Builder, *path.Builder, error) {
	return nil, nil, clues.New("not implemented")
}

func (mc mockContainerResolver) Populate(
	ctx context.Context,
	errs *fault.Bus,
	baseFolderID string,
	baseContainerPath ...string,
) error {
	return clues.New("not implemented")
}

func (mc mockContainerResolver) PathInCache(p string) (string, bool) {
	return "", false
}

func (mc mockContainerResolver) LocationInCache(p string) (string, bool) {
	return "", false
}

func (mc mockContainerResolver) AddToCache(
	ctx context.Context,
	c graph.Container,
) error {
	return clues.New("not implemented")
}

func (mc mockContainerResolver) ItemByID(id string) graph.CachedContainer {
	return mc.containersByID[id]
}

func (mc mockContainerResolver) Items() []graph.CachedContainer {
	return maps.Values(mc.containersByID)
}

var _ containerGetter = mockContainerGetter{}

type containerGetterRes struct {
	c   graph.Container
	err error
}

type mockContainerGetter struct {
	itemsByID map[string]containerGetterRes
}

func (mcg mockContainerGetter) GetContainerByID(
	ctx context.Context,
	userID string,
	containerID string,
) (graph.Container, error) {
	res := mcg.itemsByID[containerID]
	return res.c, res.err
}

// ---------------------------------------------------------------------------
// rankedContainerResolver unit tests
// ---------------------------------------------------------------------------

type RankedContainerResolverUnitSuite struct {
	tester.Suite
}

func TestRankedContainerResolverUnitSuite(t *testing.T) {
	suite.Run(t, &RankedContainerResolverUnitSuite{
		Suite: tester.NewUnitSuite(t),
	})
}

func (suite *RankedContainerResolverUnitSuite) TestItemByID() {
	// Containers available to operate on directly in tests.
	const (
		id1         = "id1"
		idNotInBase = "idNotInBase"
	)

	mcg := mockContainerGetter{
		itemsByID: map[string]containerGetterRes{
			id1:         {c: mockContainer{id: ptr.To(id1)}},
			idNotInBase: {c: mockContainer{id: ptr.To(idNotInBase)}},
		},
	}

	// Configure base containers we want.
	mcr := &mockContainerResolver{
		containersByID: map[string]graph.CachedContainer{
			id1: &mockCachedContainer{id: id1},
		},
	}

	table := []struct {
		name        string
		includes    []string
		excludes    []string
		itemID      string
		expectFound bool
	}{
		{
			name:        "NeitherIncludedNorExcluded",
			itemID:      id1,
			expectFound: true,
		},
		{
			name:        "NeitherIncludedOrExcluded NotInBase",
			itemID:      idNotInBase,
			expectFound: false,
		},
		{
			name:        "Included",
			includes:    []string{id1},
			itemID:      id1,
			expectFound: true,
		},
		{
			name:        "Excluded",
			excludes:    []string{id1},
			itemID:      id1,
			expectFound: false,
		},
		{
			name:        "IncludedAndExcluded",
			includes:    []string{id1},
			excludes:    []string{id1},
			itemID:      id1,
			expectFound: true,
		},
		{
			name:        "IncludedAndExcluded NotInBaseResolver",
			includes:    []string{idNotInBase},
			excludes:    []string{idNotInBase},
			itemID:      idNotInBase,
			expectFound: false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			rcr, err := newRankedContainerResolver(
				ctx,
				mcr,
				mcg,
				"userID",
				test.includes,
				test.excludes)
			require.NoError(t, err, clues.ToCore(err))

			item := rcr.ItemByID(test.itemID)
			if test.expectFound {
				require.NotNil(t, item)
				assert.Equal(t, test.itemID, ptr.Val(item.GetId()), "returned item ID")
			} else {
				assert.Nil(t, item, "unexpected item returned: %+v", item)
			}
		})
	}
}

func (suite *RankedContainerResolverUnitSuite) TestItems() {
	// Containers available to operate on directly in tests.
	const (
		id1         = "id1"
		id2         = "id2"
		id3         = "id3"
		idErr       = "idErr"
		idEmpty     = "idEmpty"
		idNotInBase = "idNotInBase"
	)

	mcg := mockContainerGetter{
		itemsByID: map[string]containerGetterRes{
			id1:         {c: mockContainer{id: ptr.To(id1)}},
			id2:         {c: mockContainer{id: ptr.To(id2)}},
			id3:         {c: mockContainer{id: ptr.To(id3)}},
			idErr:       {err: assert.AnError},
			idEmpty:     {c: mockContainer{}},
			idNotInBase: {c: mockContainer{id: ptr.To(idNotInBase)}},
		},
	}

	// Configure base containers we want.
	mcr := &mockContainerResolver{
		containersByID: map[string]graph.CachedContainer{
			id1: &mockCachedContainer{id: id1},
			id2: &mockCachedContainer{id: id2},
			id3: &mockCachedContainer{id: id3},
		},
	}

	// Add a bunch more containers so we're more likely to get a random order from
	// the map.
	var otherContainerIDs []string

	for i := 0; i < 100; i++ {
		id := fmt.Sprintf("extra-id%d", i)
		mcr.containersByID[id] = &mockCachedContainer{id: id}
		otherContainerIDs = append(otherContainerIDs, id)
	}

	table := []struct {
		name     string
		includes []string
		excludes []string
		// expectPrefix is the prefix of container IDs that should be in the result.
		expectPrefix []string
		// expectExtraUnordered allows specifying additional IDs in the set of IDs
		// available to work with that will appear in the unordered set. For
		// example, if only id1 was part of includes and excludes was empty then id2
		// and id3 should appear somewhere in the output but don't have a particluar
		// order requirement.
		expectExtraUnordered []string
		expectErr            assert.ErrorAssertionFunc
	}{
		{
			name:         "ReturnsRankedItems",
			includes:     []string{id2, id1, id3},
			expectPrefix: []string{id2, id1, id3},
			expectErr:    assert.NoError,
		},
		{
			name:                 "ReturnsRankedItems SomeUnordered",
			includes:             []string{id2, id1},
			expectPrefix:         []string{id2, id1},
			expectExtraUnordered: []string{id3},
			expectErr:            assert.NoError,
		},
		{
			name:         "ReturnsRankedItems SomeExcluded",
			includes:     []string{id2},
			excludes:     []string{id3, id1},
			expectPrefix: []string{id2},
			expectErr:    assert.NoError,
		},
		{
			name:                 "ReturnsRankedItems SomeIncludesNotInBase",
			includes:             []string{id2, id1, idNotInBase},
			expectPrefix:         []string{id2, id1},
			expectExtraUnordered: []string{id3},
			expectErr:            assert.NoError,
		},
		{
			name:                 "ReturnsRankedItems IncludedAndExcluded",
			includes:             []string{id2},
			excludes:             []string{id2},
			expectPrefix:         []string{id2},
			expectExtraUnordered: []string{id1, id3},
			expectErr:            assert.NoError,
		},
		{
			name:                 "ReturnsRankedItems IncludedAndExcluded NotInBase",
			includes:             []string{idNotInBase},
			excludes:             []string{idNotInBase},
			expectPrefix:         []string{},
			expectExtraUnordered: []string{id1, id2, id3},
			expectErr:            assert.NoError,
		},
		{
			name:         "ReturnsRankedItems SomeExcludesNotInBase",
			includes:     []string{id2},
			excludes:     []string{id1, idNotInBase, id3},
			expectPrefix: []string{id2},
			expectErr:    assert.NoError,
		},
		{
			name:      "FailsOnIncludeError",
			includes:  []string{id2, id1, idErr},
			expectErr: assert.Error,
		},
		{
			name:      "FailsOnExcludeError",
			excludes:  []string{id2, id1, idErr},
			expectErr: assert.Error,
		},
		{
			name:      "FailsOnIncludeContainerWithEmptyID",
			includes:  []string{id2, id1, idEmpty},
			expectErr: assert.Error,
		},
		{
			name:      "FailsOnExcludeContainerWithEmptyID",
			excludes:  []string{id2, id1, idEmpty},
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			rcr, err := newRankedContainerResolver(
				ctx,
				mcr,
				mcg,
				"userID",
				test.includes,
				test.excludes)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			items := rcr.Items()
			resIDs := make([]string, 0, len(items))

			for _, item := range items {
				resIDs = append(resIDs, ptr.Val(item.GetId()))
			}

			assert.Equal(
				t,
				test.expectPrefix,
				resIDs[:len(test.expectPrefix)],
				"ordered prefix of result")
			assert.ElementsMatch(
				t,
				append(slices.Clone(test.expectExtraUnordered), otherContainerIDs...),
				resIDs[len(test.expectPrefix):],
				"unordered remainder of result")
		})
	}
}

// ---------------------------------------------------------------------------
// unit suite
// ---------------------------------------------------------------------------

type FolderCacheUnitSuite struct {
	tester.Suite
}

func TestFolderCacheUnitSuite(t *testing.T) {
	suite.Run(t, &FolderCacheUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type containerCheckTestInfo struct {
	name  string
	c     mockContainer
	check assert.ErrorAssertionFunc
}

var (
	testID       = uuid.NewString()
	testName     = "foo"
	testParentID = uuid.NewString()
	emptyString  = ""

	containerCheckTests = []containerCheckTestInfo{
		{
			name: "NilID",
			c: mockContainer{
				id:          nil,
				displayName: &testName,
				parentID:    &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "NilDisplayName",
			c: mockContainer{
				id:          &testID,
				displayName: nil,
				parentID:    &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "EmptyID",
			c: mockContainer{
				id:          &emptyString,
				displayName: &testName,
				parentID:    &testParentID,
			},
			check: assert.NoError,
		},
		{
			name: "EmptyDisplayName",
			c: mockContainer{
				id:          &testID,
				displayName: &emptyString,
				parentID:    &testParentID,
			},
			check: assert.NoError,
		},
		{
			name: "AllValues",
			c: mockContainer{
				id:          &testID,
				displayName: &testName,
				parentID:    &testParentID,
			},
			check: assert.NoError,
		},
	}
)

func (suite *FolderCacheUnitSuite) TestCheckIDAndName() {
	for _, test := range containerCheckTests {
		suite.Run(test.name, func() {
			err := checkIDAndName(test.c)
			test.check(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *FolderCacheUnitSuite) TestCheckRequiredValues() {
	table := []containerCheckTestInfo{
		{
			name: "NilParentFolderID",
			c: mockContainer{
				id:          &testID,
				displayName: &testName,
				parentID:    nil,
			},
			check: assert.Error,
		},
		{
			name: "EmptyParentFolderID",
			c: mockContainer{
				id:          &testID,
				displayName: &testName,
				parentID:    &emptyString,
			},
			check: assert.NoError,
		},
	}

	table = append(table, containerCheckTests...)

	for _, test := range table {
		suite.Run(test.name, func() {
			err := checkRequiredValues(test.c)
			test.check(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *FolderCacheUnitSuite) TestAddFolder() {
	table := []struct {
		name  string
		cf    graph.CacheFolder
		check assert.ErrorAssertionFunc
	}{
		{
			name: "NoParentNoPath",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:          &testID,
					displayName: &testName,
					parentID:    nil,
				},
				nil,
				nil),
			check: assert.Error,
		},
		{
			name: "NoParentPath",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:          &testID,
					displayName: &testName,
					parentID:    nil,
				},
				path.Builder{}.Append("foo"),
				path.Builder{}.Append("loc")),
			check: assert.NoError,
		},
		{
			name: "NoName",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:          &testID,
					displayName: nil,
					parentID:    &testParentID,
				},
				path.Builder{}.Append("foo"),
				path.Builder{}.Append("loc")),
			check: assert.Error,
		},
		{
			name: "NoID",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:          nil,
					displayName: &testName,
					parentID:    &testParentID,
				},
				path.Builder{}.Append("foo"),
				path.Builder{}.Append("loc")),
			check: assert.Error,
		},
		{
			name: "NoPath",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:          &testID,
					displayName: &testName,
					parentID:    &testParentID,
				},
				nil,
				nil),
			check: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			fc := newContainerResolver(nil)
			err := fc.addFolder(&test.cf)
			test.check(suite.T(), err, clues.ToCore(err))
		})
	}
}

func newMockCachedContainer(name string) *mockCachedContainer {
	return &mockCachedContainer{
		id:          uuid.NewString(),
		parentID:    uuid.NewString(),
		displayName: name,
	}
}

type mockCachedContainer struct {
	id               string
	parentID         string
	displayName      string
	l                *path.Builder
	p                *path.Builder
	expectedPath     string
	expectedLocation string
}

//nolint:revive
func (m mockCachedContainer) GetId() *string { return &m.id }

//nolint:revive
func (m mockCachedContainer) GetParentFolderId() *string        { return &m.parentID }
func (m mockCachedContainer) GetDisplayName() *string           { return &m.displayName }
func (m mockCachedContainer) Location() *path.Builder           { return m.l }
func (m *mockCachedContainer) SetLocation(newLoc *path.Builder) { m.l = newLoc }
func (m mockCachedContainer) Path() *path.Builder               { return m.p }
func (m *mockCachedContainer) SetPath(newPath *path.Builder)    { m.p = newPath }

func resolverWithContainers(numContainers int, useIDInPath bool) (*containerResolver, []*mockCachedContainer) {
	containers := make([]*mockCachedContainer, 0, numContainers)

	for i := 0; i < numContainers; i++ {
		containers = append(containers, newMockCachedContainer(fmt.Sprintf("%d", i)))
	}

	// Base case for the recursive lookup.
	dn := containers[0].displayName
	apndP := containers[0].id

	containers[0].p = path.Builder{}.Append(apndP)
	containers[0].expectedPath = apndP
	containers[0].l = path.Builder{}.Append(dn)
	containers[0].expectedLocation = dn

	for i := 1; i < len(containers); i++ {
		dn := containers[i].displayName
		apndP := containers[i].id

		containers[i].parentID = containers[i-1].id
		containers[i].expectedPath = stdpath.Join(containers[i-1].expectedPath, apndP)
		containers[i].expectedLocation = stdpath.Join(containers[i-1].expectedLocation, dn)
	}

	resolver := newContainerResolver(nil)

	for _, c := range containers {
		resolver.cache[c.id] = c
	}

	return resolver, containers
}

// ---------------------------------------------------------------------------
// mock container refresher
// ---------------------------------------------------------------------------

type refreshResult struct {
	err error
	c   graph.CachedContainer
}

type mockContainerRefresher struct {
	// Folder ID -> result
	entries map[string]refreshResult
}

func (r mockContainerRefresher) refreshContainer(
	ctx context.Context,
	id string,
) (graph.CachedContainer, error) {
	rr, ok := r.entries[id]
	if !ok {
		// May not be this precise error, but it's easy to get a handle on.
		return nil, core.ErrNotFound
	}

	if rr.err != nil {
		return nil, rr.err
	}

	return rr.c, nil
}

// ---------------------------------------------------------------------------
// configured unit suite
// ---------------------------------------------------------------------------

// TestConfiguredFolderCacheUnitSuite cannot run its tests in parallel.
type ConfiguredFolderCacheUnitSuite struct {
	tester.Suite

	fc *containerResolver

	allContainers []*mockCachedContainer
}

func (suite *ConfiguredFolderCacheUnitSuite) SetupTest() {
	suite.fc, suite.allContainers = resolverWithContainers(4, false)
}

func TestConfiguredFolderCacheUnitSuite(t *testing.T) {
	suite.Run(t, &ConfiguredFolderCacheUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConfiguredFolderCacheUnitSuite) TestRefreshContainer_RefreshParent() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	resolver, containers := resolverWithContainers(4, true)
	almostLast := containers[len(containers)-2]
	last := containers[len(containers)-1]

	refresher := mockContainerRefresher{
		entries: map[string]refreshResult{
			almostLast.id: {c: almostLast},
			last.id:       {c: last},
		},
	}

	resolver.refresher = refresher

	delete(resolver.cache, almostLast.id)

	ferrs := fault.New(true)
	err := resolver.populatePaths(ctx, ferrs)
	require.NoError(t, err, "populating paths", clues.ToCore(err))

	p, l, err := resolver.IDToPath(ctx, last.id)
	require.NoError(t, err, "getting paths", clues.ToCore(err))

	assert.Equal(t, last.expectedPath, p.String())
	assert.Equal(t, last.expectedLocation, l.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestRefreshContainer_RefreshParent_NotFoundDeletes() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	resolver, containers := resolverWithContainers(4, true)
	almostLast := containers[len(containers)-2]
	last := containers[len(containers)-1]

	refresher := mockContainerRefresher{
		entries: map[string]refreshResult{
			last.id: {c: last},
		},
	}

	resolver.refresher = refresher

	delete(resolver.cache, almostLast.id)

	ferrs := fault.New(true)
	err := resolver.populatePaths(ctx, ferrs)
	require.NoError(t, err, "populating paths", clues.ToCore(err))

	_, _, err = resolver.IDToPath(ctx, last.id)
	assert.Error(t, err, "getting paths", clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestRefreshContainer_RefreshAncestor_NotFoundDeletes() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	resolver, containers := resolverWithContainers(4, true)
	gone := containers[0]
	child := containers[1]
	last := containers[len(containers)-1]

	refresher := mockContainerRefresher{
		entries: map[string]refreshResult{
			child.id: {c: child},
		},
	}

	resolver.refresher = refresher

	delete(resolver.cache, gone.id)

	ferrs := fault.New(true)
	err := resolver.populatePaths(ctx, ferrs)
	require.NoError(t, err, "populating paths", clues.ToCore(err))

	_, _, err = resolver.IDToPath(ctx, last.id)
	assert.Error(t, err, "getting paths", clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestRefreshContainer_RefreshAncestor_NewParent() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	resolver, containers := resolverWithContainers(4, true)
	other := containers[len(containers)-3]
	gone := containers[len(containers)-2]
	last := containers[len(containers)-1]

	expected := *last
	expected.parentID = other.id
	expected.expectedPath = stdpath.Join(other.expectedPath, expected.id)
	expected.expectedLocation = stdpath.Join(other.expectedLocation, expected.displayName)

	refresher := mockContainerRefresher{
		entries: map[string]refreshResult{
			last.id: {c: &expected},
		},
	}

	resolver.refresher = refresher

	delete(resolver.cache, gone.id)

	ferrs := fault.New(true)
	err := resolver.populatePaths(ctx, ferrs)
	require.NoError(t, err, "populating paths", clues.ToCore(err))

	p, l, err := resolver.IDToPath(ctx, last.id)
	require.NoError(t, err, "getting paths", clues.ToCore(err))

	assert.Equal(t, expected.expectedPath, p.String())
	assert.Equal(t, expected.expectedLocation, l.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestRefreshContainer_RefreshFolder_FolderDeleted() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	resolver, containers := resolverWithContainers(4, true)
	parent := containers[len(containers)-2]
	last := containers[len(containers)-1]

	refresher := mockContainerRefresher{
		entries: map[string]refreshResult{
			parent.id: {c: parent},
		},
	}

	resolver.refresher = refresher

	delete(resolver.cache, parent.id)

	ferrs := fault.New(true)
	err := resolver.populatePaths(ctx, ferrs)
	require.NoError(t, err, "populating paths", clues.ToCore(err))

	_, _, err = resolver.IDToPath(ctx, last.id)
	assert.Error(t, err, "getting paths", clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestDepthLimit() {
	table := []struct {
		name          string
		numContainers int
		check         assert.ErrorAssertionFunc
	}{
		{
			name:          "AtLimit",
			numContainers: maxIterations,
			check:         assert.NoError,
		},
		{
			name:          "OverLimit",
			numContainers: maxIterations + 1,
			check:         assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			resolver, containers := resolverWithContainers(test.numContainers, false)
			_, err := resolver.idToPath(ctx, containers[len(containers)-1].id, 0)
			test.check(t, err, clues.ToCore(err))
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestPopulatePaths() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	err := suite.fc.populatePaths(ctx, fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	items := suite.fc.Items()
	gotPaths := make([]string, 0, len(items))

	for _, i := range items {
		gotPaths = append(gotPaths, i.Path().String())
	}

	expectedPaths := make([]string, 0, len(suite.allContainers))
	for _, c := range suite.allContainers {
		expectedPaths = append(expectedPaths, c.expectedPath)
	}

	assert.ElementsMatch(t, expectedPaths, gotPaths)
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderNoPathsCached() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	err := suite.fc.populatePaths(ctx, fault.New(true))
	require.NoError(suite.T(), err, clues.ToCore(err))

	for _, c := range suite.allContainers {
		suite.Run(ptr.Val(c.GetDisplayName()), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			p, l, err := suite.fc.IDToPath(ctx, c.id)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, c.expectedPath, p.String())
			assert.Equal(t, c.expectedLocation, l.String())
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderCachesPaths() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	c := suite.allContainers[len(suite.allContainers)-1]

	err := suite.fc.populatePaths(ctx, fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	p, l, err := suite.fc.IDToPath(ctx, c.id)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())

	c.parentID = "foo"

	p, l, err = suite.fc.IDToPath(ctx, c.id)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsParentNotFound() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	almostLast := suite.allContainers[len(suite.allContainers)-2]

	delete(suite.fc.cache, almostLast.id)

	err := suite.fc.populatePaths(ctx, fault.New(true))
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolder_Errors_PathsNotBuilt() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	_, _, err := suite.fc.IDToPath(ctx, suite.allContainers[len(suite.allContainers)-1].id)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsNotFound() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	_, _, err := suite.fc.IDToPath(ctx, "foo")
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestAddToCache() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		dest = "testAddFolder"
		last = suite.allContainers[len(suite.allContainers)-1]
		m    = newMockCachedContainer(dest)
	)

	m.parentID = last.id
	m.expectedPath = stdpath.Join(last.expectedPath, m.id)
	m.expectedLocation = stdpath.Join(last.expectedLocation, m.displayName)

	err := suite.fc.AddToCache(ctx, m)
	require.NoError(t, err, clues.ToCore(err))

	p, l, err := suite.fc.IDToPath(ctx, m.id)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, m.expectedPath, p.String(), "ID path")
	assert.Equal(t, m.expectedLocation, l.String(), "location path")
}

type ContainerResolverIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestContainerResolverIntgSuite(t *testing.T) {
	suite.Run(t, &ContainerResolverIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ContainerResolverIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *ContainerResolverIntgSuite) TestPopulate() {
	eventFunc := func(t *testing.T) graph.ContainerResolver {
		return &eventContainerCache{
			userID: tconfig.M365UserID(t),
			enumer: suite.m365.AC.Events(),
			getter: suite.m365.AC.Events(),
		}
	}

	contactFunc := func(t *testing.T) graph.ContainerResolver {
		return &contactContainerCache{
			userID: tconfig.M365UserID(t),
			enumer: suite.m365.AC.Contacts(),
			getter: suite.m365.AC.Contacts(),
		}
	}

	tests := []struct {
		name, folderInCache, root, basePath string
		resolverFunc                        func(t *testing.T) graph.ContainerResolver
		canFind                             assert.BoolAssertionFunc
	}{
		{
			name: "Default Event Cache",
			// Fine as long as this isn't running against a migrated Exchange server.
			folderInCache: api.DefaultCalendar,
			root:          api.DefaultCalendar,
			basePath:      api.DefaultCalendar,
			resolverFunc:  eventFunc,
			canFind:       assert.True,
		},
		{
			name:          "Default Event Folder Hidden",
			folderInCache: api.DefaultContacts,
			root:          api.DefaultCalendar,
			canFind:       assert.False,
			resolverFunc:  eventFunc,
		},
		{
			name:          "Name Not in Cache",
			folderInCache: "testFooBarWhoBar",
			root:          api.DefaultCalendar,
			canFind:       assert.False,
			resolverFunc:  eventFunc,
		},
		{
			name:          "Default Contact Cache",
			folderInCache: api.DefaultContacts,
			root:          api.DefaultContacts,
			basePath:      api.DefaultContacts,
			canFind:       assert.True,
			resolverFunc:  contactFunc,
		},
		{
			name:          "Default Contact Hidden",
			folderInCache: api.DefaultContacts,
			root:          api.DefaultContacts,
			canFind:       assert.False,
			resolverFunc:  contactFunc,
		},
		{
			name:          "Name Not in Cache",
			folderInCache: "testFooBarWhoBar",
			root:          api.DefaultContacts,
			canFind:       assert.False,
			resolverFunc:  contactFunc,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			resolver := test.resolverFunc(t)

			err := resolver.Populate(ctx, fault.New(true), test.root, test.basePath)
			require.NoError(t, err, clues.ToCore(err))

			_, isFound := resolver.LocationInCache(test.folderInCache)
			test.canFind(t, isFound, "folder path", test.folderInCache)
		})
	}
}

// ---------------------------------------------------------------------------
// integration suite
// ---------------------------------------------------------------------------

func runCreateDestinationTest(
	t *testing.T,
	handler restoreHandler,
	category path.CategoryType,
	tenantID, userID, destinationName string,
	containerNames1 []string,
	containerNames2 []string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		svc = path.ExchangeService
		gcc = handler.NewContainerCache(userID)
	)

	err := gcc.Populate(ctx, fault.New(true), handler.DefaultRootContainer())
	require.NoError(t, err, clues.ToCore(err))

	path1, err := path.Build(
		tenantID,
		userID,
		svc,
		category,
		false,
		containerNames1...)
	require.NoError(t, err, clues.ToCore(err))

	containerID, gcc, err := CreateDestination(
		ctx,
		handler,
		handler.FormatRestoreDestination(destinationName, path1),
		userID,
		gcc,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	_, _, err = gcc.IDToPath(ctx, containerID)
	assert.NoError(t, err, clues.ToCore(err))

	path2, err := path.Build(
		tenantID,
		userID,
		svc,
		category,
		false,
		containerNames2...)
	require.NoError(t, err, clues.ToCore(err))

	containerID, gcc, err = CreateDestination(
		ctx,
		handler,
		handler.FormatRestoreDestination(destinationName, path2),
		userID,
		gcc,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	p, l, err := gcc.IDToPath(ctx, containerID)
	require.NoError(t, err, clues.ToCore(err))

	_, ok := gcc.LocationInCache(l.String())
	require.True(t, ok, "looking for location in cache: %s", l)

	_, ok = gcc.PathInCache(p.String())
	require.True(t, ok, "looking for path in cache: %s", p)
}
