package exchange

import (
	stdpath "path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type mockContainer struct {
	id       *string
	name     *string
	parentID *string
}

//nolint:revive
func (m mockContainer) GetId() *string {
	return m.id
}

func (m mockContainer) GetDisplayName() *string {
	return m.name
}

//nolint:revive
func (m mockContainer) GetParentFolderId() *string {
	return m.parentID
}

func (m mockContainer) GetAdditionalData() map[string]any {
	return nil
}

type FolderCacheUnitSuite struct {
	suite.Suite
}

func TestFolderCacheUnitSuite(t *testing.T) {
	suite.Run(t, new(FolderCacheUnitSuite))
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
				id:       nil,
				name:     &testName,
				parentID: &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "NilDisplayName",
			c: mockContainer{
				id:       &testID,
				name:     nil,
				parentID: &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "EmptyID",
			c: mockContainer{
				id:       &emptyString,
				name:     &testName,
				parentID: &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "EmptyDisplayName",
			c: mockContainer{
				id:       &testID,
				name:     &emptyString,
				parentID: &testParentID,
			},
			check: assert.Error,
		},
		{
			name: "AllValues",
			c: mockContainer{
				id:       &testID,
				name:     &testName,
				parentID: &testParentID,
			},
			check: assert.NoError,
		},
	}
)

func (suite *FolderCacheUnitSuite) TestCheckIDAndName() {
	for _, test := range containerCheckTests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.check(t, checkIDAndName(test.c))
		})
	}
}

func (suite *FolderCacheUnitSuite) TestCheckRequiredValues() {
	table := []containerCheckTestInfo{
		{
			name: "NilParentFolderID",
			c: mockContainer{
				id:       &testID,
				name:     &testName,
				parentID: nil,
			},
			check: assert.Error,
		},
		{
			name: "EmptyParentFolderID",
			c: mockContainer{
				id:       &testID,
				name:     &testName,
				parentID: &emptyString,
			},
			check: assert.Error,
		},
	}

	table = append(table, containerCheckTests...)

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.check(t, checkRequiredValues(test.c))
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
					id:       &testID,
					name:     &testName,
					parentID: nil,
				},
				nil,
			),
			check: assert.Error,
		},
		{
			name: "NoParentPath",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:       &testID,
					name:     &testName,
					parentID: nil,
				},
				path.Builder{}.Append("foo"),
			),
			check: assert.NoError,
		},
		{
			name: "NoName",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:       &testID,
					name:     nil,
					parentID: &testParentID,
				},
				path.Builder{}.Append("foo"),
			),
			check: assert.Error,
		},
		{
			name: "NoID",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:       nil,
					name:     &testName,
					parentID: &testParentID,
				},
				path.Builder{}.Append("foo"),
			),
			check: assert.Error,
		},
		{
			name: "NoPath",
			cf: graph.NewCacheFolder(
				&mockContainer{
					id:       &testID,
					name:     &testName,
					parentID: &testParentID,
				},
				nil,
			),
			check: assert.NoError,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			fc := newContainerResolver()
			test.check(t, fc.addFolder(test.cf))
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
	id           string
	parentID     string
	displayName  string
	p            *path.Builder
	expectedPath string
	removed      bool
}

//nolint:revive
func (m mockCachedContainer) GetId() *string {
	return &m.id
}

//nolint:revive
func (m mockCachedContainer) GetParentFolderId() *string {
	return &m.parentID
}

func (m mockCachedContainer) GetDisplayName() *string {
	return &m.displayName
}

func (m mockCachedContainer) Path() *path.Builder {
	return m.p
}

func (m *mockCachedContainer) SetPath(newPath *path.Builder) {
	m.p = newPath
}

func (m *mockCachedContainer) GetAdditionalData() map[string]any {
	return nil
}

func (m *mockCachedContainer) Deleted() bool {
	return m.removed
}

func resolverWithContainers(numContainers int) (*containerResolver, []*mockCachedContainer) {
	containers := make([]*mockCachedContainer, 0, numContainers)

	for i := 0; i < numContainers; i++ {
		containers = append(containers, newMockCachedContainer("a"))
	}

	// Base case for the recursive lookup.
	containers[0].p = path.Builder{}.Append(containers[0].displayName)
	containers[0].expectedPath = containers[0].displayName

	for i := 1; i < len(containers); i++ {
		containers[i].parentID = containers[i-1].id
		containers[i].expectedPath = stdpath.Join(
			containers[i-1].expectedPath,
			containers[i].displayName,
		)
	}

	resolver := newContainerResolver()

	for _, c := range containers {
		resolver.cache[c.id] = c
	}

	return resolver, containers
}

// TestConfiguredFolderCacheUnitSuite cannot run its tests in parallel.
type ConfiguredFolderCacheUnitSuite struct {
	suite.Suite

	fc *containerResolver

	allContainers []*mockCachedContainer
}

func (suite *ConfiguredFolderCacheUnitSuite) SetupTest() {
	suite.fc, suite.allContainers = resolverWithContainers(4)
}

func TestConfiguredFolderCacheUnitSuite(t *testing.T) {
	suite.Run(t, new(ConfiguredFolderCacheUnitSuite))
}

func (suite *ConfiguredFolderCacheUnitSuite) TestDepthLimit() {
	ctx, flush := tester.NewContext()
	defer flush()

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
		suite.T().Run(test.name, func(t *testing.T) {
			resolver, containers := resolverWithContainers(test.numContainers)
			_, err := resolver.IDToPath(ctx, containers[len(containers)-1].id)
			test.check(t, err)
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestPopulatePaths() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	require.NoError(t, suite.fc.populatePaths(ctx))

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
	ctx, flush := tester.NewContext()
	defer flush()

	for _, c := range suite.allContainers {
		suite.T().Run(*c.GetDisplayName(), func(t *testing.T) {
			p, err := suite.fc.IDToPath(ctx, c.id)
			require.NoError(t, err)

			assert.Equal(t, c.expectedPath, p.String())
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderCachesPaths() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	c := suite.allContainers[len(suite.allContainers)-1]

	p, err := suite.fc.IDToPath(ctx, c.id)
	require.NoError(t, err)

	assert.Equal(t, c.expectedPath, p.String())

	c.parentID = "foo"

	p, err = suite.fc.IDToPath(ctx, c.id)
	require.NoError(t, err)

	assert.Equal(t, c.expectedPath, p.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsParentNotFound() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	last := suite.allContainers[len(suite.allContainers)-1]
	almostLast := suite.allContainers[len(suite.allContainers)-2]

	delete(suite.fc.cache, almostLast.id)

	_, err := suite.fc.IDToPath(ctx, last.id)
	assert.Error(t, err)
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsNotFound() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	_, err := suite.fc.IDToPath(ctx, "foo")
	assert.Error(t, err)
}

func (suite *ConfiguredFolderCacheUnitSuite) TestAddToCache() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	last := suite.allContainers[len(suite.allContainers)-1]

	m := newMockCachedContainer("testAddFolder")

	m.parentID = last.id
	m.expectedPath = stdpath.Join(last.expectedPath, m.displayName)

	require.NoError(t, suite.fc.AddToCache(ctx, m))

	p, err := suite.fc.IDToPath(ctx, m.id)
	require.NoError(t, err)
	assert.Equal(t, m.expectedPath, p.String())
}
