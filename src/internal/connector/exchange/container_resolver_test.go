package exchange

import (
	"fmt"
	stdpath "path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

//nolint:revive
func (m mockContainer) GetId() *string { return m.id }

//nolint:revive
func (m mockContainer) GetParentFolderId() *string  { return m.parentID }
func (m mockContainer) GetDisplayName() *string     { return m.displayName }
func (m mockContainer) Location() *path.Builder     { return m.l }
func (m mockContainer) SetLocation(p *path.Builder) {}
func (m mockContainer) Path() *path.Builder         { return m.p }
func (m mockContainer) SetPath(p *path.Builder)     {}

func strPtr(s string) *string {
	return &s
}

// ---------------------------------------------------------------------------
// unit suite
// ---------------------------------------------------------------------------

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

	apndP := dn
	if useIDInPath {
		apndP = containers[0].id
	}

	containers[0].p = path.Builder{}.Append(apndP)
	containers[0].expectedPath = apndP
	containers[0].l = path.Builder{}.Append(dn)
	containers[0].expectedLocation = dn

	for i := 1; i < len(containers); i++ {
		dn := containers[i].displayName

		apndP := dn
		if useIDInPath {
			apndP = containers[i].id
		}

		containers[i].parentID = containers[i-1].id
		containers[i].expectedPath = stdpath.Join(containers[i-1].expectedPath, apndP)
		containers[i].expectedLocation = stdpath.Join(containers[i-1].expectedLocation, dn)
	}

	resolver := newContainerResolver()

	for _, c := range containers {
		resolver.cache[c.id] = c
	}

	return resolver, containers
}

// ---------------------------------------------------------------------------
// configured unit suite
// ---------------------------------------------------------------------------

// TestConfiguredFolderCacheUnitSuite cannot run its tests in parallel.
type ConfiguredFolderCacheUnitSuite struct {
	suite.Suite

	fc       *containerResolver
	fcWithID *containerResolver

	allContainers    []*mockCachedContainer
	containersWithID []*mockCachedContainer
}

func (suite *ConfiguredFolderCacheUnitSuite) SetupTest() {
	suite.fc, suite.allContainers = resolverWithContainers(4, false)
	suite.fcWithID, suite.containersWithID = resolverWithContainers(4, true)
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
			resolver, containers := resolverWithContainers(test.numContainers, false)
			_, _, err := resolver.IDToPath(ctx, containers[len(containers)-1].id, false)
			test.check(t, err)
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestPopulatePaths() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	require.NoError(t, suite.fc.populatePaths(ctx, false))

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
			p, l, err := suite.fc.IDToPath(ctx, c.id, false)
			require.NoError(t, err)
			assert.Equal(t, c.expectedPath, p.String())
			assert.Equal(t, c.expectedLocation, l.String())
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderNoPathsCached_useID() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, c := range suite.containersWithID {
		suite.T().Run(*c.GetDisplayName(), func(t *testing.T) {
			p, l, err := suite.fcWithID.IDToPath(ctx, c.id, true)
			require.NoError(t, err)
			assert.Equal(t, c.expectedPath, p.String())
			assert.Equal(t, c.expectedLocation, l.String())
		})
	}
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderCachesPaths() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	c := suite.allContainers[len(suite.allContainers)-1]

	p, l, err := suite.fc.IDToPath(ctx, c.id, false)
	require.NoError(t, err)
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())

	c.parentID = "foo"

	p, l, err = suite.fc.IDToPath(ctx, c.id, false)
	require.NoError(t, err)
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderCachesPaths_useID() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	c := suite.containersWithID[len(suite.containersWithID)-1]

	p, l, err := suite.fcWithID.IDToPath(ctx, c.id, true)
	require.NoError(t, err)
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())

	c.parentID = "foo"

	p, l, err = suite.fcWithID.IDToPath(ctx, c.id, true)
	require.NoError(t, err)
	assert.Equal(t, c.expectedPath, p.String())
	assert.Equal(t, c.expectedLocation, l.String())
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsParentNotFound() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	last := suite.allContainers[len(suite.allContainers)-1]
	almostLast := suite.allContainers[len(suite.allContainers)-2]

	delete(suite.fc.cache, almostLast.id)

	_, _, err := suite.fc.IDToPath(ctx, last.id, false)
	assert.Error(t, err)
}

func (suite *ConfiguredFolderCacheUnitSuite) TestLookupCachedFolderErrorsNotFound() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	_, _, err := suite.fc.IDToPath(ctx, "foo", false)
	assert.Error(t, err)
}

func (suite *ConfiguredFolderCacheUnitSuite) TestAddToCache() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		dest = "testAddFolder"
		t    = suite.T()
		last = suite.allContainers[len(suite.allContainers)-1]
		m    = newMockCachedContainer(dest)
	)

	m.parentID = last.id
	m.expectedPath = stdpath.Join(last.expectedPath, m.displayName)
	m.expectedLocation = stdpath.Join(last.expectedPath, m.displayName)

	require.Empty(t, suite.fc.DestinationNameToID(dest), "destination not yet added to cache")
	require.NoError(t, suite.fc.AddToCache(ctx, m, false))
	require.Empty(t, suite.fc.DestinationNameToID(dest),
		"destination id from cache, still empty, because this is not a calendar")

	p, l, err := suite.fc.IDToPath(ctx, m.id, false)
	require.NoError(t, err)
	assert.Equal(t, m.expectedPath, p.String())
	assert.Equal(t, m.expectedLocation, l.String())
}

// ---------------------------------------------------------------------------
// integration suite
// ---------------------------------------------------------------------------

type FolderCacheIntegrationSuite struct {
	suite.Suite
	credentials account.M365Config
	gs          graph.Servicer
}

func TestFolderCacheIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoConnectorExchangeFolderCacheTests)

	suite.Run(t, new(FolderCacheIntegrationSuite))
}

func (suite *FolderCacheIntegrationSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365

	adpt, err := graph.CreateAdapter(
		m365.AzureTenantID,
		m365.AzureClientID,
		m365.AzureClientSecret)
	require.NoError(t, err)

	suite.gs = graph.NewService(adpt)

	require.NoError(suite.T(), err)
}

// Testing to ensure that cache system works for in multiple different environments
func (suite *FolderCacheIntegrationSuite) TestCreateContainerDestination() {
	ctx, flush := tester.NewContext()
	defer flush()

	a := tester.NewM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	var (
		user            = tester.M365UserID(suite.T())
		directoryCaches = make(map[path.CategoryType]graph.ContainerResolver)
		folderName      = tester.DefaultTestRestoreDestination().ContainerName
		tests           = []struct {
			name         string
			pathFunc1    func(t *testing.T) path.Path
			pathFunc2    func(t *testing.T) path.Path
			category     path.CategoryType
			folderPrefix string
			useIDForPath bool
		}{
			{
				name:     "Mail Cache Test",
				category: path.EmailCategory,
				pathFunc1: func(t *testing.T) path.Path {
					pth, err := path.Builder{}.
						Append("Griffindor").
						Append("Croix").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EmailCategory,
							false)
					require.NoError(t, err)

					return pth
				},
				pathFunc2: func(t *testing.T) path.Path {
					pth, err := path.Builder{}.
						Append("Griffindor").
						Append("Felicius").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EmailCategory,
							false)
					require.NoError(t, err)

					return pth
				},
			},
			{
				name:     "Contact Cache Test",
				category: path.ContactsCategory,
				pathFunc1: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.
						Append("HufflePuff").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false)
					require.NoError(t, err)

					return aPath
				},
				pathFunc2: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.
						Append("Ravenclaw").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false)
					require.NoError(t, err)

					return aPath
				},
			},
			{
				name:         "Event Cache Test",
				category:     path.EventsCategory,
				useIDForPath: true,
				pathFunc1: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.
						Append("Durmstrang").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false)
					require.NoError(t, err)

					return aPath
				},
				pathFunc2: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.
						Append("Beauxbatons").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false)
					require.NoError(t, err)

					return aPath
				},
			},
		}
	)

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			folderID, err := CreateContainerDestination(
				ctx,
				m365,
				test.pathFunc1(t),
				folderName,
				directoryCaches,
				fault.New(true))
			require.NoError(t, err)

			resolver := directoryCaches[test.category]

			_, _, err = resolver.IDToPath(ctx, folderID, test.useIDForPath)
			assert.NoError(t, err)

			parentContainer := folderName
			if test.useIDForPath {
				parentContainer = folderID
			}

			secondID, err := CreateContainerDestination(
				ctx,
				m365,
				test.pathFunc2(t),
				parentContainer,
				directoryCaches,
				fault.New(true))
			require.NoError(t, err)

			_, _, err = resolver.IDToPath(ctx, secondID, test.useIDForPath)
			require.NoError(t, err)

			p := stdpath.Join(test.folderPrefix, parentContainer)
			_, ok := resolver.PathInCache(p)
			require.True(t, ok, "looking for path in cache: %s", p)
		})
	}
}
