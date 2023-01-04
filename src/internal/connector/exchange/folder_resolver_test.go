package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type CacheResolverSuite struct {
	suite.Suite
	gs graph.Servicer
}

func TestCacheResolverIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(CacheResolverSuite))
}

func (suite *CacheResolverSuite) SetupSuite() {
	t := suite.T()

	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365)
	require.NoError(t, err)

	suite.gs = service
}

func (suite *CacheResolverSuite) TestPopulate() {
	ctx, flush := tester.NewContext()
	defer flush()

	eventFunc := func(t *testing.T) graph.ContainerResolver {
		return &eventCalendarCache{
			userID: tester.M365UserID(t),
			gs:     suite.gs,
		}
	}

	contactFunc := func(t *testing.T) graph.ContainerResolver {
		return &contactFolderCache{
			userID: tester.M365UserID(t),
			gs:     suite.gs,
		}
	}

	tests := []struct {
		name, folderName, root, basePath string
		resolverFunc                     func(t *testing.T) graph.ContainerResolver
		canFind                          assert.BoolAssertionFunc
	}{
		{
			name:         "Default Event Cache",
			folderName:   DefaultCalendar,
			root:         DefaultCalendar,
			basePath:     DefaultCalendar,
			resolverFunc: eventFunc,
			canFind:      assert.True,
		},
		{
			name:         "Default Event Folder Hidden",
			root:         DefaultCalendar,
			folderName:   DefaultContactFolder,
			canFind:      assert.False,
			resolverFunc: eventFunc,
		},
		{
			name:         "Name Not in Cache",
			folderName:   "testFooBarWhoBar",
			root:         DefaultCalendar,
			canFind:      assert.False,
			resolverFunc: eventFunc,
		},
		{
			name:         "Default Contact Cache",
			folderName:   DefaultContactFolder,
			root:         DefaultContactFolder,
			basePath:     DefaultContactFolder,
			canFind:      assert.True,
			resolverFunc: contactFunc,
		},
		{
			name:         "Default Contact Hidden",
			folderName:   DefaultContactFolder,
			root:         DefaultContactFolder,
			canFind:      assert.False,
			resolverFunc: contactFunc,
		},
		{
			name:         "Name Not in Cache",
			folderName:   "testFooBarWhoBar",
			root:         DefaultContactFolder,
			canFind:      assert.False,
			resolverFunc: contactFunc,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			resolver := test.resolverFunc(t)

			require.NoError(t, resolver.Populate(ctx, test.root, test.basePath))
			_, isFound := resolver.PathInCache(test.folderName)
			test.canFind(t, isFound)
		})
	}
}
